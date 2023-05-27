/*
Blockchain Indexer API

The **Maestro Blockchain Indexer API** follows the [REST](https://restfulapi.net/) paradigm. To interact with Mapi, please  head over to [Dashboards](https://dashboard.gomaestro.org), create an API project, and copy its associated long-lived API key into your request header.  Your Mapi project is rate-limited based on your API package tier. Please see the available [Packages](https://dashboard.gomaestro.org/pricing) for more details or to upgrade your plan.  Example `GET` request for retrieving the chain tip: ``` curl -X GET --header \"api-key: <your_project_api_key>\" https://mainnet.gomaestro-api.org/v0/chain-tip ```  Example `POST` request for submitting a transaction: ``` curl -X POST --header \"Content-Type: application/cbor\" --header \"api-key: <your_project_api_key>\" --data @tx.signed https://mainnet.gomaestro-api.org/v0/transactions ```

API version: V0
Contact: info@gomaestro.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package maestro

import (
	"encoding/json"
)

// checks if the DelegatorInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DelegatorInfo{}

// DelegatorInfo Information summary of a delegator
type DelegatorInfo struct {
	// Epoch at which the delegation becomes active
	ActiveEpochNo *int64 `json:"active_epoch_no,omitempty"`
	// Delegator live stake
	Amount *int64 `json:"amount,omitempty"`
	// Transaction hash relating to the most recent delegation
	LatestDelegationTxHash *string `json:"latest_delegation_tx_hash,omitempty"`
	// Bech32 encoded stake address (reward address)
	StakeAddress *string `json:"stake_address,omitempty"`
}

// NewDelegatorInfo instantiates a new DelegatorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelegatorInfo() *DelegatorInfo {
	this := DelegatorInfo{}
	return &this
}

// NewDelegatorInfoWithDefaults instantiates a new DelegatorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelegatorInfoWithDefaults() *DelegatorInfo {
	this := DelegatorInfo{}
	return &this
}

// GetActiveEpochNo returns the ActiveEpochNo field value if set, zero value otherwise.
func (o *DelegatorInfo) GetActiveEpochNo() int64 {
	if o == nil || IsNil(o.ActiveEpochNo) {
		var ret int64
		return ret
	}
	return *o.ActiveEpochNo
}

// GetActiveEpochNoOk returns a tuple with the ActiveEpochNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatorInfo) GetActiveEpochNoOk() (*int64, bool) {
	if o == nil || IsNil(o.ActiveEpochNo) {
		return nil, false
	}
	return o.ActiveEpochNo, true
}

// HasActiveEpochNo returns a boolean if a field has been set.
func (o *DelegatorInfo) HasActiveEpochNo() bool {
	if o != nil && !IsNil(o.ActiveEpochNo) {
		return true
	}

	return false
}

// SetActiveEpochNo gets a reference to the given int64 and assigns it to the ActiveEpochNo field.
func (o *DelegatorInfo) SetActiveEpochNo(v int64) {
	o.ActiveEpochNo = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DelegatorInfo) GetAmount() int64 {
	if o == nil || IsNil(o.Amount) {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatorInfo) GetAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DelegatorInfo) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *DelegatorInfo) SetAmount(v int64) {
	o.Amount = &v
}

// GetLatestDelegationTxHash returns the LatestDelegationTxHash field value if set, zero value otherwise.
func (o *DelegatorInfo) GetLatestDelegationTxHash() string {
	if o == nil || IsNil(o.LatestDelegationTxHash) {
		var ret string
		return ret
	}
	return *o.LatestDelegationTxHash
}

// GetLatestDelegationTxHashOk returns a tuple with the LatestDelegationTxHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatorInfo) GetLatestDelegationTxHashOk() (*string, bool) {
	if o == nil || IsNil(o.LatestDelegationTxHash) {
		return nil, false
	}
	return o.LatestDelegationTxHash, true
}

// HasLatestDelegationTxHash returns a boolean if a field has been set.
func (o *DelegatorInfo) HasLatestDelegationTxHash() bool {
	if o != nil && !IsNil(o.LatestDelegationTxHash) {
		return true
	}

	return false
}

// SetLatestDelegationTxHash gets a reference to the given string and assigns it to the LatestDelegationTxHash field.
func (o *DelegatorInfo) SetLatestDelegationTxHash(v string) {
	o.LatestDelegationTxHash = &v
}

// GetStakeAddress returns the StakeAddress field value if set, zero value otherwise.
func (o *DelegatorInfo) GetStakeAddress() string {
	if o == nil || IsNil(o.StakeAddress) {
		var ret string
		return ret
	}
	return *o.StakeAddress
}

// GetStakeAddressOk returns a tuple with the StakeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatorInfo) GetStakeAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StakeAddress) {
		return nil, false
	}
	return o.StakeAddress, true
}

// HasStakeAddress returns a boolean if a field has been set.
func (o *DelegatorInfo) HasStakeAddress() bool {
	if o != nil && !IsNil(o.StakeAddress) {
		return true
	}

	return false
}

// SetStakeAddress gets a reference to the given string and assigns it to the StakeAddress field.
func (o *DelegatorInfo) SetStakeAddress(v string) {
	o.StakeAddress = &v
}

func (o DelegatorInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DelegatorInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveEpochNo) {
		toSerialize["active_epoch_no"] = o.ActiveEpochNo
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.LatestDelegationTxHash) {
		toSerialize["latest_delegation_tx_hash"] = o.LatestDelegationTxHash
	}
	if !IsNil(o.StakeAddress) {
		toSerialize["stake_address"] = o.StakeAddress
	}
	return toSerialize, nil
}

type NullableDelegatorInfo struct {
	value *DelegatorInfo
	isSet bool
}

func (v NullableDelegatorInfo) Get() *DelegatorInfo {
	return v.value
}

func (v *NullableDelegatorInfo) Set(val *DelegatorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDelegatorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDelegatorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelegatorInfo(val *DelegatorInfo) *NullableDelegatorInfo {
	return &NullableDelegatorInfo{value: val, isSet: true}
}

func (v NullableDelegatorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelegatorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


