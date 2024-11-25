/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.16.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TftpAPIFilesPutInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TftpAPIFilesPutInput{}

// TftpAPIFilesPutInput struct for TftpAPIFilesPutInput
type TftpAPIFilesPutInput struct {
	Data string `json:"data"`
	Host string `json:"host"`
	Name string `json:"name"`
}

// NewTftpAPIFilesPutInput instantiates a new TftpAPIFilesPutInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTftpAPIFilesPutInput(data string, host string, name string) *TftpAPIFilesPutInput {
	this := TftpAPIFilesPutInput{}
	this.Data = data
	this.Host = host
	this.Name = name
	return &this
}

// NewTftpAPIFilesPutInputWithDefaults instantiates a new TftpAPIFilesPutInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTftpAPIFilesPutInputWithDefaults() *TftpAPIFilesPutInput {
	this := TftpAPIFilesPutInput{}
	return &this
}

// GetData returns the Data field value
func (o *TftpAPIFilesPutInput) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TftpAPIFilesPutInput) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TftpAPIFilesPutInput) SetData(v string) {
	o.Data = v
}

// GetHost returns the Host field value
func (o *TftpAPIFilesPutInput) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *TftpAPIFilesPutInput) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *TftpAPIFilesPutInput) SetHost(v string) {
	o.Host = v
}

// GetName returns the Name field value
func (o *TftpAPIFilesPutInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TftpAPIFilesPutInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TftpAPIFilesPutInput) SetName(v string) {
	o.Name = v
}

func (o TftpAPIFilesPutInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TftpAPIFilesPutInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["host"] = o.Host
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableTftpAPIFilesPutInput struct {
	value *TftpAPIFilesPutInput
	isSet bool
}

func (v NullableTftpAPIFilesPutInput) Get() *TftpAPIFilesPutInput {
	return v.value
}

func (v *NullableTftpAPIFilesPutInput) Set(val *TftpAPIFilesPutInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTftpAPIFilesPutInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTftpAPIFilesPutInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTftpAPIFilesPutInput(val *TftpAPIFilesPutInput) *NullableTftpAPIFilesPutInput {
	return &NullableTftpAPIFilesPutInput{value: val, isSet: true}
}

func (v NullableTftpAPIFilesPutInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTftpAPIFilesPutInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
