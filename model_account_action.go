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
	"fmt"
)

// AccountAction Type of staking-related action
type AccountAction string

// List of AccountAction
const (
	REGISTRATION AccountAction = "registration"
	DEREGISTRATION AccountAction = "deregistration"
	DELEGATION AccountAction = "delegation"
	WITHDRAWAL AccountAction = "withdrawal"
)

// All allowed values of AccountAction enum
var AllowedAccountActionEnumValues = []AccountAction{
	"registration",
	"deregistration",
	"delegation",
	"withdrawal",
}

func (v *AccountAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountAction(value)
	for _, existing := range AllowedAccountActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountAction", value)
}

// NewAccountActionFromValue returns a pointer to a valid AccountAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountActionFromValue(v string) (*AccountAction, error) {
	ev := AccountAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountAction: valid values are %v", v, AllowedAccountActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountAction) IsValid() bool {
	for _, existing := range AllowedAccountActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountAction value
func (v AccountAction) Ptr() *AccountAction {
	return &v
}

type NullableAccountAction struct {
	value *AccountAction
	isSet bool
}

func (v NullableAccountAction) Get() *AccountAction {
	return v.value
}

func (v *NullableAccountAction) Set(val *AccountAction) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAction(val *AccountAction) *NullableAccountAction {
	return &NullableAccountAction{value: val, isSet: true}
}

func (v NullableAccountAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
