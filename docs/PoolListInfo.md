# PoolListInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolIdBech32** | **string** | Bech32 encoded pool ID | 
**Ticker** | Pointer to **string** | Pool ticker symbol | [optional] 

## Methods

### NewPoolListInfo

`func NewPoolListInfo(poolIdBech32 string, ) *PoolListInfo`

NewPoolListInfo instantiates a new PoolListInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolListInfoWithDefaults

`func NewPoolListInfoWithDefaults() *PoolListInfo`

NewPoolListInfoWithDefaults instantiates a new PoolListInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolIdBech32

`func (o *PoolListInfo) GetPoolIdBech32() string`

GetPoolIdBech32 returns the PoolIdBech32 field if non-nil, zero value otherwise.

### GetPoolIdBech32Ok

`func (o *PoolListInfo) GetPoolIdBech32Ok() (*string, bool)`

GetPoolIdBech32Ok returns a tuple with the PoolIdBech32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIdBech32

`func (o *PoolListInfo) SetPoolIdBech32(v string)`

SetPoolIdBech32 sets PoolIdBech32 field to given value.


### GetTicker

`func (o *PoolListInfo) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *PoolListInfo) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *PoolListInfo) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *PoolListInfo) HasTicker() bool`

HasTicker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


