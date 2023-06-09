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

// checks if the MirCert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MirCert{}

// MirCert Certificate for sending an instantaneous reward
type MirCert struct {
	// Index of the certificate in the transaction
	CertIndex int32 `json:"cert_index"`
	From MirSource `json:"from"`
	// Where the rewards funds are being sent
	To string `json:"to"`
}

// NewMirCert instantiates a new MirCert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMirCert(certIndex int32, from MirSource, to string) *MirCert {
	this := MirCert{}
	this.CertIndex = certIndex
	this.From = from
	this.To = to
	return &this
}

// NewMirCertWithDefaults instantiates a new MirCert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMirCertWithDefaults() *MirCert {
	this := MirCert{}
	return &this
}

// GetCertIndex returns the CertIndex field value
func (o *MirCert) GetCertIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CertIndex
}

// GetCertIndexOk returns a tuple with the CertIndex field value
// and a boolean to check if the value has been set.
func (o *MirCert) GetCertIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertIndex, true
}

// SetCertIndex sets field value
func (o *MirCert) SetCertIndex(v int32) {
	o.CertIndex = v
}

// GetFrom returns the From field value
func (o *MirCert) GetFrom() MirSource {
	if o == nil {
		var ret MirSource
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *MirCert) GetFromOk() (*MirSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *MirCert) SetFrom(v MirSource) {
	o.From = v
}

// GetTo returns the To field value
func (o *MirCert) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *MirCert) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *MirCert) SetTo(v string) {
	o.To = v
}

func (o MirCert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MirCert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cert_index"] = o.CertIndex
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	return toSerialize, nil
}

type NullableMirCert struct {
	value *MirCert
	isSet bool
}

func (v NullableMirCert) Get() *MirCert {
	return v.value
}

func (v *NullableMirCert) Set(val *MirCert) {
	v.value = val
	v.isSet = true
}

func (v NullableMirCert) IsSet() bool {
	return v.isSet
}

func (v *NullableMirCert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMirCert(val *MirCert) *NullableMirCert {
	return &NullableMirCert{value: val, isSet: true}
}

func (v NullableMirCert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMirCert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


