# Redeemers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | [**[]CertRedeemer**](CertRedeemer.md) |  | 
**Mints** | [**[]MintRedeemer**](MintRedeemer.md) |  | 
**Spends** | [**[]SpendRedeemer**](SpendRedeemer.md) |  | 
**Withdrawals** | [**[]WdrlRedeemer**](WdrlRedeemer.md) |  | 

## Methods

### NewRedeemers

`func NewRedeemers(certificates []CertRedeemer, mints []MintRedeemer, spends []SpendRedeemer, withdrawals []WdrlRedeemer, ) *Redeemers`

NewRedeemers instantiates a new Redeemers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeemersWithDefaults

`func NewRedeemersWithDefaults() *Redeemers`

NewRedeemersWithDefaults instantiates a new Redeemers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *Redeemers) GetCertificates() []CertRedeemer`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *Redeemers) GetCertificatesOk() (*[]CertRedeemer, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *Redeemers) SetCertificates(v []CertRedeemer)`

SetCertificates sets Certificates field to given value.


### GetMints

`func (o *Redeemers) GetMints() []MintRedeemer`

GetMints returns the Mints field if non-nil, zero value otherwise.

### GetMintsOk

`func (o *Redeemers) GetMintsOk() (*[]MintRedeemer, bool)`

GetMintsOk returns a tuple with the Mints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMints

`func (o *Redeemers) SetMints(v []MintRedeemer)`

SetMints sets Mints field to given value.


### GetSpends

`func (o *Redeemers) GetSpends() []SpendRedeemer`

GetSpends returns the Spends field if non-nil, zero value otherwise.

### GetSpendsOk

`func (o *Redeemers) GetSpendsOk() (*[]SpendRedeemer, bool)`

GetSpendsOk returns a tuple with the Spends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpends

`func (o *Redeemers) SetSpends(v []SpendRedeemer)`

SetSpends sets Spends field to given value.


### GetWithdrawals

`func (o *Redeemers) GetWithdrawals() []WdrlRedeemer`

GetWithdrawals returns the Withdrawals field if non-nil, zero value otherwise.

### GetWithdrawalsOk

`func (o *Redeemers) GetWithdrawalsOk() (*[]WdrlRedeemer, bool)`

GetWithdrawalsOk returns a tuple with the Withdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawals

`func (o *Redeemers) SetWithdrawals(v []WdrlRedeemer)`

SetWithdrawals sets Withdrawals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


