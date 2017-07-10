/*
Copyright 2015 The Kubernetes Authors.

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

package storage

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"gitlab.com/rioos/magudi/pkg/api"
	"gitlab.com/rioos/magudi/pkg/registry/cachesize"
	"gitlab.com/rioos/magudi/pkg/registry/core/secret"
)

type REST struct {
	*genericregistry.Store
}

// NewREST returns a RESTStorage object that will work against secrets.
func NewREST(optsGetter generic.RESTOptionsGetter) *REST {
	store := &genericregistry.Store{
		Copier:            api.Scheme,
		NewFunc:           func() runtime.Object { return &api.Secret{} },
		NewListFunc:       func() runtime.Object { return &api.SecretList{} },
		PredicateFunc:     secret.Matcher,
		QualifiedResource: api.Resource("secrets"),
		WatchCacheSize:    cachesize.GetWatchCacheSizeByResource("secrets"),

		CreateStrategy: secret.Strategy,
		UpdateStrategy: secret.Strategy,
		DeleteStrategy: secret.Strategy,
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: secret.GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err) // TODO: Propagate error up
	}
	return &REST{store}
}
