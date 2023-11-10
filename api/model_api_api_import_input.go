/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.6.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ApiAPIImportInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAPIImportInput{}

// ApiAPIImportInput struct for ApiAPIImportInput
type ApiAPIImportInput struct {
	Entries []ApiAPITransportEntry `json:"entries,omitempty"`
}

// NewApiAPIImportInput instantiates a new ApiAPIImportInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPIImportInput() *ApiAPIImportInput {
	this := ApiAPIImportInput{}
	return &this
}

// NewApiAPIImportInputWithDefaults instantiates a new ApiAPIImportInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPIImportInputWithDefaults() *ApiAPIImportInput {
	this := ApiAPIImportInput{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiAPIImportInput) GetEntries() []ApiAPITransportEntry {
	if o == nil {
		var ret []ApiAPITransportEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiAPIImportInput) GetEntriesOk() ([]ApiAPITransportEntry, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *ApiAPIImportInput) HasEntries() bool {
	if o != nil && IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []ApiAPITransportEntry and assigns it to the Entries field.
func (o *ApiAPIImportInput) SetEntries(v []ApiAPITransportEntry) {
	o.Entries = v
}

func (o ApiAPIImportInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAPIImportInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	return toSerialize, nil
}

type NullableApiAPIImportInput struct {
	value *ApiAPIImportInput
	isSet bool
}

func (v NullableApiAPIImportInput) Get() *ApiAPIImportInput {
	return v.value
}

func (v *NullableApiAPIImportInput) Set(val *ApiAPIImportInput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPIImportInput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPIImportInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPIImportInput(val *ApiAPIImportInput) *NullableApiAPIImportInput {
	return &NullableApiAPIImportInput{value: val, isSet: true}
}

func (v NullableApiAPIImportInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPIImportInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
