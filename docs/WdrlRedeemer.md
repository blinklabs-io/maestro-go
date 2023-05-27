# WdrlRedeemer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Datum**](Datum.md) |  | 
**ExUnits** | **int64** |  | 
**StakeAddress** | **string** |  | 

## Methods

### NewWdrlRedeemer

`func NewWdrlRedeemer(data Datum, exUnits int64, stakeAddress string, ) *WdrlRedeemer`

NewWdrlRedeemer instantiates a new WdrlRedeemer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWdrlRedeemerWithDefaults

`func NewWdrlRedeemerWithDefaults() *WdrlRedeemer`

NewWdrlRedeemerWithDefaults instantiates a new WdrlRedeemer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WdrlRedeemer) GetData() Datum`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WdrlRedeemer) GetDataOk() (*Datum, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WdrlRedeemer) SetData(v Datum)`

SetData sets Data field to given value.


### GetExUnits

`func (o *WdrlRedeemer) GetExUnits() int64`

GetExUnits returns the ExUnits field if non-nil, zero value otherwise.

### GetExUnitsOk

`func (o *WdrlRedeemer) GetExUnitsOk() (*int64, bool)`

GetExUnitsOk returns a tuple with the ExUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExUnits

`func (o *WdrlRedeemer) SetExUnits(v int64)`

SetExUnits sets ExUnits field to given value.


### GetStakeAddress

`func (o *WdrlRedeemer) GetStakeAddress() string`

GetStakeAddress returns the StakeAddress field if non-nil, zero value otherwise.

### GetStakeAddressOk

`func (o *WdrlRedeemer) GetStakeAddressOk() (*string, bool)`

GetStakeAddressOk returns a tuple with the StakeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeAddress

`func (o *WdrlRedeemer) SetStakeAddress(v string)`

SetStakeAddress sets StakeAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


