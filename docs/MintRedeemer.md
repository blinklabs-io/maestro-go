# MintRedeemer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Datum**](Datum.md) |  | 
**ExUnits** | **int64** |  | 
**Policy** | **string** |  | 

## Methods

### NewMintRedeemer

`func NewMintRedeemer(data Datum, exUnits int64, policy string, ) *MintRedeemer`

NewMintRedeemer instantiates a new MintRedeemer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintRedeemerWithDefaults

`func NewMintRedeemerWithDefaults() *MintRedeemer`

NewMintRedeemerWithDefaults instantiates a new MintRedeemer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MintRedeemer) GetData() Datum`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MintRedeemer) GetDataOk() (*Datum, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MintRedeemer) SetData(v Datum)`

SetData sets Data field to given value.


### GetExUnits

`func (o *MintRedeemer) GetExUnits() int64`

GetExUnits returns the ExUnits field if non-nil, zero value otherwise.

### GetExUnitsOk

`func (o *MintRedeemer) GetExUnitsOk() (*int64, bool)`

GetExUnitsOk returns a tuple with the ExUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExUnits

`func (o *MintRedeemer) SetExUnits(v int64)`

SetExUnits sets ExUnits field to given value.


### GetPolicy

`func (o *MintRedeemer) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *MintRedeemer) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *MintRedeemer) SetPolicy(v string)`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


