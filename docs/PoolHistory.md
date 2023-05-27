# PoolHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveStake** | Pointer to **int64** | Active stake in the epoch | [optional] 
**ActiveStakePct** | Pointer to **string** | Pool active stake as percentage of total active stake | [optional] 
**BlockCnt** | Pointer to **int64** | Blocks created in the epoch | [optional] 
**DelegRewards** | **int64** | Total rewards earned by pool delegators for the epoch | 
**DelegatorCnt** | Pointer to **int64** | Delegators in the epoch | [optional] 
**EpochNo** | **int64** | Epoch number | 
**EpochRos** | **string** | Annual return percentage for delegators for the epoch | 
**FixedCost** | **int64** | Pool fixed cost | 
**Margin** | Pointer to **float32** | Pool margin | [optional] 
**PoolFees** | **int64** | Fees collected for the epoch | 
**SaturationPct** | Pointer to **string** | Pool saturation percent | [optional] 

## Methods

### NewPoolHistory

`func NewPoolHistory(delegRewards int64, epochNo int64, epochRos string, fixedCost int64, poolFees int64, ) *PoolHistory`

NewPoolHistory instantiates a new PoolHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolHistoryWithDefaults

`func NewPoolHistoryWithDefaults() *PoolHistory`

NewPoolHistoryWithDefaults instantiates a new PoolHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveStake

`func (o *PoolHistory) GetActiveStake() int64`

GetActiveStake returns the ActiveStake field if non-nil, zero value otherwise.

### GetActiveStakeOk

`func (o *PoolHistory) GetActiveStakeOk() (*int64, bool)`

GetActiveStakeOk returns a tuple with the ActiveStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStake

`func (o *PoolHistory) SetActiveStake(v int64)`

SetActiveStake sets ActiveStake field to given value.

### HasActiveStake

`func (o *PoolHistory) HasActiveStake() bool`

HasActiveStake returns a boolean if a field has been set.

### GetActiveStakePct

`func (o *PoolHistory) GetActiveStakePct() string`

GetActiveStakePct returns the ActiveStakePct field if non-nil, zero value otherwise.

### GetActiveStakePctOk

`func (o *PoolHistory) GetActiveStakePctOk() (*string, bool)`

GetActiveStakePctOk returns a tuple with the ActiveStakePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStakePct

`func (o *PoolHistory) SetActiveStakePct(v string)`

SetActiveStakePct sets ActiveStakePct field to given value.

### HasActiveStakePct

`func (o *PoolHistory) HasActiveStakePct() bool`

HasActiveStakePct returns a boolean if a field has been set.

### GetBlockCnt

`func (o *PoolHistory) GetBlockCnt() int64`

GetBlockCnt returns the BlockCnt field if non-nil, zero value otherwise.

### GetBlockCntOk

`func (o *PoolHistory) GetBlockCntOk() (*int64, bool)`

GetBlockCntOk returns a tuple with the BlockCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCnt

`func (o *PoolHistory) SetBlockCnt(v int64)`

SetBlockCnt sets BlockCnt field to given value.

### HasBlockCnt

`func (o *PoolHistory) HasBlockCnt() bool`

HasBlockCnt returns a boolean if a field has been set.

### GetDelegRewards

`func (o *PoolHistory) GetDelegRewards() int64`

GetDelegRewards returns the DelegRewards field if non-nil, zero value otherwise.

### GetDelegRewardsOk

`func (o *PoolHistory) GetDelegRewardsOk() (*int64, bool)`

GetDelegRewardsOk returns a tuple with the DelegRewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegRewards

`func (o *PoolHistory) SetDelegRewards(v int64)`

SetDelegRewards sets DelegRewards field to given value.


### GetDelegatorCnt

`func (o *PoolHistory) GetDelegatorCnt() int64`

GetDelegatorCnt returns the DelegatorCnt field if non-nil, zero value otherwise.

### GetDelegatorCntOk

`func (o *PoolHistory) GetDelegatorCntOk() (*int64, bool)`

GetDelegatorCntOk returns a tuple with the DelegatorCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatorCnt

`func (o *PoolHistory) SetDelegatorCnt(v int64)`

SetDelegatorCnt sets DelegatorCnt field to given value.

### HasDelegatorCnt

`func (o *PoolHistory) HasDelegatorCnt() bool`

HasDelegatorCnt returns a boolean if a field has been set.

### GetEpochNo

`func (o *PoolHistory) GetEpochNo() int64`

GetEpochNo returns the EpochNo field if non-nil, zero value otherwise.

### GetEpochNoOk

`func (o *PoolHistory) GetEpochNoOk() (*int64, bool)`

GetEpochNoOk returns a tuple with the EpochNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochNo

`func (o *PoolHistory) SetEpochNo(v int64)`

SetEpochNo sets EpochNo field to given value.


### GetEpochRos

`func (o *PoolHistory) GetEpochRos() string`

GetEpochRos returns the EpochRos field if non-nil, zero value otherwise.

### GetEpochRosOk

`func (o *PoolHistory) GetEpochRosOk() (*string, bool)`

GetEpochRosOk returns a tuple with the EpochRos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochRos

`func (o *PoolHistory) SetEpochRos(v string)`

SetEpochRos sets EpochRos field to given value.


### GetFixedCost

`func (o *PoolHistory) GetFixedCost() int64`

GetFixedCost returns the FixedCost field if non-nil, zero value otherwise.

### GetFixedCostOk

`func (o *PoolHistory) GetFixedCostOk() (*int64, bool)`

GetFixedCostOk returns a tuple with the FixedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedCost

`func (o *PoolHistory) SetFixedCost(v int64)`

SetFixedCost sets FixedCost field to given value.


### GetMargin

`func (o *PoolHistory) GetMargin() float32`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *PoolHistory) GetMarginOk() (*float32, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *PoolHistory) SetMargin(v float32)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *PoolHistory) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetPoolFees

`func (o *PoolHistory) GetPoolFees() int64`

GetPoolFees returns the PoolFees field if non-nil, zero value otherwise.

### GetPoolFeesOk

`func (o *PoolHistory) GetPoolFeesOk() (*int64, bool)`

GetPoolFeesOk returns a tuple with the PoolFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolFees

`func (o *PoolHistory) SetPoolFees(v int64)`

SetPoolFees sets PoolFees field to given value.


### GetSaturationPct

`func (o *PoolHistory) GetSaturationPct() string`

GetSaturationPct returns the SaturationPct field if non-nil, zero value otherwise.

### GetSaturationPctOk

`func (o *PoolHistory) GetSaturationPctOk() (*string, bool)`

GetSaturationPctOk returns a tuple with the SaturationPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturationPct

`func (o *PoolHistory) SetSaturationPct(v string)`

SetSaturationPct sets SaturationPct field to given value.

### HasSaturationPct

`func (o *PoolHistory) HasSaturationPct() bool`

HasSaturationPct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


