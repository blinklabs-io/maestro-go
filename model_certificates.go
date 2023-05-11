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

// checks if the Certificates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Certificates{}

// Certificates Certificates found in a transaction
type Certificates struct {
	MirTransfers []MirCert `json:"mir_transfers"`
	PoolRegistrations []PoolRegCert `json:"pool_registrations"`
	PoolRetirements []PoolRetireCert `json:"pool_retirements"`
	StakeDelegations []StakeDelegCert `json:"stake_delegations"`
	StakeDeregistrations []StakeRegCert `json:"stake_deregistrations"`
	StakeRegistrations []StakeRegCert `json:"stake_registrations"`
}

// NewCertificates instantiates a new Certificates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificates(mirTransfers []MirCert, poolRegistrations []PoolRegCert, poolRetirements []PoolRetireCert, stakeDelegations []StakeDelegCert, stakeDeregistrations []StakeRegCert, stakeRegistrations []StakeRegCert) *Certificates {
	this := Certificates{}
	this.MirTransfers = mirTransfers
	this.PoolRegistrations = poolRegistrations
	this.PoolRetirements = poolRetirements
	this.StakeDelegations = stakeDelegations
	this.StakeDeregistrations = stakeDeregistrations
	this.StakeRegistrations = stakeRegistrations
	return &this
}

// NewCertificatesWithDefaults instantiates a new Certificates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificatesWithDefaults() *Certificates {
	this := Certificates{}
	return &this
}

// GetMirTransfers returns the MirTransfers field value
func (o *Certificates) GetMirTransfers() []MirCert {
	if o == nil {
		var ret []MirCert
		return ret
	}

	return o.MirTransfers
}

// GetMirTransfersOk returns a tuple with the MirTransfers field value
// and a boolean to check if the value has been set.
func (o *Certificates) GetMirTransfersOk() ([]MirCert, bool) {
	if o == nil {
		return nil, false
	}
	return o.MirTransfers, true
}

// SetMirTransfers sets field value
func (o *Certificates) SetMirTransfers(v []MirCert) {
	o.MirTransfers = v
}

// GetPoolRegistrations returns the PoolRegistrations field value
func (o *Certificates) GetPoolRegistrations() []PoolRegCert {
	if o == nil {
		var ret []PoolRegCert
		return ret
	}

	return o.PoolRegistrations
}

// GetPoolRegistrationsOk returns a tuple with the PoolRegistrations field value
// and a boolean to check if the value has been set.
func (o *Certificates) GetPoolRegistrationsOk() ([]PoolRegCert, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoolRegistrations, true
}

// SetPoolRegistrations sets field value
func (o *Certificates) SetPoolRegistrations(v []PoolRegCert) {
	o.PoolRegistrations = v
}

// GetPoolRetirements returns the PoolRetirements field value
func (o *Certificates) GetPoolRetirements() []PoolRetireCert {
	if o == nil {
		var ret []PoolRetireCert
		return ret
	}

	return o.PoolRetirements
}

// GetPoolRetirementsOk returns a tuple with the PoolRetirements field value
// and a boolean to check if the value has been set.
func (o *Certificates) GetPoolRetirementsOk() ([]PoolRetireCert, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoolRetirements, true
}

// SetPoolRetirements sets field value
func (o *Certificates) SetPoolRetirements(v []PoolRetireCert) {
	o.PoolRetirements = v
}

// GetStakeDelegations returns the StakeDelegations field value
func (o *Certificates) GetStakeDelegations() []StakeDelegCert {
	if o == nil {
		var ret []StakeDelegCert
		return ret
	}

	return o.StakeDelegations
}

// GetStakeDelegationsOk returns a tuple with the StakeDelegations field value
// and a boolean to check if the value has been set.
func (o *Certificates) GetStakeDelegationsOk() ([]StakeDelegCert, bool) {
	if o == nil {
		return nil, false
	}
	return o.StakeDelegations, true
}

// SetStakeDelegations sets field value
func (o *Certificates) SetStakeDelegations(v []StakeDelegCert) {
	o.StakeDelegations = v
}

// GetStakeDeregistrations returns the StakeDeregistrations field value
func (o *Certificates) GetStakeDeregistrations() []StakeRegCert {
	if o == nil {
		var ret []StakeRegCert
		return ret
	}

	return o.StakeDeregistrations
}

// GetStakeDeregistrationsOk returns a tuple with the StakeDeregistrations field value
// and a boolean to check if the value has been set.
func (o *Certificates) GetStakeDeregistrationsOk() ([]StakeRegCert, bool) {
	if o == nil {
		return nil, false
	}
	return o.StakeDeregistrations, true
}

// SetStakeDeregistrations sets field value
func (o *Certificates) SetStakeDeregistrations(v []StakeRegCert) {
	o.StakeDeregistrations = v
}

// GetStakeRegistrations returns the StakeRegistrations field value
func (o *Certificates) GetStakeRegistrations() []StakeRegCert {
	if o == nil {
		var ret []StakeRegCert
		return ret
	}

	return o.StakeRegistrations
}

// GetStakeRegistrationsOk returns a tuple with the StakeRegistrations field value
// and a boolean to check if the value has been set.
func (o *Certificates) GetStakeRegistrationsOk() ([]StakeRegCert, bool) {
	if o == nil {
		return nil, false
	}
	return o.StakeRegistrations, true
}

// SetStakeRegistrations sets field value
func (o *Certificates) SetStakeRegistrations(v []StakeRegCert) {
	o.StakeRegistrations = v
}

func (o Certificates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Certificates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mir_transfers"] = o.MirTransfers
	toSerialize["pool_registrations"] = o.PoolRegistrations
	toSerialize["pool_retirements"] = o.PoolRetirements
	toSerialize["stake_delegations"] = o.StakeDelegations
	toSerialize["stake_deregistrations"] = o.StakeDeregistrations
	toSerialize["stake_registrations"] = o.StakeRegistrations
	return toSerialize, nil
}

type NullableCertificates struct {
	value *Certificates
	isSet bool
}

func (v NullableCertificates) Get() *Certificates {
	return v.value
}

func (v *NullableCertificates) Set(val *Certificates) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificates) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificates(val *Certificates) *NullableCertificates {
	return &NullableCertificates{value: val, isSet: true}
}

func (v NullableCertificates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

