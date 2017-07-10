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
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	coreinternalversion "gitlab.com/rioos/magudi/pkg/api"
	admissionregistrationinternalversion "gitlab.com/rioos/magudi/pkg/apis/admissionregistration"
	appsinternalversion "gitlab.com/rioos/magudi/pkg/apis/apps"
	authenticationinternalversion "gitlab.com/rioos/magudi/pkg/apis/authentication"
	authorizationinternalversion "gitlab.com/rioos/magudi/pkg/apis/authorization"
	autoscalinginternalversion "gitlab.com/rioos/magudi/pkg/apis/autoscaling"
	batchinternalversion "gitlab.com/rioos/magudi/pkg/apis/batch"
	certificatesinternalversion "gitlab.com/rioos/magudi/pkg/apis/certificates"
	extensionsinternalversion "gitlab.com/rioos/magudi/pkg/apis/extensions"
	networkinginternalversion "gitlab.com/rioos/magudi/pkg/apis/networking"
	policyinternalversion "gitlab.com/rioos/magudi/pkg/apis/policy"
	rbacinternalversion "gitlab.com/rioos/magudi/pkg/apis/rbac"
	settingsinternalversion "gitlab.com/rioos/magudi/pkg/apis/settings"
	storageinternalversion "gitlab.com/rioos/magudi/pkg/apis/storage"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)
var parameterCodec = runtime.NewParameterCodec(scheme)

func init() {
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	AddToScheme(scheme)
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kuberentes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
func AddToScheme(scheme *runtime.Scheme) {
	admissionregistrationinternalversion.AddToScheme(scheme)
	coreinternalversion.AddToScheme(scheme)
	appsinternalversion.AddToScheme(scheme)
	authenticationinternalversion.AddToScheme(scheme)
	authorizationinternalversion.AddToScheme(scheme)
	autoscalinginternalversion.AddToScheme(scheme)
	batchinternalversion.AddToScheme(scheme)
	certificatesinternalversion.AddToScheme(scheme)
	extensionsinternalversion.AddToScheme(scheme)
	networkinginternalversion.AddToScheme(scheme)
	policyinternalversion.AddToScheme(scheme)
	rbacinternalversion.AddToScheme(scheme)
	settingsinternalversion.AddToScheme(scheme)
	storageinternalversion.AddToScheme(scheme)

}
