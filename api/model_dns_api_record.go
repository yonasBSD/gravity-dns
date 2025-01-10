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

// checks if the DnsAPIRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsAPIRecord{}

// DnsAPIRecord struct for DnsAPIRecord
type DnsAPIRecord struct {
	Data         string `json:"data"`
	Fqdn         string `json:"fqdn"`
	Hostname     string `json:"hostname"`
	MxPreference *int32 `json:"mxPreference,omitempty"`
	SrvPort      *int32 `json:"srvPort,omitempty"`
	SrvPriority  *int32 `json:"srvPriority,omitempty"`
	SrvWeight    *int32 `json:"srvWeight,omitempty"`
	Type         string `json:"type"`
	Uid          string `json:"uid"`
}

// NewDnsAPIRecord instantiates a new DnsAPIRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsAPIRecord(data string, fqdn string, hostname string, type_ string, uid string) *DnsAPIRecord {
	this := DnsAPIRecord{}
	this.Data = data
	this.Fqdn = fqdn
	this.Hostname = hostname
	this.Type = type_
	this.Uid = uid
	return &this
}

// NewDnsAPIRecordWithDefaults instantiates a new DnsAPIRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsAPIRecordWithDefaults() *DnsAPIRecord {
	this := DnsAPIRecord{}
	return &this
}

// GetData returns the Data field value
func (o *DnsAPIRecord) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DnsAPIRecord) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DnsAPIRecord) SetData(v string) {
	o.Data = v
}

// GetFqdn returns the Fqdn field value
func (o *DnsAPIRecord) GetFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *DnsAPIRecord) GetFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *DnsAPIRecord) SetFqdn(v string) {
	o.Fqdn = v
}

// GetHostname returns the Hostname field value
func (o *DnsAPIRecord) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *DnsAPIRecord) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *DnsAPIRecord) SetHostname(v string) {
	o.Hostname = v
}

// GetMxPreference returns the MxPreference field value if set, zero value otherwise.
func (o *DnsAPIRecord) GetMxPreference() int32 {
	if o == nil || IsNil(o.MxPreference) {
		var ret int32
		return ret
	}
	return *o.MxPreference
}

// GetMxPreferenceOk returns a tuple with the MxPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsAPIRecord) GetMxPreferenceOk() (*int32, bool) {
	if o == nil || IsNil(o.MxPreference) {
		return nil, false
	}
	return o.MxPreference, true
}

// HasMxPreference returns a boolean if a field has been set.
func (o *DnsAPIRecord) HasMxPreference() bool {
	if o != nil && !IsNil(o.MxPreference) {
		return true
	}

	return false
}

// SetMxPreference gets a reference to the given int32 and assigns it to the MxPreference field.
func (o *DnsAPIRecord) SetMxPreference(v int32) {
	o.MxPreference = &v
}

// GetSrvPort returns the SrvPort field value if set, zero value otherwise.
func (o *DnsAPIRecord) GetSrvPort() int32 {
	if o == nil || IsNil(o.SrvPort) {
		var ret int32
		return ret
	}
	return *o.SrvPort
}

// GetSrvPortOk returns a tuple with the SrvPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsAPIRecord) GetSrvPortOk() (*int32, bool) {
	if o == nil || IsNil(o.SrvPort) {
		return nil, false
	}
	return o.SrvPort, true
}

// HasSrvPort returns a boolean if a field has been set.
func (o *DnsAPIRecord) HasSrvPort() bool {
	if o != nil && !IsNil(o.SrvPort) {
		return true
	}

	return false
}

// SetSrvPort gets a reference to the given int32 and assigns it to the SrvPort field.
func (o *DnsAPIRecord) SetSrvPort(v int32) {
	o.SrvPort = &v
}

// GetSrvPriority returns the SrvPriority field value if set, zero value otherwise.
func (o *DnsAPIRecord) GetSrvPriority() int32 {
	if o == nil || IsNil(o.SrvPriority) {
		var ret int32
		return ret
	}
	return *o.SrvPriority
}

// GetSrvPriorityOk returns a tuple with the SrvPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsAPIRecord) GetSrvPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.SrvPriority) {
		return nil, false
	}
	return o.SrvPriority, true
}

// HasSrvPriority returns a boolean if a field has been set.
func (o *DnsAPIRecord) HasSrvPriority() bool {
	if o != nil && !IsNil(o.SrvPriority) {
		return true
	}

	return false
}

// SetSrvPriority gets a reference to the given int32 and assigns it to the SrvPriority field.
func (o *DnsAPIRecord) SetSrvPriority(v int32) {
	o.SrvPriority = &v
}

// GetSrvWeight returns the SrvWeight field value if set, zero value otherwise.
func (o *DnsAPIRecord) GetSrvWeight() int32 {
	if o == nil || IsNil(o.SrvWeight) {
		var ret int32
		return ret
	}
	return *o.SrvWeight
}

// GetSrvWeightOk returns a tuple with the SrvWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsAPIRecord) GetSrvWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.SrvWeight) {
		return nil, false
	}
	return o.SrvWeight, true
}

// HasSrvWeight returns a boolean if a field has been set.
func (o *DnsAPIRecord) HasSrvWeight() bool {
	if o != nil && !IsNil(o.SrvWeight) {
		return true
	}

	return false
}

// SetSrvWeight gets a reference to the given int32 and assigns it to the SrvWeight field.
func (o *DnsAPIRecord) SetSrvWeight(v int32) {
	o.SrvWeight = &v
}

// GetType returns the Type field value
func (o *DnsAPIRecord) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DnsAPIRecord) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DnsAPIRecord) SetType(v string) {
	o.Type = v
}

// GetUid returns the Uid field value
func (o *DnsAPIRecord) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *DnsAPIRecord) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *DnsAPIRecord) SetUid(v string) {
	o.Uid = v
}

func (o DnsAPIRecord) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsAPIRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["fqdn"] = o.Fqdn
	toSerialize["hostname"] = o.Hostname
	if !IsNil(o.MxPreference) {
		toSerialize["mxPreference"] = o.MxPreference
	}
	if !IsNil(o.SrvPort) {
		toSerialize["srvPort"] = o.SrvPort
	}
	if !IsNil(o.SrvPriority) {
		toSerialize["srvPriority"] = o.SrvPriority
	}
	if !IsNil(o.SrvWeight) {
		toSerialize["srvWeight"] = o.SrvWeight
	}
	toSerialize["type"] = o.Type
	toSerialize["uid"] = o.Uid
	return toSerialize, nil
}

type NullableDnsAPIRecord struct {
	value *DnsAPIRecord
	isSet bool
}

func (v NullableDnsAPIRecord) Get() *DnsAPIRecord {
	return v.value
}

func (v *NullableDnsAPIRecord) Set(val *DnsAPIRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsAPIRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsAPIRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsAPIRecord(val *DnsAPIRecord) *NullableDnsAPIRecord {
	return &NullableDnsAPIRecord{value: val, isSet: true}
}

func (v NullableDnsAPIRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsAPIRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
