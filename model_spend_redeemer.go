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

// checks if the SpendRedeemer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpendRedeemer{}

// SpendRedeemer struct for SpendRedeemer
type SpendRedeemer struct {
	Data Datum `json:"data"`
	ExUnits int64 `json:"ex_units"`
	InputIndex int32 `json:"input_index"`
}

// NewSpendRedeemer instantiates a new SpendRedeemer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendRedeemer(data Datum, exUnits int64, inputIndex int32) *SpendRedeemer {
	this := SpendRedeemer{}
	this.Data = data
	this.ExUnits = exUnits
	this.InputIndex = inputIndex
	return &this
}

// NewSpendRedeemerWithDefaults instantiates a new SpendRedeemer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendRedeemerWithDefaults() *SpendRedeemer {
	this := SpendRedeemer{}
	return &this
}

// GetData returns the Data field value
func (o *SpendRedeemer) GetData() Datum {
	if o == nil {
		var ret Datum
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SpendRedeemer) GetDataOk() (*Datum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SpendRedeemer) SetData(v Datum) {
	o.Data = v
}

// GetExUnits returns the ExUnits field value
func (o *SpendRedeemer) GetExUnits() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExUnits
}

// GetExUnitsOk returns a tuple with the ExUnits field value
// and a boolean to check if the value has been set.
func (o *SpendRedeemer) GetExUnitsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExUnits, true
}

// SetExUnits sets field value
func (o *SpendRedeemer) SetExUnits(v int64) {
	o.ExUnits = v
}

// GetInputIndex returns the InputIndex field value
func (o *SpendRedeemer) GetInputIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InputIndex
}

// GetInputIndexOk returns a tuple with the InputIndex field value
// and a boolean to check if the value has been set.
func (o *SpendRedeemer) GetInputIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputIndex, true
}

// SetInputIndex sets field value
func (o *SpendRedeemer) SetInputIndex(v int32) {
	o.InputIndex = v
}

func (o SpendRedeemer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpendRedeemer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["ex_units"] = o.ExUnits
	toSerialize["input_index"] = o.InputIndex
	return toSerialize, nil
}

type NullableSpendRedeemer struct {
	value *SpendRedeemer
	isSet bool
}

func (v NullableSpendRedeemer) Get() *SpendRedeemer {
	return v.value
}

func (v *NullableSpendRedeemer) Set(val *SpendRedeemer) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendRedeemer) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendRedeemer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendRedeemer(val *SpendRedeemer) *NullableSpendRedeemer {
	return &NullableSpendRedeemer{value: val, isSet: true}
}

func (v NullableSpendRedeemer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendRedeemer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


