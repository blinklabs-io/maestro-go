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

// PaymentCredKind the model 'PaymentCredKind'
type PaymentCredKind string

// List of PaymentCredKind
const (
	KEY PaymentCredKind = "key"
	SCRIPT PaymentCredKind = "script"
)

// All allowed values of PaymentCredKind enum
var AllowedPaymentCredKindEnumValues = []PaymentCredKind{
	"key",
	"script",
}

func (v *PaymentCredKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentCredKind(value)
	for _, existing := range AllowedPaymentCredKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentCredKind", value)
}

// NewPaymentCredKindFromValue returns a pointer to a valid PaymentCredKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentCredKindFromValue(v string) (*PaymentCredKind, error) {
	ev := PaymentCredKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentCredKind: valid values are %v", v, AllowedPaymentCredKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentCredKind) IsValid() bool {
	for _, existing := range AllowedPaymentCredKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentCredKind value
func (v PaymentCredKind) Ptr() *PaymentCredKind {
	return &v
}

type NullablePaymentCredKind struct {
	value *PaymentCredKind
	isSet bool
}

func (v NullablePaymentCredKind) Get() *PaymentCredKind {
	return v.value
}

func (v *NullablePaymentCredKind) Set(val *PaymentCredKind) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentCredKind) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentCredKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentCredKind(val *PaymentCredKind) *NullablePaymentCredKind {
	return &NullablePaymentCredKind{value: val, isSet: true}
}

func (v NullablePaymentCredKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentCredKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
