# ChainTip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHash** | **string** | Block hash of the most recent block | 
**Height** | **int64** | Height (number) of the most recent block | 
**Slot** | **int64** | Absolute slot of the most recent block | 

## Methods

### NewChainTip

`func NewChainTip(blockHash string, height int64, slot int64, ) *ChainTip`

NewChainTip instantiates a new ChainTip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainTipWithDefaults

`func NewChainTipWithDefaults() *ChainTip`

NewChainTipWithDefaults instantiates a new ChainTip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHash

`func (o *ChainTip) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *ChainTip) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *ChainTip) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetHeight

`func (o *ChainTip) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ChainTip) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ChainTip) SetHeight(v int64)`

SetHeight sets Height field to given value.


### GetSlot

`func (o *ChainTip) GetSlot() int64`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *ChainTip) GetSlotOk() (*int64, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *ChainTip) SetSlot(v int64)`

SetSlot sets Slot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


