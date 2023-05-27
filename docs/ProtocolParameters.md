# ProtocolParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoinsPerUtxoByte** | **int64** |  | 
**CollateralPercentage** | **int64** |  | 
**CostModels** | **map[string]map[string]int64** |  | 
**DesiredNumberOfPools** | **int64** |  | 
**MaxBlockBodySize** | **int64** |  | 
**MaxBlockHeaderSize** | **int64** |  | 
**MaxCollateralInputs** | **int64** |  | 
**MaxExecutionUnitsPerBlock** | [**ExUnit**](ExUnit.md) |  | 
**MaxExecutionUnitsPerTransaction** | [**ExUnit**](ExUnit.md) |  | 
**MaxTxSize** | **int64** |  | 
**MaxValueSize** | **int64** |  | 
**MinFeeCoefficient** | **int64** |  | 
**MinFeeConstant** | **int64** |  | 
**MinPoolCost** | **int64** |  | 
**MonetaryExpansion** | **string** |  | 
**PoolDeposit** | **int64** |  | 
**PoolInfluence** | **string** |  | 
**PoolRetirementEpochBound** | **int64** |  | 
**Prices** | [**Prices**](Prices.md) |  | 
**ProtocolVersion** | [**Version**](Version.md) |  | 
**StakeKeyDeposit** | **int64** |  | 
**TreasuryExpansion** | **string** |  | 

## Methods

### NewProtocolParameters

`func NewProtocolParameters(coinsPerUtxoByte int64, collateralPercentage int64, costModels map[string]map[string]int64, desiredNumberOfPools int64, maxBlockBodySize int64, maxBlockHeaderSize int64, maxCollateralInputs int64, maxExecutionUnitsPerBlock ExUnit, maxExecutionUnitsPerTransaction ExUnit, maxTxSize int64, maxValueSize int64, minFeeCoefficient int64, minFeeConstant int64, minPoolCost int64, monetaryExpansion string, poolDeposit int64, poolInfluence string, poolRetirementEpochBound int64, prices Prices, protocolVersion Version, stakeKeyDeposit int64, treasuryExpansion string, ) *ProtocolParameters`

NewProtocolParameters instantiates a new ProtocolParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolParametersWithDefaults

`func NewProtocolParametersWithDefaults() *ProtocolParameters`

NewProtocolParametersWithDefaults instantiates a new ProtocolParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoinsPerUtxoByte

`func (o *ProtocolParameters) GetCoinsPerUtxoByte() int64`

GetCoinsPerUtxoByte returns the CoinsPerUtxoByte field if non-nil, zero value otherwise.

### GetCoinsPerUtxoByteOk

`func (o *ProtocolParameters) GetCoinsPerUtxoByteOk() (*int64, bool)`

GetCoinsPerUtxoByteOk returns a tuple with the CoinsPerUtxoByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinsPerUtxoByte

`func (o *ProtocolParameters) SetCoinsPerUtxoByte(v int64)`

SetCoinsPerUtxoByte sets CoinsPerUtxoByte field to given value.


### GetCollateralPercentage

`func (o *ProtocolParameters) GetCollateralPercentage() int64`

GetCollateralPercentage returns the CollateralPercentage field if non-nil, zero value otherwise.

### GetCollateralPercentageOk

`func (o *ProtocolParameters) GetCollateralPercentageOk() (*int64, bool)`

GetCollateralPercentageOk returns a tuple with the CollateralPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralPercentage

`func (o *ProtocolParameters) SetCollateralPercentage(v int64)`

SetCollateralPercentage sets CollateralPercentage field to given value.


### GetCostModels

`func (o *ProtocolParameters) GetCostModels() map[string]map[string]int64`

GetCostModels returns the CostModels field if non-nil, zero value otherwise.

### GetCostModelsOk

`func (o *ProtocolParameters) GetCostModelsOk() (*map[string]map[string]int64, bool)`

GetCostModelsOk returns a tuple with the CostModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostModels

`func (o *ProtocolParameters) SetCostModels(v map[string]map[string]int64)`

SetCostModels sets CostModels field to given value.


### GetDesiredNumberOfPools

`func (o *ProtocolParameters) GetDesiredNumberOfPools() int64`

GetDesiredNumberOfPools returns the DesiredNumberOfPools field if non-nil, zero value otherwise.

### GetDesiredNumberOfPoolsOk

`func (o *ProtocolParameters) GetDesiredNumberOfPoolsOk() (*int64, bool)`

GetDesiredNumberOfPoolsOk returns a tuple with the DesiredNumberOfPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNumberOfPools

`func (o *ProtocolParameters) SetDesiredNumberOfPools(v int64)`

SetDesiredNumberOfPools sets DesiredNumberOfPools field to given value.


### GetMaxBlockBodySize

`func (o *ProtocolParameters) GetMaxBlockBodySize() int64`

GetMaxBlockBodySize returns the MaxBlockBodySize field if non-nil, zero value otherwise.

### GetMaxBlockBodySizeOk

`func (o *ProtocolParameters) GetMaxBlockBodySizeOk() (*int64, bool)`

GetMaxBlockBodySizeOk returns a tuple with the MaxBlockBodySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBlockBodySize

`func (o *ProtocolParameters) SetMaxBlockBodySize(v int64)`

SetMaxBlockBodySize sets MaxBlockBodySize field to given value.


### GetMaxBlockHeaderSize

`func (o *ProtocolParameters) GetMaxBlockHeaderSize() int64`

GetMaxBlockHeaderSize returns the MaxBlockHeaderSize field if non-nil, zero value otherwise.

### GetMaxBlockHeaderSizeOk

`func (o *ProtocolParameters) GetMaxBlockHeaderSizeOk() (*int64, bool)`

GetMaxBlockHeaderSizeOk returns a tuple with the MaxBlockHeaderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBlockHeaderSize

`func (o *ProtocolParameters) SetMaxBlockHeaderSize(v int64)`

SetMaxBlockHeaderSize sets MaxBlockHeaderSize field to given value.


### GetMaxCollateralInputs

`func (o *ProtocolParameters) GetMaxCollateralInputs() int64`

GetMaxCollateralInputs returns the MaxCollateralInputs field if non-nil, zero value otherwise.

### GetMaxCollateralInputsOk

`func (o *ProtocolParameters) GetMaxCollateralInputsOk() (*int64, bool)`

GetMaxCollateralInputsOk returns a tuple with the MaxCollateralInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCollateralInputs

`func (o *ProtocolParameters) SetMaxCollateralInputs(v int64)`

SetMaxCollateralInputs sets MaxCollateralInputs field to given value.


### GetMaxExecutionUnitsPerBlock

`func (o *ProtocolParameters) GetMaxExecutionUnitsPerBlock() ExUnit`

GetMaxExecutionUnitsPerBlock returns the MaxExecutionUnitsPerBlock field if non-nil, zero value otherwise.

### GetMaxExecutionUnitsPerBlockOk

`func (o *ProtocolParameters) GetMaxExecutionUnitsPerBlockOk() (*ExUnit, bool)`

GetMaxExecutionUnitsPerBlockOk returns a tuple with the MaxExecutionUnitsPerBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExecutionUnitsPerBlock

`func (o *ProtocolParameters) SetMaxExecutionUnitsPerBlock(v ExUnit)`

SetMaxExecutionUnitsPerBlock sets MaxExecutionUnitsPerBlock field to given value.


### GetMaxExecutionUnitsPerTransaction

`func (o *ProtocolParameters) GetMaxExecutionUnitsPerTransaction() ExUnit`

GetMaxExecutionUnitsPerTransaction returns the MaxExecutionUnitsPerTransaction field if non-nil, zero value otherwise.

### GetMaxExecutionUnitsPerTransactionOk

`func (o *ProtocolParameters) GetMaxExecutionUnitsPerTransactionOk() (*ExUnit, bool)`

GetMaxExecutionUnitsPerTransactionOk returns a tuple with the MaxExecutionUnitsPerTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExecutionUnitsPerTransaction

`func (o *ProtocolParameters) SetMaxExecutionUnitsPerTransaction(v ExUnit)`

SetMaxExecutionUnitsPerTransaction sets MaxExecutionUnitsPerTransaction field to given value.


### GetMaxTxSize

`func (o *ProtocolParameters) GetMaxTxSize() int64`

GetMaxTxSize returns the MaxTxSize field if non-nil, zero value otherwise.

### GetMaxTxSizeOk

`func (o *ProtocolParameters) GetMaxTxSizeOk() (*int64, bool)`

GetMaxTxSizeOk returns a tuple with the MaxTxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTxSize

`func (o *ProtocolParameters) SetMaxTxSize(v int64)`

SetMaxTxSize sets MaxTxSize field to given value.


### GetMaxValueSize

`func (o *ProtocolParameters) GetMaxValueSize() int64`

GetMaxValueSize returns the MaxValueSize field if non-nil, zero value otherwise.

### GetMaxValueSizeOk

`func (o *ProtocolParameters) GetMaxValueSizeOk() (*int64, bool)`

GetMaxValueSizeOk returns a tuple with the MaxValueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValueSize

`func (o *ProtocolParameters) SetMaxValueSize(v int64)`

SetMaxValueSize sets MaxValueSize field to given value.


### GetMinFeeCoefficient

`func (o *ProtocolParameters) GetMinFeeCoefficient() int64`

GetMinFeeCoefficient returns the MinFeeCoefficient field if non-nil, zero value otherwise.

### GetMinFeeCoefficientOk

`func (o *ProtocolParameters) GetMinFeeCoefficientOk() (*int64, bool)`

GetMinFeeCoefficientOk returns a tuple with the MinFeeCoefficient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFeeCoefficient

`func (o *ProtocolParameters) SetMinFeeCoefficient(v int64)`

SetMinFeeCoefficient sets MinFeeCoefficient field to given value.


### GetMinFeeConstant

`func (o *ProtocolParameters) GetMinFeeConstant() int64`

GetMinFeeConstant returns the MinFeeConstant field if non-nil, zero value otherwise.

### GetMinFeeConstantOk

`func (o *ProtocolParameters) GetMinFeeConstantOk() (*int64, bool)`

GetMinFeeConstantOk returns a tuple with the MinFeeConstant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFeeConstant

`func (o *ProtocolParameters) SetMinFeeConstant(v int64)`

SetMinFeeConstant sets MinFeeConstant field to given value.


### GetMinPoolCost

`func (o *ProtocolParameters) GetMinPoolCost() int64`

GetMinPoolCost returns the MinPoolCost field if non-nil, zero value otherwise.

### GetMinPoolCostOk

`func (o *ProtocolParameters) GetMinPoolCostOk() (*int64, bool)`

GetMinPoolCostOk returns a tuple with the MinPoolCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPoolCost

`func (o *ProtocolParameters) SetMinPoolCost(v int64)`

SetMinPoolCost sets MinPoolCost field to given value.


### GetMonetaryExpansion

`func (o *ProtocolParameters) GetMonetaryExpansion() string`

GetMonetaryExpansion returns the MonetaryExpansion field if non-nil, zero value otherwise.

### GetMonetaryExpansionOk

`func (o *ProtocolParameters) GetMonetaryExpansionOk() (*string, bool)`

GetMonetaryExpansionOk returns a tuple with the MonetaryExpansion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonetaryExpansion

`func (o *ProtocolParameters) SetMonetaryExpansion(v string)`

SetMonetaryExpansion sets MonetaryExpansion field to given value.


### GetPoolDeposit

`func (o *ProtocolParameters) GetPoolDeposit() int64`

GetPoolDeposit returns the PoolDeposit field if non-nil, zero value otherwise.

### GetPoolDepositOk

`func (o *ProtocolParameters) GetPoolDepositOk() (*int64, bool)`

GetPoolDepositOk returns a tuple with the PoolDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolDeposit

`func (o *ProtocolParameters) SetPoolDeposit(v int64)`

SetPoolDeposit sets PoolDeposit field to given value.


### GetPoolInfluence

`func (o *ProtocolParameters) GetPoolInfluence() string`

GetPoolInfluence returns the PoolInfluence field if non-nil, zero value otherwise.

### GetPoolInfluenceOk

`func (o *ProtocolParameters) GetPoolInfluenceOk() (*string, bool)`

GetPoolInfluenceOk returns a tuple with the PoolInfluence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolInfluence

`func (o *ProtocolParameters) SetPoolInfluence(v string)`

SetPoolInfluence sets PoolInfluence field to given value.


### GetPoolRetirementEpochBound

`func (o *ProtocolParameters) GetPoolRetirementEpochBound() int64`

GetPoolRetirementEpochBound returns the PoolRetirementEpochBound field if non-nil, zero value otherwise.

### GetPoolRetirementEpochBoundOk

`func (o *ProtocolParameters) GetPoolRetirementEpochBoundOk() (*int64, bool)`

GetPoolRetirementEpochBoundOk returns a tuple with the PoolRetirementEpochBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolRetirementEpochBound

`func (o *ProtocolParameters) SetPoolRetirementEpochBound(v int64)`

SetPoolRetirementEpochBound sets PoolRetirementEpochBound field to given value.


### GetPrices

`func (o *ProtocolParameters) GetPrices() Prices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ProtocolParameters) GetPricesOk() (*Prices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ProtocolParameters) SetPrices(v Prices)`

SetPrices sets Prices field to given value.


### GetProtocolVersion

`func (o *ProtocolParameters) GetProtocolVersion() Version`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *ProtocolParameters) GetProtocolVersionOk() (*Version, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *ProtocolParameters) SetProtocolVersion(v Version)`

SetProtocolVersion sets ProtocolVersion field to given value.


### GetStakeKeyDeposit

`func (o *ProtocolParameters) GetStakeKeyDeposit() int64`

GetStakeKeyDeposit returns the StakeKeyDeposit field if non-nil, zero value otherwise.

### GetStakeKeyDepositOk

`func (o *ProtocolParameters) GetStakeKeyDepositOk() (*int64, bool)`

GetStakeKeyDepositOk returns a tuple with the StakeKeyDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeKeyDeposit

`func (o *ProtocolParameters) SetStakeKeyDeposit(v int64)`

SetStakeKeyDeposit sets StakeKeyDeposit field to given value.


### GetTreasuryExpansion

`func (o *ProtocolParameters) GetTreasuryExpansion() string`

GetTreasuryExpansion returns the TreasuryExpansion field if non-nil, zero value otherwise.

### GetTreasuryExpansionOk

`func (o *ProtocolParameters) GetTreasuryExpansionOk() (*string, bool)`

GetTreasuryExpansionOk returns a tuple with the TreasuryExpansion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreasuryExpansion

`func (o *ProtocolParameters) SetTreasuryExpansion(v string)`

SetTreasuryExpansion sets TreasuryExpansion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


