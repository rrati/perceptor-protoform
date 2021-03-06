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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsSight) DeepCopyInto(out *OpsSight) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsSight.
func (in *OpsSight) DeepCopy() *OpsSight {
	if in == nil {
		return nil
	}
	out := new(OpsSight)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpsSight) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsSightList) DeepCopyInto(out *OpsSightList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpsSight, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsSightList.
func (in *OpsSightList) DeepCopy() *OpsSightList {
	if in == nil {
		return nil
	}
	out := new(OpsSightList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpsSightList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsSightSpec) DeepCopyInto(out *OpsSightSpec) {
	*out = *in
	if in.PerceptorPort != nil {
		in, out := &in.PerceptorPort, &out.PerceptorPort
		*out = new(int)
		**out = **in
	}
	if in.ScannerPort != nil {
		in, out := &in.ScannerPort, &out.ScannerPort
		*out = new(int)
		**out = **in
	}
	if in.PerceiverPort != nil {
		in, out := &in.PerceiverPort, &out.PerceiverPort
		*out = new(int)
		**out = **in
	}
	if in.ImageFacadePort != nil {
		in, out := &in.ImageFacadePort, &out.ImageFacadePort
		*out = new(int)
		**out = **in
	}
	if in.SkyfirePort != nil {
		in, out := &in.SkyfirePort, &out.SkyfirePort
		*out = new(int)
		**out = **in
	}
	if in.InternalRegistries != nil {
		in, out := &in.InternalRegistries, &out.InternalRegistries
		*out = make([]RegistryAuth, len(*in))
		copy(*out, *in)
	}
	if in.AnnotationIntervalSeconds != nil {
		in, out := &in.AnnotationIntervalSeconds, &out.AnnotationIntervalSeconds
		*out = new(int)
		**out = **in
	}
	if in.DumpIntervalMinutes != nil {
		in, out := &in.DumpIntervalMinutes, &out.DumpIntervalMinutes
		*out = new(int)
		**out = **in
	}
	if in.HubPort != nil {
		in, out := &in.HubPort, &out.HubPort
		*out = new(int)
		**out = **in
	}
	if in.HubClientTimeoutPerceptorMilliseconds != nil {
		in, out := &in.HubClientTimeoutPerceptorMilliseconds, &out.HubClientTimeoutPerceptorMilliseconds
		*out = new(int)
		**out = **in
	}
	if in.HubClientTimeoutScannerSeconds != nil {
		in, out := &in.HubClientTimeoutScannerSeconds, &out.HubClientTimeoutScannerSeconds
		*out = new(int)
		**out = **in
	}
	if in.ConcurrentScanLimit != nil {
		in, out := &in.ConcurrentScanLimit, &out.ConcurrentScanLimit
		*out = new(int)
		**out = **in
	}
	if in.TotalScanLimit != nil {
		in, out := &in.TotalScanLimit, &out.TotalScanLimit
		*out = new(int)
		**out = **in
	}
	if in.CheckForStalledScansPauseHours != nil {
		in, out := &in.CheckForStalledScansPauseHours, &out.CheckForStalledScansPauseHours
		*out = new(int)
		**out = **in
	}
	if in.StalledScanClientTimeoutHours != nil {
		in, out := &in.StalledScanClientTimeoutHours, &out.StalledScanClientTimeoutHours
		*out = new(int)
		**out = **in
	}
	if in.ModelMetricsPauseSeconds != nil {
		in, out := &in.ModelMetricsPauseSeconds, &out.ModelMetricsPauseSeconds
		*out = new(int)
		**out = **in
	}
	if in.UnknownImagePauseMilliseconds != nil {
		in, out := &in.UnknownImagePauseMilliseconds, &out.UnknownImagePauseMilliseconds
		*out = new(int)
		**out = **in
	}
	if in.ServiceAccounts != nil {
		in, out := &in.ServiceAccounts, &out.ServiceAccounts
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ContainerNames != nil {
		in, out := &in.ContainerNames, &out.ContainerNames
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ImagePerceiver != nil {
		in, out := &in.ImagePerceiver, &out.ImagePerceiver
		*out = new(bool)
		**out = **in
	}
	if in.PodPerceiver != nil {
		in, out := &in.PodPerceiver, &out.PodPerceiver
		*out = new(bool)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(bool)
		**out = **in
	}
	if in.PerceptorSkyfire != nil {
		in, out := &in.PerceptorSkyfire, &out.PerceptorSkyfire
		*out = new(bool)
		**out = **in
	}
	if in.UseMockMode != nil {
		in, out := &in.UseMockMode, &out.UseMockMode
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsSightSpec.
func (in *OpsSightSpec) DeepCopy() *OpsSightSpec {
	if in == nil {
		return nil
	}
	out := new(OpsSightSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsSightStatus) DeepCopyInto(out *OpsSightStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsSightStatus.
func (in *OpsSightStatus) DeepCopy() *OpsSightStatus {
	if in == nil {
		return nil
	}
	out := new(OpsSightStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryAuth) DeepCopyInto(out *RegistryAuth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryAuth.
func (in *RegistryAuth) DeepCopy() *RegistryAuth {
	if in == nil {
		return nil
	}
	out := new(RegistryAuth)
	in.DeepCopyInto(out)
	return out
}
