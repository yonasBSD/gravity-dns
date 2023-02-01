/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.4.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ApiAPIToolTracerouteOutput struct for ApiAPIToolTracerouteOutput
type ApiAPIToolTracerouteOutput struct {
	Hops []ApiAPIToolTracerouteOutputHop `json:"hops,omitempty"`
}

// NewApiAPIToolTracerouteOutput instantiates a new ApiAPIToolTracerouteOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPIToolTracerouteOutput() *ApiAPIToolTracerouteOutput {
	this := ApiAPIToolTracerouteOutput{}
	return &this
}

// NewApiAPIToolTracerouteOutputWithDefaults instantiates a new ApiAPIToolTracerouteOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPIToolTracerouteOutputWithDefaults() *ApiAPIToolTracerouteOutput {
	this := ApiAPIToolTracerouteOutput{}
	return &this
}

// GetHops returns the Hops field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiAPIToolTracerouteOutput) GetHops() []ApiAPIToolTracerouteOutputHop {
	if o == nil {
		var ret []ApiAPIToolTracerouteOutputHop
		return ret
	}
	return o.Hops
}

// GetHopsOk returns a tuple with the Hops field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiAPIToolTracerouteOutput) GetHopsOk() ([]ApiAPIToolTracerouteOutputHop, bool) {
	if o == nil || o.Hops == nil {
		return nil, false
	}
	return o.Hops, true
}

// HasHops returns a boolean if a field has been set.
func (o *ApiAPIToolTracerouteOutput) HasHops() bool {
	if o != nil && o.Hops != nil {
		return true
	}

	return false
}

// SetHops gets a reference to the given []ApiAPIToolTracerouteOutputHop and assigns it to the Hops field.
func (o *ApiAPIToolTracerouteOutput) SetHops(v []ApiAPIToolTracerouteOutputHop) {
	o.Hops = v
}

func (o ApiAPIToolTracerouteOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hops != nil {
		toSerialize["hops"] = o.Hops
	}
	return json.Marshal(toSerialize)
}

type NullableApiAPIToolTracerouteOutput struct {
	value *ApiAPIToolTracerouteOutput
	isSet bool
}

func (v NullableApiAPIToolTracerouteOutput) Get() *ApiAPIToolTracerouteOutput {
	return v.value
}

func (v *NullableApiAPIToolTracerouteOutput) Set(val *ApiAPIToolTracerouteOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPIToolTracerouteOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPIToolTracerouteOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPIToolTracerouteOutput(val *ApiAPIToolTracerouteOutput) *NullableApiAPIToolTracerouteOutput {
	return &NullableApiAPIToolTracerouteOutput{value: val, isSet: true}
}

func (v NullableApiAPIToolTracerouteOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPIToolTracerouteOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
