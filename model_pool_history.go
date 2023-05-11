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

// checks if the PoolHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolHistory{}

// PoolHistory Per-epoch history of a stake pool
type PoolHistory struct {
	// Active stake in the epoch
	ActiveStake *int64 `json:"active_stake,omitempty"`
	// Pool active stake as percentage of total active stake
	ActiveStakePct *string `json:"active_stake_pct,omitempty"`
	// Blocks created in the epoch
	BlockCnt *int64 `json:"block_cnt,omitempty"`
	// Total rewards earned by pool delegators for the epoch
	DelegRewards int64 `json:"deleg_rewards"`
	// Delegators in the epoch
	DelegatorCnt *int64 `json:"delegator_cnt,omitempty"`
	// Epoch number
	EpochNo int64 `json:"epoch_no"`
	// Annual return percentage for delegators for the epoch
	EpochRos string `json:"epoch_ros"`
	// Pool fixed cost
	FixedCost int64 `json:"fixed_cost"`
	// Pool margin
	Margin *float32 `json:"margin,omitempty"`
	// Fees collected for the epoch
	PoolFees int64 `json:"pool_fees"`
	// Pool saturation percent
	SaturationPct *string `json:"saturation_pct,omitempty"`
}

// NewPoolHistory instantiates a new PoolHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolHistory(delegRewards int64, epochNo int64, epochRos string, fixedCost int64, poolFees int64) *PoolHistory {
	this := PoolHistory{}
	this.DelegRewards = delegRewards
	this.EpochNo = epochNo
	this.EpochRos = epochRos
	this.FixedCost = fixedCost
	this.PoolFees = poolFees
	return &this
}

// NewPoolHistoryWithDefaults instantiates a new PoolHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolHistoryWithDefaults() *PoolHistory {
	this := PoolHistory{}
	return &this
}

// GetActiveStake returns the ActiveStake field value if set, zero value otherwise.
func (o *PoolHistory) GetActiveStake() int64 {
	if o == nil || IsNil(o.ActiveStake) {
		var ret int64
		return ret
	}
	return *o.ActiveStake
}

// GetActiveStakeOk returns a tuple with the ActiveStake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolHistory) GetActiveStakeOk() (*int64, bool) {
	if o == nil || IsNil(o.ActiveStake) {
		return nil, false
	}
	return o.ActiveStake, true
}

// HasActiveStake returns a boolean if a field has been set.
func (o *PoolHistory) HasActiveStake() bool {
	if o != nil && !IsNil(o.ActiveStake) {
		return true
	}

	return false
}

// SetActiveStake gets a reference to the given int64 and assigns it to the ActiveStake field.
func (o *PoolHistory) SetActiveStake(v int64) {
	o.ActiveStake = &v
}

// GetActiveStakePct returns the ActiveStakePct field value if set, zero value otherwise.
func (o *PoolHistory) GetActiveStakePct() string {
	if o == nil || IsNil(o.ActiveStakePct) {
		var ret string
		return ret
	}
	return *o.ActiveStakePct
}

// GetActiveStakePctOk returns a tuple with the ActiveStakePct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolHistory) GetActiveStakePctOk() (*string, bool) {
	if o == nil || IsNil(o.ActiveStakePct) {
		return nil, false
	}
	return o.ActiveStakePct, true
}

// HasActiveStakePct returns a boolean if a field has been set.
func (o *PoolHistory) HasActiveStakePct() bool {
	if o != nil && !IsNil(o.ActiveStakePct) {
		return true
	}

	return false
}

// SetActiveStakePct gets a reference to the given string and assigns it to the ActiveStakePct field.
func (o *PoolHistory) SetActiveStakePct(v string) {
	o.ActiveStakePct = &v
}

// GetBlockCnt returns the BlockCnt field value if set, zero value otherwise.
func (o *PoolHistory) GetBlockCnt() int64 {
	if o == nil || IsNil(o.BlockCnt) {
		var ret int64
		return ret
	}
	return *o.BlockCnt
}

// GetBlockCntOk returns a tuple with the BlockCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolHistory) GetBlockCntOk() (*int64, bool) {
	if o == nil || IsNil(o.BlockCnt) {
		return nil, false
	}
	return o.BlockCnt, true
}

// HasBlockCnt returns a boolean if a field has been set.
func (o *PoolHistory) HasBlockCnt() bool {
	if o != nil && !IsNil(o.BlockCnt) {
		return true
	}

	return false
}

// SetBlockCnt gets a reference to the given int64 and assigns it to the BlockCnt field.
func (o *PoolHistory) SetBlockCnt(v int64) {
	o.BlockCnt = &v
}

// GetDelegRewards returns the DelegRewards field value
func (o *PoolHistory) GetDelegRewards() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DelegRewards
}

// GetDelegRewardsOk returns a tuple with the DelegRewards field value
// and a boolean to check if the value has been set.
func (o *PoolHistory) GetDelegRewardsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DelegRewards, true
}

// SetDelegRewards sets field value
func (o *PoolHistory) SetDelegRewards(v int64) {
	o.DelegRewards = v
}

// GetDelegatorCnt returns the DelegatorCnt field value if set, zero value otherwise.
func (o *PoolHistory) GetDelegatorCnt() int64 {
	if o == nil || IsNil(o.DelegatorCnt) {
		var ret int64
		return ret
	}
	return *o.DelegatorCnt
}

// GetDelegatorCntOk returns a tuple with the DelegatorCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolHistory) GetDelegatorCntOk() (*int64, bool) {
	if o == nil || IsNil(o.DelegatorCnt) {
		return nil, false
	}
	return o.DelegatorCnt, true
}

// HasDelegatorCnt returns a boolean if a field has been set.
func (o *PoolHistory) HasDelegatorCnt() bool {
	if o != nil && !IsNil(o.DelegatorCnt) {
		return true
	}

	return false
}

// SetDelegatorCnt gets a reference to the given int64 and assigns it to the DelegatorCnt field.
func (o *PoolHistory) SetDelegatorCnt(v int64) {
	o.DelegatorCnt = &v
}

// GetEpochNo returns the EpochNo field value
func (o *PoolHistory) GetEpochNo() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EpochNo
}

// GetEpochNoOk returns a tuple with the EpochNo field value
// and a boolean to check if the value has been set.
func (o *PoolHistory) GetEpochNoOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EpochNo, true
}

// SetEpochNo sets field value
func (o *PoolHistory) SetEpochNo(v int64) {
	o.EpochNo = v
}

// GetEpochRos returns the EpochRos field value
func (o *PoolHistory) GetEpochRos() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EpochRos
}

// GetEpochRosOk returns a tuple with the EpochRos field value
// and a boolean to check if the value has been set.
func (o *PoolHistory) GetEpochRosOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EpochRos, true
}

// SetEpochRos sets field value
func (o *PoolHistory) SetEpochRos(v string) {
	o.EpochRos = v
}

// GetFixedCost returns the FixedCost field value
func (o *PoolHistory) GetFixedCost() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FixedCost
}

// GetFixedCostOk returns a tuple with the FixedCost field value
// and a boolean to check if the value has been set.
func (o *PoolHistory) GetFixedCostOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FixedCost, true
}

// SetFixedCost sets field value
func (o *PoolHistory) SetFixedCost(v int64) {
	o.FixedCost = v
}

// GetMargin returns the Margin field value if set, zero value otherwise.
func (o *PoolHistory) GetMargin() float32 {
	if o == nil || IsNil(o.Margin) {
		var ret float32
		return ret
	}
	return *o.Margin
}

// GetMarginOk returns a tuple with the Margin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolHistory) GetMarginOk() (*float32, bool) {
	if o == nil || IsNil(o.Margin) {
		return nil, false
	}
	return o.Margin, true
}

// HasMargin returns a boolean if a field has been set.
func (o *PoolHistory) HasMargin() bool {
	if o != nil && !IsNil(o.Margin) {
		return true
	}

	return false
}

// SetMargin gets a reference to the given float32 and assigns it to the Margin field.
func (o *PoolHistory) SetMargin(v float32) {
	o.Margin = &v
}

// GetPoolFees returns the PoolFees field value
func (o *PoolHistory) GetPoolFees() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PoolFees
}

// GetPoolFeesOk returns a tuple with the PoolFees field value
// and a boolean to check if the value has been set.
func (o *PoolHistory) GetPoolFeesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolFees, true
}

// SetPoolFees sets field value
func (o *PoolHistory) SetPoolFees(v int64) {
	o.PoolFees = v
}

// GetSaturationPct returns the SaturationPct field value if set, zero value otherwise.
func (o *PoolHistory) GetSaturationPct() string {
	if o == nil || IsNil(o.SaturationPct) {
		var ret string
		return ret
	}
	return *o.SaturationPct
}

// GetSaturationPctOk returns a tuple with the SaturationPct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolHistory) GetSaturationPctOk() (*string, bool) {
	if o == nil || IsNil(o.SaturationPct) {
		return nil, false
	}
	return o.SaturationPct, true
}

// HasSaturationPct returns a boolean if a field has been set.
func (o *PoolHistory) HasSaturationPct() bool {
	if o != nil && !IsNil(o.SaturationPct) {
		return true
	}

	return false
}

// SetSaturationPct gets a reference to the given string and assigns it to the SaturationPct field.
func (o *PoolHistory) SetSaturationPct(v string) {
	o.SaturationPct = &v
}

func (o PoolHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveStake) {
		toSerialize["active_stake"] = o.ActiveStake
	}
	if !IsNil(o.ActiveStakePct) {
		toSerialize["active_stake_pct"] = o.ActiveStakePct
	}
	if !IsNil(o.BlockCnt) {
		toSerialize["block_cnt"] = o.BlockCnt
	}
	toSerialize["deleg_rewards"] = o.DelegRewards
	if !IsNil(o.DelegatorCnt) {
		toSerialize["delegator_cnt"] = o.DelegatorCnt
	}
	toSerialize["epoch_no"] = o.EpochNo
	toSerialize["epoch_ros"] = o.EpochRos
	toSerialize["fixed_cost"] = o.FixedCost
	if !IsNil(o.Margin) {
		toSerialize["margin"] = o.Margin
	}
	toSerialize["pool_fees"] = o.PoolFees
	if !IsNil(o.SaturationPct) {
		toSerialize["saturation_pct"] = o.SaturationPct
	}
	return toSerialize, nil
}

type NullablePoolHistory struct {
	value *PoolHistory
	isSet bool
}

func (v NullablePoolHistory) Get() *PoolHistory {
	return v.value
}

func (v *NullablePoolHistory) Set(val *PoolHistory) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolHistory) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolHistory(val *PoolHistory) *NullablePoolHistory {
	return &NullablePoolHistory{value: val, isSet: true}
}

func (v NullablePoolHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

