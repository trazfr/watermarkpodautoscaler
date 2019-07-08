// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatermarkPodAutoscaler) DeepCopyInto(out *WatermarkPodAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatermarkPodAutoscaler.
func (in *WatermarkPodAutoscaler) DeepCopy() *WatermarkPodAutoscaler {
	if in == nil {
		return nil
	}
	out := new(WatermarkPodAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WatermarkPodAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatermarkPodAutoscalerList) DeepCopyInto(out *WatermarkPodAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WatermarkPodAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatermarkPodAutoscalerList.
func (in *WatermarkPodAutoscalerList) DeepCopy() *WatermarkPodAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(WatermarkPodAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WatermarkPodAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatermarkPodAutoscalerSpec) DeepCopyInto(out *WatermarkPodAutoscalerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatermarkPodAutoscalerSpec.
func (in *WatermarkPodAutoscalerSpec) DeepCopy() *WatermarkPodAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(WatermarkPodAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatermarkPodAutoscalerStatus) DeepCopyInto(out *WatermarkPodAutoscalerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatermarkPodAutoscalerStatus.
func (in *WatermarkPodAutoscalerStatus) DeepCopy() *WatermarkPodAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(WatermarkPodAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}