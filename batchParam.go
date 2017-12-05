package ach

// BatchParam contains information about the company(Originator) and the type of detail records to follow.
// It is a subset of BatchHeader used for simplifying the client api build process.
type BatchParam struct {
	// ServiceClassCode a three digit code identifies:
	// 	- 200 mixed debits and credits
	// 	- 220 credits only
	// 	- 225 debits only
	ServiceClassCode string `json:"service_class_code"`
	// CompanyName is the legal company name making the transaction.
	CompanyName string `json:"company_name"`
	// CompanyIdentification is assigned by your bank to identify your company. Frequently the federal tax ID
	CompanyIdentification string `json:"company_identification"`
	// StandardEntryClass identifies the payment type (product) found within the batch using a 3-character code
	StandardEntryClass string `json:"standard_entry_class"`
	// CompanyEntryDescription describes the transaction. For example "PAYROLL"
	CompanyEntryDescription string `json:"company_entry_description"`
	// CompanyDescriptiveDate a date chosen to identify the transactions in YYMMDD format.
	CompanyDescriptiveDate string `json:"company_descriptive_date"`
	// Date transactions are to be posted to the receivers’ account in YYMMDD format.
	EffectiveEntryDate string `json:"effective_entry_date"`
	// ODFIIdentification originating ODFI's routing number without the last digit
	ODFIIdentification string `json:"ODFI_identification"`
}
