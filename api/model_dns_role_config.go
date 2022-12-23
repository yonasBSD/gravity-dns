/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DnsRoleConfig struct for DnsRoleConfig
type DnsRoleConfig struct {
	Port *int32 `json:"port,omitempty"`
}

// NewDnsRoleConfig instantiates a new DnsRoleConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRoleConfig() *DnsRoleConfig {
	this := DnsRoleConfig{}
	return &this
}

// NewDnsRoleConfigWithDefaults instantiates a new DnsRoleConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRoleConfigWithDefaults() *DnsRoleConfig {
	this := DnsRoleConfig{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *DnsRoleConfig) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRoleConfig) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *DnsRoleConfig) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *DnsRoleConfig) SetPort(v int32) {
	o.Port = &v
}

func (o DnsRoleConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableDnsRoleConfig struct {
	value *DnsRoleConfig
	isSet bool
}

func (v NullableDnsRoleConfig) Get() *DnsRoleConfig {
	return v.value
}

func (v *NullableDnsRoleConfig) Set(val *DnsRoleConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRoleConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRoleConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRoleConfig(val *DnsRoleConfig) *NullableDnsRoleConfig {
	return &NullableDnsRoleConfig{value: val, isSet: true}
}

func (v NullableDnsRoleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRoleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}