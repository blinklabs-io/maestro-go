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

// checks if the AddressInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressInfo{}

// AddressInfo struct for AddressInfo
type AddressInfo struct {
	Bech32 *string `json:"bech32,omitempty"`
	Hex string `json:"hex"`
	Network *NetworkId `json:"network,omitempty"`
	PaymentCred *PaymentCredential `json:"payment_cred,omitempty"`
	StakingCred *StakingCredential `json:"staking_cred,omitempty"`
}

// NewAddressInfo instantiates a new AddressInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressInfo(hex string) *AddressInfo {
	this := AddressInfo{}
	this.Hex = hex
	return &this
}

// NewAddressInfoWithDefaults instantiates a new AddressInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressInfoWithDefaults() *AddressInfo {
	this := AddressInfo{}
	return &this
}

// GetBech32 returns the Bech32 field value if set, zero value otherwise.
func (o *AddressInfo) GetBech32() string {
	if o == nil || IsNil(o.Bech32) {
		var ret string
		return ret
	}
	return *o.Bech32
}

// GetBech32Ok returns a tuple with the Bech32 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetBech32Ok() (*string, bool) {
	if o == nil || IsNil(o.Bech32) {
		return nil, false
	}
	return o.Bech32, true
}

// HasBech32 returns a boolean if a field has been set.
func (o *AddressInfo) HasBech32() bool {
	if o != nil && !IsNil(o.Bech32) {
		return true
	}

	return false
}

// SetBech32 gets a reference to the given string and assigns it to the Bech32 field.
func (o *AddressInfo) SetBech32(v string) {
	o.Bech32 = &v
}

// GetHex returns the Hex field value
func (o *AddressInfo) GetHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hex
}

// GetHexOk returns a tuple with the Hex field value
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hex, true
}

// SetHex sets field value
func (o *AddressInfo) SetHex(v string) {
	o.Hex = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *AddressInfo) GetNetwork() NetworkId {
	if o == nil || IsNil(o.Network) {
		var ret NetworkId
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetNetworkOk() (*NetworkId, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *AddressInfo) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given NetworkId and assigns it to the Network field.
func (o *AddressInfo) SetNetwork(v NetworkId) {
	o.Network = &v
}

// GetPaymentCred returns the PaymentCred field value if set, zero value otherwise.
func (o *AddressInfo) GetPaymentCred() PaymentCredential {
	if o == nil || IsNil(o.PaymentCred) {
		var ret PaymentCredential
		return ret
	}
	return *o.PaymentCred
}

// GetPaymentCredOk returns a tuple with the PaymentCred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetPaymentCredOk() (*PaymentCredential, bool) {
	if o == nil || IsNil(o.PaymentCred) {
		return nil, false
	}
	return o.PaymentCred, true
}

// HasPaymentCred returns a boolean if a field has been set.
func (o *AddressInfo) HasPaymentCred() bool {
	if o != nil && !IsNil(o.PaymentCred) {
		return true
	}

	return false
}

// SetPaymentCred gets a reference to the given PaymentCredential and assigns it to the PaymentCred field.
func (o *AddressInfo) SetPaymentCred(v PaymentCredential) {
	o.PaymentCred = &v
}

// GetStakingCred returns the StakingCred field value if set, zero value otherwise.
func (o *AddressInfo) GetStakingCred() StakingCredential {
	if o == nil || IsNil(o.StakingCred) {
		var ret StakingCredential
		return ret
	}
	return *o.StakingCred
}

// GetStakingCredOk returns a tuple with the StakingCred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetStakingCredOk() (*StakingCredential, bool) {
	if o == nil || IsNil(o.StakingCred) {
		return nil, false
	}
	return o.StakingCred, true
}

// HasStakingCred returns a boolean if a field has been set.
func (o *AddressInfo) HasStakingCred() bool {
	if o != nil && !IsNil(o.StakingCred) {
		return true
	}

	return false
}

// SetStakingCred gets a reference to the given StakingCredential and assigns it to the StakingCred field.
func (o *AddressInfo) SetStakingCred(v StakingCredential) {
	o.StakingCred = &v
}

func (o AddressInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bech32) {
		toSerialize["bech32"] = o.Bech32
	}
	toSerialize["hex"] = o.Hex
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.PaymentCred) {
		toSerialize["payment_cred"] = o.PaymentCred
	}
	if !IsNil(o.StakingCred) {
		toSerialize["staking_cred"] = o.StakingCred
	}
	return toSerialize, nil
}

type NullableAddressInfo struct {
	value *AddressInfo
	isSet bool
}

func (v NullableAddressInfo) Get() *AddressInfo {
	return v.value
}

func (v *NullableAddressInfo) Set(val *AddressInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressInfo(val *AddressInfo) *NullableAddressInfo {
	return &NullableAddressInfo{value: val, isSet: true}
}

func (v NullableAddressInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


