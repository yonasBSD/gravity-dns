/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.6.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TypesOIDCConfig struct for TypesOIDCConfig
type TypesOIDCConfig struct {
	ClientID     *string  `json:"clientID,omitempty"`
	ClientSecret *string  `json:"clientSecret,omitempty"`
	Issuer       *string  `json:"issuer,omitempty"`
	RedirectURL  *string  `json:"redirectURL,omitempty"`
	Scopes       []string `json:"scopes,omitempty"`
}

// NewTypesOIDCConfig instantiates a new TypesOIDCConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesOIDCConfig() *TypesOIDCConfig {
	this := TypesOIDCConfig{}
	return &this
}

// NewTypesOIDCConfigWithDefaults instantiates a new TypesOIDCConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesOIDCConfigWithDefaults() *TypesOIDCConfig {
	this := TypesOIDCConfig{}
	return &this
}

// GetClientID returns the ClientID field value if set, zero value otherwise.
func (o *TypesOIDCConfig) GetClientID() string {
	if o == nil || o.ClientID == nil {
		var ret string
		return ret
	}
	return *o.ClientID
}

// GetClientIDOk returns a tuple with the ClientID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesOIDCConfig) GetClientIDOk() (*string, bool) {
	if o == nil || o.ClientID == nil {
		return nil, false
	}
	return o.ClientID, true
}

// HasClientID returns a boolean if a field has been set.
func (o *TypesOIDCConfig) HasClientID() bool {
	if o != nil && o.ClientID != nil {
		return true
	}

	return false
}

// SetClientID gets a reference to the given string and assigns it to the ClientID field.
func (o *TypesOIDCConfig) SetClientID(v string) {
	o.ClientID = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *TypesOIDCConfig) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesOIDCConfig) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *TypesOIDCConfig) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *TypesOIDCConfig) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *TypesOIDCConfig) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesOIDCConfig) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *TypesOIDCConfig) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *TypesOIDCConfig) SetIssuer(v string) {
	o.Issuer = &v
}

// GetRedirectURL returns the RedirectURL field value if set, zero value otherwise.
func (o *TypesOIDCConfig) GetRedirectURL() string {
	if o == nil || o.RedirectURL == nil {
		var ret string
		return ret
	}
	return *o.RedirectURL
}

// GetRedirectURLOk returns a tuple with the RedirectURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesOIDCConfig) GetRedirectURLOk() (*string, bool) {
	if o == nil || o.RedirectURL == nil {
		return nil, false
	}
	return o.RedirectURL, true
}

// HasRedirectURL returns a boolean if a field has been set.
func (o *TypesOIDCConfig) HasRedirectURL() bool {
	if o != nil && o.RedirectURL != nil {
		return true
	}

	return false
}

// SetRedirectURL gets a reference to the given string and assigns it to the RedirectURL field.
func (o *TypesOIDCConfig) SetRedirectURL(v string) {
	o.RedirectURL = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypesOIDCConfig) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypesOIDCConfig) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TypesOIDCConfig) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *TypesOIDCConfig) SetScopes(v []string) {
	o.Scopes = v
}

func (o TypesOIDCConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientID != nil {
		toSerialize["clientID"] = o.ClientID
	}
	if o.ClientSecret != nil {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.RedirectURL != nil {
		toSerialize["redirectURL"] = o.RedirectURL
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	return json.Marshal(toSerialize)
}

type NullableTypesOIDCConfig struct {
	value *TypesOIDCConfig
	isSet bool
}

func (v NullableTypesOIDCConfig) Get() *TypesOIDCConfig {
	return v.value
}

func (v *NullableTypesOIDCConfig) Set(val *TypesOIDCConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesOIDCConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesOIDCConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesOIDCConfig(val *TypesOIDCConfig) *NullableTypesOIDCConfig {
	return &NullableTypesOIDCConfig{value: val, isSet: true}
}

func (v NullableTypesOIDCConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesOIDCConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
