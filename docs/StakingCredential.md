# StakingCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bech32** | Pointer to **string** |  | [optional] 
**Hex** | Pointer to **string** |  | [optional] 
**Kind** | [**StakingCredKind**](StakingCredKind.md) |  | 
**Pointer** | Pointer to [**Pointer**](Pointer.md) |  | [optional] 
**RewardAddress** | Pointer to **string** |  | [optional] 

## Methods

### NewStakingCredential

`func NewStakingCredential(kind StakingCredKind, ) *StakingCredential`

NewStakingCredential instantiates a new StakingCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStakingCredentialWithDefaults

`func NewStakingCredentialWithDefaults() *StakingCredential`

NewStakingCredentialWithDefaults instantiates a new StakingCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBech32

`func (o *StakingCredential) GetBech32() string`

GetBech32 returns the Bech32 field if non-nil, zero value otherwise.

### GetBech32Ok

`func (o *StakingCredential) GetBech32Ok() (*string, bool)`

GetBech32Ok returns a tuple with the Bech32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBech32

`func (o *StakingCredential) SetBech32(v string)`

SetBech32 sets Bech32 field to given value.

### HasBech32

`func (o *StakingCredential) HasBech32() bool`

HasBech32 returns a boolean if a field has been set.

### GetHex

`func (o *StakingCredential) GetHex() string`

GetHex returns the Hex field if non-nil, zero value otherwise.

### GetHexOk

`func (o *StakingCredential) GetHexOk() (*string, bool)`

GetHexOk returns a tuple with the Hex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHex

`func (o *StakingCredential) SetHex(v string)`

SetHex sets Hex field to given value.

### HasHex

`func (o *StakingCredential) HasHex() bool`

HasHex returns a boolean if a field has been set.

### GetKind

`func (o *StakingCredential) GetKind() StakingCredKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StakingCredential) GetKindOk() (*StakingCredKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StakingCredential) SetKind(v StakingCredKind)`

SetKind sets Kind field to given value.


### GetPointer

`func (o *StakingCredential) GetPointer() Pointer`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *StakingCredential) GetPointerOk() (*Pointer, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *StakingCredential) SetPointer(v Pointer)`

SetPointer sets Pointer field to given value.

### HasPointer

`func (o *StakingCredential) HasPointer() bool`

HasPointer returns a boolean if a field has been set.

### GetRewardAddress

`func (o *StakingCredential) GetRewardAddress() string`

GetRewardAddress returns the RewardAddress field if non-nil, zero value otherwise.

### GetRewardAddressOk

`func (o *StakingCredential) GetRewardAddressOk() (*string, bool)`

GetRewardAddressOk returns a tuple with the RewardAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAddress

`func (o *StakingCredential) SetRewardAddress(v string)`

SetRewardAddress sets RewardAddress field to given value.

### HasRewardAddress

`func (o *StakingCredential) HasRewardAddress() bool`

HasRewardAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


