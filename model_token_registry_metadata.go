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

// checks if the TokenRegistryMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenRegistryMetadata{}

// TokenRegistryMetadata Token registry metadata
type TokenRegistryMetadata struct {
	// Recommended value for decimal places
	Decimals int64 `json:"decimals"`
	// Asset description
	Description string `json:"description"`
	// Base64 encoded logo PNG associated with the asset
	Logo string `json:"logo"`
	// Asset name
	Name string `json:"name"`
	// Asset ticker
	Ticker string `json:"ticker"`
	// URL associated with the asset
	Url string `json:"url"`
}

// NewTokenRegistryMetadata instantiates a new TokenRegistryMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRegistryMetadata(decimals int64, description string, logo string, name string, ticker string, url string) *TokenRegistryMetadata {
	this := TokenRegistryMetadata{}
	this.Decimals = decimals
	this.Description = description
	this.Logo = logo
	this.Name = name
	this.Ticker = ticker
	this.Url = url
	return &this
}

// NewTokenRegistryMetadataWithDefaults instantiates a new TokenRegistryMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRegistryMetadataWithDefaults() *TokenRegistryMetadata {
	this := TokenRegistryMetadata{}
	return &this
}

// GetDecimals returns the Decimals field value
func (o *TokenRegistryMetadata) GetDecimals() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value
// and a boolean to check if the value has been set.
func (o *TokenRegistryMetadata) GetDecimalsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decimals, true
}

// SetDecimals sets field value
func (o *TokenRegistryMetadata) SetDecimals(v int64) {
	o.Decimals = v
}

// GetDescription returns the Description field value
func (o *TokenRegistryMetadata) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TokenRegistryMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TokenRegistryMetadata) SetDescription(v string) {
	o.Description = v
}

// GetLogo returns the Logo field value
func (o *TokenRegistryMetadata) GetLogo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *TokenRegistryMetadata) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value
func (o *TokenRegistryMetadata) SetLogo(v string) {
	o.Logo = v
}

// GetName returns the Name field value
func (o *TokenRegistryMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TokenRegistryMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TokenRegistryMetadata) SetName(v string) {
	o.Name = v
}

// GetTicker returns the Ticker field value
func (o *TokenRegistryMetadata) GetTicker() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ticker
}

// GetTickerOk returns a tuple with the Ticker field value
// and a boolean to check if the value has been set.
func (o *TokenRegistryMetadata) GetTickerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ticker, true
}

// SetTicker sets field value
func (o *TokenRegistryMetadata) SetTicker(v string) {
	o.Ticker = v
}

// GetUrl returns the Url field value
func (o *TokenRegistryMetadata) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TokenRegistryMetadata) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TokenRegistryMetadata) SetUrl(v string) {
	o.Url = v
}

func (o TokenRegistryMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenRegistryMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["decimals"] = o.Decimals
	toSerialize["description"] = o.Description
	toSerialize["logo"] = o.Logo
	toSerialize["name"] = o.Name
	toSerialize["ticker"] = o.Ticker
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableTokenRegistryMetadata struct {
	value *TokenRegistryMetadata
	isSet bool
}

func (v NullableTokenRegistryMetadata) Get() *TokenRegistryMetadata {
	return v.value
}

func (v *NullableTokenRegistryMetadata) Set(val *TokenRegistryMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRegistryMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRegistryMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRegistryMetadata(val *TokenRegistryMetadata) *NullableTokenRegistryMetadata {
	return &NullableTokenRegistryMetadata{value: val, isSet: true}
}

func (v NullableTokenRegistryMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRegistryMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


