# EraParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EpochLength** | **int64** |  | 
**SafeZone** | Pointer to **int64** |  | [optional] 
**SlotLength** | **int64** |  | 

## Methods

### NewEraParameters

`func NewEraParameters(epochLength int64, slotLength int64, ) *EraParameters`

NewEraParameters instantiates a new EraParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEraParametersWithDefaults

`func NewEraParametersWithDefaults() *EraParameters`

NewEraParametersWithDefaults instantiates a new EraParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpochLength

`func (o *EraParameters) GetEpochLength() int64`

GetEpochLength returns the EpochLength field if non-nil, zero value otherwise.

### GetEpochLengthOk

`func (o *EraParameters) GetEpochLengthOk() (*int64, bool)`

GetEpochLengthOk returns a tuple with the EpochLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochLength

`func (o *EraParameters) SetEpochLength(v int64)`

SetEpochLength sets EpochLength field to given value.


### GetSafeZone

`func (o *EraParameters) GetSafeZone() int64`

GetSafeZone returns the SafeZone field if non-nil, zero value otherwise.

### GetSafeZoneOk

`func (o *EraParameters) GetSafeZoneOk() (*int64, bool)`

GetSafeZoneOk returns a tuple with the SafeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeZone

`func (o *EraParameters) SetSafeZone(v int64)`

SetSafeZone sets SafeZone field to given value.

### HasSafeZone

`func (o *EraParameters) HasSafeZone() bool`

HasSafeZone returns a boolean if a field has been set.

### GetSlotLength

`func (o *EraParameters) GetSlotLength() int64`

GetSlotLength returns the SlotLength field if non-nil, zero value otherwise.

### GetSlotLengthOk

`func (o *EraParameters) GetSlotLengthOk() (*int64, bool)`

GetSlotLengthOk returns a tuple with the SlotLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotLength

`func (o *EraParameters) SetSlotLength(v int64)`

SetSlotLength sets SlotLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


