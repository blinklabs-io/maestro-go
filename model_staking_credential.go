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

// checks if the StakingCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StakingCredential{}

// StakingCredential struct for StakingCredential
type StakingCredential struct {
	Bech32 *string `json:"bech32,omitempty"`
	Hex *string `json:"hex,omitempty"`
	Kind StakingCredKind `json:"kind"`
	Pointer *Pointer `json:"pointer,omitempty"`
	RewardAddress *string `json:"reward_address,omitempty"`
}

// NewStakingCredential instantiates a new StakingCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStakingCredential(kind StakingCredKind) *StakingCredential {
	this := StakingCredential{}
	this.Kind = kind
	return &this
}

// NewStakingCredentialWithDefaults instantiates a new StakingCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStakingCredentialWithDefaults() *StakingCredential {
	this := StakingCredential{}
	return &this
}

// GetBech32 returns the Bech32 field value if set, zero value otherwise.
func (o *StakingCredential) GetBech32() string {
	if o == nil || IsNil(o.Bech32) {
		var ret string
		return ret
	}
	return *o.Bech32
}

// GetBech32Ok returns a tuple with the Bech32 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakingCredential) GetBech32Ok() (*string, bool) {
	if o == nil || IsNil(o.Bech32) {
		return nil, false
	}
	return o.Bech32, true
}

// HasBech32 returns a boolean if a field has been set.
func (o *StakingCredential) HasBech32() bool {
	if o != nil && !IsNil(o.Bech32) {
		return true
	}

	return false
}

// SetBech32 gets a reference to the given string and assigns it to the Bech32 field.
func (o *StakingCredential) SetBech32(v string) {
	o.Bech32 = &v
}

// GetHex returns the Hex field value if set, zero value otherwise.
func (o *StakingCredential) GetHex() string {
	if o == nil || IsNil(o.Hex) {
		var ret string
		return ret
	}
	return *o.Hex
}

// GetHexOk returns a tuple with the Hex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakingCredential) GetHexOk() (*string, bool) {
	if o == nil || IsNil(o.Hex) {
		return nil, false
	}
	return o.Hex, true
}

// HasHex returns a boolean if a field has been set.
func (o *StakingCredential) HasHex() bool {
	if o != nil && !IsNil(o.Hex) {
		return true
	}

	return false
}

// SetHex gets a reference to the given string and assigns it to the Hex field.
func (o *StakingCredential) SetHex(v string) {
	o.Hex = &v
}

// GetKind returns the Kind field value
func (o *StakingCredential) GetKind() StakingCredKind {
	if o == nil {
		var ret StakingCredKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *StakingCredential) GetKindOk() (*StakingCredKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *StakingCredential) SetKind(v StakingCredKind) {
	o.Kind = v
}

// GetPointer returns the Pointer field value if set, zero value otherwise.
func (o *StakingCredential) GetPointer() Pointer {
	if o == nil || IsNil(o.Pointer) {
		var ret Pointer
		return ret
	}
	return *o.Pointer
}

// GetPointerOk returns a tuple with the Pointer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakingCredential) GetPointerOk() (*Pointer, bool) {
	if o == nil || IsNil(o.Pointer) {
		return nil, false
	}
	return o.Pointer, true
}

// HasPointer returns a boolean if a field has been set.
func (o *StakingCredential) HasPointer() bool {
	if o != nil && !IsNil(o.Pointer) {
		return true
	}

	return false
}

// SetPointer gets a reference to the given Pointer and assigns it to the Pointer field.
func (o *StakingCredential) SetPointer(v Pointer) {
	o.Pointer = &v
}

// GetRewardAddress returns the RewardAddress field value if set, zero value otherwise.
func (o *StakingCredential) GetRewardAddress() string {
	if o == nil || IsNil(o.RewardAddress) {
		var ret string
		return ret
	}
	return *o.RewardAddress
}

// GetRewardAddressOk returns a tuple with the RewardAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakingCredential) GetRewardAddressOk() (*string, bool) {
	if o == nil || IsNil(o.RewardAddress) {
		return nil, false
	}
	return o.RewardAddress, true
}

// HasRewardAddress returns a boolean if a field has been set.
func (o *StakingCredential) HasRewardAddress() bool {
	if o != nil && !IsNil(o.RewardAddress) {
		return true
	}

	return false
}

// SetRewardAddress gets a reference to the given string and assigns it to the RewardAddress field.
func (o *StakingCredential) SetRewardAddress(v string) {
	o.RewardAddress = &v
}

func (o StakingCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StakingCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bech32) {
		toSerialize["bech32"] = o.Bech32
	}
	if !IsNil(o.Hex) {
		toSerialize["hex"] = o.Hex
	}
	toSerialize["kind"] = o.Kind
	if !IsNil(o.Pointer) {
		toSerialize["pointer"] = o.Pointer
	}
	if !IsNil(o.RewardAddress) {
		toSerialize["reward_address"] = o.RewardAddress
	}
	return toSerialize, nil
}

type NullableStakingCredential struct {
	value *StakingCredential
	isSet bool
}

func (v NullableStakingCredential) Get() *StakingCredential {
	return v.value
}

func (v *NullableStakingCredential) Set(val *StakingCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableStakingCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableStakingCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStakingCredential(val *StakingCredential) *NullableStakingCredential {
	return &NullableStakingCredential{value: val, isSet: true}
}

func (v NullableStakingCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStakingCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


