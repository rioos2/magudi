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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	federation "gitlab.com/rioos/magudi/federation/apis/federation"
	api "gitlab.com/rioos/magudi/pkg/api"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_Cluster_To_federation_Cluster,
		Convert_federation_Cluster_To_v1beta1_Cluster,
		Convert_v1beta1_ClusterCondition_To_federation_ClusterCondition,
		Convert_federation_ClusterCondition_To_v1beta1_ClusterCondition,
		Convert_v1beta1_ClusterList_To_federation_ClusterList,
		Convert_federation_ClusterList_To_v1beta1_ClusterList,
		Convert_v1beta1_ClusterSpec_To_federation_ClusterSpec,
		Convert_federation_ClusterSpec_To_v1beta1_ClusterSpec,
		Convert_v1beta1_ClusterStatus_To_federation_ClusterStatus,
		Convert_federation_ClusterStatus_To_v1beta1_ClusterStatus,
		Convert_v1beta1_ServerAddressByClientCIDR_To_federation_ServerAddressByClientCIDR,
		Convert_federation_ServerAddressByClientCIDR_To_v1beta1_ServerAddressByClientCIDR,
	)
}

func autoConvert_v1beta1_Cluster_To_federation_Cluster(in *Cluster, out *federation.Cluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ClusterSpec_To_federation_ClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ClusterStatus_To_federation_ClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Cluster_To_federation_Cluster is an autogenerated conversion function.
func Convert_v1beta1_Cluster_To_federation_Cluster(in *Cluster, out *federation.Cluster, s conversion.Scope) error {
	return autoConvert_v1beta1_Cluster_To_federation_Cluster(in, out, s)
}

func autoConvert_federation_Cluster_To_v1beta1_Cluster(in *federation.Cluster, out *Cluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_federation_ClusterSpec_To_v1beta1_ClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_federation_ClusterStatus_To_v1beta1_ClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_federation_Cluster_To_v1beta1_Cluster is an autogenerated conversion function.
func Convert_federation_Cluster_To_v1beta1_Cluster(in *federation.Cluster, out *Cluster, s conversion.Scope) error {
	return autoConvert_federation_Cluster_To_v1beta1_Cluster(in, out, s)
}

func autoConvert_v1beta1_ClusterCondition_To_federation_ClusterCondition(in *ClusterCondition, out *federation.ClusterCondition, s conversion.Scope) error {
	out.Type = federation.ClusterConditionType(in.Type)
	out.Status = api.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1beta1_ClusterCondition_To_federation_ClusterCondition is an autogenerated conversion function.
func Convert_v1beta1_ClusterCondition_To_federation_ClusterCondition(in *ClusterCondition, out *federation.ClusterCondition, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterCondition_To_federation_ClusterCondition(in, out, s)
}

func autoConvert_federation_ClusterCondition_To_v1beta1_ClusterCondition(in *federation.ClusterCondition, out *ClusterCondition, s conversion.Scope) error {
	out.Type = ClusterConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_federation_ClusterCondition_To_v1beta1_ClusterCondition is an autogenerated conversion function.
func Convert_federation_ClusterCondition_To_v1beta1_ClusterCondition(in *federation.ClusterCondition, out *ClusterCondition, s conversion.Scope) error {
	return autoConvert_federation_ClusterCondition_To_v1beta1_ClusterCondition(in, out, s)
}

func autoConvert_v1beta1_ClusterList_To_federation_ClusterList(in *ClusterList, out *federation.ClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]federation.Cluster)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_ClusterList_To_federation_ClusterList is an autogenerated conversion function.
func Convert_v1beta1_ClusterList_To_federation_ClusterList(in *ClusterList, out *federation.ClusterList, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterList_To_federation_ClusterList(in, out, s)
}

func autoConvert_federation_ClusterList_To_v1beta1_ClusterList(in *federation.ClusterList, out *ClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]Cluster, 0)
	} else {
		out.Items = *(*[]Cluster)(unsafe.Pointer(&in.Items))
	}
	return nil
}

// Convert_federation_ClusterList_To_v1beta1_ClusterList is an autogenerated conversion function.
func Convert_federation_ClusterList_To_v1beta1_ClusterList(in *federation.ClusterList, out *ClusterList, s conversion.Scope) error {
	return autoConvert_federation_ClusterList_To_v1beta1_ClusterList(in, out, s)
}

func autoConvert_v1beta1_ClusterSpec_To_federation_ClusterSpec(in *ClusterSpec, out *federation.ClusterSpec, s conversion.Scope) error {
	out.ServerAddressByClientCIDRs = *(*[]federation.ServerAddressByClientCIDR)(unsafe.Pointer(&in.ServerAddressByClientCIDRs))
	out.SecretRef = (*api.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
	return nil
}

// Convert_v1beta1_ClusterSpec_To_federation_ClusterSpec is an autogenerated conversion function.
func Convert_v1beta1_ClusterSpec_To_federation_ClusterSpec(in *ClusterSpec, out *federation.ClusterSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterSpec_To_federation_ClusterSpec(in, out, s)
}

func autoConvert_federation_ClusterSpec_To_v1beta1_ClusterSpec(in *federation.ClusterSpec, out *ClusterSpec, s conversion.Scope) error {
	if in.ServerAddressByClientCIDRs == nil {
		out.ServerAddressByClientCIDRs = make([]ServerAddressByClientCIDR, 0)
	} else {
		out.ServerAddressByClientCIDRs = *(*[]ServerAddressByClientCIDR)(unsafe.Pointer(&in.ServerAddressByClientCIDRs))
	}
	out.SecretRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
	return nil
}

// Convert_federation_ClusterSpec_To_v1beta1_ClusterSpec is an autogenerated conversion function.
func Convert_federation_ClusterSpec_To_v1beta1_ClusterSpec(in *federation.ClusterSpec, out *ClusterSpec, s conversion.Scope) error {
	return autoConvert_federation_ClusterSpec_To_v1beta1_ClusterSpec(in, out, s)
}

func autoConvert_v1beta1_ClusterStatus_To_federation_ClusterStatus(in *ClusterStatus, out *federation.ClusterStatus, s conversion.Scope) error {
	out.Conditions = *(*[]federation.ClusterCondition)(unsafe.Pointer(&in.Conditions))
	out.Zones = *(*[]string)(unsafe.Pointer(&in.Zones))
	out.Region = in.Region
	return nil
}

// Convert_v1beta1_ClusterStatus_To_federation_ClusterStatus is an autogenerated conversion function.
func Convert_v1beta1_ClusterStatus_To_federation_ClusterStatus(in *ClusterStatus, out *federation.ClusterStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterStatus_To_federation_ClusterStatus(in, out, s)
}

func autoConvert_federation_ClusterStatus_To_v1beta1_ClusterStatus(in *federation.ClusterStatus, out *ClusterStatus, s conversion.Scope) error {
	out.Conditions = *(*[]ClusterCondition)(unsafe.Pointer(&in.Conditions))
	out.Zones = *(*[]string)(unsafe.Pointer(&in.Zones))
	out.Region = in.Region
	return nil
}

// Convert_federation_ClusterStatus_To_v1beta1_ClusterStatus is an autogenerated conversion function.
func Convert_federation_ClusterStatus_To_v1beta1_ClusterStatus(in *federation.ClusterStatus, out *ClusterStatus, s conversion.Scope) error {
	return autoConvert_federation_ClusterStatus_To_v1beta1_ClusterStatus(in, out, s)
}

func autoConvert_v1beta1_ServerAddressByClientCIDR_To_federation_ServerAddressByClientCIDR(in *ServerAddressByClientCIDR, out *federation.ServerAddressByClientCIDR, s conversion.Scope) error {
	out.ClientCIDR = in.ClientCIDR
	out.ServerAddress = in.ServerAddress
	return nil
}

// Convert_v1beta1_ServerAddressByClientCIDR_To_federation_ServerAddressByClientCIDR is an autogenerated conversion function.
func Convert_v1beta1_ServerAddressByClientCIDR_To_federation_ServerAddressByClientCIDR(in *ServerAddressByClientCIDR, out *federation.ServerAddressByClientCIDR, s conversion.Scope) error {
	return autoConvert_v1beta1_ServerAddressByClientCIDR_To_federation_ServerAddressByClientCIDR(in, out, s)
}

func autoConvert_federation_ServerAddressByClientCIDR_To_v1beta1_ServerAddressByClientCIDR(in *federation.ServerAddressByClientCIDR, out *ServerAddressByClientCIDR, s conversion.Scope) error {
	out.ClientCIDR = in.ClientCIDR
	out.ServerAddress = in.ServerAddress
	return nil
}

// Convert_federation_ServerAddressByClientCIDR_To_v1beta1_ServerAddressByClientCIDR is an autogenerated conversion function.
func Convert_federation_ServerAddressByClientCIDR_To_v1beta1_ServerAddressByClientCIDR(in *federation.ServerAddressByClientCIDR, out *ServerAddressByClientCIDR, s conversion.Scope) error {
	return autoConvert_federation_ServerAddressByClientCIDR_To_v1beta1_ServerAddressByClientCIDR(in, out, s)
}
