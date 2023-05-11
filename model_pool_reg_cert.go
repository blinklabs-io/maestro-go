/*
Blockchain Indexer API

The **Maestro Blockchain Indexer API** follows the [REST](https://restfulapi.net/) paradigm. To interact with Mapi, please  head over to [Dashboards](https://dashboard.gomaestro.org), create an API project, and copy its associated long-lived API key into your request header.  Your Mapi project is rate-limited based on your API package tier. Please see the available [Packages](https://dashboard.gomaestro.org/pricing) for more details or to upgrade your plan.  Example `GET` request for retrieving the chain tip: ``` curl -X GET --header \"api-key: <your_project_api_key>\" https://mainnet.gomaestro-api.org/v0/chain-tip ```  Example `POST` request for submitting a transaction: ``` curl -X POST --header \"Content-Type: application/cbor\" --header \"api-key: <your_project_api_key>\" --data @tx.signed https://mainnet.gomaestro-api.org/v0/transactions ```

API version: V0
Contact: info@gomaestro.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PoolRegCert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolRegCert{}

// PoolRegCert Certificate for registering or updating a stake pool
type PoolRegCert struct {
	// Index of the certificate in the transaction
	CertIndex int32 `json:"cert_index"`
	// Pool fixed cost
	FixedCost int64 `json:"fixed_cost"`
	// Epoch at which the update will become active
	FromEpoch int32 `json:"from_epoch"`
	// Pool margin
	Margin float32 `json:"margin"`
	// Hash of metadata that the metadata URL should point to
	MetadataHash *string `json:"metadata_hash,omitempty"`
	// URL pointing to pool metadata declared by the stake pool
	MetadataUrl *string `json:"metadata_url,omitempty"`
	OwnerAddresses []string `json:"owner_addresses"`
	// Pool pledge
	Pledge int64 `json:"pledge"`
	// Pool ID of the stake pool being updated
	PoolId string `json:"pool_id"`
	Relays []Relay `json:"relays"`
	// Stake address which will receive rewards from the stake pool
	RewardAddress string `json:"reward_address"`
	// VRF key hash of the stake pool
	VrfKeyHash string `json:"vrf_key_hash"`
}

// NewPoolRegCert instantiates a new PoolRegCert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolRegCert(certIndex int32, fixedCost int64, fromEpoch int32, margin float32, ownerAddresses []string, pledge int64, poolId string, relays []Relay, rewardAddress string, vrfKeyHash string) *PoolRegCert {
	this := PoolRegCert{}
	this.CertIndex = certIndex
	this.FixedCost = fixedCost
	this.FromEpoch = fromEpoch
	this.Margin = margin
	this.OwnerAddresses = ownerAddresses
	this.Pledge = pledge
	this.PoolId = poolId
	this.Relays = relays
	this.RewardAddress = rewardAddress
	this.VrfKeyHash = vrfKeyHash
	return &this
}

// NewPoolRegCertWithDefaults instantiates a new PoolRegCert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolRegCertWithDefaults() *PoolRegCert {
	this := PoolRegCert{}
	return &this
}

// GetCertIndex returns the CertIndex field value
func (o *PoolRegCert) GetCertIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CertIndex
}

// GetCertIndexOk returns a tuple with the CertIndex field value
// and a boolean to check if the value has been set.
func (o *PoolRegCert) GetCertIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertIndex, true
}

// SetCertIndex sets field value
func (o *PoolRegCert) SetCertIndex(v int32) {
	o.CertIndex = v
}

// GetFixedCost returns the FixedCost field value
func (o *PoolRegCert) GetFixedCost() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FixedCost
}

// GetFixedCostOk returns a tuple with the FixedCost field value
// and a boolean to check if the value has been set.
func (o *PoolRegCert) GetFixedCostOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FixedCost, true
}

// SetFixedCost sets field value
func (o *PoolRegCert) SetFixedCost(v int64) {
	o.FixedCost = v
}

// GetFromEpoch returns the FromEpoch field value
func (o *PoolRegCert) GetFromEpoch() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FromEpoch
}

// GetFromEpochOk returns a tuple with the FromEpoch field value
// and a boolean to check if the value has been set.
func (o *PoolRegCert) GetFromEpochOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromEpoch, true
}

// SetFromEpoch sets field value
func (o *PoolRegCert) SetFromEpoch(v int32) {
	o.FromEpoch = v
}

// GetMargin returns the Margin field value
func (o *PoolRegCert) GetMargin() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Margin
}

// GetMarginOk returns a tuple with the Margin field value
// and a boolean to check if the value has been set.
func (o *PoolRegCert) GetMarginOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Margin, true
}

// SetMargin sets field value
func (o *PoolRegCert) SetMargin(v float32) {
	o.Margin = v
}

// GetMetadataHash returns the MetadataHash field value if set, zero value otherwise.
func (o *PoolRegCert) GetMetadataHash() string {
	if o == nil || IsNil(o.MetadataHash) {
		var ret string
		return ret
	}
	return *o.MetadataHash
}

// GetMetadataHashOk returns a tuple with the MetadataHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolRegCert) GetMetadataHashOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataHash) {
		return nil, false
	}
	return o.MetadataHash, true
}

// HasMetadataHash returns a boolean if a field has been set.
func (o *PoolRegCert) HasMetadataHash() bool {
	if o != nil && !IsNil(o.MetadataHash) {
		return true
	}

	return false
}

// SetMetadataHash gets a reference to the given string and assigns it to the MetadataHash field.
func (o *PoolRegCert) SetMetadataHash(v string) {
	o.MetadataHash = &v
}

// GetMetadataUrl returns the MetadataUrl field value if set, zero value otherwise.
func (o *PoolRegCert) GetMetadataUrl() string {
	if o == nil || IsNil(o.MetadataUrl) {
		var ret string
		return ret
	}
	return *o.MetadataUrl
}

// GetMetadataUrlOk returns a tuple with the MetadataUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolRegCert) GetMetadataUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataUrl) {
		return nil, false
	}
	return o.MetadataUrl, true
}

// HasMetadataUrl returns a boolean if a field has been set.
func (o *PoolRegCert) HasMetadataUrl() bool {
	if o != nil && !IsNil(o.MetadataUrl) {
		return true
	}

	return false
}

// SetMetadataUrl gets a reference to the given string and assigns it to the MetadataUrl field.
func (o *PoolRegCert) SetMetadataUrl(v string) {
	o.MetadataUrl = &v
}

// GetOwnerAddresses returns the OwnerAddresses field value
func (o *PoolRegCert) GetOwnerAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OwnerAddresses
}

// GetOwnerAddressesOk returns a tuple with the OwnerAddresses field value
// and a boolean to check if the value has been set.
func (o *PoolRegCert) GetOwnerAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerAddresses, true
}

// SetOwnerAddresses sets field value
func (o *PoolRegCert) SetOwnerAddresses(v []string) {
	o.OwnerAddresses = v
}

// GetPledge returns the Pledge field value
func (o *PoolRegCert) GetPledge() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Pledge
}

// GetPledgeOk returns a tuple with the Pledge field value
// and a boolean to check if the value has been set.
func (o *PoolRegCert) GetPledgeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pledge, true
}

// SetPledge sets field value
func (o *PoolRegCert) SetPledge(v int64) {
	o.Pledge = v
}

// GetPoolId returns the PoolId field value
func (o *PoolRegCert) GetPoolId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value
// and a boolean to check if the value has been set.
func (o *PoolRegCert) GetPoolIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolId, true
}

// SetPoolId sets field value
func (o *PoolRegCert) SetPoolId(v string) {
	o.PoolId = v
}

// GetRelays returns the Relays field value
func (o *PoolRegCert) GetRelays() []Relay {
	if o == nil {
		var ret []Relay
		return ret
	}

	return o.Relays
}

// GetRelaysOk returns a tuple with the Relays field value
// and a boolean to check if the value has been set.
func (o *PoolRegCert) GetRelaysOk() ([]Relay, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relays, true
}

// SetRelays sets field value
func (o *PoolRegCert) SetRelays(v []Relay) {
	o.Relays = v
}

// GetRewardAddress returns the RewardAddress field value
func (o *PoolRegCert) GetRewardAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RewardAddress
}

// GetRewardAddressOk returns a tuple with the RewardAddress field value
// and a boolean to check if the value has been set.
func (o *PoolRegCert) GetRewardAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardAddress, true
}

// SetRewardAddress sets field value
func (o *PoolRegCert) SetRewardAddress(v string) {
	o.RewardAddress = v
}

// GetVrfKeyHash returns the VrfKeyHash field value
func (o *PoolRegCert) GetVrfKeyHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VrfKeyHash
}

// GetVrfKeyHashOk returns a tuple with the VrfKeyHash field value
// and a boolean to check if the value has been set.
func (o *PoolRegCert) GetVrfKeyHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VrfKeyHash, true
}

// SetVrfKeyHash sets field value
func (o *PoolRegCert) SetVrfKeyHash(v string) {
	o.VrfKeyHash = v
}

func (o PoolRegCert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolRegCert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cert_index"] = o.CertIndex
	toSerialize["fixed_cost"] = o.FixedCost
	toSerialize["from_epoch"] = o.FromEpoch
	toSerialize["margin"] = o.Margin
	if !IsNil(o.MetadataHash) {
		toSerialize["metadata_hash"] = o.MetadataHash
	}
	if !IsNil(o.MetadataUrl) {
		toSerialize["metadata_url"] = o.MetadataUrl
	}
	toSerialize["owner_addresses"] = o.OwnerAddresses
	toSerialize["pledge"] = o.Pledge
	toSerialize["pool_id"] = o.PoolId
	toSerialize["relays"] = o.Relays
	toSerialize["reward_address"] = o.RewardAddress
	toSerialize["vrf_key_hash"] = o.VrfKeyHash
	return toSerialize, nil
}

type NullablePoolRegCert struct {
	value *PoolRegCert
	isSet bool
}

func (v NullablePoolRegCert) Get() *PoolRegCert {
	return v.value
}

func (v *NullablePoolRegCert) Set(val *PoolRegCert) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolRegCert) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolRegCert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolRegCert(val *PoolRegCert) *NullablePoolRegCert {
	return &NullablePoolRegCert{value: val, isSet: true}
}

func (v NullablePoolRegCert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolRegCert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


