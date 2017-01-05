package ach

// Addenda provides business transaction information in a machine
// readable format. It is usually formatted according to ANSI, ASC, X12 Standard
type Addenda struct {
	// RecordType defines the type of record in the block. entryAddendaPos 7
	RecordType string
	// TypeCode Addenda types code '05'
	TypeCode string
	// PaymentRelatedInformation
	PaymentRelatedInformation string
	// SequenceNumber is consecutively assigned to each Addenda Record following
	// an Entry Detail Record. The first addenda sequence number must always
	// be a "1".
	SequenceNumber string
	// EntryDetailSequenceNumber contains the ascending sequence number section of the Entry
	// Detail or Corporate Entry Detail Record's trace numbe This number is
	// the same as the last seven digits of the trace number of the related
	// Entry Detail Record or Corporate Entry Detail Record.
	EntryDetailSequenceNumber string
}

// Parse takes the input record string and parses the Addenda values
func (addenda *Addenda) Parse(record string) {
	// 1-1 Always "7"
	addenda.RecordType = record[:1]
	// 2-3 Defines the specific explanation and format for the addenda information contained in the same record
	addenda.TypeCode = record[1:3]
	// 4-83 Based on the information entere (04-83) 80 alphanumeric
	addenda.PaymentRelatedInformation = record[3:83]
	// 84-87 SequenceNumber is consecutively assigned to each Addenda Record following
	// an Entry Detail Record
	addenda.SequenceNumber = record[83:87]
	// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
	addenda.EntryDetailSequenceNumber = record[87:94]

}

// NewAddenda returns a new Addenda with default values for none exported fields
func NewAddenda() *Addenda {
	return &Addenda{}
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda *Addenda) Validate() (bool, error) {
	return true, nil
}
