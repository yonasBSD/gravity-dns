/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.6.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DnsAPIRoleConfigOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsAPIRoleConfigOutput{}

// DnsAPIRoleConfigOutput struct for DnsAPIRoleConfigOutput
type DnsAPIRoleConfigOutput struct {
	Config DnsRoleConfig `json:"config"`
}

// NewDnsAPIRoleConfigOutput instantiates a new DnsAPIRoleConfigOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsAPIRoleConfigOutput(config DnsRoleConfig) *DnsAPIRoleConfigOutput {
	this := DnsAPIRoleConfigOutput{}
	this.Config = config
	return &this
}

// NewDnsAPIRoleConfigOutputWithDefaults instantiates a new DnsAPIRoleConfigOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsAPIRoleConfigOutputWithDefaults() *DnsAPIRoleConfigOutput {
	this := DnsAPIRoleConfigOutput{}
	return &this
}

// GetConfig returns the Config field value
func (o *DnsAPIRoleConfigOutput) GetConfig() DnsRoleConfig {
	if o == nil {
		var ret DnsRoleConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *DnsAPIRoleConfigOutput) GetConfigOk() (*DnsRoleConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *DnsAPIRoleConfigOutput) SetConfig(v DnsRoleConfig) {
	o.Config = v
}

func (o DnsAPIRoleConfigOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsAPIRoleConfigOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	return toSerialize, nil
}

type NullableDnsAPIRoleConfigOutput struct {
	value *DnsAPIRoleConfigOutput
	isSet bool
}

func (v NullableDnsAPIRoleConfigOutput) Get() *DnsAPIRoleConfigOutput {
	return v.value
}

func (v *NullableDnsAPIRoleConfigOutput) Set(val *DnsAPIRoleConfigOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsAPIRoleConfigOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsAPIRoleConfigOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsAPIRoleConfigOutput(val *DnsAPIRoleConfigOutput) *NullableDnsAPIRoleConfigOutput {
	return &NullableDnsAPIRoleConfigOutput{value: val, isSet: true}
}

func (v NullableDnsAPIRoleConfigOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsAPIRoleConfigOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
