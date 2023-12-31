/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v2alpha1

import (
	batch_v2alpha1 "k8s.io/api/batch/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	clientset "gitlab.com/rioos/magudi/pkg/client/clientset_generated/clientset"
	internalinterfaces "gitlab.com/rioos/magudi/pkg/client/informers/informers_generated/externalversions/internalinterfaces"
	v2alpha1 "gitlab.com/rioos/magudi/pkg/client/listers/batch/v2alpha1"
	time "time"
)

// CronJobInformer provides access to a shared informer and lister for
// CronJobs.
type CronJobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2alpha1.CronJobLister
}

type cronJobInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newCronJobInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.BatchV2alpha1().CronJobs(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.BatchV2alpha1().CronJobs(v1.NamespaceAll).Watch(options)
			},
		},
		&batch_v2alpha1.CronJob{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *cronJobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&batch_v2alpha1.CronJob{}, newCronJobInformer)
}

func (f *cronJobInformer) Lister() v2alpha1.CronJobLister {
	return v2alpha1.NewCronJobLister(f.Informer().GetIndexer())
}
