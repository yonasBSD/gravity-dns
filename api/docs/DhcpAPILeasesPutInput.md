# DhcpAPILeasesPutInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**AddressLeaseTime** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DnsZone** | Pointer to **string** |  | [optional] 
**Expiry** | Pointer to **int32** |  | [optional] 
**Hostname** | **string** |  | 

## Methods

### NewDhcpAPILeasesPutInput

`func NewDhcpAPILeasesPutInput(address string, addressLeaseTime string, hostname string, ) *DhcpAPILeasesPutInput`

NewDhcpAPILeasesPutInput instantiates a new DhcpAPILeasesPutInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpAPILeasesPutInputWithDefaults

`func NewDhcpAPILeasesPutInputWithDefaults() *DhcpAPILeasesPutInput`

NewDhcpAPILeasesPutInputWithDefaults instantiates a new DhcpAPILeasesPutInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DhcpAPILeasesPutInput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DhcpAPILeasesPutInput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DhcpAPILeasesPutInput) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressLeaseTime

`func (o *DhcpAPILeasesPutInput) GetAddressLeaseTime() string`

GetAddressLeaseTime returns the AddressLeaseTime field if non-nil, zero value otherwise.

### GetAddressLeaseTimeOk

`func (o *DhcpAPILeasesPutInput) GetAddressLeaseTimeOk() (*string, bool)`

GetAddressLeaseTimeOk returns a tuple with the AddressLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLeaseTime

`func (o *DhcpAPILeasesPutInput) SetAddressLeaseTime(v string)`

SetAddressLeaseTime sets AddressLeaseTime field to given value.


### GetDescription

`func (o *DhcpAPILeasesPutInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DhcpAPILeasesPutInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DhcpAPILeasesPutInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DhcpAPILeasesPutInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDnsZone

`func (o *DhcpAPILeasesPutInput) GetDnsZone() string`

GetDnsZone returns the DnsZone field if non-nil, zero value otherwise.

### GetDnsZoneOk

`func (o *DhcpAPILeasesPutInput) GetDnsZoneOk() (*string, bool)`

GetDnsZoneOk returns a tuple with the DnsZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZone

`func (o *DhcpAPILeasesPutInput) SetDnsZone(v string)`

SetDnsZone sets DnsZone field to given value.

### HasDnsZone

`func (o *DhcpAPILeasesPutInput) HasDnsZone() bool`

HasDnsZone returns a boolean if a field has been set.

### GetExpiry

`func (o *DhcpAPILeasesPutInput) GetExpiry() int32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *DhcpAPILeasesPutInput) GetExpiryOk() (*int32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *DhcpAPILeasesPutInput) SetExpiry(v int32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *DhcpAPILeasesPutInput) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetHostname

`func (o *DhcpAPILeasesPutInput) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DhcpAPILeasesPutInput) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DhcpAPILeasesPutInput) SetHostname(v string)`

SetHostname sets Hostname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


