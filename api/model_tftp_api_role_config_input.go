/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.27.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TftpAPIRoleConfigInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TftpAPIRoleConfigInput{}

// TftpAPIRoleConfigInput struct for TftpAPIRoleConfigInput
type TftpAPIRoleConfigInput struct {
	Config TftpRoleConfig `json:"config"`
}

// NewTftpAPIRoleConfigInput instantiates a new TftpAPIRoleConfigInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTftpAPIRoleConfigInput(config TftpRoleConfig) *TftpAPIRoleConfigInput {
	this := TftpAPIRoleConfigInput{}
	this.Config = config
	return &this
}

// NewTftpAPIRoleConfigInputWithDefaults instantiates a new TftpAPIRoleConfigInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTftpAPIRoleConfigInputWithDefaults() *TftpAPIRoleConfigInput {
	this := TftpAPIRoleConfigInput{}
	return &this
}

// GetConfig returns the Config field value
func (o *TftpAPIRoleConfigInput) GetConfig() TftpRoleConfig {
	if o == nil {
		var ret TftpRoleConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *TftpAPIRoleConfigInput) GetConfigOk() (*TftpRoleConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *TftpAPIRoleConfigInput) SetConfig(v TftpRoleConfig) {
	o.Config = v
}

func (o TftpAPIRoleConfigInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TftpAPIRoleConfigInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	return toSerialize, nil
}

type NullableTftpAPIRoleConfigInput struct {
	value *TftpAPIRoleConfigInput
	isSet bool
}

func (v NullableTftpAPIRoleConfigInput) Get() *TftpAPIRoleConfigInput {
	return v.value
}

func (v *NullableTftpAPIRoleConfigInput) Set(val *TftpAPIRoleConfigInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTftpAPIRoleConfigInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTftpAPIRoleConfigInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTftpAPIRoleConfigInput(val *TftpAPIRoleConfigInput) *NullableTftpAPIRoleConfigInput {
	return &NullableTftpAPIRoleConfigInput{value: val, isSet: true}
}

func (v NullableTftpAPIRoleConfigInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTftpAPIRoleConfigInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
