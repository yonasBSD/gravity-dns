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

// checks if the InstanceAPIInstanceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceAPIInstanceInfo{}

// InstanceAPIInstanceInfo struct for InstanceAPIInstanceInfo
type InstanceAPIInstanceInfo struct {
	BuildHash          string                 `json:"buildHash"`
	Dirs               ExtconfigExtConfigDirs `json:"dirs"`
	InstanceIP         string                 `json:"instanceIP"`
	InstanceIdentifier string                 `json:"instanceIdentifier"`
	Version            string                 `json:"version"`
}

// NewInstanceAPIInstanceInfo instantiates a new InstanceAPIInstanceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceAPIInstanceInfo(buildHash string, dirs ExtconfigExtConfigDirs, instanceIP string, instanceIdentifier string, version string) *InstanceAPIInstanceInfo {
	this := InstanceAPIInstanceInfo{}
	this.BuildHash = buildHash
	this.Dirs = dirs
	this.InstanceIP = instanceIP
	this.InstanceIdentifier = instanceIdentifier
	this.Version = version
	return &this
}

// NewInstanceAPIInstanceInfoWithDefaults instantiates a new InstanceAPIInstanceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceAPIInstanceInfoWithDefaults() *InstanceAPIInstanceInfo {
	this := InstanceAPIInstanceInfo{}
	return &this
}

// GetBuildHash returns the BuildHash field value
func (o *InstanceAPIInstanceInfo) GetBuildHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildHash
}

// GetBuildHashOk returns a tuple with the BuildHash field value
// and a boolean to check if the value has been set.
func (o *InstanceAPIInstanceInfo) GetBuildHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildHash, true
}

// SetBuildHash sets field value
func (o *InstanceAPIInstanceInfo) SetBuildHash(v string) {
	o.BuildHash = v
}

// GetDirs returns the Dirs field value
func (o *InstanceAPIInstanceInfo) GetDirs() ExtconfigExtConfigDirs {
	if o == nil {
		var ret ExtconfigExtConfigDirs
		return ret
	}

	return o.Dirs
}

// GetDirsOk returns a tuple with the Dirs field value
// and a boolean to check if the value has been set.
func (o *InstanceAPIInstanceInfo) GetDirsOk() (*ExtconfigExtConfigDirs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dirs, true
}

// SetDirs sets field value
func (o *InstanceAPIInstanceInfo) SetDirs(v ExtconfigExtConfigDirs) {
	o.Dirs = v
}

// GetInstanceIP returns the InstanceIP field value
func (o *InstanceAPIInstanceInfo) GetInstanceIP() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceIP
}

// GetInstanceIPOk returns a tuple with the InstanceIP field value
// and a boolean to check if the value has been set.
func (o *InstanceAPIInstanceInfo) GetInstanceIPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceIP, true
}

// SetInstanceIP sets field value
func (o *InstanceAPIInstanceInfo) SetInstanceIP(v string) {
	o.InstanceIP = v
}

// GetInstanceIdentifier returns the InstanceIdentifier field value
func (o *InstanceAPIInstanceInfo) GetInstanceIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceIdentifier
}

// GetInstanceIdentifierOk returns a tuple with the InstanceIdentifier field value
// and a boolean to check if the value has been set.
func (o *InstanceAPIInstanceInfo) GetInstanceIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceIdentifier, true
}

// SetInstanceIdentifier sets field value
func (o *InstanceAPIInstanceInfo) SetInstanceIdentifier(v string) {
	o.InstanceIdentifier = v
}

// GetVersion returns the Version field value
func (o *InstanceAPIInstanceInfo) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *InstanceAPIInstanceInfo) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *InstanceAPIInstanceInfo) SetVersion(v string) {
	o.Version = v
}

func (o InstanceAPIInstanceInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceAPIInstanceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["buildHash"] = o.BuildHash
	toSerialize["dirs"] = o.Dirs
	toSerialize["instanceIP"] = o.InstanceIP
	toSerialize["instanceIdentifier"] = o.InstanceIdentifier
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableInstanceAPIInstanceInfo struct {
	value *InstanceAPIInstanceInfo
	isSet bool
}

func (v NullableInstanceAPIInstanceInfo) Get() *InstanceAPIInstanceInfo {
	return v.value
}

func (v *NullableInstanceAPIInstanceInfo) Set(val *InstanceAPIInstanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceAPIInstanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceAPIInstanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceAPIInstanceInfo(val *InstanceAPIInstanceInfo) *NullableInstanceAPIInstanceInfo {
	return &NullableInstanceAPIInstanceInfo{value: val, isSet: true}
}

func (v NullableInstanceAPIInstanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceAPIInstanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
