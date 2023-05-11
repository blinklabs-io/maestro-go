# PoolMetaJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Pool description | [optional] 
**Homepage** | Pointer to **string** | Pool home page URL | [optional] 
**Name** | **string** | Pool name | 
**Ticker** | Pointer to **string** | Pool ticker symbol | [optional] 

## Methods

### NewPoolMetaJson

`func NewPoolMetaJson(name string, ) *PoolMetaJson`

NewPoolMetaJson instantiates a new PoolMetaJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolMetaJsonWithDefaults

`func NewPoolMetaJsonWithDefaults() *PoolMetaJson`

NewPoolMetaJsonWithDefaults instantiates a new PoolMetaJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PoolMetaJson) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PoolMetaJson) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PoolMetaJson) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PoolMetaJson) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHomepage

`func (o *PoolMetaJson) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *PoolMetaJson) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *PoolMetaJson) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *PoolMetaJson) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetName

`func (o *PoolMetaJson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolMetaJson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolMetaJson) SetName(v string)`

SetName sets Name field to given value.


### GetTicker

`func (o *PoolMetaJson) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *PoolMetaJson) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *PoolMetaJson) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *PoolMetaJson) HasTicker() bool`

HasTicker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


