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
	v1 "k8s.io/api/batch/v1"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "gitlab.com/rioos/magudi/pkg/api"
	api_v1 "gitlab.com/rioos/magudi/pkg/api/v1"
	batch "gitlab.com/rioos/magudi/pkg/apis/batch"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_Job_To_batch_Job,
		Convert_batch_Job_To_v1_Job,
		Convert_v1_JobCondition_To_batch_JobCondition,
		Convert_batch_JobCondition_To_v1_JobCondition,
		Convert_v1_JobList_To_batch_JobList,
		Convert_batch_JobList_To_v1_JobList,
		Convert_v1_JobSpec_To_batch_JobSpec,
		Convert_batch_JobSpec_To_v1_JobSpec,
		Convert_v1_JobStatus_To_batch_JobStatus,
		Convert_batch_JobStatus_To_v1_JobStatus,
	)
}

func autoConvert_v1_Job_To_batch_Job(in *v1.Job, out *batch.Job, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_JobSpec_To_batch_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_JobStatus_To_batch_JobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Job_To_batch_Job is an autogenerated conversion function.
func Convert_v1_Job_To_batch_Job(in *v1.Job, out *batch.Job, s conversion.Scope) error {
	return autoConvert_v1_Job_To_batch_Job(in, out, s)
}

func autoConvert_batch_Job_To_v1_Job(in *batch.Job, out *v1.Job, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_batch_JobSpec_To_v1_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_batch_JobStatus_To_v1_JobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_batch_Job_To_v1_Job is an autogenerated conversion function.
func Convert_batch_Job_To_v1_Job(in *batch.Job, out *v1.Job, s conversion.Scope) error {
	return autoConvert_batch_Job_To_v1_Job(in, out, s)
}

func autoConvert_v1_JobCondition_To_batch_JobCondition(in *v1.JobCondition, out *batch.JobCondition, s conversion.Scope) error {
	out.Type = batch.JobConditionType(in.Type)
	out.Status = api.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1_JobCondition_To_batch_JobCondition is an autogenerated conversion function.
func Convert_v1_JobCondition_To_batch_JobCondition(in *v1.JobCondition, out *batch.JobCondition, s conversion.Scope) error {
	return autoConvert_v1_JobCondition_To_batch_JobCondition(in, out, s)
}

func autoConvert_batch_JobCondition_To_v1_JobCondition(in *batch.JobCondition, out *v1.JobCondition, s conversion.Scope) error {
	out.Type = v1.JobConditionType(in.Type)
	out.Status = core_v1.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_batch_JobCondition_To_v1_JobCondition is an autogenerated conversion function.
func Convert_batch_JobCondition_To_v1_JobCondition(in *batch.JobCondition, out *v1.JobCondition, s conversion.Scope) error {
	return autoConvert_batch_JobCondition_To_v1_JobCondition(in, out, s)
}

func autoConvert_v1_JobList_To_batch_JobList(in *v1.JobList, out *batch.JobList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]batch.Job, len(*in))
		for i := range *in {
			if err := Convert_v1_Job_To_batch_Job(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_JobList_To_batch_JobList is an autogenerated conversion function.
func Convert_v1_JobList_To_batch_JobList(in *v1.JobList, out *batch.JobList, s conversion.Scope) error {
	return autoConvert_v1_JobList_To_batch_JobList(in, out, s)
}

func autoConvert_batch_JobList_To_v1_JobList(in *batch.JobList, out *v1.JobList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.Job, len(*in))
		for i := range *in {
			if err := Convert_batch_Job_To_v1_Job(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = make([]v1.Job, 0)
	}
	return nil
}

// Convert_batch_JobList_To_v1_JobList is an autogenerated conversion function.
func Convert_batch_JobList_To_v1_JobList(in *batch.JobList, out *v1.JobList, s conversion.Scope) error {
	return autoConvert_batch_JobList_To_v1_JobList(in, out, s)
}

func autoConvert_v1_JobSpec_To_batch_JobSpec(in *v1.JobSpec, out *batch.JobSpec, s conversion.Scope) error {
	out.Parallelism = (*int32)(unsafe.Pointer(in.Parallelism))
	out.Completions = (*int32)(unsafe.Pointer(in.Completions))
	out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
	out.Selector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.ManualSelector = (*bool)(unsafe.Pointer(in.ManualSelector))
	if err := api_v1.Convert_v1_PodTemplateSpec_To_api_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_batch_JobSpec_To_v1_JobSpec(in *batch.JobSpec, out *v1.JobSpec, s conversion.Scope) error {
	out.Parallelism = (*int32)(unsafe.Pointer(in.Parallelism))
	out.Completions = (*int32)(unsafe.Pointer(in.Completions))
	out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
	out.Selector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.ManualSelector = (*bool)(unsafe.Pointer(in.ManualSelector))
	if err := api_v1.Convert_api_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_JobStatus_To_batch_JobStatus(in *v1.JobStatus, out *batch.JobStatus, s conversion.Scope) error {
	out.Conditions = *(*[]batch.JobCondition)(unsafe.Pointer(&in.Conditions))
	out.StartTime = (*meta_v1.Time)(unsafe.Pointer(in.StartTime))
	out.CompletionTime = (*meta_v1.Time)(unsafe.Pointer(in.CompletionTime))
	out.Active = in.Active
	out.Succeeded = in.Succeeded
	out.Failed = in.Failed
	return nil
}

// Convert_v1_JobStatus_To_batch_JobStatus is an autogenerated conversion function.
func Convert_v1_JobStatus_To_batch_JobStatus(in *v1.JobStatus, out *batch.JobStatus, s conversion.Scope) error {
	return autoConvert_v1_JobStatus_To_batch_JobStatus(in, out, s)
}

func autoConvert_batch_JobStatus_To_v1_JobStatus(in *batch.JobStatus, out *v1.JobStatus, s conversion.Scope) error {
	out.Conditions = *(*[]v1.JobCondition)(unsafe.Pointer(&in.Conditions))
	out.StartTime = (*meta_v1.Time)(unsafe.Pointer(in.StartTime))
	out.CompletionTime = (*meta_v1.Time)(unsafe.Pointer(in.CompletionTime))
	out.Active = in.Active
	out.Succeeded = in.Succeeded
	out.Failed = in.Failed
	return nil
}

// Convert_batch_JobStatus_To_v1_JobStatus is an autogenerated conversion function.
func Convert_batch_JobStatus_To_v1_JobStatus(in *batch.JobStatus, out *v1.JobStatus, s conversion.Scope) error {
	return autoConvert_batch_JobStatus_To_v1_JobStatus(in, out, s)
}
