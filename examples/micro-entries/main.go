// Package main shows an example of creating micro-deposits to validate an account.
package main

import (
	"log"
	"os"
	"time"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/describe"
)

func main() {
	est, _ := time.LoadLocation("America/New_York")
	now := time.Now().In(est)

	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "231380104" // Routing Number of the ACH Operator or receiving point to which the file is being sent
	fh.ImmediateOrigin = "121042882"      // Routing Number of the ACH Operator or sending point that is sending the file
	fh.FileCreationDate = now.Format("060102")
	fh.ImmediateDestinationName = "Receiver Bank Name"
	fh.ImmediateOriginName = "Origin Bank Name"

	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.MixedDebitsAndCredits
	bh.CompanyName = "Name on Account" // The name of the company/person that has relationship with receiver
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.WEB                           // Or CCD
	bh.CompanyEntryDescription = "AcctVerify"                     // will be on receiving accounts statement
	bh.EffectiveEntryDate = now.AddDate(0, 0, 1).Format("060102") // YYMMDD
	bh.ODFIIdentification = "121042882"                           // Originating Routing Number

	credit1 := ach.NewEntryDetail()
	credit1.TransactionCode = ach.CheckingCredit
	credit1.SetRDFI("231380104")          // Receivers bank transit routing number
	credit1.DFIAccountNumber = "12345678" // Receivers bank account number
	credit1.Amount = 7                    // $0.07
	credit1.SetTraceNumber(bh.ODFIIdentification, 1)
	credit1.IndividualName = "Receiver Account Name" // Identifies the receiver of the transaction

	credit2 := ach.NewEntryDetail()
	credit2.TransactionCode = ach.CheckingCredit
	credit2.SetRDFI("231380104")
	credit2.DFIAccountNumber = "12345678"
	credit2.Amount = 22 // $0.22
	credit2.SetTraceNumber(bh.ODFIIdentification, 2)
	credit2.IndividualName = "Receiver Account Name"

	// Talk with your ODFI or partner bank if you need to debit your account (to fund the outgoing credits).

	// Debit the account you're validating (optional)
	debit := ach.NewEntryDetail()
	debit.TransactionCode = ach.CheckingDebit
	debit.SetRDFI("231380104")          // Receivers bank transit routing number
	debit.DFIAccountNumber = "12345678" // Receivers bank account number
	debit.Amount = credit1.Amount + credit2.Amount
	debit.SetTraceNumber(bh.ODFIIdentification, 3)
	debit.IndividualName = "Receiver Account Name" // Identifies the receiver of the transaction

	batch := ach.NewBatchWEB(bh)
	batch.AddEntry(credit1)
	batch.AddEntry(credit2)
	batch.AddEntry(debit)
	err := batch.Create()
	if err != nil {
		log.Fatalf("ERROR building WEB batch: %v", err)
	}

	file := ach.NewFile()
	file.SetHeader(fh)
	file.AddBatch(batch)
	err = file.Create()
	if err != nil {
		log.Fatalf("ERROR building file: %v", err)
	}

	describe.File(os.Stdout, *file, nil)
}
