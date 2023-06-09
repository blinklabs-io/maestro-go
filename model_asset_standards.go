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

// checks if the AssetStandards type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetStandards{}

// AssetStandards Asset information corresponding to popular standards
type AssetStandards struct {
	Cip25Metadata map[string]interface{} `json:"cip25_metadata,omitempty"`
	Cip68Metadata map[string]interface{} `json:"cip68_metadata,omitempty"`
}

// NewAssetStandards instantiates a new AssetStandards object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetStandards() *AssetStandards {
	this := AssetStandards{}
	return &this
}

// NewAssetStandardsWithDefaults instantiates a new AssetStandards object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetStandardsWithDefaults() *AssetStandards {
	this := AssetStandards{}
	return &this
}

// GetCip25Metadata returns the Cip25Metadata field value if set, zero value otherwise.
func (o *AssetStandards) GetCip25Metadata() map[string]interface{} {
	if o == nil || IsNil(o.Cip25Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Cip25Metadata
}

// GetCip25MetadataOk returns a tuple with the Cip25Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetStandards) GetCip25MetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Cip25Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Cip25Metadata, true
}

// HasCip25Metadata returns a boolean if a field has been set.
func (o *AssetStandards) HasCip25Metadata() bool {
	if o != nil && !IsNil(o.Cip25Metadata) {
		return true
	}

	return false
}

// SetCip25Metadata gets a reference to the given map[string]interface{} and assigns it to the Cip25Metadata field.
func (o *AssetStandards) SetCip25Metadata(v map[string]interface{}) {
	o.Cip25Metadata = v
}

// GetCip68Metadata returns the Cip68Metadata field value if set, zero value otherwise.
func (o *AssetStandards) GetCip68Metadata() map[string]interface{} {
	if o == nil || IsNil(o.Cip68Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Cip68Metadata
}

// GetCip68MetadataOk returns a tuple with the Cip68Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetStandards) GetCip68MetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Cip68Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Cip68Metadata, true
}

// HasCip68Metadata returns a boolean if a field has been set.
func (o *AssetStandards) HasCip68Metadata() bool {
	if o != nil && !IsNil(o.Cip68Metadata) {
		return true
	}

	return false
}

// SetCip68Metadata gets a reference to the given map[string]interface{} and assigns it to the Cip68Metadata field.
func (o *AssetStandards) SetCip68Metadata(v map[string]interface{}) {
	o.Cip68Metadata = v
}

func (o AssetStandards) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetStandards) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cip25Metadata) {
		toSerialize["cip25_metadata"] = o.Cip25Metadata
	}
	if !IsNil(o.Cip68Metadata) {
		toSerialize["cip68_metadata"] = o.Cip68Metadata
	}
	return toSerialize, nil
}

type NullableAssetStandards struct {
	value *AssetStandards
	isSet bool
}

func (v NullableAssetStandards) Get() *AssetStandards {
	return v.value
}

func (v *NullableAssetStandards) Set(val *AssetStandards) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetStandards) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetStandards) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetStandards(val *AssetStandards) *NullableAssetStandards {
	return &NullableAssetStandards{value: val, isSet: true}
}

func (v NullableAssetStandards) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetStandards) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


