//go:build !ignore_autogenerated

// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package volume

import (
	"k8s.io/api/core/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesVolume) DeepCopyInto(out *KubernetesVolume) {
	*out = *in
	if in.HostPathLegacy != nil {
		in, out := &in.HostPathLegacy, &out.HostPathLegacy
		*out = new(v1.HostPathVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.HostPath != nil {
		in, out := &in.HostPath, &out.HostPath
		*out = new(v1.HostPathVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.EmptyDir != nil {
		in, out := &in.EmptyDir, &out.EmptyDir
		*out = new(v1.EmptyDirVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretSource != nil {
		in, out := &in.SecretSource, &out.SecretSource
		*out = new(v1.SecretVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesVolume.
func (in *KubernetesVolume) DeepCopy() *KubernetesVolume {
	if in == nil {
		return nil
	}
	out := new(KubernetesVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaim) DeepCopyInto(out *PersistentVolumeClaim) {
	*out = *in
	in.PersistentVolumeClaimSpec.DeepCopyInto(&out.PersistentVolumeClaimSpec)
	out.PersistentVolumeSource = in.PersistentVolumeSource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaim.
func (in *PersistentVolumeClaim) DeepCopy() *PersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}
