# PoolMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaHash** | Pointer to **string** | Hash of the pool metadata | [optional] 
**MetaJson** | Pointer to [**PoolMetaJson**](PoolMetaJson.md) |  | [optional] 
**MetaUrl** | Pointer to **string** | URL pointing to the pool metadata | [optional] 
**PoolIdBech32** | **string** | Bech32 encoded pool ID | 

## Methods

### NewPoolMetadata

`func NewPoolMetadata(poolIdBech32 string, ) *PoolMetadata`

NewPoolMetadata instantiates a new PoolMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolMetadataWithDefaults

`func NewPoolMetadataWithDefaults() *PoolMetadata`

NewPoolMetadataWithDefaults instantiates a new PoolMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaHash

`func (o *PoolMetadata) GetMetaHash() string`

GetMetaHash returns the MetaHash field if non-nil, zero value otherwise.

### GetMetaHashOk

`func (o *PoolMetadata) GetMetaHashOk() (*string, bool)`

GetMetaHashOk returns a tuple with the MetaHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaHash

`func (o *PoolMetadata) SetMetaHash(v string)`

SetMetaHash sets MetaHash field to given value.

### HasMetaHash

`func (o *PoolMetadata) HasMetaHash() bool`

HasMetaHash returns a boolean if a field has been set.

### GetMetaJson

`func (o *PoolMetadata) GetMetaJson() PoolMetaJson`

GetMetaJson returns the MetaJson field if non-nil, zero value otherwise.

### GetMetaJsonOk

`func (o *PoolMetadata) GetMetaJsonOk() (*PoolMetaJson, bool)`

GetMetaJsonOk returns a tuple with the MetaJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaJson

`func (o *PoolMetadata) SetMetaJson(v PoolMetaJson)`

SetMetaJson sets MetaJson field to given value.

### HasMetaJson

`func (o *PoolMetadata) HasMetaJson() bool`

HasMetaJson returns a boolean if a field has been set.

### GetMetaUrl

`func (o *PoolMetadata) GetMetaUrl() string`

GetMetaUrl returns the MetaUrl field if non-nil, zero value otherwise.

### GetMetaUrlOk

`func (o *PoolMetadata) GetMetaUrlOk() (*string, bool)`

GetMetaUrlOk returns a tuple with the MetaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaUrl

`func (o *PoolMetadata) SetMetaUrl(v string)`

SetMetaUrl sets MetaUrl field to given value.

### HasMetaUrl

`func (o *PoolMetadata) HasMetaUrl() bool`

HasMetaUrl returns a boolean if a field has been set.

### GetPoolIdBech32

`func (o *PoolMetadata) GetPoolIdBech32() string`

GetPoolIdBech32 returns the PoolIdBech32 field if non-nil, zero value otherwise.

### GetPoolIdBech32Ok

`func (o *PoolMetadata) GetPoolIdBech32Ok() (*string, bool)`

GetPoolIdBech32Ok returns a tuple with the PoolIdBech32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIdBech32

`func (o *PoolMetadata) SetPoolIdBech32(v string)`

SetPoolIdBech32 sets PoolIdBech32 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


