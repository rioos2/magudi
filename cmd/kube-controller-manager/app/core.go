/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package app implements a server that runs a set of active
// components.  This includes replication controllers, service endpoints and
// nodes.
//
package app

import (
	"fmt"
	"net"

	"github.com/golang/glog"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"gitlab.com/sankish/magudi/pkg/api"
	"gitlab.com/sankish/magudi/pkg/client/clientset_generated/clientset"
	endpointcontroller "gitlab.com/sankish/magudi/pkg/controller/endpoint"
	"gitlab.com/sankish/magudi/pkg/controller/garbagecollector"
	"gitlab.com/sankish/magudi/pkg/controller/garbagecollector/metaonly"
	namespacecontroller "gitlab.com/sankish/magudi/pkg/controller/namespace"
	"gitlab.com/sankish/magudi/pkg/controller/podgc"
	replicationcontroller "gitlab.com/sankish/magudi/pkg/controller/replication"
	routecontroller "gitlab.com/sankish/magudi/pkg/controller/route"
	servicecontroller "gitlab.com/sankish/magudi/pkg/controller/service"
	serviceaccountcontroller "gitlab.com/sankish/magudi/pkg/controller/serviceaccount"
	ttlcontroller "gitlab.com/sankish/magudi/pkg/controller/ttl"
)

func startServiceController(ctx ControllerContext) (bool, error) {
	serviceController, err := servicecontroller.New(
		ctx.Cloud,
		ctx.ClientBuilder.ClientOrDie("service-controller"),
		ctx.InformerFactory.Core().V1().Services(),
		ctx.InformerFactory.Core().V1().Nodes(),
		ctx.Options.ClusterName,
	)
	if err != nil {
		glog.Errorf("Failed to start service controller: %v", err)
		return false, nil
	}
	go serviceController.Run(ctx.Stop, int(ctx.Options.ConcurrentServiceSyncs))
	return true, nil
}

func startRouteController(ctx ControllerContext) (bool, error) {
	_, clusterCIDR, err := net.ParseCIDR(ctx.Options.ClusterCIDR)
	if err != nil {
		glog.Warningf("Unsuccessful parsing of cluster CIDR %v: %v", ctx.Options.ClusterCIDR, err)
	}
	// TODO demorgans
	if ctx.Options.AllocateNodeCIDRs && ctx.Options.ConfigureCloudRoutes {
		if ctx.Cloud == nil {
			glog.Warning("configure-cloud-routes is set, but no cloud provider specified. Will not configure cloud provider routes.")
			return false, nil
		} else if routes, ok := ctx.Cloud.Routes(); !ok {
			glog.Warning("configure-cloud-routes is set, but cloud provider does not support routes. Will not configure cloud provider routes.")
			return false, nil
		} else {
			routeController := routecontroller.New(routes, ctx.ClientBuilder.ClientOrDie("route-controller"), ctx.InformerFactory.Core().V1().Nodes(), ctx.Options.ClusterName, clusterCIDR)
			go routeController.Run(ctx.Stop, ctx.Options.RouteReconciliationPeriod.Duration)
			return true, nil
		}
	} else {
		glog.Infof("Will not configure cloud provider routes for allocate-node-cidrs: %v, configure-cloud-routes: %v.", ctx.Options.AllocateNodeCIDRs, ctx.Options.ConfigureCloudRoutes)
		return false, nil
	}
}


func startEndpointController(ctx ControllerContext) (bool, error) {
	go endpointcontroller.NewEndpointController(
		ctx.InformerFactory.Core().V1().Pods(),
		ctx.InformerFactory.Core().V1().Services(),
		ctx.InformerFactory.Core().V1().Endpoints(),
		ctx.ClientBuilder.ClientOrDie("endpoint-controller"),
	).Run(int(ctx.Options.ConcurrentEndpointSyncs), ctx.Stop)
	return true, nil
}

func startReplicationController(ctx ControllerContext) (bool, error) {
	go replicationcontroller.NewReplicationManager(
		ctx.InformerFactory.Core().V1().Pods(),
		ctx.InformerFactory.Core().V1().ReplicationControllers(),
		ctx.ClientBuilder.ClientOrDie("replication-controller"),
		replicationcontroller.BurstReplicas,
	).Run(int(ctx.Options.ConcurrentRCSyncs), ctx.Stop)
	return true, nil
}

func startPodGCController(ctx ControllerContext) (bool, error) {
	go podgc.NewPodGC(
		ctx.ClientBuilder.ClientOrDie("pod-garbage-collector"),
		ctx.InformerFactory.Core().V1().Pods(),
		int(ctx.Options.TerminatedPodGCThreshold),
	).Run(ctx.Stop)
	return true, nil
}


func startNamespaceController(ctx ControllerContext) (bool, error) {
	// TODO: should use a dynamic RESTMapper built from the discovery results.
	restMapper := api.Registry.RESTMapper()

	// the namespace cleanup controller is very chatty.  It makes lots of discovery calls and then it makes lots of delete calls
	// the ratelimiter negatively affects its speed.  Deleting 100 total items in a namespace (that's only a few of each resource
	// including events), takes ~10 seconds by default.
	nsKubeconfig := ctx.ClientBuilder.ConfigOrDie("namespace-controller")
	nsKubeconfig.QPS *= 10
	nsKubeconfig.Burst *= 10
	namespaceKubeClient := clientset.NewForConfigOrDie(nsKubeconfig)
	namespaceClientPool := dynamic.NewClientPool(nsKubeconfig, restMapper, dynamic.LegacyAPIPathResolverFunc)

	discoverResourcesFn := namespaceKubeClient.Discovery().ServerPreferredNamespacedResources

	namespaceController := namespacecontroller.NewNamespaceController(
		namespaceKubeClient,
		namespaceClientPool,
		discoverResourcesFn,
		ctx.InformerFactory.Core().V1().Namespaces(),
		ctx.Options.NamespaceSyncPeriod.Duration,
		v1.FinalizerKubernetes,
	)
	go namespaceController.Run(int(ctx.Options.ConcurrentNamespaceSyncs), ctx.Stop)

	return true, nil
}

func startServiceAccountController(ctx ControllerContext) (bool, error) {
	go serviceaccountcontroller.NewServiceAccountsController(
		ctx.InformerFactory.Core().V1().ServiceAccounts(),
		ctx.InformerFactory.Core().V1().Namespaces(),
		ctx.ClientBuilder.ClientOrDie("service-account-controller"),
		serviceaccountcontroller.DefaultServiceAccountsControllerOptions(),
	).Run(1, ctx.Stop)
	return true, nil
}

func startTTLController(ctx ControllerContext) (bool, error) {
	go ttlcontroller.NewTTLController(
		ctx.InformerFactory.Core().V1().Nodes(),
		ctx.ClientBuilder.ClientOrDie("ttl-controller"),
	).Run(5, ctx.Stop)
	return true, nil
}

func startGarbageCollectorController(ctx ControllerContext) (bool, error) {
	if !ctx.Options.EnableGarbageCollector {
		return false, nil
	}

	// TODO: should use a dynamic RESTMapper built from the discovery results.
	restMapper := api.Registry.RESTMapper()

	gcClientset := ctx.ClientBuilder.ClientOrDie("generic-garbage-collector")
	preferredResources, err := gcClientset.Discovery().ServerPreferredResources()
	if err != nil {
		return true, fmt.Errorf("failed to get supported resources from server: %v", err)
	}
	deletableResources := discovery.FilteredBy(discovery.SupportsAllVerbs{Verbs: []string{"get", "list", "watch", "patch", "update", "delete"}}, preferredResources)
	deletableGroupVersionResources, err := discovery.GroupVersionResources(deletableResources)
	if err != nil {
		return true, fmt.Errorf("Failed to parse resources from server: %v", err)
	}

	config := ctx.ClientBuilder.ConfigOrDie("generic-garbage-collector")
	config.ContentConfig.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: metaonly.NewMetadataCodecFactory()}
	metaOnlyClientPool := dynamic.NewClientPool(config, restMapper, dynamic.LegacyAPIPathResolverFunc)
	config.ContentConfig = dynamic.ContentConfig()
	clientPool := dynamic.NewClientPool(config, restMapper, dynamic.LegacyAPIPathResolverFunc)

	ignoredResources := make(map[schema.GroupResource]struct{})
	for _, r := range ctx.Options.GCIgnoredResources {
		ignoredResources[schema.GroupResource{Group: r.Group, Resource: r.Resource}] = struct{}{}
	}

	garbageCollector, err := garbagecollector.NewGarbageCollector(
		metaOnlyClientPool,
		clientPool,
		restMapper,
		deletableGroupVersionResources,
		ignoredResources,
		ctx.InformerFactory,
	)
	if err != nil {
		return true, fmt.Errorf("Failed to start the generic garbage collector: %v", err)
	}
	workers := int(ctx.Options.ConcurrentGCSyncs)
	go garbageCollector.Run(workers, ctx.Stop)

	return true, nil
}
