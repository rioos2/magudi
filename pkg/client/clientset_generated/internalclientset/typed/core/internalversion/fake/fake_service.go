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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	api "gitlab.com/rioos/magudi/pkg/api"
)

// FakeServices implements ServiceInterface
type FakeServices struct {
	Fake *FakeCore
	ns   string
}

var servicesResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "services"}

var servicesKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "Service"}

func (c *FakeServices) Create(service *api.Service) (result *api.Service, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicesResource, c.ns, service), &api.Service{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.Service), err
}

func (c *FakeServices) Update(service *api.Service) (result *api.Service, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicesResource, c.ns, service), &api.Service{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.Service), err
}

func (c *FakeServices) UpdateStatus(service *api.Service) (*api.Service, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicesResource, "status", c.ns, service), &api.Service{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.Service), err
}

func (c *FakeServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicesResource, c.ns, name), &api.Service{})

	return err
}

func (c *FakeServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &api.ServiceList{})
	return err
}

func (c *FakeServices) Get(name string, options v1.GetOptions) (result *api.Service, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicesResource, c.ns, name), &api.Service{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.Service), err
}

func (c *FakeServices) List(opts v1.ListOptions) (result *api.ServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicesResource, servicesKind, c.ns, opts), &api.ServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &api.ServiceList{}
	for _, item := range obj.(*api.ServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested services.
func (c *FakeServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicesResource, c.ns, opts))

}

// Patch applies the patch and returns the patched service.
func (c *FakeServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *api.Service, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicesResource, c.ns, name, data, subresources...), &api.Service{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.Service), err
}
