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

package clientset

// These imports are the API groups the client will support.
import (
	"fmt"

	"gitlab.com/rioos/magudi/pkg/api"
	_ "gitlab.com/rioos/magudi/pkg/api/install"
	_ "gitlab.com/rioos/magudi/pkg/apis/apps/install"
	_ "gitlab.com/rioos/magudi/pkg/apis/authentication/install"
	_ "gitlab.com/rioos/magudi/pkg/apis/authorization/install"
	_ "gitlab.com/rioos/magudi/pkg/apis/autoscaling/install"
	_ "gitlab.com/rioos/magudi/pkg/apis/batch/install"
	_ "gitlab.com/rioos/magudi/pkg/apis/certificates/install"
	_ "gitlab.com/rioos/magudi/pkg/apis/extensions/install"
	_ "gitlab.com/rioos/magudi/pkg/apis/policy/install"
	_ "gitlab.com/rioos/magudi/pkg/apis/rbac/install"
	_ "gitlab.com/rioos/magudi/pkg/apis/settings/install"
	_ "gitlab.com/rioos/magudi/pkg/apis/storage/install"
)

func init() {
	if missingVersions := api.Registry.ValidateEnvRequestedVersions(); len(missingVersions) != 0 {
		panic(fmt.Sprintf("KUBE_API_VERSIONS contains versions that are not installed: %q.", missingVersions))
	}
}
