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
	rbac "gitlab.com/rioos/magudi/pkg/apis/rbac"
)

// FakeRoleBindings implements RoleBindingInterface
type FakeRoleBindings struct {
	Fake *FakeRbac
	ns   string
}

var rolebindingsResource = schema.GroupVersionResource{Group: "rbac.authorization.k8s.io", Version: "", Resource: "rolebindings"}

var rolebindingsKind = schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "", Kind: "RoleBinding"}

func (c *FakeRoleBindings) Create(roleBinding *rbac.RoleBinding) (result *rbac.RoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rolebindingsResource, c.ns, roleBinding), &rbac.RoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rbac.RoleBinding), err
}

func (c *FakeRoleBindings) Update(roleBinding *rbac.RoleBinding) (result *rbac.RoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rolebindingsResource, c.ns, roleBinding), &rbac.RoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rbac.RoleBinding), err
}

func (c *FakeRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rolebindingsResource, c.ns, name), &rbac.RoleBinding{})

	return err
}

func (c *FakeRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rolebindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &rbac.RoleBindingList{})
	return err
}

func (c *FakeRoleBindings) Get(name string, options v1.GetOptions) (result *rbac.RoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rolebindingsResource, c.ns, name), &rbac.RoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rbac.RoleBinding), err
}

func (c *FakeRoleBindings) List(opts v1.ListOptions) (result *rbac.RoleBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rolebindingsResource, rolebindingsKind, c.ns, opts), &rbac.RoleBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rbac.RoleBindingList{}
	for _, item := range obj.(*rbac.RoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested roleBindings.
func (c *FakeRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rolebindingsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched roleBinding.
func (c *FakeRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *rbac.RoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rolebindingsResource, c.ns, name, data, subresources...), &rbac.RoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rbac.RoleBinding), err
}
