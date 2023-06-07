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

// BackupRoleConfig struct for BackupRoleConfig
type BackupRoleConfig struct {
	AccessKey *string `json:"accessKey,omitempty"`
	Bucket    *string `json:"bucket,omitempty"`
	CronExpr  *string `json:"cronExpr,omitempty"`
	Endpoint  *string `json:"endpoint,omitempty"`
	Path      *string `json:"path,omitempty"`
	SecretKey *string `json:"secretKey,omitempty"`
}

// NewBackupRoleConfig instantiates a new BackupRoleConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRoleConfig() *BackupRoleConfig {
	this := BackupRoleConfig{}
	return &this
}

// NewBackupRoleConfigWithDefaults instantiates a new BackupRoleConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRoleConfigWithDefaults() *BackupRoleConfig {
	this := BackupRoleConfig{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *BackupRoleConfig) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRoleConfig) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *BackupRoleConfig) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *BackupRoleConfig) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *BackupRoleConfig) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRoleConfig) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *BackupRoleConfig) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *BackupRoleConfig) SetBucket(v string) {
	o.Bucket = &v
}

// GetCronExpr returns the CronExpr field value if set, zero value otherwise.
func (o *BackupRoleConfig) GetCronExpr() string {
	if o == nil || o.CronExpr == nil {
		var ret string
		return ret
	}
	return *o.CronExpr
}

// GetCronExprOk returns a tuple with the CronExpr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRoleConfig) GetCronExprOk() (*string, bool) {
	if o == nil || o.CronExpr == nil {
		return nil, false
	}
	return o.CronExpr, true
}

// HasCronExpr returns a boolean if a field has been set.
func (o *BackupRoleConfig) HasCronExpr() bool {
	if o != nil && o.CronExpr != nil {
		return true
	}

	return false
}

// SetCronExpr gets a reference to the given string and assigns it to the CronExpr field.
func (o *BackupRoleConfig) SetCronExpr(v string) {
	o.CronExpr = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *BackupRoleConfig) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRoleConfig) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *BackupRoleConfig) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *BackupRoleConfig) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BackupRoleConfig) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRoleConfig) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BackupRoleConfig) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *BackupRoleConfig) SetPath(v string) {
	o.Path = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *BackupRoleConfig) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRoleConfig) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *BackupRoleConfig) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *BackupRoleConfig) SetSecretKey(v string) {
	o.SecretKey = &v
}

func (o BackupRoleConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKey != nil {
		toSerialize["accessKey"] = o.AccessKey
	}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	if o.CronExpr != nil {
		toSerialize["cronExpr"] = o.CronExpr
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.SecretKey != nil {
		toSerialize["secretKey"] = o.SecretKey
	}
	return json.Marshal(toSerialize)
}

type NullableBackupRoleConfig struct {
	value *BackupRoleConfig
	isSet bool
}

func (v NullableBackupRoleConfig) Get() *BackupRoleConfig {
	return v.value
}

func (v *NullableBackupRoleConfig) Set(val *BackupRoleConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRoleConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRoleConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRoleConfig(val *BackupRoleConfig) *NullableBackupRoleConfig {
	return &NullableBackupRoleConfig{value: val, isSet: true}
}

func (v NullableBackupRoleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRoleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
