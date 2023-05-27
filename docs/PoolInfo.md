# PoolInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveEpochNo** | **int64** | Epoch when the update takes effect | 
**ActiveStake** | Pointer to **int64** | Active stake | [optional] 
**BlockCount** | Pointer to **int64** | Number of blocks created | [optional] 
**FixedCost** | **int64** | Pool fixed cost | 
**LiveDelegators** | **int64** | Number of current delegators | 
**LivePledge** | Pointer to **int64** | Account balance of pool owners | [optional] 
**LiveSaturation** | Pointer to **string** | Live saturation | [optional] 
**LiveStake** | Pointer to **int64** | Live stake | [optional] 
**Margin** | **float32** | Pool margin | 
**MetaHash** | Pointer to **string** | Hash of the pool metadata | [optional] 
**MetaJson** | Pointer to [**PoolMetaJson**](PoolMetaJson.md) |  | [optional] 
**MetaUrl** | Pointer to **string** | URL pointing to the pool metadata | [optional] 
**OpCert** | Pointer to **string** | Pool operational certificate | [optional] 
**OpCertCounter** | Pointer to **int64** | Operational certificate counter | [optional] 
**Owners** | **[]string** |  | 
**Pledge** | **int64** | Pool pledge | 
**PoolIdBech32** | **string** | Bech32 encoded pool ID | 
**PoolIdHex** | **string** | Hex encoded pool ID | 
**PoolStatus** | Pointer to **string** | Status of the pool | [optional] 
**Relays** | [**[]Relay**](Relay.md) |  | 
**RetiringEpoch** | Pointer to **int32** | Epoch at which the pool will be retired | [optional] 
**RewardAddr** | Pointer to **string** | Reward address associated with the pool | [optional] 
**Sigma** | Pointer to **string** | Pool stake share | [optional] 
**VrfKeyHash** | **string** | VRF key hash | 

## Methods

### NewPoolInfo

`func NewPoolInfo(activeEpochNo int64, fixedCost int64, liveDelegators int64, margin float32, owners []string, pledge int64, poolIdBech32 string, poolIdHex string, relays []Relay, vrfKeyHash string, ) *PoolInfo`

NewPoolInfo instantiates a new PoolInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolInfoWithDefaults

`func NewPoolInfoWithDefaults() *PoolInfo`

NewPoolInfoWithDefaults instantiates a new PoolInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveEpochNo

`func (o *PoolInfo) GetActiveEpochNo() int64`

GetActiveEpochNo returns the ActiveEpochNo field if non-nil, zero value otherwise.

### GetActiveEpochNoOk

`func (o *PoolInfo) GetActiveEpochNoOk() (*int64, bool)`

GetActiveEpochNoOk returns a tuple with the ActiveEpochNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveEpochNo

`func (o *PoolInfo) SetActiveEpochNo(v int64)`

SetActiveEpochNo sets ActiveEpochNo field to given value.


### GetActiveStake

`func (o *PoolInfo) GetActiveStake() int64`

GetActiveStake returns the ActiveStake field if non-nil, zero value otherwise.

### GetActiveStakeOk

`func (o *PoolInfo) GetActiveStakeOk() (*int64, bool)`

GetActiveStakeOk returns a tuple with the ActiveStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStake

`func (o *PoolInfo) SetActiveStake(v int64)`

SetActiveStake sets ActiveStake field to given value.

### HasActiveStake

`func (o *PoolInfo) HasActiveStake() bool`

HasActiveStake returns a boolean if a field has been set.

### GetBlockCount

`func (o *PoolInfo) GetBlockCount() int64`

GetBlockCount returns the BlockCount field if non-nil, zero value otherwise.

### GetBlockCountOk

`func (o *PoolInfo) GetBlockCountOk() (*int64, bool)`

GetBlockCountOk returns a tuple with the BlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCount

`func (o *PoolInfo) SetBlockCount(v int64)`

SetBlockCount sets BlockCount field to given value.

### HasBlockCount

`func (o *PoolInfo) HasBlockCount() bool`

HasBlockCount returns a boolean if a field has been set.

### GetFixedCost

`func (o *PoolInfo) GetFixedCost() int64`

GetFixedCost returns the FixedCost field if non-nil, zero value otherwise.

### GetFixedCostOk

`func (o *PoolInfo) GetFixedCostOk() (*int64, bool)`

GetFixedCostOk returns a tuple with the FixedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedCost

`func (o *PoolInfo) SetFixedCost(v int64)`

SetFixedCost sets FixedCost field to given value.


### GetLiveDelegators

`func (o *PoolInfo) GetLiveDelegators() int64`

GetLiveDelegators returns the LiveDelegators field if non-nil, zero value otherwise.

### GetLiveDelegatorsOk

`func (o *PoolInfo) GetLiveDelegatorsOk() (*int64, bool)`

GetLiveDelegatorsOk returns a tuple with the LiveDelegators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveDelegators

`func (o *PoolInfo) SetLiveDelegators(v int64)`

SetLiveDelegators sets LiveDelegators field to given value.


### GetLivePledge

`func (o *PoolInfo) GetLivePledge() int64`

GetLivePledge returns the LivePledge field if non-nil, zero value otherwise.

### GetLivePledgeOk

`func (o *PoolInfo) GetLivePledgeOk() (*int64, bool)`

GetLivePledgeOk returns a tuple with the LivePledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivePledge

`func (o *PoolInfo) SetLivePledge(v int64)`

SetLivePledge sets LivePledge field to given value.

### HasLivePledge

`func (o *PoolInfo) HasLivePledge() bool`

HasLivePledge returns a boolean if a field has been set.

### GetLiveSaturation

`func (o *PoolInfo) GetLiveSaturation() string`

GetLiveSaturation returns the LiveSaturation field if non-nil, zero value otherwise.

### GetLiveSaturationOk

`func (o *PoolInfo) GetLiveSaturationOk() (*string, bool)`

GetLiveSaturationOk returns a tuple with the LiveSaturation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveSaturation

`func (o *PoolInfo) SetLiveSaturation(v string)`

SetLiveSaturation sets LiveSaturation field to given value.

### HasLiveSaturation

`func (o *PoolInfo) HasLiveSaturation() bool`

HasLiveSaturation returns a boolean if a field has been set.

### GetLiveStake

`func (o *PoolInfo) GetLiveStake() int64`

GetLiveStake returns the LiveStake field if non-nil, zero value otherwise.

### GetLiveStakeOk

`func (o *PoolInfo) GetLiveStakeOk() (*int64, bool)`

GetLiveStakeOk returns a tuple with the LiveStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStake

`func (o *PoolInfo) SetLiveStake(v int64)`

SetLiveStake sets LiveStake field to given value.

### HasLiveStake

`func (o *PoolInfo) HasLiveStake() bool`

HasLiveStake returns a boolean if a field has been set.

### GetMargin

`func (o *PoolInfo) GetMargin() float32`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *PoolInfo) GetMarginOk() (*float32, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *PoolInfo) SetMargin(v float32)`

SetMargin sets Margin field to given value.


### GetMetaHash

`func (o *PoolInfo) GetMetaHash() string`

GetMetaHash returns the MetaHash field if non-nil, zero value otherwise.

### GetMetaHashOk

`func (o *PoolInfo) GetMetaHashOk() (*string, bool)`

GetMetaHashOk returns a tuple with the MetaHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaHash

`func (o *PoolInfo) SetMetaHash(v string)`

SetMetaHash sets MetaHash field to given value.

### HasMetaHash

`func (o *PoolInfo) HasMetaHash() bool`

HasMetaHash returns a boolean if a field has been set.

### GetMetaJson

`func (o *PoolInfo) GetMetaJson() PoolMetaJson`

GetMetaJson returns the MetaJson field if non-nil, zero value otherwise.

### GetMetaJsonOk

`func (o *PoolInfo) GetMetaJsonOk() (*PoolMetaJson, bool)`

GetMetaJsonOk returns a tuple with the MetaJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaJson

`func (o *PoolInfo) SetMetaJson(v PoolMetaJson)`

SetMetaJson sets MetaJson field to given value.

### HasMetaJson

`func (o *PoolInfo) HasMetaJson() bool`

HasMetaJson returns a boolean if a field has been set.

### GetMetaUrl

`func (o *PoolInfo) GetMetaUrl() string`

GetMetaUrl returns the MetaUrl field if non-nil, zero value otherwise.

### GetMetaUrlOk

`func (o *PoolInfo) GetMetaUrlOk() (*string, bool)`

GetMetaUrlOk returns a tuple with the MetaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaUrl

`func (o *PoolInfo) SetMetaUrl(v string)`

SetMetaUrl sets MetaUrl field to given value.

### HasMetaUrl

`func (o *PoolInfo) HasMetaUrl() bool`

HasMetaUrl returns a boolean if a field has been set.

### GetOpCert

`func (o *PoolInfo) GetOpCert() string`

GetOpCert returns the OpCert field if non-nil, zero value otherwise.

### GetOpCertOk

`func (o *PoolInfo) GetOpCertOk() (*string, bool)`

GetOpCertOk returns a tuple with the OpCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpCert

`func (o *PoolInfo) SetOpCert(v string)`

SetOpCert sets OpCert field to given value.

### HasOpCert

`func (o *PoolInfo) HasOpCert() bool`

HasOpCert returns a boolean if a field has been set.

### GetOpCertCounter

`func (o *PoolInfo) GetOpCertCounter() int64`

GetOpCertCounter returns the OpCertCounter field if non-nil, zero value otherwise.

### GetOpCertCounterOk

`func (o *PoolInfo) GetOpCertCounterOk() (*int64, bool)`

GetOpCertCounterOk returns a tuple with the OpCertCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpCertCounter

`func (o *PoolInfo) SetOpCertCounter(v int64)`

SetOpCertCounter sets OpCertCounter field to given value.

### HasOpCertCounter

`func (o *PoolInfo) HasOpCertCounter() bool`

HasOpCertCounter returns a boolean if a field has been set.

### GetOwners

`func (o *PoolInfo) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *PoolInfo) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *PoolInfo) SetOwners(v []string)`

SetOwners sets Owners field to given value.


### GetPledge

`func (o *PoolInfo) GetPledge() int64`

GetPledge returns the Pledge field if non-nil, zero value otherwise.

### GetPledgeOk

`func (o *PoolInfo) GetPledgeOk() (*int64, bool)`

GetPledgeOk returns a tuple with the Pledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPledge

`func (o *PoolInfo) SetPledge(v int64)`

SetPledge sets Pledge field to given value.


### GetPoolIdBech32

`func (o *PoolInfo) GetPoolIdBech32() string`

GetPoolIdBech32 returns the PoolIdBech32 field if non-nil, zero value otherwise.

### GetPoolIdBech32Ok

`func (o *PoolInfo) GetPoolIdBech32Ok() (*string, bool)`

GetPoolIdBech32Ok returns a tuple with the PoolIdBech32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIdBech32

`func (o *PoolInfo) SetPoolIdBech32(v string)`

SetPoolIdBech32 sets PoolIdBech32 field to given value.


### GetPoolIdHex

`func (o *PoolInfo) GetPoolIdHex() string`

GetPoolIdHex returns the PoolIdHex field if non-nil, zero value otherwise.

### GetPoolIdHexOk

`func (o *PoolInfo) GetPoolIdHexOk() (*string, bool)`

GetPoolIdHexOk returns a tuple with the PoolIdHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIdHex

`func (o *PoolInfo) SetPoolIdHex(v string)`

SetPoolIdHex sets PoolIdHex field to given value.


### GetPoolStatus

`func (o *PoolInfo) GetPoolStatus() string`

GetPoolStatus returns the PoolStatus field if non-nil, zero value otherwise.

### GetPoolStatusOk

`func (o *PoolInfo) GetPoolStatusOk() (*string, bool)`

GetPoolStatusOk returns a tuple with the PoolStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolStatus

`func (o *PoolInfo) SetPoolStatus(v string)`

SetPoolStatus sets PoolStatus field to given value.

### HasPoolStatus

`func (o *PoolInfo) HasPoolStatus() bool`

HasPoolStatus returns a boolean if a field has been set.

### GetRelays

`func (o *PoolInfo) GetRelays() []Relay`

GetRelays returns the Relays field if non-nil, zero value otherwise.

### GetRelaysOk

`func (o *PoolInfo) GetRelaysOk() (*[]Relay, bool)`

GetRelaysOk returns a tuple with the Relays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelays

`func (o *PoolInfo) SetRelays(v []Relay)`

SetRelays sets Relays field to given value.


### GetRetiringEpoch

`func (o *PoolInfo) GetRetiringEpoch() int32`

GetRetiringEpoch returns the RetiringEpoch field if non-nil, zero value otherwise.

### GetRetiringEpochOk

`func (o *PoolInfo) GetRetiringEpochOk() (*int32, bool)`

GetRetiringEpochOk returns a tuple with the RetiringEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetiringEpoch

`func (o *PoolInfo) SetRetiringEpoch(v int32)`

SetRetiringEpoch sets RetiringEpoch field to given value.

### HasRetiringEpoch

`func (o *PoolInfo) HasRetiringEpoch() bool`

HasRetiringEpoch returns a boolean if a field has been set.

### GetRewardAddr

`func (o *PoolInfo) GetRewardAddr() string`

GetRewardAddr returns the RewardAddr field if non-nil, zero value otherwise.

### GetRewardAddrOk

`func (o *PoolInfo) GetRewardAddrOk() (*string, bool)`

GetRewardAddrOk returns a tuple with the RewardAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAddr

`func (o *PoolInfo) SetRewardAddr(v string)`

SetRewardAddr sets RewardAddr field to given value.

### HasRewardAddr

`func (o *PoolInfo) HasRewardAddr() bool`

HasRewardAddr returns a boolean if a field has been set.

### GetSigma

`func (o *PoolInfo) GetSigma() string`

GetSigma returns the Sigma field if non-nil, zero value otherwise.

### GetSigmaOk

`func (o *PoolInfo) GetSigmaOk() (*string, bool)`

GetSigmaOk returns a tuple with the Sigma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigma

`func (o *PoolInfo) SetSigma(v string)`

SetSigma sets Sigma field to given value.

### HasSigma

`func (o *PoolInfo) HasSigma() bool`

HasSigma returns a boolean if a field has been set.

### GetVrfKeyHash

`func (o *PoolInfo) GetVrfKeyHash() string`

GetVrfKeyHash returns the VrfKeyHash field if non-nil, zero value otherwise.

### GetVrfKeyHashOk

`func (o *PoolInfo) GetVrfKeyHashOk() (*string, bool)`

GetVrfKeyHashOk returns a tuple with the VrfKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfKeyHash

`func (o *PoolInfo) SetVrfKeyHash(v string)`

SetVrfKeyHash sets VrfKeyHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


