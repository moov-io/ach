// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/moov-io/ach"
)

func dumpFiles(paths []string) error {
	for i := range paths {
		if i > 0 {
			fmt.Println("") // extra newline between multiple ACH files
		}
		fmt.Printf("Describing ACH file '%s'\n\n", paths[i])
		file, err := readACHFile(paths[i])
		if err != nil {
			fmt.Printf("WARN: problem reading %s:\n %v\n\n", paths[i], err)
		}
		dumpFile(file)
	}
	return nil
}

func dumpFile(file *ach.File) {
	if file == nil {
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	fh, fc := file.Header, file.Control

	// FileHeader
	fmt.Fprintln(w, "  Origin\tOriginName\tDestination\tDestinationName\tFileCreationDate\tFileCreationTime")
	fmt.Fprintf(w, "  %s\t%s\t%s\t%s\t%s\t%s\n", fh.ImmediateOrigin, fh.ImmediateOriginName, fh.ImmediateDestination, fh.ImmediateDestinationName, fh.FileCreationDate, fh.FileCreationTime)

	// FileControl
	fmt.Fprintln(w, "\n  BatchCount\tBlockCount\tEntryAddendaCount\tTotalDebitAmount\tTotalCreditAmount")
	fmt.Fprintf(w, "  %d\t%d\t%d\t%d\t%d\n", fc.BatchCount, fc.BlockCount, fc.EntryAddendaCount, fc.TotalDebitEntryDollarAmountInFile, fc.TotalCreditEntryDollarAmountInFile)

	// Batches
	for i := range file.Batches {
		fmt.Fprintln(w, "\n  BatchNumber\tType\tServiceClass\tAddendas\tTotalDebits\tTotalCredits")

		bh := file.Batches[i].GetHeader()
		bc := file.Batches[i].GetControl()
		fmt.Fprintf(w, "  %d\t%s\t%d\t%d\t%d\t%d\n", bh.BatchNumber, bh.StandardEntryClassCode, bh.ServiceClassCode, bc.EntryAddendaCount, bc.TotalDebitEntryDollarAmount, bc.TotalCreditEntryDollarAmount)

		entries := file.Batches[i].GetEntries()
		for j := range entries {
			fmt.Fprintln(w, "\n    TxCode\tAccountNumber\tAmount\tName\tTraceNumber\tCategory")

			e := entries[j]
			fmt.Fprintf(w, "    %d\t%s\t%d\t%s\t%s\t%s\n", e.TransactionCode, e.DFIAccountNumber, e.Amount, e.IndividualName, e.TraceNumber, e.Category)

			dumpAddenda02(w, e.Addenda02)
			for i := range e.Addenda05 {
				if i == 0 {
					fmt.Fprintln(w, "\n      Addenda05")
				}
				dumpAddenda05(w, e.Addenda05[i])
			}
			dumpAddenda98(w, e.Addenda98)
			dumpAddenda99(w, e.Addenda99)
		}
	}
	// TODO(adam): Do different stuff with -v enabeld
}

func dumpAddenda02(w *tabwriter.Writer, a *ach.Addenda02) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "\n      Addenda02")
	fmt.Fprintln(w, "      ReferenceInfoOne\tReferenceInfoTwo\tTerminalIdentification\tTransactionSerial\tDate\tAuthCodeOrExires\tLocation\tCity\tState\tTraceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		a.ReferenceInformationOne, a.ReferenceInformationTwo, a.TerminalIdentificationCode, a.TransactionSerialNumber,
		a.TransactionDate, a.AuthorizationCodeOrExpireDate, a.TerminalLocation, a.TerminalCity, a.TerminalState, a.TraceNumber)
}

func dumpAddenda05(w *tabwriter.Writer, a *ach.Addenda05) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "      PaymentRelatedInformation\tSequenceNumber\tEntryDetailSequenceNumber")
	fmt.Fprintf(w, "      %s\t%d\t%d\n", a.PaymentRelatedInformation, a.SequenceNumber, a.EntryDetailSequenceNumber)
}

func dumpAddenda98(w *tabwriter.Writer, a *ach.Addenda98) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "\n      Addenda98")
	fmt.Fprintln(w, "      ChangeCode\tOriginalTrace\tOriginalDFI\tCorrectedData\tTraceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%s\t%s\t%s\n", a.ChangeCode, a.OriginalTrace, a.OriginalDFI, a.CorrectedData, a.TraceNumber)
}

func dumpAddenda99(w *tabwriter.Writer, a *ach.Addenda99) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "\n      Addenda99")
	fmt.Fprintln(w, "      ReturnCode\tOriginalTrace\tDateOfDeath\tOriginalDFI\tAddendaInformation\tTraceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%s\t%s\t%s\t%s\n", a.ReturnCode, a.OriginalTrace, a.DateOfDeath, a.OriginalDFI, a.AddendaInformation, a.TraceNumber)
}

func readACHFile(path string) (*ach.File, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("problem opening %s: %v", path, err)
	}
	defer fd.Close()

	f, err := ach.NewReader(fd).Read()
	return &f, err
}
