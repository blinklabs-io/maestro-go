# AccountReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Reward amount | 
**EarnedEpoch** | **int32** | Epoch in which the reward was earned | 
**PoolId** | **string** | Bech32 encoded pool ID (if relevant to reward type) | 
**SpendableEpoch** | **int32** | Epoch at which the reward is spendable | 
**Type** | [**AccountStakingRewardType**](AccountStakingRewardType.md) |  | 

## Methods

### NewAccountReward

`func NewAccountReward(amount int64, earnedEpoch int32, poolId string, spendableEpoch int32, type_ AccountStakingRewardType, ) *AccountReward`

NewAccountReward instantiates a new AccountReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRewardWithDefaults

`func NewAccountRewardWithDefaults() *AccountReward`

NewAccountRewardWithDefaults instantiates a new AccountReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AccountReward) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountReward) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountReward) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetEarnedEpoch

`func (o *AccountReward) GetEarnedEpoch() int32`

GetEarnedEpoch returns the EarnedEpoch field if non-nil, zero value otherwise.

### GetEarnedEpochOk

`func (o *AccountReward) GetEarnedEpochOk() (*int32, bool)`

GetEarnedEpochOk returns a tuple with the EarnedEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnedEpoch

`func (o *AccountReward) SetEarnedEpoch(v int32)`

SetEarnedEpoch sets EarnedEpoch field to given value.


### GetPoolId

`func (o *AccountReward) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *AccountReward) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *AccountReward) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetSpendableEpoch

`func (o *AccountReward) GetSpendableEpoch() int32`

GetSpendableEpoch returns the SpendableEpoch field if non-nil, zero value otherwise.

### GetSpendableEpochOk

`func (o *AccountReward) GetSpendableEpochOk() (*int32, bool)`

GetSpendableEpochOk returns a tuple with the SpendableEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendableEpoch

`func (o *AccountReward) SetSpendableEpoch(v int32)`

SetSpendableEpoch sets SpendableEpoch field to given value.


### GetType

`func (o *AccountReward) GetType() AccountStakingRewardType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountReward) GetTypeOk() (*AccountStakingRewardType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountReward) SetType(v AccountStakingRewardType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


