/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.13.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DnsAPIRoleConfigInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsAPIRoleConfigInput{}

// DnsAPIRoleConfigInput struct for DnsAPIRoleConfigInput
type DnsAPIRoleConfigInput struct {
	Config DnsRoleConfig `json:"config"`
}

// NewDnsAPIRoleConfigInput instantiates a new DnsAPIRoleConfigInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsAPIRoleConfigInput(config DnsRoleConfig) *DnsAPIRoleConfigInput {
	this := DnsAPIRoleConfigInput{}
	this.Config = config
	return &this
}

// NewDnsAPIRoleConfigInputWithDefaults instantiates a new DnsAPIRoleConfigInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsAPIRoleConfigInputWithDefaults() *DnsAPIRoleConfigInput {
	this := DnsAPIRoleConfigInput{}
	return &this
}

// GetConfig returns the Config field value
func (o *DnsAPIRoleConfigInput) GetConfig() DnsRoleConfig {
	if o == nil {
		var ret DnsRoleConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *DnsAPIRoleConfigInput) GetConfigOk() (*DnsRoleConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *DnsAPIRoleConfigInput) SetConfig(v DnsRoleConfig) {
	o.Config = v
}

func (o DnsAPIRoleConfigInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsAPIRoleConfigInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	return toSerialize, nil
}

type NullableDnsAPIRoleConfigInput struct {
	value *DnsAPIRoleConfigInput
	isSet bool
}

func (v NullableDnsAPIRoleConfigInput) Get() *DnsAPIRoleConfigInput {
	return v.value
}

func (v *NullableDnsAPIRoleConfigInput) Set(val *DnsAPIRoleConfigInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsAPIRoleConfigInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsAPIRoleConfigInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsAPIRoleConfigInput(val *DnsAPIRoleConfigInput) *NullableDnsAPIRoleConfigInput {
	return &NullableDnsAPIRoleConfigInput{value: val, isSet: true}
}

func (v NullableDnsAPIRoleConfigInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsAPIRoleConfigInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
