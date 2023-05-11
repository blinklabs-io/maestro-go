# AssetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetName** | **string** | Hex encoding of the asset name | 
**AssetNameAscii** | Pointer to **string** | ASCII representation of the asset name | [optional] 
**AssetStandards** | [**AssetStandards**](AssetStandards.md) |  | 
**BurnTxCount** | **int64** | Number of transactions which burned some of the asset | 
**Fingerprint** | **string** | CIP-14 fingerprint of the asset | 
**FirstMintTime** | **int32** | UNIX timestamp of the first mint transaction | 
**FirstMintTx** | **string** | Transaction hash of the first transaction which minted the asset | 
**LatestMintTxMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**MintTxCount** | **int64** | Number of transactions which minted some of the asset | 
**TokenRegistryMetadata** | Pointer to [**TokenRegistryMetadata**](TokenRegistryMetadata.md) |  | [optional] 
**TotalSupply** | **int64** | Current amount of the asset minted | 

## Methods

### NewAssetInfo

`func NewAssetInfo(assetName string, assetStandards AssetStandards, burnTxCount int64, fingerprint string, firstMintTime int32, firstMintTx string, mintTxCount int64, totalSupply int64, ) *AssetInfo`

NewAssetInfo instantiates a new AssetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetInfoWithDefaults

`func NewAssetInfoWithDefaults() *AssetInfo`

NewAssetInfoWithDefaults instantiates a new AssetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetName

`func (o *AssetInfo) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *AssetInfo) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *AssetInfo) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.


### GetAssetNameAscii

`func (o *AssetInfo) GetAssetNameAscii() string`

GetAssetNameAscii returns the AssetNameAscii field if non-nil, zero value otherwise.

### GetAssetNameAsciiOk

`func (o *AssetInfo) GetAssetNameAsciiOk() (*string, bool)`

GetAssetNameAsciiOk returns a tuple with the AssetNameAscii field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetNameAscii

`func (o *AssetInfo) SetAssetNameAscii(v string)`

SetAssetNameAscii sets AssetNameAscii field to given value.

### HasAssetNameAscii

`func (o *AssetInfo) HasAssetNameAscii() bool`

HasAssetNameAscii returns a boolean if a field has been set.

### GetAssetStandards

`func (o *AssetInfo) GetAssetStandards() AssetStandards`

GetAssetStandards returns the AssetStandards field if non-nil, zero value otherwise.

### GetAssetStandardsOk

`func (o *AssetInfo) GetAssetStandardsOk() (*AssetStandards, bool)`

GetAssetStandardsOk returns a tuple with the AssetStandards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetStandards

`func (o *AssetInfo) SetAssetStandards(v AssetStandards)`

SetAssetStandards sets AssetStandards field to given value.


### GetBurnTxCount

`func (o *AssetInfo) GetBurnTxCount() int64`

GetBurnTxCount returns the BurnTxCount field if non-nil, zero value otherwise.

### GetBurnTxCountOk

`func (o *AssetInfo) GetBurnTxCountOk() (*int64, bool)`

GetBurnTxCountOk returns a tuple with the BurnTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnTxCount

`func (o *AssetInfo) SetBurnTxCount(v int64)`

SetBurnTxCount sets BurnTxCount field to given value.


### GetFingerprint

`func (o *AssetInfo) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *AssetInfo) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *AssetInfo) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetFirstMintTime

`func (o *AssetInfo) GetFirstMintTime() int32`

GetFirstMintTime returns the FirstMintTime field if non-nil, zero value otherwise.

### GetFirstMintTimeOk

`func (o *AssetInfo) GetFirstMintTimeOk() (*int32, bool)`

GetFirstMintTimeOk returns a tuple with the FirstMintTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMintTime

`func (o *AssetInfo) SetFirstMintTime(v int32)`

SetFirstMintTime sets FirstMintTime field to given value.


### GetFirstMintTx

`func (o *AssetInfo) GetFirstMintTx() string`

GetFirstMintTx returns the FirstMintTx field if non-nil, zero value otherwise.

### GetFirstMintTxOk

`func (o *AssetInfo) GetFirstMintTxOk() (*string, bool)`

GetFirstMintTxOk returns a tuple with the FirstMintTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMintTx

`func (o *AssetInfo) SetFirstMintTx(v string)`

SetFirstMintTx sets FirstMintTx field to given value.


### GetLatestMintTxMetadata

`func (o *AssetInfo) GetLatestMintTxMetadata() map[string]interface{}`

GetLatestMintTxMetadata returns the LatestMintTxMetadata field if non-nil, zero value otherwise.

### GetLatestMintTxMetadataOk

`func (o *AssetInfo) GetLatestMintTxMetadataOk() (*map[string]interface{}, bool)`

GetLatestMintTxMetadataOk returns a tuple with the LatestMintTxMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestMintTxMetadata

`func (o *AssetInfo) SetLatestMintTxMetadata(v map[string]interface{})`

SetLatestMintTxMetadata sets LatestMintTxMetadata field to given value.

### HasLatestMintTxMetadata

`func (o *AssetInfo) HasLatestMintTxMetadata() bool`

HasLatestMintTxMetadata returns a boolean if a field has been set.

### GetMintTxCount

`func (o *AssetInfo) GetMintTxCount() int64`

GetMintTxCount returns the MintTxCount field if non-nil, zero value otherwise.

### GetMintTxCountOk

`func (o *AssetInfo) GetMintTxCountOk() (*int64, bool)`

GetMintTxCountOk returns a tuple with the MintTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintTxCount

`func (o *AssetInfo) SetMintTxCount(v int64)`

SetMintTxCount sets MintTxCount field to given value.


### GetTokenRegistryMetadata

`func (o *AssetInfo) GetTokenRegistryMetadata() TokenRegistryMetadata`

GetTokenRegistryMetadata returns the TokenRegistryMetadata field if non-nil, zero value otherwise.

### GetTokenRegistryMetadataOk

`func (o *AssetInfo) GetTokenRegistryMetadataOk() (*TokenRegistryMetadata, bool)`

GetTokenRegistryMetadataOk returns a tuple with the TokenRegistryMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRegistryMetadata

`func (o *AssetInfo) SetTokenRegistryMetadata(v TokenRegistryMetadata)`

SetTokenRegistryMetadata sets TokenRegistryMetadata field to given value.

### HasTokenRegistryMetadata

`func (o *AssetInfo) HasTokenRegistryMetadata() bool`

HasTokenRegistryMetadata returns a boolean if a field has been set.

### GetTotalSupply

`func (o *AssetInfo) GetTotalSupply() int64`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *AssetInfo) GetTotalSupplyOk() (*int64, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *AssetInfo) SetTotalSupply(v int64)`

SetTotalSupply sets TotalSupply field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


