# AccountHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveStake** | **int64** | Active stake of the account in the epoch | 
**EpochNo** | **int32** | Epoch number | 
**PoolId** | Pointer to **string** | Bech32 encoded pool ID the account was delegated to | [optional] 

## Methods

### NewAccountHistory

`func NewAccountHistory(activeStake int64, epochNo int32, ) *AccountHistory`

NewAccountHistory instantiates a new AccountHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountHistoryWithDefaults

`func NewAccountHistoryWithDefaults() *AccountHistory`

NewAccountHistoryWithDefaults instantiates a new AccountHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveStake

`func (o *AccountHistory) GetActiveStake() int64`

GetActiveStake returns the ActiveStake field if non-nil, zero value otherwise.

### GetActiveStakeOk

`func (o *AccountHistory) GetActiveStakeOk() (*int64, bool)`

GetActiveStakeOk returns a tuple with the ActiveStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStake

`func (o *AccountHistory) SetActiveStake(v int64)`

SetActiveStake sets ActiveStake field to given value.


### GetEpochNo

`func (o *AccountHistory) GetEpochNo() int32`

GetEpochNo returns the EpochNo field if non-nil, zero value otherwise.

### GetEpochNoOk

`func (o *AccountHistory) GetEpochNoOk() (*int32, bool)`

GetEpochNoOk returns a tuple with the EpochNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochNo

`func (o *AccountHistory) SetEpochNo(v int32)`

SetEpochNo sets EpochNo field to given value.


### GetPoolId

`func (o *AccountHistory) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *AccountHistory) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *AccountHistory) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *AccountHistory) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


