/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.26.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ApiAPIMembersOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAPIMembersOutput{}

// ApiAPIMembersOutput struct for ApiAPIMembersOutput
type ApiAPIMembersOutput struct {
	Members []ApiAPIMember `json:"members,omitempty"`
}

// NewApiAPIMembersOutput instantiates a new ApiAPIMembersOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPIMembersOutput() *ApiAPIMembersOutput {
	this := ApiAPIMembersOutput{}
	return &this
}

// NewApiAPIMembersOutputWithDefaults instantiates a new ApiAPIMembersOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPIMembersOutputWithDefaults() *ApiAPIMembersOutput {
	this := ApiAPIMembersOutput{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiAPIMembersOutput) GetMembers() []ApiAPIMember {
	if o == nil {
		var ret []ApiAPIMember
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiAPIMembersOutput) GetMembersOk() ([]ApiAPIMember, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ApiAPIMembersOutput) HasMembers() bool {
	if o != nil && IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []ApiAPIMember and assigns it to the Members field.
func (o *ApiAPIMembersOutput) SetMembers(v []ApiAPIMember) {
	o.Members = v
}

func (o ApiAPIMembersOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAPIMembersOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	return toSerialize, nil
}

type NullableApiAPIMembersOutput struct {
	value *ApiAPIMembersOutput
	isSet bool
}

func (v NullableApiAPIMembersOutput) Get() *ApiAPIMembersOutput {
	return v.value
}

func (v *NullableApiAPIMembersOutput) Set(val *ApiAPIMembersOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPIMembersOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPIMembersOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPIMembersOutput(val *ApiAPIMembersOutput) *NullableApiAPIMembersOutput {
	return &NullableApiAPIMembersOutput{value: val, isSet: true}
}

func (v NullableApiAPIMembersOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPIMembersOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
