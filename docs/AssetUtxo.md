# AssetUtxo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Address which controls the UTxO | 
**Amount** | **int64** | Amount of the asset contained in the UTxO | 
**Index** | **int32** | UTxO transaction index | 
**TxHash** | **string** | UTxO transaction hash | 

## Methods

### NewAssetUtxo

`func NewAssetUtxo(address string, amount int64, index int32, txHash string, ) *AssetUtxo`

NewAssetUtxo instantiates a new AssetUtxo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetUtxoWithDefaults

`func NewAssetUtxoWithDefaults() *AssetUtxo`

NewAssetUtxoWithDefaults instantiates a new AssetUtxo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AssetUtxo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AssetUtxo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AssetUtxo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *AssetUtxo) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AssetUtxo) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AssetUtxo) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetIndex

`func (o *AssetUtxo) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AssetUtxo) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AssetUtxo) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetTxHash

`func (o *AssetUtxo) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *AssetUtxo) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *AssetUtxo) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


