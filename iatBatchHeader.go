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
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/moov-io/ach/internal/iso3166"
	"github.com/moov-io/ach/internal/iso4217"
)

// msgServiceClass

// IATBatchHeader identifies the originating entity and the type of transactions
// contained in the batch for SEC Code IAT. This record also contains the effective
// date, or desired settlement date, for all entries contained in this batch. The
// settlement date field is not entered as it is determined by the ACH operator.
//
// An IAT entry is a credit or debit ACH entry that is part of a payment transaction
// involving a financial agency's office (i.e., depository financial institution or
// business issuing money orders) that is not located in the territorial jurisdiction
// of the United States. IAT entries can be made to or from a corporate or consumer
// account and must be accompanied by seven (7) mandatory addenda records identifying
// the name and physical address of the Originator, name and physical address of the
// Receiver, Receiver's account number, Receiver's bank identity and reason for the payment.
type IATBatchHeader struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`

	// ServiceClassCode ACH Mixed Debits and Credits '200'
	// ACH Credits Only '220'
	// ACH Debits Only '225'
	ServiceClassCode int `json:"serviceClassCode"`

	// IATIndicator - Leave Blank - It is only used for corrected IAT entries
	IATIndicator string `json:"IATIndicator,omitempty"`

	// ForeignExchangeIndicator is a code indicating currency conversion
	//
	// FV Fixed-to-Variable – Entry is originated in a fixed-value amount
	// and is to be received in a variable amount resulting from the
	// execution of the foreign exchange conversion.
	//
	// VF Variable-to-Fixed – Entry is originated in a variable-value
	// amount based on a specific foreign exchange rate for conversion to a
	// fixed-value amount in which the entry is to be received.
	//
	// FF Fixed-to-Fixed – Entry is originated in a fixed-value amount and
	// is to be received in the same fixed-value amount in the same
	// currency denomination. There is no foreign exchange conversion for
	// entries transmitted using this code. For entries originated in a fixed value
	// amount, the foreign Exchange Reference Field will be space
	// filled.
	ForeignExchangeIndicator string `json:"foreignExchangeIndicator"`

	// ForeignExchangeReferenceIndicator is a code used to indicate the content of the
	// Foreign Exchange Reference Field and is filled by the gateway operator.
	// Valid entries are:
	// 1 - Foreign Exchange Rate;
	// 2 - Foreign Exchange Reference Number; or
	// 3 - Space Filled
	ForeignExchangeReferenceIndicator int `json:"foreignExchangeReferenceIndicator"`

	// ForeignExchangeReference  Contains either the foreign exchange rate used to execute
	// the foreign exchange conversion of a cross-border entry or another reference to the foreign
	// exchange transaction.
	ForeignExchangeReference string `json:"foreignExchangeReference"`

	// ISODestinationCountryCode is the two-character code, as approved by the International
	// Organization for Standardization (ISO), to identify the country in which the entry is
	// to be received. Values can be found on the International Organization for Standardization
	// website: www.iso.org.  For entries destined to account holder in the U.S., this would be US.
	ISODestinationCountryCode string `json:"ISODestinationCountryCode"`

	// OriginatorIdentification identifies the following:
	// For U.S. entities: the number assigned will be your tax ID
	// For non-U.S. entities: the number assigned will be your DDA number,
	// or the last 9 characters of your account number if it exceeds 9 characters
	OriginatorIdentification string `json:"originatorIdentification"`

	// StandardEntryClassCode for consumer and non consumer international payments is IAT
	// Identifies the payment type (product) found within an ACH batch-using a 3-character code.
	// The SEC Code pertains to all items within batch.
	// Determines format of the detail records.
	// Determines addenda records (required or optional PLUS one or up to 9,999 records).
	// Determines rules to follow (return time frames).
	// Some SEC codes require specific data in predetermined fields within the ACH record
	StandardEntryClassCode string `json:"standardEntryClassCode"`

	// CompanyEntryDescription A description of the entries contained in the batch
	//
	//The Originator establishes the value of this field to provide a
	// description of the purpose of the entry to be displayed back to
	// the receive For example, "GAS BILL," "REG. SALARY," "INS. PREM,"
	// "SOC. SEC.," "DTC," "TRADE PAY," "PURCHASE," etc.
	//
	// This field must contain the word "REVERSAL" (left justified) when the
	// batch contains reversing entries.
	//
	// This field must contain the word "RECLAIM" (left justified) when the
	// batch contains reclamation entries.
	//
	// This field must contain the word "NONSETTLED" (left justified) when the
	// batch contains entries which could not settle.
	CompanyEntryDescription string `json:"companyEntryDescription,omitempty"`

	// ISOOriginatingCurrencyCode is the three-character code, as approved by the International
	// Organization for Standardization (ISO), to identify the currency denomination in which the
	// entry was first originated. If the source of funds is within the territorial jurisdiction
	// of the U.S., enter 'USD', otherwise refer to International Organization for Standardization
	// website for value: www.iso.org -- (Account Currency)
	ISOOriginatingCurrencyCode string `json:"ISOOriginatingCurrencyCode"`

	// ISODestinationCurrencyCode is the three-character code, as approved by the International
	// Organization for Standardization (ISO), to identify the currency denomination in which the
	// entry will ultimately be settled. If the final destination of funds is within the territorial
	// jurisdiction of the U.S., enter “USD”, otherwise refer to International Organization for
	// Standardization website for value: www.iso.org -- (Payment Currency)
	ISODestinationCurrencyCode string `json:"ISODestinationCurrencyCode"`

	// EffectiveEntryDate the date on which the entries are to settle. Format: YYMMDD (Y=Year, M=Month, D=Day)
	EffectiveEntryDate string `json:"effectiveEntryDate,omitempty"`

	// SettlementDate Leave blank, this field is inserted by the ACH operator
	SettlementDate string `json:"settlementDate,omitempty"`

	// OriginatorStatusCode refers to the ODFI initiating the Entry.
	// 0 ADV File prepared by an ACH Operator.
	// 1 This code identifies the Originator as a depository financial institution.
	// 2 This code identifies the Originator as a Federal Government entity or agency.
	OriginatorStatusCode int `json:"originatorStatusCode,omitempty"`

	// ODFIIdentification First 8 digits of the originating DFI transit routing number
	// For Inbound IAT Entries, this field contains the routing number of the U.S. Gateway
	// Operator.  For Outbound IAT Entries, this field contains the standard routing number,
	// as assigned by Accuity, that identifies the U.S. ODFI initiating the Entry.
	// Format - TTTTAAAA
	ODFIIdentification string `json:"ODFIIdentification"`

	// BatchNumber is assigned in ascending sequence to each batch by the ODFI
	// or its Sending Point in a given file of entries. Since the batch number
	// in the Batch Header Record and the Batch Control Record is the same,
	// the ascending sequence number should be assigned by batch and not by
	// record.
	BatchNumber int `json:"batchNumber"`

	// validator is composed for data validation
	validator

	// converters is composed for ACH to golang Converters
	converters
}

const (
	// IATCOR is the valid value for IATBatchHeader.IATIndicator for IAT Notification Of Changr
	IATCOR = "IATCOR"
)

// NewIATBatchHeader returns a new BatchHeader with default values for non exported fields
func NewIATBatchHeader() *IATBatchHeader {
	iatBh := &IATBatchHeader{
		OriginatorStatusCode: 0, //Prepared by an Originator
		BatchNumber:          1,
	}
	return iatBh
}

// Parse takes the input record string and parses the BatchHeader values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (iatBh *IATBatchHeader) Parse(record string) {
	if utf8.RuneCountInString(record) != 94 {
		return
	}

	// 1-1 Always "5"
	// 2-4 MixedCreditsAnDebits (220), CReditsOnly 9220), DebitsOnly (225)"
	iatBh.ServiceClassCode = iatBh.parseNumField(record[1:4])
	// 05-20  Blank except for corrected IAT entries
	iatBh.IATIndicator = iatBh.parseStringField(record[4:20])
	// 21-22 A code indicating currency conversion
	// “FV” Fixed-to-Variable
	// “VF” Variable-to-Fixed
	// “FF” Fixed-to-Fixed
	iatBh.ForeignExchangeIndicator = iatBh.parseStringField(record[20:22])
	// 23-23 Foreign Exchange Reference Indicator – Refers to “Foreign Exchange Reference”
	// field and is filled by the gateway operator. Valid entries are:
	// 1 - Foreign Exchange Rate;
	// 2 - Foreign Exchange Reference Number; or
	// 3 - Space Filled
	iatBh.ForeignExchangeReferenceIndicator = iatBh.parseNumField(record[22:23])
	// 24-38 Contains either the foreign exchange rate used to execute the
	// foreign exchange conversion of a cross-border entry or another
	// reference to the foreign exchange transaction.
	iatBh.ForeignExchangeReference = iatBh.parseStringField(record[23:38])
	// 39-40  Receiver ISO Country Code - For entries
	// destined to account holder in the U.S., this would be 'US'.
	iatBh.ISODestinationCountryCode = iatBh.parseStringField(record[38:40])
	// 41-50 For U.S. entities: the number assigned will be your tax ID
	// For non-U.S. entities: the number assigned will be your DDA number,
	// or the last 9 characters of your account number if it exceeds 9 characters
	iatBh.OriginatorIdentification = iatBh.parseStringField(record[40:50])
	// 51-53 IAT for both consumer and non consumer international payments
	iatBh.StandardEntryClassCode = record[50:53]
	// 54-63 Your description of the transaction. This text will appear on the receivers' bank statement.
	// For example: "Payroll   "
	iatBh.CompanyEntryDescription = strings.TrimSpace(record[53:63])
	// 64-66 Originator ISO Currency Code
	iatBh.ISOOriginatingCurrencyCode = iatBh.parseStringField(record[63:66])
	// 67-69 Receiver ISO Currency Code
	iatBh.ISODestinationCurrencyCode = iatBh.parseStringField(record[66:69])
	// 70-75 Date transactions are to be posted to the receivers' account.
	// You almost always want the transaction to post as soon as possible, so put tomorrow's date in YYMMDD format
	iatBh.EffectiveEntryDate = iatBh.validateSimpleDate(record[69:75])
	// 76-78 Always blank (just fill with spaces)
	iatBh.SettlementDate = iatBh.validateSettlementDate(record[75:78])
	// 79-79 Always 1
	iatBh.OriginatorStatusCode = iatBh.parseNumField(record[78:79])
	// 80-87 Your ODFI's routing number without the last digit. The last digit is simply a
	// checksum digit, which is why it is not necessary
	iatBh.ODFIIdentification = iatBh.parseStringField(record[79:87])
	// 88-94 Sequential number of this Batch Header Record
	// For example, put "1" if this is the first Batch Header Record in the file
	iatBh.BatchNumber = iatBh.parseNumField(record[87:94])
}

// String writes the BatchHeader struct to a 94 character string.
func (iatBh *IATBatchHeader) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(batchHeaderPos)
	buf.WriteString(fmt.Sprintf("%v", iatBh.ServiceClassCode))
	buf.WriteString(iatBh.IATIndicatorField())
	buf.WriteString(iatBh.ForeignExchangeIndicatorField())
	buf.WriteString(iatBh.ForeignExchangeReferenceIndicatorField())
	buf.WriteString(iatBh.ForeignExchangeReferenceField())
	buf.WriteString(iatBh.ISODestinationCountryCodeField())
	buf.WriteString(iatBh.OriginatorIdentificationField())
	buf.WriteString(iatBh.StandardEntryClassCode)
	buf.WriteString(iatBh.CompanyEntryDescriptionField())
	buf.WriteString(iatBh.ISOOriginatingCurrencyCodeField())
	buf.WriteString(iatBh.ISODestinationCurrencyCodeField())
	buf.WriteString(iatBh.EffectiveEntryDateField())
	buf.WriteString(iatBh.SettlementDateField())
	buf.WriteString(fmt.Sprintf("%v", iatBh.OriginatorStatusCode))
	buf.WriteString(iatBh.ODFIIdentificationField())
	buf.WriteString(iatBh.BatchNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (iatBh *IATBatchHeader) Validate() error {
	if err := iatBh.fieldInclusion(); err != nil {
		return err
	}
	if err := iatBh.isServiceClass(iatBh.ServiceClassCode); err != nil {
		return fieldError("ServiceClassCode", err, strconv.Itoa(iatBh.ServiceClassCode))
	}
	if err := iatBh.isForeignExchangeIndicator(iatBh.ForeignExchangeIndicator); err != nil {
		return fieldError("ForeignExchangeIndicator", err, iatBh.ForeignExchangeIndicator)
	}
	if err := iatBh.isForeignExchangeReferenceIndicator(iatBh.ForeignExchangeReferenceIndicator); err != nil {
		return fieldError("ForeignExchangeReferenceIndicator", err, strconv.Itoa(iatBh.ForeignExchangeReferenceIndicator))
	}
	if !iso3166.Valid(iatBh.ISODestinationCountryCode) {
		return fieldError("ISODestinationCountryCode", ErrValidISO3166, iatBh.ISODestinationCountryCode)
	}
	if err := iatBh.isSECCode(iatBh.StandardEntryClassCode); err != nil {
		return fieldError("StandardEntryClassCode", err, iatBh.StandardEntryClassCode)
	}
	if err := iatBh.isAlphanumeric(iatBh.CompanyEntryDescription); err != nil {
		return fieldError("CompanyEntryDescription", err, iatBh.CompanyEntryDescription)
	}
	if !iso4217.Valid(iatBh.ISOOriginatingCurrencyCode) {
		return fieldError("ISOOriginatingCurrencyCode", ErrValidISO4217, iatBh.ISOOriginatingCurrencyCode)
	}
	if !iso4217.Valid(iatBh.ISODestinationCurrencyCode) {
		return fieldError("ISODestinationCurrencyCode", ErrValidISO4217, iatBh.ISODestinationCurrencyCode)
	}
	if err := iatBh.isOriginatorStatusCode(iatBh.OriginatorStatusCode); err != nil {
		return fieldError("OriginatorStatusCode", err, strconv.Itoa(iatBh.OriginatorStatusCode))
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (iatBh *IATBatchHeader) fieldInclusion() error {
	if iatBh.ServiceClassCode == 0 {
		return fieldError("ServiceClassCode", ErrFieldInclusion, strconv.Itoa(iatBh.ServiceClassCode))
	}
	if iatBh.ForeignExchangeIndicator == "" {
		return fieldError("ForeignExchangeIndicator", ErrFieldInclusion, iatBh.ForeignExchangeIndicator)
	}
	if iatBh.ForeignExchangeReferenceIndicator == 0 {
		return fieldError("ForeignExchangeReferenceIndicator", ErrFieldRequired, strconv.Itoa(iatBh.ForeignExchangeReferenceIndicator))
	}
	// ToDo: It can be space filled based on ForeignExchangeReferenceIndicator just use a validator to handle -
	// ToDo: Calling Field ok for validation?
	/*	if iatBh.ForeignExchangeReference == "" {
		return fieldError("ForeignExchangeReference", ErrFieldRequired, iatBh.ForeignExchangeReference)
	}*/
	if iatBh.ISODestinationCountryCode == "" {
		return fieldError("ISODestinationCountryCode", ErrFieldInclusion, iatBh.ISODestinationCountryCode)
	}
	if iatBh.OriginatorIdentification == "" {
		return fieldError("OriginatorIdentification", ErrFieldInclusion, iatBh.OriginatorIdentification)
	}
	if iatBh.StandardEntryClassCode == "" {
		return fieldError("StandardEntryClassCode", ErrFieldInclusion, iatBh.StandardEntryClassCode)
	}
	if iatBh.CompanyEntryDescription == "" {
		return fieldError("CompanyEntryDescription", ErrFieldInclusion, iatBh.CompanyEntryDescription)
	}
	if iatBh.ISOOriginatingCurrencyCode == "" {
		return fieldError("ISOOriginatingCurrencyCode", ErrFieldInclusion, iatBh.ISOOriginatingCurrencyCode)
	}
	if iatBh.ISODestinationCurrencyCode == "" {
		return fieldError("ISODestinationCurrencyCode", ErrFieldInclusion, iatBh.ISODestinationCurrencyCode)
	}
	if iatBh.ODFIIdentification == "" {
		return fieldError("ODFIIdentification", ErrFieldInclusion, iatBh.ODFIIdentificationField())
	}
	return nil
}

// IATIndicatorField gets the IATIndicator left padded
func (iatBh *IATBatchHeader) IATIndicatorField() string {
	// should this be left padded
	return iatBh.alphaField(iatBh.IATIndicator, 16)
}

// ForeignExchangeIndicatorField gets the ForeignExchangeIndicator
func (iatBh *IATBatchHeader) ForeignExchangeIndicatorField() string {
	return iatBh.alphaField(iatBh.ForeignExchangeIndicator, 2)
}

// ForeignExchangeReferenceIndicatorField gets the ForeignExchangeReferenceIndicator
func (iatBh *IATBatchHeader) ForeignExchangeReferenceIndicatorField() string {
	return iatBh.numericField(iatBh.ForeignExchangeReferenceIndicator, 1)
}

// ForeignExchangeReferenceField gets the ForeignExchangeReference left padded
func (iatBh *IATBatchHeader) ForeignExchangeReferenceField() string {
	if iatBh.ForeignExchangeReferenceIndicator == 3 {
		//blank space
		return "               "
	}
	return iatBh.alphaField(iatBh.ForeignExchangeReference, 15)
}

// ISODestinationCountryCodeField gets the ISODestinationCountryCode
func (iatBh *IATBatchHeader) ISODestinationCountryCodeField() string {
	return iatBh.alphaField(iatBh.ISODestinationCountryCode, 2)
}

// OriginatorIdentificationField gets the OriginatorIdentification left padded
func (iatBh *IATBatchHeader) OriginatorIdentificationField() string {
	return iatBh.alphaField(iatBh.OriginatorIdentification, 10)
}

// CompanyEntryDescriptionField gets the CompanyEntryDescription left padded
func (iatBh *IATBatchHeader) CompanyEntryDescriptionField() string {
	return iatBh.alphaField(iatBh.CompanyEntryDescription, 10)
}

// ISOOriginatingCurrencyCodeField gets the ISOOriginatingCurrencyCode
func (iatBh *IATBatchHeader) ISOOriginatingCurrencyCodeField() string {
	return iatBh.alphaField(iatBh.ISOOriginatingCurrencyCode, 3)
}

// ISODestinationCurrencyCodeField gets the ISODestinationCurrencyCode
func (iatBh *IATBatchHeader) ISODestinationCurrencyCodeField() string {
	return iatBh.alphaField(iatBh.ISODestinationCurrencyCode, 3)
}

// EffectiveEntryDateField get the EffectiveEntryDate in YYMMDD format
func (iatBh *IATBatchHeader) EffectiveEntryDateField() string {
	return iatBh.stringField(iatBh.EffectiveEntryDate, 6) // YYMMDD
}

// ODFIIdentificationField get the odfi number zero padded
func (iatBh *IATBatchHeader) ODFIIdentificationField() string {
	return iatBh.stringField(iatBh.ODFIIdentification, 8)
}

// BatchNumberField get the batch number zero padded
func (iatBh *IATBatchHeader) BatchNumberField() string {
	return iatBh.numericField(iatBh.BatchNumber, 7)
}

// SettlementDateField gets the SettlementDate
func (iatBh *IATBatchHeader) SettlementDateField() string {
	return iatBh.alphaField(iatBh.SettlementDate, 3)
}
