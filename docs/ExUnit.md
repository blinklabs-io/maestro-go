# ExUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memory** | **int64** | Memory execution units | 
**Steps** | **int64** | CPU execution units | 

## Methods

### NewExUnit

`func NewExUnit(memory int64, steps int64, ) *ExUnit`

NewExUnit instantiates a new ExUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExUnitWithDefaults

`func NewExUnitWithDefaults() *ExUnit`

NewExUnitWithDefaults instantiates a new ExUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemory

`func (o *ExUnit) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ExUnit) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ExUnit) SetMemory(v int64)`

SetMemory sets Memory field to given value.


### GetSteps

`func (o *ExUnit) GetSteps() int64`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ExUnit) GetStepsOk() (*int64, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ExUnit) SetSteps(v int64)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


