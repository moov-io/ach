// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ach

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	lowerAlphaCharacters  = "abcdefghijklmnopqrstuvwxyz"
	numericCharacters     = "0123456789"
	asciiCharacters       = ` !"#$%&'()*+,-./:;<=>?@[\]^_{|}~` + "`"
	ebcdicExtraCharacters = `¢¬¦±`

	validAlphaNumericCharacters          = lowerAlphaCharacters + strings.ToUpper(lowerAlphaCharacters) + numericCharacters + asciiCharacters + ebcdicExtraCharacters
	validUppercaseAlphaNumericCharacters = strings.ToUpper(lowerAlphaCharacters) + numericCharacters + asciiCharacters + ebcdicExtraCharacters
)

// validator is common validation and formatting of golang types to ach type strings
type validator struct{}

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
	return ErrCardTransactionType
}

// isCreditCardYear validates a 2 digit year for credit cards, but
// only accepts a range of years. 2018 to 2050
func (v *validator) isCreditCardYear(s string) error {
	if s < "18" || s > "50" {
		return ErrValidYear
	}
	return nil
}

// isMonth validates a 2 digit month 01-12
func (v *validator) isMonth(s string) error {
	switch s {
	case
		"01", "02", "03", "04", "05", "06",
		"07", "08", "09", "10", "11", "12":
		return nil
	}
	return ErrValidMonth
}

// isDay validates a 2 digit day based on a 2 digit month
// months are 01-12, days are 01-29, 01-30, or 01-31
func (v *validator) isDay(m string, d string) error {
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
	return ErrValidDay
}

// validateSimpleDate will return the incoming string only if it matches a valid YYMMDD
// date format. (Y=Year, M=Month, D=Day)
func (v *validator) validateSimpleDate(s string) string {
	_, err := time.Parse("060102", s) // YYMMDD
	if err != nil {
		return ""
	}
	return s
}

var (
	// hhmmRegex defines a regex for all valid 24-hour clock timestamps.
	// Format: HHmm (H=hour, m=minute) - (first H can only be 0, 1, or 2)
	hhmmRegex = regexp.MustCompile(`^([0-2]{1}[\d]{1}[0-5]{1}\d{1})$`)
)

// validateSimpleTime will return the incoming string only if it is a valid 24-hour clock time.
func (v *validator) validateSimpleTime(s string) string {
	if hhmmRegex.MatchString(s) {
		return s // successfully matched and validated
	}
	return ""
}

// isForeignExchangeIndicator ensures foreign exchange indicators of an
// IATBatchHeader is valid
func (v *validator) isForeignExchangeIndicator(code string) error {
	switch code {
	case
		"FV", "VF", "FF":
		return nil
	}
	return ErrForeignExchangeIndicator
}

// isForeignExchangeReferenceIndicator ensures foreign exchange reference
// indicator of am IATBatchHeader is valid
func (v *validator) isForeignExchangeReferenceIndicator(code int) error {
	switch code {
	case
		1, 2, 3:
		return nil
	}
	return ErrForeignExchangeReferenceIndicator
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
	return ErrIDNumberQualifier
}

// isOriginatorStatusCode ensures status code of a batch is valid
func (v *validator) isOriginatorStatusCode(code int) error {
	switch code {
	case
		// ADV file - prepared by an ACH Operator
		0,
		// Originator is a financial institution
		1,
		// Originator is a Government Agency or other agency not subject to ACH Rules
		2:
		return nil
	}
	return ErrOrigStatusCode
}

// isSECCode returns true if a SEC Code of a Batch is found
func (v *validator) isSECCode(code string) error {
	switch code {
	case
		ACK, ADV, ARC, ATX, BOC, CCD, CIE, COR, CTX, DNE, ENR,
		IAT, MTE, POS, PPD, POP, RCK, SHR, TEL, TRC, TRX, WEB, XCK:
		return nil
	}
	return ErrSECCode
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
	return ErrServiceClass
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
	return ErrAddendaTypeCode
}

// isTransactionCode ensures TransactionCode of an Entry is valid
//
// The Tran Code is a two-digit code in positions 2 - 3 of the Entry Detail Record (6 Record) within an ACH File.
// The first digit of the Tran Code indicates the account type to which the entry will post, where the number:
//
//	"2" designates a Checking Account.
//	"3" designates a Savings Account.
//	"4" designates a General Ledger Account.
//	"5" designates Loan Account.
//
// The second digit of the Tran Code identifies the entry as:
//
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
	return StandardTransactionCode(code)
}

// StandardTransactionCode checks the provided TransactionCode to verify it is a valid NACHA value.
func StandardTransactionCode(code int) error {
	switch code {
	// TransactionCode if the receivers account is:
	case
		// Demand Credit Records (for checking, NOW, and share draft accounts)

		// Automated Return or Notification of Change for original transaction code '22', '23, '24'
		CheckingReturnNOCCredit,
		// Credit (deposit) to checking account '22'
		CheckingCredit,
		// Prenote for credit to checking account '23'
		CheckingPrenoteCredit,
		// Zero dollar with remittance data
		CheckingZeroDollarRemittanceCredit,

		// Demand Debit Records (for checking, NOW, and share draft accounts)

		// Automated Return or Notification of Change for original transaction code 27, 28, or 29
		CheckingReturnNOCDebit,
		// Debit (withdrawal) to checking account '27'
		CheckingDebit,
		// Prenote for debit to checking account '28'
		CheckingPrenoteDebit,
		// Zero dollar with remittance data (for CCD, CTX, and IAT Entries only)
		CheckingZeroDollarRemittanceDebit,

		// Savings Account Credit Records

		// Return or Notification of Change for original transaction code 32, 33, or 34
		SavingsReturnNOCCredit,
		// Credit to savings account '32'
		SavingsCredit,
		// Prenote for credit to savings account '33'
		SavingsPrenoteCredit,
		// Zero dollar with remittance data (for CCD, CTX, and IAT Entries only); Acknowledgment Entries (ACK and ATX Entries only)
		SavingsZeroDollarRemittanceCredit,

		// Savings Account Debit Records

		// Automated Return or Notification of Change for original transaction code '37', '38', '39
		SavingsReturnNOCDebit,
		// Debit to savings account '37'
		SavingsDebit,
		// Prenote for debit to savings account '38'
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
	return ErrTransactionCode
}

func (v *validator) isPrenote(code int) bool {
	switch code {
	case CheckingPrenoteCredit, CheckingPrenoteDebit,
		SavingsPrenoteCredit, SavingsPrenoteDebit,
		GLPrenoteCredit, GLPrenoteDebit, LoanPrenoteCredit:
		return true
	}
	return false
}

// isTransactionTypeCode verifies Addenda10 TransactionTypeCode is a valid value
// This code is used as a Secondary SEC code to help identify the source and purpose of the transaction.
//
//	ANN = Annuity, BUS = Business/Commercial, DEP = Deposit, LOA = Loan, MIS = Miscellaneous, MOR = Mortgage
//	PEN = Pension, REM = Remittance2, RLS = Rent/Lease, SAL = Salary/Payroll, TAX = Tax
//
//	ARC = Accounts Receivable Entry, BOC = Back Office Conversion Entry, IAT = International ACH Transaction,
//	MTE = Machine Transfer Entry, POP = Point of Purchase Entry, POS Point of Sale, RCK = Re-presented Check Entry,
//	SHR = Shared Network Transaction, TEL = Telephone-Initiated Transaction, WEB = Internet-Initiated Transaction
//
// Also, according to the Nacha rules, "There is no requirement to add Secondary SEC Codes for PPD, CCD, CTX, and
// other SEC codes not included in the list above."
func (v *validator) isTransactionTypeCode(s string) error {
	switch strings.ToUpper(s) {
	case
		"ANN", "BUS", "DEP", "LOA", "MIS", "MOR",
		"PEN", "REM", "RLS", "SAL", "TAX",
		ARC, BOC, IAT, MTE, POP, POS, RCK, SHR, TEL, WEB:
		return nil
	}
	return ErrTransactionTypeCode
}

func (v *validator) includesValidCharacters(input string, charset string) error {
	for _, i := range input {
		var found bool
		for _, c := range charset {
			if i == c {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("invalid character: %v", i)
		}
	}
	return nil
}

// isUpperAlphanumeric checks if string only contains ASCII alphanumeric upper case characters
func (v *validator) isUpperAlphanumeric(s string) error {
	err := v.includesValidCharacters(s, validUppercaseAlphaNumericCharacters)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUpperAlpha, err)
	}
	return nil
}

// isAlphanumeric checks if a string only contains ASCII alphanumeric characters
func (v *validator) isAlphanumeric(s string) error {
	err := v.includesValidCharacters(s, validAlphaNumericCharacters)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrNonAlphanumeric, err)
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
func CalculateCheckDigit(routingNumber string) int {
	if n := utf8.RuneCountInString(routingNumber); n != 8 && n != 9 {
		return -1
	}

	var sum int
	for i, r := range routingNumber {
		// Don't process check digit of routing number
		if i >= 8 {
			break
		}

		// Reject anything that's not a digit
		if r < '0' || r > '9' {
			return -1 // only digits are allowed
		}

		// Calculate the check digit
		var n int32 = (r - '0')

		switch i {
		case 0, 3, 6:
			sum += int(n * 3)
		case 1, 4, 7:
			sum += int(n * 7)
		case 2, 5:
			sum += int(n)
		}
	}

	return roundUp10(sum) - sum
}

func (v *validator) CalculateCheckDigit(routingNumber string) int {
	return CalculateCheckDigit(routingNumber)
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

	check := fmt.Sprintf("%d", CalculateCheckDigit(routingNumber))
	last := string(routingNumber[len(routingNumber)-1])
	if check != last {
		return fmt.Errorf("routing number checksum mismatch: expected %s but got %s", check, last)
	}
	return nil
}

// roundUp10 round number up to the next ten spot.
func roundUp10(n int) int {
	return int(math.Ceil(float64(n)/10.0)) * 10
}

func (v *validator) validateSettlementDate(s string) string {
	emptyField := "   "

	if s == emptyField || utf8.RuneCountInString(s) != len(emptyField) {
		return emptyField
	}

	day, err := strconv.Atoi(s)
	if err != nil {
		return emptyField
	}

	if day < 1 || day > 366 {
		return emptyField
	}

	return s

}
