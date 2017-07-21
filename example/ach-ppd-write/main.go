package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/moov-io/ach"
)

func main() {
	// Example transfer to write an ACH PPD file to send/credit a external institutions account
	// Important: All financial instituions are different and will require registration and exact field values.

	// Set originator bank ODFI and destination Operator for the financial institution
	// this is the funding/recieving source of the transfer
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = 9876543210 // A blank space followed by your ODFI's transit/routing numbe
	fh.ImmediateOrigin = 1234567890      // Organization or Company FED ID usually 1 and FEIN/SSN. Assigned by your ODFI
	fh.FileCreationDate = time.Now()     // Todays Date
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"

	// BatchHeader identifies the originating entity and the type of transactions contained in the batch
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = 220          // ACH credit pushes money out, 225 debits/pulls money in.
	bh.CompanyName = "Name on Account" // The name of the company/person that has relationship with reciever
	bh.CompanyIdentification = strconv.Itoa(fh.ImmediateOrigin)
	bh.StandardEntryClassCode = "PPD"         // Consumer destination vs Company CCD
	bh.CompanyEntryDescription = "REG.SALARY" // will be on recieving accounts statement
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1)
	bh.ODFIIdentification = 12345678 // first 8 digits of your bank account

	// Identifies the recievers account information
	// can be multiple entrys per batch
	entry := ach.NewEntryDetail()
	// Identifies the entry as a debit and credit entry AND to what type of account (Savings, DDA, Loan
	entry.TransactionCode = 22                     // Code 22: Credit (deposit) to checking account
	entry.SetRDFI(9101298)                         // Recievers bank transit rounting number
	entry.DFIAccountNumber = "12345678"            // Recievers bank account number
	entry.Amount = 100000000                       // Amount of transaction with no decimil. One dollar and eleven cents = 111
	entry.IndividualName = "Reciever Account Name" // Identifies the reciever of the transaction

	// build the batch
	batch := ach.NewBatchPPD()
	batch.SetHeader(bh)
	batch.AddEntry(entry)
	if err := batch.Build(); err != nil {
		log.Fatalf("Unexpected error building batch: %s\n", err)
	}

	// build the file
	file := ach.NewFile()
	file.SetHeader(fh)
	file.AddBatch(batch)
	if err := file.Build(); err != nil {
		log.Fatalf("Unexpected error building file: %s\n", err)
	}

	// write the file to std out. Anything io.Writer
	w := ach.NewWriter(os.Stdout)
	if err := w.WriteAll([]*ach.File{file}); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()
}
