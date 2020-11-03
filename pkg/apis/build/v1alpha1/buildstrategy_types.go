// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// LabelBuildStrategyName is a label key for defining the build strategy name
	LabelBuildStrategyName = "buildstrategy.build.dev/name"

	// LabelBuildStrategyGeneration is a label key for defining the build strategy generation
	LabelBuildStrategyGeneration = "buildstrategy.build.dev/generation"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BuildStrategy is the Schema representing a strategy in the namespace scope to build images from source code.
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=buildstrategies,scope=Namespaced,shortName=bs;bss
type BuildStrategy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BuildStrategySpec   `json:"spec,omitempty"`
	Status BuildStrategyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BuildStrategyList contains a list of BuildStrategy
type BuildStrategyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BuildStrategy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BuildStrategy{}, &BuildStrategyList{})
}
