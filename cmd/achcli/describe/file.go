// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package describe

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
	"unicode/utf8"

	"github.com/moov-io/ach"
)

type Opts struct {
	MaskNames          bool
	MaskAccountNumbers bool
}

func File(ww io.Writer, file *ach.File, opts *Opts) {
	if file == nil {
		return
	}
	if opts == nil {
		opts = &Opts{}
	}

	w := tabwriter.NewWriter(ww, 0, 0, 2, ' ', 0)
	defer w.Flush()

	fh, fc := file.Header, file.Control

	// FileHeader
	fmt.Fprintln(w, "  Origin\tOriginName\tDestination\tDestinationName\tFileCreationDate\tFileCreationTime")
	fmt.Fprintf(w, "  %s\t%s\t%s\t%s\t%s\t%s\n", fh.ImmediateOrigin, fh.ImmediateOriginName, fh.ImmediateDestination, fh.ImmediateDestinationName, fh.FileCreationDate, fh.FileCreationTime)

	// Batches
	for i := range file.Batches {
		fmt.Fprintln(w, "\n  BatchNumber\tSECCode\tServiceClassCode\tCompanyName\tDiscretionaryData\tIdentification\tEntryDescription\tEffectiveEntryDate\tDescriptiveDate")

		bh := file.Batches[i].GetHeader()
		if bh != nil {
			fmt.Fprintf(w, "  %d\t%s\t%d %s\t%s\t%s\t%s\t%s\t%s\t%s\n",
				bh.BatchNumber,
				bh.StandardEntryClassCode,
				bh.ServiceClassCode,
				serviceClassCodes[bh.ServiceClassCode],
				bh.CompanyName,
				bh.CompanyDiscretionaryData,
				bh.CompanyIdentification,
				bh.CompanyEntryDescription,
				bh.EffectiveEntryDate,
				bh.CompanyDescriptiveDate,
			)
		}

		entries := file.Batches[i].GetEntries()
		for j := range entries {
			fmt.Fprintln(w, "\n    TransactionCode\tRDFIIdentification\tAccountNumber\tAmount\tName\tTraceNumber\tCategory")

			e := entries[j]
			accountNumber := e.DFIAccountNumber
			if opts.MaskAccountNumbers {
				accountNumber = maskAccountNumber(strings.TrimSpace(accountNumber))
			}

			fmt.Fprintf(w, "    %d %s\t%s\t%s\t%d\t%s\t%s\t%s\n", e.TransactionCode, transactionCodes[e.TransactionCode], e.RDFIIdentification, accountNumber, e.Amount, e.IndividualName, e.TraceNumber, e.Category)

			dumpAddenda02(w, e.Addenda02)
			for i := range e.Addenda05 {
				if i == 0 {
					fmt.Fprintln(w, "\n      Addenda05")
				}
				dumpAddenda05(w, e.Addenda05[i])
			}
			dumpAddenda98(w, e.Addenda98)
			dumpAddenda99(w, e.Addenda99)
			dumpAddenda99Dishonored(w, e.Addenda99Dishonored)
			dumpAddenda99Contested(w, e.Addenda99Contested)
		}

		bc := file.Batches[i].GetControl()
		if bc != nil {
			fmt.Fprintln(w, "\n  ServiceClassCode\tEntryAddendaCount\tEntryHash\tTotalDebits\tTotalCredits\tMACCode\tODFIIdentification\tBatchNumber")
			fmt.Fprintf(w, "  %d %s\t%d\t%d\t%d\t%d\t%s\t%s\t%d\n", bc.ServiceClassCode, serviceClassCodes[bh.ServiceClassCode], bc.EntryAddendaCount, bc.EntryHash, bc.TotalDebitEntryDollarAmount, bc.TotalCreditEntryDollarAmount, bc.MessageAuthenticationCode, bc.ODFIIdentification, bc.BatchNumber)
		}
	}

	// FileControl
	fmt.Fprintln(w, "\n  BatchCount\tBlockCount\tEntryAddendaCount\tTotalDebitAmount\tTotalCreditAmount")
	fmt.Fprintf(w, "  %d\t%d\t%d\t%d\t%d\n", fc.BatchCount, fc.BlockCount, fc.EntryAddendaCount, fc.TotalDebitEntryDollarAmountInFile, fc.TotalCreditEntryDollarAmountInFile)
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

func dumpAddenda99Dishonored(w *tabwriter.Writer, a *ach.Addenda99Dishonored) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "\n      Dishonored Addenda99")
	fmt.Fprintln(w, "      Dis. ReturnCode\tOrig. TraceNumber\tRDFI Identification\tTraceNumber\tSettlementDate\tReturnCode\tAddendaInformation\tTraceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		a.DishonoredReturnReasonCode, a.OriginalEntryTraceNumber, a.OriginalReceivingDFIIdentification, a.ReturnTraceNumber,
		a.ReturnSettlementDate, a.ReturnReasonCode, a.AddendaInformation, a.TraceNumber)
}

func dumpAddenda99Contested(w *tabwriter.Writer, a *ach.Addenda99Contested) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "\n      Contested Dishonored Addenda99")
	fmt.Fprintln(w, "      ContestedReturnCode\tOrig. TraceNumber\tOrig Date Returned\tOrig. RDFIIdentification\tOrig. SettlementDate\tReturnTraceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%s\t%s\t%s\t%s\n",
		a.ContestedReturnCode, a.OriginalEntryTraceNumber, a.DateOriginalEntryReturned, a.OriginalReceivingDFIIdentification,
		a.OriginalSettlementDate, a.ReturnTraceNumber)

	fmt.Fprintln(w, "      ReturnSettlementDate\tReturnReasonCode\tDishonoredTraceNumber\tDishonoredSettlementDate\tDishonoredReasonCode\tTraceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%s\t%s\t%s\t%s\n",
		a.ReturnSettlementDate, a.ReturnReasonCode, a.DishonoredReturnTraceNumber, a.DishonoredReturnSettlementDate, a.DishonoredReturnReasonCode, a.TraceNumber)
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

func maskAccountNumber(s string) string {
	length := utf8.RuneCountInString(s)
	if length < 5 {
		return "****" // too short, we can't keep anything
	}
	return strings.Repeat("*", length-4) + s[length-4:]
}
