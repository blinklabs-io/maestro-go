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

// checks if the AccountInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountInfo{}

// AccountInfo Summary of information regarding a stake account
type AccountInfo struct {
	// Bech32 pool ID that the stake key is delegated to
	DelegatedPool *string `json:"delegated_pool,omitempty"`
	// True if the stake key is registered
	Registered bool `json:"registered"`
	// The amount of rewards that are available to be withdrawn
	RewardsAvailable int64 `json:"rewards_available"`
	// Bech32 encoded stake address
	StakeAddress string `json:"stake_address"`
	// Total balance controlled by the stake key (sum of UTxO and rewards)
	TotalBalance int64 `json:"total_balance"`
	// Total rewards earned
	TotalRewarded int64 `json:"total_rewarded"`
	// Total rewards withdrawn
	TotalWithdrawn int64 `json:"total_withdrawn"`
	// Amount locked in UTxOs controlled by addresses with the stake key
	UtxoBalance int64 `json:"utxo_balance"`
}

// NewAccountInfo instantiates a new AccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInfo(registered bool, rewardsAvailable int64, stakeAddress string, totalBalance int64, totalRewarded int64, totalWithdrawn int64, utxoBalance int64) *AccountInfo {
	this := AccountInfo{}
	this.Registered = registered
	this.RewardsAvailable = rewardsAvailable
	this.StakeAddress = stakeAddress
	this.TotalBalance = totalBalance
	this.TotalRewarded = totalRewarded
	this.TotalWithdrawn = totalWithdrawn
	this.UtxoBalance = utxoBalance
	return &this
}

// NewAccountInfoWithDefaults instantiates a new AccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInfoWithDefaults() *AccountInfo {
	this := AccountInfo{}
	return &this
}

// GetDelegatedPool returns the DelegatedPool field value if set, zero value otherwise.
func (o *AccountInfo) GetDelegatedPool() string {
	if o == nil || IsNil(o.DelegatedPool) {
		var ret string
		return ret
	}
	return *o.DelegatedPool
}

// GetDelegatedPoolOk returns a tuple with the DelegatedPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetDelegatedPoolOk() (*string, bool) {
	if o == nil || IsNil(o.DelegatedPool) {
		return nil, false
	}
	return o.DelegatedPool, true
}

// HasDelegatedPool returns a boolean if a field has been set.
func (o *AccountInfo) HasDelegatedPool() bool {
	if o != nil && !IsNil(o.DelegatedPool) {
		return true
	}

	return false
}

// SetDelegatedPool gets a reference to the given string and assigns it to the DelegatedPool field.
func (o *AccountInfo) SetDelegatedPool(v string) {
	o.DelegatedPool = &v
}

// GetRegistered returns the Registered field value
func (o *AccountInfo) GetRegistered() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Registered
}

// GetRegisteredOk returns a tuple with the Registered field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetRegisteredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Registered, true
}

// SetRegistered sets field value
func (o *AccountInfo) SetRegistered(v bool) {
	o.Registered = v
}

// GetRewardsAvailable returns the RewardsAvailable field value
func (o *AccountInfo) GetRewardsAvailable() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RewardsAvailable
}

// GetRewardsAvailableOk returns a tuple with the RewardsAvailable field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetRewardsAvailableOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardsAvailable, true
}

// SetRewardsAvailable sets field value
func (o *AccountInfo) SetRewardsAvailable(v int64) {
	o.RewardsAvailable = v
}

// GetStakeAddress returns the StakeAddress field value
func (o *AccountInfo) GetStakeAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StakeAddress
}

// GetStakeAddressOk returns a tuple with the StakeAddress field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetStakeAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakeAddress, true
}

// SetStakeAddress sets field value
func (o *AccountInfo) SetStakeAddress(v string) {
	o.StakeAddress = v
}

// GetTotalBalance returns the TotalBalance field value
func (o *AccountInfo) GetTotalBalance() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalBalance
}

// GetTotalBalanceOk returns a tuple with the TotalBalance field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetTotalBalanceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalBalance, true
}

// SetTotalBalance sets field value
func (o *AccountInfo) SetTotalBalance(v int64) {
	o.TotalBalance = v
}

// GetTotalRewarded returns the TotalRewarded field value
func (o *AccountInfo) GetTotalRewarded() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalRewarded
}

// GetTotalRewardedOk returns a tuple with the TotalRewarded field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetTotalRewardedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRewarded, true
}

// SetTotalRewarded sets field value
func (o *AccountInfo) SetTotalRewarded(v int64) {
	o.TotalRewarded = v
}

// GetTotalWithdrawn returns the TotalWithdrawn field value
func (o *AccountInfo) GetTotalWithdrawn() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalWithdrawn
}

// GetTotalWithdrawnOk returns a tuple with the TotalWithdrawn field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetTotalWithdrawnOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalWithdrawn, true
}

// SetTotalWithdrawn sets field value
func (o *AccountInfo) SetTotalWithdrawn(v int64) {
	o.TotalWithdrawn = v
}

// GetUtxoBalance returns the UtxoBalance field value
func (o *AccountInfo) GetUtxoBalance() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UtxoBalance
}

// GetUtxoBalanceOk returns a tuple with the UtxoBalance field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetUtxoBalanceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UtxoBalance, true
}

// SetUtxoBalance sets field value
func (o *AccountInfo) SetUtxoBalance(v int64) {
	o.UtxoBalance = v
}

func (o AccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DelegatedPool) {
		toSerialize["delegated_pool"] = o.DelegatedPool
	}
	toSerialize["registered"] = o.Registered
	toSerialize["rewards_available"] = o.RewardsAvailable
	toSerialize["stake_address"] = o.StakeAddress
	toSerialize["total_balance"] = o.TotalBalance
	toSerialize["total_rewarded"] = o.TotalRewarded
	toSerialize["total_withdrawn"] = o.TotalWithdrawn
	toSerialize["utxo_balance"] = o.UtxoBalance
	return toSerialize, nil
}

type NullableAccountInfo struct {
	value *AccountInfo
	isSet bool
}

func (v NullableAccountInfo) Get() *AccountInfo {
	return v.value
}

func (v *NullableAccountInfo) Set(val *AccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInfo(val *AccountInfo) *NullableAccountInfo {
	return &NullableAccountInfo{value: val, isSet: true}
}

func (v NullableAccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


