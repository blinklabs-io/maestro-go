# MirCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertIndex** | **int32** | Index of the certificate in the transaction | 
**From** | [**MirSource**](MirSource.md) |  | 
**To** | **string** | Where the rewards funds are being sent | 

## Methods

### NewMirCert

`func NewMirCert(certIndex int32, from MirSource, to string, ) *MirCert`

NewMirCert instantiates a new MirCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMirCertWithDefaults

`func NewMirCertWithDefaults() *MirCert`

NewMirCertWithDefaults instantiates a new MirCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertIndex

`func (o *MirCert) GetCertIndex() int32`

GetCertIndex returns the CertIndex field if non-nil, zero value otherwise.

### GetCertIndexOk

`func (o *MirCert) GetCertIndexOk() (*int32, bool)`

GetCertIndexOk returns a tuple with the CertIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIndex

`func (o *MirCert) SetCertIndex(v int32)`

SetCertIndex sets CertIndex field to given value.


### GetFrom

`func (o *MirCert) GetFrom() MirSource`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MirCert) GetFromOk() (*MirSource, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MirCert) SetFrom(v MirSource)`

SetFrom sets From field to given value.


### GetTo

`func (o *MirCert) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MirCert) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MirCert) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


