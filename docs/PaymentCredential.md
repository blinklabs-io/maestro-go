# PaymentCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bech32** | **string** |  | 
**Hex** | **string** |  | 
**Kind** | [**PaymentCredKind**](PaymentCredKind.md) |  | 

## Methods

### NewPaymentCredential

`func NewPaymentCredential(bech32 string, hex string, kind PaymentCredKind, ) *PaymentCredential`

NewPaymentCredential instantiates a new PaymentCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCredentialWithDefaults

`func NewPaymentCredentialWithDefaults() *PaymentCredential`

NewPaymentCredentialWithDefaults instantiates a new PaymentCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBech32

`func (o *PaymentCredential) GetBech32() string`

GetBech32 returns the Bech32 field if non-nil, zero value otherwise.

### GetBech32Ok

`func (o *PaymentCredential) GetBech32Ok() (*string, bool)`

GetBech32Ok returns a tuple with the Bech32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBech32

`func (o *PaymentCredential) SetBech32(v string)`

SetBech32 sets Bech32 field to given value.


### GetHex

`func (o *PaymentCredential) GetHex() string`

GetHex returns the Hex field if non-nil, zero value otherwise.

### GetHexOk

`func (o *PaymentCredential) GetHexOk() (*string, bool)`

GetHexOk returns a tuple with the Hex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHex

`func (o *PaymentCredential) SetHex(v string)`

SetHex sets Hex field to given value.


### GetKind

`func (o *PaymentCredential) GetKind() PaymentCredKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PaymentCredential) GetKindOk() (*PaymentCredKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PaymentCredential) SetKind(v PaymentCredKind)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


