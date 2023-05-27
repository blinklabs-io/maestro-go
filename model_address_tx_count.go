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

// checks if the AddressTxCount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressTxCount{}

// AddressTxCount struct for AddressTxCount
type AddressTxCount struct {
	// Number of transactions
	Count int64 `json:"count"`
}

// NewAddressTxCount instantiates a new AddressTxCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTxCount(count int64) *AddressTxCount {
	this := AddressTxCount{}
	this.Count = count
	return &this
}

// NewAddressTxCountWithDefaults instantiates a new AddressTxCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTxCountWithDefaults() *AddressTxCount {
	this := AddressTxCount{}
	return &this
}

// GetCount returns the Count field value
func (o *AddressTxCount) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *AddressTxCount) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *AddressTxCount) SetCount(v int64) {
	o.Count = v
}

func (o AddressTxCount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTxCount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	return toSerialize, nil
}

type NullableAddressTxCount struct {
	value *AddressTxCount
	isSet bool
}

func (v NullableAddressTxCount) Get() *AddressTxCount {
	return v.value
}

func (v *NullableAddressTxCount) Set(val *AddressTxCount) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTxCount) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTxCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTxCount(val *AddressTxCount) *NullableAddressTxCount {
	return &NullableAddressTxCount{value: val, isSet: true}
}

func (v NullableAddressTxCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTxCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


