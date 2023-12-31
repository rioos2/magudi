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

package admission

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	authentication "gitlab.com/rioos/magudi/pkg/apis/authentication"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_admission_AdmissionReview, InType: reflect.TypeOf(&AdmissionReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_admission_AdmissionReviewSpec, InType: reflect.TypeOf(&AdmissionReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_admission_AdmissionReviewStatus, InType: reflect.TypeOf(&AdmissionReviewStatus{})},
	)
}

// DeepCopy_admission_AdmissionReview is an autogenerated deepcopy function.
func DeepCopy_admission_AdmissionReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AdmissionReview)
		out := out.(*AdmissionReview)
		*out = *in
		if err := DeepCopy_admission_AdmissionReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_admission_AdmissionReviewStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_admission_AdmissionReviewSpec is an autogenerated deepcopy function.
func DeepCopy_admission_AdmissionReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AdmissionReviewSpec)
		out := out.(*AdmissionReviewSpec)
		*out = *in
		// in.Object is kind 'Interface'
		if in.Object != nil {
			if newVal, err := c.DeepCopy(&in.Object); err != nil {
				return err
			} else {
				out.Object = *newVal.(*runtime.Object)
			}
		}
		// in.OldObject is kind 'Interface'
		if in.OldObject != nil {
			if newVal, err := c.DeepCopy(&in.OldObject); err != nil {
				return err
			} else {
				out.OldObject = *newVal.(*runtime.Object)
			}
		}
		if err := authentication.DeepCopy_authentication_UserInfo(&in.UserInfo, &out.UserInfo, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_admission_AdmissionReviewStatus is an autogenerated deepcopy function.
func DeepCopy_admission_AdmissionReviewStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AdmissionReviewStatus)
		out := out.(*AdmissionReviewStatus)
		*out = *in
		if in.Result != nil {
			in, out := &in.Result, &out.Result
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.Status)
			}
		}
		return nil
	}
}
