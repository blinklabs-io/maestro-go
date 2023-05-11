# PoolBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsSlot** | Pointer to **int64** | Absolute slot of the block | [optional] 
**BlockHash** | **string** | Block hash | 
**BlockHeight** | **int32** | Block height (block number) | 
**BlockTime** | **int32** | UNIX timestamp when the block was mined | 
**EpochNo** | Pointer to **int32** | Epoch number | [optional] 
**EpochSlot** | Pointer to **int32** | Epoch slot | [optional] 

## Methods

### NewPoolBlock

`func NewPoolBlock(blockHash string, blockHeight int32, blockTime int32, ) *PoolBlock`

NewPoolBlock instantiates a new PoolBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolBlockWithDefaults

`func NewPoolBlockWithDefaults() *PoolBlock`

NewPoolBlockWithDefaults instantiates a new PoolBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsSlot

`func (o *PoolBlock) GetAbsSlot() int64`

GetAbsSlot returns the AbsSlot field if non-nil, zero value otherwise.

### GetAbsSlotOk

`func (o *PoolBlock) GetAbsSlotOk() (*int64, bool)`

GetAbsSlotOk returns a tuple with the AbsSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsSlot

`func (o *PoolBlock) SetAbsSlot(v int64)`

SetAbsSlot sets AbsSlot field to given value.

### HasAbsSlot

`func (o *PoolBlock) HasAbsSlot() bool`

HasAbsSlot returns a boolean if a field has been set.

### GetBlockHash

`func (o *PoolBlock) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *PoolBlock) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *PoolBlock) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockHeight

`func (o *PoolBlock) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *PoolBlock) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *PoolBlock) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetBlockTime

`func (o *PoolBlock) GetBlockTime() int32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *PoolBlock) GetBlockTimeOk() (*int32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *PoolBlock) SetBlockTime(v int32)`

SetBlockTime sets BlockTime field to given value.


### GetEpochNo

`func (o *PoolBlock) GetEpochNo() int32`

GetEpochNo returns the EpochNo field if non-nil, zero value otherwise.

### GetEpochNoOk

`func (o *PoolBlock) GetEpochNoOk() (*int32, bool)`

GetEpochNoOk returns a tuple with the EpochNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochNo

`func (o *PoolBlock) SetEpochNo(v int32)`

SetEpochNo sets EpochNo field to given value.

### HasEpochNo

`func (o *PoolBlock) HasEpochNo() bool`

HasEpochNo returns a boolean if a field has been set.

### GetEpochSlot

`func (o *PoolBlock) GetEpochSlot() int32`

GetEpochSlot returns the EpochSlot field if non-nil, zero value otherwise.

### GetEpochSlotOk

`func (o *PoolBlock) GetEpochSlotOk() (*int32, bool)`

GetEpochSlotOk returns a tuple with the EpochSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochSlot

`func (o *PoolBlock) SetEpochSlot(v int32)`

SetEpochSlot sets EpochSlot field to given value.

### HasEpochSlot

`func (o *PoolBlock) HasEpochSlot() bool`

HasEpochSlot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


