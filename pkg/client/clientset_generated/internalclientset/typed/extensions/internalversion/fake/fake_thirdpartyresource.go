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
	extensions "gitlab.com/rioos/magudi/pkg/apis/extensions"
)

// FakeThirdPartyResources implements ThirdPartyResourceInterface
type FakeThirdPartyResources struct {
	Fake *FakeExtensions
}

var thirdpartyresourcesResource = schema.GroupVersionResource{Group: "extensions", Version: "", Resource: "thirdpartyresources"}

var thirdpartyresourcesKind = schema.GroupVersionKind{Group: "extensions", Version: "", Kind: "ThirdPartyResource"}

func (c *FakeThirdPartyResources) Create(thirdPartyResource *extensions.ThirdPartyResource) (result *extensions.ThirdPartyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(thirdpartyresourcesResource, thirdPartyResource), &extensions.ThirdPartyResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.ThirdPartyResource), err
}

func (c *FakeThirdPartyResources) Update(thirdPartyResource *extensions.ThirdPartyResource) (result *extensions.ThirdPartyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(thirdpartyresourcesResource, thirdPartyResource), &extensions.ThirdPartyResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.ThirdPartyResource), err
}

func (c *FakeThirdPartyResources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(thirdpartyresourcesResource, name), &extensions.ThirdPartyResource{})
	return err
}

func (c *FakeThirdPartyResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(thirdpartyresourcesResource, listOptions)

	_, err := c.Fake.Invokes(action, &extensions.ThirdPartyResourceList{})
	return err
}

func (c *FakeThirdPartyResources) Get(name string, options v1.GetOptions) (result *extensions.ThirdPartyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(thirdpartyresourcesResource, name), &extensions.ThirdPartyResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.ThirdPartyResource), err
}

func (c *FakeThirdPartyResources) List(opts v1.ListOptions) (result *extensions.ThirdPartyResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(thirdpartyresourcesResource, thirdpartyresourcesKind, opts), &extensions.ThirdPartyResourceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &extensions.ThirdPartyResourceList{}
	for _, item := range obj.(*extensions.ThirdPartyResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested thirdPartyResources.
func (c *FakeThirdPartyResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(thirdpartyresourcesResource, opts))
}

// Patch applies the patch and returns the patched thirdPartyResource.
func (c *FakeThirdPartyResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *extensions.ThirdPartyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(thirdpartyresourcesResource, name, data, subresources...), &extensions.ThirdPartyResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.ThirdPartyResource), err
}
