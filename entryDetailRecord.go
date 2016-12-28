package ach

// EntryDetailRecord contains the actual transaction data for an individual entry.
// Fields include those designating the entry as a deposit (credit) or
// withdrawal (debit), the transit routing number for the entry recipient’s financial
// institution, the account number (left justify,no zero fill), name, and dollar amount.
type EntryDetailRecord struct {
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

	Addenda AddendaRecord
}

// AddendaRecord provides business transaction information in a machine
// readable format. It is usually formatted according to ANSI, ASC, X12 Standard
type AddendaRecord struct {
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
	// Detail or Corporate Entry Detail Record's trace number. This number is
	// the same as the last seven digits of the trace number of the related
	// Entry Detail Record or Corporate Entry Detail Record.
	EntryDetailSequenceNumber string
}
