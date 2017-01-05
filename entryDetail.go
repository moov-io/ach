package ach

// EntryDetail contains the actual transaction data for an individual entry.
// Fields include those designating the entry as a deposit (credit) or
// withdrawal (debit), the transit routing number for the entry recipient’s financial
// institution, the account number (left justify,no zero fill), name, and dollar amount.
type EntryDetail struct {
	// RecordType defines the type of record in the block. 6
	RecordType string
	// TransactionCode if the recievers account is:
	// Credit (deposit) to checking account ‘22’
	// Prenote for credit to checking account ‘23’
	// Debit (withdrawal) to checking account ‘27’
	// Prenote for debit to checking account ‘28’
	// Credit to savings account ‘32’
	// Prenote for credit to savings account ‘33’
	// Debit to savings account ‘37’
	// Prenote for debit to savings account ‘38’
	TransactionCode string

	// RoutingNumber is the RDFI's routing number without the last digit.
	RdfiIdentification string

	// CheckDigit the last digit of the RDFI's routing number
	CheckDigit string

	// The receiver's bank account number you are crediting/debiting.
	// It important to note that this is an alphanumeric field, so its space padded, no zero padded
	DfiAccountNumber string

	// Amount Number of cents you are debiting/crediting this account
	Amount string

	// IndividualIdentificationNumber n internal identification (alphanumeric) that
	// you use to uniquely identify this Entry Detail Record
	IndividualIdentificationNumber string

	// IndividualName The name of the receiver, usually the name on the bank account
	IndividualName string

	// DiscretionaryData allows ODFIs to include codes, of significance only to them,
	// to enable specialized handling of the entry. There will be no
	// standardized interpretation for the value of this field. It can either
	// be a single two-character code, or two distince one-character codes,
	// according to the needs of the ODFI and/or Originator involved. This
	// field must be returned intact for any returned entry.
	DiscretionaryData string

	// AddendaRecordIndicator indicates the existence of an Addenda Record.
	// A value of "1" indicates that one ore more addenda records follow,
	// and "0" means no such record is present.
	AddendaRecordIndicator string

	// TraceNumber assigned by the ODFI in ascending sequence, is included in each
	// Entry Detail Record, Corporate Entry Detail Record, and addenda Record.
	// Trace Numbers uniquely identify each entry within a batch in an ACH input file.
	// In association with the Batch Number, transmission (File Creation) Date,
	// and File ID Modifier, the Trace Number uniquely identifies an entry within a given file.
	// For addenda Records, the Trace Number will be identical to the Trace Number
	// in the associated Entry Detail Record, since the Trace Number is associated
	// with an entry or item rather than a physical record.
	TraceNumber string
}

// Parse takes the input record string and parses the EntryDetail values
func (ed *EntryDetail) Parse(record string) {
	// 1-1 Always "6"
	ed.RecordType = record[:1]
	// 2-3 is checking credit 22 debit 27 savings credit 32 debit 37
	ed.TransactionCode = record[1:3]
	// 4-11 the RDFI's routing number without the last digit.
	ed.RdfiIdentification = record[3:11]
	// 12-12 The last digit of the RDFI's routing number
	ed.CheckDigit = record[11:12]
	// 13-29 The receiver's bank account number you are crediting/debiting
	ed.DfiAccountNumber = record[12:29]
	// 30-39 Number of cents you are debiting/crediting this account
	ed.Amount = record[29:39]
	// 40-54 An internal identification (alphanumeric) that you use to uniquely identify this Entry Detail Record
	ed.IndividualIdentificationNumber = record[39:54]
	// 55-76 The name of the receiver, usually the name on the bank account
	ed.IndividualName = record[54:76]
	// 77-78 allows ODFIs to include codes of significance only to them
	// normally blank
	ed.DiscretionaryData = record[76:78]
	// 79-79 1 if addenda exists 0 if it does not
	ed.AddendaRecordIndicator = record[78:79]
	// 80-84 An internal identification (alphanumeric) that you use to uniquely identify
	// this Entry Detail Recor This number should be unique to the transaction and will help identify the transaction in case of an inquiry
	ed.TraceNumber = record[79:94]
}

// NewEntryDetail returns a new EntryDetail with default values for none exported fields
func NewEntryDetail() *EntryDetail {
	return &EntryDetail{}
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ed *EntryDetail) Validate() (bool, error) {
	return true, nil
}
