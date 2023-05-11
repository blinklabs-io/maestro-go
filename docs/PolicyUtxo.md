# PolicyUtxo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Address which controls the UTxO | 
**Assets** | [**[]AssetInPolicy**](AssetInPolicy.md) |  | 
**Index** | **int32** | UTxO transaction index | 
**TxHash** | **string** | UTxO transaction hash | 

## Methods

### NewPolicyUtxo

`func NewPolicyUtxo(address string, assets []AssetInPolicy, index int32, txHash string, ) *PolicyUtxo`

NewPolicyUtxo instantiates a new PolicyUtxo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyUtxoWithDefaults

`func NewPolicyUtxoWithDefaults() *PolicyUtxo`

NewPolicyUtxoWithDefaults instantiates a new PolicyUtxo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PolicyUtxo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PolicyUtxo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PolicyUtxo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAssets

`func (o *PolicyUtxo) GetAssets() []AssetInPolicy`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *PolicyUtxo) GetAssetsOk() (*[]AssetInPolicy, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *PolicyUtxo) SetAssets(v []AssetInPolicy)`

SetAssets sets Assets field to given value.


### GetIndex

`func (o *PolicyUtxo) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PolicyUtxo) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PolicyUtxo) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetTxHash

`func (o *PolicyUtxo) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *PolicyUtxo) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *PolicyUtxo) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


