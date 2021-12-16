//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

package v1beta1

import (
	"github.com/hetznercloud/hcloud-go/hcloud"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachine) DeepCopyInto(out *HCloudMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachine.
func (in *HCloudMachine) DeepCopy() *HCloudMachine {
	if in == nil {
		return nil
	}
	out := new(HCloudMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HCloudMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachineList) DeepCopyInto(out *HCloudMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HCloudMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachineList.
func (in *HCloudMachineList) DeepCopy() *HCloudMachineList {
	if in == nil {
		return nil
	}
	out := new(HCloudMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HCloudMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachineSpec) DeepCopyInto(out *HCloudMachineSpec) {
	*out = *in
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = make([]SSHKey, len(*in))
		copy(*out, *in)
	}
	if in.PlacementGroupName != nil {
		in, out := &in.PlacementGroupName, &out.PlacementGroupName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachineSpec.
func (in *HCloudMachineSpec) DeepCopy() *HCloudMachineSpec {
	if in == nil {
		return nil
	}
	out := new(HCloudMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachineStatus) DeepCopyInto(out *HCloudMachineStatus) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
	if in.InstanceState != nil {
		in, out := &in.InstanceState, &out.InstanceState
		*out = new(hcloud.ServerStatus)
		**out = **in
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachineStatus.
func (in *HCloudMachineStatus) DeepCopy() *HCloudMachineStatus {
	if in == nil {
		return nil
	}
	out := new(HCloudMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachineTemplate) DeepCopyInto(out *HCloudMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachineTemplate.
func (in *HCloudMachineTemplate) DeepCopy() *HCloudMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(HCloudMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HCloudMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachineTemplateList) DeepCopyInto(out *HCloudMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HCloudMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachineTemplateList.
func (in *HCloudMachineTemplateList) DeepCopy() *HCloudMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(HCloudMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HCloudMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachineTemplateResource) DeepCopyInto(out *HCloudMachineTemplateResource) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachineTemplateResource.
func (in *HCloudMachineTemplateResource) DeepCopy() *HCloudMachineTemplateResource {
	if in == nil {
		return nil
	}
	out := new(HCloudMachineTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachineTemplateSpec) DeepCopyInto(out *HCloudMachineTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachineTemplateSpec.
func (in *HCloudMachineTemplateSpec) DeepCopy() *HCloudMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(HCloudMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudNetworkSpec) DeepCopyInto(out *HCloudNetworkSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudNetworkSpec.
func (in *HCloudNetworkSpec) DeepCopy() *HCloudNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(HCloudNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudPlacementGroupSpec) DeepCopyInto(out *HCloudPlacementGroupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudPlacementGroupSpec.
func (in *HCloudPlacementGroupSpec) DeepCopy() *HCloudPlacementGroupSpec {
	if in == nil {
		return nil
	}
	out := new(HCloudPlacementGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudPlacementGroupStatus) DeepCopyInto(out *HCloudPlacementGroupStatus) {
	*out = *in
	if in.Server != nil {
		in, out := &in.Server, &out.Server
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudPlacementGroupStatus.
func (in *HCloudPlacementGroupStatus) DeepCopy() *HCloudPlacementGroupStatus {
	if in == nil {
		return nil
	}
	out := new(HCloudPlacementGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalMachine) DeepCopyInto(out *HetznerBareMetalMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalMachine.
func (in *HetznerBareMetalMachine) DeepCopy() *HetznerBareMetalMachine {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerBareMetalMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalMachineHost) DeepCopyInto(out *HetznerBareMetalMachineHost) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalMachineHost.
func (in *HetznerBareMetalMachineHost) DeepCopy() *HetznerBareMetalMachineHost {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalMachineHost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerBareMetalMachineHost) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalMachineHostList) DeepCopyInto(out *HetznerBareMetalMachineHostList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HetznerBareMetalMachineHost, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalMachineHostList.
func (in *HetznerBareMetalMachineHostList) DeepCopy() *HetznerBareMetalMachineHostList {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalMachineHostList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerBareMetalMachineHostList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalMachineHostSpec) DeepCopyInto(out *HetznerBareMetalMachineHostSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalMachineHostSpec.
func (in *HetznerBareMetalMachineHostSpec) DeepCopy() *HetznerBareMetalMachineHostSpec {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalMachineHostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalMachineHostStatus) DeepCopyInto(out *HetznerBareMetalMachineHostStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalMachineHostStatus.
func (in *HetznerBareMetalMachineHostStatus) DeepCopy() *HetznerBareMetalMachineHostStatus {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalMachineHostStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalMachineList) DeepCopyInto(out *HetznerBareMetalMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HetznerBareMetalMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalMachineList.
func (in *HetznerBareMetalMachineList) DeepCopy() *HetznerBareMetalMachineList {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerBareMetalMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalMachineSpec) DeepCopyInto(out *HetznerBareMetalMachineSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalMachineSpec.
func (in *HetznerBareMetalMachineSpec) DeepCopy() *HetznerBareMetalMachineSpec {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalMachineStatus) DeepCopyInto(out *HetznerBareMetalMachineStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalMachineStatus.
func (in *HetznerBareMetalMachineStatus) DeepCopy() *HetznerBareMetalMachineStatus {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalMachineTemplate) DeepCopyInto(out *HetznerBareMetalMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalMachineTemplate.
func (in *HetznerBareMetalMachineTemplate) DeepCopy() *HetznerBareMetalMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerBareMetalMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalMachineTemplateList) DeepCopyInto(out *HetznerBareMetalMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HetznerBareMetalMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalMachineTemplateList.
func (in *HetznerBareMetalMachineTemplateList) DeepCopy() *HetznerBareMetalMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerBareMetalMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalMachineTemplateSpec) DeepCopyInto(out *HetznerBareMetalMachineTemplateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalMachineTemplateSpec.
func (in *HetznerBareMetalMachineTemplateSpec) DeepCopy() *HetznerBareMetalMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalMachineTemplateStatus) DeepCopyInto(out *HetznerBareMetalMachineTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalMachineTemplateStatus.
func (in *HetznerBareMetalMachineTemplateStatus) DeepCopy() *HetznerBareMetalMachineTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalMachineTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalRemediationTemplate) DeepCopyInto(out *HetznerBareMetalRemediationTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalRemediationTemplate.
func (in *HetznerBareMetalRemediationTemplate) DeepCopy() *HetznerBareMetalRemediationTemplate {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalRemediationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerBareMetalRemediationTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalRemediationTemplateList) DeepCopyInto(out *HetznerBareMetalRemediationTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HetznerBareMetalRemediationTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalRemediationTemplateList.
func (in *HetznerBareMetalRemediationTemplateList) DeepCopy() *HetznerBareMetalRemediationTemplateList {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalRemediationTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerBareMetalRemediationTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalRemediationTemplateSpec) DeepCopyInto(out *HetznerBareMetalRemediationTemplateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalRemediationTemplateSpec.
func (in *HetznerBareMetalRemediationTemplateSpec) DeepCopy() *HetznerBareMetalRemediationTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalRemediationTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerBareMetalRemediationTemplateStatus) DeepCopyInto(out *HetznerBareMetalRemediationTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerBareMetalRemediationTemplateStatus.
func (in *HetznerBareMetalRemediationTemplateStatus) DeepCopy() *HetznerBareMetalRemediationTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(HetznerBareMetalRemediationTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerCluster) DeepCopyInto(out *HetznerCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerCluster.
func (in *HetznerCluster) DeepCopy() *HetznerCluster {
	if in == nil {
		return nil
	}
	out := new(HetznerCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerClusterList) DeepCopyInto(out *HetznerClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HetznerCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerClusterList.
func (in *HetznerClusterList) DeepCopy() *HetznerClusterList {
	if in == nil {
		return nil
	}
	out := new(HetznerClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerClusterSpec) DeepCopyInto(out *HetznerClusterSpec) {
	*out = *in
	out.HCloudNetwork = in.HCloudNetwork
	if in.ControlPlaneRegions != nil {
		in, out := &in.ControlPlaneRegions, &out.ControlPlaneRegions
		*out = make([]Region, len(*in))
		copy(*out, *in)
	}
	in.SSHKeys.DeepCopyInto(&out.SSHKeys)
	if in.ControlPlaneEndpoint != nil {
		in, out := &in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint
		*out = new(apiv1beta1.APIEndpoint)
		**out = **in
	}
	in.ControlPlaneLoadBalancer.DeepCopyInto(&out.ControlPlaneLoadBalancer)
	if in.HCloudPlacementGroup != nil {
		in, out := &in.HCloudPlacementGroup, &out.HCloudPlacementGroup
		*out = make([]HCloudPlacementGroupSpec, len(*in))
		copy(*out, *in)
	}
	out.HetznerSecret = in.HetznerSecret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerClusterSpec.
func (in *HetznerClusterSpec) DeepCopy() *HetznerClusterSpec {
	if in == nil {
		return nil
	}
	out := new(HetznerClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerClusterStatus) DeepCopyInto(out *HetznerClusterStatus) {
	*out = *in
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(NetworkStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ControlPlaneLoadBalancer != nil {
		in, out := &in.ControlPlaneLoadBalancer, &out.ControlPlaneLoadBalancer
		*out = new(LoadBalancerStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.HCloudPlacementGroup != nil {
		in, out := &in.HCloudPlacementGroup, &out.HCloudPlacementGroup
		*out = make([]HCloudPlacementGroupStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(apiv1beta1.FailureDomains, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerClusterStatus.
func (in *HetznerClusterStatus) DeepCopy() *HetznerClusterStatus {
	if in == nil {
		return nil
	}
	out := new(HetznerClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerClusterTemplate) DeepCopyInto(out *HetznerClusterTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerClusterTemplate.
func (in *HetznerClusterTemplate) DeepCopy() *HetznerClusterTemplate {
	if in == nil {
		return nil
	}
	out := new(HetznerClusterTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerClusterTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerClusterTemplateList) DeepCopyInto(out *HetznerClusterTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HetznerClusterTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerClusterTemplateList.
func (in *HetznerClusterTemplateList) DeepCopy() *HetznerClusterTemplateList {
	if in == nil {
		return nil
	}
	out := new(HetznerClusterTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerClusterTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerClusterTemplateResource) DeepCopyInto(out *HetznerClusterTemplateResource) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerClusterTemplateResource.
func (in *HetznerClusterTemplateResource) DeepCopy() *HetznerClusterTemplateResource {
	if in == nil {
		return nil
	}
	out := new(HetznerClusterTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerClusterTemplateSpec) DeepCopyInto(out *HetznerClusterTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerClusterTemplateSpec.
func (in *HetznerClusterTemplateSpec) DeepCopy() *HetznerClusterTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(HetznerClusterTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerSSHKeys) DeepCopyInto(out *HetznerSSHKeys) {
	*out = *in
	if in.HCloud != nil {
		in, out := &in.HCloud, &out.HCloud
		*out = make([]SSHKey, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerSSHKeys.
func (in *HetznerSSHKeys) DeepCopy() *HetznerSSHKeys {
	if in == nil {
		return nil
	}
	out := new(HetznerSSHKeys)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerSecretKeyRef) DeepCopyInto(out *HetznerSecretKeyRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerSecretKeyRef.
func (in *HetznerSecretKeyRef) DeepCopy() *HetznerSecretKeyRef {
	if in == nil {
		return nil
	}
	out := new(HetznerSecretKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerSecretRef) DeepCopyInto(out *HetznerSecretRef) {
	*out = *in
	out.Key = in.Key
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerSecretRef.
func (in *HetznerSecretRef) DeepCopy() *HetznerSecretRef {
	if in == nil {
		return nil
	}
	out := new(HetznerSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerServiceSpec) DeepCopyInto(out *LoadBalancerServiceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerServiceSpec.
func (in *LoadBalancerServiceSpec) DeepCopy() *LoadBalancerServiceSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSpec) DeepCopyInto(out *LoadBalancerSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ExtraServices != nil {
		in, out := &in.ExtraServices, &out.ExtraServices
		*out = make([]LoadBalancerServiceSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSpec.
func (in *LoadBalancerSpec) DeepCopy() *LoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerStatus) DeepCopyInto(out *LoadBalancerStatus) {
	*out = *in
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerStatus.
func (in *LoadBalancerStatus) DeepCopy() *LoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkStatus) DeepCopyInto(out *NetworkStatus) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AttachedServers != nil {
		in, out := &in.AttachedServers, &out.AttachedServers
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkStatus.
func (in *NetworkStatus) DeepCopy() *NetworkStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHKey) DeepCopyInto(out *SSHKey) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHKey.
func (in *SSHKey) DeepCopy() *SSHKey {
	if in == nil {
		return nil
	}
	out := new(SSHKey)
	in.DeepCopyInto(out)
	return out
}
