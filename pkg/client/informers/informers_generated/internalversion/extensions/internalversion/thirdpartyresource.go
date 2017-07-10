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

package internalversion

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	extensions "gitlab.com/rioos/magudi/pkg/apis/extensions"
	internalclientset "gitlab.com/rioos/magudi/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "gitlab.com/rioos/magudi/pkg/client/informers/informers_generated/internalversion/internalinterfaces"
	internalversion "gitlab.com/rioos/magudi/pkg/client/listers/extensions/internalversion"
	time "time"
)

// ThirdPartyResourceInformer provides access to a shared informer and lister for
// ThirdPartyResources.
type ThirdPartyResourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ThirdPartyResourceLister
}

type thirdPartyResourceInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newThirdPartyResourceInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Extensions().ThirdPartyResources().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Extensions().ThirdPartyResources().Watch(options)
			},
		},
		&extensions.ThirdPartyResource{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *thirdPartyResourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&extensions.ThirdPartyResource{}, newThirdPartyResourceInformer)
}

func (f *thirdPartyResourceInformer) Lister() internalversion.ThirdPartyResourceLister {
	return internalversion.NewThirdPartyResourceLister(f.Informer().GetIndexer())
}
