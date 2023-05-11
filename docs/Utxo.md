# Utxo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Address which controls the UTxO | 
**Assets** | [**[]Asset**](Asset.md) |  | 
**Datum** | Pointer to [**DatumOption**](DatumOption.md) |  | [optional] 
**Index** | **int32** | UTxO transaction index | 
**ReferenceScript** | Pointer to [**Script**](Script.md) |  | [optional] 
**TxHash** | **string** | UTxO transaction hash | 

## Methods

### NewUtxo

`func NewUtxo(address string, assets []Asset, index int32, txHash string, ) *Utxo`

NewUtxo instantiates a new Utxo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtxoWithDefaults

`func NewUtxoWithDefaults() *Utxo`

NewUtxoWithDefaults instantiates a new Utxo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Utxo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Utxo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Utxo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAssets

`func (o *Utxo) GetAssets() []Asset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Utxo) GetAssetsOk() (*[]Asset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Utxo) SetAssets(v []Asset)`

SetAssets sets Assets field to given value.


### GetDatum

`func (o *Utxo) GetDatum() DatumOption`

GetDatum returns the Datum field if non-nil, zero value otherwise.

### GetDatumOk

`func (o *Utxo) GetDatumOk() (*DatumOption, bool)`

GetDatumOk returns a tuple with the Datum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatum

`func (o *Utxo) SetDatum(v DatumOption)`

SetDatum sets Datum field to given value.

### HasDatum

`func (o *Utxo) HasDatum() bool`

HasDatum returns a boolean if a field has been set.

### GetIndex

`func (o *Utxo) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Utxo) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Utxo) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetReferenceScript

`func (o *Utxo) GetReferenceScript() Script`

GetReferenceScript returns the ReferenceScript field if non-nil, zero value otherwise.

### GetReferenceScriptOk

`func (o *Utxo) GetReferenceScriptOk() (*Script, bool)`

GetReferenceScriptOk returns a tuple with the ReferenceScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceScript

`func (o *Utxo) SetReferenceScript(v Script)`

SetReferenceScript sets ReferenceScript field to given value.

### HasReferenceScript

`func (o *Utxo) HasReferenceScript() bool`

HasReferenceScript returns a boolean if a field has been set.

### GetTxHash

`func (o *Utxo) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *Utxo) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *Utxo) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


