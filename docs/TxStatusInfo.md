# TxStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxHash** | **string** |  | 
**TxStatus** | [**TxStatus**](TxStatus.md) |  | 

## Methods

### NewTxStatusInfo

`func NewTxStatusInfo(txHash string, txStatus TxStatus, ) *TxStatusInfo`

NewTxStatusInfo instantiates a new TxStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxStatusInfoWithDefaults

`func NewTxStatusInfoWithDefaults() *TxStatusInfo`

NewTxStatusInfoWithDefaults instantiates a new TxStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxHash

`func (o *TxStatusInfo) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *TxStatusInfo) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *TxStatusInfo) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetTxStatus

`func (o *TxStatusInfo) GetTxStatus() TxStatus`

GetTxStatus returns the TxStatus field if non-nil, zero value otherwise.

### GetTxStatusOk

`func (o *TxStatusInfo) GetTxStatusOk() (*TxStatus, bool)`

GetTxStatusOk returns a tuple with the TxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatus

`func (o *TxStatusInfo) SetTxStatus(v TxStatus)`

SetTxStatus sets TxStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


