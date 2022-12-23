/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.2.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TypesAPIMetricsRecord struct for TypesAPIMetricsRecord
type TypesAPIMetricsRecord struct {
	Handler string `json:"handler"`
	Node    string `json:"node"`
	Time    string `json:"time"`
	Value   int32  `json:"value"`
}

// NewTypesAPIMetricsRecord instantiates a new TypesAPIMetricsRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesAPIMetricsRecord(handler string, node string, time string, value int32) *TypesAPIMetricsRecord {
	this := TypesAPIMetricsRecord{}
	this.Handler = handler
	this.Node = node
	this.Time = time
	this.Value = value
	return &this
}

// NewTypesAPIMetricsRecordWithDefaults instantiates a new TypesAPIMetricsRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesAPIMetricsRecordWithDefaults() *TypesAPIMetricsRecord {
	this := TypesAPIMetricsRecord{}
	return &this
}

// GetHandler returns the Handler field value
func (o *TypesAPIMetricsRecord) GetHandler() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handler
}

// GetHandlerOk returns a tuple with the Handler field value
// and a boolean to check if the value has been set.
func (o *TypesAPIMetricsRecord) GetHandlerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handler, true
}

// SetHandler sets field value
func (o *TypesAPIMetricsRecord) SetHandler(v string) {
	o.Handler = v
}

// GetNode returns the Node field value
func (o *TypesAPIMetricsRecord) GetNode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Node
}

// GetNodeOk returns a tuple with the Node field value
// and a boolean to check if the value has been set.
func (o *TypesAPIMetricsRecord) GetNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Node, true
}

// SetNode sets field value
func (o *TypesAPIMetricsRecord) SetNode(v string) {
	o.Node = v
}

// GetTime returns the Time field value
func (o *TypesAPIMetricsRecord) GetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *TypesAPIMetricsRecord) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *TypesAPIMetricsRecord) SetTime(v string) {
	o.Time = v
}

// GetValue returns the Value field value
func (o *TypesAPIMetricsRecord) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TypesAPIMetricsRecord) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TypesAPIMetricsRecord) SetValue(v int32) {
	o.Value = v
}

func (o TypesAPIMetricsRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handler"] = o.Handler
	}
	if true {
		toSerialize["node"] = o.Node
	}
	if true {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableTypesAPIMetricsRecord struct {
	value *TypesAPIMetricsRecord
	isSet bool
}

func (v NullableTypesAPIMetricsRecord) Get() *TypesAPIMetricsRecord {
	return v.value
}

func (v *NullableTypesAPIMetricsRecord) Set(val *TypesAPIMetricsRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesAPIMetricsRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesAPIMetricsRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesAPIMetricsRecord(val *TypesAPIMetricsRecord) *NullableTypesAPIMetricsRecord {
	return &NullableTypesAPIMetricsRecord{value: val, isSet: true}
}

func (v NullableTypesAPIMetricsRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesAPIMetricsRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
