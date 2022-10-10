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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DBaaSPolicySpec enables admin capabilities within a namespace and sets default inventory policy.
// Policy defaults can be overridden on a per-inventory basis.
type DBaaSPolicySpec struct {
	DBaaSInventoryPolicy `json:",inline"`
}

// DBaaSInventoryPolicy sets inventory policy
type DBaaSInventoryPolicy struct {
	// Disable provisioning against inventory accounts
	DisableProvisions *bool `json:"disableProvisions,omitempty"`

	// Namespaces where DBaaSConnections/DBaaSInstances are allowed to reference a policy's inventories.
	// Each inventory can individually override this. Use "*" to allow all namespaces.
	// If not set in either the policy or inventory object, connections will only be allowed in the inventory's namespace.
	ConnectionNamespaces *[]string `json:"connectionNamespaces,omitempty"`

	// Use a label selector to determine namespaces where DBaaSConnections/DBaaSInstances are allowed to reference a policy's inventories.
	// Each inventory can individually override this. A label selector is a label query over a set of resources. The result of matchLabels and
	// matchExpressions are ANDed. An empty label selector matches all objects. A null
	// label selector matches no objects.
	ConnectionNsSelector *metav1.LabelSelector `json:"connectionNsSelector,omitempty"`
}

// DBaaSPolicyStatus defines the observed state of DBaaSPolicy
type DBaaSPolicyStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Active",type=string,JSONPath=`.status.conditions[0].status`
//+kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// DBaaSPolicy enables admin capabilities within a namespace and sets default inventory policy.
// Policy defaults can be overridden on a per-inventory basis.
//+operator-sdk:csv:customresourcedefinitions:displayName="Provider Account Policy"
type DBaaSPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DBaaSPolicySpec   `json:"spec,omitempty"`
	Status DBaaSPolicyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DBaaSPolicyList contains a list of DBaaSPolicy
type DBaaSPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBaaSPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DBaaSPolicy{}, &DBaaSPolicyList{})
}
