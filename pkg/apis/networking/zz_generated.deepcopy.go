// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package networking

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	api "gitlab.com/sankish/magudi/pkg/api"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_networking_NetworkPolicy, InType: reflect.TypeOf(&NetworkPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_networking_NetworkPolicyIngressRule, InType: reflect.TypeOf(&NetworkPolicyIngressRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_networking_NetworkPolicyList, InType: reflect.TypeOf(&NetworkPolicyList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_networking_NetworkPolicyPeer, InType: reflect.TypeOf(&NetworkPolicyPeer{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_networking_NetworkPolicyPort, InType: reflect.TypeOf(&NetworkPolicyPort{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_networking_NetworkPolicySpec, InType: reflect.TypeOf(&NetworkPolicySpec{})},
	)
}

// DeepCopy_networking_NetworkPolicy is an autogenerated deepcopy function.
func DeepCopy_networking_NetworkPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicy)
		out := out.(*NetworkPolicy)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_networking_NetworkPolicySpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_networking_NetworkPolicyIngressRule is an autogenerated deepcopy function.
func DeepCopy_networking_NetworkPolicyIngressRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicyIngressRule)
		out := out.(*NetworkPolicyIngressRule)
		*out = *in
		if in.Ports != nil {
			in, out := &in.Ports, &out.Ports
			*out = make([]NetworkPolicyPort, len(*in))
			for i := range *in {
				if err := DeepCopy_networking_NetworkPolicyPort(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.From != nil {
			in, out := &in.From, &out.From
			*out = make([]NetworkPolicyPeer, len(*in))
			for i := range *in {
				if err := DeepCopy_networking_NetworkPolicyPeer(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_networking_NetworkPolicyList is an autogenerated deepcopy function.
func DeepCopy_networking_NetworkPolicyList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicyList)
		out := out.(*NetworkPolicyList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]NetworkPolicy, len(*in))
			for i := range *in {
				if err := DeepCopy_networking_NetworkPolicy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_networking_NetworkPolicyPeer is an autogenerated deepcopy function.
func DeepCopy_networking_NetworkPolicyPeer(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicyPeer)
		out := out.(*NetworkPolicyPeer)
		*out = *in
		if in.PodSelector != nil {
			in, out := &in.PodSelector, &out.PodSelector
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.LabelSelector)
			}
		}
		if in.NamespaceSelector != nil {
			in, out := &in.NamespaceSelector, &out.NamespaceSelector
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.LabelSelector)
			}
		}
		return nil
	}
}

// DeepCopy_networking_NetworkPolicyPort is an autogenerated deepcopy function.
func DeepCopy_networking_NetworkPolicyPort(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicyPort)
		out := out.(*NetworkPolicyPort)
		*out = *in
		if in.Protocol != nil {
			in, out := &in.Protocol, &out.Protocol
			*out = new(api.Protocol)
			**out = **in
		}
		if in.Port != nil {
			in, out := &in.Port, &out.Port
			*out = new(intstr.IntOrString)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_networking_NetworkPolicySpec is an autogenerated deepcopy function.
func DeepCopy_networking_NetworkPolicySpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicySpec)
		out := out.(*NetworkPolicySpec)
		*out = *in
		if newVal, err := c.DeepCopy(&in.PodSelector); err != nil {
			return err
		} else {
			out.PodSelector = *newVal.(*v1.LabelSelector)
		}
		if in.Ingress != nil {
			in, out := &in.Ingress, &out.Ingress
			*out = make([]NetworkPolicyIngressRule, len(*in))
			for i := range *in {
				if err := DeepCopy_networking_NetworkPolicyIngressRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}
