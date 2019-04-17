// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package achutil

import (
	"bytes"
	"fmt"

	"github.com/moov-io/ach"
)

func DumpFile(file *ach.File) *bytes.Buffer {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("ID: %s\n", file.ID))
	dumpHeader(&buf, file)
	dumpBatches(&buf, file)
	dumpControl(&buf, file)

	return &buf
}

func dumpHeader(buf *bytes.Buffer, file *ach.File) {
	buf.WriteString("File Header\n")
	buf.WriteString(fmt.Sprintf("  ID: %s\n", file.Header.ID))
	buf.WriteString(fmt.Sprintf("  ImmediateDestination: %s\n", file.Header.ImmediateDestination))
	buf.WriteString(fmt.Sprintf("  ImmediateOrigin: %s\n", file.Header.ImmediateOrigin))
	buf.WriteString(fmt.Sprintf("  FileCreationDate: %s\n", file.Header.FileCreationDate))
	buf.WriteString(fmt.Sprintf("  FileCreationTime: %s\n", file.Header.FileCreationTime))
	buf.WriteString(fmt.Sprintf("  FileIDModifier: %s\n", file.Header.FileIDModifier))
	buf.WriteString(fmt.Sprintf("  ImmediateDestinationName: %s\n", file.Header.ImmediateDestinationName))
	buf.WriteString(fmt.Sprintf("  ImmediateOriginName: %s\n", file.Header.ImmediateOriginName))
	buf.WriteString(fmt.Sprintf("  ReferenceCode: %s\n", file.Header.ReferenceCode))
}

func dumpBatches(buf *bytes.Buffer, file *ach.File) {
	buf.WriteString(fmt.Sprintf("Batches: %d total\n", len(file.Batches)))
	for i := range file.Batches {
		if i > 0 {
			buf.WriteString("\n") // don't add extra newline for first batch
		}
		buf.WriteString(fmt.Sprintf("  Batch Header (#%d)\n", i))
		buf.WriteString(fmt.Sprintf("    ID: %s\n", file.Batches[i].GetHeader().ID))
		buf.WriteString(fmt.Sprintf("    ServiceClassCode: %d\n", file.Batches[i].GetHeader().ServiceClassCode))
		buf.WriteString(fmt.Sprintf("    CompanyName: %s\n", file.Batches[i].GetHeader().CompanyName))
		buf.WriteString(fmt.Sprintf("    CompanyDiscretionaryData: %s\n", file.Batches[i].GetHeader().CompanyDiscretionaryData))
		buf.WriteString(fmt.Sprintf("    CompanyIdentification: %s\n", file.Batches[i].GetHeader().CompanyIdentification))
		buf.WriteString(fmt.Sprintf("    StandardEntryClassCode: %s\n", file.Batches[i].GetHeader().StandardEntryClassCode))
		buf.WriteString(fmt.Sprintf("    CompanyEntryDescription: %s\n", file.Batches[i].GetHeader().CompanyEntryDescription))
		buf.WriteString(fmt.Sprintf("    CompanyDescriptiveDate: %s\n", file.Batches[i].GetHeader().CompanyDescriptiveDate))
		buf.WriteString(fmt.Sprintf("    EffectiveEntryDate: %s\n", file.Batches[i].GetHeader().EffectiveEntryDate))
		buf.WriteString(fmt.Sprintf("    OriginatorStatusCode: %d\n", file.Batches[i].GetHeader().OriginatorStatusCode))
		buf.WriteString(fmt.Sprintf("    ODFIIdentification: %s\n", file.Batches[i].GetHeader().ODFIIdentification))
		buf.WriteString(fmt.Sprintf("    BatchNumber: %d\n", file.Batches[i].GetHeader().BatchNumber))

		entries := file.Batches[i].GetEntries()
		for j := range entries {
			buf.WriteString(fmt.Sprintf("  Entry Detail (#%d)\n", j))
			buf.WriteString(fmt.Sprintf("    ID: %s\n", entries[j].ID))
			buf.WriteString(fmt.Sprintf("    TransactionCode: %d\n", entries[j].TransactionCode))
			buf.WriteString(fmt.Sprintf("    RDFIIdentification: %s\n", entries[j].RDFIIdentification))
			buf.WriteString(fmt.Sprintf("    CheckDigit: %s\n", entries[j].CheckDigit))
			buf.WriteString(fmt.Sprintf("    DFIAccountNumber: %s\n", entries[j].DFIAccountNumber))
			buf.WriteString(fmt.Sprintf("    Amount: %d\n", entries[j].Amount))
			buf.WriteString(fmt.Sprintf("    IdentificationNumber: %s\n", entries[j].IdentificationNumber))
			buf.WriteString(fmt.Sprintf("    IndividualName: %s\n", entries[j].IndividualName))
			buf.WriteString(fmt.Sprintf("    DiscretionaryData: %s\n", entries[j].DiscretionaryData))
			buf.WriteString(fmt.Sprintf("    AddendaRecordIndicator: %d\n", entries[j].AddendaRecordIndicator))
			buf.WriteString(fmt.Sprintf("    TraceNumber: %s\n", entries[j].TraceNumber))
			buf.WriteString(fmt.Sprintf("    Category: %s\n", entries[j].Category))

			if entries[j].Addenda02 != nil {
				buf.WriteString("    Addenda02\n")
				buf.WriteString(fmt.Sprintf("      ID: %s\n", entries[j].Addenda02.ID))
				buf.WriteString(fmt.Sprintf("      TypeCode: %s\n", entries[j].Addenda02.TypeCode))
				buf.WriteString(fmt.Sprintf("      ReferenceInformationOne: %s\n", entries[j].Addenda02.ReferenceInformationOne))
				buf.WriteString(fmt.Sprintf("      ReferenceInformationTwo: %s\n", entries[j].Addenda02.ReferenceInformationTwo))
				buf.WriteString(fmt.Sprintf("      TerminalIdentificationCode: %s\n", entries[j].Addenda02.TerminalIdentificationCode))
				buf.WriteString(fmt.Sprintf("      TransactionSerialNumber: %s\n", entries[j].Addenda02.TransactionSerialNumber))
				buf.WriteString(fmt.Sprintf("      TransactionDate: %s\n", entries[j].Addenda02.TransactionDate))
				buf.WriteString(fmt.Sprintf("      AuthorizationCodeOrExpireDate: %s\n", entries[j].Addenda02.AuthorizationCodeOrExpireDate))
				buf.WriteString(fmt.Sprintf("      TerminalLocation: %s\n", entries[j].Addenda02.TerminalLocation))
				buf.WriteString(fmt.Sprintf("      TerminalCity: %s\n", entries[j].Addenda02.TerminalCity))
				buf.WriteString(fmt.Sprintf("      TerminalState: %s\n", entries[j].Addenda02.TerminalState))
			}

			for k := range entries[j].Addenda05 {
				buf.WriteString(fmt.Sprintf("    Addenda05 (#%d)\n", k))
				buf.WriteString(fmt.Sprintf("      ID: %s\n", entries[j].Addenda05[k].ID))
				buf.WriteString(fmt.Sprintf("      TypeCode: %s\n", entries[j].Addenda05[k].TypeCode))
				buf.WriteString(fmt.Sprintf("      PaymentRelatedInformation: %s\n", entries[j].Addenda05[k].PaymentRelatedInformation))
				buf.WriteString(fmt.Sprintf("      SequenceNumber: %d\n", entries[j].Addenda05[k].SequenceNumber))
				buf.WriteString(fmt.Sprintf("      EntryDetailSequenceNumber: %d\n", entries[j].Addenda05[k].EntryDetailSequenceNumber))
			}

			if entries[j].Addenda98 != nil {
				buf.WriteString("    Addenda98\n")
				buf.WriteString(fmt.Sprintf("      ID: %s\n", entries[j].Addenda98.ID))
				buf.WriteString(fmt.Sprintf("      TypeCode: %s\n", entries[j].Addenda98.TypeCode))
				buf.WriteString(fmt.Sprintf("      ChangeCode: %s\n", entries[j].Addenda98.ChangeCode))
				buf.WriteString(fmt.Sprintf("      OriginalTrace: %s\n", entries[j].Addenda98.OriginalTrace))
				buf.WriteString(fmt.Sprintf("      OriginalDFI: %s\n", entries[j].Addenda98.OriginalDFI))
				buf.WriteString(fmt.Sprintf("      CorrectedData: %s\n", entries[j].Addenda98.CorrectedData))
				buf.WriteString(fmt.Sprintf("      TraceNumber: %s\n", entries[j].Addenda98.TraceNumber))
			}

			if entries[j].Addenda99 != nil {
				buf.WriteString("    Addenda99\n")
				buf.WriteString(fmt.Sprintf("      ID: %s\n", entries[j].Addenda99.ID))
				buf.WriteString(fmt.Sprintf("      TypeCode: %s\n", entries[j].Addenda99.TypeCode))
				buf.WriteString(fmt.Sprintf("      ReturnCode: %s\n", entries[j].Addenda99.ReturnCode))
				buf.WriteString(fmt.Sprintf("      OriginalTrace: %s\n", entries[j].Addenda99.OriginalTrace))
				buf.WriteString(fmt.Sprintf("      OriginalDFI: %s\n", entries[j].Addenda99.OriginalDFI))
				buf.WriteString(fmt.Sprintf("      AddendaInformation: %s\n", entries[j].Addenda99.AddendaInformation))
				buf.WriteString(fmt.Sprintf("      TraceNumber: %s\n", entries[j].Addenda99.TraceNumber))
			}
		}

		buf.WriteString("  Batch Control\n")
		buf.WriteString(fmt.Sprintf("    ID: %s\n", file.Batches[i].GetControl().ID))
		buf.WriteString(fmt.Sprintf("    ServiceClassCode: %d\n", file.Batches[i].GetControl().ServiceClassCode))
		buf.WriteString(fmt.Sprintf("    EntryAddendaCount: %d\n", file.Batches[i].GetControl().EntryAddendaCount))
		buf.WriteString(fmt.Sprintf("    EntryHash: %d\n", file.Batches[i].GetControl().EntryHash))
		buf.WriteString(fmt.Sprintf("    TotalDebitEntryDollarAmount: %d\n", file.Batches[i].GetControl().TotalDebitEntryDollarAmount))
		buf.WriteString(fmt.Sprintf("    TotalCreditEntryDollarAmount: %d\n", file.Batches[i].GetControl().TotalCreditEntryDollarAmount))
		buf.WriteString(fmt.Sprintf("    CompanyIdentification: %s\n", file.Batches[i].GetControl().CompanyIdentification))
		buf.WriteString(fmt.Sprintf("    MessageAuthenticationCode: %s\n", file.Batches[i].GetControl().MessageAuthenticationCode))
		buf.WriteString(fmt.Sprintf("    ODFIIdentification: %s\n", file.Batches[i].GetControl().ODFIIdentification))
		buf.WriteString(fmt.Sprintf("    BatchNumber: %d\n", file.Batches[i].GetControl().BatchNumber))
	}
}

func dumpControl(buf *bytes.Buffer, file *ach.File) {
	buf.WriteString("File Control\n")
	buf.WriteString(fmt.Sprintf("  ID: %s\n", file.Control.ID))
	buf.WriteString(fmt.Sprintf("  BatchCount: %d\n", file.Control.BatchCount))
	buf.WriteString(fmt.Sprintf("  BlockCount: %d\n", file.Control.BlockCount))
	buf.WriteString(fmt.Sprintf("  EntryAddendaCount: %d\n", file.Control.EntryAddendaCount))
	buf.WriteString(fmt.Sprintf("  EntryHash: %d\n", file.Control.EntryHash))
	buf.WriteString(fmt.Sprintf("  TotalDebitEntryDollarAmountInFile: %d\n", file.Control.TotalDebitEntryDollarAmountInFile))
	buf.WriteString(fmt.Sprintf("  TotalCreditEntryDollarAmountInFile: %d\n", file.Control.TotalCreditEntryDollarAmountInFile))
}
