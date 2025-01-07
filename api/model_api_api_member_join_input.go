/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.21.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ApiAPIMemberJoinInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAPIMemberJoinInput{}

// ApiAPIMemberJoinInput struct for ApiAPIMemberJoinInput
type ApiAPIMemberJoinInput struct {
	Identifier *string `json:"identifier,omitempty"`
	Peer       *string `json:"peer,omitempty"`
	Roles      *string `json:"roles,omitempty"`
}

// NewApiAPIMemberJoinInput instantiates a new ApiAPIMemberJoinInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPIMemberJoinInput() *ApiAPIMemberJoinInput {
	this := ApiAPIMemberJoinInput{}
	return &this
}

// NewApiAPIMemberJoinInputWithDefaults instantiates a new ApiAPIMemberJoinInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPIMemberJoinInputWithDefaults() *ApiAPIMemberJoinInput {
	this := ApiAPIMemberJoinInput{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *ApiAPIMemberJoinInput) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIMemberJoinInput) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *ApiAPIMemberJoinInput) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *ApiAPIMemberJoinInput) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetPeer returns the Peer field value if set, zero value otherwise.
func (o *ApiAPIMemberJoinInput) GetPeer() string {
	if o == nil || IsNil(o.Peer) {
		var ret string
		return ret
	}
	return *o.Peer
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIMemberJoinInput) GetPeerOk() (*string, bool) {
	if o == nil || IsNil(o.Peer) {
		return nil, false
	}
	return o.Peer, true
}

// HasPeer returns a boolean if a field has been set.
func (o *ApiAPIMemberJoinInput) HasPeer() bool {
	if o != nil && !IsNil(o.Peer) {
		return true
	}

	return false
}

// SetPeer gets a reference to the given string and assigns it to the Peer field.
func (o *ApiAPIMemberJoinInput) SetPeer(v string) {
	o.Peer = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ApiAPIMemberJoinInput) GetRoles() string {
	if o == nil || IsNil(o.Roles) {
		var ret string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIMemberJoinInput) GetRolesOk() (*string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ApiAPIMemberJoinInput) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given string and assigns it to the Roles field.
func (o *ApiAPIMemberJoinInput) SetRoles(v string) {
	o.Roles = &v
}

func (o ApiAPIMemberJoinInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAPIMemberJoinInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Peer) {
		toSerialize["peer"] = o.Peer
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableApiAPIMemberJoinInput struct {
	value *ApiAPIMemberJoinInput
	isSet bool
}

func (v NullableApiAPIMemberJoinInput) Get() *ApiAPIMemberJoinInput {
	return v.value
}

func (v *NullableApiAPIMemberJoinInput) Set(val *ApiAPIMemberJoinInput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPIMemberJoinInput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPIMemberJoinInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPIMemberJoinInput(val *ApiAPIMemberJoinInput) *NullableApiAPIMemberJoinInput {
	return &NullableApiAPIMemberJoinInput{value: val, isSet: true}
}

func (v NullableApiAPIMemberJoinInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPIMemberJoinInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
