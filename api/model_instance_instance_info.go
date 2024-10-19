/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the InstanceInstanceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceInstanceInfo{}

// InstanceInstanceInfo struct for InstanceInstanceInfo
type InstanceInstanceInfo struct {
	Identifier string `json:"identifier"`
	Ip         string `json:"ip"`
	Roles      string `json:"roles"`
	Version    string `json:"version"`
}

// NewInstanceInstanceInfo instantiates a new InstanceInstanceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceInstanceInfo(identifier string, ip string, roles string, version string) *InstanceInstanceInfo {
	this := InstanceInstanceInfo{}
	this.Identifier = identifier
	this.Ip = ip
	this.Roles = roles
	this.Version = version
	return &this
}

// NewInstanceInstanceInfoWithDefaults instantiates a new InstanceInstanceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceInstanceInfoWithDefaults() *InstanceInstanceInfo {
	this := InstanceInstanceInfo{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *InstanceInstanceInfo) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *InstanceInstanceInfo) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *InstanceInstanceInfo) SetIdentifier(v string) {
	o.Identifier = v
}

// GetIp returns the Ip field value
func (o *InstanceInstanceInfo) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *InstanceInstanceInfo) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *InstanceInstanceInfo) SetIp(v string) {
	o.Ip = v
}

// GetRoles returns the Roles field value
func (o *InstanceInstanceInfo) GetRoles() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *InstanceInstanceInfo) GetRolesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *InstanceInstanceInfo) SetRoles(v string) {
	o.Roles = v
}

// GetVersion returns the Version field value
func (o *InstanceInstanceInfo) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *InstanceInstanceInfo) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *InstanceInstanceInfo) SetVersion(v string) {
	o.Version = v
}

func (o InstanceInstanceInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceInstanceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifier"] = o.Identifier
	toSerialize["ip"] = o.Ip
	toSerialize["roles"] = o.Roles
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableInstanceInstanceInfo struct {
	value *InstanceInstanceInfo
	isSet bool
}

func (v NullableInstanceInstanceInfo) Get() *InstanceInstanceInfo {
	return v.value
}

func (v *NullableInstanceInstanceInfo) Set(val *InstanceInstanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceInstanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceInstanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceInstanceInfo(val *InstanceInstanceInfo) *NullableInstanceInstanceInfo {
	return &NullableInstanceInstanceInfo{value: val, isSet: true}
}

func (v NullableInstanceInstanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceInstanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
