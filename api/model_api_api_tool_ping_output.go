/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ApiAPIToolPingOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAPIToolPingOutput{}

// ApiAPIToolPingOutput struct for ApiAPIToolPingOutput
type ApiAPIToolPingOutput struct {
	AvgRtt                *int32   `json:"avgRtt,omitempty"`
	MaxRtt                *int32   `json:"maxRtt,omitempty"`
	MinRtt                *int32   `json:"minRtt,omitempty"`
	PacketLoss            *float32 `json:"packetLoss,omitempty"`
	PacketsRecv           *int32   `json:"packetsRecv,omitempty"`
	PacketsRecvDuplicates *int32   `json:"packetsRecvDuplicates,omitempty"`
	PacketsSent           *int32   `json:"packetsSent,omitempty"`
	StdDevRtt             *int32   `json:"stdDevRtt,omitempty"`
}

// NewApiAPIToolPingOutput instantiates a new ApiAPIToolPingOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPIToolPingOutput() *ApiAPIToolPingOutput {
	this := ApiAPIToolPingOutput{}
	return &this
}

// NewApiAPIToolPingOutputWithDefaults instantiates a new ApiAPIToolPingOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPIToolPingOutputWithDefaults() *ApiAPIToolPingOutput {
	this := ApiAPIToolPingOutput{}
	return &this
}

// GetAvgRtt returns the AvgRtt field value if set, zero value otherwise.
func (o *ApiAPIToolPingOutput) GetAvgRtt() int32 {
	if o == nil || IsNil(o.AvgRtt) {
		var ret int32
		return ret
	}
	return *o.AvgRtt
}

// GetAvgRttOk returns a tuple with the AvgRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetAvgRttOk() (*int32, bool) {
	if o == nil || IsNil(o.AvgRtt) {
		return nil, false
	}
	return o.AvgRtt, true
}

// HasAvgRtt returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasAvgRtt() bool {
	if o != nil && !IsNil(o.AvgRtt) {
		return true
	}

	return false
}

// SetAvgRtt gets a reference to the given int32 and assigns it to the AvgRtt field.
func (o *ApiAPIToolPingOutput) SetAvgRtt(v int32) {
	o.AvgRtt = &v
}

// GetMaxRtt returns the MaxRtt field value if set, zero value otherwise.
func (o *ApiAPIToolPingOutput) GetMaxRtt() int32 {
	if o == nil || IsNil(o.MaxRtt) {
		var ret int32
		return ret
	}
	return *o.MaxRtt
}

// GetMaxRttOk returns a tuple with the MaxRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetMaxRttOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRtt) {
		return nil, false
	}
	return o.MaxRtt, true
}

// HasMaxRtt returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasMaxRtt() bool {
	if o != nil && !IsNil(o.MaxRtt) {
		return true
	}

	return false
}

// SetMaxRtt gets a reference to the given int32 and assigns it to the MaxRtt field.
func (o *ApiAPIToolPingOutput) SetMaxRtt(v int32) {
	o.MaxRtt = &v
}

// GetMinRtt returns the MinRtt field value if set, zero value otherwise.
func (o *ApiAPIToolPingOutput) GetMinRtt() int32 {
	if o == nil || IsNil(o.MinRtt) {
		var ret int32
		return ret
	}
	return *o.MinRtt
}

// GetMinRttOk returns a tuple with the MinRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetMinRttOk() (*int32, bool) {
	if o == nil || IsNil(o.MinRtt) {
		return nil, false
	}
	return o.MinRtt, true
}

// HasMinRtt returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasMinRtt() bool {
	if o != nil && !IsNil(o.MinRtt) {
		return true
	}

	return false
}

// SetMinRtt gets a reference to the given int32 and assigns it to the MinRtt field.
func (o *ApiAPIToolPingOutput) SetMinRtt(v int32) {
	o.MinRtt = &v
}

// GetPacketLoss returns the PacketLoss field value if set, zero value otherwise.
func (o *ApiAPIToolPingOutput) GetPacketLoss() float32 {
	if o == nil || IsNil(o.PacketLoss) {
		var ret float32
		return ret
	}
	return *o.PacketLoss
}

// GetPacketLossOk returns a tuple with the PacketLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetPacketLossOk() (*float32, bool) {
	if o == nil || IsNil(o.PacketLoss) {
		return nil, false
	}
	return o.PacketLoss, true
}

// HasPacketLoss returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasPacketLoss() bool {
	if o != nil && !IsNil(o.PacketLoss) {
		return true
	}

	return false
}

// SetPacketLoss gets a reference to the given float32 and assigns it to the PacketLoss field.
func (o *ApiAPIToolPingOutput) SetPacketLoss(v float32) {
	o.PacketLoss = &v
}

// GetPacketsRecv returns the PacketsRecv field value if set, zero value otherwise.
func (o *ApiAPIToolPingOutput) GetPacketsRecv() int32 {
	if o == nil || IsNil(o.PacketsRecv) {
		var ret int32
		return ret
	}
	return *o.PacketsRecv
}

// GetPacketsRecvOk returns a tuple with the PacketsRecv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetPacketsRecvOk() (*int32, bool) {
	if o == nil || IsNil(o.PacketsRecv) {
		return nil, false
	}
	return o.PacketsRecv, true
}

// HasPacketsRecv returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasPacketsRecv() bool {
	if o != nil && !IsNil(o.PacketsRecv) {
		return true
	}

	return false
}

// SetPacketsRecv gets a reference to the given int32 and assigns it to the PacketsRecv field.
func (o *ApiAPIToolPingOutput) SetPacketsRecv(v int32) {
	o.PacketsRecv = &v
}

// GetPacketsRecvDuplicates returns the PacketsRecvDuplicates field value if set, zero value otherwise.
func (o *ApiAPIToolPingOutput) GetPacketsRecvDuplicates() int32 {
	if o == nil || IsNil(o.PacketsRecvDuplicates) {
		var ret int32
		return ret
	}
	return *o.PacketsRecvDuplicates
}

// GetPacketsRecvDuplicatesOk returns a tuple with the PacketsRecvDuplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetPacketsRecvDuplicatesOk() (*int32, bool) {
	if o == nil || IsNil(o.PacketsRecvDuplicates) {
		return nil, false
	}
	return o.PacketsRecvDuplicates, true
}

// HasPacketsRecvDuplicates returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasPacketsRecvDuplicates() bool {
	if o != nil && !IsNil(o.PacketsRecvDuplicates) {
		return true
	}

	return false
}

// SetPacketsRecvDuplicates gets a reference to the given int32 and assigns it to the PacketsRecvDuplicates field.
func (o *ApiAPIToolPingOutput) SetPacketsRecvDuplicates(v int32) {
	o.PacketsRecvDuplicates = &v
}

// GetPacketsSent returns the PacketsSent field value if set, zero value otherwise.
func (o *ApiAPIToolPingOutput) GetPacketsSent() int32 {
	if o == nil || IsNil(o.PacketsSent) {
		var ret int32
		return ret
	}
	return *o.PacketsSent
}

// GetPacketsSentOk returns a tuple with the PacketsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetPacketsSentOk() (*int32, bool) {
	if o == nil || IsNil(o.PacketsSent) {
		return nil, false
	}
	return o.PacketsSent, true
}

// HasPacketsSent returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasPacketsSent() bool {
	if o != nil && !IsNil(o.PacketsSent) {
		return true
	}

	return false
}

// SetPacketsSent gets a reference to the given int32 and assigns it to the PacketsSent field.
func (o *ApiAPIToolPingOutput) SetPacketsSent(v int32) {
	o.PacketsSent = &v
}

// GetStdDevRtt returns the StdDevRtt field value if set, zero value otherwise.
func (o *ApiAPIToolPingOutput) GetStdDevRtt() int32 {
	if o == nil || IsNil(o.StdDevRtt) {
		var ret int32
		return ret
	}
	return *o.StdDevRtt
}

// GetStdDevRttOk returns a tuple with the StdDevRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetStdDevRttOk() (*int32, bool) {
	if o == nil || IsNil(o.StdDevRtt) {
		return nil, false
	}
	return o.StdDevRtt, true
}

// HasStdDevRtt returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasStdDevRtt() bool {
	if o != nil && !IsNil(o.StdDevRtt) {
		return true
	}

	return false
}

// SetStdDevRtt gets a reference to the given int32 and assigns it to the StdDevRtt field.
func (o *ApiAPIToolPingOutput) SetStdDevRtt(v int32) {
	o.StdDevRtt = &v
}

func (o ApiAPIToolPingOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAPIToolPingOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvgRtt) {
		toSerialize["avgRtt"] = o.AvgRtt
	}
	if !IsNil(o.MaxRtt) {
		toSerialize["maxRtt"] = o.MaxRtt
	}
	if !IsNil(o.MinRtt) {
		toSerialize["minRtt"] = o.MinRtt
	}
	if !IsNil(o.PacketLoss) {
		toSerialize["packetLoss"] = o.PacketLoss
	}
	if !IsNil(o.PacketsRecv) {
		toSerialize["packetsRecv"] = o.PacketsRecv
	}
	if !IsNil(o.PacketsRecvDuplicates) {
		toSerialize["packetsRecvDuplicates"] = o.PacketsRecvDuplicates
	}
	if !IsNil(o.PacketsSent) {
		toSerialize["packetsSent"] = o.PacketsSent
	}
	if !IsNil(o.StdDevRtt) {
		toSerialize["stdDevRtt"] = o.StdDevRtt
	}
	return toSerialize, nil
}

type NullableApiAPIToolPingOutput struct {
	value *ApiAPIToolPingOutput
	isSet bool
}

func (v NullableApiAPIToolPingOutput) Get() *ApiAPIToolPingOutput {
	return v.value
}

func (v *NullableApiAPIToolPingOutput) Set(val *ApiAPIToolPingOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPIToolPingOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPIToolPingOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPIToolPingOutput(val *ApiAPIToolPingOutput) *NullableApiAPIToolPingOutput {
	return &NullableApiAPIToolPingOutput{value: val, isSet: true}
}

func (v NullableApiAPIToolPingOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPIToolPingOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
