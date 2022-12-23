/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DhcpAPILease struct for DhcpAPILease
type DhcpAPILease struct {
	Address          string  `json:"address"`
	AddressLeaseTime string  `json:"addressLeaseTime"`
	DnsZone          *string `json:"dnsZone,omitempty"`
	Hostname         string  `json:"hostname"`
	Identifier       string  `json:"identifier"`
	ScopeKey         string  `json:"scopeKey"`
}

// NewDhcpAPILease instantiates a new DhcpAPILease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpAPILease(address string, addressLeaseTime string, hostname string, identifier string, scopeKey string) *DhcpAPILease {
	this := DhcpAPILease{}
	this.Address = address
	this.AddressLeaseTime = addressLeaseTime
	this.Hostname = hostname
	this.Identifier = identifier
	this.ScopeKey = scopeKey
	return &this
}

// NewDhcpAPILeaseWithDefaults instantiates a new DhcpAPILease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpAPILeaseWithDefaults() *DhcpAPILease {
	this := DhcpAPILease{}
	return &this
}

// GetAddress returns the Address field value
func (o *DhcpAPILease) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *DhcpAPILease) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *DhcpAPILease) SetAddress(v string) {
	o.Address = v
}

// GetAddressLeaseTime returns the AddressLeaseTime field value
func (o *DhcpAPILease) GetAddressLeaseTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLeaseTime
}

// GetAddressLeaseTimeOk returns a tuple with the AddressLeaseTime field value
// and a boolean to check if the value has been set.
func (o *DhcpAPILease) GetAddressLeaseTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLeaseTime, true
}

// SetAddressLeaseTime sets field value
func (o *DhcpAPILease) SetAddressLeaseTime(v string) {
	o.AddressLeaseTime = v
}

// GetDnsZone returns the DnsZone field value if set, zero value otherwise.
func (o *DhcpAPILease) GetDnsZone() string {
	if o == nil || o.DnsZone == nil {
		var ret string
		return ret
	}
	return *o.DnsZone
}

// GetDnsZoneOk returns a tuple with the DnsZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpAPILease) GetDnsZoneOk() (*string, bool) {
	if o == nil || o.DnsZone == nil {
		return nil, false
	}
	return o.DnsZone, true
}

// HasDnsZone returns a boolean if a field has been set.
func (o *DhcpAPILease) HasDnsZone() bool {
	if o != nil && o.DnsZone != nil {
		return true
	}

	return false
}

// SetDnsZone gets a reference to the given string and assigns it to the DnsZone field.
func (o *DhcpAPILease) SetDnsZone(v string) {
	o.DnsZone = &v
}

// GetHostname returns the Hostname field value
func (o *DhcpAPILease) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *DhcpAPILease) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *DhcpAPILease) SetHostname(v string) {
	o.Hostname = v
}

// GetIdentifier returns the Identifier field value
func (o *DhcpAPILease) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *DhcpAPILease) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *DhcpAPILease) SetIdentifier(v string) {
	o.Identifier = v
}

// GetScopeKey returns the ScopeKey field value
func (o *DhcpAPILease) GetScopeKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScopeKey
}

// GetScopeKeyOk returns a tuple with the ScopeKey field value
// and a boolean to check if the value has been set.
func (o *DhcpAPILease) GetScopeKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScopeKey, true
}

// SetScopeKey sets field value
func (o *DhcpAPILease) SetScopeKey(v string) {
	o.ScopeKey = v
}

func (o DhcpAPILease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["addressLeaseTime"] = o.AddressLeaseTime
	}
	if o.DnsZone != nil {
		toSerialize["dnsZone"] = o.DnsZone
	}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["scopeKey"] = o.ScopeKey
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpAPILease struct {
	value *DhcpAPILease
	isSet bool
}

func (v NullableDhcpAPILease) Get() *DhcpAPILease {
	return v.value
}

func (v *NullableDhcpAPILease) Set(val *DhcpAPILease) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpAPILease) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpAPILease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpAPILease(val *DhcpAPILease) *NullableDhcpAPILease {
	return &NullableDhcpAPILease{value: val, isSet: true}
}

func (v NullableDhcpAPILease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpAPILease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
