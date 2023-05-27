# AccountUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsSlot** | **int32** | Absolute slot of the block which contained the transaction | 
**Action** | [**AccountAction**](AccountAction.md) |  | 
**EpochNo** | **int32** | Epoch number in which the transaction occured | 
**TxHash** | **string** | Transaction hash of the transaction which performed the action | 

## Methods

### NewAccountUpdate

`func NewAccountUpdate(absSlot int32, action AccountAction, epochNo int32, txHash string, ) *AccountUpdate`

NewAccountUpdate instantiates a new AccountUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdateWithDefaults

`func NewAccountUpdateWithDefaults() *AccountUpdate`

NewAccountUpdateWithDefaults instantiates a new AccountUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsSlot

`func (o *AccountUpdate) GetAbsSlot() int32`

GetAbsSlot returns the AbsSlot field if non-nil, zero value otherwise.

### GetAbsSlotOk

`func (o *AccountUpdate) GetAbsSlotOk() (*int32, bool)`

GetAbsSlotOk returns a tuple with the AbsSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsSlot

`func (o *AccountUpdate) SetAbsSlot(v int32)`

SetAbsSlot sets AbsSlot field to given value.


### GetAction

`func (o *AccountUpdate) GetAction() AccountAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AccountUpdate) GetActionOk() (*AccountAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AccountUpdate) SetAction(v AccountAction)`

SetAction sets Action field to given value.


### GetEpochNo

`func (o *AccountUpdate) GetEpochNo() int32`

GetEpochNo returns the EpochNo field if non-nil, zero value otherwise.

### GetEpochNoOk

`func (o *AccountUpdate) GetEpochNoOk() (*int32, bool)`

GetEpochNoOk returns a tuple with the EpochNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochNo

`func (o *AccountUpdate) SetEpochNo(v int32)`

SetEpochNo sets EpochNo field to given value.


### GetTxHash

`func (o *AccountUpdate) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *AccountUpdate) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *AccountUpdate) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


