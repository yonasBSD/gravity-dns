/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.25.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ApiAPIMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAPIMember{}

// ApiAPIMember struct for ApiAPIMember
type ApiAPIMember struct {
	Id   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewApiAPIMember instantiates a new ApiAPIMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPIMember() *ApiAPIMember {
	this := ApiAPIMember{}
	return &this
}

// NewApiAPIMemberWithDefaults instantiates a new ApiAPIMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPIMemberWithDefaults() *ApiAPIMember {
	this := ApiAPIMember{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiAPIMember) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIMember) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiAPIMember) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ApiAPIMember) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiAPIMember) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIMember) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiAPIMember) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiAPIMember) SetName(v string) {
	o.Name = &v
}

func (o ApiAPIMember) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAPIMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableApiAPIMember struct {
	value *ApiAPIMember
	isSet bool
}

func (v NullableApiAPIMember) Get() *ApiAPIMember {
	return v.value
}

func (v *NullableApiAPIMember) Set(val *ApiAPIMember) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPIMember) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPIMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPIMember(val *ApiAPIMember) *NullableApiAPIMember {
	return &NullableApiAPIMember{value: val, isSet: true}
}

func (v NullableApiAPIMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPIMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
