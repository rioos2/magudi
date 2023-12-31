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

// FakePods implements PodInterface
type FakePods struct {
	Fake *FakeCore
	ns   string
}

var podsResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "pods"}

var podsKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "Pod"}

func (c *FakePods) Create(pod *api.Pod) (result *api.Pod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podsResource, c.ns, pod), &api.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.Pod), err
}

func (c *FakePods) Update(pod *api.Pod) (result *api.Pod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podsResource, c.ns, pod), &api.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.Pod), err
}

func (c *FakePods) UpdateStatus(pod *api.Pod) (*api.Pod, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(podsResource, "status", c.ns, pod), &api.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.Pod), err
}

func (c *FakePods) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(podsResource, c.ns, name), &api.Pod{})

	return err
}

func (c *FakePods) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &api.PodList{})
	return err
}

func (c *FakePods) Get(name string, options v1.GetOptions) (result *api.Pod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podsResource, c.ns, name), &api.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.Pod), err
}

func (c *FakePods) List(opts v1.ListOptions) (result *api.PodList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podsResource, podsKind, c.ns, opts), &api.PodList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &api.PodList{}
	for _, item := range obj.(*api.PodList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pods.
func (c *FakePods) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched pod.
func (c *FakePods) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *api.Pod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podsResource, c.ns, name, data, subresources...), &api.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.Pod), err
}
