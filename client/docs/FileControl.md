# FileControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Moov API File ID | [optional] 
**BatchCount** | **int32** | Count of Batches in the File | [optional] 
**BlockCount** | **int32** | BlockCount total number of records in the file (include all headers and trailer) divided by 10 (This number must be evenly divisible by 10. If not, additional records consisting of all 9’s are added to the file after the initial ‘9’ record to fill out the block 10.)  | [optional] 
**EntryAddendaCount** | **int32** | Total detail and addenda records in the file | [optional] 
**EntryHash** | **int32** | EntryHash calculated in the same manner as the batch has total but includes total from entire file | [optional] 
**TotalDebit** | **int32** | Accumulated Batch debit totals within the file. | [optional] 
**TotalCredit** | **int32** | Accumulated Batch credit totals within the file. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


