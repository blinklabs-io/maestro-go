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

// checks if the UtxoWithBytes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtxoWithBytes{}

// UtxoWithBytes Transaction output
type UtxoWithBytes struct {
	// Address which controls the UTxO
	Address string `json:"address"`
	Assets []Asset `json:"assets"`
	Datum *DatumOption `json:"datum,omitempty"`
	// UTxO transaction index
	Index int32 `json:"index"`
	ReferenceScript *Script `json:"reference_script,omitempty"`
	// UTxO transaction hash
	TxHash string `json:"tx_hash"`
	// Hex encoded transaction output CBOR bytes
	TxoutCbor *string `json:"txout_cbor,omitempty"`
}

// NewUtxoWithBytes instantiates a new UtxoWithBytes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtxoWithBytes(address string, assets []Asset, index int32, txHash string) *UtxoWithBytes {
	this := UtxoWithBytes{}
	this.Address = address
	this.Assets = assets
	this.Index = index
	this.TxHash = txHash
	return &this
}

// NewUtxoWithBytesWithDefaults instantiates a new UtxoWithBytes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtxoWithBytesWithDefaults() *UtxoWithBytes {
	this := UtxoWithBytes{}
	return &this
}

// GetAddress returns the Address field value
func (o *UtxoWithBytes) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *UtxoWithBytes) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *UtxoWithBytes) SetAddress(v string) {
	o.Address = v
}

// GetAssets returns the Assets field value
func (o *UtxoWithBytes) GetAssets() []Asset {
	if o == nil {
		var ret []Asset
		return ret
	}

	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value
// and a boolean to check if the value has been set.
func (o *UtxoWithBytes) GetAssetsOk() ([]Asset, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assets, true
}

// SetAssets sets field value
func (o *UtxoWithBytes) SetAssets(v []Asset) {
	o.Assets = v
}

// GetDatum returns the Datum field value if set, zero value otherwise.
func (o *UtxoWithBytes) GetDatum() DatumOption {
	if o == nil || IsNil(o.Datum) {
		var ret DatumOption
		return ret
	}
	return *o.Datum
}

// GetDatumOk returns a tuple with the Datum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtxoWithBytes) GetDatumOk() (*DatumOption, bool) {
	if o == nil || IsNil(o.Datum) {
		return nil, false
	}
	return o.Datum, true
}

// HasDatum returns a boolean if a field has been set.
func (o *UtxoWithBytes) HasDatum() bool {
	if o != nil && !IsNil(o.Datum) {
		return true
	}

	return false
}

// SetDatum gets a reference to the given DatumOption and assigns it to the Datum field.
func (o *UtxoWithBytes) SetDatum(v DatumOption) {
	o.Datum = &v
}

// GetIndex returns the Index field value
func (o *UtxoWithBytes) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *UtxoWithBytes) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *UtxoWithBytes) SetIndex(v int32) {
	o.Index = v
}

// GetReferenceScript returns the ReferenceScript field value if set, zero value otherwise.
func (o *UtxoWithBytes) GetReferenceScript() Script {
	if o == nil || IsNil(o.ReferenceScript) {
		var ret Script
		return ret
	}
	return *o.ReferenceScript
}

// GetReferenceScriptOk returns a tuple with the ReferenceScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtxoWithBytes) GetReferenceScriptOk() (*Script, bool) {
	if o == nil || IsNil(o.ReferenceScript) {
		return nil, false
	}
	return o.ReferenceScript, true
}

// HasReferenceScript returns a boolean if a field has been set.
func (o *UtxoWithBytes) HasReferenceScript() bool {
	if o != nil && !IsNil(o.ReferenceScript) {
		return true
	}

	return false
}

// SetReferenceScript gets a reference to the given Script and assigns it to the ReferenceScript field.
func (o *UtxoWithBytes) SetReferenceScript(v Script) {
	o.ReferenceScript = &v
}

// GetTxHash returns the TxHash field value
func (o *UtxoWithBytes) GetTxHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value
// and a boolean to check if the value has been set.
func (o *UtxoWithBytes) GetTxHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxHash, true
}

// SetTxHash sets field value
func (o *UtxoWithBytes) SetTxHash(v string) {
	o.TxHash = v
}

// GetTxoutCbor returns the TxoutCbor field value if set, zero value otherwise.
func (o *UtxoWithBytes) GetTxoutCbor() string {
	if o == nil || IsNil(o.TxoutCbor) {
		var ret string
		return ret
	}
	return *o.TxoutCbor
}

// GetTxoutCborOk returns a tuple with the TxoutCbor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtxoWithBytes) GetTxoutCborOk() (*string, bool) {
	if o == nil || IsNil(o.TxoutCbor) {
		return nil, false
	}
	return o.TxoutCbor, true
}

// HasTxoutCbor returns a boolean if a field has been set.
func (o *UtxoWithBytes) HasTxoutCbor() bool {
	if o != nil && !IsNil(o.TxoutCbor) {
		return true
	}

	return false
}

// SetTxoutCbor gets a reference to the given string and assigns it to the TxoutCbor field.
func (o *UtxoWithBytes) SetTxoutCbor(v string) {
	o.TxoutCbor = &v
}

func (o UtxoWithBytes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtxoWithBytes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["assets"] = o.Assets
	if !IsNil(o.Datum) {
		toSerialize["datum"] = o.Datum
	}
	toSerialize["index"] = o.Index
	if !IsNil(o.ReferenceScript) {
		toSerialize["reference_script"] = o.ReferenceScript
	}
	toSerialize["tx_hash"] = o.TxHash
	if !IsNil(o.TxoutCbor) {
		toSerialize["txout_cbor"] = o.TxoutCbor
	}
	return toSerialize, nil
}

type NullableUtxoWithBytes struct {
	value *UtxoWithBytes
	isSet bool
}

func (v NullableUtxoWithBytes) Get() *UtxoWithBytes {
	return v.value
}

func (v *NullableUtxoWithBytes) Set(val *UtxoWithBytes) {
	v.value = val
	v.isSet = true
}

func (v NullableUtxoWithBytes) IsSet() bool {
	return v.isSet
}

func (v *NullableUtxoWithBytes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtxoWithBytes(val *UtxoWithBytes) *NullableUtxoWithBytes {
	return &NullableUtxoWithBytes{value: val, isSet: true}
}

func (v NullableUtxoWithBytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtxoWithBytes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


