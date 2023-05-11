# TokenRegistryMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decimals** | **int64** | Recommended value for decimal places | 
**Description** | **string** | Asset description | 
**Logo** | **string** | Base64 encoded logo PNG associated with the asset | 
**Name** | **string** | Asset name | 
**Ticker** | **string** | Asset ticker | 
**Url** | **string** | URL associated with the asset | 

## Methods

### NewTokenRegistryMetadata

`func NewTokenRegistryMetadata(decimals int64, description string, logo string, name string, ticker string, url string, ) *TokenRegistryMetadata`

NewTokenRegistryMetadata instantiates a new TokenRegistryMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRegistryMetadataWithDefaults

`func NewTokenRegistryMetadataWithDefaults() *TokenRegistryMetadata`

NewTokenRegistryMetadataWithDefaults instantiates a new TokenRegistryMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecimals

`func (o *TokenRegistryMetadata) GetDecimals() int64`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenRegistryMetadata) GetDecimalsOk() (*int64, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenRegistryMetadata) SetDecimals(v int64)`

SetDecimals sets Decimals field to given value.


### GetDescription

`func (o *TokenRegistryMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenRegistryMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenRegistryMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLogo

`func (o *TokenRegistryMetadata) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *TokenRegistryMetadata) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *TokenRegistryMetadata) SetLogo(v string)`

SetLogo sets Logo field to given value.


### GetName

`func (o *TokenRegistryMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenRegistryMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenRegistryMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetTicker

`func (o *TokenRegistryMetadata) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *TokenRegistryMetadata) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *TokenRegistryMetadata) SetTicker(v string)`

SetTicker sets Ticker field to given value.


### GetUrl

`func (o *TokenRegistryMetadata) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TokenRegistryMetadata) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TokenRegistryMetadata) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


