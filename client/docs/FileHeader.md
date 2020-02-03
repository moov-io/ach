# FileHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImmediateOrigin** | **string** | contains the Routing Number of the ACH Operator or sending point that is sending the file. | 
**ImmediateOriginName** | **string** | The name of the ACH operator or sending point that is sending the file. | 
**ImmediateDestination** | **string** | contains the Routing Number of the ACH Operator or receiving point to which the file is being sent | 
**ImmediateDestinationName** | **string** | The name of the ACH or receiving point for which that file is destined. | 
**FileCreationTime** | **string** | The File Creation Date is the date when the file was prepared by an ODFI. (Format HHmm - H&#x3D;Hour, m&#x3D;Minute) | [optional] 
**FileCreationDate** | **string** | The File Creation Time is the time when the file was prepared by an ODFI. (Format YYMMDD - Y&#x3D;Year, M&#x3D;Month, D&#x3D;Day) | [optional] 
**FileIDModifier** | **string** | Incremented value for each file for RDFI&#39;s. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


