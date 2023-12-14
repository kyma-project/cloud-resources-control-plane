/*
Copyright 2023.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Enum=Regional;Zonal
type AwsFileSystemType string

const (
	AwsFileSystemTypeRegional = "Regional"
	AwsFileSystemTypeZonal    = "Zonal"
)

// +kubebuilder:validation:Enum=Enhanced;Bursting
type AwsThroughputMode string

const (
	AwsThroughputModeEnhanced = "Enhanced"
	AwsThroughputModeBursting = "Bursting"
)

// NfsShareSpec defines the desired state of NfsShare
type NfsShareSpec struct {
	// +kubebuilder:validation:Required
	Kyma string `json:"kyma"`

	// +optional
	Gcp *NfsShareGcp `json:"gcp,omitempty"`

	// +optional
	Azure *NfsShareAzure `json:"azure,omitempty"`

	// +optional
	Aws *NfsShareAws `json:"aws,omitempty"`
}

type NfsShareGcp struct {
}

type NfsShareAzure struct {
}

type NfsShareAws struct {
	Type       AwsFileSystemType `json:"type,omitempty"`
	Throughput AwsThroughputMode `json:"throughput,omitempty"`
}

// NfsShareStatus defines the observed state of NfsShare
type NfsShareStatus struct {
	State StatusState `json:"state,omitempty"`

	// +optional
	Scope *ScopeX `json:"scope,omitempty"`

	// List of status conditions to indicate the status of a Peering.
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NfsShare is the Schema for the nfsshares API
type NfsShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NfsShareSpec   `json:"spec,omitempty"`
	Status NfsShareStatus `json:"status,omitempty"`
}

func (in *NfsShare) Kyma() string {
	return in.Spec.Kyma
}

func (in *NfsShare) Scope() *ScopeX {
	return in.Status.Scope
}
func (in *NfsShare) SetScope(scope *ScopeX) {
	in.Status.Scope = scope
}

//+kubebuilder:object:root=true

// NfsShareList contains a list of NfsShare
type NfsShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NfsShare `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NfsShare{}, &NfsShareList{})
}
