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

// checks if the PoolListInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolListInfo{}

// PoolListInfo Stake pool identifier
type PoolListInfo struct {
	// Bech32 encoded pool ID
	PoolIdBech32 string `json:"pool_id_bech32"`
	// Pool ticker symbol
	Ticker *string `json:"ticker,omitempty"`
}

// NewPoolListInfo instantiates a new PoolListInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolListInfo(poolIdBech32 string) *PoolListInfo {
	this := PoolListInfo{}
	this.PoolIdBech32 = poolIdBech32
	return &this
}

// NewPoolListInfoWithDefaults instantiates a new PoolListInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolListInfoWithDefaults() *PoolListInfo {
	this := PoolListInfo{}
	return &this
}

// GetPoolIdBech32 returns the PoolIdBech32 field value
func (o *PoolListInfo) GetPoolIdBech32() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoolIdBech32
}

// GetPoolIdBech32Ok returns a tuple with the PoolIdBech32 field value
// and a boolean to check if the value has been set.
func (o *PoolListInfo) GetPoolIdBech32Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolIdBech32, true
}

// SetPoolIdBech32 sets field value
func (o *PoolListInfo) SetPoolIdBech32(v string) {
	o.PoolIdBech32 = v
}

// GetTicker returns the Ticker field value if set, zero value otherwise.
func (o *PoolListInfo) GetTicker() string {
	if o == nil || IsNil(o.Ticker) {
		var ret string
		return ret
	}
	return *o.Ticker
}

// GetTickerOk returns a tuple with the Ticker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolListInfo) GetTickerOk() (*string, bool) {
	if o == nil || IsNil(o.Ticker) {
		return nil, false
	}
	return o.Ticker, true
}

// HasTicker returns a boolean if a field has been set.
func (o *PoolListInfo) HasTicker() bool {
	if o != nil && !IsNil(o.Ticker) {
		return true
	}

	return false
}

// SetTicker gets a reference to the given string and assigns it to the Ticker field.
func (o *PoolListInfo) SetTicker(v string) {
	o.Ticker = &v
}

func (o PoolListInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolListInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pool_id_bech32"] = o.PoolIdBech32
	if !IsNil(o.Ticker) {
		toSerialize["ticker"] = o.Ticker
	}
	return toSerialize, nil
}

type NullablePoolListInfo struct {
	value *PoolListInfo
	isSet bool
}

func (v NullablePoolListInfo) Get() *PoolListInfo {
	return v.value
}

func (v *NullablePoolListInfo) Set(val *PoolListInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolListInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolListInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolListInfo(val *PoolListInfo) *NullablePoolListInfo {
	return &NullablePoolListInfo{value: val, isSet: true}
}

func (v NullablePoolListInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolListInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


