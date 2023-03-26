/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.4.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DiscoveryAPISubnetsPutInput struct for DiscoveryAPISubnetsPutInput
type DiscoveryAPISubnetsPutInput struct {
	DiscoveryTTL int32  `json:"discoveryTTL"`
	DnsResolver  string `json:"dnsResolver"`
	SubnetCidr   string `json:"subnetCidr"`
}

// NewDiscoveryAPISubnetsPutInput instantiates a new DiscoveryAPISubnetsPutInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryAPISubnetsPutInput(discoveryTTL int32, dnsResolver string, subnetCidr string) *DiscoveryAPISubnetsPutInput {
	this := DiscoveryAPISubnetsPutInput{}
	this.DiscoveryTTL = discoveryTTL
	this.DnsResolver = dnsResolver
	this.SubnetCidr = subnetCidr
	return &this
}

// NewDiscoveryAPISubnetsPutInputWithDefaults instantiates a new DiscoveryAPISubnetsPutInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryAPISubnetsPutInputWithDefaults() *DiscoveryAPISubnetsPutInput {
	this := DiscoveryAPISubnetsPutInput{}
	return &this
}

// GetDiscoveryTTL returns the DiscoveryTTL field value
func (o *DiscoveryAPISubnetsPutInput) GetDiscoveryTTL() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DiscoveryTTL
}

// GetDiscoveryTTLOk returns a tuple with the DiscoveryTTL field value
// and a boolean to check if the value has been set.
func (o *DiscoveryAPISubnetsPutInput) GetDiscoveryTTLOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscoveryTTL, true
}

// SetDiscoveryTTL sets field value
func (o *DiscoveryAPISubnetsPutInput) SetDiscoveryTTL(v int32) {
	o.DiscoveryTTL = v
}

// GetDnsResolver returns the DnsResolver field value
func (o *DiscoveryAPISubnetsPutInput) GetDnsResolver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DnsResolver
}

// GetDnsResolverOk returns a tuple with the DnsResolver field value
// and a boolean to check if the value has been set.
func (o *DiscoveryAPISubnetsPutInput) GetDnsResolverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnsResolver, true
}

// SetDnsResolver sets field value
func (o *DiscoveryAPISubnetsPutInput) SetDnsResolver(v string) {
	o.DnsResolver = v
}

// GetSubnetCidr returns the SubnetCidr field value
func (o *DiscoveryAPISubnetsPutInput) GetSubnetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetCidr
}

// GetSubnetCidrOk returns a tuple with the SubnetCidr field value
// and a boolean to check if the value has been set.
func (o *DiscoveryAPISubnetsPutInput) GetSubnetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetCidr, true
}

// SetSubnetCidr sets field value
func (o *DiscoveryAPISubnetsPutInput) SetSubnetCidr(v string) {
	o.SubnetCidr = v
}

func (o DiscoveryAPISubnetsPutInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["discoveryTTL"] = o.DiscoveryTTL
	}
	if true {
		toSerialize["dnsResolver"] = o.DnsResolver
	}
	if true {
		toSerialize["subnetCidr"] = o.SubnetCidr
	}
	return json.Marshal(toSerialize)
}

type NullableDiscoveryAPISubnetsPutInput struct {
	value *DiscoveryAPISubnetsPutInput
	isSet bool
}

func (v NullableDiscoveryAPISubnetsPutInput) Get() *DiscoveryAPISubnetsPutInput {
	return v.value
}

func (v *NullableDiscoveryAPISubnetsPutInput) Set(val *DiscoveryAPISubnetsPutInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryAPISubnetsPutInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryAPISubnetsPutInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryAPISubnetsPutInput(val *DiscoveryAPISubnetsPutInput) *NullableDiscoveryAPISubnetsPutInput {
	return &NullableDiscoveryAPISubnetsPutInput{value: val, isSet: true}
}

func (v NullableDiscoveryAPISubnetsPutInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryAPISubnetsPutInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
