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
	"fmt"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"gitlab.com/rioos/magudi/pkg/api"
	"gitlab.com/rioos/magudi/pkg/api/ref"
	k8s_api_v1 "gitlab.com/rioos/magudi/pkg/api/v1"
)

// The EventExpansion interface allows manually adding extra methods to the EventInterface.
type EventExpansion interface {
	// CreateWithEventNamespace is the same as a Create, except that it sends the request to the event.Namespace.
	CreateWithEventNamespace(event *api.Event) (*api.Event, error)
	// UpdateWithEventNamespace is the same as a Update, except that it sends the request to the event.Namespace.
	UpdateWithEventNamespace(event *api.Event) (*api.Event, error)
	PatchWithEventNamespace(event *api.Event, data []byte) (*api.Event, error)
	// Search finds events about the specified object
	Search(scheme *runtime.Scheme, objOrRef runtime.Object) (*api.EventList, error)
	// Returns the appropriate field selector based on the API version being used to communicate with the server.
	// The returned field selector can be used with List and Watch to filter desired events.
	GetFieldSelector(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID *string) fields.Selector
}

// CreateWithEventNamespace makes a new event. Returns the copy of the event the server returns,
// or an error. The namespace to create the event within is deduced from the
// event; it must either match this event client's namespace, or this event
// client must have been created with the "" namespace.
func (e *events) CreateWithEventNamespace(event *api.Event) (*api.Event, error) {
	if e.ns != "" && event.Namespace != e.ns {
		return nil, fmt.Errorf("can't create an event with namespace '%v' in namespace '%v'", event.Namespace, e.ns)
	}
	result := &api.Event{}
	err := e.client.Post().
		NamespaceIfScoped(event.Namespace, len(event.Namespace) > 0).
		Resource("events").
		Body(event).
		Do().
		Into(result)
	return result, err
}

// UpdateWithEventNamespace modifies an existing event. It returns the copy of the event that the server returns,
// or an error. The namespace and key to update the event within is deduced from the event. The
// namespace must either match this event client's namespace, or this event client must have been
// created with the "" namespace. Update also requires the ResourceVersion to be set in the event
// object.
func (e *events) UpdateWithEventNamespace(event *api.Event) (*api.Event, error) {
	result := &api.Event{}
	err := e.client.Put().
		NamespaceIfScoped(event.Namespace, len(event.Namespace) > 0).
		Resource("events").
		Name(event.Name).
		Body(event).
		Do().
		Into(result)
	return result, err
}

// PatchWithEventNamespace modifies an existing event. It returns the copy of
// the event that the server returns, or an error. The namespace and name of the
// target event is deduced from the incompleteEvent. The namespace must either
// match this event client's namespace, or this event client must have been
// created with the "" namespace.
func (e *events) PatchWithEventNamespace(incompleteEvent *api.Event, data []byte) (*api.Event, error) {
	if e.ns != "" && incompleteEvent.Namespace != e.ns {
		return nil, fmt.Errorf("can't patch an event with namespace '%v' in namespace '%v'", incompleteEvent.Namespace, e.ns)
	}
	result := &api.Event{}
	err := e.client.Patch(types.StrategicMergePatchType).
		NamespaceIfScoped(incompleteEvent.Namespace, len(incompleteEvent.Namespace) > 0).
		Resource("events").
		Name(incompleteEvent.Name).
		Body(data).
		Do().
		Into(result)
	return result, err
}

// Search finds events about the specified object. The namespace of the
// object must match this event's client namespace unless the event client
// was made with the "" namespace.
func (e *events) Search(scheme *runtime.Scheme, objOrRef runtime.Object) (*api.EventList, error) {
	ref, err := ref.GetReference(scheme, objOrRef)
	if err != nil {
		return nil, err
	}
	if e.ns != "" && ref.Namespace != e.ns {
		return nil, fmt.Errorf("won't be able to find any events of namespace '%v' in namespace '%v'", ref.Namespace, e.ns)
	}
	stringRefKind := string(ref.Kind)
	var refKind *string
	if stringRefKind != "" {
		refKind = &stringRefKind
	}
	stringRefUID := string(ref.UID)
	var refUID *string
	if stringRefUID != "" {
		refUID = &stringRefUID
	}
	fieldSelector := e.GetFieldSelector(&ref.Name, &ref.Namespace, refKind, refUID)
	return e.List(metav1.ListOptions{FieldSelector: fieldSelector.String()})
}

// Returns the appropriate field selector based on the API version being used to communicate with the server.
// The returned field selector can be used with List and Watch to filter desired events.
func (e *events) GetFieldSelector(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID *string) fields.Selector {
	apiVersion := e.client.APIVersion().String()
	field := fields.Set{}
	if involvedObjectName != nil {
		field[GetInvolvedObjectNameFieldLabel(apiVersion)] = *involvedObjectName
	}
	if involvedObjectNamespace != nil {
		field["involvedObject.namespace"] = *involvedObjectNamespace
	}
	if involvedObjectKind != nil {
		field["involvedObject.kind"] = *involvedObjectKind
	}
	if involvedObjectUID != nil {
		field["involvedObject.uid"] = *involvedObjectUID
	}
	return field.AsSelector()
}

// Returns the appropriate field label to use for name of the involved object as per the given API version.
func GetInvolvedObjectNameFieldLabel(version string) string {
	return "involvedObject.name"
}

// TODO: This is a temporary arrangement and will be removed once all clients are moved to use the clientset.
type EventSinkImpl struct {
	Interface EventInterface
}

func (e *EventSinkImpl) Create(event *v1.Event) (*v1.Event, error) {
	internalEvent := &api.Event{}
	err := k8s_api_v1.Convert_v1_Event_To_api_Event(event, internalEvent, nil)
	if err != nil {
		return nil, err
	}
	_, err = e.Interface.CreateWithEventNamespace(internalEvent)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (e *EventSinkImpl) Update(event *v1.Event) (*v1.Event, error) {
	internalEvent := &api.Event{}
	err := k8s_api_v1.Convert_v1_Event_To_api_Event(event, internalEvent, nil)
	if err != nil {
		return nil, err
	}
	_, err = e.Interface.UpdateWithEventNamespace(internalEvent)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (e *EventSinkImpl) Patch(event *v1.Event, data []byte) (*v1.Event, error) {
	internalEvent := &api.Event{}
	err := k8s_api_v1.Convert_v1_Event_To_api_Event(event, internalEvent, nil)
	if err != nil {
		return nil, err
	}
	internalEvent, err = e.Interface.PatchWithEventNamespace(internalEvent, data)
	if err != nil {
		return nil, err
	}
	externalEvent := &v1.Event{}
	err = k8s_api_v1.Convert_api_Event_To_v1_Event(internalEvent, externalEvent, nil)
	if err != nil {
		// Patch succeeded, no need to report the failed conversion
		return event, nil
	}
	return externalEvent, nil
}
