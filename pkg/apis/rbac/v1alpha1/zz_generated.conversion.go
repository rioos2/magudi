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

package v1alpha1

import (
	v1alpha1 "k8s.io/api/rbac/v1alpha1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	rbac "gitlab.com/rioos/magudi/pkg/apis/rbac"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_ClusterRole_To_rbac_ClusterRole,
		Convert_rbac_ClusterRole_To_v1alpha1_ClusterRole,
		Convert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding,
		Convert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding,
		Convert_v1alpha1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList,
		Convert_rbac_ClusterRoleBindingList_To_v1alpha1_ClusterRoleBindingList,
		Convert_v1alpha1_ClusterRoleList_To_rbac_ClusterRoleList,
		Convert_rbac_ClusterRoleList_To_v1alpha1_ClusterRoleList,
		Convert_v1alpha1_PolicyRule_To_rbac_PolicyRule,
		Convert_rbac_PolicyRule_To_v1alpha1_PolicyRule,
		Convert_v1alpha1_Role_To_rbac_Role,
		Convert_rbac_Role_To_v1alpha1_Role,
		Convert_v1alpha1_RoleBinding_To_rbac_RoleBinding,
		Convert_rbac_RoleBinding_To_v1alpha1_RoleBinding,
		Convert_v1alpha1_RoleBindingList_To_rbac_RoleBindingList,
		Convert_rbac_RoleBindingList_To_v1alpha1_RoleBindingList,
		Convert_v1alpha1_RoleList_To_rbac_RoleList,
		Convert_rbac_RoleList_To_v1alpha1_RoleList,
		Convert_v1alpha1_RoleRef_To_rbac_RoleRef,
		Convert_rbac_RoleRef_To_v1alpha1_RoleRef,
		Convert_v1alpha1_Subject_To_rbac_Subject,
		Convert_rbac_Subject_To_v1alpha1_Subject,
	)
}

func autoConvert_v1alpha1_ClusterRole_To_rbac_ClusterRole(in *v1alpha1.ClusterRole, out *rbac.ClusterRole, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Rules = *(*[]rbac.PolicyRule)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_v1alpha1_ClusterRole_To_rbac_ClusterRole is an autogenerated conversion function.
func Convert_v1alpha1_ClusterRole_To_rbac_ClusterRole(in *v1alpha1.ClusterRole, out *rbac.ClusterRole, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterRole_To_rbac_ClusterRole(in, out, s)
}

func autoConvert_rbac_ClusterRole_To_v1alpha1_ClusterRole(in *rbac.ClusterRole, out *v1alpha1.ClusterRole, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if in.Rules == nil {
		out.Rules = make([]v1alpha1.PolicyRule, 0)
	} else {
		out.Rules = *(*[]v1alpha1.PolicyRule)(unsafe.Pointer(&in.Rules))
	}
	return nil
}

// Convert_rbac_ClusterRole_To_v1alpha1_ClusterRole is an autogenerated conversion function.
func Convert_rbac_ClusterRole_To_v1alpha1_ClusterRole(in *rbac.ClusterRole, out *v1alpha1.ClusterRole, s conversion.Scope) error {
	return autoConvert_rbac_ClusterRole_To_v1alpha1_ClusterRole(in, out, s)
}

func autoConvert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(in *v1alpha1.ClusterRoleBinding, out *rbac.ClusterRoleBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]rbac.Subject, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_Subject_To_rbac_Subject(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Subjects = nil
	}
	if err := Convert_v1alpha1_RoleRef_To_rbac_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding is an autogenerated conversion function.
func Convert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(in *v1alpha1.ClusterRoleBinding, out *rbac.ClusterRoleBinding, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(in, out, s)
}

func autoConvert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding(in *rbac.ClusterRoleBinding, out *v1alpha1.ClusterRoleBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]v1alpha1.Subject, len(*in))
		for i := range *in {
			if err := Convert_rbac_Subject_To_v1alpha1_Subject(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Subjects = make([]v1alpha1.Subject, 0)
	}
	if err := Convert_rbac_RoleRef_To_v1alpha1_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}

// Convert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding is an autogenerated conversion function.
func Convert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding(in *rbac.ClusterRoleBinding, out *v1alpha1.ClusterRoleBinding, s conversion.Scope) error {
	return autoConvert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding(in, out, s)
}

func autoConvert_v1alpha1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList(in *v1alpha1.ClusterRoleBindingList, out *rbac.ClusterRoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]rbac.ClusterRoleBinding, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList is an autogenerated conversion function.
func Convert_v1alpha1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList(in *v1alpha1.ClusterRoleBindingList, out *rbac.ClusterRoleBindingList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList(in, out, s)
}

func autoConvert_rbac_ClusterRoleBindingList_To_v1alpha1_ClusterRoleBindingList(in *rbac.ClusterRoleBindingList, out *v1alpha1.ClusterRoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1alpha1.ClusterRoleBinding, len(*in))
		for i := range *in {
			if err := Convert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = make([]v1alpha1.ClusterRoleBinding, 0)
	}
	return nil
}

// Convert_rbac_ClusterRoleBindingList_To_v1alpha1_ClusterRoleBindingList is an autogenerated conversion function.
func Convert_rbac_ClusterRoleBindingList_To_v1alpha1_ClusterRoleBindingList(in *rbac.ClusterRoleBindingList, out *v1alpha1.ClusterRoleBindingList, s conversion.Scope) error {
	return autoConvert_rbac_ClusterRoleBindingList_To_v1alpha1_ClusterRoleBindingList(in, out, s)
}

func autoConvert_v1alpha1_ClusterRoleList_To_rbac_ClusterRoleList(in *v1alpha1.ClusterRoleList, out *rbac.ClusterRoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]rbac.ClusterRole)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ClusterRoleList_To_rbac_ClusterRoleList is an autogenerated conversion function.
func Convert_v1alpha1_ClusterRoleList_To_rbac_ClusterRoleList(in *v1alpha1.ClusterRoleList, out *rbac.ClusterRoleList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterRoleList_To_rbac_ClusterRoleList(in, out, s)
}

func autoConvert_rbac_ClusterRoleList_To_v1alpha1_ClusterRoleList(in *rbac.ClusterRoleList, out *v1alpha1.ClusterRoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]v1alpha1.ClusterRole, 0)
	} else {
		out.Items = *(*[]v1alpha1.ClusterRole)(unsafe.Pointer(&in.Items))
	}
	return nil
}

// Convert_rbac_ClusterRoleList_To_v1alpha1_ClusterRoleList is an autogenerated conversion function.
func Convert_rbac_ClusterRoleList_To_v1alpha1_ClusterRoleList(in *rbac.ClusterRoleList, out *v1alpha1.ClusterRoleList, s conversion.Scope) error {
	return autoConvert_rbac_ClusterRoleList_To_v1alpha1_ClusterRoleList(in, out, s)
}

func autoConvert_v1alpha1_PolicyRule_To_rbac_PolicyRule(in *v1alpha1.PolicyRule, out *rbac.PolicyRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	return nil
}

// Convert_v1alpha1_PolicyRule_To_rbac_PolicyRule is an autogenerated conversion function.
func Convert_v1alpha1_PolicyRule_To_rbac_PolicyRule(in *v1alpha1.PolicyRule, out *rbac.PolicyRule, s conversion.Scope) error {
	return autoConvert_v1alpha1_PolicyRule_To_rbac_PolicyRule(in, out, s)
}

func autoConvert_rbac_PolicyRule_To_v1alpha1_PolicyRule(in *rbac.PolicyRule, out *v1alpha1.PolicyRule, s conversion.Scope) error {
	if in.Verbs == nil {
		out.Verbs = make([]string, 0)
	} else {
		out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	}
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	return nil
}

// Convert_rbac_PolicyRule_To_v1alpha1_PolicyRule is an autogenerated conversion function.
func Convert_rbac_PolicyRule_To_v1alpha1_PolicyRule(in *rbac.PolicyRule, out *v1alpha1.PolicyRule, s conversion.Scope) error {
	return autoConvert_rbac_PolicyRule_To_v1alpha1_PolicyRule(in, out, s)
}

func autoConvert_v1alpha1_Role_To_rbac_Role(in *v1alpha1.Role, out *rbac.Role, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Rules = *(*[]rbac.PolicyRule)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_v1alpha1_Role_To_rbac_Role is an autogenerated conversion function.
func Convert_v1alpha1_Role_To_rbac_Role(in *v1alpha1.Role, out *rbac.Role, s conversion.Scope) error {
	return autoConvert_v1alpha1_Role_To_rbac_Role(in, out, s)
}

func autoConvert_rbac_Role_To_v1alpha1_Role(in *rbac.Role, out *v1alpha1.Role, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if in.Rules == nil {
		out.Rules = make([]v1alpha1.PolicyRule, 0)
	} else {
		out.Rules = *(*[]v1alpha1.PolicyRule)(unsafe.Pointer(&in.Rules))
	}
	return nil
}

// Convert_rbac_Role_To_v1alpha1_Role is an autogenerated conversion function.
func Convert_rbac_Role_To_v1alpha1_Role(in *rbac.Role, out *v1alpha1.Role, s conversion.Scope) error {
	return autoConvert_rbac_Role_To_v1alpha1_Role(in, out, s)
}

func autoConvert_v1alpha1_RoleBinding_To_rbac_RoleBinding(in *v1alpha1.RoleBinding, out *rbac.RoleBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]rbac.Subject, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_Subject_To_rbac_Subject(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Subjects = nil
	}
	if err := Convert_v1alpha1_RoleRef_To_rbac_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_RoleBinding_To_rbac_RoleBinding is an autogenerated conversion function.
func Convert_v1alpha1_RoleBinding_To_rbac_RoleBinding(in *v1alpha1.RoleBinding, out *rbac.RoleBinding, s conversion.Scope) error {
	return autoConvert_v1alpha1_RoleBinding_To_rbac_RoleBinding(in, out, s)
}

func autoConvert_rbac_RoleBinding_To_v1alpha1_RoleBinding(in *rbac.RoleBinding, out *v1alpha1.RoleBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]v1alpha1.Subject, len(*in))
		for i := range *in {
			if err := Convert_rbac_Subject_To_v1alpha1_Subject(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Subjects = make([]v1alpha1.Subject, 0)
	}
	if err := Convert_rbac_RoleRef_To_v1alpha1_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}

// Convert_rbac_RoleBinding_To_v1alpha1_RoleBinding is an autogenerated conversion function.
func Convert_rbac_RoleBinding_To_v1alpha1_RoleBinding(in *rbac.RoleBinding, out *v1alpha1.RoleBinding, s conversion.Scope) error {
	return autoConvert_rbac_RoleBinding_To_v1alpha1_RoleBinding(in, out, s)
}

func autoConvert_v1alpha1_RoleBindingList_To_rbac_RoleBindingList(in *v1alpha1.RoleBindingList, out *rbac.RoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]rbac.RoleBinding, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_RoleBinding_To_rbac_RoleBinding(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_RoleBindingList_To_rbac_RoleBindingList is an autogenerated conversion function.
func Convert_v1alpha1_RoleBindingList_To_rbac_RoleBindingList(in *v1alpha1.RoleBindingList, out *rbac.RoleBindingList, s conversion.Scope) error {
	return autoConvert_v1alpha1_RoleBindingList_To_rbac_RoleBindingList(in, out, s)
}

func autoConvert_rbac_RoleBindingList_To_v1alpha1_RoleBindingList(in *rbac.RoleBindingList, out *v1alpha1.RoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1alpha1.RoleBinding, len(*in))
		for i := range *in {
			if err := Convert_rbac_RoleBinding_To_v1alpha1_RoleBinding(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = make([]v1alpha1.RoleBinding, 0)
	}
	return nil
}

// Convert_rbac_RoleBindingList_To_v1alpha1_RoleBindingList is an autogenerated conversion function.
func Convert_rbac_RoleBindingList_To_v1alpha1_RoleBindingList(in *rbac.RoleBindingList, out *v1alpha1.RoleBindingList, s conversion.Scope) error {
	return autoConvert_rbac_RoleBindingList_To_v1alpha1_RoleBindingList(in, out, s)
}

func autoConvert_v1alpha1_RoleList_To_rbac_RoleList(in *v1alpha1.RoleList, out *rbac.RoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]rbac.Role)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_RoleList_To_rbac_RoleList is an autogenerated conversion function.
func Convert_v1alpha1_RoleList_To_rbac_RoleList(in *v1alpha1.RoleList, out *rbac.RoleList, s conversion.Scope) error {
	return autoConvert_v1alpha1_RoleList_To_rbac_RoleList(in, out, s)
}

func autoConvert_rbac_RoleList_To_v1alpha1_RoleList(in *rbac.RoleList, out *v1alpha1.RoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]v1alpha1.Role, 0)
	} else {
		out.Items = *(*[]v1alpha1.Role)(unsafe.Pointer(&in.Items))
	}
	return nil
}

// Convert_rbac_RoleList_To_v1alpha1_RoleList is an autogenerated conversion function.
func Convert_rbac_RoleList_To_v1alpha1_RoleList(in *rbac.RoleList, out *v1alpha1.RoleList, s conversion.Scope) error {
	return autoConvert_rbac_RoleList_To_v1alpha1_RoleList(in, out, s)
}

func autoConvert_v1alpha1_RoleRef_To_rbac_RoleRef(in *v1alpha1.RoleRef, out *rbac.RoleRef, s conversion.Scope) error {
	out.APIGroup = in.APIGroup
	out.Kind = in.Kind
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_RoleRef_To_rbac_RoleRef is an autogenerated conversion function.
func Convert_v1alpha1_RoleRef_To_rbac_RoleRef(in *v1alpha1.RoleRef, out *rbac.RoleRef, s conversion.Scope) error {
	return autoConvert_v1alpha1_RoleRef_To_rbac_RoleRef(in, out, s)
}

func autoConvert_rbac_RoleRef_To_v1alpha1_RoleRef(in *rbac.RoleRef, out *v1alpha1.RoleRef, s conversion.Scope) error {
	out.APIGroup = in.APIGroup
	out.Kind = in.Kind
	out.Name = in.Name
	return nil
}

// Convert_rbac_RoleRef_To_v1alpha1_RoleRef is an autogenerated conversion function.
func Convert_rbac_RoleRef_To_v1alpha1_RoleRef(in *rbac.RoleRef, out *v1alpha1.RoleRef, s conversion.Scope) error {
	return autoConvert_rbac_RoleRef_To_v1alpha1_RoleRef(in, out, s)
}

func autoConvert_v1alpha1_Subject_To_rbac_Subject(in *v1alpha1.Subject, out *rbac.Subject, s conversion.Scope) error {
	out.Kind = in.Kind
	// INFO: in.APIVersion opted out of conversion generation
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}

func autoConvert_rbac_Subject_To_v1alpha1_Subject(in *rbac.Subject, out *v1alpha1.Subject, s conversion.Scope) error {
	out.Kind = in.Kind
	// WARNING: in.APIGroup requires manual conversion: does not exist in peer-type
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}
