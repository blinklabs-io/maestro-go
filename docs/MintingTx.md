# MintingTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockTimestamp** | **int32** | UNIX timestamp of the block which included transaction | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**MintAmount** | **int64** | Amount of the asset minted or burned (negative if burned) | 
**TxHash** | **string** | Transaction hash | 

## Methods

### NewMintingTx

`func NewMintingTx(blockTimestamp int32, mintAmount int64, txHash string, ) *MintingTx`

NewMintingTx instantiates a new MintingTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintingTxWithDefaults

`func NewMintingTxWithDefaults() *MintingTx`

NewMintingTxWithDefaults instantiates a new MintingTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTimestamp

`func (o *MintingTx) GetBlockTimestamp() int32`

GetBlockTimestamp returns the BlockTimestamp field if non-nil, zero value otherwise.

### GetBlockTimestampOk

`func (o *MintingTx) GetBlockTimestampOk() (*int32, bool)`

GetBlockTimestampOk returns a tuple with the BlockTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimestamp

`func (o *MintingTx) SetBlockTimestamp(v int32)`

SetBlockTimestamp sets BlockTimestamp field to given value.


### GetMetadata

`func (o *MintingTx) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MintingTx) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MintingTx) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MintingTx) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMintAmount

`func (o *MintingTx) GetMintAmount() int64`

GetMintAmount returns the MintAmount field if non-nil, zero value otherwise.

### GetMintAmountOk

`func (o *MintingTx) GetMintAmountOk() (*int64, bool)`

GetMintAmountOk returns a tuple with the MintAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAmount

`func (o *MintingTx) SetMintAmount(v int64)`

SetMintAmount sets MintAmount field to given value.


### GetTxHash

`func (o *MintingTx) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *MintingTx) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *MintingTx) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


