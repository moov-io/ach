// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"unicode/utf8"
)

var (
	upperAlphanumericRegex = regexp.MustCompile(`[^ A-Z0-9!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	alphanumericRegex      = regexp.MustCompile(`[^ \w!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	// Errors specific to validation
	msgAlphanumeric        = "has non alphanumeric characters"
	msgUpperAlpha          = "is not uppercase A-Z or 0-9"
	msgFieldInclusion      = "is a mandatory field and has a default value"
	msgFieldRequired       = "is a required field"
	msgValidFieldLength    = "is not length %d"
	msgServiceClass        = "is an invalid Service Class Code"
	msgSECCode             = "is an invalid Standard Entry Class Code"
	msgOrigStatusCode      = "is an invalid Originator Status Code"
	msgAddendaTypeCode     = "is an invalid Addenda Type Code"
	msgTransactionCode     = "is an invalid Transaction Code"
	msgValidCheckDigit     = "does not match calculated check digit %d"
	msgCardTransactionType = "is an invalid Card Transaction Type"
	msgValidMonth          = "is an invalid month"
	msgValidDay            = "is an invalid day"
	msgValidYear           = "is an invalid year"
	// IAT
	msgForeignExchangeIndicator          = "is an invalid Foreign Exchange Indicator"
	msgForeignExchangeReferenceIndicator = "is an invalid Foreign Exchange Reference Indicator"
	msgTransactionTypeCode               = "is an invalid Addenda10 Transaction Type Code"
	msgIDNumberQualifier                 = "is an invalid Identification Number Qualifier"
)

// validator is common validation and formatting of golang types to ach type strings
type validator struct{}

// FieldError is returned for errors at a field level in a record
type FieldError struct {
	FieldName string // field name where error happened
	Value     string // value that cause error
	Msg       string // context of the error.
}

// Error message is constructed
// FieldName Msg Value
// Example1: BatchCount $% has none alphanumeric characters
// Example2: BatchCount 5 is out-of-balance with file count 6
func (e *FieldError) Error() string {
	return fmt.Sprintf("%s %s %s", e.FieldName, e.Value, e.Msg)
}

// isCardTransactionType ensures card transaction type of a batchPOS is valid
func (v *validator) isCardTransactionType(code string) error {
	switch code {
	case
		// Purchase of goods or services
		"01",
		// Cash
		"02",
		// Return Reversal
		"03",
		// Purchase Reversal
		"11",
		// Cash Reversal
		"12",
		// Return
		"13",
		// Adjustment
		"21",
		// Miscellaneous Transaction
		"99":
		return nil
	}
	return errors.New(msgCardTransactionType)
}

// isYear validates a 2 digit year 18-50 (2018 - 2050)
// ToDo:  Add/remove more years as card expiration dates need/don't need them
func (v *validator) isYear(s string) error {
	switch s {
	case
		"18", "19",
		"20", "21", "22", "23", "24", "25", "26", "27", "28", "29",
		"30", "31", "32", "33", "34", "35", "36", "37", "38", "39",
		"40", "41", "42", "43", "44", "45", "46", "47", "48", "49",
		"50":
		return nil
	}
	return errors.New(msgValidYear)
}

// isMonth validates a 2 digit month 01-12
func (v *validator) isMonth(s string) error {
	switch s {
	case
		"01", "02", "03", "04", "05", "06",
		"07", "08", "09", "10", "11", "12":
		return nil
	}
	return errors.New(msgValidMonth)
}

// isDay validates a 2 digit day based on a 2 digit month
// month 01-12 day 01-31 based on month
func (v *validator) isDay(m string, d string) error {
	// ToDo: Future Consideration Leap Year - not sure if cards actually have 0229
	switch m {
	// February
	case "02":
		switch d {
		case
			"01", "02", "03", "04", "05", "06",
			"07", "08", "09", "10", "11", "12",
			"13", "14", "15", "16", "17", "18",
			"19", "20", "21", "22", "23", "24",
			"25", "26", "27", "28", "29":
			return nil
		}
	// April, June, September, November
	case "04", "06", "09", "11":
		switch d {
		case
			"01", "02", "03", "04", "05", "06",
			"07", "08", "09", "10", "11", "12",
			"13", "14", "15", "16", "17", "18",
			"19", "20", "21", "22", "23", "24",
			"25", "26", "27", "28", "29", "30":
			return nil
		}
	// January, March, May, July, August, October, December
	case "01", "03", "05", "07", "08", "10", "12":
		switch d {
		case
			"01", "02", "03", "04", "05", "06",
			"07", "08", "09", "10", "11", "12",
			"13", "14", "15", "16", "17", "18",
			"19", "20", "21", "22", "23", "24",
			"25", "26", "27", "28", "29", "30", "31":
			return nil
		}
	}
	return errors.New(msgValidDay)
}

// isForeignExchangeIndicator ensures foreign exchange indicators of an
// IATBatchHeader is valid
func (v *validator) isForeignExchangeIndicator(code string) error {
	switch code {
	case
		"FV", "VF", "FF":
		return nil
	}
	return errors.New(msgForeignExchangeIndicator)
}

// isForeignExchangeReferenceIndicator ensures foreign exchange reference
// indicator of am IATBatchHeader is valid
func (v *validator) isForeignExchangeReferenceIndicator(code int) error {
	switch code {
	case
		1, 2, 3:
		return nil
	}
	return errors.New(msgForeignExchangeReferenceIndicator)
}

// isIDNumberQualifier ensures ODFI Identification Number Qualifier is valid
// For Inbound IATs: The 2-digit code that identifies the numbering scheme used in the
// Foreign DFI Identification Number field:
// 01 = National Clearing System
// 02 = BIC Code
// 03 = IBAN Code
// used for both ODFIIDNumberQualifier and RDFIIDNumberQualifier
func (v *validator) isIDNumberQualifier(s string) error {
	switch s {
	case
		"01", "02", "03":
		return nil
	}
	return errors.New(msgIDNumberQualifier)
}

// isOriginatorStatusCode ensures status code of a batch is valid
func (v *validator) isOriginatorStatusCode(code int) error {
	switch code {
	case
		// ADV file - prepared by an ACH Operator
		0,
		//Originator is a financial institution
		1,
		// Originator is a Government Agency or other agency not subject to ACH Rules
		2:
		return nil
	}
	return errors.New(msgOrigStatusCode)
}

// isSECCode returns true if a SEC Code of a Batch is found
func (v *validator) isSECCode(code string) error {
	switch code {
	case
		ACK, ADV, ARC, ATX, BOC, CCD, CIE, COR, CTX, DNE, ENR,
		IAT, MTE, POS, PPD, POP, RCK, SHR, TEL, TRC, TRX, WEB, XCK:
		return nil
	}
	return errors.New(msgSECCode)
}

// iServiceClass returns true if a valid service class code of a batch is found
func (v *validator) isServiceClass(code int) error {
	switch code {
	case
		// Mixed Debits and Credits
		MixedDebitsAndCredits,
		// Credits Only
		CreditsOnly,
		// Debits Only
		DebitsOnly,
		// Automated Accounting Advices
		AutomatedAccountingAdvices:
		return nil
	}
	return errors.New(msgServiceClass)
}

// isTypeCode returns true if a valid type code of an Addendum is found
//
// The Addenda Type Code defines the specific interpretation and format for the addenda information contained in the Entry.
func (v *validator) isTypeCode(code string) error {
	switch code {
	case
		// For POS, SHR or MTE Entries
		"02",
		// Addenda Record
		"08",
		// Notification of Change and Refused Notification of Change Entry
		"98",
		// Return, Dishonored Return and Contested Dishonored Return Entries
		"99",
		//  IAT forward Entries and IAT Returns
		"10", "11", "12", "13", "14", "15", "16", "17", "18",
		// ACK, ATX, CCD, CIE, CTX, DNE, ENR, PPD, TRX and WEB Entries
		"05":
		return nil
	}
	return errors.New(msgAddendaTypeCode)
}

// isTransactionCode ensures TransactionCode of an Entry is valid
//
// The Tran Code is a two-digit code in positions 2 - 3 of the Entry Detail Record (6 Record) within an ACH File.
// The first digit of the Tran Code indicates the account type to which the entry will post, where the number:
//	"2" designates a Checking Account.
//	"3" designates a Savings Account.
// 	"4" designates a General Ledger Account.
// 	"5" designates Loan Account.
//The second digit of the Tran Code identifies the entry as:
//	an original forward entry, where the number:
//		"2" designates a credit. or
//		"7" designates a debit.
//	a return or NOC, where the number:
//		"1" designates the return/NOC of a credit, or
//		"6" designates a return/NOC of a debit.
//	a pre-note or non-monetary informational transaction, where the number:
//		"3" designates a credit, or
//		"8" designates a debit.
func (v *validator) isTransactionCode(code int) error {
	switch code {
	// TransactionCode if the receivers account is:
	case
		// Demand Credit Records (for checking, NOW, and share draft accounts)

		// Automated Return or Notification of Change for original transaction code '22', '23, '24'
		CheckingReturnNOCCredit,
		// Credit (deposit) to checking account ‘22’
		CheckingCredit,
		// Prenote for credit to checking account ‘23’
		CheckingPrenoteCredit,
		// Zero dollar with remittance data
		CheckingZeroDollarRemittanceCredit,

		// Demand Debit Records (for checking, NOW, and share draft accounts)

		// Automated Return or Notification of Change for original transaction code 27, 28, or 29
		CheckingReturnNOCDebit,
		// Debit (withdrawal) to checking account ‘27’
		CheckingDebit,
		// Prenote for debit to checking account ‘28’
		CheckingPrenoteDebit,
		// Zero dollar with remittance data (for CCD, CTX, and IAT Entries only)
		CheckingZeroDollarRemittanceDebit,

		// Savings Account Credit Records

		// Return or Notification of Change for original transaction code 32, 33, or 34
		SavingsReturnNOCCredit,
		// Credit to savings account ‘32’
		SavingsCredit,
		// Prenote for credit to savings account ‘33’
		SavingsPrenoteCredit,
		// Zero dollar with remittance data (for CCD, CTX, and IAT Entries only); Acknowledgment Entries (ACK and ATX Entries only)
		SavingsZeroDollarRemittanceCredit,

		// Savings Account Debit Records

		// Automated Return or Notification of Change for original transaction code '37', '38', '39
		SavingsReturnNOCDebit,
		// Debit to savings account ‘37’
		SavingsDebit,
		// Prenote for debit to savings account ‘38’
		SavingsPrenoteDebit,
		// Zero dollar with remittance data
		SavingsZeroDollarRemittanceDebit,

		// Financial Institution General Ledger Credit Records

		//Return or Notification of Change for original transaction code 42, 43, or 44
		GLReturnNOCCredit,
		// General Ledger Credit
		GLCredit,
		// Prenotification of General Ledger Credit (non-dollar)
		GLPrenoteCredit,
		// Zero dollar with remittance data
		GLZeroDollarRemittanceCredit,

		// Financial Institution General Ledger Debit Records

		// Return or Notification of Change for original transaction code 47, 48, or 49
		GLReturnNOCDebit,
		//General Ledger Debit
		GLDebit,
		// Prenotification of General Ledger Debit (non-dollar)
		GLPrenoteDebit,
		// Zero dollar with remittance data
		GLZeroDollarRemittanceDebit,

		// Loan Account Credit Records
		// Return or Notification of Change for original transaction code 52, 53, or 54
		LoanReturnNOCCredit,
		// Loan Account Credit
		LoanCredit,
		// Prenotification of Loan Account Credit (non-dollar)
		LoanPrenoteCredit,
		// Zero dollar with remittance data
		LoanZeroDollarRemittanceCredit,

		// Loan Account Debit Records (for Reversals Only)

		// Loan Account Debit (Reversals Only)
		LoanDebit,
		// Return or Notification of Change for original transaction code 55
		LoanReturnNOCDebit,

		// Accounting Records (for use in ADV Files only)
		// These transaction codes represent accounting Entries.

		// Credit for ACH debits originated
		CreditForDebitsOriginated,
		//Debit for ACH credits originated
		DebitForCreditsOriginated,
		// Credit for ACH credits received
		CreditForCreditsReceived,
		// Debit for ACH debits received
		DebitForDebitsReceived,
		// Credit for ACH credits in Rejected batches
		CreditForCreditsRejected,
		// Debit for ACH debits in Rejected batches
		DebitForDebitsRejectedBatches,
		// Summary credit for respondent ACH activity
		CreditSummary,
		// Summary debit for respondent ACH activity
		DebitSummary:
		return nil
	}
	return errors.New(msgTransactionCode)
}

// isTransactionTypeCode verifies Addenda10 TransactionTypeCode is a valid value
// ANN = Annuity, BUS = Business/Commercial, DEP = Deposit, LOA = Loan, MIS = Miscellaneous, MOR = Mortgage
// PEN = Pension, RLS = Rent/Lease, REM = Remittance2, SAL = Salary/Payroll, TAX = Tax, TEL = Telephone-Initiated Transaction
// WEB = Internet-Initiated Transaction, ARC = Accounts Receivable Entry, BOC = Back Office Conversion Entry,
// POP = Point of Purchase Entry, RCK = Re-presented Check Entry
func (v *validator) isTransactionTypeCode(s string) error {
	switch s {
	case "ANN", "BUS", "DEP", "LOA", "MIS", "MOR",
		"PEN", "RLS", "REM", "SAL", "TAX", TEL, WEB,
		ARC, BOC, POP, RCK:
		return nil
	}
	return errors.New(msgTransactionTypeCode)
}

// isUpperAlphanumeric checks if string only contains ASCII alphanumeric upper case characters
func (v *validator) isUpperAlphanumeric(s string) error {
	if upperAlphanumericRegex.MatchString(s) {
		return errors.New(msgUpperAlpha)
	}
	return nil
}

// isAlphanumeric checks if a string only contains ASCII alphanumeric characters
func (v *validator) isAlphanumeric(s string) error {
	if alphanumericRegex.MatchString(s) {
		// ^[ A-Za-z0-9_@./#&+-]*$/
		return errors.New(msgAlphanumeric)
	}
	return nil
}

// CalculateCheckDigit returns a check digit for a routing number
// Multiply each digit in the Routing number by a weighting factor. The weighting factors for each digit are:
// Position: 1 2 3 4 5 6 7 8
// Weights : 3 7 1 3 7 1 3 7
// Add the results of the eight multiplications
// Subtract the sum from the next highest multiple of 10.
// The result is the Check Digit
func (v *validator) CalculateCheckDigit(routingNumber string) int {
	if n := utf8.RuneCountInString(routingNumber); n != 8 && n != 9 {
		return -1
	}

	var routeIndex [8]string
	for i := 0; i < 8; i++ {
		if routingNumber[i] < '0' || routingNumber[i] > '9' {
			return -1 // only digits are allowed
		}
		routeIndex[i] = string(routingNumber[i])
	}
	n, _ := strconv.Atoi(routeIndex[0])
	sum := n * 3
	n, _ = strconv.Atoi(routeIndex[1])
	sum = sum + (n * 7)
	n, _ = strconv.Atoi(routeIndex[2])
	sum = sum + n // multiply by 1
	n, _ = strconv.Atoi(routeIndex[3])
	sum = sum + (n * 3)
	n, _ = strconv.Atoi(routeIndex[4])
	sum = sum + (n * 7)
	n, _ = strconv.Atoi(routeIndex[5])
	sum = sum + n // multiply by 1
	n, _ = strconv.Atoi(routeIndex[6])
	sum = sum + (n * 3)
	n, _ = strconv.Atoi(routeIndex[7])
	sum = sum + (n * 7)

	return v.roundUp10(sum) - sum
}

// CheckRoutingNumber returns a nil error if the provided routingNumber is valid according to
// NACHA rules. See CalculateCheckDigit for details on computing the check digit.
func CheckRoutingNumber(routingNumber string) error {
	if routingNumber == "" {
		return errors.New("no routing number provided")
	}
	if n := utf8.RuneCountInString(routingNumber); n != 9 {
		return fmt.Errorf("invalid routing number length of %d", n)
	}

	v := new(validator)
	check := fmt.Sprintf("%d", v.CalculateCheckDigit(routingNumber))
	last := string(routingNumber[len(routingNumber)-1])
	if check != last {
		return fmt.Errorf("routing number checksum mismatch: expected %s but got %s", check, last)
	}
	return nil
}

// roundUp10 round number up to the next ten spot.
func (v *validator) roundUp10(n int) int {
	return int(math.Ceil(float64(n)/10.0)) * 10
}
