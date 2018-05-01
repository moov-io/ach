package ach

import (
	"fmt"
	"strings"
)

// Addenda98 is a Addendumer addenda record format for Notification OF Change(98)
// The field contents for Notification of Change Entries must match the field contents of the original Entries
type Addenda98 struct {
	// RecordType defines the type of record in the block. entryAddendaPos 7
	recordType string
	// TypeCode Addenda types code '98'
	typeCode string
	// ChangeCode field contains a standard code used by an ACH Operator or RDFI to describe the reason for a change Entry.
	// Must exist in changeCodeDict
	ChangeCode string
	// OriginalTrace This field contains the Trace Number as originally included on the forward Entry or Prenotification.
	// The RDFI must include the Original Entry Trace Number in the Addenda Record of an Entry being returned to an ODFI,
	// in the Addenda Record of an 98, within an Acknowledgment Entry, or with an RDFI request for a copy of an authorization.
	OriginalTrace int
	// OriginalDFI field contains the Receiving DFI Identification (addenda.RDFIIdentification) as originally included on the forward Entry or Prenotification that the RDFI is returning or correcting.
	OriginalDFI string
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

	// Error messages specific to Addenda98
	msgAddenda98ChangeCode    = "found is not a valid addenda Change Code"
	msgAddenda98TypeCode      = "is not Addenda98 type code of 98"
	msgAddenda98CorrectedData = "must contain the corrected information corresponding to the Change Code"
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

// NewAddenda98 returns an reference to an instantiated Addenda98 with default values
func NewAddenda98() *Addenda98 {
	addenda98 := &Addenda98{
		recordType: "7",
		typeCode:   "98",
	}
	return addenda98
}

// Parse takes the input record string and parses the Addenda98 values
func (addenda98 *Addenda98) Parse(record string) {
	// 1-1 Always "7"
	addenda98.recordType = "7"
	// 2-3 Always "98"
	addenda98.typeCode = record[1:3]
	// 4-6
	addenda98.ChangeCode = record[3:6]
	// 7-21
	addenda98.OriginalTrace = addenda98.parseNumField(record[6:21])
	// 28-35
	addenda98.OriginalDFI = addenda98.parseStringField(record[27:35])
	// 36-64
	addenda98.CorrectedData = strings.TrimSpace(record[35:64])
	// 80-94
	addenda98.TraceNumber = addenda98.parseNumField(record[79:94])
}

// String writes the Addenda98 struct to a 94 character string
func (addenda98 *Addenda98) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v",
		addenda98.recordType,
		addenda98.TypeCode(),
		addenda98.ChangeCode,
		addenda98.OriginalTraceField(),
		"      ", //6 char reserved field
		addenda98.OriginalDFIField(),
		addenda98.CorrectedDataField(),
		"               ", // 15 char reserved field
		addenda98.TraceNumberField(),
	)
}

// Validate verifies NACHA rules for Addenda98
func (addenda98 *Addenda98) Validate() error {
	if addenda98.recordType != "7" {
		msg := fmt.Sprintf(msgRecordType, 7)
		return &FieldError{FieldName: "recordType", Value: addenda98.recordType, Msg: msg}
	}
	// Type Code must be 98
	if addenda98.typeCode != "98" {
		return &FieldError{FieldName: "TypeCode", Value: addenda98.typeCode, Msg: msgAddendaTypeCode}
	}

	// Addenda98 requires a valid ChangeCode
	_, ok := changeCodeDict[addenda98.ChangeCode]
	if !ok {
		return &FieldError{FieldName: "ChangeCode", Value: addenda98.ChangeCode, Msg: msgAddenda98ChangeCode}
	}

	// Addenda98 Record must contain the corrected information corresponding to the Change Code used
	if addenda98.CorrectedData == "" {
		return &FieldError{FieldName: "CorrectedData", Value: addenda98.CorrectedData, Msg: msgAddenda98CorrectedData}
	}

	return nil
}

// TypeCode defines the format of the underlying addenda record
func (addenda98 *Addenda98) TypeCode() string {
	return addenda98.typeCode
}

// OriginalTraceField returns a zero padded OriginalTrace string
func (addenda98 *Addenda98) OriginalTraceField() string {
	return addenda98.numericField(addenda98.OriginalTrace, 15)
}

// OriginalDFIField returns a zero padded OriginalDFI string
func (addenda98 *Addenda98) OriginalDFIField() string {
	return addenda98.stringRTNField(addenda98.OriginalDFI, 8)
}

//CorrectedDataField returns a space padded CorrectedData string
func (addenda98 *Addenda98) CorrectedDataField() string {
	return addenda98.alphaField(addenda98.CorrectedData, 29)
}

// TraceNumberField returns a zero padded traceNumber string
func (addenda98 *Addenda98) TraceNumberField() string {
	return addenda98.numericField(addenda98.TraceNumber, 15)
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
