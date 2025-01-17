/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.25.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TftpAPIFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TftpAPIFile{}

// TftpAPIFile struct for TftpAPIFile
type TftpAPIFile struct {
	Host      string `json:"host"`
	Name      string `json:"name"`
	SizeBytes int32  `json:"sizeBytes"`
}

// NewTftpAPIFile instantiates a new TftpAPIFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTftpAPIFile(host string, name string, sizeBytes int32) *TftpAPIFile {
	this := TftpAPIFile{}
	this.Host = host
	this.Name = name
	this.SizeBytes = sizeBytes
	return &this
}

// NewTftpAPIFileWithDefaults instantiates a new TftpAPIFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTftpAPIFileWithDefaults() *TftpAPIFile {
	this := TftpAPIFile{}
	return &this
}

// GetHost returns the Host field value
func (o *TftpAPIFile) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *TftpAPIFile) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *TftpAPIFile) SetHost(v string) {
	o.Host = v
}

// GetName returns the Name field value
func (o *TftpAPIFile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TftpAPIFile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TftpAPIFile) SetName(v string) {
	o.Name = v
}

// GetSizeBytes returns the SizeBytes field value
func (o *TftpAPIFile) GetSizeBytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SizeBytes
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value
// and a boolean to check if the value has been set.
func (o *TftpAPIFile) GetSizeBytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeBytes, true
}

// SetSizeBytes sets field value
func (o *TftpAPIFile) SetSizeBytes(v int32) {
	o.SizeBytes = v
}

func (o TftpAPIFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TftpAPIFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["name"] = o.Name
	toSerialize["sizeBytes"] = o.SizeBytes
	return toSerialize, nil
}

type NullableTftpAPIFile struct {
	value *TftpAPIFile
	isSet bool
}

func (v NullableTftpAPIFile) Get() *TftpAPIFile {
	return v.value
}

func (v *NullableTftpAPIFile) Set(val *TftpAPIFile) {
	v.value = val
	v.isSet = true
}

func (v NullableTftpAPIFile) IsSet() bool {
	return v.isSet
}

func (v *NullableTftpAPIFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTftpAPIFile(val *TftpAPIFile) *NullableTftpAPIFile {
	return &NullableTftpAPIFile{value: val, isSet: true}
}

func (v NullableTftpAPIFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTftpAPIFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
