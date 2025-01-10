/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AuthAPIMeOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthAPIMeOutput{}

// AuthAPIMeOutput struct for AuthAPIMeOutput
type AuthAPIMeOutput struct {
	Authenticated bool             `json:"authenticated"`
	Permissions   []AuthPermission `json:"permissions"`
	Username      string           `json:"username"`
}

// NewAuthAPIMeOutput instantiates a new AuthAPIMeOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAPIMeOutput(authenticated bool, permissions []AuthPermission, username string) *AuthAPIMeOutput {
	this := AuthAPIMeOutput{}
	this.Authenticated = authenticated
	this.Permissions = permissions
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

// GetPermissions returns the Permissions field value
// If the value is explicit nil, the zero value for []AuthPermission will be returned
func (o *AuthAPIMeOutput) GetPermissions() []AuthPermission {
	if o == nil {
		var ret []AuthPermission
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthAPIMeOutput) GetPermissionsOk() ([]AuthPermission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *AuthAPIMeOutput) SetPermissions(v []AuthPermission) {
	o.Permissions = v
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthAPIMeOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authenticated"] = o.Authenticated
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	toSerialize["username"] = o.Username
	return toSerialize, nil
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
