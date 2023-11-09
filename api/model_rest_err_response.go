/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.6.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the RestErrResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestErrResponse{}

// RestErrResponse struct for RestErrResponse
type RestErrResponse struct {
	// Application-specific error code.
	Code *int32 `json:"code,omitempty"`
	// Application context.
	Context map[string]interface{} `json:"context,omitempty"`
	// Error message.
	Error *string `json:"error,omitempty"`
	// Status text.
	Status *string `json:"status,omitempty"`
}

// NewRestErrResponse instantiates a new RestErrResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestErrResponse() *RestErrResponse {
	this := RestErrResponse{}
	return &this
}

// NewRestErrResponseWithDefaults instantiates a new RestErrResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestErrResponseWithDefaults() *RestErrResponse {
	this := RestErrResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *RestErrResponse) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestErrResponse) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *RestErrResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *RestErrResponse) SetCode(v int32) {
	o.Code = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *RestErrResponse) GetContext() map[string]interface{} {
	if o == nil || IsNil(o.Context) {
		var ret map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestErrResponse) GetContextOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Context) {
		return map[string]interface{}{}, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *RestErrResponse) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *RestErrResponse) SetContext(v map[string]interface{}) {
	o.Context = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RestErrResponse) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestErrResponse) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RestErrResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *RestErrResponse) SetError(v string) {
	o.Error = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RestErrResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestErrResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RestErrResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RestErrResponse) SetStatus(v string) {
	o.Status = &v
}

func (o RestErrResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestErrResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableRestErrResponse struct {
	value *RestErrResponse
	isSet bool
}

func (v NullableRestErrResponse) Get() *RestErrResponse {
	return v.value
}

func (v *NullableRestErrResponse) Set(val *RestErrResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRestErrResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRestErrResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestErrResponse(val *RestErrResponse) *NullableRestErrResponse {
	return &NullableRestErrResponse{value: val, isSet: true}
}

func (v NullableRestErrResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestErrResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
