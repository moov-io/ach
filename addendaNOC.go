package ach

import (
	"fmt"
	"strings"
)

// AddendaNOC is a Addendumer addenda record format for Notification OF Change(NOC)
// The field contents for Notification of Change Entries must match the field contents of the original Entries
type AddendaNOC struct {
	// RecordType defines the type of record in the block. entryAddendaPos 7
	recordType string
	// TypeCode Addenda types code '98'
	typeCode string
	// ChangeCode field contains a standard code used by an ACH Operator or RDFI to describe the reason for a change Entry.
	// Must exist in changeCodeDict
	ChangeCode string
	// OriginalTrace This field contains the Trace Number as originally included on the forward Entry or Prenotification.
	// The RDFI must include the Original Entry Trace Number in the Addenda Record of an Entry being returned to an ODFI,
	// in the Addenda Record of an NOC, within an Acknowledgment Entry, or with an RDFI request for a copy of an authorization.
	OriginalTrace int
	// OriginalDFI field contains the Receiving DFI Identification (addenda.RDFIIdentification) as originally included on the forward Entry or Prenotification that the RDFI is returning or correcting.
	OriginalDFI int
	// CorrectedData
	CorrectedData string
	// TraceNumber matches the Entry Detail Trace Number of the entry being returned.
	TraceNumber int

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
}

var (
	changeCodeDict = map[string]*changeCode{}

	// Error messages specific to AddendaNOC
	msgAddendaNOCChangeCode    = "found is not a valid addenda Change Code"
	msgAddendaNOCTypeCode      = "is not AddendaNOC type code of 98"
	msgAddendaNOCCorrectedData = "must contain the corrected information corresponding to the Change Code"
)

func init() {
	// populate the changeCode map with lookup values
	changeCodeDict = makeChangeCodeDict()
}

// changeCode holds a change Code, Reason/Title, and Description
// table of return codes exists in Part 4.2 of the NACHA corporate rules and guidelines
type changeCode struct {
	Code, Reason, Description string
}

// NewAddendaNOC returns an reference to an instantiated AddendaNOC with default values
func NewAddendaNOC(params ...AddendaParam) *AddendaNOC {
	addendaNOC := &AddendaNOC{
		recordType: "7",
		typeCode:   "98",
	}
	if len(params) > 0 {
		addendaNOC.ChangeCode = params[0].ChangeCode
		addendaNOC.OriginalTrace = addendaNOC.parseNumField(params[0].OriginalTrace)
		addendaNOC.OriginalDFI = addendaNOC.parseNumField(params[0].OriginalDFI)
		addendaNOC.CorrectedData = params[0].CorrectedData
		addendaNOC.TraceNumber = addendaNOC.parseNumField(params[0].TraceNumber)
	}
	return addendaNOC
}

// Parse takes the input record string and parses the AddendaNOC values
func (addendaNOC *AddendaNOC) Parse(record string) {
	// 1-1 Always "7"
	addendaNOC.recordType = "7"
	// 2-3 Always "98"
	addendaNOC.typeCode = record[1:3]
	// 4-6
	addendaNOC.ChangeCode = record[3:6]
	// 7-21
	addendaNOC.OriginalTrace = addendaNOC.parseNumField(record[6:21])
	// 28-35
	addendaNOC.OriginalDFI = addendaNOC.parseNumField(record[27:35])
	// 36-64
	addendaNOC.CorrectedData = strings.TrimSpace(record[35:64])
	// 80-94
	addendaNOC.TraceNumber = addendaNOC.parseNumField(record[79:94])
}

// String writes the AddendaNOC struct to a 94 character string
func (addendaNOC *AddendaNOC) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v",
		addendaNOC.recordType,
		addendaNOC.TypeCode(),
		addendaNOC.ChangeCode,
		addendaNOC.OriginalTraceField(),
		"      ", //6 char reserved field
		addendaNOC.OriginalDFIField(),
		addendaNOC.CorrectedDataField(),
		"               ", // 15 char reserved field
		addendaNOC.TraceNumberField(),
	)
}

// Validate verifies NACHA rules for AddendaNOC
func (addendaNOC *AddendaNOC) Validate() error {
	if addendaNOC.recordType != "7" {
		msg := fmt.Sprintf(msgRecordType, 7)
		return &FieldError{FieldName: "recordType", Value: addendaNOC.recordType, Msg: msg}
	}
	// Type Code must be 98
	if addendaNOC.typeCode != "98" {
		return &FieldError{FieldName: "TypeCode", Value: addendaNOC.typeCode, Msg: msgAddendaTypeCode}
	}

	// AddendaNOC requires a valid ChangeCode
	_, ok := changeCodeDict[addendaNOC.ChangeCode]
	if !ok {
		return &FieldError{FieldName: "ChangeCode", Value: addendaNOC.ChangeCode, Msg: msgAddendaNOCChangeCode}
	}

	// AddendaNOC Record must contain the corrected information corresponding to the Change Code used
	if addendaNOC.CorrectedData == "" {
		return &FieldError{FieldName: "CorrectedData", Value: addendaNOC.CorrectedData, Msg: msgAddendaNOCCorrectedData}
	}

	return nil
}

// TypeCode defines the format of the underlying addenda record
func (addendaNOC *AddendaNOC) TypeCode() string {
	return addendaNOC.typeCode
}

// OriginalTraceField returns a zero padded OriginalTrace string
func (addendaNOC *AddendaNOC) OriginalTraceField() string {
	return addendaNOC.numericField(addendaNOC.OriginalTrace, 15)
}

// OriginalDFIField returns a zero padded OriginalDFI string
func (addendaNOC *AddendaNOC) OriginalDFIField() string {
	return addendaNOC.numericField(addendaNOC.OriginalDFI, 8)
}

//CorrectedDataField returns a space padded CorrectedData string
func (addendaNOC *AddendaNOC) CorrectedDataField() string {
	return addendaNOC.alphaField(addendaNOC.CorrectedData, 29)
}

// TraceNumberField returns a zero padded traceNumber string
func (addendaNOC *AddendaNOC) TraceNumberField() string {
	return addendaNOC.numericField(addendaNOC.TraceNumber, 15)
}

func makeChangeCodeDict() map[string]*changeCode {
	dict := make(map[string]*changeCode)

	codes := []changeCode{
		{"C01", "Incorrect bank account number", "Bank account number incorrect or formatted incorrectly"},
		{"C02", "Incorrect transit/routing number", "Once valid transit/routing number must be changed"},
		{"C03", "Incorrect transit/routing number and bank account number", "Once valid transit/routing number must be changed and causes a change to bank account number structure"},
		{"C04", "Bank account name change", "Customer has changed name or ODFI submitted name incorrectly"},
		{"C05", "Incorrect payment code", "Entry posted to demand account should contain savings payment codes or vice versa"},
		{"C06", "Incorrect bank account number and transit code", "Bank account number must be changed and payment code should indicate posting to another account type (demand/savings)"},
		{"C07", "Incorrect transit/routing number, bank account number and payment code", "Changes required in three fields indicated"},
		{"C09", "Incorrect individual ID number", "Individual's ID number is incorrect"},
		{"C10", "Incorrect company name", "Company name is no longer valid and should be changed."},
		{"C11", "Incorrect company identification", "Company ID is no longer valid and should be changed"},
		{"C12", "Incorrect company name and company ID", "Both the company name and company id are no longer valid and must be changed"},
	}
	// populate the map
	for _, code := range codes {
		dict[code.Code] = &code
	}
	return dict
}
