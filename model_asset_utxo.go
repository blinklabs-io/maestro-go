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

// checks if the AssetUtxo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetUtxo{}

// AssetUtxo UTxO which contains a specific asset
type AssetUtxo struct {
	// Address which controls the UTxO
	Address string `json:"address"`
	// Amount of the asset contained in the UTxO
	Amount int64 `json:"amount"`
	// UTxO transaction index
	Index int32 `json:"index"`
	// UTxO transaction hash
	TxHash string `json:"tx_hash"`
}

// NewAssetUtxo instantiates a new AssetUtxo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetUtxo(address string, amount int64, index int32, txHash string) *AssetUtxo {
	this := AssetUtxo{}
	this.Address = address
	this.Amount = amount
	this.Index = index
	this.TxHash = txHash
	return &this
}

// NewAssetUtxoWithDefaults instantiates a new AssetUtxo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetUtxoWithDefaults() *AssetUtxo {
	this := AssetUtxo{}
	return &this
}

// GetAddress returns the Address field value
func (o *AssetUtxo) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AssetUtxo) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AssetUtxo) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *AssetUtxo) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AssetUtxo) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AssetUtxo) SetAmount(v int64) {
	o.Amount = v
}

// GetIndex returns the Index field value
func (o *AssetUtxo) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *AssetUtxo) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *AssetUtxo) SetIndex(v int32) {
	o.Index = v
}

// GetTxHash returns the TxHash field value
func (o *AssetUtxo) GetTxHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value
// and a boolean to check if the value has been set.
func (o *AssetUtxo) GetTxHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxHash, true
}

// SetTxHash sets field value
func (o *AssetUtxo) SetTxHash(v string) {
	o.TxHash = v
}

func (o AssetUtxo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetUtxo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["amount"] = o.Amount
	toSerialize["index"] = o.Index
	toSerialize["tx_hash"] = o.TxHash
	return toSerialize, nil
}

type NullableAssetUtxo struct {
	value *AssetUtxo
	isSet bool
}

func (v NullableAssetUtxo) Get() *AssetUtxo {
	return v.value
}

func (v *NullableAssetUtxo) Set(val *AssetUtxo) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetUtxo) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetUtxo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetUtxo(val *AssetUtxo) *NullableAssetUtxo {
	return &NullableAssetUtxo{value: val, isSet: true}
}

func (v NullableAssetUtxo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetUtxo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


