/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.4.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DnsAPIRecordsGetOutput struct for DnsAPIRecordsGetOutput
type DnsAPIRecordsGetOutput struct {
	Records []DnsAPIRecord `json:"records"`
}

// NewDnsAPIRecordsGetOutput instantiates a new DnsAPIRecordsGetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsAPIRecordsGetOutput(records []DnsAPIRecord) *DnsAPIRecordsGetOutput {
	this := DnsAPIRecordsGetOutput{}
	this.Records = records
	return &this
}

// NewDnsAPIRecordsGetOutputWithDefaults instantiates a new DnsAPIRecordsGetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsAPIRecordsGetOutputWithDefaults() *DnsAPIRecordsGetOutput {
	this := DnsAPIRecordsGetOutput{}
	return &this
}

// GetRecords returns the Records field value
// If the value is explicit nil, the zero value for []DnsAPIRecord will be returned
func (o *DnsAPIRecordsGetOutput) GetRecords() []DnsAPIRecord {
	if o == nil {
		var ret []DnsAPIRecord
		return ret
	}

	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnsAPIRecordsGetOutput) GetRecordsOk() ([]DnsAPIRecord, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return o.Records, true
}

// SetRecords sets field value
func (o *DnsAPIRecordsGetOutput) SetRecords(v []DnsAPIRecord) {
	o.Records = v
}

func (o DnsAPIRecordsGetOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	return json.Marshal(toSerialize)
}

type NullableDnsAPIRecordsGetOutput struct {
	value *DnsAPIRecordsGetOutput
	isSet bool
}

func (v NullableDnsAPIRecordsGetOutput) Get() *DnsAPIRecordsGetOutput {
	return v.value
}

func (v *NullableDnsAPIRecordsGetOutput) Set(val *DnsAPIRecordsGetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsAPIRecordsGetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsAPIRecordsGetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsAPIRecordsGetOutput(val *DnsAPIRecordsGetOutput) *NullableDnsAPIRecordsGetOutput {
	return &NullableDnsAPIRecordsGetOutput{value: val, isSet: true}
}

func (v NullableDnsAPIRecordsGetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsAPIRecordsGetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
