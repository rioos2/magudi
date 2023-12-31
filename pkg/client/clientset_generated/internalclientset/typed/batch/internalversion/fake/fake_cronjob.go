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
	batch "gitlab.com/rioos/magudi/pkg/apis/batch"
)

// FakeCronJobs implements CronJobInterface
type FakeCronJobs struct {
	Fake *FakeBatch
	ns   string
}

var cronjobsResource = schema.GroupVersionResource{Group: "batch", Version: "", Resource: "cronjobs"}

var cronjobsKind = schema.GroupVersionKind{Group: "batch", Version: "", Kind: "CronJob"}

func (c *FakeCronJobs) Create(cronJob *batch.CronJob) (result *batch.CronJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cronjobsResource, c.ns, cronJob), &batch.CronJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.CronJob), err
}

func (c *FakeCronJobs) Update(cronJob *batch.CronJob) (result *batch.CronJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cronjobsResource, c.ns, cronJob), &batch.CronJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.CronJob), err
}

func (c *FakeCronJobs) UpdateStatus(cronJob *batch.CronJob) (*batch.CronJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cronjobsResource, "status", c.ns, cronJob), &batch.CronJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.CronJob), err
}

func (c *FakeCronJobs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cronjobsResource, c.ns, name), &batch.CronJob{})

	return err
}

func (c *FakeCronJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cronjobsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &batch.CronJobList{})
	return err
}

func (c *FakeCronJobs) Get(name string, options v1.GetOptions) (result *batch.CronJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cronjobsResource, c.ns, name), &batch.CronJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.CronJob), err
}

func (c *FakeCronJobs) List(opts v1.ListOptions) (result *batch.CronJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cronjobsResource, cronjobsKind, c.ns, opts), &batch.CronJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &batch.CronJobList{}
	for _, item := range obj.(*batch.CronJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cronJobs.
func (c *FakeCronJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cronjobsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched cronJob.
func (c *FakeCronJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.CronJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cronjobsResource, c.ns, name, data, subresources...), &batch.CronJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.CronJob), err
}
