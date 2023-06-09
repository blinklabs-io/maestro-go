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

// checks if the EraParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EraParameters{}

// EraParameters struct for EraParameters
type EraParameters struct {
	EpochLength int64 `json:"epoch_length"`
	SafeZone *int64 `json:"safe_zone,omitempty"`
	SlotLength int64 `json:"slot_length"`
}

// NewEraParameters instantiates a new EraParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEraParameters(epochLength int64, slotLength int64) *EraParameters {
	this := EraParameters{}
	this.EpochLength = epochLength
	this.SlotLength = slotLength
	return &this
}

// NewEraParametersWithDefaults instantiates a new EraParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEraParametersWithDefaults() *EraParameters {
	this := EraParameters{}
	return &this
}

// GetEpochLength returns the EpochLength field value
func (o *EraParameters) GetEpochLength() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EpochLength
}

// GetEpochLengthOk returns a tuple with the EpochLength field value
// and a boolean to check if the value has been set.
func (o *EraParameters) GetEpochLengthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EpochLength, true
}

// SetEpochLength sets field value
func (o *EraParameters) SetEpochLength(v int64) {
	o.EpochLength = v
}

// GetSafeZone returns the SafeZone field value if set, zero value otherwise.
func (o *EraParameters) GetSafeZone() int64 {
	if o == nil || IsNil(o.SafeZone) {
		var ret int64
		return ret
	}
	return *o.SafeZone
}

// GetSafeZoneOk returns a tuple with the SafeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EraParameters) GetSafeZoneOk() (*int64, bool) {
	if o == nil || IsNil(o.SafeZone) {
		return nil, false
	}
	return o.SafeZone, true
}

// HasSafeZone returns a boolean if a field has been set.
func (o *EraParameters) HasSafeZone() bool {
	if o != nil && !IsNil(o.SafeZone) {
		return true
	}

	return false
}

// SetSafeZone gets a reference to the given int64 and assigns it to the SafeZone field.
func (o *EraParameters) SetSafeZone(v int64) {
	o.SafeZone = &v
}

// GetSlotLength returns the SlotLength field value
func (o *EraParameters) GetSlotLength() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SlotLength
}

// GetSlotLengthOk returns a tuple with the SlotLength field value
// and a boolean to check if the value has been set.
func (o *EraParameters) GetSlotLengthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlotLength, true
}

// SetSlotLength sets field value
func (o *EraParameters) SetSlotLength(v int64) {
	o.SlotLength = v
}

func (o EraParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EraParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["epoch_length"] = o.EpochLength
	if !IsNil(o.SafeZone) {
		toSerialize["safe_zone"] = o.SafeZone
	}
	toSerialize["slot_length"] = o.SlotLength
	return toSerialize, nil
}

type NullableEraParameters struct {
	value *EraParameters
	isSet bool
}

func (v NullableEraParameters) Get() *EraParameters {
	return v.value
}

func (v *NullableEraParameters) Set(val *EraParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableEraParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableEraParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEraParameters(val *EraParameters) *NullableEraParameters {
	return &NullableEraParameters{value: val, isSet: true}
}

func (v NullableEraParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEraParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


