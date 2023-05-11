# PoolRelay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolIdBech32** | **string** | Bech32 encoded pool ID | 
**Relays** | [**[]Relay**](Relay.md) |  | 

## Methods

### NewPoolRelay

`func NewPoolRelay(poolIdBech32 string, relays []Relay, ) *PoolRelay`

NewPoolRelay instantiates a new PoolRelay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolRelayWithDefaults

`func NewPoolRelayWithDefaults() *PoolRelay`

NewPoolRelayWithDefaults instantiates a new PoolRelay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolIdBech32

`func (o *PoolRelay) GetPoolIdBech32() string`

GetPoolIdBech32 returns the PoolIdBech32 field if non-nil, zero value otherwise.

### GetPoolIdBech32Ok

`func (o *PoolRelay) GetPoolIdBech32Ok() (*string, bool)`

GetPoolIdBech32Ok returns a tuple with the PoolIdBech32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIdBech32

`func (o *PoolRelay) SetPoolIdBech32(v string)`

SetPoolIdBech32 sets PoolIdBech32 field to given value.


### GetRelays

`func (o *PoolRelay) GetRelays() []Relay`

GetRelays returns the Relays field if non-nil, zero value otherwise.

### GetRelaysOk

`func (o *PoolRelay) GetRelaysOk() (*[]Relay, bool)`

GetRelaysOk returns a tuple with the Relays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelays

`func (o *PoolRelay) SetRelays(v []Relay)`

SetRelays sets Relays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


