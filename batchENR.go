// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strconv"
	"strings"
)

// BatchENR is a non-monetary entry that enrolls a person with an agency of the US government
// for a depository financial institution.
//
// Allowed TransactionCode values: 22 Demand Credit, 27 Demand Debit, 32 Savings Credit, 37 Savings Debit
type BatchENR struct {
	batch
}

func NewBatchENR(bh *BatchHeader) *BatchENR {
	batch := new(BatchENR)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate ensures the batch meets NACHA rules specific to this batch type.
func (batch *BatchENR) Validate() error {
	if err := batch.verify(); err != nil {
		return err
	}

	// Batch Header checks
	if batch.Header.StandardEntryClassCode != "ENR" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "ENR")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}
	if batch.Header.CompanyEntryDescription != "AUTOENROLL" {
		msg := fmt.Sprintf(msgBatchCompanyEntryDescription, batch.Header.CompanyEntryDescription, "ENR, must be AUTOENROLL")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "CompanyEntryDescription", Msg: msg}
	}

	// Range over Entries
	for _, entry := range batch.Entries {
		if err := entry.Validate(); err != nil {
			return err
		}

		if entry.Amount != 0 {
			msg := fmt.Sprintf(msgBatchAmountZero, entry.Amount, "ENR")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Amount", Msg: msg}
		}

		switch entry.TransactionCode {
		case 22, 27, 32, 37:
		default:
			msg := fmt.Sprintf(msgBatchTransactionCode, entry.TransactionCode, "ENR")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}

		// ENR must have one Addenda05
		// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode
		if err := batch.addendaFieldInclusion(entry); err != nil {
			return err
		}
	}
	return nil
}

// Create builds the batch sequence numbers and batch control.
func (batch *BatchENR) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	return batch.Validate()
}

type ENRPaymentInformation struct {
	// TransactionCode is the Transaction Code of the holder's account
	// Values: 22 (Demand  Credit), 27 (Demand Debit), 32 (Savings Credit), 37 (Savings Debit)
	TransactionCode int

	// RDFIIdentification is the Receiving Depository Identification Number. Typically the first 8 of their ABA routing number.
	RDFIIdentification string

	// CheckDigit is the last digit from an ABA routing number.
	CheckDigit string

	// DFIAccountNumber contains the holder's account number.
	DFIAccountNumber string

	// IndividualIdentification contains the customer's Social Security Number (SSN) for automated enrollments and the
	// taxpayer ID for companies.
	IndividualIdentification string

	// IndividualName is the account holders full name.
	IndividualName string

	// EnrolleeClassificationCode (also called Representative Payee Indicator) returns a code from a specific Addenda05 record.
	// These codes represent:
	//  0: (no)  - Initiated by beneficiary
	//  1: (yes) - Initiated by someone other than named beneficiary
	//  A: Enrollee is a consumer
	//  b: Enrollee is a company
	EnrolleeClassificationCode int
}

func (info *ENRPaymentInformation) String() string {
	line := "TransactionCode: %d, RDFIIdentification: %s, CheckDigit: %s, DFIAccountNumber: %s, IndividualIdentification: %v, IndividualName: %s, EnrolleeClassificationCode: %d"
	return fmt.Sprintf(line, info.TransactionCode, info.RDFIIdentification, info.CheckDigit, info.DFIAccountNumber, info.IndividualIdentification != "", info.IndividualName, info.EnrolleeClassificationCode)
}

// ParsePaymentInformation returns an ENRPaymentInformation for a given Addenda05 record. The information is parsed from the addenda's
// PaymentRelatedInformation field.
//
// The returned information is not validated for correctness.
func (batch *BatchENR) ParsePaymentInformation(addenda05 *Addenda05) (*ENRPaymentInformation, error) {
	parts := strings.Split(strings.TrimSuffix(addenda05.PaymentRelatedInformation, `\`), "*") // PaymentRelatedInformation is terminated by '\'
	if len(parts) != 8 {
		return nil, fmt.Errorf("ENR: unable to parse Addenda05 (%s) PaymentRelatedInformation", addenda05.ID)
	}

	txCode, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("ENR: unable to parse TransactionCode (%s) from Addenda05.ID=%s", parts[0], addenda05.ID)
	}
	enrolleeCode, err := strconv.Atoi(parts[7])
	if err != nil {
		return nil, fmt.Errorf("ENR: unable to parse EnrolleeClassificationCode (%s) from Addenda05.ID=%s", parts[7], addenda05.ID)
	}

	return &ENRPaymentInformation{
		TransactionCode:            txCode,
		RDFIIdentification:         parts[1],
		CheckDigit:                 parts[2],
		DFIAccountNumber:           parts[3],
		IndividualIdentification:   parts[4],
		IndividualName:             fmt.Sprintf("%s %s", parts[6], parts[5]),
		EnrolleeClassificationCode: enrolleeCode,
	}, nil
}
