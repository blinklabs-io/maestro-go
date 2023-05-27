# DatumOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **string** | Hex encoded datum CBOR bytes (&#x60;null&#x60; if datum type is &#x60;hash&#x60; and corresponding datum bytes have not been seen on-chain) | [optional] 
**Hash** | **string** | Datum hash | 
**Json** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | [**DatumOptionType**](DatumOptionType.md) |  | 

## Methods

### NewDatumOption

`func NewDatumOption(hash string, type_ DatumOptionType, ) *DatumOption`

NewDatumOption instantiates a new DatumOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatumOptionWithDefaults

`func NewDatumOptionWithDefaults() *DatumOption`

NewDatumOptionWithDefaults instantiates a new DatumOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *DatumOption) GetBytes() string`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *DatumOption) GetBytesOk() (*string, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *DatumOption) SetBytes(v string)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *DatumOption) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetHash

`func (o *DatumOption) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *DatumOption) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *DatumOption) SetHash(v string)`

SetHash sets Hash field to given value.


### GetJson

`func (o *DatumOption) GetJson() map[string]interface{}`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DatumOption) GetJsonOk() (*map[string]interface{}, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DatumOption) SetJson(v map[string]interface{})`

SetJson sets Json field to given value.

### HasJson

`func (o *DatumOption) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetType

`func (o *DatumOption) GetType() DatumOptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatumOption) GetTypeOk() (*DatumOptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatumOption) SetType(v DatumOptionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


