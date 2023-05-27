# EpochInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlkCount** | **int32** | Total blocks in the epoch | 
**EndTime** | **int64** | UNIX timestamp when the epoch ended | 
**EpochNo** | **int32** | Epoch number | 
**Fees** | **string** | Total fees collected in the epoch | 
**StartTime** | **int64** | UNIX timestamp when the epoch began | 
**TxCount** | **int32** | Total transactions in the epoch | 

## Methods

### NewEpochInfo

`func NewEpochInfo(blkCount int32, endTime int64, epochNo int32, fees string, startTime int64, txCount int32, ) *EpochInfo`

NewEpochInfo instantiates a new EpochInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpochInfoWithDefaults

`func NewEpochInfoWithDefaults() *EpochInfo`

NewEpochInfoWithDefaults instantiates a new EpochInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlkCount

`func (o *EpochInfo) GetBlkCount() int32`

GetBlkCount returns the BlkCount field if non-nil, zero value otherwise.

### GetBlkCountOk

`func (o *EpochInfo) GetBlkCountOk() (*int32, bool)`

GetBlkCountOk returns a tuple with the BlkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlkCount

`func (o *EpochInfo) SetBlkCount(v int32)`

SetBlkCount sets BlkCount field to given value.


### GetEndTime

`func (o *EpochInfo) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *EpochInfo) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *EpochInfo) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.


### GetEpochNo

`func (o *EpochInfo) GetEpochNo() int32`

GetEpochNo returns the EpochNo field if non-nil, zero value otherwise.

### GetEpochNoOk

`func (o *EpochInfo) GetEpochNoOk() (*int32, bool)`

GetEpochNoOk returns a tuple with the EpochNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochNo

`func (o *EpochInfo) SetEpochNo(v int32)`

SetEpochNo sets EpochNo field to given value.


### GetFees

`func (o *EpochInfo) GetFees() string`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *EpochInfo) GetFeesOk() (*string, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *EpochInfo) SetFees(v string)`

SetFees sets Fees field to given value.


### GetStartTime

`func (o *EpochInfo) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EpochInfo) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EpochInfo) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.


### GetTxCount

`func (o *EpochInfo) GetTxCount() int32`

GetTxCount returns the TxCount field if non-nil, zero value otherwise.

### GetTxCountOk

`func (o *EpochInfo) GetTxCountOk() (*int32, bool)`

GetTxCountOk returns a tuple with the TxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxCount

`func (o *EpochInfo) SetTxCount(v int32)`

SetTxCount sets TxCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


