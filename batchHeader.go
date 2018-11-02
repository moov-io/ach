// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// msgServiceClass

// BatchHeader identifies the originating entity and the type of transactions
// contained in the batch (i.e., the standard entry class, PPD for consumer, CCD
// or CTX for corporate). This record also contains the effective date, or desired
// settlement date, for all entries contained in this batch. The settlement date
// field is not entered as it is determined by the ACH operator
type BatchHeader struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block. 5
	recordType string

	// ServiceClassCode ACH Mixed Debits and Credits ‘200’
	// ACH Credits Only ‘220’
	// ACH Debits Only ‘225'
	ServiceClassCode int `json:"serviceClassCode"`

	// CompanyName the company originating the entries in the batch
	CompanyName string `json:"companyName"`

	// CompanyDiscretionaryData allows Originators and/or ODFIs to include codes (one or more),
	// of significance only to them, to enable specialized handling of all
	// subsequent entries in that batch. There will be no standardized
	// interpretation for the value of the field. This field must be returned
	// intact on any return entry.
	CompanyDiscretionaryData string `json:"companyDiscretionaryData,omitempty"`

	// CompanyIdentification The 9 digit FEIN number (proceeded by a predetermined
	// alpha or numeric character) of the entity in the company name field
	CompanyIdentification string `json:"companyIdentification"`

	// StandardEntryClassCode
	// Identifies the payment type (product) found within an ACH batch-using a 3-character code.
	// The SEC Code pertains to all items within batch.
	// Determines format of the detail records.
	// Determines addenda records (required or optional PLUS one or up to 9,999 records).
	// Determines rules to follow (return time frames).
	// Some SEC codes require specific data in predetermined fields within the ACH record
	StandardEntryClassCode string `json:"standardEntryClassCode,omitempty"`

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

	// CompanyDescriptiveDate except as otherwise noted below, the Originator establishes this field
	// as the date it would like to see displayed to the receiver for
	// descriptive purposes. This field is never used to control timing of any
	// computer or manual operation. It is solely for descriptive purposes.
	// The RDFI should not assume any specific format. Examples of possible
	// entries in this field are "011392,", "01 92," "JAN 13," "JAN 92," etc.
	CompanyDescriptiveDate string `json:"companyDescriptiveDate,omitempty"`

	// EffectiveEntryDate the date on which the entries are to settle
	EffectiveEntryDate time.Time `json:"effectiveEntryDate,omitempty"`

	// SettlementDate Leave blank, this field is inserted by the ACH operator
	settlementDate string

	// OriginatorStatusCode refers to the ODFI initiating the Entry.
	// 0 ADV File prepared by an ACH Operator.
	// 1 This code identifies the Originator as a depository financial institution.
	// 2 This code identifies the Originator as a Federal Government entity or agency.
	OriginatorStatusCode int `json:"originatorStatusCode,omitempty"`

	//ODFIIdentification First 8 digits of the originating DFI transit routing number
	ODFIIdentification string `json:"ODFIIdentification"`

	// BatchNumber is assigned in ascending sequence to each batch by the ODFI
	// or its Sending Point in a given file of entries. Since the batch number
	// in the Batch Header Record and the Batch Control Record is the same,
	// the ascending sequence number should be assigned by batch and not by
	// record.
	BatchNumber int `json:"batchNumber,omitempty"`

	// validator is composed for data validation
	validator

	// converters is composed for ACH to golang Converters
	converters
}

// NewBatchHeader returns a new BatchHeader with default values for non exported fields
func NewBatchHeader() *BatchHeader {
	bh := &BatchHeader{
		recordType:           "5",
		OriginatorStatusCode: 0, //Prepared by an Originator
		BatchNumber:          1,
	}
	return bh
}

// Parse takes the input record string and parses the BatchHeader values
func (bh *BatchHeader) Parse(record string) {
	if utf8.RuneCountInString(record) != 94 {
		return
	}

	// 1-1 Always "5"
	bh.recordType = "5"
	// 2-4 If the entries are credits, always "220". If the entries are debits, always "225"
	bh.ServiceClassCode = bh.parseNumField(record[1:4])
	// 5-20 Your company's name. This name may appear on the receivers’ statements prepared by the RDFI.
	bh.CompanyName = strings.TrimSpace(record[4:20])
	// 21-40 Optional field you may use to describe the batch for internal accounting purposes
	bh.CompanyDiscretionaryData = strings.TrimSpace(record[20:40])
	// 41-50 A 10-digit number assigned to you by the ODFI once they approve you to
	// originate ACH files through them. This is the same as the "Immediate origin" field in File Header Record
	bh.CompanyIdentification = strings.TrimSpace(record[40:50])
	// 51-53 If the entries are PPD (credits/debits towards consumer account), use "PPD".
	// If the entries are CCD (credits/debits towards corporate account), use "CCD".
	// The difference between the 2 SEC codes are outside of the scope of this post.
	bh.StandardEntryClassCode = record[50:53]
	// 54-63 Your description of the transaction. This text will appear on the receivers’ bank statement.
	// For example: "Payroll   "
	bh.CompanyEntryDescription = strings.TrimSpace(record[53:63])
	// 64-69 The date you choose to identify the transactions in YYMMDD format.
	// This date may be printed on the receivers’ bank statement by the RDFI
	bh.CompanyDescriptiveDate = strings.TrimSpace(record[63:69])
	// 70-75 Date transactions are to be posted to the receivers’ account.
	// You almost always want the transaction to post as soon as possible, so put tomorrow's date in YYMMDD format
	bh.EffectiveEntryDate = bh.parseSimpleDate(record[69:75])
	// 76-79 Always blank (just fill with spaces)
	bh.settlementDate = "   "
	// 79-79 Always 1
	bh.OriginatorStatusCode = bh.parseNumField(record[78:79])
	// 80-87 Your ODFI's routing number without the last digit. The last digit is simply a
	// checksum digit, which is why it is not necessary
	bh.ODFIIdentification = bh.parseStringField(record[79:87])
	// 88-94 Sequential number of this Batch Header Record
	// For example, put "1" if this is the first Batch Header Record in the file
	bh.BatchNumber = bh.parseNumField(record[87:94])
}

// String writes the BatchHeader struct to a 94 character string.
func (bh *BatchHeader) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(bh.recordType)
	buf.WriteString(fmt.Sprintf("%v", bh.ServiceClassCode))
	buf.WriteString(bh.CompanyNameField())
	buf.WriteString(bh.CompanyDiscretionaryDataField())
	buf.WriteString(bh.CompanyIdentificationField())
	buf.WriteString(bh.StandardEntryClassCode)
	buf.WriteString(bh.CompanyEntryDescriptionField())
	buf.WriteString(bh.CompanyDescriptiveDateField())
	buf.WriteString(bh.EffectiveEntryDateField())
	buf.WriteString(bh.settlementDateField())
	buf.WriteString(fmt.Sprintf("%v", bh.OriginatorStatusCode))
	buf.WriteString(bh.ODFIIdentificationField())
	buf.WriteString(bh.BatchNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bh *BatchHeader) Validate() error {
	if err := bh.fieldInclusion(); err != nil {
		return err
	}
	if bh.recordType != "5" {
		msg := fmt.Sprintf(msgRecordType, 5)
		return &FieldError{FieldName: "recordType", Value: bh.recordType, Msg: msg}
	}
	if err := bh.isServiceClass(bh.ServiceClassCode); err != nil {
		return &FieldError{FieldName: "ServiceClassCode", Value: strconv.Itoa(bh.ServiceClassCode), Msg: err.Error()}
	}
	if err := bh.isSECCode(bh.StandardEntryClassCode); err != nil {
		return &FieldError{FieldName: "StandardEntryClassCode", Value: bh.StandardEntryClassCode, Msg: err.Error()}
	}
	if err := bh.isOriginatorStatusCode(bh.OriginatorStatusCode); err != nil {
		return &FieldError{FieldName: "OriginatorStatusCode", Value: strconv.Itoa(bh.OriginatorStatusCode), Msg: err.Error()}
	}
	if err := bh.isAlphanumeric(bh.CompanyName); err != nil {
		return &FieldError{FieldName: "CompanyName", Value: bh.CompanyName, Msg: err.Error()}
	}
	if err := bh.isAlphanumeric(bh.CompanyDiscretionaryData); err != nil {
		return &FieldError{FieldName: "CompanyDiscretionaryData", Value: bh.CompanyDiscretionaryData, Msg: err.Error()}
	}
	if err := bh.isAlphanumeric(bh.CompanyIdentification); err != nil {
		return &FieldError{FieldName: "CompanyIdentification", Value: bh.CompanyIdentification, Msg: err.Error()}
	}
	if err := bh.isAlphanumeric(bh.CompanyEntryDescription); err != nil {
		return &FieldError{FieldName: "CompanyEntryDescription", Value: bh.CompanyEntryDescription, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (bh *BatchHeader) fieldInclusion() error {
	if bh.recordType == "" {
		return &FieldError{
			FieldName: "recordType",
			Value:     bh.recordType,
			Msg:       msgFieldInclusion + ", did you use NewBatchHeader()?",
		}
	}
	if bh.ServiceClassCode == 0 {
		return &FieldError{
			FieldName: "ServiceClassCode",
			Value:     strconv.Itoa(bh.ServiceClassCode),
			Msg:       msgFieldInclusion + ", did you use NewBatchHeader()?",
		}
	}
	if bh.CompanyName == "" {
		return &FieldError{
			FieldName: "CompanyName",
			Value:     bh.CompanyName,
			Msg:       msgFieldInclusion + ", did you use NewBatchHeader()?",
		}
	}
	if bh.CompanyIdentification == "" {
		return &FieldError{
			FieldName: "CompanyIdentification",
			Value:     bh.CompanyIdentification,
			Msg:       msgFieldInclusion + ", did you use NewBatchHeader()?",
		}
	}
	if bh.StandardEntryClassCode == "" {
		return &FieldError{
			FieldName: "StandardEntryClassCode",
			Value:     bh.StandardEntryClassCode,
			Msg:       msgFieldInclusion + ", did you use NewBatchHeader()?",
		}
	}
	if bh.CompanyEntryDescription == "" {
		return &FieldError{
			FieldName: "CompanyEntryDescription",
			Value:     bh.CompanyEntryDescription,
			Msg:       msgFieldInclusion + ", did you use NewBatchHeader()?",
		}
	}
	if bh.ODFIIdentification == "" {
		return &FieldError{
			FieldName: "ODFIIdentification",
			Value:     bh.ODFIIdentificationField(),
			Msg:       msgFieldInclusion + ", did you use NewBatchHeader()?",
		}
	}
	return nil
}

// CompanyNameField get the CompanyName left padded
func (bh *BatchHeader) CompanyNameField() string {
	return bh.alphaField(bh.CompanyName, 16)
}

// CompanyDiscretionaryDataField get the CompanyDiscretionaryData left padded
func (bh *BatchHeader) CompanyDiscretionaryDataField() string {
	return bh.alphaField(bh.CompanyDiscretionaryData, 20)
}

// CompanyIdentificationField get the CompanyIdentification left padded
func (bh *BatchHeader) CompanyIdentificationField() string {
	return bh.alphaField(bh.CompanyIdentification, 10)
}

// CompanyEntryDescriptionField get the CompanyEntryDescription left padded
func (bh *BatchHeader) CompanyEntryDescriptionField() string {
	return bh.alphaField(bh.CompanyEntryDescription, 10)
}

// CompanyDescriptiveDateField get the CompanyDescriptiveDate left padded
func (bh *BatchHeader) CompanyDescriptiveDateField() string {
	return bh.alphaField(bh.CompanyDescriptiveDate, 6)
}

// EffectiveEntryDateField get the EffectiveEntryDate in YYMMDD format
func (bh *BatchHeader) EffectiveEntryDateField() string {
	// ENR records require EffectiveEntryDate to be space filled. NACHA Page OR108
	if bh.CompanyEntryDescription == "AUTOENROLL" {
		return bh.alphaField("", 6) // YYMMDD
	} else {
		return bh.formatSimpleDate(bh.EffectiveEntryDate)
	}
}

// ODFIIdentificationField get the odfi number zero padded
func (bh *BatchHeader) ODFIIdentificationField() string {
	return bh.stringField(bh.ODFIIdentification, 8)
}

// BatchNumberField get the batch number zero padded
func (bh *BatchHeader) BatchNumberField() string {
	return bh.numericField(bh.BatchNumber, 7)
}

func (bh *BatchHeader) settlementDateField() string {
	return bh.alphaField(bh.settlementDate, 3)
}
