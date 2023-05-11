# PoolRetireCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterEpoch** | **int32** | Pool will be retired at the end of this epoch | 
**CertIndex** | **int32** | Index of the certificate in the transaction | 
**PoolId** | **string** | Bech32 pool ID of the pool being retired | 

## Methods

### NewPoolRetireCert

`func NewPoolRetireCert(afterEpoch int32, certIndex int32, poolId string, ) *PoolRetireCert`

NewPoolRetireCert instantiates a new PoolRetireCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolRetireCertWithDefaults

`func NewPoolRetireCertWithDefaults() *PoolRetireCert`

NewPoolRetireCertWithDefaults instantiates a new PoolRetireCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfterEpoch

`func (o *PoolRetireCert) GetAfterEpoch() int32`

GetAfterEpoch returns the AfterEpoch field if non-nil, zero value otherwise.

### GetAfterEpochOk

`func (o *PoolRetireCert) GetAfterEpochOk() (*int32, bool)`

GetAfterEpochOk returns a tuple with the AfterEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterEpoch

`func (o *PoolRetireCert) SetAfterEpoch(v int32)`

SetAfterEpoch sets AfterEpoch field to given value.


### GetCertIndex

`func (o *PoolRetireCert) GetCertIndex() int32`

GetCertIndex returns the CertIndex field if non-nil, zero value otherwise.

### GetCertIndexOk

`func (o *PoolRetireCert) GetCertIndexOk() (*int32, bool)`

GetCertIndexOk returns a tuple with the CertIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIndex

`func (o *PoolRetireCert) SetCertIndex(v int32)`

SetCertIndex sets CertIndex field to given value.


### GetPoolId

`func (o *PoolRetireCert) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *PoolRetireCert) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *PoolRetireCert) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


