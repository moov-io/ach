package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/moov-io/ach"
)

func redact(args []string, dryRun bool) error {
	m := &Redactor{
		maskDigit:     9,
		maskCharacter: '*',
	}

	filename := args[0]
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("problem opening %s: %v", filename, err)
	}

	file, err := ach.NewReader(f).Read()
	if err != nil {
		return fmt.Errorf("reading file: %v", err)
	}
	dumpFile(&file)

	header := &file.Header
	control := &file.Control
	m.maskStrings(
		&header.ID,
		&header.ImmediateOrigin,
		&header.ImmediateDestination,
		&header.FileCreationDate,
		&header.FileCreationTime,
		&header.ImmediateDestinationName,
		&header.ImmediateOriginName,
		&header.ReferenceCode,
	)

	m.maskInts(
		&control.TotalCreditEntryDollarAmountInFile,
		&control.TotalDebitEntryDollarAmountInFile,
		&control.EntryHash,
	)

	// File ADV Control
	m.maskStrings(
		&file.ADVControl.ID,
	)
	m.maskInts(
		&file.ADVControl.EntryHash,
		&file.ADVControl.TotalDebitEntryDollarAmountInFile,
		&file.ADVControl.TotalCreditEntryDollarAmountInFile,
	)

	m.maskBatches(file.Batches)
	m.maskBatches(file.NotificationOfChange)
	m.maskBatches(file.ReturnEntries)
	m.maskIATBatches(file.IATBatches)

	var buf bytes.Buffer
	achWriter := ach.NewWriter(&buf)
	achWriter.BypassValidation = true
	if err := achWriter.Write(&file); err != nil {
		return fmt.Errorf("writing file: %v", err)
	}

	var w io.Writer
	if dryRun {
		w = os.Stdout
	} else { // Write to new file
		ext := filepath.Ext(filename)
		newFilename := fmt.Sprintf("%s-redacted%s", strings.TrimSuffix(filename, filepath.Ext(filename)), ext)

		w, err = os.Create(newFilename)
		if err != nil {
			return fmt.Errorf("creating new file %s: %v", newFilename, err)
		}
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		return fmt.Errorf("writing data: %v", err)
	}

	return nil
}

type Redactor struct {
	maskDigit     int
	maskCharacter rune
}

func (r *Redactor) maskStrings(strs ...*string) {
	for _, s := range strs {
		r.maskString(s)
	}
}

func (r *Redactor) maskString(s *string) {
	*s = strings.TrimSpace(*s)

	if len(*s) == 0 {
		return
	}

	tmp := *s
	allowed := 0 // todo: How many characters do we want to expose? 0, 1 or 2?
	length := utf8.RuneCountInString(tmp)
	if length < allowed {
		*s = strings.Repeat(string(r.maskCharacter), allowed)
	}
	*s = strings.Repeat(string(r.maskCharacter), length-allowed) + tmp[length-allowed:]
}

func (r *Redactor) maskInts(nums ...*int) {
	for _, n := range nums {
		r.maskInt(n)
	}
}

func (r *Redactor) maskInt(n *int) {
	if *n == 0 {
		*n = r.maskDigit
	}

	// Get number of digits
	var numDigits int
	remaining := *n
	for remaining != 0 {
		remaining /= 10

		numDigits += 1
	}

	// replace digits with maskDigit
	tmp := 0
	for i := 0; i < numDigits; i++ {
		tmp *= 10
		tmp += r.maskDigit
	}

	*n = tmp
}

func (r *Redactor) maskBatches(batches []ach.Batcher) {
	for _, b := range batches {
		if header := b.GetHeader(); header != nil {
			r.maskStrings(
				&header.ID,
				&header.CompanyName,
				&header.CompanyDiscretionaryData,
				&header.CompanyIdentification,
				&header.CompanyEntryDescription,
				//&header.StandardEntryClassCode, //todo: PII?
				&header.CompanyDescriptiveDate,
				&header.EffectiveEntryDate,
				&header.ODFIIdentification,
			)

			r.maskInts(
				//&header.ServiceClassCode, //todo: PII?
				//&header.OriginatorStatusCode, //todo: PII?
				&header.BatchNumber,
			)
		}

		// Entries
		for _, e := range b.GetEntries() {
			r.maskStrings(
				&e.ID,
				&e.RDFIIdentification,
				&e.CheckDigit,
				&e.DFIAccountNumber,
				&e.IdentificationNumber,
				&e.IndividualName,
				&e.DiscretionaryData,
				&e.TraceNumber,
				&e.Category,
			)
			r.maskInts(
				&e.Amount,
			)

			if a := e.Addenda02; a != nil {
				r.maskStrings(
					&a.ID,
					&a.ReferenceInformationOne,
					&a.ReferenceInformationTwo,
					&a.TerminalIdentificationCode,
					&a.TransactionSerialNumber,
					&a.TransactionDate,
					&a.AuthorizationCodeOrExpireDate,
					&a.TerminalLocation,
					&a.TerminalCity,
					&a.TerminalState,
					&a.TraceNumber,
				)
			}

			for _, a := range e.Addenda05 {
				r.maskStrings(
					&a.ID,
					&a.PaymentRelatedInformation,
				)
				r.maskInts(
					&a.EntryDetailSequenceNumber,
				)
			}

			if a := e.Addenda98; a != nil {
				r.maskStrings(
					&a.ID,
					//&a.ChangeCode, //todo: PII?
					&a.OriginalTrace,
					&a.OriginalDFI,
					&a.CorrectedData,
					&a.TraceNumber,
				)
			}

			r.maskAddenda99(e.Addenda99)
		}

		// ADVEntries
		for _, e := range b.GetADVEntries() {
			r.maskStrings(
				&e.ID,
				&e.RDFIIdentification,
				&e.CheckDigit,
				&e.DFIAccountNumber,
				&e.AdviceRoutingNumber,
				&e.ACHOperatorData,
				&e.ACHOperatorRoutingNumber,
				&e.FileIdentification,
				&e.IndividualName,
				&e.DiscretionaryData,
				&e.Category,
			)
			r.maskInts(
				&e.Amount,
				&e.JulianDay,
			)

			r.maskAddenda99(e.Addenda99)
		}

		r.maskBatchControl(b.GetControl())

		if control := b.GetADVControl(); control != nil {
			r.maskStrings(
				&control.ID,
				&control.ODFIIdentification,
			)
			r.maskInts(
				&control.EntryHash,
				&control.TotalDebitEntryDollarAmount,
				&control.TotalCreditEntryDollarAmount,
				&control.BatchNumber,
			)
		}
	}
}

func (r *Redactor) maskIATBatches(batches []ach.IATBatch) {
	for _, b := range batches {
		r.maskStrings(&b.ID)

		if h := b.Header; h != nil {
			r.maskStrings(
				&h.ID,
				&h.OriginatorIdentification,
				//&h.StandardEntryClassCode, // todo: PII?
				&h.CompanyEntryDescription,
				&h.EffectiveEntryDate,
				&h.ODFIIdentification,
			)

			r.maskInts(
				// todo: add additional PII fields here
				&h.BatchNumber,
			)
		}

		r.maskBatchControl(b.GetControl())

		for _, e := range b.GetEntries() {
			r.maskStrings(
				&e.ID,
				&e.RDFIIdentification,
				&e.CheckDigit,
				&e.DFIAccountNumber,
				&e.TraceNumber,
				&e.Category,
			)
			r.maskInts(
				&e.Amount,
				&e.TransactionCode,
			)

			if a := e.Addenda10; a != nil {
				r.maskStrings(
					&a.ID,
					//&a.TransactionTypeCode, // todo: PII?
					&a.ForeignTraceNumber,
					&a.Name,
				)
				r.maskInts(
					&a.ForeignPaymentAmount,
					&a.EntryDetailSequenceNumber,
				)
			}

			if a := e.Addenda11; a != nil {
				r.maskStrings(
					&a.ID,
					&a.OriginatorName,
					&a.OriginatorStreetAddress,
				)
				r.maskInts(&a.EntryDetailSequenceNumber)
			}

			if a := e.Addenda12; a != nil {
				r.maskStrings(
					&a.ID,
					&a.OriginatorCityStateProvince,
					&a.OriginatorCountryPostalCode,
				)
				r.maskInts(&a.EntryDetailSequenceNumber)
			}

			if a := e.Addenda13; a != nil {
				r.maskStrings(
					&a.ID,
					&a.ODFIName,
					&a.ODFIIDNumberQualifier,
					&a.ODFIIdentification,
					&a.ODFIBranchCountryCode,
				)
				r.maskInts(&a.EntryDetailSequenceNumber)
			}

			if a := e.Addenda14; a != nil {
				r.maskStrings(
					&a.ID,
					&a.RDFIName,
					&a.RDFIIDNumberQualifier,
					&a.RDFIIdentification,
					&a.RDFIBranchCountryCode,
				)
				r.maskInts(&a.EntryDetailSequenceNumber)
			}

			if a := e.Addenda15; a != nil {
				r.maskStrings(
					&a.ID,
					&a.ReceiverIDNumber,
					&a.ReceiverStreetAddress,
				)
				r.maskInts(&a.EntryDetailSequenceNumber)
			}

			if a := e.Addenda16; a != nil {
				r.maskStrings(
					&a.ID,
					&a.ReceiverCityStateProvince,
					&a.ReceiverCountryPostalCode,
				)
				r.maskInts(&a.EntryDetailSequenceNumber)
			}

			for _, a := range e.Addenda17 {
				r.maskStrings(
					&a.ID,
					&a.PaymentRelatedInformation,
				)
				r.maskInts(&a.EntryDetailSequenceNumber)
			}

			for _, a := range e.Addenda18 {
				r.maskStrings(
					&a.ID,
					&a.ForeignCorrespondentBankName,
					//&a.ForeignCorrespondentBankIDNumberQualifier,
					&a.ForeignCorrespondentBankIDNumber,
					&a.ForeignCorrespondentBankBranchCountryCode,
				)
				r.maskInts(&a.EntryDetailSequenceNumber)
			}

			if a := e.Addenda98; a != nil {
				r.maskStrings(
					&a.ID,
					&a.ChangeCode,
					&a.OriginalTrace,
					&a.OriginalDFI,
					&a.CorrectedData,

					&a.TraceNumber,
				)
			}

			if a := e.Addenda99; a != nil {
				r.maskStrings(
					&a.ID,
					&a.OriginalTrace,
					&a.DateOfDeath,
					&a.OriginalDFI,
					&a.AddendaInformation,
					&a.TraceNumber,
				)
			}
		}
	}
}

func (r *Redactor) maskBatchControl(c *ach.BatchControl) {
	if c == nil {
		return
	}

	r.maskStrings(
		&c.ID,
		&c.CompanyIdentification,
		&c.MessageAuthenticationCode,
		&c.ODFIIdentification,
	)
	r.maskInts(
		&c.EntryHash,
		&c.TotalDebitEntryDollarAmount,
		&c.TotalCreditEntryDollarAmount,
		&c.BatchNumber,
	)
}

func (r *Redactor) maskAddenda99(a *ach.Addenda99) {
	if a == nil {
		return
	}

	r.maskStrings(
		&a.ID,
		//&a.ReturnCode, //todo: PII?
		&a.OriginalTrace,
		&a.DateOfDeath,
		&a.AddendaInformation,
		&a.TraceNumber,
	)
}
