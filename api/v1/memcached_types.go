/*
Copyright 2025.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConditionReady is the condition type for the Ready condition
const (
	ConditionReady       = "Ready"
	ConditionAvailable   = "Available"
	ConditionDegraded    = "Degraded"
	ConditionDowngraded  = "Downgraded"
	ConditionProgressing = "Progressing"
)

// MemcachedFinalizer is the name of the finalizer
const MemcachedFinalizer = "cache.devspace.com/finalizer"

// MemcachedSpec defines the desired state of Memcached
type MemcachedSpec struct {
	// Size is the size of the memcached deployment
	Size int32 `json:"size"`
}

// MemcachedStatus defines the observed state of Memcached
type MemcachedStatus struct {
	// Conditions store the status conditions of the Memcached instances
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Size",type="integer",JSONPath=".spec.size"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// Memcached is the Schema for the memcacheds API
type Memcached struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemcachedSpec   `json:"spec,omitempty"`
	Status MemcachedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemcachedList contains a list of Memcached
type MemcachedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Memcached `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Memcached{}, &MemcachedList{})
}
