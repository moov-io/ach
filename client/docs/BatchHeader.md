# BatchHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Batch Header ID | [optional] 
**ServiceClassCode** | **int32** | Service Class Code - ACH Credits Only 220 and ACH Debits Only 225 | 
**CompanyName** | **string** | Company originating the entries in the batch | 
**CompanyDiscretionaryData** | **string** | The 9 digit FEIN number (proceeded by a predetermined alpha or numeric character) of the entity in the company name field | [optional] 
**StandardEntryClassCode** | **string** | Identifies the payment type (product) found within an ACH batch-using a 3-character code. | [optional] 
**CompanyEntryDescription** | **string** | A description of the entries contained in the batch. The Originator establishes the value of this field to provide a description of the purpose of the entry to be displayed back to the receive For example, \&quot;GAS BILL,\&quot; \&quot;REG. SALARY,\&quot; \&quot;INS. PREM,\&quot;, \&quot;SOC. SEC.,\&quot; \&quot;DTC,\&quot; \&quot;TRADE PAY,\&quot; \&quot;PURCHASE,\&quot; etc. This field must contain the word \&quot;REVERSAL\&quot; (left justified) when the batch contains reversing entries. This field must contain the word \&quot;RECLAIM\&quot; (left justified) when the batch contains reclamation entries. This field must contain the word \&quot;NONSETTLED\&quot; (left justified) when the batch contains entries which could not settle.  | [optional] 
**CompanyDescriptiveDate** | **string** | The Originator establishes this field as the date it would like to see displayed to the receiver for descriptive purposes. This field is never used to control timing of any computer or manual operation. It is solely for descriptive purposes. The RDFI should not assume any specific format.  | [optional] 
**EffectiveEntryDate** | **string** | Date on which the entries are to settle. Format YYMMDD (Y&#x3D;Year, M&#x3D;Month, D&#x3D;Day) | [optional] 
**OriginatorStatusCode** | **int32** | ODFI initiating the Entry. 0 ADV File prepared by an ACH Operator. 1 This code identifies the Originator as a depository financial institution. 2 This code identifies the Originator as a Federal Government entity or agency.  | [optional] 
**ODFIIdentification** | **string** | First 8 digits of the originating DFI transit routing number | 
**BatchNumber** | **string** | BatchNumber is assigned in ascending sequence to each batch by the ODFI or its Sending Point in a given file of entries. Since the batch number in the Batch Header Record and the Batch Control Record is the same, the ascending sequence number should be assigned by batch and not by record.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


