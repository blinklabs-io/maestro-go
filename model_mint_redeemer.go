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

// checks if the MintRedeemer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MintRedeemer{}

// MintRedeemer struct for MintRedeemer
type MintRedeemer struct {
	Data Datum `json:"data"`
	ExUnits int64 `json:"ex_units"`
	Policy string `json:"policy"`
}

// NewMintRedeemer instantiates a new MintRedeemer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMintRedeemer(data Datum, exUnits int64, policy string) *MintRedeemer {
	this := MintRedeemer{}
	this.Data = data
	this.ExUnits = exUnits
	this.Policy = policy
	return &this
}

// NewMintRedeemerWithDefaults instantiates a new MintRedeemer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMintRedeemerWithDefaults() *MintRedeemer {
	this := MintRedeemer{}
	return &this
}

// GetData returns the Data field value
func (o *MintRedeemer) GetData() Datum {
	if o == nil {
		var ret Datum
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MintRedeemer) GetDataOk() (*Datum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MintRedeemer) SetData(v Datum) {
	o.Data = v
}

// GetExUnits returns the ExUnits field value
func (o *MintRedeemer) GetExUnits() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExUnits
}

// GetExUnitsOk returns a tuple with the ExUnits field value
// and a boolean to check if the value has been set.
func (o *MintRedeemer) GetExUnitsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExUnits, true
}

// SetExUnits sets field value
func (o *MintRedeemer) SetExUnits(v int64) {
	o.ExUnits = v
}

// GetPolicy returns the Policy field value
func (o *MintRedeemer) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *MintRedeemer) GetPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *MintRedeemer) SetPolicy(v string) {
	o.Policy = v
}

func (o MintRedeemer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MintRedeemer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["ex_units"] = o.ExUnits
	toSerialize["policy"] = o.Policy
	return toSerialize, nil
}

type NullableMintRedeemer struct {
	value *MintRedeemer
	isSet bool
}

func (v NullableMintRedeemer) Get() *MintRedeemer {
	return v.value
}

func (v *NullableMintRedeemer) Set(val *MintRedeemer) {
	v.value = val
	v.isSet = true
}

func (v NullableMintRedeemer) IsSet() bool {
	return v.isSet
}

func (v *NullableMintRedeemer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMintRedeemer(val *MintRedeemer) *NullableMintRedeemer {
	return &NullableMintRedeemer{value: val, isSet: true}
}

func (v NullableMintRedeemer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMintRedeemer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


