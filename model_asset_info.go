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

// checks if the AssetInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetInfo{}

// AssetInfo struct for AssetInfo
type AssetInfo struct {
	// Hex encoding of the asset name
	AssetName string `json:"asset_name"`
	// ASCII representation of the asset name
	AssetNameAscii *string `json:"asset_name_ascii,omitempty"`
	AssetStandards AssetStandards `json:"asset_standards"`
	// Number of transactions which burned some of the asset
	BurnTxCount int64 `json:"burn_tx_count"`
	// CIP-14 fingerprint of the asset
	Fingerprint string `json:"fingerprint"`
	// UNIX timestamp of the first mint transaction
	FirstMintTime int32 `json:"first_mint_time"`
	// Transaction hash of the first transaction which minted the asset
	FirstMintTx string `json:"first_mint_tx"`
	LatestMintTxMetadata map[string]interface{} `json:"latest_mint_tx_metadata,omitempty"`
	// Number of transactions which minted some of the asset
	MintTxCount int64 `json:"mint_tx_count"`
	TokenRegistryMetadata *TokenRegistryMetadata `json:"token_registry_metadata,omitempty"`
	// Current amount of the asset minted
	TotalSupply int64 `json:"total_supply"`
}

// NewAssetInfo instantiates a new AssetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetInfo(assetName string, assetStandards AssetStandards, burnTxCount int64, fingerprint string, firstMintTime int32, firstMintTx string, mintTxCount int64, totalSupply int64) *AssetInfo {
	this := AssetInfo{}
	this.AssetName = assetName
	this.AssetStandards = assetStandards
	this.BurnTxCount = burnTxCount
	this.Fingerprint = fingerprint
	this.FirstMintTime = firstMintTime
	this.FirstMintTx = firstMintTx
	this.MintTxCount = mintTxCount
	this.TotalSupply = totalSupply
	return &this
}

// NewAssetInfoWithDefaults instantiates a new AssetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetInfoWithDefaults() *AssetInfo {
	this := AssetInfo{}
	return &this
}

// GetAssetName returns the AssetName field value
func (o *AssetInfo) GetAssetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetName
}

// GetAssetNameOk returns a tuple with the AssetName field value
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetAssetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetName, true
}

// SetAssetName sets field value
func (o *AssetInfo) SetAssetName(v string) {
	o.AssetName = v
}

// GetAssetNameAscii returns the AssetNameAscii field value if set, zero value otherwise.
func (o *AssetInfo) GetAssetNameAscii() string {
	if o == nil || IsNil(o.AssetNameAscii) {
		var ret string
		return ret
	}
	return *o.AssetNameAscii
}

// GetAssetNameAsciiOk returns a tuple with the AssetNameAscii field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetAssetNameAsciiOk() (*string, bool) {
	if o == nil || IsNil(o.AssetNameAscii) {
		return nil, false
	}
	return o.AssetNameAscii, true
}

// HasAssetNameAscii returns a boolean if a field has been set.
func (o *AssetInfo) HasAssetNameAscii() bool {
	if o != nil && !IsNil(o.AssetNameAscii) {
		return true
	}

	return false
}

// SetAssetNameAscii gets a reference to the given string and assigns it to the AssetNameAscii field.
func (o *AssetInfo) SetAssetNameAscii(v string) {
	o.AssetNameAscii = &v
}

// GetAssetStandards returns the AssetStandards field value
func (o *AssetInfo) GetAssetStandards() AssetStandards {
	if o == nil {
		var ret AssetStandards
		return ret
	}

	return o.AssetStandards
}

// GetAssetStandardsOk returns a tuple with the AssetStandards field value
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetAssetStandardsOk() (*AssetStandards, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetStandards, true
}

// SetAssetStandards sets field value
func (o *AssetInfo) SetAssetStandards(v AssetStandards) {
	o.AssetStandards = v
}

// GetBurnTxCount returns the BurnTxCount field value
func (o *AssetInfo) GetBurnTxCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BurnTxCount
}

// GetBurnTxCountOk returns a tuple with the BurnTxCount field value
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetBurnTxCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnTxCount, true
}

// SetBurnTxCount sets field value
func (o *AssetInfo) SetBurnTxCount(v int64) {
	o.BurnTxCount = v
}

// GetFingerprint returns the Fingerprint field value
func (o *AssetInfo) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value
func (o *AssetInfo) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetFirstMintTime returns the FirstMintTime field value
func (o *AssetInfo) GetFirstMintTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FirstMintTime
}

// GetFirstMintTimeOk returns a tuple with the FirstMintTime field value
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetFirstMintTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstMintTime, true
}

// SetFirstMintTime sets field value
func (o *AssetInfo) SetFirstMintTime(v int32) {
	o.FirstMintTime = v
}

// GetFirstMintTx returns the FirstMintTx field value
func (o *AssetInfo) GetFirstMintTx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstMintTx
}

// GetFirstMintTxOk returns a tuple with the FirstMintTx field value
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetFirstMintTxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstMintTx, true
}

// SetFirstMintTx sets field value
func (o *AssetInfo) SetFirstMintTx(v string) {
	o.FirstMintTx = v
}

// GetLatestMintTxMetadata returns the LatestMintTxMetadata field value if set, zero value otherwise.
func (o *AssetInfo) GetLatestMintTxMetadata() map[string]interface{} {
	if o == nil || IsNil(o.LatestMintTxMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.LatestMintTxMetadata
}

// GetLatestMintTxMetadataOk returns a tuple with the LatestMintTxMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetLatestMintTxMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.LatestMintTxMetadata) {
		return map[string]interface{}{}, false
	}
	return o.LatestMintTxMetadata, true
}

// HasLatestMintTxMetadata returns a boolean if a field has been set.
func (o *AssetInfo) HasLatestMintTxMetadata() bool {
	if o != nil && !IsNil(o.LatestMintTxMetadata) {
		return true
	}

	return false
}

// SetLatestMintTxMetadata gets a reference to the given map[string]interface{} and assigns it to the LatestMintTxMetadata field.
func (o *AssetInfo) SetLatestMintTxMetadata(v map[string]interface{}) {
	o.LatestMintTxMetadata = v
}

// GetMintTxCount returns the MintTxCount field value
func (o *AssetInfo) GetMintTxCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MintTxCount
}

// GetMintTxCountOk returns a tuple with the MintTxCount field value
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetMintTxCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MintTxCount, true
}

// SetMintTxCount sets field value
func (o *AssetInfo) SetMintTxCount(v int64) {
	o.MintTxCount = v
}

// GetTokenRegistryMetadata returns the TokenRegistryMetadata field value if set, zero value otherwise.
func (o *AssetInfo) GetTokenRegistryMetadata() TokenRegistryMetadata {
	if o == nil || IsNil(o.TokenRegistryMetadata) {
		var ret TokenRegistryMetadata
		return ret
	}
	return *o.TokenRegistryMetadata
}

// GetTokenRegistryMetadataOk returns a tuple with the TokenRegistryMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetTokenRegistryMetadataOk() (*TokenRegistryMetadata, bool) {
	if o == nil || IsNil(o.TokenRegistryMetadata) {
		return nil, false
	}
	return o.TokenRegistryMetadata, true
}

// HasTokenRegistryMetadata returns a boolean if a field has been set.
func (o *AssetInfo) HasTokenRegistryMetadata() bool {
	if o != nil && !IsNil(o.TokenRegistryMetadata) {
		return true
	}

	return false
}

// SetTokenRegistryMetadata gets a reference to the given TokenRegistryMetadata and assigns it to the TokenRegistryMetadata field.
func (o *AssetInfo) SetTokenRegistryMetadata(v TokenRegistryMetadata) {
	o.TokenRegistryMetadata = &v
}

// GetTotalSupply returns the TotalSupply field value
func (o *AssetInfo) GetTotalSupply() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalSupply
}

// GetTotalSupplyOk returns a tuple with the TotalSupply field value
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetTotalSupplyOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSupply, true
}

// SetTotalSupply sets field value
func (o *AssetInfo) SetTotalSupply(v int64) {
	o.TotalSupply = v
}

func (o AssetInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asset_name"] = o.AssetName
	if !IsNil(o.AssetNameAscii) {
		toSerialize["asset_name_ascii"] = o.AssetNameAscii
	}
	toSerialize["asset_standards"] = o.AssetStandards
	toSerialize["burn_tx_count"] = o.BurnTxCount
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["first_mint_time"] = o.FirstMintTime
	toSerialize["first_mint_tx"] = o.FirstMintTx
	if !IsNil(o.LatestMintTxMetadata) {
		toSerialize["latest_mint_tx_metadata"] = o.LatestMintTxMetadata
	}
	toSerialize["mint_tx_count"] = o.MintTxCount
	if !IsNil(o.TokenRegistryMetadata) {
		toSerialize["token_registry_metadata"] = o.TokenRegistryMetadata
	}
	toSerialize["total_supply"] = o.TotalSupply
	return toSerialize, nil
}

type NullableAssetInfo struct {
	value *AssetInfo
	isSet bool
}

func (v NullableAssetInfo) Get() *AssetInfo {
	return v.value
}

func (v *NullableAssetInfo) Set(val *AssetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetInfo(val *AssetInfo) *NullableAssetInfo {
	return &NullableAssetInfo{value: val, isSet: true}
}

func (v NullableAssetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


