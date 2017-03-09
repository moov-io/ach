package main

import (
	"fmt"
	"os"
	"time"

	"github.com/moov-io/ach"
)

func main() {
	// Example transfer to write an ACH PPD file to send/credit a external institutions account
	// Important: All financial instituions are different and will require registration and exact field values.

	// Set originator bank ODFI and destination Operator for the financial institution
	// this is the funding/recieving source of the transfer
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = 9876543210 // Your Bank transit/routing number
	fh.ImmediateOrigin = 1234567890      // Organization or Company federal usually 1 and tax ID
	fh.FileCreationDate = time.Now()
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"

	// BatchHeader identifies the originating entity and the type of transactions contained in the batch
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = 220                  // ACH credit pushes money out, 225 debits/pulls money in.
	bh.CompanyName = "Name on Bank Account"    // The name of the company/person that has relationship with reciever
	bh.CompanyIdentification = "123456789"     // Identify the originator of the transaction. Usually: 1 followed by 9 digit tax id
	bh.StandardEntryClassCode = "PPD"          // Consumer destination vs Company CCD
	bh.CompanyEntryDescription = "Loan Payoff" // will be on recieving accounts statement
	bh.ODFIIdentification = 12345678           // first 8 digits of your bank account

	// Identifies the recievers account information
	// can be multiple entrys per batch
	entry := ach.NewEntryDetail()
	// Identifies the entry as a debit and credit entry AND to what type of account (Savings, DDA, Loan
	entry.TransactionCode = 22
	entry.SetRDFI(9101298)                            // First 8 digits of recievers bank transit rounting number
	entry.DFIAccountNumber = "12345678"               // recievers bank account number
	entry.Amount = 100000000                          // Amount of transaction with no decimil. One dollar and eleven cents = 111
	entry.IndividualName = "Name on Reciever Account" // Identifies the reciever of the transaction

	// build the batch
	batch := ach.NewBatch().SetHeader(bh)
	batch.AddEntryDetail(entry)
	batch.Build()

	// build the file
	file := ach.NewFile()
	file.SetHeader(fh)
	file.AddBatch(batch)
	file.Build()

	// ensure everything is validated
	if err := file.ValidateAll(); err != nil {
		fmt.Printf("Could not validate built file: %v", err)
	}

	// write the file to std out. Anything io.Writer
	w := ach.NewWriter(os.Stdout)
	err := w.WriteAll([]*ach.File{file})
	if err != nil {
		fmt.Printf("Unexpected error: %s\n", err)
	}
	w.Flush()

}
