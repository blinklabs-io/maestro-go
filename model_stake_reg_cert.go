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

// checks if the StakeRegCert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StakeRegCert{}

// StakeRegCert Certificate for registering a stake key
type StakeRegCert struct {
	// Index of the certificate in the transaction
	CertIndex int32 `json:"cert_index"`
	// Stake address corresponding to stake key being updated
	StakeAddress string `json:"stake_address"`
}

// NewStakeRegCert instantiates a new StakeRegCert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStakeRegCert(certIndex int32, stakeAddress string) *StakeRegCert {
	this := StakeRegCert{}
	this.CertIndex = certIndex
	this.StakeAddress = stakeAddress
	return &this
}

// NewStakeRegCertWithDefaults instantiates a new StakeRegCert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStakeRegCertWithDefaults() *StakeRegCert {
	this := StakeRegCert{}
	return &this
}

// GetCertIndex returns the CertIndex field value
func (o *StakeRegCert) GetCertIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CertIndex
}

// GetCertIndexOk returns a tuple with the CertIndex field value
// and a boolean to check if the value has been set.
func (o *StakeRegCert) GetCertIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertIndex, true
}

// SetCertIndex sets field value
func (o *StakeRegCert) SetCertIndex(v int32) {
	o.CertIndex = v
}

// GetStakeAddress returns the StakeAddress field value
func (o *StakeRegCert) GetStakeAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StakeAddress
}

// GetStakeAddressOk returns a tuple with the StakeAddress field value
// and a boolean to check if the value has been set.
func (o *StakeRegCert) GetStakeAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakeAddress, true
}

// SetStakeAddress sets field value
func (o *StakeRegCert) SetStakeAddress(v string) {
	o.StakeAddress = v
}

func (o StakeRegCert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StakeRegCert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cert_index"] = o.CertIndex
	toSerialize["stake_address"] = o.StakeAddress
	return toSerialize, nil
}

type NullableStakeRegCert struct {
	value *StakeRegCert
	isSet bool
}

func (v NullableStakeRegCert) Get() *StakeRegCert {
	return v.value
}

func (v *NullableStakeRegCert) Set(val *StakeRegCert) {
	v.value = val
	v.isSet = true
}

func (v NullableStakeRegCert) IsSet() bool {
	return v.isSet
}

func (v *NullableStakeRegCert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStakeRegCert(val *StakeRegCert) *NullableStakeRegCert {
	return &NullableStakeRegCert{value: val, isSet: true}
}

func (v NullableStakeRegCert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStakeRegCert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


