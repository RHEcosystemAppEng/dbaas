/*
Copyright 2021.

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

// This object defines the desired state of a DBaaSInventory object.
type DBaaSOperatorInventorySpec struct {
	// A reference to a DBaaSProvider custom resource (CR).
	ProviderRef NamespacedName `json:"providerRef"`

	// The properties that will be copied into the provider’s inventory.
	DBaaSInventorySpec `json:",inline"`

	// The policy for this inventory.
	Policy *DBaaSInventoryPolicy `json:"policy,omitempty"`
}

//+kubebuilder:storageversion
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// The schema for the DBaaSInventory API.
// Inventory objects must be created in a valid namespace, determined by the existence of a DBaaSPolicy object.
//+operator-sdk:csv:customresourcedefinitions:displayName="Provider Account"
type DBaaSInventory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DBaaSOperatorInventorySpec `json:"spec,omitempty"`
	Status DBaaSInventoryStatus       `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Contains a list of DBaaSInventories.
type DBaaSInventoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBaaSInventory `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DBaaSInventory{}, &DBaaSInventoryList{})
}