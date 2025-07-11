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

// checks if the DiscoveryAPIDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveryAPIDevice{}

// DiscoveryAPIDevice struct for DiscoveryAPIDevice
type DiscoveryAPIDevice struct {
	Hostname   string `json:"hostname"`
	Identifier string `json:"identifier"`
	Ip         string `json:"ip"`
	Mac        string `json:"mac"`
}

// NewDiscoveryAPIDevice instantiates a new DiscoveryAPIDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryAPIDevice(hostname string, identifier string, ip string, mac string) *DiscoveryAPIDevice {
	this := DiscoveryAPIDevice{}
	this.Hostname = hostname
	this.Identifier = identifier
	this.Ip = ip
	this.Mac = mac
	return &this
}

// NewDiscoveryAPIDeviceWithDefaults instantiates a new DiscoveryAPIDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryAPIDeviceWithDefaults() *DiscoveryAPIDevice {
	this := DiscoveryAPIDevice{}
	return &this
}

// GetHostname returns the Hostname field value
func (o *DiscoveryAPIDevice) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *DiscoveryAPIDevice) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *DiscoveryAPIDevice) SetHostname(v string) {
	o.Hostname = v
}

// GetIdentifier returns the Identifier field value
func (o *DiscoveryAPIDevice) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *DiscoveryAPIDevice) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *DiscoveryAPIDevice) SetIdentifier(v string) {
	o.Identifier = v
}

// GetIp returns the Ip field value
func (o *DiscoveryAPIDevice) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *DiscoveryAPIDevice) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *DiscoveryAPIDevice) SetIp(v string) {
	o.Ip = v
}

// GetMac returns the Mac field value
func (o *DiscoveryAPIDevice) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *DiscoveryAPIDevice) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *DiscoveryAPIDevice) SetMac(v string) {
	o.Mac = v
}

func (o DiscoveryAPIDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryAPIDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hostname"] = o.Hostname
	toSerialize["identifier"] = o.Identifier
	toSerialize["ip"] = o.Ip
	toSerialize["mac"] = o.Mac
	return toSerialize, nil
}

type NullableDiscoveryAPIDevice struct {
	value *DiscoveryAPIDevice
	isSet bool
}

func (v NullableDiscoveryAPIDevice) Get() *DiscoveryAPIDevice {
	return v.value
}

func (v *NullableDiscoveryAPIDevice) Set(val *DiscoveryAPIDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryAPIDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryAPIDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryAPIDevice(val *DiscoveryAPIDevice) *NullableDiscoveryAPIDevice {
	return &NullableDiscoveryAPIDevice{value: val, isSet: true}
}

func (v NullableDiscoveryAPIDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryAPIDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
