# SpendRedeemer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Datum**](Datum.md) |  | 
**ExUnits** | **int64** |  | 
**InputIndex** | **int32** |  | 

## Methods

### NewSpendRedeemer

`func NewSpendRedeemer(data Datum, exUnits int64, inputIndex int32, ) *SpendRedeemer`

NewSpendRedeemer instantiates a new SpendRedeemer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendRedeemerWithDefaults

`func NewSpendRedeemerWithDefaults() *SpendRedeemer`

NewSpendRedeemerWithDefaults instantiates a new SpendRedeemer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SpendRedeemer) GetData() Datum`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SpendRedeemer) GetDataOk() (*Datum, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SpendRedeemer) SetData(v Datum)`

SetData sets Data field to given value.


### GetExUnits

`func (o *SpendRedeemer) GetExUnits() int64`

GetExUnits returns the ExUnits field if non-nil, zero value otherwise.

### GetExUnitsOk

`func (o *SpendRedeemer) GetExUnitsOk() (*int64, bool)`

GetExUnitsOk returns a tuple with the ExUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExUnits

`func (o *SpendRedeemer) SetExUnits(v int64)`

SetExUnits sets ExUnits field to given value.


### GetInputIndex

`func (o *SpendRedeemer) GetInputIndex() int32`

GetInputIndex returns the InputIndex field if non-nil, zero value otherwise.

### GetInputIndexOk

`func (o *SpendRedeemer) GetInputIndexOk() (*int32, bool)`

GetInputIndexOk returns a tuple with the InputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputIndex

`func (o *SpendRedeemer) SetInputIndex(v int32)`

SetInputIndex sets InputIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


