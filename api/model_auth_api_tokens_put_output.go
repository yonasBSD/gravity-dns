/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthAPITokensPutOutput struct for AuthAPITokensPutOutput
type AuthAPITokensPutOutput struct {
	Key string `json:"key"`
}

// NewAuthAPITokensPutOutput instantiates a new AuthAPITokensPutOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAPITokensPutOutput(key string) *AuthAPITokensPutOutput {
	this := AuthAPITokensPutOutput{}
	this.Key = key
	return &this
}

// NewAuthAPITokensPutOutputWithDefaults instantiates a new AuthAPITokensPutOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAPITokensPutOutputWithDefaults() *AuthAPITokensPutOutput {
	this := AuthAPITokensPutOutput{}
	return &this
}

// GetKey returns the Key field value
func (o *AuthAPITokensPutOutput) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AuthAPITokensPutOutput) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AuthAPITokensPutOutput) SetKey(v string) {
	o.Key = v
}

func (o AuthAPITokensPutOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableAuthAPITokensPutOutput struct {
	value *AuthAPITokensPutOutput
	isSet bool
}

func (v NullableAuthAPITokensPutOutput) Get() *AuthAPITokensPutOutput {
	return v.value
}

func (v *NullableAuthAPITokensPutOutput) Set(val *AuthAPITokensPutOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAPITokensPutOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAPITokensPutOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAPITokensPutOutput(val *AuthAPITokensPutOutput) *NullableAuthAPITokensPutOutput {
	return &NullableAuthAPITokensPutOutput{value: val, isSet: true}
}

func (v NullableAuthAPITokensPutOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAPITokensPutOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
