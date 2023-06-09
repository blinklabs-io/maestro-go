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

// checks if the PaymentCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentCredential{}

// PaymentCredential struct for PaymentCredential
type PaymentCredential struct {
	Bech32 string `json:"bech32"`
	Hex string `json:"hex"`
	Kind PaymentCredKind `json:"kind"`
}

// NewPaymentCredential instantiates a new PaymentCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentCredential(bech32 string, hex string, kind PaymentCredKind) *PaymentCredential {
	this := PaymentCredential{}
	this.Bech32 = bech32
	this.Hex = hex
	this.Kind = kind
	return &this
}

// NewPaymentCredentialWithDefaults instantiates a new PaymentCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentCredentialWithDefaults() *PaymentCredential {
	this := PaymentCredential{}
	return &this
}

// GetBech32 returns the Bech32 field value
func (o *PaymentCredential) GetBech32() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bech32
}

// GetBech32Ok returns a tuple with the Bech32 field value
// and a boolean to check if the value has been set.
func (o *PaymentCredential) GetBech32Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bech32, true
}

// SetBech32 sets field value
func (o *PaymentCredential) SetBech32(v string) {
	o.Bech32 = v
}

// GetHex returns the Hex field value
func (o *PaymentCredential) GetHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hex
}

// GetHexOk returns a tuple with the Hex field value
// and a boolean to check if the value has been set.
func (o *PaymentCredential) GetHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hex, true
}

// SetHex sets field value
func (o *PaymentCredential) SetHex(v string) {
	o.Hex = v
}

// GetKind returns the Kind field value
func (o *PaymentCredential) GetKind() PaymentCredKind {
	if o == nil {
		var ret PaymentCredKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *PaymentCredential) GetKindOk() (*PaymentCredKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *PaymentCredential) SetKind(v PaymentCredKind) {
	o.Kind = v
}

func (o PaymentCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bech32"] = o.Bech32
	toSerialize["hex"] = o.Hex
	toSerialize["kind"] = o.Kind
	return toSerialize, nil
}

type NullablePaymentCredential struct {
	value *PaymentCredential
	isSet bool
}

func (v NullablePaymentCredential) Get() *PaymentCredential {
	return v.value
}

func (v *NullablePaymentCredential) Set(val *PaymentCredential) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentCredential) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentCredential(val *PaymentCredential) *NullablePaymentCredential {
	return &NullablePaymentCredential{value: val, isSet: true}
}

func (v NullablePaymentCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


