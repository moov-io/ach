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

	// IATBatches
	for i := range file.IATBatches {
		iatBatch := file.IATBatches[i]
		bh := iatBatch.GetHeader()
		if bh != nil {
			fmt.Fprintln(w, "\n  BatchNumber\tSECCode\tServiceClassCode\tIATIndicator\tDestinationCountryCode\tFE Indicator\tFE ReferenceIndicator\tFE Reference\tCompanyEntryDescription")
			fmt.Fprintf(w, "  %d\t%s\t%d %s\t%s\t%s\t%s\t%d\t%s\t%s\n",
				bh.BatchNumber,
				bh.StandardEntryClassCode,
				bh.ServiceClassCode,
				serviceClassCodes[bh.ServiceClassCode],
				bh.IATIndicator,
				bh.ISODestinationCountryCode,
				bh.ForeignExchangeIndicator,
				bh.ForeignExchangeReferenceIndicator,
				bh.ForeignExchangeReference,
				bh.CompanyEntryDescription,
			)

			fmt.Fprintln(w, "\n    OriginatorIdentification\tISOOriginatingCurrencyCode\tISODestinationCurrencyCode\tODFIIdentification\tEffectiveEntryDate\tOriginatorStatusCode")
			fmt.Fprintf(w, "    %s\t%s\t%s\t%s\t%s\t%d\n",
				bh.OriginatorIdentification,
				bh.ISOOriginatingCurrencyCode,
				bh.ISODestinationCurrencyCode,
				bh.ODFIIdentification,
				bh.EffectiveEntryDate,
				bh.OriginatorStatusCode,
			)
		}

		entries := iatBatch.GetEntries()
		for j := range entries {
			fmt.Fprintln(w, "\n    TransactionCode\tRDFIIdentification\tAccountNumber\tAmount\tAddendaRecords\tTraceNumber\tCategory")

			e := entries[j]
			accountNumber := e.DFIAccountNumber
			if opts.MaskAccountNumbers {
				accountNumber = maskAccountNumber(strings.TrimSpace(accountNumber))
			}

			fmt.Fprintf(w, "    %d %s\t%s\t%s\t%d\t%d\t%s\t%s\n", e.TransactionCode, transactionCodes[e.TransactionCode], e.RDFIIdentification, accountNumber, e.Amount, e.AddendaRecords, e.TraceNumber, e.Category)

			dumpAddenda10(w, e.Addenda10)
			dumpAddenda11(w, e.Addenda11)
			dumpAddenda12(w, e.Addenda12)
			dumpAddenda13(w, e.Addenda13)
			dumpAddenda14(w, e.Addenda14)
			dumpAddenda15(w, e.Addenda15)
			dumpAddenda16(w, e.Addenda16)

			for i := range e.Addenda17 {
				dumpAddenda17(w, e.Addenda17[i])
			}
			for i := range e.Addenda18 {
				dumpAddenda18(w, e.Addenda18[i])
			}

			dumpAddenda98(w, e.Addenda98)
			dumpAddenda99(w, e.Addenda99)
		}

		bc := iatBatch.GetControl()
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

func dumpAddenda10(w *tabwriter.Writer, a *ach.Addenda10) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "      TypeCode\tTransactionTypeCode\tForeignPaymentAmount\tForeignTraceNumber\tName\tEntryDetailSequenceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%d\t%s\t%s\t%d\n", a.TypeCode, a.TransactionTypeCode, a.ForeignPaymentAmount, a.ForeignTraceNumber, a.Name, a.EntryDetailSequenceNumber)
}

func dumpAddenda11(w *tabwriter.Writer, a *ach.Addenda11) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "      TypeCode\tOriginatorName\tOriginatorStreetAddress\tEntryDetailSequenceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%s\t%d\n", a.TypeCode, a.OriginatorName, a.OriginatorStreetAddress, a.EntryDetailSequenceNumber)
}

func dumpAddenda12(w *tabwriter.Writer, a *ach.Addenda12) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "      TypeCode\tOriginatorCityStateProvince\tOriginatorCountryPostalCode\tEntryDetailSequenceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%s\t%d\n", a.TypeCode, a.OriginatorCityStateProvince, a.OriginatorCountryPostalCode, a.EntryDetailSequenceNumber)
}

func dumpAddenda13(w *tabwriter.Writer, a *ach.Addenda13) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "      TypeCode\tODFIName\tODFIIDNumberQualifier\tODFIIdentification\tODFIBranchCountryCode\tEntryDetailSequenceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%s\t%s\t%s\t%d\n", a.TypeCode, a.ODFIName, a.ODFIIDNumberQualifier, a.ODFIIdentification, a.ODFIBranchCountryCode, a.EntryDetailSequenceNumber)
}

func dumpAddenda14(w *tabwriter.Writer, a *ach.Addenda14) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "      TypeCode\tRDFIName\tRDFIIDNumberQualifier\tRDFIIdentification\tRDFIBranchCountryCode\tEntryDetailSequenceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%s\t%s\t%s\t%d\n", a.TypeCode, a.RDFIName, a.RDFIIDNumberQualifier, a.RDFIIdentification, a.RDFIBranchCountryCode, a.EntryDetailSequenceNumber)
}

func dumpAddenda15(w *tabwriter.Writer, a *ach.Addenda15) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "      TypeCode\tReceiverIDNumber\tReceiverStreetAddress\tEntryDetailSequenceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%s\t%d\n", a.TypeCode, a.ReceiverIDNumber, a.ReceiverStreetAddress, a.EntryDetailSequenceNumber)
}

func dumpAddenda16(w *tabwriter.Writer, a *ach.Addenda16) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "      TypeCode\tReceiverCityStateProvince\tReceiverCountryPostalCode\tEntryDetailSequenceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%s\t%d\n", a.TypeCode, a.ReceiverCityStateProvince, a.ReceiverCountryPostalCode, a.EntryDetailSequenceNumber)
}

func dumpAddenda17(w *tabwriter.Writer, a *ach.Addenda17) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "      TypeCode\tPaymentRelatedInformation\tSequenceNumber\tEntryDetailSequenceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%d\t%d\n", a.TypeCode, a.PaymentRelatedInformation, a.SequenceNumber, a.EntryDetailSequenceNumber)
}

func dumpAddenda18(w *tabwriter.Writer, a *ach.Addenda18) {
	if a == nil {
		return
	}

	fmt.Fprintln(w, "      TypeCode\tForeignCorrespondentBankName\tForeignCorrespondentBankIDNumberQualifier\tForeignCorrespondentBankIDNumber\tForeignCorrespondentBankBranchCountryCode\tSequenceNumber\tEntryDetailSequenceNumber")
	fmt.Fprintf(w, "      %s\t%s\t%s\t%s\t%s\t%d\t%d\n", a.TypeCode, a.ForeignCorrespondentBankName, a.ForeignCorrespondentBankIDNumberQualifier, a.ForeignCorrespondentBankIDNumber, a.ForeignCorrespondentBankBranchCountryCode, a.SequenceNumber, a.EntryDetailSequenceNumber)
}

func maskAccountNumber(s string) string {
	length := utf8.RuneCountInString(s)
	if length < 5 {
		return "****" // too short, we can't keep anything
	}
	return strings.Repeat("*", length-4) + s[length-4:]
}
