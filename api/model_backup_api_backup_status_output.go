/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BackupAPIBackupStatusOutput struct for BackupAPIBackupStatusOutput
type BackupAPIBackupStatusOutput struct {
	Status []BackupAPIBackupStatus `json:"status"`
}

// NewBackupAPIBackupStatusOutput instantiates a new BackupAPIBackupStatusOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupAPIBackupStatusOutput(status []BackupAPIBackupStatus) *BackupAPIBackupStatusOutput {
	this := BackupAPIBackupStatusOutput{}
	this.Status = status
	return &this
}

// NewBackupAPIBackupStatusOutputWithDefaults instantiates a new BackupAPIBackupStatusOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupAPIBackupStatusOutputWithDefaults() *BackupAPIBackupStatusOutput {
	this := BackupAPIBackupStatusOutput{}
	return &this
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for []BackupAPIBackupStatus will be returned
func (o *BackupAPIBackupStatusOutput) GetStatus() []BackupAPIBackupStatus {
	if o == nil {
		var ret []BackupAPIBackupStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupAPIBackupStatusOutput) GetStatusOk() ([]BackupAPIBackupStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// SetStatus sets field value
func (o *BackupAPIBackupStatusOutput) SetStatus(v []BackupAPIBackupStatus) {
	o.Status = v
}

func (o BackupAPIBackupStatusOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableBackupAPIBackupStatusOutput struct {
	value *BackupAPIBackupStatusOutput
	isSet bool
}

func (v NullableBackupAPIBackupStatusOutput) Get() *BackupAPIBackupStatusOutput {
	return v.value
}

func (v *NullableBackupAPIBackupStatusOutput) Set(val *BackupAPIBackupStatusOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupAPIBackupStatusOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupAPIBackupStatusOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupAPIBackupStatusOutput(val *BackupAPIBackupStatusOutput) *NullableBackupAPIBackupStatusOutput {
	return &NullableBackupAPIBackupStatusOutput{value: val, isSet: true}
}

func (v NullableBackupAPIBackupStatusOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupAPIBackupStatusOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
