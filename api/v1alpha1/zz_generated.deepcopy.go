//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The l7mp/stunner team.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayConfig) DeepCopyInto(out *GatewayConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayConfig.
func (in *GatewayConfig) DeepCopy() *GatewayConfig {
	if in == nil {
		return nil
	}
	out := new(GatewayConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayConfigList) DeepCopyInto(out *GatewayConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GatewayConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayConfigList.
func (in *GatewayConfigList) DeepCopy() *GatewayConfigList {
	if in == nil {
		return nil
	}
	out := new(GatewayConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayConfigSpec) DeepCopyInto(out *GatewayConfigSpec) {
	*out = *in
	if in.StunnerConfig != nil {
		in, out := &in.StunnerConfig, &out.StunnerConfig
		*out = new(string)
		**out = **in
	}
	if in.Realm != nil {
		in, out := &in.Realm, &out.Realm
		*out = new(string)
		**out = **in
	}
	if in.MetricsEndpoint != nil {
		in, out := &in.MetricsEndpoint, &out.MetricsEndpoint
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckEndpoint != nil {
		in, out := &in.HealthCheckEndpoint, &out.HealthCheckEndpoint
		*out = new(string)
		**out = **in
	}
	if in.AuthType != nil {
		in, out := &in.AuthType, &out.AuthType
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.SharedSecret != nil {
		in, out := &in.SharedSecret, &out.SharedSecret
		*out = new(string)
		**out = **in
	}
	if in.AuthLifetime != nil {
		in, out := &in.AuthLifetime, &out.AuthLifetime
		*out = new(int32)
		**out = **in
	}
	if in.LoadBalancerServiceAnnotations != nil {
		in, out := &in.LoadBalancerServiceAnnotations, &out.LoadBalancerServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.MinPort != nil {
		in, out := &in.MinPort, &out.MinPort
		*out = new(int32)
		**out = **in
	}
	if in.MaxPort != nil {
		in, out := &in.MaxPort, &out.MaxPort
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayConfigSpec.
func (in *GatewayConfigSpec) DeepCopy() *GatewayConfigSpec {
	if in == nil {
		return nil
	}
	out := new(GatewayConfigSpec)
	in.DeepCopyInto(out)
	return out
}
