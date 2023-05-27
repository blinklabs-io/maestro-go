# Datum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | **string** | Hex encoded datum CBOR bytes | 
**Json** | **map[string]interface{}** |  | 

## Methods

### NewDatum

`func NewDatum(bytes string, json map[string]interface{}, ) *Datum`

NewDatum instantiates a new Datum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatumWithDefaults

`func NewDatumWithDefaults() *Datum`

NewDatumWithDefaults instantiates a new Datum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *Datum) GetBytes() string`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *Datum) GetBytesOk() (*string, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *Datum) SetBytes(v string)`

SetBytes sets Bytes field to given value.


### GetJson

`func (o *Datum) GetJson() map[string]interface{}`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *Datum) GetJsonOk() (*map[string]interface{}, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *Datum) SetJson(v map[string]interface{})`

SetJson sets Json field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


