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
	"fmt"
)

// NetworkId the model 'NetworkId'
type NetworkId string

// List of NetworkId
const (
	NETWORKID_MAINNET NetworkId = "mainnet"
	NETWORKID_TESTNET NetworkId = "testnet"
)

// All allowed values of NetworkId enum
var AllowedNetworkIdEnumValues = []NetworkId{
	"mainnet",
	"testnet",
}

func (v *NetworkId) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkId(value)
	for _, existing := range AllowedNetworkIdEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkId", value)
}

// NewNetworkIdFromValue returns a pointer to a valid NetworkId
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkIdFromValue(v string) (*NetworkId, error) {
	ev := NetworkId(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkId: valid values are %v", v, AllowedNetworkIdEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkId) IsValid() bool {
	for _, existing := range AllowedNetworkIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkId value
func (v NetworkId) Ptr() *NetworkId {
	return &v
}

type NullableNetworkId struct {
	value *NetworkId
	isSet bool
}

func (v NullableNetworkId) Get() *NetworkId {
	return v.value
}

func (v *NullableNetworkId) Set(val *NetworkId) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkId) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkId(val *NetworkId) *NullableNetworkId {
	return &NullableNetworkId{value: val, isSet: true}
}

func (v NullableNetworkId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

