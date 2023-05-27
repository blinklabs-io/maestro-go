# StakeRegCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertIndex** | **int32** | Index of the certificate in the transaction | 
**StakeAddress** | **string** | Stake address corresponding to stake key being updated | 

## Methods

### NewStakeRegCert

`func NewStakeRegCert(certIndex int32, stakeAddress string, ) *StakeRegCert`

NewStakeRegCert instantiates a new StakeRegCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStakeRegCertWithDefaults

`func NewStakeRegCertWithDefaults() *StakeRegCert`

NewStakeRegCertWithDefaults instantiates a new StakeRegCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertIndex

`func (o *StakeRegCert) GetCertIndex() int32`

GetCertIndex returns the CertIndex field if non-nil, zero value otherwise.

### GetCertIndexOk

`func (o *StakeRegCert) GetCertIndexOk() (*int32, bool)`

GetCertIndexOk returns a tuple with the CertIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIndex

`func (o *StakeRegCert) SetCertIndex(v int32)`

SetCertIndex sets CertIndex field to given value.


### GetStakeAddress

`func (o *StakeRegCert) GetStakeAddress() string`

GetStakeAddress returns the StakeAddress field if non-nil, zero value otherwise.

### GetStakeAddressOk

`func (o *StakeRegCert) GetStakeAddressOk() (*string, bool)`

GetStakeAddressOk returns a tuple with the StakeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeAddress

`func (o *StakeRegCert) SetStakeAddress(v string)`

SetStakeAddress sets StakeAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


