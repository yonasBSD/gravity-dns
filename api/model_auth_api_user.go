/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.6.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AuthAPIUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthAPIUser{}

// AuthAPIUser struct for AuthAPIUser
type AuthAPIUser struct {
	Username string `json:"username"`
}

// NewAuthAPIUser instantiates a new AuthAPIUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAPIUser(username string) *AuthAPIUser {
	this := AuthAPIUser{}
	this.Username = username
	return &this
}

// NewAuthAPIUserWithDefaults instantiates a new AuthAPIUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAPIUserWithDefaults() *AuthAPIUser {
	this := AuthAPIUser{}
	return &this
}

// GetUsername returns the Username field value
func (o *AuthAPIUser) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AuthAPIUser) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AuthAPIUser) SetUsername(v string) {
	o.Username = v
}

func (o AuthAPIUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthAPIUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableAuthAPIUser struct {
	value *AuthAPIUser
	isSet bool
}

func (v NullableAuthAPIUser) Get() *AuthAPIUser {
	return v.value
}

func (v *NullableAuthAPIUser) Set(val *AuthAPIUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAPIUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAPIUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAPIUser(val *AuthAPIUser) *NullableAuthAPIUser {
	return &NullableAuthAPIUser{value: val, isSet: true}
}

func (v NullableAuthAPIUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAPIUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
