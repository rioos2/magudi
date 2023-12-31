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

package internalversion

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	api "gitlab.com/rioos/magudi/pkg/api"
	scheme "gitlab.com/rioos/magudi/pkg/client/clientset_generated/internalclientset/scheme"
)

// PersistentVolumeClaimsGetter has a method to return a PersistentVolumeClaimInterface.
// A group's client should implement this interface.
type PersistentVolumeClaimsGetter interface {
	PersistentVolumeClaims(namespace string) PersistentVolumeClaimInterface
}

// PersistentVolumeClaimInterface has methods to work with PersistentVolumeClaim resources.
type PersistentVolumeClaimInterface interface {
	Create(*api.PersistentVolumeClaim) (*api.PersistentVolumeClaim, error)
	Update(*api.PersistentVolumeClaim) (*api.PersistentVolumeClaim, error)
	UpdateStatus(*api.PersistentVolumeClaim) (*api.PersistentVolumeClaim, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*api.PersistentVolumeClaim, error)
	List(opts v1.ListOptions) (*api.PersistentVolumeClaimList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *api.PersistentVolumeClaim, err error)
	PersistentVolumeClaimExpansion
}

// persistentVolumeClaims implements PersistentVolumeClaimInterface
type persistentVolumeClaims struct {
	client rest.Interface
	ns     string
}

// newPersistentVolumeClaims returns a PersistentVolumeClaims
func newPersistentVolumeClaims(c *CoreClient, namespace string) *persistentVolumeClaims {
	return &persistentVolumeClaims{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a persistentVolumeClaim and creates it.  Returns the server's representation of the persistentVolumeClaim, and an error, if there is any.
func (c *persistentVolumeClaims) Create(persistentVolumeClaim *api.PersistentVolumeClaim) (result *api.PersistentVolumeClaim, err error) {
	result = &api.PersistentVolumeClaim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		Body(persistentVolumeClaim).
		Do().
		Into(result)
	return
}

// Update takes the representation of a persistentVolumeClaim and updates it. Returns the server's representation of the persistentVolumeClaim, and an error, if there is any.
func (c *persistentVolumeClaims) Update(persistentVolumeClaim *api.PersistentVolumeClaim) (result *api.PersistentVolumeClaim, err error) {
	result = &api.PersistentVolumeClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		Name(persistentVolumeClaim.Name).
		Body(persistentVolumeClaim).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclientstatus=false comment above the type to avoid generating UpdateStatus().

func (c *persistentVolumeClaims) UpdateStatus(persistentVolumeClaim *api.PersistentVolumeClaim) (result *api.PersistentVolumeClaim, err error) {
	result = &api.PersistentVolumeClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		Name(persistentVolumeClaim.Name).
		SubResource("status").
		Body(persistentVolumeClaim).
		Do().
		Into(result)
	return
}

// Delete takes name of the persistentVolumeClaim and deletes it. Returns an error if one occurs.
func (c *persistentVolumeClaims) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *persistentVolumeClaims) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the persistentVolumeClaim, and returns the corresponding persistentVolumeClaim object, and an error if there is any.
func (c *persistentVolumeClaims) Get(name string, options v1.GetOptions) (result *api.PersistentVolumeClaim, err error) {
	result = &api.PersistentVolumeClaim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PersistentVolumeClaims that match those selectors.
func (c *persistentVolumeClaims) List(opts v1.ListOptions) (result *api.PersistentVolumeClaimList, err error) {
	result = &api.PersistentVolumeClaimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested persistentVolumeClaims.
func (c *persistentVolumeClaims) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched persistentVolumeClaim.
func (c *persistentVolumeClaims) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *api.PersistentVolumeClaim, err error) {
	result = &api.PersistentVolumeClaim{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
