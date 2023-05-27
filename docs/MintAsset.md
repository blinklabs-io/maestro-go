# MintAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | **int64** | Amount of the asset minted or burned (negative is burn) | 
**Unit** | **string** | Asset (in the form &#x60;hex(policy_id)#hex(asset_name)&#x60;) | 

## Methods

### NewMintAsset

`func NewMintAsset(quantity int64, unit string, ) *MintAsset`

NewMintAsset instantiates a new MintAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintAssetWithDefaults

`func NewMintAssetWithDefaults() *MintAsset`

NewMintAssetWithDefaults instantiates a new MintAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *MintAsset) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MintAsset) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MintAsset) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetUnit

`func (o *MintAsset) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MintAsset) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MintAsset) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


