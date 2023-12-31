/*
Copyright 2016 The Kubernetes Authors.

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
	"k8s.io/apimachinery/pkg/types"
	"gitlab.com/rioos/magudi/pkg/api"
)

// The NodeExpansion interface allows manually adding extra methods to the NodeInterface.
type NodeExpansion interface {
	// PatchStatus modifies the status of an existing node. It returns the copy
	// of the node that the server returns, or an error.
	PatchStatus(nodeName string, data []byte) (*api.Node, error)
}

// PatchStatus modifies the status of an existing node. It returns the copy of
// the node that the server returns, or an error.
func (c *nodes) PatchStatus(nodeName string, data []byte) (*api.Node, error) {
	result := &api.Node{}
	err := c.client.Patch(types.StrategicMergePatchType).
		Resource("nodes").
		Name(nodeName).
		SubResource("status").
		Body(data).
		Do().
		Into(result)
	return result, err
}
