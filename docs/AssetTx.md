# AssetTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHeight** | **int32** | The height of the block which included the transaction | 
**EpochNo** | **int32** | Epoch in which the transaction occurred | 
**TxHash** | **string** | Transaction hash | 

## Methods

### NewAssetTx

`func NewAssetTx(blockHeight int32, epochNo int32, txHash string, ) *AssetTx`

NewAssetTx instantiates a new AssetTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTxWithDefaults

`func NewAssetTxWithDefaults() *AssetTx`

NewAssetTxWithDefaults instantiates a new AssetTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHeight

`func (o *AssetTx) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *AssetTx) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *AssetTx) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetEpochNo

`func (o *AssetTx) GetEpochNo() int32`

GetEpochNo returns the EpochNo field if non-nil, zero value otherwise.

### GetEpochNoOk

`func (o *AssetTx) GetEpochNoOk() (*int32, bool)`

GetEpochNoOk returns a tuple with the EpochNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochNo

`func (o *AssetTx) SetEpochNo(v int32)`

SetEpochNo sets EpochNo field to given value.


### GetTxHash

`func (o *AssetTx) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *AssetTx) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *AssetTx) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


