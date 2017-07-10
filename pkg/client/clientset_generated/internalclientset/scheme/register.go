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

package scheme

import (
	announced "k8s.io/apimachinery/pkg/apimachinery/announced"
	registered "k8s.io/apimachinery/pkg/apimachinery/registered"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	core "gitlab.com/sankish/magudi/pkg/api/install"
	admissionregistration "gitlab.com/sankish/magudi/pkg/apis/admissionregistration/install"
	apps "gitlab.com/sankish/magudi/pkg/apis/apps/install"
	authentication "gitlab.com/sankish/magudi/pkg/apis/authentication/install"
	authorization "gitlab.com/sankish/magudi/pkg/apis/authorization/install"
	autoscaling "gitlab.com/sankish/magudi/pkg/apis/autoscaling/install"
	batch "gitlab.com/sankish/magudi/pkg/apis/batch/install"
	certificates "gitlab.com/sankish/magudi/pkg/apis/certificates/install"
	extensions "gitlab.com/sankish/magudi/pkg/apis/extensions/install"
	networking "gitlab.com/sankish/magudi/pkg/apis/networking/install"
	policy "gitlab.com/sankish/magudi/pkg/apis/policy/install"
	rbac "gitlab.com/sankish/magudi/pkg/apis/rbac/install"
	settings "gitlab.com/sankish/magudi/pkg/apis/settings/install"
	storage "gitlab.com/sankish/magudi/pkg/apis/storage/install"
	os "os"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)

var Registry = registered.NewOrDie(os.Getenv("KUBE_API_VERSIONS"))
var GroupFactoryRegistry = make(announced.APIGroupFactoryRegistry)

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	Install(GroupFactoryRegistry, Registry, Scheme)
}

// Install registers the API group and adds types to a scheme
func Install(groupFactoryRegistry announced.APIGroupFactoryRegistry, registry *registered.APIRegistrationManager, scheme *runtime.Scheme) {
	admissionregistration.Install(groupFactoryRegistry, registry, scheme)
	core.Install(groupFactoryRegistry, registry, scheme)
	apps.Install(groupFactoryRegistry, registry, scheme)
	authentication.Install(groupFactoryRegistry, registry, scheme)
	authorization.Install(groupFactoryRegistry, registry, scheme)
	autoscaling.Install(groupFactoryRegistry, registry, scheme)
	batch.Install(groupFactoryRegistry, registry, scheme)
	certificates.Install(groupFactoryRegistry, registry, scheme)
	extensions.Install(groupFactoryRegistry, registry, scheme)
	networking.Install(groupFactoryRegistry, registry, scheme)
	policy.Install(groupFactoryRegistry, registry, scheme)
	rbac.Install(groupFactoryRegistry, registry, scheme)
	settings.Install(groupFactoryRegistry, registry, scheme)
	storage.Install(groupFactoryRegistry, registry, scheme)

	ExtraInstall(groupFactoryRegistry, registry, scheme)
}
