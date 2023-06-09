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

// checks if the CurrentEpochInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrentEpochInfo{}

// CurrentEpochInfo Information summary of the current epoch
type CurrentEpochInfo struct {
	// Total blocks in the epoch so far
	BlkCount int32 `json:"blk_count"`
	// Epoch number
	EpochNo int32 `json:"epoch_no"`
	// Total fees collected in the epoch so far
	Fees string `json:"fees"`
	// UNIX timestamp when the epoch began
	StartTime int64 `json:"start_time"`
	// Total transactions in the epoch so far
	TxCount int32 `json:"tx_count"`
}

// NewCurrentEpochInfo instantiates a new CurrentEpochInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentEpochInfo(blkCount int32, epochNo int32, fees string, startTime int64, txCount int32) *CurrentEpochInfo {
	this := CurrentEpochInfo{}
	this.BlkCount = blkCount
	this.EpochNo = epochNo
	this.Fees = fees
	this.StartTime = startTime
	this.TxCount = txCount
	return &this
}

// NewCurrentEpochInfoWithDefaults instantiates a new CurrentEpochInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentEpochInfoWithDefaults() *CurrentEpochInfo {
	this := CurrentEpochInfo{}
	return &this
}

// GetBlkCount returns the BlkCount field value
func (o *CurrentEpochInfo) GetBlkCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlkCount
}

// GetBlkCountOk returns a tuple with the BlkCount field value
// and a boolean to check if the value has been set.
func (o *CurrentEpochInfo) GetBlkCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlkCount, true
}

// SetBlkCount sets field value
func (o *CurrentEpochInfo) SetBlkCount(v int32) {
	o.BlkCount = v
}

// GetEpochNo returns the EpochNo field value
func (o *CurrentEpochInfo) GetEpochNo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EpochNo
}

// GetEpochNoOk returns a tuple with the EpochNo field value
// and a boolean to check if the value has been set.
func (o *CurrentEpochInfo) GetEpochNoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EpochNo, true
}

// SetEpochNo sets field value
func (o *CurrentEpochInfo) SetEpochNo(v int32) {
	o.EpochNo = v
}

// GetFees returns the Fees field value
func (o *CurrentEpochInfo) GetFees() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value
// and a boolean to check if the value has been set.
func (o *CurrentEpochInfo) GetFeesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fees, true
}

// SetFees sets field value
func (o *CurrentEpochInfo) SetFees(v string) {
	o.Fees = v
}

// GetStartTime returns the StartTime field value
func (o *CurrentEpochInfo) GetStartTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *CurrentEpochInfo) GetStartTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *CurrentEpochInfo) SetStartTime(v int64) {
	o.StartTime = v
}

// GetTxCount returns the TxCount field value
func (o *CurrentEpochInfo) GetTxCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TxCount
}

// GetTxCountOk returns a tuple with the TxCount field value
// and a boolean to check if the value has been set.
func (o *CurrentEpochInfo) GetTxCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxCount, true
}

// SetTxCount sets field value
func (o *CurrentEpochInfo) SetTxCount(v int32) {
	o.TxCount = v
}

func (o CurrentEpochInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentEpochInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["blk_count"] = o.BlkCount
	toSerialize["epoch_no"] = o.EpochNo
	toSerialize["fees"] = o.Fees
	toSerialize["start_time"] = o.StartTime
	toSerialize["tx_count"] = o.TxCount
	return toSerialize, nil
}

type NullableCurrentEpochInfo struct {
	value *CurrentEpochInfo
	isSet bool
}

func (v NullableCurrentEpochInfo) Get() *CurrentEpochInfo {
	return v.value
}

func (v *NullableCurrentEpochInfo) Set(val *CurrentEpochInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentEpochInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentEpochInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentEpochInfo(val *CurrentEpochInfo) *NullableCurrentEpochInfo {
	return &NullableCurrentEpochInfo{value: val, isSet: true}
}

func (v NullableCurrentEpochInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentEpochInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


