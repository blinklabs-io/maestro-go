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

// checks if the MintingTx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MintingTx{}

// MintingTx Transaction which minted or burned a specific asset
type MintingTx struct {
	// UNIX timestamp of the block which included transaction
	BlockTimestamp int32 `json:"block_timestamp"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Amount of the asset minted or burned (negative if burned)
	MintAmount int64 `json:"mint_amount"`
	// Transaction hash
	TxHash string `json:"tx_hash"`
}

// NewMintingTx instantiates a new MintingTx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMintingTx(blockTimestamp int32, mintAmount int64, txHash string) *MintingTx {
	this := MintingTx{}
	this.BlockTimestamp = blockTimestamp
	this.MintAmount = mintAmount
	this.TxHash = txHash
	return &this
}

// NewMintingTxWithDefaults instantiates a new MintingTx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMintingTxWithDefaults() *MintingTx {
	this := MintingTx{}
	return &this
}

// GetBlockTimestamp returns the BlockTimestamp field value
func (o *MintingTx) GetBlockTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockTimestamp
}

// GetBlockTimestampOk returns a tuple with the BlockTimestamp field value
// and a boolean to check if the value has been set.
func (o *MintingTx) GetBlockTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTimestamp, true
}

// SetBlockTimestamp sets field value
func (o *MintingTx) SetBlockTimestamp(v int32) {
	o.BlockTimestamp = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *MintingTx) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MintingTx) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *MintingTx) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *MintingTx) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMintAmount returns the MintAmount field value
func (o *MintingTx) GetMintAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MintAmount
}

// GetMintAmountOk returns a tuple with the MintAmount field value
// and a boolean to check if the value has been set.
func (o *MintingTx) GetMintAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MintAmount, true
}

// SetMintAmount sets field value
func (o *MintingTx) SetMintAmount(v int64) {
	o.MintAmount = v
}

// GetTxHash returns the TxHash field value
func (o *MintingTx) GetTxHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value
// and a boolean to check if the value has been set.
func (o *MintingTx) GetTxHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxHash, true
}

// SetTxHash sets field value
func (o *MintingTx) SetTxHash(v string) {
	o.TxHash = v
}

func (o MintingTx) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MintingTx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block_timestamp"] = o.BlockTimestamp
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["mint_amount"] = o.MintAmount
	toSerialize["tx_hash"] = o.TxHash
	return toSerialize, nil
}

type NullableMintingTx struct {
	value *MintingTx
	isSet bool
}

func (v NullableMintingTx) Get() *MintingTx {
	return v.value
}

func (v *NullableMintingTx) Set(val *MintingTx) {
	v.value = val
	v.isSet = true
}

func (v NullableMintingTx) IsSet() bool {
	return v.isSet
}

func (v *NullableMintingTx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMintingTx(val *MintingTx) *NullableMintingTx {
	return &NullableMintingTx{value: val, isSet: true}
}

func (v NullableMintingTx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMintingTx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

