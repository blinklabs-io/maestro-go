# AccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegatedPool** | Pointer to **string** | Bech32 pool ID that the stake key is delegated to | [optional] 
**Registered** | **bool** | True if the stake key is registered | 
**RewardsAvailable** | **int64** | The amount of rewards that are available to be withdrawn | 
**StakeAddress** | **string** | Bech32 encoded stake address | 
**TotalBalance** | **int64** | Total balance controlled by the stake key (sum of UTxO and rewards) | 
**TotalRewarded** | **int64** | Total rewards earned | 
**TotalWithdrawn** | **int64** | Total rewards withdrawn | 
**UtxoBalance** | **int64** | Amount locked in UTxOs controlled by addresses with the stake key | 

## Methods

### NewAccountInfo

`func NewAccountInfo(registered bool, rewardsAvailable int64, stakeAddress string, totalBalance int64, totalRewarded int64, totalWithdrawn int64, utxoBalance int64, ) *AccountInfo`

NewAccountInfo instantiates a new AccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInfoWithDefaults

`func NewAccountInfoWithDefaults() *AccountInfo`

NewAccountInfoWithDefaults instantiates a new AccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegatedPool

`func (o *AccountInfo) GetDelegatedPool() string`

GetDelegatedPool returns the DelegatedPool field if non-nil, zero value otherwise.

### GetDelegatedPoolOk

`func (o *AccountInfo) GetDelegatedPoolOk() (*string, bool)`

GetDelegatedPoolOk returns a tuple with the DelegatedPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedPool

`func (o *AccountInfo) SetDelegatedPool(v string)`

SetDelegatedPool sets DelegatedPool field to given value.

### HasDelegatedPool

`func (o *AccountInfo) HasDelegatedPool() bool`

HasDelegatedPool returns a boolean if a field has been set.

### GetRegistered

`func (o *AccountInfo) GetRegistered() bool`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *AccountInfo) GetRegisteredOk() (*bool, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *AccountInfo) SetRegistered(v bool)`

SetRegistered sets Registered field to given value.


### GetRewardsAvailable

`func (o *AccountInfo) GetRewardsAvailable() int64`

GetRewardsAvailable returns the RewardsAvailable field if non-nil, zero value otherwise.

### GetRewardsAvailableOk

`func (o *AccountInfo) GetRewardsAvailableOk() (*int64, bool)`

GetRewardsAvailableOk returns a tuple with the RewardsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsAvailable

`func (o *AccountInfo) SetRewardsAvailable(v int64)`

SetRewardsAvailable sets RewardsAvailable field to given value.


### GetStakeAddress

`func (o *AccountInfo) GetStakeAddress() string`

GetStakeAddress returns the StakeAddress field if non-nil, zero value otherwise.

### GetStakeAddressOk

`func (o *AccountInfo) GetStakeAddressOk() (*string, bool)`

GetStakeAddressOk returns a tuple with the StakeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeAddress

`func (o *AccountInfo) SetStakeAddress(v string)`

SetStakeAddress sets StakeAddress field to given value.


### GetTotalBalance

`func (o *AccountInfo) GetTotalBalance() int64`

GetTotalBalance returns the TotalBalance field if non-nil, zero value otherwise.

### GetTotalBalanceOk

`func (o *AccountInfo) GetTotalBalanceOk() (*int64, bool)`

GetTotalBalanceOk returns a tuple with the TotalBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalance

`func (o *AccountInfo) SetTotalBalance(v int64)`

SetTotalBalance sets TotalBalance field to given value.


### GetTotalRewarded

`func (o *AccountInfo) GetTotalRewarded() int64`

GetTotalRewarded returns the TotalRewarded field if non-nil, zero value otherwise.

### GetTotalRewardedOk

`func (o *AccountInfo) GetTotalRewardedOk() (*int64, bool)`

GetTotalRewardedOk returns a tuple with the TotalRewarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRewarded

`func (o *AccountInfo) SetTotalRewarded(v int64)`

SetTotalRewarded sets TotalRewarded field to given value.


### GetTotalWithdrawn

`func (o *AccountInfo) GetTotalWithdrawn() int64`

GetTotalWithdrawn returns the TotalWithdrawn field if non-nil, zero value otherwise.

### GetTotalWithdrawnOk

`func (o *AccountInfo) GetTotalWithdrawnOk() (*int64, bool)`

GetTotalWithdrawnOk returns a tuple with the TotalWithdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithdrawn

`func (o *AccountInfo) SetTotalWithdrawn(v int64)`

SetTotalWithdrawn sets TotalWithdrawn field to given value.


### GetUtxoBalance

`func (o *AccountInfo) GetUtxoBalance() int64`

GetUtxoBalance returns the UtxoBalance field if non-nil, zero value otherwise.

### GetUtxoBalanceOk

`func (o *AccountInfo) GetUtxoBalanceOk() (*int64, bool)`

GetUtxoBalanceOk returns a tuple with the UtxoBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoBalance

`func (o *AccountInfo) SetUtxoBalance(v int64)`

SetUtxoBalance sets UtxoBalance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


