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

// checks if the Prices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Prices{}

// Prices struct for Prices
type Prices struct {
	Memory string `json:"memory"`
	Steps string `json:"steps"`
}

// NewPrices instantiates a new Prices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrices(memory string, steps string) *Prices {
	this := Prices{}
	this.Memory = memory
	this.Steps = steps
	return &this
}

// NewPricesWithDefaults instantiates a new Prices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricesWithDefaults() *Prices {
	this := Prices{}
	return &this
}

// GetMemory returns the Memory field value
func (o *Prices) GetMemory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *Prices) GetMemoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value
func (o *Prices) SetMemory(v string) {
	o.Memory = v
}

// GetSteps returns the Steps field value
func (o *Prices) GetSteps() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *Prices) GetStepsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Steps, true
}

// SetSteps sets field value
func (o *Prices) SetSteps(v string) {
	o.Steps = v
}

func (o Prices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Prices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["memory"] = o.Memory
	toSerialize["steps"] = o.Steps
	return toSerialize, nil
}

type NullablePrices struct {
	value *Prices
	isSet bool
}

func (v NullablePrices) Get() *Prices {
	return v.value
}

func (v *NullablePrices) Set(val *Prices) {
	v.value = val
	v.isSet = true
}

func (v NullablePrices) IsSet() bool {
	return v.isSet
}

func (v *NullablePrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrices(val *Prices) *NullablePrices {
	return &NullablePrices{value: val, isSet: true}
}

func (v NullablePrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


