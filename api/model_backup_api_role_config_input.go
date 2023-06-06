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

// BackupAPIRoleConfigInput struct for BackupAPIRoleConfigInput
type BackupAPIRoleConfigInput struct {
	Config BackupRoleConfig `json:"config"`
}

// NewBackupAPIRoleConfigInput instantiates a new BackupAPIRoleConfigInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupAPIRoleConfigInput(config BackupRoleConfig) *BackupAPIRoleConfigInput {
	this := BackupAPIRoleConfigInput{}
	this.Config = config
	return &this
}

// NewBackupAPIRoleConfigInputWithDefaults instantiates a new BackupAPIRoleConfigInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupAPIRoleConfigInputWithDefaults() *BackupAPIRoleConfigInput {
	this := BackupAPIRoleConfigInput{}
	return &this
}

// GetConfig returns the Config field value
func (o *BackupAPIRoleConfigInput) GetConfig() BackupRoleConfig {
	if o == nil {
		var ret BackupRoleConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *BackupAPIRoleConfigInput) GetConfigOk() (*BackupRoleConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *BackupAPIRoleConfigInput) SetConfig(v BackupRoleConfig) {
	o.Config = v
}

func (o BackupAPIRoleConfigInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableBackupAPIRoleConfigInput struct {
	value *BackupAPIRoleConfigInput
	isSet bool
}

func (v NullableBackupAPIRoleConfigInput) Get() *BackupAPIRoleConfigInput {
	return v.value
}

func (v *NullableBackupAPIRoleConfigInput) Set(val *BackupAPIRoleConfigInput) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupAPIRoleConfigInput) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupAPIRoleConfigInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupAPIRoleConfigInput(val *BackupAPIRoleConfigInput) *NullableBackupAPIRoleConfigInput {
	return &NullableBackupAPIRoleConfigInput{value: val, isSet: true}
}

func (v NullableBackupAPIRoleConfigInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupAPIRoleConfigInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
