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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta2

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
)

// DeviceCapacityApplyConfiguration represents a declarative configuration of the DeviceCapacity type for use
// with apply.
type DeviceCapacityApplyConfiguration struct {
	Value         *resource.Quantity                       `json:"value,omitempty"`
	RequestPolicy *CapacityRequestPolicyApplyConfiguration `json:"requestPolicy,omitempty"`
}

// DeviceCapacityApplyConfiguration constructs a declarative configuration of the DeviceCapacity type for use with
// apply.
func DeviceCapacity() *DeviceCapacityApplyConfiguration {
	return &DeviceCapacityApplyConfiguration{}
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *DeviceCapacityApplyConfiguration) WithValue(value resource.Quantity) *DeviceCapacityApplyConfiguration {
	b.Value = &value
	return b
}

// WithRequestPolicy sets the RequestPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RequestPolicy field is set to the value of the last call.
func (b *DeviceCapacityApplyConfiguration) WithRequestPolicy(value *CapacityRequestPolicyApplyConfiguration) *DeviceCapacityApplyConfiguration {
	b.RequestPolicy = value
	return b
}
