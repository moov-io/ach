// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strconv"
)

// BatchCTX holds the BatchHeader and BatchControl and all EntryDetail for CTX Entries.
//
// The Corporate Trade Exchange (CTX) application provides the ability to collect and disburse
// funds and information between companies. Generally it is used by businesses paying one another
// for goods or services. These payments replace checks with an electronic process of debiting and
// crediting invoices between the financial institutions of participating companies.
type BatchCTX struct {
	batch
}

var (
	msgBatchCTXAddenda      = "9999 is the maximum addenda records for SEC code CTX"
	msgBatchCTXAddendaCount = "%v entry detail addenda records not equal to addendum %v"
	msgBatchCTXAddendaType  = "%T found where Addenda05 is required for SEC code CTX"
)

// NewBatchCTX returns a *BatchCTX
func NewBatchCTX(bh *BatchHeader) *BatchCTX {
	batch := new(BatchCTX)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchCTX) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration based validation for this type.

	// Add type specific validation.

	if batch.Header.StandardEntryClassCode != "CTX" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "CTX")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {

		// Addenda validations - CTX Addenda must be Addenda05

		// A maximum of 9999 addenda records for CTX entry details
		if len(entry.Addendum) > 9999 {
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addendum", Msg: msgBatchCTXAddenda}
		}

		addendaRecords, _ := strconv.Atoi(entry.CTXAddendaRecordsField())
		if len(entry.Addendum) != addendaRecords {
			msg := fmt.Sprintf(msgBatchCTXAddendaCount, addendaRecords, len(entry.Addendum))
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addendum", Msg: msg}
		}

		if len(entry.Addendum) > 0 {
			for i := range entry.Addendum {
				addenda05, ok := entry.Addendum[i].(*Addenda05)
				if !ok {
					msg := fmt.Sprintf(msgBatchCTXAddendaType, entry.Addendum[i])
					return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addendum", Msg: msg}
				}
				if err := addenda05.Validate(); err != nil {
					// convert the field error in to a batch error for a consistent api
					if e, ok := err.(*FieldError); ok {
						return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: e.FieldName, Msg: e.Msg}
					}
				}
			}
		}
	}
	return nil
}

// Create takes Batch Header and Entries and builds a valid batch
func (batch *BatchCTX) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}
