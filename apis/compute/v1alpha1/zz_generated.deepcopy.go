//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Droplet) DeepCopyInto(out *Droplet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Droplet.
func (in *Droplet) DeepCopy() *Droplet {
	if in == nil {
		return nil
	}
	out := new(Droplet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Droplet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DropletInitParameters) DeepCopyInto(out *DropletInitParameters) {
	*out = *in
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = new(bool)
		**out = **in
	}
	if in.DropletAgent != nil {
		in, out := &in.DropletAgent, &out.DropletAgent
		*out = new(bool)
		**out = **in
	}
	if in.GracefulShutdown != nil {
		in, out := &in.GracefulShutdown, &out.GracefulShutdown
		*out = new(bool)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(bool)
		**out = **in
	}
	if in.IPv6Address != nil {
		in, out := &in.IPv6Address, &out.IPv6Address
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrivateNetworking != nil {
		in, out := &in.PrivateNetworking, &out.PrivateNetworking
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ResizeDisk != nil {
		in, out := &in.ResizeDisk, &out.ResizeDisk
		*out = new(bool)
		**out = **in
	}
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
	if in.VPCUUID != nil {
		in, out := &in.VPCUUID, &out.VPCUUID
		*out = new(string)
		**out = **in
	}
	if in.VolumeIds != nil {
		in, out := &in.VolumeIds, &out.VolumeIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DropletInitParameters.
func (in *DropletInitParameters) DeepCopy() *DropletInitParameters {
	if in == nil {
		return nil
	}
	out := new(DropletInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DropletList) DeepCopyInto(out *DropletList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Droplet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DropletList.
func (in *DropletList) DeepCopy() *DropletList {
	if in == nil {
		return nil
	}
	out := new(DropletList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DropletList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DropletObservation) DeepCopyInto(out *DropletObservation) {
	*out = *in
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = new(bool)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = new(float64)
		**out = **in
	}
	if in.DropletAgent != nil {
		in, out := &in.DropletAgent, &out.DropletAgent
		*out = new(bool)
		**out = **in
	}
	if in.GracefulShutdown != nil {
		in, out := &in.GracefulShutdown, &out.GracefulShutdown
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IPv4Address != nil {
		in, out := &in.IPv4Address, &out.IPv4Address
		*out = new(string)
		**out = **in
	}
	if in.IPv4AddressPrivate != nil {
		in, out := &in.IPv4AddressPrivate, &out.IPv4AddressPrivate
		*out = new(string)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(bool)
		**out = **in
	}
	if in.IPv6Address != nil {
		in, out := &in.IPv6Address, &out.IPv6Address
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Locked != nil {
		in, out := &in.Locked, &out.Locked
		*out = new(bool)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PriceHourly != nil {
		in, out := &in.PriceHourly, &out.PriceHourly
		*out = new(float64)
		**out = **in
	}
	if in.PriceMonthly != nil {
		in, out := &in.PriceMonthly, &out.PriceMonthly
		*out = new(float64)
		**out = **in
	}
	if in.PrivateNetworking != nil {
		in, out := &in.PrivateNetworking, &out.PrivateNetworking
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ResizeDisk != nil {
		in, out := &in.ResizeDisk, &out.ResizeDisk
		*out = new(bool)
		**out = **in
	}
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Urn != nil {
		in, out := &in.Urn, &out.Urn
		*out = new(string)
		**out = **in
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
	if in.VPCUUID != nil {
		in, out := &in.VPCUUID, &out.VPCUUID
		*out = new(string)
		**out = **in
	}
	if in.Vcpus != nil {
		in, out := &in.Vcpus, &out.Vcpus
		*out = new(float64)
		**out = **in
	}
	if in.VolumeIds != nil {
		in, out := &in.VolumeIds, &out.VolumeIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DropletObservation.
func (in *DropletObservation) DeepCopy() *DropletObservation {
	if in == nil {
		return nil
	}
	out := new(DropletObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DropletParameters) DeepCopyInto(out *DropletParameters) {
	*out = *in
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = new(bool)
		**out = **in
	}
	if in.DropletAgent != nil {
		in, out := &in.DropletAgent, &out.DropletAgent
		*out = new(bool)
		**out = **in
	}
	if in.GracefulShutdown != nil {
		in, out := &in.GracefulShutdown, &out.GracefulShutdown
		*out = new(bool)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(bool)
		**out = **in
	}
	if in.IPv6Address != nil {
		in, out := &in.IPv6Address, &out.IPv6Address
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrivateNetworking != nil {
		in, out := &in.PrivateNetworking, &out.PrivateNetworking
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ResizeDisk != nil {
		in, out := &in.ResizeDisk, &out.ResizeDisk
		*out = new(bool)
		**out = **in
	}
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
	if in.VPCUUID != nil {
		in, out := &in.VPCUUID, &out.VPCUUID
		*out = new(string)
		**out = **in
	}
	if in.VolumeIds != nil {
		in, out := &in.VolumeIds, &out.VolumeIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DropletParameters.
func (in *DropletParameters) DeepCopy() *DropletParameters {
	if in == nil {
		return nil
	}
	out := new(DropletParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DropletSpec) DeepCopyInto(out *DropletSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DropletSpec.
func (in *DropletSpec) DeepCopy() *DropletSpec {
	if in == nil {
		return nil
	}
	out := new(DropletSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DropletStatus) DeepCopyInto(out *DropletStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DropletStatus.
func (in *DropletStatus) DeepCopy() *DropletStatus {
	if in == nil {
		return nil
	}
	out := new(DropletStatus)
	in.DeepCopyInto(out)
	return out
}