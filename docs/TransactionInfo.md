# TransactionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalSigners** | **[]string** |  | 
**BlockAbsoluteSlot** | **int64** | Absolute slot of the block which includes the transaction | 
**BlockHash** | **string** | Hash of the block which includes the transaction | 
**BlockHeight** | **int32** | Block height (number) of the block which includes the transaction | 
**BlockTimestamp** | **int64** | UNIX timestamp of the block which includes the transaction | 
**BlockTxIndex** | **int32** | The transaction&#39;s position within the block which includes it | 
**Certificates** | [**Certificates**](Certificates.md) |  | 
**CollateralInputs** | [**[]Utxo**](Utxo.md) |  | 
**CollateralReturn** | Pointer to [**Utxo**](Utxo.md) |  | [optional] 
**Inputs** | [**[]Utxo**](Utxo.md) |  | 
**InvalidBefore** | Pointer to **int64** | The slot before which the transaction would not be accepted onto the chain | [optional] 
**InvalidHereafter** | Pointer to **int64** | The slot from which the transaction would not be accepted onto the chain | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Mint** | [**[]MintAsset**](MintAsset.md) |  | 
**Outputs** | [**[]Utxo**](Utxo.md) |  | 
**Redeemers** | [**Redeemers**](Redeemers.md) |  | 
**ReferenceInputs** | [**[]Utxo**](Utxo.md) |  | 
**ScriptsExecuted** | [**[]Script**](Script.md) |  | 
**ScriptsSuccessful** | **bool** | False if any executed Plutus scripts failed (aka phase-two validity), meaning collateral was processed. | 
**TxHash** | **string** | Transaction hash (identifier) | 
**Withdrawals** | [**[]Withdrawal**](Withdrawal.md) |  | 

## Methods

### NewTransactionInfo

`func NewTransactionInfo(additionalSigners []string, blockAbsoluteSlot int64, blockHash string, blockHeight int32, blockTimestamp int64, blockTxIndex int32, certificates Certificates, collateralInputs []Utxo, inputs []Utxo, mint []MintAsset, outputs []Utxo, redeemers Redeemers, referenceInputs []Utxo, scriptsExecuted []Script, scriptsSuccessful bool, txHash string, withdrawals []Withdrawal, ) *TransactionInfo`

NewTransactionInfo instantiates a new TransactionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionInfoWithDefaults

`func NewTransactionInfoWithDefaults() *TransactionInfo`

NewTransactionInfoWithDefaults instantiates a new TransactionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalSigners

`func (o *TransactionInfo) GetAdditionalSigners() []string`

GetAdditionalSigners returns the AdditionalSigners field if non-nil, zero value otherwise.

### GetAdditionalSignersOk

`func (o *TransactionInfo) GetAdditionalSignersOk() (*[]string, bool)`

GetAdditionalSignersOk returns a tuple with the AdditionalSigners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSigners

`func (o *TransactionInfo) SetAdditionalSigners(v []string)`

SetAdditionalSigners sets AdditionalSigners field to given value.


### GetBlockAbsoluteSlot

`func (o *TransactionInfo) GetBlockAbsoluteSlot() int64`

GetBlockAbsoluteSlot returns the BlockAbsoluteSlot field if non-nil, zero value otherwise.

### GetBlockAbsoluteSlotOk

`func (o *TransactionInfo) GetBlockAbsoluteSlotOk() (*int64, bool)`

GetBlockAbsoluteSlotOk returns a tuple with the BlockAbsoluteSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAbsoluteSlot

`func (o *TransactionInfo) SetBlockAbsoluteSlot(v int64)`

SetBlockAbsoluteSlot sets BlockAbsoluteSlot field to given value.


### GetBlockHash

`func (o *TransactionInfo) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TransactionInfo) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TransactionInfo) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockHeight

`func (o *TransactionInfo) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *TransactionInfo) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *TransactionInfo) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetBlockTimestamp

`func (o *TransactionInfo) GetBlockTimestamp() int64`

GetBlockTimestamp returns the BlockTimestamp field if non-nil, zero value otherwise.

### GetBlockTimestampOk

`func (o *TransactionInfo) GetBlockTimestampOk() (*int64, bool)`

GetBlockTimestampOk returns a tuple with the BlockTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimestamp

`func (o *TransactionInfo) SetBlockTimestamp(v int64)`

SetBlockTimestamp sets BlockTimestamp field to given value.


### GetBlockTxIndex

`func (o *TransactionInfo) GetBlockTxIndex() int32`

GetBlockTxIndex returns the BlockTxIndex field if non-nil, zero value otherwise.

### GetBlockTxIndexOk

`func (o *TransactionInfo) GetBlockTxIndexOk() (*int32, bool)`

GetBlockTxIndexOk returns a tuple with the BlockTxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTxIndex

`func (o *TransactionInfo) SetBlockTxIndex(v int32)`

SetBlockTxIndex sets BlockTxIndex field to given value.


### GetCertificates

`func (o *TransactionInfo) GetCertificates() Certificates`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *TransactionInfo) GetCertificatesOk() (*Certificates, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *TransactionInfo) SetCertificates(v Certificates)`

SetCertificates sets Certificates field to given value.


### GetCollateralInputs

`func (o *TransactionInfo) GetCollateralInputs() []Utxo`

GetCollateralInputs returns the CollateralInputs field if non-nil, zero value otherwise.

### GetCollateralInputsOk

`func (o *TransactionInfo) GetCollateralInputsOk() (*[]Utxo, bool)`

GetCollateralInputsOk returns a tuple with the CollateralInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralInputs

`func (o *TransactionInfo) SetCollateralInputs(v []Utxo)`

SetCollateralInputs sets CollateralInputs field to given value.


### GetCollateralReturn

`func (o *TransactionInfo) GetCollateralReturn() Utxo`

GetCollateralReturn returns the CollateralReturn field if non-nil, zero value otherwise.

### GetCollateralReturnOk

`func (o *TransactionInfo) GetCollateralReturnOk() (*Utxo, bool)`

GetCollateralReturnOk returns a tuple with the CollateralReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralReturn

`func (o *TransactionInfo) SetCollateralReturn(v Utxo)`

SetCollateralReturn sets CollateralReturn field to given value.

### HasCollateralReturn

`func (o *TransactionInfo) HasCollateralReturn() bool`

HasCollateralReturn returns a boolean if a field has been set.

### GetInputs

`func (o *TransactionInfo) GetInputs() []Utxo`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *TransactionInfo) GetInputsOk() (*[]Utxo, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *TransactionInfo) SetInputs(v []Utxo)`

SetInputs sets Inputs field to given value.


### GetInvalidBefore

`func (o *TransactionInfo) GetInvalidBefore() int64`

GetInvalidBefore returns the InvalidBefore field if non-nil, zero value otherwise.

### GetInvalidBeforeOk

`func (o *TransactionInfo) GetInvalidBeforeOk() (*int64, bool)`

GetInvalidBeforeOk returns a tuple with the InvalidBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidBefore

`func (o *TransactionInfo) SetInvalidBefore(v int64)`

SetInvalidBefore sets InvalidBefore field to given value.

### HasInvalidBefore

`func (o *TransactionInfo) HasInvalidBefore() bool`

HasInvalidBefore returns a boolean if a field has been set.

### GetInvalidHereafter

`func (o *TransactionInfo) GetInvalidHereafter() int64`

GetInvalidHereafter returns the InvalidHereafter field if non-nil, zero value otherwise.

### GetInvalidHereafterOk

`func (o *TransactionInfo) GetInvalidHereafterOk() (*int64, bool)`

GetInvalidHereafterOk returns a tuple with the InvalidHereafter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidHereafter

`func (o *TransactionInfo) SetInvalidHereafter(v int64)`

SetInvalidHereafter sets InvalidHereafter field to given value.

### HasInvalidHereafter

`func (o *TransactionInfo) HasInvalidHereafter() bool`

HasInvalidHereafter returns a boolean if a field has been set.

### GetMetadata

`func (o *TransactionInfo) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TransactionInfo) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TransactionInfo) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TransactionInfo) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMint

`func (o *TransactionInfo) GetMint() []MintAsset`

GetMint returns the Mint field if non-nil, zero value otherwise.

### GetMintOk

`func (o *TransactionInfo) GetMintOk() (*[]MintAsset, bool)`

GetMintOk returns a tuple with the Mint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMint

`func (o *TransactionInfo) SetMint(v []MintAsset)`

SetMint sets Mint field to given value.


### GetOutputs

`func (o *TransactionInfo) GetOutputs() []Utxo`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *TransactionInfo) GetOutputsOk() (*[]Utxo, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *TransactionInfo) SetOutputs(v []Utxo)`

SetOutputs sets Outputs field to given value.


### GetRedeemers

`func (o *TransactionInfo) GetRedeemers() Redeemers`

GetRedeemers returns the Redeemers field if non-nil, zero value otherwise.

### GetRedeemersOk

`func (o *TransactionInfo) GetRedeemersOk() (*Redeemers, bool)`

GetRedeemersOk returns a tuple with the Redeemers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemers

`func (o *TransactionInfo) SetRedeemers(v Redeemers)`

SetRedeemers sets Redeemers field to given value.


### GetReferenceInputs

`func (o *TransactionInfo) GetReferenceInputs() []Utxo`

GetReferenceInputs returns the ReferenceInputs field if non-nil, zero value otherwise.

### GetReferenceInputsOk

`func (o *TransactionInfo) GetReferenceInputsOk() (*[]Utxo, bool)`

GetReferenceInputsOk returns a tuple with the ReferenceInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceInputs

`func (o *TransactionInfo) SetReferenceInputs(v []Utxo)`

SetReferenceInputs sets ReferenceInputs field to given value.


### GetScriptsExecuted

`func (o *TransactionInfo) GetScriptsExecuted() []Script`

GetScriptsExecuted returns the ScriptsExecuted field if non-nil, zero value otherwise.

### GetScriptsExecutedOk

`func (o *TransactionInfo) GetScriptsExecutedOk() (*[]Script, bool)`

GetScriptsExecutedOk returns a tuple with the ScriptsExecuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptsExecuted

`func (o *TransactionInfo) SetScriptsExecuted(v []Script)`

SetScriptsExecuted sets ScriptsExecuted field to given value.


### GetScriptsSuccessful

`func (o *TransactionInfo) GetScriptsSuccessful() bool`

GetScriptsSuccessful returns the ScriptsSuccessful field if non-nil, zero value otherwise.

### GetScriptsSuccessfulOk

`func (o *TransactionInfo) GetScriptsSuccessfulOk() (*bool, bool)`

GetScriptsSuccessfulOk returns a tuple with the ScriptsSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptsSuccessful

`func (o *TransactionInfo) SetScriptsSuccessful(v bool)`

SetScriptsSuccessful sets ScriptsSuccessful field to given value.


### GetTxHash

`func (o *TransactionInfo) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *TransactionInfo) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *TransactionInfo) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetWithdrawals

`func (o *TransactionInfo) GetWithdrawals() []Withdrawal`

GetWithdrawals returns the Withdrawals field if non-nil, zero value otherwise.

### GetWithdrawalsOk

`func (o *TransactionInfo) GetWithdrawalsOk() (*[]Withdrawal, bool)`

GetWithdrawalsOk returns a tuple with the Withdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawals

`func (o *TransactionInfo) SetWithdrawals(v []Withdrawal)`

SetWithdrawals sets Withdrawals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


