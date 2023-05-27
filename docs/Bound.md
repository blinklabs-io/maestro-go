# Bound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Epoch** | **int64** |  | 
**Slot** | **int64** |  | 
**Time** | **int64** |  | 

## Methods

### NewBound

`func NewBound(epoch int64, slot int64, time int64, ) *Bound`

NewBound instantiates a new Bound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundWithDefaults

`func NewBoundWithDefaults() *Bound`

NewBoundWithDefaults instantiates a new Bound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpoch

`func (o *Bound) GetEpoch() int64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *Bound) GetEpochOk() (*int64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *Bound) SetEpoch(v int64)`

SetEpoch sets Epoch field to given value.


### GetSlot

`func (o *Bound) GetSlot() int64`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *Bound) GetSlotOk() (*int64, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *Bound) SetSlot(v int64)`

SetSlot sets Slot field to given value.


### GetTime

`func (o *Bound) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Bound) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Bound) SetTime(v int64)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


