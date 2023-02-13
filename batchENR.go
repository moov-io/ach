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
)

// BatchENR is a non-monetary entry that enrolls a person with an agency of the US government
// for a depository financial institution.
//
// Allowed TransactionCode values: 22 Demand Credit, 27 Demand Debit, 32 Savings Credit, 37 Savings Debit
type BatchENR struct {
	Batch
}

// NewBatchENR returns a *BatchENR
func NewBatchENR(bh *BatchHeader) *BatchENR {
	batch := new(BatchENR)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	batch.SetID(bh.ID)
	return batch
}

// Validate ensures the batch meets NACHA rules specific to this batch type.
func (batch *BatchENR) Validate() error {
	if err := batch.verify(); err != nil {
		return err
	}

	// Batch Header checks
	if batch.Header.StandardEntryClassCode != ENR {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, ENR)
	}
	if batch.Header.CompanyEntryDescription != "AUTOENROLL" {
		return batch.Error("CompanyEntryDescription", ErrBatchCompanyEntryDescriptionAutoenroll, batch.Header.CompanyEntryDescription)
	}

	// Range over Entries
	for _, entry := range batch.Entries {
		if err := entry.Validate(); err != nil {
			return err
		}

		if entry.Amount != 0 {
			return batch.Error("Amount", ErrBatchAmountNonZero, entry.Amount)
		}

		switch entry.TransactionCode {
		case CheckingCredit, CheckingDebit, SavingsCredit, SavingsDebit:
			// nothing
		default:
			return batch.Error("TransactionCode", ErrBatchTransactionCode, entry.TransactionCode)
		}
		// // Verify the Amount is valid for SEC code and TransactionCode
		// if err := batch.ValidAmountForCodes(entry); err != nil { // TODO(adam): https://github.com/moov-io/ach/issues/1172
		// 	return err
		// }
		// Verify the TransactionCode is valid for a ServiceClassCode
		if err := batch.ValidTranCodeForServiceClassCode(entry); err != nil {
			return err
		}
		// ENR must have one Addenda05
		// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode
		if err := batch.addendaFieldInclusion(entry); err != nil {
			return err
		}
	}
	return nil
}

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *BatchENR) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	return batch.Validate()
}

// ENRPaymentInformation structure
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
