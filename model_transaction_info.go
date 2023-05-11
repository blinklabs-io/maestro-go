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

// checks if the TransactionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionInfo{}

// TransactionInfo Transaction Information
type TransactionInfo struct {
	AdditionalSigners []string `json:"additional_signers"`
	// Absolute slot of the block which includes the transaction
	BlockAbsoluteSlot int64 `json:"block_absolute_slot"`
	// Hash of the block which includes the transaction
	BlockHash string `json:"block_hash"`
	// Block height (number) of the block which includes the transaction
	BlockHeight int32 `json:"block_height"`
	// UNIX timestamp of the block which includes the transaction
	BlockTimestamp int64 `json:"block_timestamp"`
	// The transaction's position within the block which includes it
	BlockTxIndex int32 `json:"block_tx_index"`
	Certificates Certificates `json:"certificates"`
	CollateralInputs []Utxo `json:"collateral_inputs"`
	CollateralReturn *Utxo `json:"collateral_return,omitempty"`
	Inputs []Utxo `json:"inputs"`
	// The slot before which the transaction would not be accepted onto the chain
	InvalidBefore *int64 `json:"invalid_before,omitempty"`
	// The slot from which the transaction would not be accepted onto the chain
	InvalidHereafter *int64 `json:"invalid_hereafter,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Mint []MintAsset `json:"mint"`
	Outputs []Utxo `json:"outputs"`
	Redeemers Redeemers `json:"redeemers"`
	ReferenceInputs []Utxo `json:"reference_inputs"`
	ScriptsExecuted []Script `json:"scripts_executed"`
	// False if any executed Plutus scripts failed (aka phase-two validity), meaning collateral was processed.
	ScriptsSuccessful bool `json:"scripts_successful"`
	// Transaction hash (identifier)
	TxHash string `json:"tx_hash"`
	Withdrawals []Withdrawal `json:"withdrawals"`
}

// NewTransactionInfo instantiates a new TransactionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionInfo(additionalSigners []string, blockAbsoluteSlot int64, blockHash string, blockHeight int32, blockTimestamp int64, blockTxIndex int32, certificates Certificates, collateralInputs []Utxo, inputs []Utxo, mint []MintAsset, outputs []Utxo, redeemers Redeemers, referenceInputs []Utxo, scriptsExecuted []Script, scriptsSuccessful bool, txHash string, withdrawals []Withdrawal) *TransactionInfo {
	this := TransactionInfo{}
	this.AdditionalSigners = additionalSigners
	this.BlockAbsoluteSlot = blockAbsoluteSlot
	this.BlockHash = blockHash
	this.BlockHeight = blockHeight
	this.BlockTimestamp = blockTimestamp
	this.BlockTxIndex = blockTxIndex
	this.Certificates = certificates
	this.CollateralInputs = collateralInputs
	this.Inputs = inputs
	this.Mint = mint
	this.Outputs = outputs
	this.Redeemers = redeemers
	this.ReferenceInputs = referenceInputs
	this.ScriptsExecuted = scriptsExecuted
	this.ScriptsSuccessful = scriptsSuccessful
	this.TxHash = txHash
	this.Withdrawals = withdrawals
	return &this
}

// NewTransactionInfoWithDefaults instantiates a new TransactionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionInfoWithDefaults() *TransactionInfo {
	this := TransactionInfo{}
	return &this
}

// GetAdditionalSigners returns the AdditionalSigners field value
func (o *TransactionInfo) GetAdditionalSigners() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AdditionalSigners
}

// GetAdditionalSignersOk returns a tuple with the AdditionalSigners field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetAdditionalSignersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionalSigners, true
}

// SetAdditionalSigners sets field value
func (o *TransactionInfo) SetAdditionalSigners(v []string) {
	o.AdditionalSigners = v
}

// GetBlockAbsoluteSlot returns the BlockAbsoluteSlot field value
func (o *TransactionInfo) GetBlockAbsoluteSlot() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BlockAbsoluteSlot
}

// GetBlockAbsoluteSlotOk returns a tuple with the BlockAbsoluteSlot field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetBlockAbsoluteSlotOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockAbsoluteSlot, true
}

// SetBlockAbsoluteSlot sets field value
func (o *TransactionInfo) SetBlockAbsoluteSlot(v int64) {
	o.BlockAbsoluteSlot = v
}

// GetBlockHash returns the BlockHash field value
func (o *TransactionInfo) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *TransactionInfo) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *TransactionInfo) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *TransactionInfo) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetBlockTimestamp returns the BlockTimestamp field value
func (o *TransactionInfo) GetBlockTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BlockTimestamp
}

// GetBlockTimestampOk returns a tuple with the BlockTimestamp field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetBlockTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTimestamp, true
}

// SetBlockTimestamp sets field value
func (o *TransactionInfo) SetBlockTimestamp(v int64) {
	o.BlockTimestamp = v
}

// GetBlockTxIndex returns the BlockTxIndex field value
func (o *TransactionInfo) GetBlockTxIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockTxIndex
}

// GetBlockTxIndexOk returns a tuple with the BlockTxIndex field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetBlockTxIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTxIndex, true
}

// SetBlockTxIndex sets field value
func (o *TransactionInfo) SetBlockTxIndex(v int32) {
	o.BlockTxIndex = v
}

// GetCertificates returns the Certificates field value
func (o *TransactionInfo) GetCertificates() Certificates {
	if o == nil {
		var ret Certificates
		return ret
	}

	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetCertificatesOk() (*Certificates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificates, true
}

// SetCertificates sets field value
func (o *TransactionInfo) SetCertificates(v Certificates) {
	o.Certificates = v
}

// GetCollateralInputs returns the CollateralInputs field value
func (o *TransactionInfo) GetCollateralInputs() []Utxo {
	if o == nil {
		var ret []Utxo
		return ret
	}

	return o.CollateralInputs
}

// GetCollateralInputsOk returns a tuple with the CollateralInputs field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetCollateralInputsOk() ([]Utxo, bool) {
	if o == nil {
		return nil, false
	}
	return o.CollateralInputs, true
}

// SetCollateralInputs sets field value
func (o *TransactionInfo) SetCollateralInputs(v []Utxo) {
	o.CollateralInputs = v
}

// GetCollateralReturn returns the CollateralReturn field value if set, zero value otherwise.
func (o *TransactionInfo) GetCollateralReturn() Utxo {
	if o == nil || IsNil(o.CollateralReturn) {
		var ret Utxo
		return ret
	}
	return *o.CollateralReturn
}

// GetCollateralReturnOk returns a tuple with the CollateralReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetCollateralReturnOk() (*Utxo, bool) {
	if o == nil || IsNil(o.CollateralReturn) {
		return nil, false
	}
	return o.CollateralReturn, true
}

// HasCollateralReturn returns a boolean if a field has been set.
func (o *TransactionInfo) HasCollateralReturn() bool {
	if o != nil && !IsNil(o.CollateralReturn) {
		return true
	}

	return false
}

// SetCollateralReturn gets a reference to the given Utxo and assigns it to the CollateralReturn field.
func (o *TransactionInfo) SetCollateralReturn(v Utxo) {
	o.CollateralReturn = &v
}

// GetInputs returns the Inputs field value
func (o *TransactionInfo) GetInputs() []Utxo {
	if o == nil {
		var ret []Utxo
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetInputsOk() ([]Utxo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *TransactionInfo) SetInputs(v []Utxo) {
	o.Inputs = v
}

// GetInvalidBefore returns the InvalidBefore field value if set, zero value otherwise.
func (o *TransactionInfo) GetInvalidBefore() int64 {
	if o == nil || IsNil(o.InvalidBefore) {
		var ret int64
		return ret
	}
	return *o.InvalidBefore
}

// GetInvalidBeforeOk returns a tuple with the InvalidBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetInvalidBeforeOk() (*int64, bool) {
	if o == nil || IsNil(o.InvalidBefore) {
		return nil, false
	}
	return o.InvalidBefore, true
}

// HasInvalidBefore returns a boolean if a field has been set.
func (o *TransactionInfo) HasInvalidBefore() bool {
	if o != nil && !IsNil(o.InvalidBefore) {
		return true
	}

	return false
}

// SetInvalidBefore gets a reference to the given int64 and assigns it to the InvalidBefore field.
func (o *TransactionInfo) SetInvalidBefore(v int64) {
	o.InvalidBefore = &v
}

// GetInvalidHereafter returns the InvalidHereafter field value if set, zero value otherwise.
func (o *TransactionInfo) GetInvalidHereafter() int64 {
	if o == nil || IsNil(o.InvalidHereafter) {
		var ret int64
		return ret
	}
	return *o.InvalidHereafter
}

// GetInvalidHereafterOk returns a tuple with the InvalidHereafter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetInvalidHereafterOk() (*int64, bool) {
	if o == nil || IsNil(o.InvalidHereafter) {
		return nil, false
	}
	return o.InvalidHereafter, true
}

// HasInvalidHereafter returns a boolean if a field has been set.
func (o *TransactionInfo) HasInvalidHereafter() bool {
	if o != nil && !IsNil(o.InvalidHereafter) {
		return true
	}

	return false
}

// SetInvalidHereafter gets a reference to the given int64 and assigns it to the InvalidHereafter field.
func (o *TransactionInfo) SetInvalidHereafter(v int64) {
	o.InvalidHereafter = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TransactionInfo) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TransactionInfo) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *TransactionInfo) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMint returns the Mint field value
func (o *TransactionInfo) GetMint() []MintAsset {
	if o == nil {
		var ret []MintAsset
		return ret
	}

	return o.Mint
}

// GetMintOk returns a tuple with the Mint field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetMintOk() ([]MintAsset, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mint, true
}

// SetMint sets field value
func (o *TransactionInfo) SetMint(v []MintAsset) {
	o.Mint = v
}

// GetOutputs returns the Outputs field value
func (o *TransactionInfo) GetOutputs() []Utxo {
	if o == nil {
		var ret []Utxo
		return ret
	}

	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetOutputsOk() ([]Utxo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Outputs, true
}

// SetOutputs sets field value
func (o *TransactionInfo) SetOutputs(v []Utxo) {
	o.Outputs = v
}

// GetRedeemers returns the Redeemers field value
func (o *TransactionInfo) GetRedeemers() Redeemers {
	if o == nil {
		var ret Redeemers
		return ret
	}

	return o.Redeemers
}

// GetRedeemersOk returns a tuple with the Redeemers field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetRedeemersOk() (*Redeemers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redeemers, true
}

// SetRedeemers sets field value
func (o *TransactionInfo) SetRedeemers(v Redeemers) {
	o.Redeemers = v
}

// GetReferenceInputs returns the ReferenceInputs field value
func (o *TransactionInfo) GetReferenceInputs() []Utxo {
	if o == nil {
		var ret []Utxo
		return ret
	}

	return o.ReferenceInputs
}

// GetReferenceInputsOk returns a tuple with the ReferenceInputs field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetReferenceInputsOk() ([]Utxo, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceInputs, true
}

// SetReferenceInputs sets field value
func (o *TransactionInfo) SetReferenceInputs(v []Utxo) {
	o.ReferenceInputs = v
}

// GetScriptsExecuted returns the ScriptsExecuted field value
func (o *TransactionInfo) GetScriptsExecuted() []Script {
	if o == nil {
		var ret []Script
		return ret
	}

	return o.ScriptsExecuted
}

// GetScriptsExecutedOk returns a tuple with the ScriptsExecuted field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetScriptsExecutedOk() ([]Script, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScriptsExecuted, true
}

// SetScriptsExecuted sets field value
func (o *TransactionInfo) SetScriptsExecuted(v []Script) {
	o.ScriptsExecuted = v
}

// GetScriptsSuccessful returns the ScriptsSuccessful field value
func (o *TransactionInfo) GetScriptsSuccessful() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ScriptsSuccessful
}

// GetScriptsSuccessfulOk returns a tuple with the ScriptsSuccessful field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetScriptsSuccessfulOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptsSuccessful, true
}

// SetScriptsSuccessful sets field value
func (o *TransactionInfo) SetScriptsSuccessful(v bool) {
	o.ScriptsSuccessful = v
}

// GetTxHash returns the TxHash field value
func (o *TransactionInfo) GetTxHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetTxHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxHash, true
}

// SetTxHash sets field value
func (o *TransactionInfo) SetTxHash(v string) {
	o.TxHash = v
}

// GetWithdrawals returns the Withdrawals field value
func (o *TransactionInfo) GetWithdrawals() []Withdrawal {
	if o == nil {
		var ret []Withdrawal
		return ret
	}

	return o.Withdrawals
}

// GetWithdrawalsOk returns a tuple with the Withdrawals field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetWithdrawalsOk() ([]Withdrawal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Withdrawals, true
}

// SetWithdrawals sets field value
func (o *TransactionInfo) SetWithdrawals(v []Withdrawal) {
	o.Withdrawals = v
}

func (o TransactionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["additional_signers"] = o.AdditionalSigners
	toSerialize["block_absolute_slot"] = o.BlockAbsoluteSlot
	toSerialize["block_hash"] = o.BlockHash
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["block_timestamp"] = o.BlockTimestamp
	toSerialize["block_tx_index"] = o.BlockTxIndex
	toSerialize["certificates"] = o.Certificates
	toSerialize["collateral_inputs"] = o.CollateralInputs
	if !IsNil(o.CollateralReturn) {
		toSerialize["collateral_return"] = o.CollateralReturn
	}
	toSerialize["inputs"] = o.Inputs
	if !IsNil(o.InvalidBefore) {
		toSerialize["invalid_before"] = o.InvalidBefore
	}
	if !IsNil(o.InvalidHereafter) {
		toSerialize["invalid_hereafter"] = o.InvalidHereafter
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["mint"] = o.Mint
	toSerialize["outputs"] = o.Outputs
	toSerialize["redeemers"] = o.Redeemers
	toSerialize["reference_inputs"] = o.ReferenceInputs
	toSerialize["scripts_executed"] = o.ScriptsExecuted
	toSerialize["scripts_successful"] = o.ScriptsSuccessful
	toSerialize["tx_hash"] = o.TxHash
	toSerialize["withdrawals"] = o.Withdrawals
	return toSerialize, nil
}

type NullableTransactionInfo struct {
	value *TransactionInfo
	isSet bool
}

func (v NullableTransactionInfo) Get() *TransactionInfo {
	return v.value
}

func (v *NullableTransactionInfo) Set(val *TransactionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionInfo(val *TransactionInfo) *NullableTransactionInfo {
	return &NullableTransactionInfo{value: val, isSet: true}
}

func (v NullableTransactionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


