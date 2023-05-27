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

// checks if the Script type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Script{}

// Script Native or Plutus script
type Script struct {
	// Script bytes (`null` if `native` script)
	Bytes *string `json:"bytes,omitempty"`
	// Script hash
	Hash string `json:"hash"`
	Json map[string]interface{} `json:"json,omitempty"`
	Type ScriptType `json:"type"`
}

// NewScript instantiates a new Script object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScript(hash string, type_ ScriptType) *Script {
	this := Script{}
	this.Hash = hash
	this.Type = type_
	return &this
}

// NewScriptWithDefaults instantiates a new Script object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptWithDefaults() *Script {
	this := Script{}
	return &this
}

// GetBytes returns the Bytes field value if set, zero value otherwise.
func (o *Script) GetBytes() string {
	if o == nil || IsNil(o.Bytes) {
		var ret string
		return ret
	}
	return *o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetBytesOk() (*string, bool) {
	if o == nil || IsNil(o.Bytes) {
		return nil, false
	}
	return o.Bytes, true
}

// HasBytes returns a boolean if a field has been set.
func (o *Script) HasBytes() bool {
	if o != nil && !IsNil(o.Bytes) {
		return true
	}

	return false
}

// SetBytes gets a reference to the given string and assigns it to the Bytes field.
func (o *Script) SetBytes(v string) {
	o.Bytes = &v
}

// GetHash returns the Hash field value
func (o *Script) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *Script) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *Script) SetHash(v string) {
	o.Hash = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *Script) GetJson() map[string]interface{} {
	if o == nil || IsNil(o.Json) {
		var ret map[string]interface{}
		return ret
	}
	return o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetJsonOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Json) {
		return map[string]interface{}{}, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *Script) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given map[string]interface{} and assigns it to the Json field.
func (o *Script) SetJson(v map[string]interface{}) {
	o.Json = v
}

// GetType returns the Type field value
func (o *Script) GetType() ScriptType {
	if o == nil {
		var ret ScriptType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Script) GetTypeOk() (*ScriptType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Script) SetType(v ScriptType) {
	o.Type = v
}

func (o Script) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Script) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bytes) {
		toSerialize["bytes"] = o.Bytes
	}
	toSerialize["hash"] = o.Hash
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableScript struct {
	value *Script
	isSet bool
}

func (v NullableScript) Get() *Script {
	return v.value
}

func (v *NullableScript) Set(val *Script) {
	v.value = val
	v.isSet = true
}

func (v NullableScript) IsSet() bool {
	return v.isSet
}

func (v *NullableScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScript(val *Script) *NullableScript {
	return &NullableScript{value: val, isSet: true}
}

func (v NullableScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


