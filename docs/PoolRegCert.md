# PoolRegCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertIndex** | **int32** | Index of the certificate in the transaction | 
**FixedCost** | **int64** | Pool fixed cost | 
**FromEpoch** | **int32** | Epoch at which the update will become active | 
**Margin** | **float32** | Pool margin | 
**MetadataHash** | Pointer to **string** | Hash of metadata that the metadata URL should point to | [optional] 
**MetadataUrl** | Pointer to **string** | URL pointing to pool metadata declared by the stake pool | [optional] 
**OwnerAddresses** | **[]string** |  | 
**Pledge** | **int64** | Pool pledge | 
**PoolId** | **string** | Pool ID of the stake pool being updated | 
**Relays** | [**[]Relay**](Relay.md) |  | 
**RewardAddress** | **string** | Stake address which will receive rewards from the stake pool | 
**VrfKeyHash** | **string** | VRF key hash of the stake pool | 

## Methods

### NewPoolRegCert

`func NewPoolRegCert(certIndex int32, fixedCost int64, fromEpoch int32, margin float32, ownerAddresses []string, pledge int64, poolId string, relays []Relay, rewardAddress string, vrfKeyHash string, ) *PoolRegCert`

NewPoolRegCert instantiates a new PoolRegCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolRegCertWithDefaults

`func NewPoolRegCertWithDefaults() *PoolRegCert`

NewPoolRegCertWithDefaults instantiates a new PoolRegCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertIndex

`func (o *PoolRegCert) GetCertIndex() int32`

GetCertIndex returns the CertIndex field if non-nil, zero value otherwise.

### GetCertIndexOk

`func (o *PoolRegCert) GetCertIndexOk() (*int32, bool)`

GetCertIndexOk returns a tuple with the CertIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIndex

`func (o *PoolRegCert) SetCertIndex(v int32)`

SetCertIndex sets CertIndex field to given value.


### GetFixedCost

`func (o *PoolRegCert) GetFixedCost() int64`

GetFixedCost returns the FixedCost field if non-nil, zero value otherwise.

### GetFixedCostOk

`func (o *PoolRegCert) GetFixedCostOk() (*int64, bool)`

GetFixedCostOk returns a tuple with the FixedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedCost

`func (o *PoolRegCert) SetFixedCost(v int64)`

SetFixedCost sets FixedCost field to given value.


### GetFromEpoch

`func (o *PoolRegCert) GetFromEpoch() int32`

GetFromEpoch returns the FromEpoch field if non-nil, zero value otherwise.

### GetFromEpochOk

`func (o *PoolRegCert) GetFromEpochOk() (*int32, bool)`

GetFromEpochOk returns a tuple with the FromEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEpoch

`func (o *PoolRegCert) SetFromEpoch(v int32)`

SetFromEpoch sets FromEpoch field to given value.


### GetMargin

`func (o *PoolRegCert) GetMargin() float32`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *PoolRegCert) GetMarginOk() (*float32, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *PoolRegCert) SetMargin(v float32)`

SetMargin sets Margin field to given value.


### GetMetadataHash

`func (o *PoolRegCert) GetMetadataHash() string`

GetMetadataHash returns the MetadataHash field if non-nil, zero value otherwise.

### GetMetadataHashOk

`func (o *PoolRegCert) GetMetadataHashOk() (*string, bool)`

GetMetadataHashOk returns a tuple with the MetadataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataHash

`func (o *PoolRegCert) SetMetadataHash(v string)`

SetMetadataHash sets MetadataHash field to given value.

### HasMetadataHash

`func (o *PoolRegCert) HasMetadataHash() bool`

HasMetadataHash returns a boolean if a field has been set.

### GetMetadataUrl

`func (o *PoolRegCert) GetMetadataUrl() string`

GetMetadataUrl returns the MetadataUrl field if non-nil, zero value otherwise.

### GetMetadataUrlOk

`func (o *PoolRegCert) GetMetadataUrlOk() (*string, bool)`

GetMetadataUrlOk returns a tuple with the MetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUrl

`func (o *PoolRegCert) SetMetadataUrl(v string)`

SetMetadataUrl sets MetadataUrl field to given value.

### HasMetadataUrl

`func (o *PoolRegCert) HasMetadataUrl() bool`

HasMetadataUrl returns a boolean if a field has been set.

### GetOwnerAddresses

`func (o *PoolRegCert) GetOwnerAddresses() []string`

GetOwnerAddresses returns the OwnerAddresses field if non-nil, zero value otherwise.

### GetOwnerAddressesOk

`func (o *PoolRegCert) GetOwnerAddressesOk() (*[]string, bool)`

GetOwnerAddressesOk returns a tuple with the OwnerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAddresses

`func (o *PoolRegCert) SetOwnerAddresses(v []string)`

SetOwnerAddresses sets OwnerAddresses field to given value.


### GetPledge

`func (o *PoolRegCert) GetPledge() int64`

GetPledge returns the Pledge field if non-nil, zero value otherwise.

### GetPledgeOk

`func (o *PoolRegCert) GetPledgeOk() (*int64, bool)`

GetPledgeOk returns a tuple with the Pledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPledge

`func (o *PoolRegCert) SetPledge(v int64)`

SetPledge sets Pledge field to given value.


### GetPoolId

`func (o *PoolRegCert) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *PoolRegCert) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *PoolRegCert) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetRelays

`func (o *PoolRegCert) GetRelays() []Relay`

GetRelays returns the Relays field if non-nil, zero value otherwise.

### GetRelaysOk

`func (o *PoolRegCert) GetRelaysOk() (*[]Relay, bool)`

GetRelaysOk returns a tuple with the Relays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelays

`func (o *PoolRegCert) SetRelays(v []Relay)`

SetRelays sets Relays field to given value.


### GetRewardAddress

`func (o *PoolRegCert) GetRewardAddress() string`

GetRewardAddress returns the RewardAddress field if non-nil, zero value otherwise.

### GetRewardAddressOk

`func (o *PoolRegCert) GetRewardAddressOk() (*string, bool)`

GetRewardAddressOk returns a tuple with the RewardAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAddress

`func (o *PoolRegCert) SetRewardAddress(v string)`

SetRewardAddress sets RewardAddress field to given value.


### GetVrfKeyHash

`func (o *PoolRegCert) GetVrfKeyHash() string`

GetVrfKeyHash returns the VrfKeyHash field if non-nil, zero value otherwise.

### GetVrfKeyHashOk

`func (o *PoolRegCert) GetVrfKeyHashOk() (*string, bool)`

GetVrfKeyHashOk returns a tuple with the VrfKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfKeyHash

`func (o *PoolRegCert) SetVrfKeyHash(v string)`

SetVrfKeyHash sets VrfKeyHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


