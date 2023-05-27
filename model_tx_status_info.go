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

// checks if the TxStatusInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TxStatusInfo{}

// TxStatusInfo struct for TxStatusInfo
type TxStatusInfo struct {
	TxHash string `json:"tx_hash"`
	TxStatus TxStatus `json:"tx_status"`
}

// NewTxStatusInfo instantiates a new TxStatusInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxStatusInfo(txHash string, txStatus TxStatus) *TxStatusInfo {
	this := TxStatusInfo{}
	this.TxHash = txHash
	this.TxStatus = txStatus
	return &this
}

// NewTxStatusInfoWithDefaults instantiates a new TxStatusInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxStatusInfoWithDefaults() *TxStatusInfo {
	this := TxStatusInfo{}
	return &this
}

// GetTxHash returns the TxHash field value
func (o *TxStatusInfo) GetTxHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value
// and a boolean to check if the value has been set.
func (o *TxStatusInfo) GetTxHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxHash, true
}

// SetTxHash sets field value
func (o *TxStatusInfo) SetTxHash(v string) {
	o.TxHash = v
}

// GetTxStatus returns the TxStatus field value
func (o *TxStatusInfo) GetTxStatus() TxStatus {
	if o == nil {
		var ret TxStatus
		return ret
	}

	return o.TxStatus
}

// GetTxStatusOk returns a tuple with the TxStatus field value
// and a boolean to check if the value has been set.
func (o *TxStatusInfo) GetTxStatusOk() (*TxStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxStatus, true
}

// SetTxStatus sets field value
func (o *TxStatusInfo) SetTxStatus(v TxStatus) {
	o.TxStatus = v
}

func (o TxStatusInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TxStatusInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx_hash"] = o.TxHash
	toSerialize["tx_status"] = o.TxStatus
	return toSerialize, nil
}

type NullableTxStatusInfo struct {
	value *TxStatusInfo
	isSet bool
}

func (v NullableTxStatusInfo) Get() *TxStatusInfo {
	return v.value
}

func (v *NullableTxStatusInfo) Set(val *TxStatusInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTxStatusInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTxStatusInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxStatusInfo(val *TxStatusInfo) *NullableTxStatusInfo {
	return &NullableTxStatusInfo{value: val, isSet: true}
}

func (v NullableTxStatusInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxStatusInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


