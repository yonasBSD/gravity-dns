/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.23.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AuthAPILoginInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthAPILoginInput{}

// AuthAPILoginInput struct for AuthAPILoginInput
type AuthAPILoginInput struct {
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewAuthAPILoginInput instantiates a new AuthAPILoginInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAPILoginInput() *AuthAPILoginInput {
	this := AuthAPILoginInput{}
	return &this
}

// NewAuthAPILoginInputWithDefaults instantiates a new AuthAPILoginInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAPILoginInputWithDefaults() *AuthAPILoginInput {
	this := AuthAPILoginInput{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AuthAPILoginInput) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAPILoginInput) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AuthAPILoginInput) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AuthAPILoginInput) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AuthAPILoginInput) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAPILoginInput) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AuthAPILoginInput) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AuthAPILoginInput) SetUsername(v string) {
	o.Username = &v
}

func (o AuthAPILoginInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthAPILoginInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableAuthAPILoginInput struct {
	value *AuthAPILoginInput
	isSet bool
}

func (v NullableAuthAPILoginInput) Get() *AuthAPILoginInput {
	return v.value
}

func (v *NullableAuthAPILoginInput) Set(val *AuthAPILoginInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAPILoginInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAPILoginInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAPILoginInput(val *AuthAPILoginInput) *NullableAuthAPILoginInput {
	return &NullableAuthAPILoginInput{value: val, isSet: true}
}

func (v NullableAuthAPILoginInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAPILoginInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
