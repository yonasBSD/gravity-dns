/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// BackupAPIBackupStatus struct for BackupAPIBackupStatus
type BackupAPIBackupStatus struct {
	Duration *int32     `json:"duration,omitempty"`
	Error    *string    `json:"error,omitempty"`
	Filename *string    `json:"filename,omitempty"`
	Node     *string    `json:"node,omitempty"`
	Size     *int32     `json:"size,omitempty"`
	Status   *string    `json:"status,omitempty"`
	Time     *time.Time `json:"time,omitempty"`
}

// NewBackupAPIBackupStatus instantiates a new BackupAPIBackupStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupAPIBackupStatus() *BackupAPIBackupStatus {
	this := BackupAPIBackupStatus{}
	return &this
}

// NewBackupAPIBackupStatusWithDefaults instantiates a new BackupAPIBackupStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupAPIBackupStatusWithDefaults() *BackupAPIBackupStatus {
	this := BackupAPIBackupStatus{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *BackupAPIBackupStatus) GetDuration() int32 {
	if o == nil || o.Duration == nil {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupAPIBackupStatus) GetDurationOk() (*int32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *BackupAPIBackupStatus) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *BackupAPIBackupStatus) SetDuration(v int32) {
	o.Duration = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BackupAPIBackupStatus) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupAPIBackupStatus) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BackupAPIBackupStatus) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *BackupAPIBackupStatus) SetError(v string) {
	o.Error = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *BackupAPIBackupStatus) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupAPIBackupStatus) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *BackupAPIBackupStatus) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *BackupAPIBackupStatus) SetFilename(v string) {
	o.Filename = &v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *BackupAPIBackupStatus) GetNode() string {
	if o == nil || o.Node == nil {
		var ret string
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupAPIBackupStatus) GetNodeOk() (*string, bool) {
	if o == nil || o.Node == nil {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *BackupAPIBackupStatus) HasNode() bool {
	if o != nil && o.Node != nil {
		return true
	}

	return false
}

// SetNode gets a reference to the given string and assigns it to the Node field.
func (o *BackupAPIBackupStatus) SetNode(v string) {
	o.Node = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BackupAPIBackupStatus) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupAPIBackupStatus) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BackupAPIBackupStatus) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *BackupAPIBackupStatus) SetSize(v int32) {
	o.Size = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BackupAPIBackupStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupAPIBackupStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BackupAPIBackupStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BackupAPIBackupStatus) SetStatus(v string) {
	o.Status = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *BackupAPIBackupStatus) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupAPIBackupStatus) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *BackupAPIBackupStatus) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *BackupAPIBackupStatus) SetTime(v time.Time) {
	o.Time = &v
}

func (o BackupAPIBackupStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.Node != nil {
		toSerialize["node"] = o.Node
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableBackupAPIBackupStatus struct {
	value *BackupAPIBackupStatus
	isSet bool
}

func (v NullableBackupAPIBackupStatus) Get() *BackupAPIBackupStatus {
	return v.value
}

func (v *NullableBackupAPIBackupStatus) Set(val *BackupAPIBackupStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupAPIBackupStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupAPIBackupStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupAPIBackupStatus(val *BackupAPIBackupStatus) *NullableBackupAPIBackupStatus {
	return &NullableBackupAPIBackupStatus{value: val, isSet: true}
}

func (v NullableBackupAPIBackupStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupAPIBackupStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
