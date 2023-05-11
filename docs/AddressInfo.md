# AddressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bech32** | Pointer to **string** |  | [optional] 
**Hex** | **string** |  | 
**Network** | Pointer to [**NetworkId**](NetworkId.md) |  | [optional] 
**PaymentCred** | Pointer to [**PaymentCredential**](PaymentCredential.md) |  | [optional] 
**StakingCred** | Pointer to [**StakingCredential**](StakingCredential.md) |  | [optional] 

## Methods

### NewAddressInfo

`func NewAddressInfo(hex string, ) *AddressInfo`

NewAddressInfo instantiates a new AddressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressInfoWithDefaults

`func NewAddressInfoWithDefaults() *AddressInfo`

NewAddressInfoWithDefaults instantiates a new AddressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBech32

`func (o *AddressInfo) GetBech32() string`

GetBech32 returns the Bech32 field if non-nil, zero value otherwise.

### GetBech32Ok

`func (o *AddressInfo) GetBech32Ok() (*string, bool)`

GetBech32Ok returns a tuple with the Bech32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBech32

`func (o *AddressInfo) SetBech32(v string)`

SetBech32 sets Bech32 field to given value.

### HasBech32

`func (o *AddressInfo) HasBech32() bool`

HasBech32 returns a boolean if a field has been set.

### GetHex

`func (o *AddressInfo) GetHex() string`

GetHex returns the Hex field if non-nil, zero value otherwise.

### GetHexOk

`func (o *AddressInfo) GetHexOk() (*string, bool)`

GetHexOk returns a tuple with the Hex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHex

`func (o *AddressInfo) SetHex(v string)`

SetHex sets Hex field to given value.


### GetNetwork

`func (o *AddressInfo) GetNetwork() NetworkId`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AddressInfo) GetNetworkOk() (*NetworkId, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AddressInfo) SetNetwork(v NetworkId)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *AddressInfo) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPaymentCred

`func (o *AddressInfo) GetPaymentCred() PaymentCredential`

GetPaymentCred returns the PaymentCred field if non-nil, zero value otherwise.

### GetPaymentCredOk

`func (o *AddressInfo) GetPaymentCredOk() (*PaymentCredential, bool)`

GetPaymentCredOk returns a tuple with the PaymentCred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCred

`func (o *AddressInfo) SetPaymentCred(v PaymentCredential)`

SetPaymentCred sets PaymentCred field to given value.

### HasPaymentCred

`func (o *AddressInfo) HasPaymentCred() bool`

HasPaymentCred returns a boolean if a field has been set.

### GetStakingCred

`func (o *AddressInfo) GetStakingCred() StakingCredential`

GetStakingCred returns the StakingCred field if non-nil, zero value otherwise.

### GetStakingCredOk

`func (o *AddressInfo) GetStakingCredOk() (*StakingCredential, bool)`

GetStakingCredOk returns a tuple with the StakingCred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingCred

`func (o *AddressInfo) SetStakingCred(v StakingCredential)`

SetStakingCred sets StakingCred field to given value.

### HasStakingCred

`func (o *AddressInfo) HasStakingCred() bool`

HasStakingCred returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


