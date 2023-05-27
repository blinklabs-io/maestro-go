# DelegatorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveEpochNo** | Pointer to **int64** | Epoch at which the delegation becomes active | [optional] 
**Amount** | Pointer to **int64** | Delegator live stake | [optional] 
**LatestDelegationTxHash** | Pointer to **string** | Transaction hash relating to the most recent delegation | [optional] 
**StakeAddress** | Pointer to **string** | Bech32 encoded stake address (reward address) | [optional] 

## Methods

### NewDelegatorInfo

`func NewDelegatorInfo() *DelegatorInfo`

NewDelegatorInfo instantiates a new DelegatorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegatorInfoWithDefaults

`func NewDelegatorInfoWithDefaults() *DelegatorInfo`

NewDelegatorInfoWithDefaults instantiates a new DelegatorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveEpochNo

`func (o *DelegatorInfo) GetActiveEpochNo() int64`

GetActiveEpochNo returns the ActiveEpochNo field if non-nil, zero value otherwise.

### GetActiveEpochNoOk

`func (o *DelegatorInfo) GetActiveEpochNoOk() (*int64, bool)`

GetActiveEpochNoOk returns a tuple with the ActiveEpochNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveEpochNo

`func (o *DelegatorInfo) SetActiveEpochNo(v int64)`

SetActiveEpochNo sets ActiveEpochNo field to given value.

### HasActiveEpochNo

`func (o *DelegatorInfo) HasActiveEpochNo() bool`

HasActiveEpochNo returns a boolean if a field has been set.

### GetAmount

`func (o *DelegatorInfo) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DelegatorInfo) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DelegatorInfo) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DelegatorInfo) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetLatestDelegationTxHash

`func (o *DelegatorInfo) GetLatestDelegationTxHash() string`

GetLatestDelegationTxHash returns the LatestDelegationTxHash field if non-nil, zero value otherwise.

### GetLatestDelegationTxHashOk

`func (o *DelegatorInfo) GetLatestDelegationTxHashOk() (*string, bool)`

GetLatestDelegationTxHashOk returns a tuple with the LatestDelegationTxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDelegationTxHash

`func (o *DelegatorInfo) SetLatestDelegationTxHash(v string)`

SetLatestDelegationTxHash sets LatestDelegationTxHash field to given value.

### HasLatestDelegationTxHash

`func (o *DelegatorInfo) HasLatestDelegationTxHash() bool`

HasLatestDelegationTxHash returns a boolean if a field has been set.

### GetStakeAddress

`func (o *DelegatorInfo) GetStakeAddress() string`

GetStakeAddress returns the StakeAddress field if non-nil, zero value otherwise.

### GetStakeAddressOk

`func (o *DelegatorInfo) GetStakeAddressOk() (*string, bool)`

GetStakeAddressOk returns a tuple with the StakeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeAddress

`func (o *DelegatorInfo) SetStakeAddress(v string)`

SetStakeAddress sets StakeAddress field to given value.

### HasStakeAddress

`func (o *DelegatorInfo) HasStakeAddress() bool`

HasStakeAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


