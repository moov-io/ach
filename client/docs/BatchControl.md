# BatchControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Batch ID | [optional] 
**ServiceClassCode** | **int32** | Same as ServiceClassCode in BatchHeaderRecord | [optional] 
**EntryAddendaCount** | **int32** | EntryAddendaCount is a tally of each Entry Detail Record and each Addenda Record processed, within either the batch or file as appropriate. | [optional] 
**EntryHash** | **int32** | Validate the Receiving DFI Identification in each Entry Detail Record is hashed to provide a check against inadvertent alteration of data contents due to hardware failure or program error. In this context the Entry Hash is the sum of the corresponding fields in the Entry Detail Records on the file.  | [optional] 
**TotalDebit** | **int32** | Contains accumulated Entry debit totals within the batch. | [optional] 
**TotalCredit** | **int32** | Contains accumulated Entry credit totals within the batch. | [optional] 
**CompanyIdentification** | **string** | Alphanumeric code used to identify an Originator The Company Identification Field must be included on all prenotification records and on each entry initiated pursuant to such prenotification. The Company ID may begin with the ANSI one-digit Identification Code Designator (ICD), followed by the identification number The ANSI Identification Numbers and related Identification Code IRS Employer Identification Number (EIN) \&quot;1\&quot; Data Universal Numbering Systems (DUNS) \&quot;3\&quot; User Assigned Number \&quot;9\&quot;  | [optional] 
**MessageAuthentication** | **string** | MAC is an eight character code derived from a special key used in conjunction with the DES algorithm. The purpose of the MAC is to validate the authenticity of ACH entries. The DES algorithm and key message standards must be in accordance with standards adopted by the American National Standards Institute. The remaining eleven characters of this field are blank. | [optional] 
**ODFIIdentification** | **string** | The routing number is used to identify the DFI originating entries within a given branch. | [optional] 
**BatchNumber** | **string** | BatchNumber is assigned in ascending sequence to each batch by the ODFI or its Sending Point in a given file of entries. Since the batch number in the Batch Header Record and the Batch Control Record is the same, the ascending sequence number should be assigned by batch and not by record. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


