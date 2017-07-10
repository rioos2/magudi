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

package v1

import (
	autoscaling_v1 "k8s.io/api/autoscaling/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	clientset "gitlab.com/rioos/magudi/pkg/client/clientset_generated/clientset"
	internalinterfaces "gitlab.com/rioos/magudi/pkg/client/informers/informers_generated/externalversions/internalinterfaces"
	v1 "gitlab.com/rioos/magudi/pkg/client/listers/autoscaling/v1"
	time "time"
)

// HorizontalPodAutoscalerInformer provides access to a shared informer and lister for
// HorizontalPodAutoscalers.
type HorizontalPodAutoscalerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.HorizontalPodAutoscalerLister
}

type horizontalPodAutoscalerInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newHorizontalPodAutoscalerInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.AutoscalingV1().HorizontalPodAutoscalers(meta_v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.AutoscalingV1().HorizontalPodAutoscalers(meta_v1.NamespaceAll).Watch(options)
			},
		},
		&autoscaling_v1.HorizontalPodAutoscaler{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *horizontalPodAutoscalerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&autoscaling_v1.HorizontalPodAutoscaler{}, newHorizontalPodAutoscalerInformer)
}

func (f *horizontalPodAutoscalerInformer) Lister() v1.HorizontalPodAutoscalerLister {
	return v1.NewHorizontalPodAutoscalerLister(f.Informer().GetIndexer())
}
