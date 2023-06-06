/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.6.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthAPIMeOutput struct for AuthAPIMeOutput
type AuthAPIMeOutput struct {
	Authenticated bool   `json:"authenticated"`
	Username      string `json:"username"`
}

// NewAuthAPIMeOutput instantiates a new AuthAPIMeOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAPIMeOutput(authenticated bool, username string) *AuthAPIMeOutput {
	this := AuthAPIMeOutput{}
	this.Authenticated = authenticated
	this.Username = username
	return &this
}

// NewAuthAPIMeOutputWithDefaults instantiates a new AuthAPIMeOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAPIMeOutputWithDefaults() *AuthAPIMeOutput {
	this := AuthAPIMeOutput{}
	return &this
}

// GetAuthenticated returns the Authenticated field value
func (o *AuthAPIMeOutput) GetAuthenticated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Authenticated
}

// GetAuthenticatedOk returns a tuple with the Authenticated field value
// and a boolean to check if the value has been set.
func (o *AuthAPIMeOutput) GetAuthenticatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authenticated, true
}

// SetAuthenticated sets field value
func (o *AuthAPIMeOutput) SetAuthenticated(v bool) {
	o.Authenticated = v
}

// GetUsername returns the Username field value
func (o *AuthAPIMeOutput) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AuthAPIMeOutput) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AuthAPIMeOutput) SetUsername(v string) {
	o.Username = v
}

func (o AuthAPIMeOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authenticated"] = o.Authenticated
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableAuthAPIMeOutput struct {
	value *AuthAPIMeOutput
	isSet bool
}

func (v NullableAuthAPIMeOutput) Get() *AuthAPIMeOutput {
	return v.value
}

func (v *NullableAuthAPIMeOutput) Set(val *AuthAPIMeOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAPIMeOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAPIMeOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAPIMeOutput(val *AuthAPIMeOutput) *NullableAuthAPIMeOutput {
	return &NullableAuthAPIMeOutput{value: val, isSet: true}
}

func (v NullableAuthAPIMeOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAPIMeOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
