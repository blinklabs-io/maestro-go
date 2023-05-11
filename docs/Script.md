# Script

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **string** | Script bytes (&#x60;null&#x60; if &#x60;native&#x60; script) | [optional] 
**Hash** | **string** | Script hash | 
**Json** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | [**ScriptType**](ScriptType.md) |  | 

## Methods

### NewScript

`func NewScript(hash string, type_ ScriptType, ) *Script`

NewScript instantiates a new Script object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptWithDefaults

`func NewScriptWithDefaults() *Script`

NewScriptWithDefaults instantiates a new Script object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *Script) GetBytes() string`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *Script) GetBytesOk() (*string, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *Script) SetBytes(v string)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *Script) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetHash

`func (o *Script) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Script) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Script) SetHash(v string)`

SetHash sets Hash field to given value.


### GetJson

`func (o *Script) GetJson() map[string]interface{}`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *Script) GetJsonOk() (*map[string]interface{}, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *Script) SetJson(v map[string]interface{})`

SetJson sets Json field to given value.

### HasJson

`func (o *Script) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetType

`func (o *Script) GetType() ScriptType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Script) GetTypeOk() (*ScriptType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Script) SetType(v ScriptType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


