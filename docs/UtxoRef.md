# UtxoRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | UTxO transaction index | 
**TxHash** | **string** | UTxO transaction hash | 

## Methods

### NewUtxoRef

`func NewUtxoRef(index int32, txHash string, ) *UtxoRef`

NewUtxoRef instantiates a new UtxoRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtxoRefWithDefaults

`func NewUtxoRefWithDefaults() *UtxoRef`

NewUtxoRefWithDefaults instantiates a new UtxoRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *UtxoRef) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *UtxoRef) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *UtxoRef) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetTxHash

`func (o *UtxoRef) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *UtxoRef) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *UtxoRef) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


