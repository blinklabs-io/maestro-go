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

// checks if the PoolBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolBlock{}

// PoolBlock Block created by a stake pool
type PoolBlock struct {
	// Absolute slot of the block
	AbsSlot *int64 `json:"abs_slot,omitempty"`
	// Block hash
	BlockHash string `json:"block_hash"`
	// Block height (block number)
	BlockHeight int32 `json:"block_height"`
	// UNIX timestamp when the block was mined
	BlockTime int32 `json:"block_time"`
	// Epoch number
	EpochNo *int32 `json:"epoch_no,omitempty"`
	// Epoch slot
	EpochSlot *int32 `json:"epoch_slot,omitempty"`
}

// NewPoolBlock instantiates a new PoolBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolBlock(blockHash string, blockHeight int32, blockTime int32) *PoolBlock {
	this := PoolBlock{}
	this.BlockHash = blockHash
	this.BlockHeight = blockHeight
	this.BlockTime = blockTime
	return &this
}

// NewPoolBlockWithDefaults instantiates a new PoolBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolBlockWithDefaults() *PoolBlock {
	this := PoolBlock{}
	return &this
}

// GetAbsSlot returns the AbsSlot field value if set, zero value otherwise.
func (o *PoolBlock) GetAbsSlot() int64 {
	if o == nil || IsNil(o.AbsSlot) {
		var ret int64
		return ret
	}
	return *o.AbsSlot
}

// GetAbsSlotOk returns a tuple with the AbsSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolBlock) GetAbsSlotOk() (*int64, bool) {
	if o == nil || IsNil(o.AbsSlot) {
		return nil, false
	}
	return o.AbsSlot, true
}

// HasAbsSlot returns a boolean if a field has been set.
func (o *PoolBlock) HasAbsSlot() bool {
	if o != nil && !IsNil(o.AbsSlot) {
		return true
	}

	return false
}

// SetAbsSlot gets a reference to the given int64 and assigns it to the AbsSlot field.
func (o *PoolBlock) SetAbsSlot(v int64) {
	o.AbsSlot = &v
}

// GetBlockHash returns the BlockHash field value
func (o *PoolBlock) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *PoolBlock) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *PoolBlock) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *PoolBlock) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *PoolBlock) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *PoolBlock) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetBlockTime returns the BlockTime field value
func (o *PoolBlock) GetBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value
// and a boolean to check if the value has been set.
func (o *PoolBlock) GetBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTime, true
}

// SetBlockTime sets field value
func (o *PoolBlock) SetBlockTime(v int32) {
	o.BlockTime = v
}

// GetEpochNo returns the EpochNo field value if set, zero value otherwise.
func (o *PoolBlock) GetEpochNo() int32 {
	if o == nil || IsNil(o.EpochNo) {
		var ret int32
		return ret
	}
	return *o.EpochNo
}

// GetEpochNoOk returns a tuple with the EpochNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolBlock) GetEpochNoOk() (*int32, bool) {
	if o == nil || IsNil(o.EpochNo) {
		return nil, false
	}
	return o.EpochNo, true
}

// HasEpochNo returns a boolean if a field has been set.
func (o *PoolBlock) HasEpochNo() bool {
	if o != nil && !IsNil(o.EpochNo) {
		return true
	}

	return false
}

// SetEpochNo gets a reference to the given int32 and assigns it to the EpochNo field.
func (o *PoolBlock) SetEpochNo(v int32) {
	o.EpochNo = &v
}

// GetEpochSlot returns the EpochSlot field value if set, zero value otherwise.
func (o *PoolBlock) GetEpochSlot() int32 {
	if o == nil || IsNil(o.EpochSlot) {
		var ret int32
		return ret
	}
	return *o.EpochSlot
}

// GetEpochSlotOk returns a tuple with the EpochSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolBlock) GetEpochSlotOk() (*int32, bool) {
	if o == nil || IsNil(o.EpochSlot) {
		return nil, false
	}
	return o.EpochSlot, true
}

// HasEpochSlot returns a boolean if a field has been set.
func (o *PoolBlock) HasEpochSlot() bool {
	if o != nil && !IsNil(o.EpochSlot) {
		return true
	}

	return false
}

// SetEpochSlot gets a reference to the given int32 and assigns it to the EpochSlot field.
func (o *PoolBlock) SetEpochSlot(v int32) {
	o.EpochSlot = &v
}

func (o PoolBlock) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbsSlot) {
		toSerialize["abs_slot"] = o.AbsSlot
	}
	toSerialize["block_hash"] = o.BlockHash
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["block_time"] = o.BlockTime
	if !IsNil(o.EpochNo) {
		toSerialize["epoch_no"] = o.EpochNo
	}
	if !IsNil(o.EpochSlot) {
		toSerialize["epoch_slot"] = o.EpochSlot
	}
	return toSerialize, nil
}

type NullablePoolBlock struct {
	value *PoolBlock
	isSet bool
}

func (v NullablePoolBlock) Get() *PoolBlock {
	return v.value
}

func (v *NullablePoolBlock) Set(val *PoolBlock) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolBlock) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolBlock(val *PoolBlock) *NullablePoolBlock {
	return &NullablePoolBlock{value: val, isSet: true}
}

func (v NullablePoolBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


