/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DnsAPIZonesGetOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsAPIZonesGetOutput{}

// DnsAPIZonesGetOutput struct for DnsAPIZonesGetOutput
type DnsAPIZonesGetOutput struct {
	Zones []DnsAPIZone `json:"zones"`
}

// NewDnsAPIZonesGetOutput instantiates a new DnsAPIZonesGetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsAPIZonesGetOutput(zones []DnsAPIZone) *DnsAPIZonesGetOutput {
	this := DnsAPIZonesGetOutput{}
	this.Zones = zones
	return &this
}

// NewDnsAPIZonesGetOutputWithDefaults instantiates a new DnsAPIZonesGetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsAPIZonesGetOutputWithDefaults() *DnsAPIZonesGetOutput {
	this := DnsAPIZonesGetOutput{}
	return &this
}

// GetZones returns the Zones field value
// If the value is explicit nil, the zero value for []DnsAPIZone will be returned
func (o *DnsAPIZonesGetOutput) GetZones() []DnsAPIZone {
	if o == nil {
		var ret []DnsAPIZone
		return ret
	}

	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnsAPIZonesGetOutput) GetZonesOk() ([]DnsAPIZone, bool) {
	if o == nil || IsNil(o.Zones) {
		return nil, false
	}
	return o.Zones, true
}

// SetZones sets field value
func (o *DnsAPIZonesGetOutput) SetZones(v []DnsAPIZone) {
	o.Zones = v
}

func (o DnsAPIZonesGetOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsAPIZonesGetOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Zones != nil {
		toSerialize["zones"] = o.Zones
	}
	return toSerialize, nil
}

type NullableDnsAPIZonesGetOutput struct {
	value *DnsAPIZonesGetOutput
	isSet bool
}

func (v NullableDnsAPIZonesGetOutput) Get() *DnsAPIZonesGetOutput {
	return v.value
}

func (v *NullableDnsAPIZonesGetOutput) Set(val *DnsAPIZonesGetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsAPIZonesGetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsAPIZonesGetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsAPIZonesGetOutput(val *DnsAPIZonesGetOutput) *NullableDnsAPIZonesGetOutput {
	return &NullableDnsAPIZonesGetOutput{value: val, isSet: true}
}

func (v NullableDnsAPIZonesGetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsAPIZonesGetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
