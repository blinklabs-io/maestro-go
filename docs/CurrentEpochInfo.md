# CurrentEpochInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlkCount** | **int32** | Total blocks in the epoch so far | 
**EpochNo** | **int32** | Epoch number | 
**Fees** | **string** | Total fees collected in the epoch so far | 
**StartTime** | **int64** | UNIX timestamp when the epoch began | 
**TxCount** | **int32** | Total transactions in the epoch so far | 

## Methods

### NewCurrentEpochInfo

`func NewCurrentEpochInfo(blkCount int32, epochNo int32, fees string, startTime int64, txCount int32, ) *CurrentEpochInfo`

NewCurrentEpochInfo instantiates a new CurrentEpochInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentEpochInfoWithDefaults

`func NewCurrentEpochInfoWithDefaults() *CurrentEpochInfo`

NewCurrentEpochInfoWithDefaults instantiates a new CurrentEpochInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlkCount

`func (o *CurrentEpochInfo) GetBlkCount() int32`

GetBlkCount returns the BlkCount field if non-nil, zero value otherwise.

### GetBlkCountOk

`func (o *CurrentEpochInfo) GetBlkCountOk() (*int32, bool)`

GetBlkCountOk returns a tuple with the BlkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlkCount

`func (o *CurrentEpochInfo) SetBlkCount(v int32)`

SetBlkCount sets BlkCount field to given value.


### GetEpochNo

`func (o *CurrentEpochInfo) GetEpochNo() int32`

GetEpochNo returns the EpochNo field if non-nil, zero value otherwise.

### GetEpochNoOk

`func (o *CurrentEpochInfo) GetEpochNoOk() (*int32, bool)`

GetEpochNoOk returns a tuple with the EpochNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochNo

`func (o *CurrentEpochInfo) SetEpochNo(v int32)`

SetEpochNo sets EpochNo field to given value.


### GetFees

`func (o *CurrentEpochInfo) GetFees() string`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *CurrentEpochInfo) GetFeesOk() (*string, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *CurrentEpochInfo) SetFees(v string)`

SetFees sets Fees field to given value.


### GetStartTime

`func (o *CurrentEpochInfo) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CurrentEpochInfo) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CurrentEpochInfo) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.


### GetTxCount

`func (o *CurrentEpochInfo) GetTxCount() int32`

GetTxCount returns the TxCount field if non-nil, zero value otherwise.

### GetTxCountOk

`func (o *CurrentEpochInfo) GetTxCountOk() (*int32, bool)`

GetTxCountOk returns a tuple with the TxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxCount

`func (o *CurrentEpochInfo) SetTxCount(v int32)`

SetTxCount sets TxCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


