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

package v1

import (
	v1 "k8s.io/api/autoscaling/v1"
	core_v1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "gitlab.com/rioos/magudi/pkg/api"
	autoscaling "gitlab.com/rioos/magudi/pkg/apis/autoscaling"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference,
		Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference,
		Convert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler,
		Convert_autoscaling_HorizontalPodAutoscaler_To_v1_HorizontalPodAutoscaler,
		Convert_v1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition,
		Convert_autoscaling_HorizontalPodAutoscalerCondition_To_v1_HorizontalPodAutoscalerCondition,
		Convert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList,
		Convert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList,
		Convert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec,
		Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec,
		Convert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus,
		Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus,
		Convert_v1_MetricSpec_To_autoscaling_MetricSpec,
		Convert_autoscaling_MetricSpec_To_v1_MetricSpec,
		Convert_v1_MetricStatus_To_autoscaling_MetricStatus,
		Convert_autoscaling_MetricStatus_To_v1_MetricStatus,
		Convert_v1_ObjectMetricSource_To_autoscaling_ObjectMetricSource,
		Convert_autoscaling_ObjectMetricSource_To_v1_ObjectMetricSource,
		Convert_v1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus,
		Convert_autoscaling_ObjectMetricStatus_To_v1_ObjectMetricStatus,
		Convert_v1_PodsMetricSource_To_autoscaling_PodsMetricSource,
		Convert_autoscaling_PodsMetricSource_To_v1_PodsMetricSource,
		Convert_v1_PodsMetricStatus_To_autoscaling_PodsMetricStatus,
		Convert_autoscaling_PodsMetricStatus_To_v1_PodsMetricStatus,
		Convert_v1_ResourceMetricSource_To_autoscaling_ResourceMetricSource,
		Convert_autoscaling_ResourceMetricSource_To_v1_ResourceMetricSource,
		Convert_v1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus,
		Convert_autoscaling_ResourceMetricStatus_To_v1_ResourceMetricStatus,
		Convert_v1_Scale_To_autoscaling_Scale,
		Convert_autoscaling_Scale_To_v1_Scale,
		Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec,
		Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec,
		Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus,
		Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus,
	)
}

func autoConvert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in *v1.CrossVersionObjectReference, out *autoscaling.CrossVersionObjectReference, s conversion.Scope) error {
	out.Kind = in.Kind
	out.Name = in.Name
	out.APIVersion = in.APIVersion
	return nil
}

// Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference is an autogenerated conversion function.
func Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in *v1.CrossVersionObjectReference, out *autoscaling.CrossVersionObjectReference, s conversion.Scope) error {
	return autoConvert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in, out, s)
}

func autoConvert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(in *autoscaling.CrossVersionObjectReference, out *v1.CrossVersionObjectReference, s conversion.Scope) error {
	out.Kind = in.Kind
	out.Name = in.Name
	out.APIVersion = in.APIVersion
	return nil
}

// Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference is an autogenerated conversion function.
func Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(in *autoscaling.CrossVersionObjectReference, out *v1.CrossVersionObjectReference, s conversion.Scope) error {
	return autoConvert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(in, out, s)
}

func autoConvert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(in *v1.HorizontalPodAutoscaler, out *autoscaling.HorizontalPodAutoscaler, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_autoscaling_HorizontalPodAutoscaler_To_v1_HorizontalPodAutoscaler(in *autoscaling.HorizontalPodAutoscaler, out *v1.HorizontalPodAutoscaler, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition(in *v1.HorizontalPodAutoscalerCondition, out *autoscaling.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
	out.Type = autoscaling.HorizontalPodAutoscalerConditionType(in.Type)
	out.Status = autoscaling.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition is an autogenerated conversion function.
func Convert_v1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition(in *v1.HorizontalPodAutoscalerCondition, out *autoscaling.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
	return autoConvert_v1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscalerCondition_To_v1_HorizontalPodAutoscalerCondition(in *autoscaling.HorizontalPodAutoscalerCondition, out *v1.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
	out.Type = v1.HorizontalPodAutoscalerConditionType(in.Type)
	out.Status = core_v1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_autoscaling_HorizontalPodAutoscalerCondition_To_v1_HorizontalPodAutoscalerCondition is an autogenerated conversion function.
func Convert_autoscaling_HorizontalPodAutoscalerCondition_To_v1_HorizontalPodAutoscalerCondition(in *autoscaling.HorizontalPodAutoscalerCondition, out *v1.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscalerCondition_To_v1_HorizontalPodAutoscalerCondition(in, out, s)
}

func autoConvert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in *v1.HorizontalPodAutoscalerList, out *autoscaling.HorizontalPodAutoscalerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]autoscaling.HorizontalPodAutoscaler, len(*in))
		for i := range *in {
			if err := Convert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList is an autogenerated conversion function.
func Convert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in *v1.HorizontalPodAutoscalerList, out *autoscaling.HorizontalPodAutoscalerList, s conversion.Scope) error {
	return autoConvert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList(in *autoscaling.HorizontalPodAutoscalerList, out *v1.HorizontalPodAutoscalerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.HorizontalPodAutoscaler, len(*in))
		for i := range *in {
			if err := Convert_autoscaling_HorizontalPodAutoscaler_To_v1_HorizontalPodAutoscaler(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = make([]v1.HorizontalPodAutoscaler, 0)
	}
	return nil
}

// Convert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList is an autogenerated conversion function.
func Convert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList(in *autoscaling.HorizontalPodAutoscalerList, out *v1.HorizontalPodAutoscalerList, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList(in, out, s)
}

func autoConvert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(in *v1.HorizontalPodAutoscalerSpec, out *autoscaling.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	if err := Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(&in.ScaleTargetRef, &out.ScaleTargetRef, s); err != nil {
		return err
	}
	out.MinReplicas = (*int32)(unsafe.Pointer(in.MinReplicas))
	out.MaxReplicas = in.MaxReplicas
	// WARNING: in.TargetCPUUtilizationPercentage requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec(in *autoscaling.HorizontalPodAutoscalerSpec, out *v1.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	if err := Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(&in.ScaleTargetRef, &out.ScaleTargetRef, s); err != nil {
		return err
	}
	out.MinReplicas = (*int32)(unsafe.Pointer(in.MinReplicas))
	out.MaxReplicas = in.MaxReplicas
	// WARNING: in.Metrics requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(in *v1.HorizontalPodAutoscalerStatus, out *autoscaling.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.LastScaleTime = (*meta_v1.Time)(unsafe.Pointer(in.LastScaleTime))
	out.CurrentReplicas = in.CurrentReplicas
	out.DesiredReplicas = in.DesiredReplicas
	// WARNING: in.CurrentCPUUtilizationPercentage requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus(in *autoscaling.HorizontalPodAutoscalerStatus, out *v1.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.LastScaleTime = (*meta_v1.Time)(unsafe.Pointer(in.LastScaleTime))
	out.CurrentReplicas = in.CurrentReplicas
	out.DesiredReplicas = in.DesiredReplicas
	// WARNING: in.CurrentMetrics requires manual conversion: does not exist in peer-type
	// WARNING: in.Conditions requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1_MetricSpec_To_autoscaling_MetricSpec(in *v1.MetricSpec, out *autoscaling.MetricSpec, s conversion.Scope) error {
	out.Type = autoscaling.MetricSourceType(in.Type)
	out.Object = (*autoscaling.ObjectMetricSource)(unsafe.Pointer(in.Object))
	out.Pods = (*autoscaling.PodsMetricSource)(unsafe.Pointer(in.Pods))
	out.Resource = (*autoscaling.ResourceMetricSource)(unsafe.Pointer(in.Resource))
	return nil
}

// Convert_v1_MetricSpec_To_autoscaling_MetricSpec is an autogenerated conversion function.
func Convert_v1_MetricSpec_To_autoscaling_MetricSpec(in *v1.MetricSpec, out *autoscaling.MetricSpec, s conversion.Scope) error {
	return autoConvert_v1_MetricSpec_To_autoscaling_MetricSpec(in, out, s)
}

func autoConvert_autoscaling_MetricSpec_To_v1_MetricSpec(in *autoscaling.MetricSpec, out *v1.MetricSpec, s conversion.Scope) error {
	out.Type = v1.MetricSourceType(in.Type)
	out.Object = (*v1.ObjectMetricSource)(unsafe.Pointer(in.Object))
	out.Pods = (*v1.PodsMetricSource)(unsafe.Pointer(in.Pods))
	out.Resource = (*v1.ResourceMetricSource)(unsafe.Pointer(in.Resource))
	return nil
}

// Convert_autoscaling_MetricSpec_To_v1_MetricSpec is an autogenerated conversion function.
func Convert_autoscaling_MetricSpec_To_v1_MetricSpec(in *autoscaling.MetricSpec, out *v1.MetricSpec, s conversion.Scope) error {
	return autoConvert_autoscaling_MetricSpec_To_v1_MetricSpec(in, out, s)
}

func autoConvert_v1_MetricStatus_To_autoscaling_MetricStatus(in *v1.MetricStatus, out *autoscaling.MetricStatus, s conversion.Scope) error {
	out.Type = autoscaling.MetricSourceType(in.Type)
	out.Object = (*autoscaling.ObjectMetricStatus)(unsafe.Pointer(in.Object))
	out.Pods = (*autoscaling.PodsMetricStatus)(unsafe.Pointer(in.Pods))
	out.Resource = (*autoscaling.ResourceMetricStatus)(unsafe.Pointer(in.Resource))
	return nil
}

// Convert_v1_MetricStatus_To_autoscaling_MetricStatus is an autogenerated conversion function.
func Convert_v1_MetricStatus_To_autoscaling_MetricStatus(in *v1.MetricStatus, out *autoscaling.MetricStatus, s conversion.Scope) error {
	return autoConvert_v1_MetricStatus_To_autoscaling_MetricStatus(in, out, s)
}

func autoConvert_autoscaling_MetricStatus_To_v1_MetricStatus(in *autoscaling.MetricStatus, out *v1.MetricStatus, s conversion.Scope) error {
	out.Type = v1.MetricSourceType(in.Type)
	out.Object = (*v1.ObjectMetricStatus)(unsafe.Pointer(in.Object))
	out.Pods = (*v1.PodsMetricStatus)(unsafe.Pointer(in.Pods))
	out.Resource = (*v1.ResourceMetricStatus)(unsafe.Pointer(in.Resource))
	return nil
}

// Convert_autoscaling_MetricStatus_To_v1_MetricStatus is an autogenerated conversion function.
func Convert_autoscaling_MetricStatus_To_v1_MetricStatus(in *autoscaling.MetricStatus, out *v1.MetricStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_MetricStatus_To_v1_MetricStatus(in, out, s)
}

func autoConvert_v1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(in *v1.ObjectMetricSource, out *autoscaling.ObjectMetricSource, s conversion.Scope) error {
	if err := Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(&in.Target, &out.Target, s); err != nil {
		return err
	}
	out.MetricName = in.MetricName
	out.TargetValue = in.TargetValue
	return nil
}

// Convert_v1_ObjectMetricSource_To_autoscaling_ObjectMetricSource is an autogenerated conversion function.
func Convert_v1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(in *v1.ObjectMetricSource, out *autoscaling.ObjectMetricSource, s conversion.Scope) error {
	return autoConvert_v1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(in, out, s)
}

func autoConvert_autoscaling_ObjectMetricSource_To_v1_ObjectMetricSource(in *autoscaling.ObjectMetricSource, out *v1.ObjectMetricSource, s conversion.Scope) error {
	if err := Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(&in.Target, &out.Target, s); err != nil {
		return err
	}
	out.MetricName = in.MetricName
	out.TargetValue = in.TargetValue
	return nil
}

// Convert_autoscaling_ObjectMetricSource_To_v1_ObjectMetricSource is an autogenerated conversion function.
func Convert_autoscaling_ObjectMetricSource_To_v1_ObjectMetricSource(in *autoscaling.ObjectMetricSource, out *v1.ObjectMetricSource, s conversion.Scope) error {
	return autoConvert_autoscaling_ObjectMetricSource_To_v1_ObjectMetricSource(in, out, s)
}

func autoConvert_v1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(in *v1.ObjectMetricStatus, out *autoscaling.ObjectMetricStatus, s conversion.Scope) error {
	if err := Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(&in.Target, &out.Target, s); err != nil {
		return err
	}
	out.MetricName = in.MetricName
	out.CurrentValue = in.CurrentValue
	return nil
}

// Convert_v1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus is an autogenerated conversion function.
func Convert_v1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(in *v1.ObjectMetricStatus, out *autoscaling.ObjectMetricStatus, s conversion.Scope) error {
	return autoConvert_v1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(in, out, s)
}

func autoConvert_autoscaling_ObjectMetricStatus_To_v1_ObjectMetricStatus(in *autoscaling.ObjectMetricStatus, out *v1.ObjectMetricStatus, s conversion.Scope) error {
	if err := Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(&in.Target, &out.Target, s); err != nil {
		return err
	}
	out.MetricName = in.MetricName
	out.CurrentValue = in.CurrentValue
	return nil
}

// Convert_autoscaling_ObjectMetricStatus_To_v1_ObjectMetricStatus is an autogenerated conversion function.
func Convert_autoscaling_ObjectMetricStatus_To_v1_ObjectMetricStatus(in *autoscaling.ObjectMetricStatus, out *v1.ObjectMetricStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_ObjectMetricStatus_To_v1_ObjectMetricStatus(in, out, s)
}

func autoConvert_v1_PodsMetricSource_To_autoscaling_PodsMetricSource(in *v1.PodsMetricSource, out *autoscaling.PodsMetricSource, s conversion.Scope) error {
	out.MetricName = in.MetricName
	out.TargetAverageValue = in.TargetAverageValue
	return nil
}

// Convert_v1_PodsMetricSource_To_autoscaling_PodsMetricSource is an autogenerated conversion function.
func Convert_v1_PodsMetricSource_To_autoscaling_PodsMetricSource(in *v1.PodsMetricSource, out *autoscaling.PodsMetricSource, s conversion.Scope) error {
	return autoConvert_v1_PodsMetricSource_To_autoscaling_PodsMetricSource(in, out, s)
}

func autoConvert_autoscaling_PodsMetricSource_To_v1_PodsMetricSource(in *autoscaling.PodsMetricSource, out *v1.PodsMetricSource, s conversion.Scope) error {
	out.MetricName = in.MetricName
	out.TargetAverageValue = in.TargetAverageValue
	return nil
}

// Convert_autoscaling_PodsMetricSource_To_v1_PodsMetricSource is an autogenerated conversion function.
func Convert_autoscaling_PodsMetricSource_To_v1_PodsMetricSource(in *autoscaling.PodsMetricSource, out *v1.PodsMetricSource, s conversion.Scope) error {
	return autoConvert_autoscaling_PodsMetricSource_To_v1_PodsMetricSource(in, out, s)
}

func autoConvert_v1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(in *v1.PodsMetricStatus, out *autoscaling.PodsMetricStatus, s conversion.Scope) error {
	out.MetricName = in.MetricName
	out.CurrentAverageValue = in.CurrentAverageValue
	return nil
}

// Convert_v1_PodsMetricStatus_To_autoscaling_PodsMetricStatus is an autogenerated conversion function.
func Convert_v1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(in *v1.PodsMetricStatus, out *autoscaling.PodsMetricStatus, s conversion.Scope) error {
	return autoConvert_v1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(in, out, s)
}

func autoConvert_autoscaling_PodsMetricStatus_To_v1_PodsMetricStatus(in *autoscaling.PodsMetricStatus, out *v1.PodsMetricStatus, s conversion.Scope) error {
	out.MetricName = in.MetricName
	out.CurrentAverageValue = in.CurrentAverageValue
	return nil
}

// Convert_autoscaling_PodsMetricStatus_To_v1_PodsMetricStatus is an autogenerated conversion function.
func Convert_autoscaling_PodsMetricStatus_To_v1_PodsMetricStatus(in *autoscaling.PodsMetricStatus, out *v1.PodsMetricStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_PodsMetricStatus_To_v1_PodsMetricStatus(in, out, s)
}

func autoConvert_v1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(in *v1.ResourceMetricSource, out *autoscaling.ResourceMetricSource, s conversion.Scope) error {
	out.Name = api.ResourceName(in.Name)
	out.TargetAverageUtilization = (*int32)(unsafe.Pointer(in.TargetAverageUtilization))
	out.TargetAverageValue = (*resource.Quantity)(unsafe.Pointer(in.TargetAverageValue))
	return nil
}

// Convert_v1_ResourceMetricSource_To_autoscaling_ResourceMetricSource is an autogenerated conversion function.
func Convert_v1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(in *v1.ResourceMetricSource, out *autoscaling.ResourceMetricSource, s conversion.Scope) error {
	return autoConvert_v1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(in, out, s)
}

func autoConvert_autoscaling_ResourceMetricSource_To_v1_ResourceMetricSource(in *autoscaling.ResourceMetricSource, out *v1.ResourceMetricSource, s conversion.Scope) error {
	out.Name = core_v1.ResourceName(in.Name)
	out.TargetAverageUtilization = (*int32)(unsafe.Pointer(in.TargetAverageUtilization))
	out.TargetAverageValue = (*resource.Quantity)(unsafe.Pointer(in.TargetAverageValue))
	return nil
}

// Convert_autoscaling_ResourceMetricSource_To_v1_ResourceMetricSource is an autogenerated conversion function.
func Convert_autoscaling_ResourceMetricSource_To_v1_ResourceMetricSource(in *autoscaling.ResourceMetricSource, out *v1.ResourceMetricSource, s conversion.Scope) error {
	return autoConvert_autoscaling_ResourceMetricSource_To_v1_ResourceMetricSource(in, out, s)
}

func autoConvert_v1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(in *v1.ResourceMetricStatus, out *autoscaling.ResourceMetricStatus, s conversion.Scope) error {
	out.Name = api.ResourceName(in.Name)
	out.CurrentAverageUtilization = (*int32)(unsafe.Pointer(in.CurrentAverageUtilization))
	out.CurrentAverageValue = in.CurrentAverageValue
	return nil
}

// Convert_v1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus is an autogenerated conversion function.
func Convert_v1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(in *v1.ResourceMetricStatus, out *autoscaling.ResourceMetricStatus, s conversion.Scope) error {
	return autoConvert_v1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(in, out, s)
}

func autoConvert_autoscaling_ResourceMetricStatus_To_v1_ResourceMetricStatus(in *autoscaling.ResourceMetricStatus, out *v1.ResourceMetricStatus, s conversion.Scope) error {
	out.Name = core_v1.ResourceName(in.Name)
	out.CurrentAverageUtilization = (*int32)(unsafe.Pointer(in.CurrentAverageUtilization))
	out.CurrentAverageValue = in.CurrentAverageValue
	return nil
}

// Convert_autoscaling_ResourceMetricStatus_To_v1_ResourceMetricStatus is an autogenerated conversion function.
func Convert_autoscaling_ResourceMetricStatus_To_v1_ResourceMetricStatus(in *autoscaling.ResourceMetricStatus, out *v1.ResourceMetricStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_ResourceMetricStatus_To_v1_ResourceMetricStatus(in, out, s)
}

func autoConvert_v1_Scale_To_autoscaling_Scale(in *v1.Scale, out *autoscaling.Scale, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Scale_To_autoscaling_Scale is an autogenerated conversion function.
func Convert_v1_Scale_To_autoscaling_Scale(in *v1.Scale, out *autoscaling.Scale, s conversion.Scope) error {
	return autoConvert_v1_Scale_To_autoscaling_Scale(in, out, s)
}

func autoConvert_autoscaling_Scale_To_v1_Scale(in *autoscaling.Scale, out *v1.Scale, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_autoscaling_Scale_To_v1_Scale is an autogenerated conversion function.
func Convert_autoscaling_Scale_To_v1_Scale(in *autoscaling.Scale, out *v1.Scale, s conversion.Scope) error {
	return autoConvert_autoscaling_Scale_To_v1_Scale(in, out, s)
}

func autoConvert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in *v1.ScaleSpec, out *autoscaling.ScaleSpec, s conversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

// Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec is an autogenerated conversion function.
func Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in *v1.ScaleSpec, out *autoscaling.ScaleSpec, s conversion.Scope) error {
	return autoConvert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in, out, s)
}

func autoConvert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in *autoscaling.ScaleSpec, out *v1.ScaleSpec, s conversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

// Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec is an autogenerated conversion function.
func Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in *autoscaling.ScaleSpec, out *v1.ScaleSpec, s conversion.Scope) error {
	return autoConvert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in, out, s)
}

func autoConvert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in *v1.ScaleStatus, out *autoscaling.ScaleStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	out.Selector = in.Selector
	return nil
}

// Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus is an autogenerated conversion function.
func Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in *v1.ScaleStatus, out *autoscaling.ScaleStatus, s conversion.Scope) error {
	return autoConvert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in, out, s)
}

func autoConvert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in *autoscaling.ScaleStatus, out *v1.ScaleStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	out.Selector = in.Selector
	return nil
}

// Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus is an autogenerated conversion function.
func Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in *autoscaling.ScaleStatus, out *v1.ScaleStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in, out, s)
}
