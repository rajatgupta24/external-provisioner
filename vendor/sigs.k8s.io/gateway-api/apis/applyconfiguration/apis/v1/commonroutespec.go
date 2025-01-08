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

// CommonRouteSpecApplyConfiguration represents a declarative configuration of the CommonRouteSpec type for use
// with apply.
type CommonRouteSpecApplyConfiguration struct {
	ParentRefs []ParentReferenceApplyConfiguration `json:"parentRefs,omitempty"`
}

// CommonRouteSpecApplyConfiguration constructs a declarative configuration of the CommonRouteSpec type for use with
// apply.
func CommonRouteSpec() *CommonRouteSpecApplyConfiguration {
	return &CommonRouteSpecApplyConfiguration{}
}

// WithParentRefs adds the given value to the ParentRefs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ParentRefs field.
func (b *CommonRouteSpecApplyConfiguration) WithParentRefs(values ...*ParentReferenceApplyConfiguration) *CommonRouteSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithParentRefs")
		}
		b.ParentRefs = append(b.ParentRefs, *values[i])
	}
	return b
}
