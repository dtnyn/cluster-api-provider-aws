//go:build !ignore_autogenerated

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

package v1beta2

import (
	"k8s.io/apimachinery/pkg/runtime"
	apiv1beta2 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
	"sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFargateProfile) DeepCopyInto(out *AWSFargateProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFargateProfile.
func (in *AWSFargateProfile) DeepCopy() *AWSFargateProfile {
	if in == nil {
		return nil
	}
	out := new(AWSFargateProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSFargateProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFargateProfileList) DeepCopyInto(out *AWSFargateProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSFargateProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFargateProfileList.
func (in *AWSFargateProfileList) DeepCopy() *AWSFargateProfileList {
	if in == nil {
		return nil
	}
	out := new(AWSFargateProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSFargateProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSLaunchTemplate) DeepCopyInto(out *AWSLaunchTemplate) {
	*out = *in
	in.AMI.DeepCopyInto(&out.AMI)
	if in.RootVolume != nil {
		in, out := &in.RootVolume, &out.RootVolume
		*out = new(apiv1beta2.Volume)
		(*in).DeepCopyInto(*out)
	}
	if in.SSHKeyName != nil {
		in, out := &in.SSHKeyName, &out.SSHKeyName
		*out = new(string)
		**out = **in
	}
	if in.VersionNumber != nil {
		in, out := &in.VersionNumber, &out.VersionNumber
		*out = new(int64)
		**out = **in
	}
	if in.AdditionalSecurityGroups != nil {
		in, out := &in.AdditionalSecurityGroups, &out.AdditionalSecurityGroups
		*out = make([]apiv1beta2.AWSResourceReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SpotMarketOptions != nil {
		in, out := &in.SpotMarketOptions, &out.SpotMarketOptions
		*out = new(apiv1beta2.SpotMarketOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceMetadataOptions != nil {
		in, out := &in.InstanceMetadataOptions, &out.InstanceMetadataOptions
		*out = new(apiv1beta2.InstanceMetadataOptions)
		**out = **in
	}
	if in.PrivateDNSName != nil {
		in, out := &in.PrivateDNSName, &out.PrivateDNSName
		*out = new(apiv1beta2.PrivateDNSName)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSLaunchTemplate.
func (in *AWSLaunchTemplate) DeepCopy() *AWSLaunchTemplate {
	if in == nil {
		return nil
	}
	out := new(AWSLaunchTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachinePool) DeepCopyInto(out *AWSMachinePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachinePool.
func (in *AWSMachinePool) DeepCopy() *AWSMachinePool {
	if in == nil {
		return nil
	}
	out := new(AWSMachinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSMachinePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachinePoolInstanceStatus) DeepCopyInto(out *AWSMachinePoolInstanceStatus) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachinePoolInstanceStatus.
func (in *AWSMachinePoolInstanceStatus) DeepCopy() *AWSMachinePoolInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(AWSMachinePoolInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachinePoolList) DeepCopyInto(out *AWSMachinePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSMachinePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachinePoolList.
func (in *AWSMachinePoolList) DeepCopy() *AWSMachinePoolList {
	if in == nil {
		return nil
	}
	out := new(AWSMachinePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSMachinePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachinePoolSpec) DeepCopyInto(out *AWSMachinePoolSpec) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AvailabilityZoneSubnetType != nil {
		in, out := &in.AvailabilityZoneSubnetType, &out.AvailabilityZoneSubnetType
		*out = new(AZSubnetType)
		**out = **in
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]apiv1beta2.AWSResourceReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalTags != nil {
		in, out := &in.AdditionalTags, &out.AdditionalTags
		*out = make(apiv1beta2.Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.AWSLaunchTemplate.DeepCopyInto(&out.AWSLaunchTemplate)
	if in.MixedInstancesPolicy != nil {
		in, out := &in.MixedInstancesPolicy, &out.MixedInstancesPolicy
		*out = new(MixedInstancesPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.ProviderIDList != nil {
		in, out := &in.ProviderIDList, &out.ProviderIDList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.DefaultCoolDown = in.DefaultCoolDown
	out.DefaultInstanceWarmup = in.DefaultInstanceWarmup
	if in.RefreshPreferences != nil {
		in, out := &in.RefreshPreferences, &out.RefreshPreferences
		*out = new(RefreshPreferences)
		(*in).DeepCopyInto(*out)
	}
	if in.SuspendProcesses != nil {
		in, out := &in.SuspendProcesses, &out.SuspendProcesses
		*out = new(SuspendProcessesTypes)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachinePoolSpec.
func (in *AWSMachinePoolSpec) DeepCopy() *AWSMachinePoolSpec {
	if in == nil {
		return nil
	}
	out := new(AWSMachinePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachinePoolStatus) DeepCopyInto(out *AWSMachinePoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]AWSMachinePoolInstanceStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LaunchTemplateVersion != nil {
		in, out := &in.LaunchTemplateVersion, &out.LaunchTemplateVersion
		*out = new(string)
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
	if in.ASGStatus != nil {
		in, out := &in.ASGStatus, &out.ASGStatus
		*out = new(ASGStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachinePoolStatus.
func (in *AWSMachinePoolStatus) DeepCopy() *AWSMachinePoolStatus {
	if in == nil {
		return nil
	}
	out := new(AWSMachinePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedMachinePool) DeepCopyInto(out *AWSManagedMachinePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedMachinePool.
func (in *AWSManagedMachinePool) DeepCopy() *AWSManagedMachinePool {
	if in == nil {
		return nil
	}
	out := new(AWSManagedMachinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSManagedMachinePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedMachinePoolList) DeepCopyInto(out *AWSManagedMachinePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSManagedMachinePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedMachinePoolList.
func (in *AWSManagedMachinePoolList) DeepCopy() *AWSManagedMachinePoolList {
	if in == nil {
		return nil
	}
	out := new(AWSManagedMachinePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSManagedMachinePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedMachinePoolSpec) DeepCopyInto(out *AWSManagedMachinePoolSpec) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AvailabilityZoneSubnetType != nil {
		in, out := &in.AvailabilityZoneSubnetType, &out.AvailabilityZoneSubnetType
		*out = new(AZSubnetType)
		**out = **in
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalTags != nil {
		in, out := &in.AdditionalTags, &out.AdditionalTags
		*out = make(apiv1beta2.Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RoleAdditionalPolicies != nil {
		in, out := &in.RoleAdditionalPolicies, &out.RoleAdditionalPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AMIVersion != nil {
		in, out := &in.AMIVersion, &out.AMIVersion
		*out = new(string)
		**out = **in
	}
	if in.AMIType != nil {
		in, out := &in.AMIType, &out.AMIType
		*out = new(ManagedMachineAMIType)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make(Taints, len(*in))
		copy(*out, *in)
	}
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(int32)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.Scaling != nil {
		in, out := &in.Scaling, &out.Scaling
		*out = new(ManagedMachinePoolScaling)
		(*in).DeepCopyInto(*out)
	}
	if in.RemoteAccess != nil {
		in, out := &in.RemoteAccess, &out.RemoteAccess
		*out = new(ManagedRemoteAccess)
		(*in).DeepCopyInto(*out)
	}
	if in.ProviderIDList != nil {
		in, out := &in.ProviderIDList, &out.ProviderIDList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CapacityType != nil {
		in, out := &in.CapacityType, &out.CapacityType
		*out = new(ManagedMachinePoolCapacityType)
		**out = **in
	}
	if in.UpdateConfig != nil {
		in, out := &in.UpdateConfig, &out.UpdateConfig
		*out = new(UpdateConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AWSLaunchTemplate != nil {
		in, out := &in.AWSLaunchTemplate, &out.AWSLaunchTemplate
		*out = new(AWSLaunchTemplate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedMachinePoolSpec.
func (in *AWSManagedMachinePoolSpec) DeepCopy() *AWSManagedMachinePoolSpec {
	if in == nil {
		return nil
	}
	out := new(AWSManagedMachinePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedMachinePoolStatus) DeepCopyInto(out *AWSManagedMachinePoolStatus) {
	*out = *in
	if in.LaunchTemplateID != nil {
		in, out := &in.LaunchTemplateID, &out.LaunchTemplateID
		*out = new(string)
		**out = **in
	}
	if in.LaunchTemplateVersion != nil {
		in, out := &in.LaunchTemplateVersion, &out.LaunchTemplateVersion
		*out = new(string)
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
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedMachinePoolStatus.
func (in *AWSManagedMachinePoolStatus) DeepCopy() *AWSManagedMachinePoolStatus {
	if in == nil {
		return nil
	}
	out := new(AWSManagedMachinePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingGroup) DeepCopyInto(out *AutoScalingGroup) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(apiv1beta2.Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DesiredCapacity != nil {
		in, out := &in.DesiredCapacity, &out.DesiredCapacity
		*out = new(int32)
		**out = **in
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.DefaultCoolDown = in.DefaultCoolDown
	out.DefaultInstanceWarmup = in.DefaultInstanceWarmup
	if in.MixedInstancesPolicy != nil {
		in, out := &in.MixedInstancesPolicy, &out.MixedInstancesPolicy
		*out = new(MixedInstancesPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]apiv1beta2.Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CurrentlySuspendProcesses != nil {
		in, out := &in.CurrentlySuspendProcesses, &out.CurrentlySuspendProcesses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingGroup.
func (in *AutoScalingGroup) DeepCopy() *AutoScalingGroup {
	if in == nil {
		return nil
	}
	out := new(AutoScalingGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceMapping) DeepCopyInto(out *BlockDeviceMapping) {
	*out = *in
	out.Ebs = in.Ebs
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceMapping.
func (in *BlockDeviceMapping) DeepCopy() *BlockDeviceMapping {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EBS) DeepCopyInto(out *EBS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EBS.
func (in *EBS) DeepCopy() *EBS {
	if in == nil {
		return nil
	}
	out := new(EBS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FargateProfileSpec) DeepCopyInto(out *FargateProfileSpec) {
	*out = *in
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalTags != nil {
		in, out := &in.AdditionalTags, &out.AdditionalTags
		*out = make(apiv1beta2.Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = make([]FargateSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FargateProfileSpec.
func (in *FargateProfileSpec) DeepCopy() *FargateProfileSpec {
	if in == nil {
		return nil
	}
	out := new(FargateProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FargateProfileStatus) DeepCopyInto(out *FargateProfileStatus) {
	*out = *in
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
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FargateProfileStatus.
func (in *FargateProfileStatus) DeepCopy() *FargateProfileStatus {
	if in == nil {
		return nil
	}
	out := new(FargateProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FargateSelector) DeepCopyInto(out *FargateSelector) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FargateSelector.
func (in *FargateSelector) DeepCopy() *FargateSelector {
	if in == nil {
		return nil
	}
	out := new(FargateSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstancesDistribution) DeepCopyInto(out *InstancesDistribution) {
	*out = *in
	if in.OnDemandBaseCapacity != nil {
		in, out := &in.OnDemandBaseCapacity, &out.OnDemandBaseCapacity
		*out = new(int64)
		**out = **in
	}
	if in.OnDemandPercentageAboveBaseCapacity != nil {
		in, out := &in.OnDemandPercentageAboveBaseCapacity, &out.OnDemandPercentageAboveBaseCapacity
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstancesDistribution.
func (in *InstancesDistribution) DeepCopy() *InstancesDistribution {
	if in == nil {
		return nil
	}
	out := new(InstancesDistribution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedMachinePoolScaling) DeepCopyInto(out *ManagedMachinePoolScaling) {
	*out = *in
	if in.MinSize != nil {
		in, out := &in.MinSize, &out.MinSize
		*out = new(int32)
		**out = **in
	}
	if in.MaxSize != nil {
		in, out := &in.MaxSize, &out.MaxSize
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedMachinePoolScaling.
func (in *ManagedMachinePoolScaling) DeepCopy() *ManagedMachinePoolScaling {
	if in == nil {
		return nil
	}
	out := new(ManagedMachinePoolScaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedRemoteAccess) DeepCopyInto(out *ManagedRemoteAccess) {
	*out = *in
	if in.SSHKeyName != nil {
		in, out := &in.SSHKeyName, &out.SSHKeyName
		*out = new(string)
		**out = **in
	}
	if in.SourceSecurityGroups != nil {
		in, out := &in.SourceSecurityGroups, &out.SourceSecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedRemoteAccess.
func (in *ManagedRemoteAccess) DeepCopy() *ManagedRemoteAccess {
	if in == nil {
		return nil
	}
	out := new(ManagedRemoteAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MixedInstancesPolicy) DeepCopyInto(out *MixedInstancesPolicy) {
	*out = *in
	if in.InstancesDistribution != nil {
		in, out := &in.InstancesDistribution, &out.InstancesDistribution
		*out = new(InstancesDistribution)
		(*in).DeepCopyInto(*out)
	}
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]Overrides, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MixedInstancesPolicy.
func (in *MixedInstancesPolicy) DeepCopy() *MixedInstancesPolicy {
	if in == nil {
		return nil
	}
	out := new(MixedInstancesPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Overrides) DeepCopyInto(out *Overrides) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Overrides.
func (in *Overrides) DeepCopy() *Overrides {
	if in == nil {
		return nil
	}
	out := new(Overrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Processes) DeepCopyInto(out *Processes) {
	*out = *in
	if in.Launch != nil {
		in, out := &in.Launch, &out.Launch
		*out = new(bool)
		**out = **in
	}
	if in.Terminate != nil {
		in, out := &in.Terminate, &out.Terminate
		*out = new(bool)
		**out = **in
	}
	if in.AddToLoadBalancer != nil {
		in, out := &in.AddToLoadBalancer, &out.AddToLoadBalancer
		*out = new(bool)
		**out = **in
	}
	if in.AlarmNotification != nil {
		in, out := &in.AlarmNotification, &out.AlarmNotification
		*out = new(bool)
		**out = **in
	}
	if in.AZRebalance != nil {
		in, out := &in.AZRebalance, &out.AZRebalance
		*out = new(bool)
		**out = **in
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(bool)
		**out = **in
	}
	if in.InstanceRefresh != nil {
		in, out := &in.InstanceRefresh, &out.InstanceRefresh
		*out = new(bool)
		**out = **in
	}
	if in.ReplaceUnhealthy != nil {
		in, out := &in.ReplaceUnhealthy, &out.ReplaceUnhealthy
		*out = new(bool)
		**out = **in
	}
	if in.ScheduledActions != nil {
		in, out := &in.ScheduledActions, &out.ScheduledActions
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Processes.
func (in *Processes) DeepCopy() *Processes {
	if in == nil {
		return nil
	}
	out := new(Processes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROSACluster) DeepCopyInto(out *ROSACluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROSACluster.
func (in *ROSACluster) DeepCopy() *ROSACluster {
	if in == nil {
		return nil
	}
	out := new(ROSACluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ROSACluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROSAClusterList) DeepCopyInto(out *ROSAClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ROSACluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROSAClusterList.
func (in *ROSAClusterList) DeepCopy() *ROSAClusterList {
	if in == nil {
		return nil
	}
	out := new(ROSAClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ROSAClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROSAClusterSpec) DeepCopyInto(out *ROSAClusterSpec) {
	*out = *in
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROSAClusterSpec.
func (in *ROSAClusterSpec) DeepCopy() *ROSAClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ROSAClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROSAClusterStatus) DeepCopyInto(out *ROSAClusterStatus) {
	*out = *in
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(v1beta1.FailureDomains, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROSAClusterStatus.
func (in *ROSAClusterStatus) DeepCopy() *ROSAClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ROSAClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROSAMachinePool) DeepCopyInto(out *ROSAMachinePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROSAMachinePool.
func (in *ROSAMachinePool) DeepCopy() *ROSAMachinePool {
	if in == nil {
		return nil
	}
	out := new(ROSAMachinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ROSAMachinePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROSAMachinePoolList) DeepCopyInto(out *ROSAMachinePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ROSAMachinePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROSAMachinePoolList.
func (in *ROSAMachinePoolList) DeepCopy() *ROSAMachinePoolList {
	if in == nil {
		return nil
	}
	out := new(ROSAMachinePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ROSAMachinePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RefreshPreferences) DeepCopyInto(out *RefreshPreferences) {
	*out = *in
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(string)
		**out = **in
	}
	if in.InstanceWarmup != nil {
		in, out := &in.InstanceWarmup, &out.InstanceWarmup
		*out = new(int64)
		**out = **in
	}
	if in.MinHealthyPercentage != nil {
		in, out := &in.MinHealthyPercentage, &out.MinHealthyPercentage
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RefreshPreferences.
func (in *RefreshPreferences) DeepCopy() *RefreshPreferences {
	if in == nil {
		return nil
	}
	out := new(RefreshPreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RosaMachinePoolAutoScaling) DeepCopyInto(out *RosaMachinePoolAutoScaling) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RosaMachinePoolAutoScaling.
func (in *RosaMachinePoolAutoScaling) DeepCopy() *RosaMachinePoolAutoScaling {
	if in == nil {
		return nil
	}
	out := new(RosaMachinePoolAutoScaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RosaMachinePoolSpec) DeepCopyInto(out *RosaMachinePoolSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]RosaTaint, len(*in))
		copy(*out, *in)
	}
	if in.Autoscaling != nil {
		in, out := &in.Autoscaling, &out.Autoscaling
		*out = new(RosaMachinePoolAutoScaling)
		**out = **in
	}
	if in.TuningConfigs != nil {
		in, out := &in.TuningConfigs, &out.TuningConfigs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalSecurityGroups != nil {
		in, out := &in.AdditionalSecurityGroups, &out.AdditionalSecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ProviderIDList != nil {
		in, out := &in.ProviderIDList, &out.ProviderIDList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RosaMachinePoolSpec.
func (in *RosaMachinePoolSpec) DeepCopy() *RosaMachinePoolSpec {
	if in == nil {
		return nil
	}
	out := new(RosaMachinePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RosaMachinePoolStatus) DeepCopyInto(out *RosaMachinePoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RosaMachinePoolStatus.
func (in *RosaMachinePoolStatus) DeepCopy() *RosaMachinePoolStatus {
	if in == nil {
		return nil
	}
	out := new(RosaMachinePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RosaTaint) DeepCopyInto(out *RosaTaint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RosaTaint.
func (in *RosaTaint) DeepCopy() *RosaTaint {
	if in == nil {
		return nil
	}
	out := new(RosaTaint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SuspendProcessesTypes) DeepCopyInto(out *SuspendProcessesTypes) {
	*out = *in
	if in.Processes != nil {
		in, out := &in.Processes, &out.Processes
		*out = new(Processes)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SuspendProcessesTypes.
func (in *SuspendProcessesTypes) DeepCopy() *SuspendProcessesTypes {
	if in == nil {
		return nil
	}
	out := new(SuspendProcessesTypes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Tags) DeepCopyInto(out *Tags) {
	{
		in := &in
		*out = make(Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tags.
func (in Tags) DeepCopy() Tags {
	if in == nil {
		return nil
	}
	out := new(Tags)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Taint) DeepCopyInto(out *Taint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Taint.
func (in *Taint) DeepCopy() *Taint {
	if in == nil {
		return nil
	}
	out := new(Taint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Taints) DeepCopyInto(out *Taints) {
	{
		in := &in
		*out = make(Taints, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Taints.
func (in Taints) DeepCopy() Taints {
	if in == nil {
		return nil
	}
	out := new(Taints)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateConfig) DeepCopyInto(out *UpdateConfig) {
	*out = *in
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(int)
		**out = **in
	}
	if in.MaxUnavailablePercentage != nil {
		in, out := &in.MaxUnavailablePercentage, &out.MaxUnavailablePercentage
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateConfig.
func (in *UpdateConfig) DeepCopy() *UpdateConfig {
	if in == nil {
		return nil
	}
	out := new(UpdateConfig)
	in.DeepCopyInto(out)
	return out
}
