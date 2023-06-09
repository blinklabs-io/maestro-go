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

// AccountStakingRewardType Staking-related reward type
type AccountStakingRewardType string

// List of AccountStakingRewardType
const (
	ACCOUNTSTAKINGREWARDTYPE_MEMBER AccountStakingRewardType = "member"
	ACCOUNTSTAKINGREWARDTYPE_LEADER AccountStakingRewardType = "leader"
	ACCOUNTSTAKINGREWARDTYPE_REFUND AccountStakingRewardType = "refund"
)

// All allowed values of AccountStakingRewardType enum
var AllowedAccountStakingRewardTypeEnumValues = []AccountStakingRewardType{
	"member",
	"leader",
	"refund",
}

func (v *AccountStakingRewardType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountStakingRewardType(value)
	for _, existing := range AllowedAccountStakingRewardTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountStakingRewardType", value)
}

// NewAccountStakingRewardTypeFromValue returns a pointer to a valid AccountStakingRewardType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountStakingRewardTypeFromValue(v string) (*AccountStakingRewardType, error) {
	ev := AccountStakingRewardType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountStakingRewardType: valid values are %v", v, AllowedAccountStakingRewardTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountStakingRewardType) IsValid() bool {
	for _, existing := range AllowedAccountStakingRewardTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountStakingRewardType value
func (v AccountStakingRewardType) Ptr() *AccountStakingRewardType {
	return &v
}

type NullableAccountStakingRewardType struct {
	value *AccountStakingRewardType
	isSet bool
}

func (v NullableAccountStakingRewardType) Get() *AccountStakingRewardType {
	return v.value
}

func (v *NullableAccountStakingRewardType) Set(val *AccountStakingRewardType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountStakingRewardType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountStakingRewardType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountStakingRewardType(val *AccountStakingRewardType) *NullableAccountStakingRewardType {
	return &NullableAccountStakingRewardType{value: val, isSet: true}
}

func (v NullableAccountStakingRewardType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountStakingRewardType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

