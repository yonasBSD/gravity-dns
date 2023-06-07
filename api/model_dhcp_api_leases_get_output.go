/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.6.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DhcpAPILeasesGetOutput struct for DhcpAPILeasesGetOutput
type DhcpAPILeasesGetOutput struct {
	Leases []DhcpAPILease `json:"leases"`
}

// NewDhcpAPILeasesGetOutput instantiates a new DhcpAPILeasesGetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpAPILeasesGetOutput(leases []DhcpAPILease) *DhcpAPILeasesGetOutput {
	this := DhcpAPILeasesGetOutput{}
	this.Leases = leases
	return &this
}

// NewDhcpAPILeasesGetOutputWithDefaults instantiates a new DhcpAPILeasesGetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpAPILeasesGetOutputWithDefaults() *DhcpAPILeasesGetOutput {
	this := DhcpAPILeasesGetOutput{}
	return &this
}

// GetLeases returns the Leases field value
// If the value is explicit nil, the zero value for []DhcpAPILease will be returned
func (o *DhcpAPILeasesGetOutput) GetLeases() []DhcpAPILease {
	if o == nil {
		var ret []DhcpAPILease
		return ret
	}

	return o.Leases
}

// GetLeasesOk returns a tuple with the Leases field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DhcpAPILeasesGetOutput) GetLeasesOk() ([]DhcpAPILease, bool) {
	if o == nil || o.Leases == nil {
		return nil, false
	}
	return o.Leases, true
}

// SetLeases sets field value
func (o *DhcpAPILeasesGetOutput) SetLeases(v []DhcpAPILease) {
	o.Leases = v
}

func (o DhcpAPILeasesGetOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Leases != nil {
		toSerialize["leases"] = o.Leases
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpAPILeasesGetOutput struct {
	value *DhcpAPILeasesGetOutput
	isSet bool
}

func (v NullableDhcpAPILeasesGetOutput) Get() *DhcpAPILeasesGetOutput {
	return v.value
}

func (v *NullableDhcpAPILeasesGetOutput) Set(val *DhcpAPILeasesGetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpAPILeasesGetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpAPILeasesGetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpAPILeasesGetOutput(val *DhcpAPILeasesGetOutput) *NullableDhcpAPILeasesGetOutput {
	return &NullableDhcpAPILeasesGetOutput{value: val, isSet: true}
}

func (v NullableDhcpAPILeasesGetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpAPILeasesGetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
