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

// checks if the AssetHolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetHolder{}

// AssetHolder Holder of a specific asset
type AssetHolder struct {
	// Address of the holder
	Address string `json:"address"`
	// Amount of the asset owned by the holder
	Amount int64 `json:"amount"`
}

// NewAssetHolder instantiates a new AssetHolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetHolder(address string, amount int64) *AssetHolder {
	this := AssetHolder{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewAssetHolderWithDefaults instantiates a new AssetHolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetHolderWithDefaults() *AssetHolder {
	this := AssetHolder{}
	return &this
}

// GetAddress returns the Address field value
func (o *AssetHolder) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AssetHolder) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AssetHolder) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *AssetHolder) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AssetHolder) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AssetHolder) SetAmount(v int64) {
	o.Amount = v
}

func (o AssetHolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetHolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

type NullableAssetHolder struct {
	value *AssetHolder
	isSet bool
}

func (v NullableAssetHolder) Get() *AssetHolder {
	return v.value
}

func (v *NullableAssetHolder) Set(val *AssetHolder) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetHolder) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetHolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetHolder(val *AssetHolder) *NullableAssetHolder {
	return &NullableAssetHolder{value: val, isSet: true}
}

func (v NullableAssetHolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetHolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

