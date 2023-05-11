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

// checks if the Data type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Data{}

// Data struct for Data
type Data struct {
	Hash string `json:"hash"`
	Value map[string]interface{} `json:"value"`
}

// NewData instantiates a new Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewData(hash string, value map[string]interface{}) *Data {
	this := Data{}
	this.Hash = hash
	this.Value = value
	return &this
}

// NewDataWithDefaults instantiates a new Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataWithDefaults() *Data {
	this := Data{}
	return &this
}

// GetHash returns the Hash field value
func (o *Data) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *Data) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *Data) SetHash(v string) {
	o.Hash = v
}

// GetValue returns the Value field value
func (o *Data) GetValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Data) GetValueOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *Data) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o Data) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Data) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hash"] = o.Hash
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableData struct {
	value *Data
	isSet bool
}

func (v NullableData) Get() *Data {
	return v.value
}

func (v *NullableData) Set(val *Data) {
	v.value = val
	v.isSet = true
}

func (v NullableData) IsSet() bool {
	return v.isSet
}

func (v *NullableData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableData(val *Data) *NullableData {
	return &NullableData{value: val, isSet: true}
}

func (v NullableData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


