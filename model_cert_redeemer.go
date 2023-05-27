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

// checks if the CertRedeemer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertRedeemer{}

// CertRedeemer struct for CertRedeemer
type CertRedeemer struct {
	CertIndex int32 `json:"cert_index"`
	Data Datum `json:"data"`
	ExUnits int64 `json:"ex_units"`
}

// NewCertRedeemer instantiates a new CertRedeemer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertRedeemer(certIndex int32, data Datum, exUnits int64) *CertRedeemer {
	this := CertRedeemer{}
	this.CertIndex = certIndex
	this.Data = data
	this.ExUnits = exUnits
	return &this
}

// NewCertRedeemerWithDefaults instantiates a new CertRedeemer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertRedeemerWithDefaults() *CertRedeemer {
	this := CertRedeemer{}
	return &this
}

// GetCertIndex returns the CertIndex field value
func (o *CertRedeemer) GetCertIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CertIndex
}

// GetCertIndexOk returns a tuple with the CertIndex field value
// and a boolean to check if the value has been set.
func (o *CertRedeemer) GetCertIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertIndex, true
}

// SetCertIndex sets field value
func (o *CertRedeemer) SetCertIndex(v int32) {
	o.CertIndex = v
}

// GetData returns the Data field value
func (o *CertRedeemer) GetData() Datum {
	if o == nil {
		var ret Datum
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CertRedeemer) GetDataOk() (*Datum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CertRedeemer) SetData(v Datum) {
	o.Data = v
}

// GetExUnits returns the ExUnits field value
func (o *CertRedeemer) GetExUnits() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExUnits
}

// GetExUnitsOk returns a tuple with the ExUnits field value
// and a boolean to check if the value has been set.
func (o *CertRedeemer) GetExUnitsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExUnits, true
}

// SetExUnits sets field value
func (o *CertRedeemer) SetExUnits(v int64) {
	o.ExUnits = v
}

func (o CertRedeemer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertRedeemer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cert_index"] = o.CertIndex
	toSerialize["data"] = o.Data
	toSerialize["ex_units"] = o.ExUnits
	return toSerialize, nil
}

type NullableCertRedeemer struct {
	value *CertRedeemer
	isSet bool
}

func (v NullableCertRedeemer) Get() *CertRedeemer {
	return v.value
}

func (v *NullableCertRedeemer) Set(val *CertRedeemer) {
	v.value = val
	v.isSet = true
}

func (v NullableCertRedeemer) IsSet() bool {
	return v.isSet
}

func (v *NullableCertRedeemer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertRedeemer(val *CertRedeemer) *NullableCertRedeemer {
	return &NullableCertRedeemer{value: val, isSet: true}
}

func (v NullableCertRedeemer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertRedeemer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


