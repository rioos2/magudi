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

package v2alpha1

import (
	autoscalingv2alpha1 "k8s.io/api/autoscaling/v2alpha1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"gitlab.com/rioos/magudi/pkg/apis/autoscaling"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

func SetDefaults_HorizontalPodAutoscaler(obj *autoscalingv2alpha1.HorizontalPodAutoscaler) {
	if obj.Spec.MinReplicas == nil {
		minReplicas := int32(1)
		obj.Spec.MinReplicas = &minReplicas
	}

	if len(obj.Spec.Metrics) == 0 {
		utilizationDefaultVal := int32(autoscaling.DefaultCPUUtilization)
		obj.Spec.Metrics = []autoscalingv2alpha1.MetricSpec{
			{
				Type: autoscalingv2alpha1.ResourceMetricSourceType,
				Resource: &autoscalingv2alpha1.ResourceMetricSource{
					Name: v1.ResourceCPU,
					TargetAverageUtilization: &utilizationDefaultVal,
				},
			},
		}
	}
}
