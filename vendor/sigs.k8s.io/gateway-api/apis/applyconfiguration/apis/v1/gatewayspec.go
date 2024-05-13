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

package v1

import (
	v1 "sigs.k8s.io/gateway-api/apis/v1"
)

// GatewaySpecApplyConfiguration represents an declarative configuration of the GatewaySpec type for use
// with apply.
type GatewaySpecApplyConfiguration struct {
	GatewayClassName *v1.ObjectName                           `json:"gatewayClassName,omitempty"`
	Listeners        []ListenerApplyConfiguration             `json:"listeners,omitempty"`
	Addresses        []GatewayAddressApplyConfiguration       `json:"addresses,omitempty"`
	Infrastructure   *GatewayInfrastructureApplyConfiguration `json:"infrastructure,omitempty"`
}

// GatewaySpecApplyConfiguration constructs an declarative configuration of the GatewaySpec type for use with
// apply.
func GatewaySpec() *GatewaySpecApplyConfiguration {
	return &GatewaySpecApplyConfiguration{}
}

// WithGatewayClassName sets the GatewayClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GatewayClassName field is set to the value of the last call.
func (b *GatewaySpecApplyConfiguration) WithGatewayClassName(value v1.ObjectName) *GatewaySpecApplyConfiguration {
	b.GatewayClassName = &value
	return b
}

// WithListeners adds the given value to the Listeners field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Listeners field.
func (b *GatewaySpecApplyConfiguration) WithListeners(values ...*ListenerApplyConfiguration) *GatewaySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithListeners")
		}
		b.Listeners = append(b.Listeners, *values[i])
	}
	return b
}

// WithAddresses adds the given value to the Addresses field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Addresses field.
func (b *GatewaySpecApplyConfiguration) WithAddresses(values ...*GatewayAddressApplyConfiguration) *GatewaySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAddresses")
		}
		b.Addresses = append(b.Addresses, *values[i])
	}
	return b
}

// WithInfrastructure sets the Infrastructure field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Infrastructure field is set to the value of the last call.
func (b *GatewaySpecApplyConfiguration) WithInfrastructure(value *GatewayInfrastructureApplyConfiguration) *GatewaySpecApplyConfiguration {
	b.Infrastructure = value
	return b
}
