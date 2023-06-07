/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.6.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

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
	if o == nil || o.AvgRtt == nil {
		var ret int32
		return ret
	}
	return *o.AvgRtt
}

// GetAvgRttOk returns a tuple with the AvgRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetAvgRttOk() (*int32, bool) {
	if o == nil || o.AvgRtt == nil {
		return nil, false
	}
	return o.AvgRtt, true
}

// HasAvgRtt returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasAvgRtt() bool {
	if o != nil && o.AvgRtt != nil {
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
	if o == nil || o.MaxRtt == nil {
		var ret int32
		return ret
	}
	return *o.MaxRtt
}

// GetMaxRttOk returns a tuple with the MaxRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetMaxRttOk() (*int32, bool) {
	if o == nil || o.MaxRtt == nil {
		return nil, false
	}
	return o.MaxRtt, true
}

// HasMaxRtt returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasMaxRtt() bool {
	if o != nil && o.MaxRtt != nil {
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
	if o == nil || o.MinRtt == nil {
		var ret int32
		return ret
	}
	return *o.MinRtt
}

// GetMinRttOk returns a tuple with the MinRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetMinRttOk() (*int32, bool) {
	if o == nil || o.MinRtt == nil {
		return nil, false
	}
	return o.MinRtt, true
}

// HasMinRtt returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasMinRtt() bool {
	if o != nil && o.MinRtt != nil {
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
	if o == nil || o.PacketLoss == nil {
		var ret float32
		return ret
	}
	return *o.PacketLoss
}

// GetPacketLossOk returns a tuple with the PacketLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetPacketLossOk() (*float32, bool) {
	if o == nil || o.PacketLoss == nil {
		return nil, false
	}
	return o.PacketLoss, true
}

// HasPacketLoss returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasPacketLoss() bool {
	if o != nil && o.PacketLoss != nil {
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
	if o == nil || o.PacketsRecv == nil {
		var ret int32
		return ret
	}
	return *o.PacketsRecv
}

// GetPacketsRecvOk returns a tuple with the PacketsRecv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetPacketsRecvOk() (*int32, bool) {
	if o == nil || o.PacketsRecv == nil {
		return nil, false
	}
	return o.PacketsRecv, true
}

// HasPacketsRecv returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasPacketsRecv() bool {
	if o != nil && o.PacketsRecv != nil {
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
	if o == nil || o.PacketsRecvDuplicates == nil {
		var ret int32
		return ret
	}
	return *o.PacketsRecvDuplicates
}

// GetPacketsRecvDuplicatesOk returns a tuple with the PacketsRecvDuplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetPacketsRecvDuplicatesOk() (*int32, bool) {
	if o == nil || o.PacketsRecvDuplicates == nil {
		return nil, false
	}
	return o.PacketsRecvDuplicates, true
}

// HasPacketsRecvDuplicates returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasPacketsRecvDuplicates() bool {
	if o != nil && o.PacketsRecvDuplicates != nil {
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
	if o == nil || o.PacketsSent == nil {
		var ret int32
		return ret
	}
	return *o.PacketsSent
}

// GetPacketsSentOk returns a tuple with the PacketsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetPacketsSentOk() (*int32, bool) {
	if o == nil || o.PacketsSent == nil {
		return nil, false
	}
	return o.PacketsSent, true
}

// HasPacketsSent returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasPacketsSent() bool {
	if o != nil && o.PacketsSent != nil {
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
	if o == nil || o.StdDevRtt == nil {
		var ret int32
		return ret
	}
	return *o.StdDevRtt
}

// GetStdDevRttOk returns a tuple with the StdDevRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolPingOutput) GetStdDevRttOk() (*int32, bool) {
	if o == nil || o.StdDevRtt == nil {
		return nil, false
	}
	return o.StdDevRtt, true
}

// HasStdDevRtt returns a boolean if a field has been set.
func (o *ApiAPIToolPingOutput) HasStdDevRtt() bool {
	if o != nil && o.StdDevRtt != nil {
		return true
	}

	return false
}

// SetStdDevRtt gets a reference to the given int32 and assigns it to the StdDevRtt field.
func (o *ApiAPIToolPingOutput) SetStdDevRtt(v int32) {
	o.StdDevRtt = &v
}

func (o ApiAPIToolPingOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvgRtt != nil {
		toSerialize["avgRtt"] = o.AvgRtt
	}
	if o.MaxRtt != nil {
		toSerialize["maxRtt"] = o.MaxRtt
	}
	if o.MinRtt != nil {
		toSerialize["minRtt"] = o.MinRtt
	}
	if o.PacketLoss != nil {
		toSerialize["packetLoss"] = o.PacketLoss
	}
	if o.PacketsRecv != nil {
		toSerialize["packetsRecv"] = o.PacketsRecv
	}
	if o.PacketsRecvDuplicates != nil {
		toSerialize["packetsRecvDuplicates"] = o.PacketsRecvDuplicates
	}
	if o.PacketsSent != nil {
		toSerialize["packetsSent"] = o.PacketsSent
	}
	if o.StdDevRtt != nil {
		toSerialize["stdDevRtt"] = o.StdDevRtt
	}
	return json.Marshal(toSerialize)
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
