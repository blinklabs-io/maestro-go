# PolicyHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Address of the holder | 
**Assets** | [**[]AssetInPolicy**](AssetInPolicy.md) |  | 

## Methods

### NewPolicyHolder

`func NewPolicyHolder(address string, assets []AssetInPolicy, ) *PolicyHolder`

NewPolicyHolder instantiates a new PolicyHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyHolderWithDefaults

`func NewPolicyHolderWithDefaults() *PolicyHolder`

NewPolicyHolderWithDefaults instantiates a new PolicyHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PolicyHolder) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PolicyHolder) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PolicyHolder) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAssets

`func (o *PolicyHolder) GetAssets() []AssetInPolicy`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *PolicyHolder) GetAssetsOk() (*[]AssetInPolicy, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *PolicyHolder) SetAssets(v []AssetInPolicy)`

SetAssets sets Assets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


