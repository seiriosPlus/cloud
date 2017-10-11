// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package controller

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MasterSpec).DeepCopyInto(out.(*MasterSpec))
			return nil
		}, InType: reflect.TypeOf(&MasterSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PserverSpec).DeepCopyInto(out.(*PserverSpec))
			return nil
		}, InType: reflect.TypeOf(&PserverSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TrainerSpec).DeepCopyInto(out.(*TrainerSpec))
			return nil
		}, InType: reflect.TypeOf(&TrainerSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TrainingJob).DeepCopyInto(out.(*TrainingJob))
			return nil
		}, InType: reflect.TypeOf(&TrainingJob{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TrainingJobList).DeepCopyInto(out.(*TrainingJobList))
			return nil
		}, InType: reflect.TypeOf(&TrainingJobList{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TrainingJobSpec).DeepCopyInto(out.(*TrainingJobSpec))
			return nil
		}, InType: reflect.TypeOf(&TrainingJobSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TrainingJobStatus).DeepCopyInto(out.(*TrainingJobStatus))
			return nil
		}, InType: reflect.TypeOf(&TrainingJobStatus{})},
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterSpec) DeepCopyInto(out *MasterSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterSpec.
func (in *MasterSpec) DeepCopy() *MasterSpec {
	if in == nil {
		return nil
	}
	out := new(MasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PserverSpec) DeepCopyInto(out *PserverSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PserverSpec.
func (in *PserverSpec) DeepCopy() *PserverSpec {
	if in == nil {
		return nil
	}
	out := new(PserverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainerSpec) DeepCopyInto(out *TrainerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainerSpec.
func (in *TrainerSpec) DeepCopy() *TrainerSpec {
	if in == nil {
		return nil
	}
	out := new(TrainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainingJob) DeepCopyInto(out *TrainingJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainingJob.
func (in *TrainingJob) DeepCopy() *TrainingJob {
	if in == nil {
		return nil
	}
	out := new(TrainingJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrainingJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainingJobList) DeepCopyInto(out *TrainingJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrainingJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainingJobList.
func (in *TrainingJobList) DeepCopy() *TrainingJobList {
	if in == nil {
		return nil
	}
	out := new(TrainingJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainingJobSpec) DeepCopyInto(out *TrainingJobSpec) {
	*out = *in
	out.Trainer = in.Trainer
	out.Pserver = in.Pserver
	out.Master = in.Master
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainingJobSpec.
func (in *TrainingJobSpec) DeepCopy() *TrainingJobSpec {
	if in == nil {
		return nil
	}
	out := new(TrainingJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainingJobStatus) DeepCopyInto(out *TrainingJobStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainingJobStatus.
func (in *TrainingJobStatus) DeepCopy() *TrainingJobStatus {
	if in == nil {
		return nil
	}
	out := new(TrainingJobStatus)
	in.DeepCopyInto(out)
	return out
}