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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IpRangeSpec defines the desired state of IpRange
type IpRangeSpec struct {
	// +kubebuilder:validation:Required
	Kyma string `json:"kyma"`

	// +kubebuilder:validation:Required
	Scope *ScopeRef `json:"scope"`

	// +kubebuilder:validation:Required
	Cidr string `json:"cidr"`
}

// IpRangeStatus defines the observed state of IpRange
type IpRangeStatus struct {
	State StatusState `json:"state,omitempty"`

	// List of status conditions to indicate the status of a Peering.
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// IpRange is the Schema for the ipranges API
type IpRange struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IpRangeSpec   `json:"spec,omitempty"`
	Status IpRangeStatus `json:"status,omitempty"`
}

func (in *IpRange) KymaName() string {
	return in.Spec.Kyma
}

func (in *IpRange) ScopeRef() *ScopeRef {
	return in.Spec.Scope
}

func (in *IpRange) SetScopeRef(scopeRef *ScopeRef) {
	in.Spec.Scope = scopeRef
}

func (in *IpRange) Conditions() *[]metav1.Condition {
	return &in.Status.Conditions
}

//+kubebuilder:object:root=true

// IpRangeList contains a list of IpRange
type IpRangeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IpRange `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IpRange{}, &IpRangeList{})
}
