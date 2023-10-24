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
	"strings"
	"unicode/utf8"
)

// When a Return Entry is prepared, the original Company/Batch Header Record, the original Entry Detail Record,
// and the Company/Batch Control Record are copied for return to the Originator.
//
// The Return Entry is a new Entry. These Entries must be assigned new batch and trace numbers, new identification numbers for the returning institution,
// appropriate transaction codes, etc., as required per format specifications.
//
// See Appendix Four: Return Entries in the NACHA Corporate

var (
	returnCodeDict = map[string]*ReturnCode{}
)

func init() {
	// populate the ReturnCode map with lookup values
	returnCodeDict = makeReturnCodeDict()
}

// Addenda99 utilized for Notification of Change Entry (COR) and Return types.
type Addenda99 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// TypeCode Addenda types code '99'
	TypeCode string `json:"typeCode"`
	// ReturnCode field contains a standard code used by an ACH Operator or RDFI to describe the reason for returning an Entry.
	// Must exist in returnCodeDict
	ReturnCode string `json:"returnCode"`
	// OriginalTrace This field contains the Trace Number as originally included on the forward Entry or Prenotification.
	// The RDFI must include the Original Entry Trace Number in the Addenda Record of an Entry being returned to an ODFI,
	// in the Addenda Record of an 98, within an Acknowledgment Entry, or with an RDFI request for a copy of an authorization.
	OriginalTrace string `json:"originalTrace"`
	// DateOfDeath The field date of death is to be supplied on Entries being returned for reason of death (return reason codes R14 and R15). Format: YYMMDD (Y=Year, M=Month, D=Day)
	DateOfDeath string `json:"dateOfDeath"`
	// OriginalDFI field contains the Receiving DFI Identification (addenda.RDFIIdentification) as originally included on the forward Entry or Prenotification that the RDFI is returning or correcting.
	OriginalDFI string `json:"originalDFI"`
	// AddendaInformation
	AddendaInformation string `json:"addendaInformation,omitempty"`
	// TraceNumber matches the Entry Detail Trace Number of the entry being returned.
	//
	// Use TraceNumberField for a properly formatted string representation.
	TraceNumber string `json:"traceNumber,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters

	validateOpts *ValidateOpts
}

// ReturnCode holds a return Code, Reason/Title, and Description
//
// Table of return codes exists in Part 4.2 of the NACHA corporate rules and guidelines
type ReturnCode struct {
	Code        string `json:"code"`
	Reason      string `json:"reason"`
	Description string `json:"description"`
}

// NewAddenda99 returns a new Addenda99 with default values for none exported fields
func NewAddenda99() *Addenda99 {
	Addenda99 := &Addenda99{
		TypeCode: "99",
	}
	return Addenda99
}

// Parse takes the input record string and parses the Addenda99 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (Addenda99 *Addenda99) Parse(record string) {
	runeCount := utf8.RuneCountInString(record)
	if runeCount != 94 {
		return
	}

	buf := getBuffer()
	defer saveBuffer(buf)

	reset := func() string {
		out := buf.String()
		buf.Reset()
		return out
	}

	// We're going to process the record rune-by-rune and at each field cutoff save the value.
	var idx int
	for _, r := range record {
		idx++

		// Append rune to buffer
		buf.WriteRune(r)

		// At each cutoff save the buffer and reset
		switch idx {
		case 0, 1:
			// 1-1 Always 7
			reset()
		case 3:
			// 2-3 Defines the specific explanation and format for the addenda information contained in the same record
			Addenda99.TypeCode = reset()
		case 6:
			// 4-6
			Addenda99.ReturnCode = reset()
		case 21:
			// 7-21
			Addenda99.OriginalTrace = strings.TrimSpace(reset())
		case 27:
			// 22-27, might be a date or blank
			Addenda99.DateOfDeath = Addenda99.validateSimpleDate(reset())
		case 35:
			// 28-35
			Addenda99.OriginalDFI = Addenda99.parseStringField(reset())
		case 79:
			// 36-79
			Addenda99.AddendaInformation = strings.TrimSpace(reset())
		case 94:
			// 80-94
			Addenda99.TraceNumber = strings.TrimSpace(reset())
		}
	}
}

// String writes the Addenda99 struct to a 94 character string
func (Addenda99 *Addenda99) String() string {
	if Addenda99 == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryAddendaPos)
	buf.WriteString(Addenda99.TypeCode)
	buf.WriteString(Addenda99.ReturnCode)
	buf.WriteString(Addenda99.OriginalTraceField())
	buf.WriteString(Addenda99.DateOfDeathField())
	buf.WriteString(Addenda99.OriginalDFIField())
	buf.WriteString(Addenda99.AddendaInformationField())
	buf.WriteString(Addenda99.TraceNumberField())
	return buf.String()
}

// Validate verifies NACHA rules for Addenda99
func (Addenda99 *Addenda99) Validate() error {
	if Addenda99.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, Addenda99.TypeCode)
	}
	if Addenda99.TypeCode != "99" {
		return fieldError("TypeCode", ErrAddendaTypeCode, Addenda99.TypeCode)
	}

	if Addenda99.validateOpts == nil || !Addenda99.validateOpts.CustomReturnCodes {
		_, ok := returnCodeDict[Addenda99.ReturnCode]
		if !ok {
			// Return Addenda requires a valid ReturnCode
			return fieldError("ReturnCode", ErrAddenda99ReturnCode, Addenda99.ReturnCode)
		}
	}

	return nil
}

// SetValidation stores ValidateOpts on the Batch which are to be used to override
// the default NACHA validation rules.
func (Addenda99 *Addenda99) SetValidation(opts *ValidateOpts) {
	if Addenda99 == nil {
		return
	}
	Addenda99.validateOpts = opts
}

// OriginalTraceField returns a zero padded OriginalTrace string
func (Addenda99 *Addenda99) OriginalTraceField() string {
	return Addenda99.stringField(Addenda99.OriginalTrace, 15)
}

// DateOfDeathField returns a space padded DateOfDeath string
func (Addenda99 *Addenda99) DateOfDeathField() string {
	// Return space padded 6 characters if it is a zero value of DateOfDeath
	if Addenda99.DateOfDeath == "" {
		return Addenda99.alphaField("", 6)
	}
	return Addenda99.formatSimpleDate(Addenda99.DateOfDeath)
}

// OriginalDFIField returns a zero padded OriginalDFI string
func (Addenda99 *Addenda99) OriginalDFIField() string {
	return Addenda99.stringField(Addenda99.OriginalDFI, 8)
}

// AddendaInformationField returns a space padded AddendaInformation string
func (Addenda99 *Addenda99) AddendaInformationField() string {
	return Addenda99.alphaField(Addenda99.AddendaInformation, 44)
}

// IATPaymentAmount sets original forward entry payment amount characters 1-10 of underlying AddendaInformation
func (Addenda99 *Addenda99) IATPaymentAmount(s string) {
	Addenda99.AddendaInformation = Addenda99.stringField(s, 10)
}

// IATAddendaInformation sets Addenda Information for IAT return items, characters 10-44 of
// underlying AddendaInformation
func (Addenda99 *Addenda99) IATAddendaInformation(s string) {
	Addenda99.AddendaInformation = Addenda99.AddendaInformation + Addenda99.alphaField(s, 34)
}

// IATPaymentAmountField returns original forward entry payment amount int, characters 1-10 of
// underlying AddendaInformation
func (Addenda99 *Addenda99) IATPaymentAmountField() int {
	return Addenda99.parseNumField(Addenda99.AddendaInformation[0:10])
}

// IATAddendaInformationField returns a space padded AddendaInformation string, characters 10-44 of
// underlying AddendaInformation
func (Addenda99 *Addenda99) IATAddendaInformationField() string {
	return Addenda99.alphaField(Addenda99.AddendaInformation[9:44], 34)
}

// TraceNumberField returns a zero padded TraceNumber string
func (Addenda99 *Addenda99) TraceNumberField() string {
	return Addenda99.stringField(Addenda99.TraceNumber, 15)
}

// ReturnCodeField gives the ReturnCode struct for the given Addenda99 record
func (Addenda99 *Addenda99) ReturnCodeField() *ReturnCode {
	code, ok := returnCodeDict[Addenda99.ReturnCode]
	if ok {
		return code
	}
	return nil
}

// LookupReturnCode will return a struct representing the reason and description for
// the provided NACHA return code.
func LookupReturnCode(code string) *ReturnCode {
	if code, exists := returnCodeDict[strings.ToUpper(code)]; exists {
		return code
	}
	return nil
}

func makeReturnCodeDict() map[string]*ReturnCode {
	dict := make(map[string]*ReturnCode)

	codes := []ReturnCode{
		// Return Reason Codes for RDFIs
		{"R01", "Insufficient Funds", "Available balance is not sufficient to cover the dollar value of the debit entry"},
		{"R02", "Account Closed", "Previously active account has been closed by customer or RDFI"},
		// R03 may not be used to return ARC, BOC or POP entries solely because they do not contain an Individual Name.
		{"R03", "No Account/Unable to Locate Account", "Account number structure is valid and passes editing process, but does not correspond to individual or is not an open account"},
		{"R04", "Invalid Account Number", "Account number structure not valid; entry may fail check digit validation or may contain an incorrect number of digits."},
		{"R05", "Improper Debit to Consumer Account", "A CCD, CTX, or CBR debit entry was transmitted to a Consumer Account of the Receiver and was not authorized by the Receiver"},
		{"R06", "Returned per ODFI's Request", "ODFI has requested RDFI to return the ACH entry (optional to RDFI - ODFI indemnifies RDFI)"},
		// R07 Prohibited use for ARC, BOC, POP and RCK.
		{"R07", "Authorization Revoked by Customer", "Consumer, who previously authorized ACH payment, has revoked authorization from Originator (must be returned no later than 60 days from settlement date and customer must sign affidavit)"},
		{"R08", "Payment Stopped", "Receiver of a recurring debit transaction has stopped payment to a specific ACH debit. RDFI should verify the Receiver's intent when a request for stop payment is made to insure this is not intended to be a revocation of authorization"},
		{"R09", "Uncollected Funds", "Sufficient book or ledger balance exists to satisfy dollar value of the transaction, but the dollar value of transaction is in process of collection (i.e., uncollected checks) or cash reserve balance below dollar value of the debit entry."},
		{"R10", "Customer Advises Originator is Not Known to Receiver and/or Originator is Not Authorized by Receiver to Debit Receiver’s Account", "The receiver does not know the Originator’s identity and/or has not authorized the Originator to debit. Alternatively, for ARC and BOC entries, the signature on the source document is not authentic or authorized. For POP entries, the signature on the written authorization is not authentic or authorized."},
		{"R11", "Customer Advises Entry Not in Accordance with the Terms of the Authorization", "The Originator and Receiver have a relationship, and an authorization to debit exists, but there is an error or defect in the payment such that the entry does not conform to the terms of the authorization. The Originator may correct the error and submit a new entry within 60 days of the return entry's settlement date without the need for re-authorization by the Receiver."},
		{"R12", "Branch Sold to Another DFI", "Financial institution receives entry destined for an account at a branch that has been sold to another financial institution."},
		{"R13", "RDFI not qualified to participate", "Financial institution does not receive commercial ACH entries"},
		{"R14", "Representative payee deceased or unable to continue in that capacity", "The representative payee authorized to accept entries on behalf of a beneficiary is either deceased or unable to continue in that capacity"},
		{"R15", "Beneficiary or bank account holder", "(Other than representative payee) deceased* - (1) the beneficiary entitled to payments is deceased or (2) the bank account holder other than a representative payee is deceased"},
		{"R16", "Bank account frozen", "Funds in bank account are unavailable due to action by RDFI or legal order"},
		{"R17", "File Record Edit Criteria/Entry with Invalid Account Number Initiated Under Questionable Circumstances", "(1) Field(s) cannot be processed by RDFI; or (2) the Entry contains an invalid DFI Account Number (account closed/no account/unable to locate account/invalid account number) and is believed by the RDFI to have been initiated under questionable circumstances; or (3) either the RDFI or Receiver has identified a Reversing Entry as one that was improperly initiated by the Originator or ODFI."},
		{"R18", "Improper effective entry date", "Entries have been presented prior to the first available processing window for the effective date."},
		{"R19", "Amount field error", "Improper formatting of the amount field"},
		{"R20", "Non-payment bank account", "Entry destined for non-payment bank account defined by reg."},
		{"R21", "Invalid company ID number", "The company ID information not valid (normally CIE entries)"},
		{"R22", "Invalid individual ID number", "Individual id used by receiver is incorrect (CIE entries)"},
		{"R23", "Credit entry refused by receiver", "Receiver returned entry because minimum or exact amount not remitted, bank account is subject to litigation, or payment represents an overpayment, originator is not known to receiver or receiver has not authorized this credit entry to this bank account"},
		{"R24", "Duplicate entry", "RDFI has received a duplicate entry"},
		{"R25", "Addenda error", "Improper formatting of the addenda record information"},
		{"R26", "Mandatory field error", "Improper information in one of the mandatory fields"},
		{"R27", "Trace number error", "Original entry trace number is not valid for return entry; or addenda trace numbers do not correspond with entry detail record"},
		{"R28", "Transit routing number check digit error", "Check digit for the transit routing number is incorrect"},
		{"R29", "Corporate customer advises not authorized", "RDFI has been notified by corporate receiver that debit entry of originator is not authorized"},
		{"R30", "RDFI not participant in check truncation program", "Financial institution not participating in automated check safekeeping application"},
		{"R31", "Permissible return entry (CCD and CTX only)", "RDFI has been notified by the ODFI that it agrees to accept a CCD or CTX return entry"},
		{"R32", "RDFI non-settlement", "RDFI is not able to settle the entry"},
		{"R33", "Return of XCK entry", "RDFI determines at its sole discretion to return an XCK entry; an XCK return entry may be initiated by midnight of the sixtieth day following the settlement date if the XCK entry"},
		{"R34", "Limited participation RDFI", "RDFI participation has been limited by a federal or state supervisor"},
		{"R35", "Return of improper debit entry", "ACH debit not permitted for use with the CIE standard entry class code (except for reversals)"},
		{"R36", "Return of improper credit entry", "ACH credit entries (with the exception of reversing entries) are not permitted for use with ARC, BOC, POP, RCK, TEL, and XCK."},
		{"R37", "Source Document Presented for Payment (Adjustment Entry)", "The source document to which an ARC, BOC or POP entry relates has been presented for payment. RDFI must obtain a Written Statement and return the entry within 60 days following Settlement Date"},
		{"R38", "Stop Payment on Source Document (Adjustment Entry)", "A stop payment has been placed on the source document to which the ARC or BOC entry relates. RDFI must return no later than 60 days following Settlement Date. No Written Statement is required as the original stop payment form covers the return"},
		{"R39", "Improper Source Document", "The RDFI has determined the source document used for the ARC, BOC or POP entry to its Receiver's account is improper."},
		// Return Codes to be used for ENR entries and are initiated by a Federal Government Agency
		{"R40", "Return of ENR Entry by Federal Government Agency (ENR Only)", "This return reason code may only be used to return ENR entries and is at the federal Government Agency's Sole discretion"},
		{"R41", "Invalid Transaction Code (ENR only)", "Either the Transaction Code included in Field 3 of the Addenda Record does not conform to the ACH Record Format Specifications contained in Appendix Three (ACH Record Format Specifications) or it is not appropriate with regard to an Automated Enrollment Entry."},
		{"R42", "Routing Number/Check Digit Error (ENR Only)", "The Routing Number and the Check Digit included in Field 3 of the Addenda Record is either not a valid number or it does not conform to the Modulus 10 formula."},
		{"R43", "Invalid DFI Account Number (ENR Only)", "The Receiver's account number included in Field 3 of the Addenda Record must include at least one alphameric character."},
		{"R44", "Invalid Individual ID Number/Identification Number (ENR only)", "The Individual ID Number/Identification Number provided in Field 3 of the Addenda Record does not match a corresponding ID number in the Federal Government Agency's records."},
		{"R45", "Invalid Individual Name/Company Name (ENR only)", "The name of the consumer or company provided in Field 3 of the Addenda Record either does not match a corresponding name in the Federal Government Agency's records or fails to include at least one alphameric character."},
		{"R46", "Invalid Representative Payee Indicator (ENR Only)", "The Representative Payee Indicator Code included in Field 3 of the Addenda Record has been omitted or it is not consistent with the Federal Government Agency's records."},
		{"R47", "Duplicate Enrollment (ENR Only)", "The Entry is a duplicate of an Automated Enrollment Entry previously initiated by a DFI."},
		// Return Codes to be used for RCK entries only and are initiated by a RDFI
		{"R50", "State Law Affecting RCK Acceptance", "RDFI is located in a state that has not adopted Revised Article 4 of the UCC or the RDFI is located in a state that requires all canceled checks to be returned within the periodic statement"},
		{"R51", "Item Related to RCK Entry is Ineligible or RCK Entry is Improper", "The item to which the RCK entry relates was not eligible, Originator did not provide notice of the RCK policy, signature on the item was not genuine, the item has been altered or amount of the entry was not accurately obtained from the item. RDFI must obtain a Written Statement and return the entry within 60 days following Settlement Date"},
		{"R52", "Stop Payment on Item (Adjustment Entry)", "A stop payment has been placed on the item to which the RCK entry relates. RDFI must return no later than 60 days following Settlement Date. No Written Statement is required as the original stop payment form covers the return."},
		{"R53", "Item and RCK Entry Presented for Payment (Adjustment Entry)", "Both the RCK entry and check have been presented for payment. RDFI must obtain a Written Statement and return the entry within 60 days following Settlement Date"},
		// Return Codes to be used by the ODFI for dishonored return entries
		{"R61", "Misrouted Return", "The financial institution preparing the Return Entry (the RDFI of the original Entry) has placed the incorrect Routing Number in the Receiving DFI Identification field."},
		{"R62", "Return of Erroneous or Reversing Debt", "The Originator’s/ODFI’s use of the reversal process resulted in, or failed to correct, an unintended credit to the Receiver."},
		{"R67", "Duplicate Return", "The ODFI has received more than one Return for the same Entry."},
		{"R68", "Untimely Return", "The Return Entry has not been sent within the time frame established by these Rules."},
		{"R69", "Field Error(s)", "One or more of the field requirements are incorrect."},
		{"R70", "Permissible Return Entry Not Accepted/Return Not Requested by ODFI", "The ODFI has received a Return Entry identified by the RDFI as being returned with the permission of, or at the request of, the ODFI, but the ODFI has not agreed to accept the Entry or has not requested the return of the Entry."},
		// Return Codes to be used by the RDFI for contested dishonored return entries
		{"R71", "Misrouted Dishonored Return", "The financial institution preparing the dishonored Return Entry (the ODFI of the original Entry) has placed the incorrect Routing Number in the Receiving DFI Identification field."},
		{"R72", "Untimely Dishonored Return", "The dishonored Return Entry has not been sent within the designated time frame."},
		{"R73", "Timely Original Return", "The RDFI is certifying that the original Return Entry was sent within the time frame designated in these Rules."},
		{"R74", "Corrected Return", "The RDFI is correcting a previous Return Entry that was dishonored using Return Reason Code R69 (Field Error(s)) because it contained incomplete or incorrect information."},
		{"R75", "Return Not a Duplicate", "The Return Entry was not a duplicate of an Entry previously returned by the RDFI."},
		{"R76", "No Errors Found", "The original Return Entry did not contain the errors indicated by the ODFI in the dishonored Return Entry."},
		{"R77", "Non-Acceptance of R62 Dishonored Return", "The RDFI returned the Erroneous Entry and the related Reversing Entry. Alternatively, the funds relating to the R62 dishonored Return are not recoverably from the Receiver."},
		//Return Codes to be used by Gateways for the return of international payments
		{"R80", "IAT Entry Coding Error", "The IAT Entry is being returned due to one or more of the following conditions: Invalid DFI/Bank Branch Country Code, invalid DFI/Bank Identification Number Qualifier, invalid Foreign Exchange Indicator, invalid ISO Originating Currency Code, invalid ISO Destination Currency Code, invalid ISO Destination Country Code, invalid Transaction Type Code"},
		{"R81", "Non-Participant in IAT Program", "The IAT Entry is being returned because the Gateway does not have an agreement with either the ODFI or the Gateway's customer to transmit Outbound IAT Entries."},
		{"R82", "Invalid Foreign Receiving DFI Identification", "The reference used to identify the Foreign Receiving DFI of an Outbound IAT Entry is invalid."},
		{"R83", "Foreign Receiving DFI Unable to Settle", "The IAT Entry is being returned due to settlement problems in the foreign payment system."},
		{"R84", "Entry Not Processed by Gateway", "For Outbound IAT Entries, the Entry has not been processed and is being returned at the Gateway's discretion because either (1) the processing of such Entry may expose the Gateway to excessive risk, or (2) the foreign payment system does not support the functions needed to process the transaction."},
		{"R85", "Incorrectly Coded Outbound International Payment", "The RDFI/Gateway has identified the Entry as an Outbound international payment and is returning the Entry because it bears an SEC Code that lacks information required by the Gateway for OFAC compliance."},
	}
	// populate the map
	for i := range codes {
		dict[codes[i].Code] = &codes[i]
	}
	return dict
}
