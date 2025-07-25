/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.27.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AuthAPIUsersGetOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthAPIUsersGetOutput{}

// AuthAPIUsersGetOutput struct for AuthAPIUsersGetOutput
type AuthAPIUsersGetOutput struct {
	Users []AuthAPIUser `json:"users"`
}

// NewAuthAPIUsersGetOutput instantiates a new AuthAPIUsersGetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAPIUsersGetOutput(users []AuthAPIUser) *AuthAPIUsersGetOutput {
	this := AuthAPIUsersGetOutput{}
	this.Users = users
	return &this
}

// NewAuthAPIUsersGetOutputWithDefaults instantiates a new AuthAPIUsersGetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAPIUsersGetOutputWithDefaults() *AuthAPIUsersGetOutput {
	this := AuthAPIUsersGetOutput{}
	return &this
}

// GetUsers returns the Users field value
// If the value is explicit nil, the zero value for []AuthAPIUser will be returned
func (o *AuthAPIUsersGetOutput) GetUsers() []AuthAPIUser {
	if o == nil {
		var ret []AuthAPIUser
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthAPIUsersGetOutput) GetUsersOk() ([]AuthAPIUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *AuthAPIUsersGetOutput) SetUsers(v []AuthAPIUser) {
	o.Users = v
}

func (o AuthAPIUsersGetOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthAPIUsersGetOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableAuthAPIUsersGetOutput struct {
	value *AuthAPIUsersGetOutput
	isSet bool
}

func (v NullableAuthAPIUsersGetOutput) Get() *AuthAPIUsersGetOutput {
	return v.value
}

func (v *NullableAuthAPIUsersGetOutput) Set(val *AuthAPIUsersGetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAPIUsersGetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAPIUsersGetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAPIUsersGetOutput(val *AuthAPIUsersGetOutput) *NullableAuthAPIUsersGetOutput {
	return &NullableAuthAPIUsersGetOutput{value: val, isSet: true}
}

func (v NullableAuthAPIUsersGetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAPIUsersGetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
