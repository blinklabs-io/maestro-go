# PoolUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveEpochNo** | **int64** | Epoch when the update takes effect | 
**BlockTime** | Pointer to **int32** | UNIX timestamp of the block containing the transaction | [optional] 
**FixedCost** | **int64** | Pool fixed cost | 
**Margin** | **float32** | Pool margin | 
**MetaHash** | Pointer to **string** | Hash of the pool metadata | [optional] 
**MetaJson** | Pointer to [**PoolMetaJson**](PoolMetaJson.md) |  | [optional] 
**MetaUrl** | Pointer to **string** | URL pointing to the pool metadata | [optional] 
**Owners** | **[]string** |  | 
**Pledge** | **int64** | Pool pledge | 
**PoolIdBech32** | **string** | Bech32 encoded pool ID | 
**PoolIdHex** | **string** | Hex encoded pool ID | 
**PoolStatus** | Pointer to **string** | Status of the pool | [optional] 
**Relays** | [**[]Relay**](Relay.md) |  | 
**RetiringEpoch** | Pointer to **int32** | Epoch at which the pool will be retired | [optional] 
**RewardAddr** | Pointer to **string** | Reward address associated with the pool | [optional] 
**TxHash** | **string** | Transaction hash for the transaction which contained the update | 
**VrfKeyHash** | **string** | VRF key hash | 

## Methods

### NewPoolUpdate

`func NewPoolUpdate(activeEpochNo int64, fixedCost int64, margin float32, owners []string, pledge int64, poolIdBech32 string, poolIdHex string, relays []Relay, txHash string, vrfKeyHash string, ) *PoolUpdate`

NewPoolUpdate instantiates a new PoolUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolUpdateWithDefaults

`func NewPoolUpdateWithDefaults() *PoolUpdate`

NewPoolUpdateWithDefaults instantiates a new PoolUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveEpochNo

`func (o *PoolUpdate) GetActiveEpochNo() int64`

GetActiveEpochNo returns the ActiveEpochNo field if non-nil, zero value otherwise.

### GetActiveEpochNoOk

`func (o *PoolUpdate) GetActiveEpochNoOk() (*int64, bool)`

GetActiveEpochNoOk returns a tuple with the ActiveEpochNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveEpochNo

`func (o *PoolUpdate) SetActiveEpochNo(v int64)`

SetActiveEpochNo sets ActiveEpochNo field to given value.


### GetBlockTime

`func (o *PoolUpdate) GetBlockTime() int32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *PoolUpdate) GetBlockTimeOk() (*int32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *PoolUpdate) SetBlockTime(v int32)`

SetBlockTime sets BlockTime field to given value.

### HasBlockTime

`func (o *PoolUpdate) HasBlockTime() bool`

HasBlockTime returns a boolean if a field has been set.

### GetFixedCost

`func (o *PoolUpdate) GetFixedCost() int64`

GetFixedCost returns the FixedCost field if non-nil, zero value otherwise.

### GetFixedCostOk

`func (o *PoolUpdate) GetFixedCostOk() (*int64, bool)`

GetFixedCostOk returns a tuple with the FixedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedCost

`func (o *PoolUpdate) SetFixedCost(v int64)`

SetFixedCost sets FixedCost field to given value.


### GetMargin

`func (o *PoolUpdate) GetMargin() float32`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *PoolUpdate) GetMarginOk() (*float32, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *PoolUpdate) SetMargin(v float32)`

SetMargin sets Margin field to given value.


### GetMetaHash

`func (o *PoolUpdate) GetMetaHash() string`

GetMetaHash returns the MetaHash field if non-nil, zero value otherwise.

### GetMetaHashOk

`func (o *PoolUpdate) GetMetaHashOk() (*string, bool)`

GetMetaHashOk returns a tuple with the MetaHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaHash

`func (o *PoolUpdate) SetMetaHash(v string)`

SetMetaHash sets MetaHash field to given value.

### HasMetaHash

`func (o *PoolUpdate) HasMetaHash() bool`

HasMetaHash returns a boolean if a field has been set.

### GetMetaJson

`func (o *PoolUpdate) GetMetaJson() PoolMetaJson`

GetMetaJson returns the MetaJson field if non-nil, zero value otherwise.

### GetMetaJsonOk

`func (o *PoolUpdate) GetMetaJsonOk() (*PoolMetaJson, bool)`

GetMetaJsonOk returns a tuple with the MetaJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaJson

`func (o *PoolUpdate) SetMetaJson(v PoolMetaJson)`

SetMetaJson sets MetaJson field to given value.

### HasMetaJson

`func (o *PoolUpdate) HasMetaJson() bool`

HasMetaJson returns a boolean if a field has been set.

### GetMetaUrl

`func (o *PoolUpdate) GetMetaUrl() string`

GetMetaUrl returns the MetaUrl field if non-nil, zero value otherwise.

### GetMetaUrlOk

`func (o *PoolUpdate) GetMetaUrlOk() (*string, bool)`

GetMetaUrlOk returns a tuple with the MetaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaUrl

`func (o *PoolUpdate) SetMetaUrl(v string)`

SetMetaUrl sets MetaUrl field to given value.

### HasMetaUrl

`func (o *PoolUpdate) HasMetaUrl() bool`

HasMetaUrl returns a boolean if a field has been set.

### GetOwners

`func (o *PoolUpdate) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *PoolUpdate) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *PoolUpdate) SetOwners(v []string)`

SetOwners sets Owners field to given value.


### GetPledge

`func (o *PoolUpdate) GetPledge() int64`

GetPledge returns the Pledge field if non-nil, zero value otherwise.

### GetPledgeOk

`func (o *PoolUpdate) GetPledgeOk() (*int64, bool)`

GetPledgeOk returns a tuple with the Pledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPledge

`func (o *PoolUpdate) SetPledge(v int64)`

SetPledge sets Pledge field to given value.


### GetPoolIdBech32

`func (o *PoolUpdate) GetPoolIdBech32() string`

GetPoolIdBech32 returns the PoolIdBech32 field if non-nil, zero value otherwise.

### GetPoolIdBech32Ok

`func (o *PoolUpdate) GetPoolIdBech32Ok() (*string, bool)`

GetPoolIdBech32Ok returns a tuple with the PoolIdBech32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIdBech32

`func (o *PoolUpdate) SetPoolIdBech32(v string)`

SetPoolIdBech32 sets PoolIdBech32 field to given value.


### GetPoolIdHex

`func (o *PoolUpdate) GetPoolIdHex() string`

GetPoolIdHex returns the PoolIdHex field if non-nil, zero value otherwise.

### GetPoolIdHexOk

`func (o *PoolUpdate) GetPoolIdHexOk() (*string, bool)`

GetPoolIdHexOk returns a tuple with the PoolIdHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIdHex

`func (o *PoolUpdate) SetPoolIdHex(v string)`

SetPoolIdHex sets PoolIdHex field to given value.


### GetPoolStatus

`func (o *PoolUpdate) GetPoolStatus() string`

GetPoolStatus returns the PoolStatus field if non-nil, zero value otherwise.

### GetPoolStatusOk

`func (o *PoolUpdate) GetPoolStatusOk() (*string, bool)`

GetPoolStatusOk returns a tuple with the PoolStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolStatus

`func (o *PoolUpdate) SetPoolStatus(v string)`

SetPoolStatus sets PoolStatus field to given value.

### HasPoolStatus

`func (o *PoolUpdate) HasPoolStatus() bool`

HasPoolStatus returns a boolean if a field has been set.

### GetRelays

`func (o *PoolUpdate) GetRelays() []Relay`

GetRelays returns the Relays field if non-nil, zero value otherwise.

### GetRelaysOk

`func (o *PoolUpdate) GetRelaysOk() (*[]Relay, bool)`

GetRelaysOk returns a tuple with the Relays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelays

`func (o *PoolUpdate) SetRelays(v []Relay)`

SetRelays sets Relays field to given value.


### GetRetiringEpoch

`func (o *PoolUpdate) GetRetiringEpoch() int32`

GetRetiringEpoch returns the RetiringEpoch field if non-nil, zero value otherwise.

### GetRetiringEpochOk

`func (o *PoolUpdate) GetRetiringEpochOk() (*int32, bool)`

GetRetiringEpochOk returns a tuple with the RetiringEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetiringEpoch

`func (o *PoolUpdate) SetRetiringEpoch(v int32)`

SetRetiringEpoch sets RetiringEpoch field to given value.

### HasRetiringEpoch

`func (o *PoolUpdate) HasRetiringEpoch() bool`

HasRetiringEpoch returns a boolean if a field has been set.

### GetRewardAddr

`func (o *PoolUpdate) GetRewardAddr() string`

GetRewardAddr returns the RewardAddr field if non-nil, zero value otherwise.

### GetRewardAddrOk

`func (o *PoolUpdate) GetRewardAddrOk() (*string, bool)`

GetRewardAddrOk returns a tuple with the RewardAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAddr

`func (o *PoolUpdate) SetRewardAddr(v string)`

SetRewardAddr sets RewardAddr field to given value.

### HasRewardAddr

`func (o *PoolUpdate) HasRewardAddr() bool`

HasRewardAddr returns a boolean if a field has been set.

### GetTxHash

`func (o *PoolUpdate) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *PoolUpdate) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *PoolUpdate) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetVrfKeyHash

`func (o *PoolUpdate) GetVrfKeyHash() string`

GetVrfKeyHash returns the VrfKeyHash field if non-nil, zero value otherwise.

### GetVrfKeyHashOk

`func (o *PoolUpdate) GetVrfKeyHashOk() (*string, bool)`

GetVrfKeyHashOk returns a tuple with the VrfKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfKeyHash

`func (o *PoolUpdate) SetVrfKeyHash(v string)`

SetVrfKeyHash sets VrfKeyHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


