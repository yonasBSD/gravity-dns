/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.4.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ApiAPIToolPortmapInput struct for ApiAPIToolPortmapInput
type ApiAPIToolPortmapInput struct {
	Host *string `json:"host,omitempty"`
}

// NewApiAPIToolPortmapInput instantiates a new ApiAPIToolPortmapInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPIToolPortmapInput() *ApiAPIToolPortmapInput {
	this := ApiAPIToolPortmapInput{}
	return &this
}

// NewApiAPIToolPortmapInputWithDefaults instantiates a new ApiAPIToolPortmapInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPIToolPortmapInputWithDefaults() *ApiAPIToolPortmapInput {
	this := ApiAPIToolPortmapInput{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ApiAPIToolPortmapInput) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPortmapInput) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ApiAPIToolPortmapInput) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ApiAPIToolPortmapInput) SetHost(v string) {
	o.Host = &v
}

func (o ApiAPIToolPortmapInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	return json.Marshal(toSerialize)
}

type NullableApiAPIToolPortmapInput struct {
	value *ApiAPIToolPortmapInput
	isSet bool
}

func (v NullableApiAPIToolPortmapInput) Get() *ApiAPIToolPortmapInput {
	return v.value
}

func (v *NullableApiAPIToolPortmapInput) Set(val *ApiAPIToolPortmapInput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPIToolPortmapInput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPIToolPortmapInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPIToolPortmapInput(val *ApiAPIToolPortmapInput) *NullableApiAPIToolPortmapInput {
	return &NullableApiAPIToolPortmapInput{value: val, isSet: true}
}

func (v NullableApiAPIToolPortmapInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPIToolPortmapInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
