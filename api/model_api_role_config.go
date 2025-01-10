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

// checks if the ApiRoleConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRoleConfig{}

// ApiRoleConfig struct for ApiRoleConfig
type ApiRoleConfig struct {
	CookieSecret    *string          `json:"cookieSecret,omitempty"`
	ListenOverride  *string          `json:"listenOverride,omitempty"`
	Oidc            *TypesOIDCConfig `json:"oidc,omitempty"`
	Port            *int32           `json:"port,omitempty"`
	SessionDuration *string          `json:"sessionDuration,omitempty"`
}

// NewApiRoleConfig instantiates a new ApiRoleConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRoleConfig() *ApiRoleConfig {
	this := ApiRoleConfig{}
	return &this
}

// NewApiRoleConfigWithDefaults instantiates a new ApiRoleConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRoleConfigWithDefaults() *ApiRoleConfig {
	this := ApiRoleConfig{}
	return &this
}

// GetCookieSecret returns the CookieSecret field value if set, zero value otherwise.
func (o *ApiRoleConfig) GetCookieSecret() string {
	if o == nil || IsNil(o.CookieSecret) {
		var ret string
		return ret
	}
	return *o.CookieSecret
}

// GetCookieSecretOk returns a tuple with the CookieSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleConfig) GetCookieSecretOk() (*string, bool) {
	if o == nil || IsNil(o.CookieSecret) {
		return nil, false
	}
	return o.CookieSecret, true
}

// HasCookieSecret returns a boolean if a field has been set.
func (o *ApiRoleConfig) HasCookieSecret() bool {
	if o != nil && !IsNil(o.CookieSecret) {
		return true
	}

	return false
}

// SetCookieSecret gets a reference to the given string and assigns it to the CookieSecret field.
func (o *ApiRoleConfig) SetCookieSecret(v string) {
	o.CookieSecret = &v
}

// GetListenOverride returns the ListenOverride field value if set, zero value otherwise.
func (o *ApiRoleConfig) GetListenOverride() string {
	if o == nil || IsNil(o.ListenOverride) {
		var ret string
		return ret
	}
	return *o.ListenOverride
}

// GetListenOverrideOk returns a tuple with the ListenOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleConfig) GetListenOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.ListenOverride) {
		return nil, false
	}
	return o.ListenOverride, true
}

// HasListenOverride returns a boolean if a field has been set.
func (o *ApiRoleConfig) HasListenOverride() bool {
	if o != nil && !IsNil(o.ListenOverride) {
		return true
	}

	return false
}

// SetListenOverride gets a reference to the given string and assigns it to the ListenOverride field.
func (o *ApiRoleConfig) SetListenOverride(v string) {
	o.ListenOverride = &v
}

// GetOidc returns the Oidc field value if set, zero value otherwise.
func (o *ApiRoleConfig) GetOidc() TypesOIDCConfig {
	if o == nil || IsNil(o.Oidc) {
		var ret TypesOIDCConfig
		return ret
	}
	return *o.Oidc
}

// GetOidcOk returns a tuple with the Oidc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleConfig) GetOidcOk() (*TypesOIDCConfig, bool) {
	if o == nil || IsNil(o.Oidc) {
		return nil, false
	}
	return o.Oidc, true
}

// HasOidc returns a boolean if a field has been set.
func (o *ApiRoleConfig) HasOidc() bool {
	if o != nil && !IsNil(o.Oidc) {
		return true
	}

	return false
}

// SetOidc gets a reference to the given TypesOIDCConfig and assigns it to the Oidc field.
func (o *ApiRoleConfig) SetOidc(v TypesOIDCConfig) {
	o.Oidc = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ApiRoleConfig) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleConfig) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ApiRoleConfig) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ApiRoleConfig) SetPort(v int32) {
	o.Port = &v
}

// GetSessionDuration returns the SessionDuration field value if set, zero value otherwise.
func (o *ApiRoleConfig) GetSessionDuration() string {
	if o == nil || IsNil(o.SessionDuration) {
		var ret string
		return ret
	}
	return *o.SessionDuration
}

// GetSessionDurationOk returns a tuple with the SessionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleConfig) GetSessionDurationOk() (*string, bool) {
	if o == nil || IsNil(o.SessionDuration) {
		return nil, false
	}
	return o.SessionDuration, true
}

// HasSessionDuration returns a boolean if a field has been set.
func (o *ApiRoleConfig) HasSessionDuration() bool {
	if o != nil && !IsNil(o.SessionDuration) {
		return true
	}

	return false
}

// SetSessionDuration gets a reference to the given string and assigns it to the SessionDuration field.
func (o *ApiRoleConfig) SetSessionDuration(v string) {
	o.SessionDuration = &v
}

func (o ApiRoleConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRoleConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CookieSecret) {
		toSerialize["cookieSecret"] = o.CookieSecret
	}
	if !IsNil(o.ListenOverride) {
		toSerialize["listenOverride"] = o.ListenOverride
	}
	if !IsNil(o.Oidc) {
		toSerialize["oidc"] = o.Oidc
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.SessionDuration) {
		toSerialize["sessionDuration"] = o.SessionDuration
	}
	return toSerialize, nil
}

type NullableApiRoleConfig struct {
	value *ApiRoleConfig
	isSet bool
}

func (v NullableApiRoleConfig) Get() *ApiRoleConfig {
	return v.value
}

func (v *NullableApiRoleConfig) Set(val *ApiRoleConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRoleConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRoleConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRoleConfig(val *ApiRoleConfig) *NullableApiRoleConfig {
	return &NullableApiRoleConfig{value: val, isSet: true}
}

func (v NullableApiRoleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRoleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
