# AssetHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Address of the holder | 
**Amount** | **int64** | Amount of the asset owned by the holder | 

## Methods

### NewAssetHolder

`func NewAssetHolder(address string, amount int64, ) *AssetHolder`

NewAssetHolder instantiates a new AssetHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetHolderWithDefaults

`func NewAssetHolderWithDefaults() *AssetHolder`

NewAssetHolderWithDefaults instantiates a new AssetHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AssetHolder) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AssetHolder) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AssetHolder) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *AssetHolder) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AssetHolder) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AssetHolder) SetAmount(v int64)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


