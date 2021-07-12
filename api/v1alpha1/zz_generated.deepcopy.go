// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationField) DeepCopyInto(out *AuthenticationField) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationField.
func (in *AuthenticationField) DeepCopy() *AuthenticationField {
	if in == nil {
		return nil
	}
	out := new(AuthenticationField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaaSConnection) DeepCopyInto(out *DBaaSConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaaSConnection.
func (in *DBaaSConnection) DeepCopy() *DBaaSConnection {
	if in == nil {
		return nil
	}
	out := new(DBaaSConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DBaaSConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaaSConnectionList) DeepCopyInto(out *DBaaSConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DBaaSConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaaSConnectionList.
func (in *DBaaSConnectionList) DeepCopy() *DBaaSConnectionList {
	if in == nil {
		return nil
	}
	out := new(DBaaSConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DBaaSConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaaSConnectionSpec) DeepCopyInto(out *DBaaSConnectionSpec) {
	*out = *in
	out.InventoryRef = in.InventoryRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaaSConnectionSpec.
func (in *DBaaSConnectionSpec) DeepCopy() *DBaaSConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(DBaaSConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaaSConnectionStatus) DeepCopyInto(out *DBaaSConnectionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CredentialsRef != nil {
		in, out := &in.CredentialsRef, &out.CredentialsRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.ConnectionInfoRef != nil {
		in, out := &in.ConnectionInfoRef, &out.ConnectionInfoRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaaSConnectionStatus.
func (in *DBaaSConnectionStatus) DeepCopy() *DBaaSConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(DBaaSConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaaSInventory) DeepCopyInto(out *DBaaSInventory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaaSInventory.
func (in *DBaaSInventory) DeepCopy() *DBaaSInventory {
	if in == nil {
		return nil
	}
	out := new(DBaaSInventory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DBaaSInventory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaaSInventoryList) DeepCopyInto(out *DBaaSInventoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DBaaSInventory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaaSInventoryList.
func (in *DBaaSInventoryList) DeepCopy() *DBaaSInventoryList {
	if in == nil {
		return nil
	}
	out := new(DBaaSInventoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DBaaSInventoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaaSInventorySpec) DeepCopyInto(out *DBaaSInventorySpec) {
	*out = *in
	if in.CredentialsRef != nil {
		in, out := &in.CredentialsRef, &out.CredentialsRef
		*out = new(NamespacedName)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaaSInventorySpec.
func (in *DBaaSInventorySpec) DeepCopy() *DBaaSInventorySpec {
	if in == nil {
		return nil
	}
	out := new(DBaaSInventorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaaSInventoryStatus) DeepCopyInto(out *DBaaSInventoryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaaSInventoryStatus.
func (in *DBaaSInventoryStatus) DeepCopy() *DBaaSInventoryStatus {
	if in == nil {
		return nil
	}
	out := new(DBaaSInventoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaaSOperatorInventorySpec) DeepCopyInto(out *DBaaSOperatorInventorySpec) {
	*out = *in
	out.Provider = in.Provider
	in.DBaaSInventorySpec.DeepCopyInto(&out.DBaaSInventorySpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaaSOperatorInventorySpec.
func (in *DBaaSOperatorInventorySpec) DeepCopy() *DBaaSOperatorInventorySpec {
	if in == nil {
		return nil
	}
	out := new(DBaaSOperatorInventorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaaSProvider) DeepCopyInto(out *DBaaSProvider) {
	*out = *in
	out.Provider = in.Provider
	if in.AuthenticationFields != nil {
		in, out := &in.AuthenticationFields, &out.AuthenticationFields
		*out = make([]AuthenticationField, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaaSProvider.
func (in *DBaaSProvider) DeepCopy() *DBaaSProvider {
	if in == nil {
		return nil
	}
	out := new(DBaaSProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaaSProviderList) DeepCopyInto(out *DBaaSProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DBaaSProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaaSProviderList.
func (in *DBaaSProviderList) DeepCopy() *DBaaSProviderList {
	if in == nil {
		return nil
	}
	out := new(DBaaSProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaasTenant) DeepCopyInto(out *DBaasTenant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaasTenant.
func (in *DBaasTenant) DeepCopy() *DBaasTenant {
	if in == nil {
		return nil
	}
	out := new(DBaasTenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DBaasTenant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaasTenantList) DeepCopyInto(out *DBaasTenantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DBaasTenant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaasTenantList.
func (in *DBaasTenantList) DeepCopy() *DBaasTenantList {
	if in == nil {
		return nil
	}
	out := new(DBaasTenantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DBaasTenantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaasTenantSpec) DeepCopyInto(out *DBaasTenantSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaasTenantSpec.
func (in *DBaasTenantSpec) DeepCopy() *DBaasTenantSpec {
	if in == nil {
		return nil
	}
	out := new(DBaasTenantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBaasTenantStatus) DeepCopyInto(out *DBaasTenantStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBaasTenantStatus.
func (in *DBaasTenantStatus) DeepCopy() *DBaasTenantStatus {
	if in == nil {
		return nil
	}
	out := new(DBaasTenantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseProvider) DeepCopyInto(out *DatabaseProvider) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseProvider.
func (in *DatabaseProvider) DeepCopy() *DatabaseProvider {
	if in == nil {
		return nil
	}
	out := new(DatabaseProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	if in.InstanceInfo != nil {
		in, out := &in.InstanceInfo, &out.InstanceInfo
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedName) DeepCopyInto(out *NamespacedName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedName.
func (in *NamespacedName) DeepCopy() *NamespacedName {
	if in == nil {
		return nil
	}
	out := new(NamespacedName)
	in.DeepCopyInto(out)
	return out
}
