/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.6.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ApiAPIToolPingInput struct for ApiAPIToolPingInput
type ApiAPIToolPingInput struct {
	Host *string `json:"host,omitempty"`
}

// NewApiAPIToolPingInput instantiates a new ApiAPIToolPingInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPIToolPingInput() *ApiAPIToolPingInput {
	this := ApiAPIToolPingInput{}
	return &this
}

// NewApiAPIToolPingInputWithDefaults instantiates a new ApiAPIToolPingInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPIToolPingInputWithDefaults() *ApiAPIToolPingInput {
	this := ApiAPIToolPingInput{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ApiAPIToolPingInput) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingInput) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ApiAPIToolPingInput) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ApiAPIToolPingInput) SetHost(v string) {
	o.Host = &v
}

func (o ApiAPIToolPingInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	return json.Marshal(toSerialize)
}

type NullableApiAPIToolPingInput struct {
	value *ApiAPIToolPingInput
	isSet bool
}

func (v NullableApiAPIToolPingInput) Get() *ApiAPIToolPingInput {
	return v.value
}

func (v *NullableApiAPIToolPingInput) Set(val *ApiAPIToolPingInput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPIToolPingInput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPIToolPingInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPIToolPingInput(val *ApiAPIToolPingInput) *NullableApiAPIToolPingInput {
	return &NullableApiAPIToolPingInput{value: val, isSet: true}
}

func (v NullableApiAPIToolPingInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPIToolPingInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
