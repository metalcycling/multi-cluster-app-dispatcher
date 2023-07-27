//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 The Kubernetes Authors.

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
/*
Copyright 2019, 2021, 2022, 2023 The Multi-Cluster App Dispatcher Authors.

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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapper) DeepCopyInto(out *AppWrapper) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapper.
func (in *AppWrapper) DeepCopy() *AppWrapper {
	if in == nil {
		return nil
	}
	out := new(AppWrapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppWrapper) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapperCondition) DeepCopyInto(out *AppWrapperCondition) {
	*out = *in
	in.LastUpdateMicroTime.DeepCopyInto(&out.LastUpdateMicroTime)
	in.LastTransitionMicroTime.DeepCopyInto(&out.LastTransitionMicroTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapperCondition.
func (in *AppWrapperCondition) DeepCopy() *AppWrapperCondition {
	if in == nil {
		return nil
	}
	out := new(AppWrapperCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapperGenericResource) DeepCopyInto(out *AppWrapperGenericResource) {
	*out = *in
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		*out = new(int32)
		**out = **in
	}
	in.GenericTemplate.DeepCopyInto(&out.GenericTemplate)
	if in.CustomPodResources != nil {
		in, out := &in.CustomPodResources, &out.CustomPodResources
		*out = make([]CustomPodResourceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapperGenericResource.
func (in *AppWrapperGenericResource) DeepCopy() *AppWrapperGenericResource {
	if in == nil {
		return nil
	}
	out := new(AppWrapperGenericResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapperList) DeepCopyInto(out *AppWrapperList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppWrapper, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapperList.
func (in *AppWrapperList) DeepCopy() *AppWrapperList {
	if in == nil {
		return nil
	}
	out := new(AppWrapperList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppWrapperList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapperResourceList) DeepCopyInto(out *AppWrapperResourceList) {
	*out = *in
	if in.GenericItems != nil {
		in, out := &in.GenericItems, &out.GenericItems
		*out = make([]AppWrapperGenericResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapperResourceList.
func (in *AppWrapperResourceList) DeepCopy() *AppWrapperResourceList {
	if in == nil {
		return nil
	}
	out := new(AppWrapperResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapperService) DeepCopyInto(out *AppWrapperService) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapperService.
func (in *AppWrapperService) DeepCopy() *AppWrapperService {
	if in == nil {
		return nil
	}
	out := new(AppWrapperService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapperSpec) DeepCopyInto(out *AppWrapperSpec) {
	*out = *in
	in.Service.DeepCopyInto(&out.Service)
	in.AggrResources.DeepCopyInto(&out.AggrResources)
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.SchedSpec.DeepCopyInto(&out.SchedSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapperSpec.
func (in *AppWrapperSpec) DeepCopy() *AppWrapperSpec {
	if in == nil {
		return nil
	}
	out := new(AppWrapperSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapperStatus) DeepCopyInto(out *AppWrapperStatus) {
	*out = *in
	in.ControllerFirstTimestamp.DeepCopyInto(&out.ControllerFirstTimestamp)
	in.ControllerFirstDispatchTimestamp.DeepCopyInto(&out.ControllerFirstDispatchTimestamp)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AppWrapperCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PendingPodConditions != nil {
		in, out := &in.PendingPodConditions, &out.PendingPodConditions
		*out = make([]PendingPodSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapperStatus.
func (in *AppWrapperStatus) DeepCopy() *AppWrapperStatus {
	if in == nil {
		return nil
	}
	out := new(AppWrapperStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterReference) DeepCopyInto(out *ClusterReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterReference.
func (in *ClusterReference) DeepCopy() *ClusterReference {
	if in == nil {
		return nil
	}
	out := new(ClusterReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSchedulingSpec) DeepCopyInto(out *ClusterSchedulingSpec) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]ClusterReference, len(*in))
		copy(*out, *in)
	}
	if in.ClusterSelector != nil {
		in, out := &in.ClusterSelector, &out.ClusterSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSchedulingSpec.
func (in *ClusterSchedulingSpec) DeepCopy() *ClusterSchedulingSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSchedulingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomPodResourceTemplate) DeepCopyInto(out *CustomPodResourceTemplate) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomPodResourceTemplate.
func (in *CustomPodResourceTemplate) DeepCopy() *CustomPodResourceTemplate {
	if in == nil {
		return nil
	}
	out := new(CustomPodResourceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DispatchDurationSpec) DeepCopyInto(out *DispatchDurationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DispatchDurationSpec.
func (in *DispatchDurationSpec) DeepCopy() *DispatchDurationSpec {
	if in == nil {
		return nil
	}
	out := new(DispatchDurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DispatchingWindowSpec) DeepCopyInto(out *DispatchingWindowSpec) {
	*out = *in
	in.Start.DeepCopyInto(&out.Start)
	in.End.DeepCopyInto(&out.End)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DispatchingWindowSpec.
func (in *DispatchingWindowSpec) DeepCopy() *DispatchingWindowSpec {
	if in == nil {
		return nil
	}
	out := new(DispatchingWindowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PendingPodSpec) DeepCopyInto(out *PendingPodSpec) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]corev1.PodCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PendingPodSpec.
func (in *PendingPodSpec) DeepCopy() *PendingPodSpec {
	if in == nil {
		return nil
	}
	out := new(PendingPodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueJob) DeepCopyInto(out *QueueJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueJob.
func (in *QueueJob) DeepCopy() *QueueJob {
	if in == nil {
		return nil
	}
	out := new(QueueJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueueJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueJobList) DeepCopyInto(out *QueueJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QueueJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueJobList.
func (in *QueueJobList) DeepCopy() *QueueJobList {
	if in == nil {
		return nil
	}
	out := new(QueueJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueueJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueJobSpec) DeepCopyInto(out *QueueJobSpec) {
	*out = *in
	in.SchedSpec.DeepCopyInto(&out.SchedSpec)
	if in.TaskSpecs != nil {
		in, out := &in.TaskSpecs, &out.TaskSpecs
		*out = make([]TaskSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueJobSpec.
func (in *QueueJobSpec) DeepCopy() *QueueJobSpec {
	if in == nil {
		return nil
	}
	out := new(QueueJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueJobStatus) DeepCopyInto(out *QueueJobStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueJobStatus.
func (in *QueueJobStatus) DeepCopy() *QueueJobStatus {
	if in == nil {
		return nil
	}
	out := new(QueueJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequeuingTemplate) DeepCopyInto(out *RequeuingTemplate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequeuingTemplate.
func (in *RequeuingTemplate) DeepCopy() *RequeuingTemplate {
	if in == nil {
		return nil
	}
	out := new(RequeuingTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleTimeSpec) DeepCopyInto(out *ScheduleTimeSpec) {
	*out = *in
	in.Min.DeepCopyInto(&out.Min)
	in.Desired.DeepCopyInto(&out.Desired)
	in.Max.DeepCopyInto(&out.Max)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleTimeSpec.
func (in *ScheduleTimeSpec) DeepCopy() *ScheduleTimeSpec {
	if in == nil {
		return nil
	}
	out := new(ScheduleTimeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingSpec) DeepCopyInto(out *SchedulingSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingSpec.
func (in *SchedulingSpec) DeepCopy() *SchedulingSpec {
	if in == nil {
		return nil
	}
	out := new(SchedulingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SchedulingSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingSpecList) DeepCopyInto(out *SchedulingSpecList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SchedulingSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingSpecList.
func (in *SchedulingSpecList) DeepCopy() *SchedulingSpecList {
	if in == nil {
		return nil
	}
	out := new(SchedulingSpecList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SchedulingSpecList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingSpecTemplate) DeepCopyInto(out *SchedulingSpecTemplate) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Requeuing = in.Requeuing
	out.DispatchDuration = in.DispatchDuration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingSpecTemplate.
func (in *SchedulingSpecTemplate) DeepCopy() *SchedulingSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(SchedulingSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskSpec) DeepCopyInto(out *TaskSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskSpec.
func (in *TaskSpec) DeepCopy() *TaskSpec {
	if in == nil {
		return nil
	}
	out := new(TaskSpec)
	in.DeepCopyInto(out)
	return out
}
