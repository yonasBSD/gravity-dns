/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.26.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TftpRoleConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TftpRoleConfig{}

// TftpRoleConfig struct for TftpRoleConfig
type TftpRoleConfig struct {
	EnableLocal *bool  `json:"enableLocal,omitempty"`
	Port        *int32 `json:"port,omitempty"`
}

// NewTftpRoleConfig instantiates a new TftpRoleConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTftpRoleConfig() *TftpRoleConfig {
	this := TftpRoleConfig{}
	return &this
}

// NewTftpRoleConfigWithDefaults instantiates a new TftpRoleConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTftpRoleConfigWithDefaults() *TftpRoleConfig {
	this := TftpRoleConfig{}
	return &this
}

// GetEnableLocal returns the EnableLocal field value if set, zero value otherwise.
func (o *TftpRoleConfig) GetEnableLocal() bool {
	if o == nil || IsNil(o.EnableLocal) {
		var ret bool
		return ret
	}
	return *o.EnableLocal
}

// GetEnableLocalOk returns a tuple with the EnableLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TftpRoleConfig) GetEnableLocalOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableLocal) {
		return nil, false
	}
	return o.EnableLocal, true
}

// HasEnableLocal returns a boolean if a field has been set.
func (o *TftpRoleConfig) HasEnableLocal() bool {
	if o != nil && !IsNil(o.EnableLocal) {
		return true
	}

	return false
}

// SetEnableLocal gets a reference to the given bool and assigns it to the EnableLocal field.
func (o *TftpRoleConfig) SetEnableLocal(v bool) {
	o.EnableLocal = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *TftpRoleConfig) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TftpRoleConfig) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *TftpRoleConfig) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *TftpRoleConfig) SetPort(v int32) {
	o.Port = &v
}

func (o TftpRoleConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TftpRoleConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableLocal) {
		toSerialize["enableLocal"] = o.EnableLocal
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableTftpRoleConfig struct {
	value *TftpRoleConfig
	isSet bool
}

func (v NullableTftpRoleConfig) Get() *TftpRoleConfig {
	return v.value
}

func (v *NullableTftpRoleConfig) Set(val *TftpRoleConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTftpRoleConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTftpRoleConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTftpRoleConfig(val *TftpRoleConfig) *NullableTftpRoleConfig {
	return &NullableTftpRoleConfig{value: val, isSet: true}
}

func (v NullableTftpRoleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTftpRoleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
