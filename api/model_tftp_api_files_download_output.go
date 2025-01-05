/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.20.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TftpAPIFilesDownloadOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TftpAPIFilesDownloadOutput{}

// TftpAPIFilesDownloadOutput struct for TftpAPIFilesDownloadOutput
type TftpAPIFilesDownloadOutput struct {
	Data string `json:"data"`
}

// NewTftpAPIFilesDownloadOutput instantiates a new TftpAPIFilesDownloadOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTftpAPIFilesDownloadOutput(data string) *TftpAPIFilesDownloadOutput {
	this := TftpAPIFilesDownloadOutput{}
	this.Data = data
	return &this
}

// NewTftpAPIFilesDownloadOutputWithDefaults instantiates a new TftpAPIFilesDownloadOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTftpAPIFilesDownloadOutputWithDefaults() *TftpAPIFilesDownloadOutput {
	this := TftpAPIFilesDownloadOutput{}
	return &this
}

// GetData returns the Data field value
func (o *TftpAPIFilesDownloadOutput) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TftpAPIFilesDownloadOutput) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TftpAPIFilesDownloadOutput) SetData(v string) {
	o.Data = v
}

func (o TftpAPIFilesDownloadOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TftpAPIFilesDownloadOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableTftpAPIFilesDownloadOutput struct {
	value *TftpAPIFilesDownloadOutput
	isSet bool
}

func (v NullableTftpAPIFilesDownloadOutput) Get() *TftpAPIFilesDownloadOutput {
	return v.value
}

func (v *NullableTftpAPIFilesDownloadOutput) Set(val *TftpAPIFilesDownloadOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableTftpAPIFilesDownloadOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableTftpAPIFilesDownloadOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTftpAPIFilesDownloadOutput(val *TftpAPIFilesDownloadOutput) *NullableTftpAPIFilesDownloadOutput {
	return &NullableTftpAPIFilesDownloadOutput{value: val, isSet: true}
}

func (v NullableTftpAPIFilesDownloadOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTftpAPIFilesDownloadOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
