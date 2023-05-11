# EraSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to [**Bound**](Bound.md) |  | [optional] 
**Parameters** | [**EraParameters**](EraParameters.md) |  | 
**Start** | [**Bound**](Bound.md) |  | 

## Methods

### NewEraSummary

`func NewEraSummary(parameters EraParameters, start Bound, ) *EraSummary`

NewEraSummary instantiates a new EraSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEraSummaryWithDefaults

`func NewEraSummaryWithDefaults() *EraSummary`

NewEraSummaryWithDefaults instantiates a new EraSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *EraSummary) GetEnd() Bound`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *EraSummary) GetEndOk() (*Bound, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *EraSummary) SetEnd(v Bound)`

SetEnd sets End field to given value.

### HasEnd

`func (o *EraSummary) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetParameters

`func (o *EraSummary) GetParameters() EraParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *EraSummary) GetParametersOk() (*EraParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *EraSummary) SetParameters(v EraParameters)`

SetParameters sets Parameters field to given value.


### GetStart

`func (o *EraSummary) GetStart() Bound`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EraSummary) GetStartOk() (*Bound, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EraSummary) SetStart(v Bound)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


