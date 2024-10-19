/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DhcpAPIScopeStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DhcpAPIScopeStatistics{}

// DhcpAPIScopeStatistics struct for DhcpAPIScopeStatistics
type DhcpAPIScopeStatistics struct {
	Usable int32 `json:"usable"`
	Used   int32 `json:"used"`
}

// NewDhcpAPIScopeStatistics instantiates a new DhcpAPIScopeStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpAPIScopeStatistics(usable int32, used int32) *DhcpAPIScopeStatistics {
	this := DhcpAPIScopeStatistics{}
	this.Usable = usable
	this.Used = used
	return &this
}

// NewDhcpAPIScopeStatisticsWithDefaults instantiates a new DhcpAPIScopeStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpAPIScopeStatisticsWithDefaults() *DhcpAPIScopeStatistics {
	this := DhcpAPIScopeStatistics{}
	return &this
}

// GetUsable returns the Usable field value
func (o *DhcpAPIScopeStatistics) GetUsable() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Usable
}

// GetUsableOk returns a tuple with the Usable field value
// and a boolean to check if the value has been set.
func (o *DhcpAPIScopeStatistics) GetUsableOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usable, true
}

// SetUsable sets field value
func (o *DhcpAPIScopeStatistics) SetUsable(v int32) {
	o.Usable = v
}

// GetUsed returns the Used field value
func (o *DhcpAPIScopeStatistics) GetUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *DhcpAPIScopeStatistics) GetUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value
func (o *DhcpAPIScopeStatistics) SetUsed(v int32) {
	o.Used = v
}

func (o DhcpAPIScopeStatistics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DhcpAPIScopeStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["usable"] = o.Usable
	toSerialize["used"] = o.Used
	return toSerialize, nil
}

type NullableDhcpAPIScopeStatistics struct {
	value *DhcpAPIScopeStatistics
	isSet bool
}

func (v NullableDhcpAPIScopeStatistics) Get() *DhcpAPIScopeStatistics {
	return v.value
}

func (v *NullableDhcpAPIScopeStatistics) Set(val *DhcpAPIScopeStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpAPIScopeStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpAPIScopeStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpAPIScopeStatistics(val *DhcpAPIScopeStatistics) *NullableDhcpAPIScopeStatistics {
	return &NullableDhcpAPIScopeStatistics{value: val, isSet: true}
}

func (v NullableDhcpAPIScopeStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpAPIScopeStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}