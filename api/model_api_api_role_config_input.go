/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.6.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ApiAPIRoleConfigInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAPIRoleConfigInput{}

// ApiAPIRoleConfigInput struct for ApiAPIRoleConfigInput
type ApiAPIRoleConfigInput struct {
	Config ApiRoleConfig `json:"config"`
}

// NewApiAPIRoleConfigInput instantiates a new ApiAPIRoleConfigInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPIRoleConfigInput(config ApiRoleConfig) *ApiAPIRoleConfigInput {
	this := ApiAPIRoleConfigInput{}
	this.Config = config
	return &this
}

// NewApiAPIRoleConfigInputWithDefaults instantiates a new ApiAPIRoleConfigInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPIRoleConfigInputWithDefaults() *ApiAPIRoleConfigInput {
	this := ApiAPIRoleConfigInput{}
	return &this
}

// GetConfig returns the Config field value
func (o *ApiAPIRoleConfigInput) GetConfig() ApiRoleConfig {
	if o == nil {
		var ret ApiRoleConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ApiAPIRoleConfigInput) GetConfigOk() (*ApiRoleConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *ApiAPIRoleConfigInput) SetConfig(v ApiRoleConfig) {
	o.Config = v
}

func (o ApiAPIRoleConfigInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAPIRoleConfigInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	return toSerialize, nil
}

type NullableApiAPIRoleConfigInput struct {
	value *ApiAPIRoleConfigInput
	isSet bool
}

func (v NullableApiAPIRoleConfigInput) Get() *ApiAPIRoleConfigInput {
	return v.value
}

func (v *NullableApiAPIRoleConfigInput) Set(val *ApiAPIRoleConfigInput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPIRoleConfigInput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPIRoleConfigInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPIRoleConfigInput(val *ApiAPIRoleConfigInput) *NullableApiAPIRoleConfigInput {
	return &NullableApiAPIRoleConfigInput{value: val, isSet: true}
}

func (v NullableApiAPIRoleConfigInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPIRoleConfigInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
