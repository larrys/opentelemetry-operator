// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package collector

import (
	"maps"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/open-telemetry/opentelemetry-operator/apis/v1alpha1"
	"github.com/open-telemetry/opentelemetry-operator/apis/v1beta1"
	"github.com/open-telemetry/opentelemetry-operator/internal/manifests"
)

// TargetAllocator builds the TargetAllocator CR for the given instance.
func TargetAllocator(params manifests.Params) (*v1alpha1.TargetAllocator, error) {
	taSpec := params.OtelCol.Spec.TargetAllocator
	if !taSpec.Enabled {
		return nil, nil
	}

	// setting all the labels normally here leads to some undesirable results, like the name label being wrong
	// instead, only set managed-by and leave everything else as-is
	labels := maps.Clone(params.OtelCol.Labels)
	if labels == nil {
		labels = make(map[string]string, 1)
	}
	labels["app.kubernetes.io/managed-by"] = "opentelemetry-operator"

	return &v1alpha1.TargetAllocator{
		ObjectMeta: metav1.ObjectMeta{
			Name:        params.OtelCol.Name,
			Namespace:   params.OtelCol.Namespace,
			Annotations: params.OtelCol.Annotations,
			Labels:      labels,
		},
		Spec: v1alpha1.TargetAllocatorSpec{
			OpenTelemetryCommonFields: v1beta1.OpenTelemetryCommonFields{
				Replicas:                  taSpec.Replicas,
				NodeSelector:              taSpec.NodeSelector,
				Resources:                 taSpec.Resources,
				ServiceAccount:            taSpec.ServiceAccount,
				Image:                     taSpec.Image,
				Affinity:                  taSpec.Affinity,
				SecurityContext:           taSpec.SecurityContext,
				PodSecurityContext:        taSpec.PodSecurityContext,
				TopologySpreadConstraints: taSpec.TopologySpreadConstraints,
				Tolerations:               taSpec.Tolerations,
				Env:                       taSpec.Env,
				PodAnnotations:            params.OtelCol.Spec.PodAnnotations,
				PodDisruptionBudget:       taSpec.PodDisruptionBudget,
			},
			AllocationStrategy:           taSpec.AllocationStrategy,
			FilterStrategy:               taSpec.FilterStrategy,
			PrometheusCR:                 taSpec.PrometheusCR,
			Observability:                taSpec.Observability,
			CollectorNotReadyGracePeriod: taSpec.CollectorNotReadyGracePeriod,
		},
	}, nil
}
