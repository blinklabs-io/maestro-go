# Withdrawal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** |  | 
**StakeAddress** | **string** |  | 

## Methods

### NewWithdrawal

`func NewWithdrawal(amount int64, stakeAddress string, ) *Withdrawal`

NewWithdrawal instantiates a new Withdrawal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawalWithDefaults

`func NewWithdrawalWithDefaults() *Withdrawal`

NewWithdrawalWithDefaults instantiates a new Withdrawal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Withdrawal) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Withdrawal) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Withdrawal) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetStakeAddress

`func (o *Withdrawal) GetStakeAddress() string`

GetStakeAddress returns the StakeAddress field if non-nil, zero value otherwise.

### GetStakeAddressOk

`func (o *Withdrawal) GetStakeAddressOk() (*string, bool)`

GetStakeAddressOk returns a tuple with the StakeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeAddress

`func (o *Withdrawal) SetStakeAddress(v string)`

SetStakeAddress sets StakeAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


