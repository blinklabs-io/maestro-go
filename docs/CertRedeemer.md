# CertRedeemer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertIndex** | **int32** |  | 
**Data** | [**Datum**](Datum.md) |  | 
**ExUnits** | **int64** |  | 

## Methods

### NewCertRedeemer

`func NewCertRedeemer(certIndex int32, data Datum, exUnits int64, ) *CertRedeemer`

NewCertRedeemer instantiates a new CertRedeemer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertRedeemerWithDefaults

`func NewCertRedeemerWithDefaults() *CertRedeemer`

NewCertRedeemerWithDefaults instantiates a new CertRedeemer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertIndex

`func (o *CertRedeemer) GetCertIndex() int32`

GetCertIndex returns the CertIndex field if non-nil, zero value otherwise.

### GetCertIndexOk

`func (o *CertRedeemer) GetCertIndexOk() (*int32, bool)`

GetCertIndexOk returns a tuple with the CertIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIndex

`func (o *CertRedeemer) SetCertIndex(v int32)`

SetCertIndex sets CertIndex field to given value.


### GetData

`func (o *CertRedeemer) GetData() Datum`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CertRedeemer) GetDataOk() (*Datum, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CertRedeemer) SetData(v Datum)`

SetData sets Data field to given value.


### GetExUnits

`func (o *CertRedeemer) GetExUnits() int64`

GetExUnits returns the ExUnits field if non-nil, zero value otherwise.

### GetExUnitsOk

`func (o *CertRedeemer) GetExUnitsOk() (*int64, bool)`

GetExUnitsOk returns a tuple with the ExUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExUnits

`func (o *CertRedeemer) SetExUnits(v int64)`

SetExUnits sets ExUnits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


