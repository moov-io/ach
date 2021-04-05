package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/moov-io/ach"
)

/*
What should this command do?
- mask all PII data
- verify number of lines are the same
- write to stdOut if no output is specified
- run this against a file with 10K Lines to see if it's fast
*/
func mask(args []string) error {
	m := &Masker{
		maskDigit:     9,
		maskCharacter: '*',
	}

	args = []string{""}
	args[0] = filepath.Join("test", "testdata", "ppd-debit.ach")
	filename := args[0]
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("problem opening %s: %v", filename, err)
	}

	var buf bytes.Buffer

	var reader io.Reader = f
	reader = io.TeeReader(reader, &buf)

	file, err := ach.NewReader(reader).Read()
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

	// Write to new file
	//ext := filepath.Ext(filename)
	//newFilename := fmt.Sprintf("%s_masked%s", strings.TrimSuffix(filename, filepath.Ext(filename)), ext)
	//newFile, err := os.Create(newFilename)
	//if err != nil {
	//	return fmt.Errorf("creating new file %s: %v", newFilename, err)
	//}

	var buff bytes.Buffer
	w := ach.NewWriter(&buff)
	w.BypassValidation = true
	err = w.Write(&file)
	if err != nil {
		return fmt.Errorf("writing file: %v", err)
	}

	// Open the file again to count the number of lines
	//newFile, err = os.Open(newFilename)
	//if err != nil {
	//	return fmt.Errorf("opening file: %v", err)
	//}

	// todo: print to standard out if dry run
	oldCount := countLines(&buf)
	newCount := countLines(io.TeeReader(&buff, os.Stdout))
	if oldCount != newCount {
		return fmt.Errorf("masked file does not have the same number of lines as the original, expected %d but got %d lines", oldCount, newCount)
	}

	return nil
}

func countLines(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	count := 0

	for scanner.Scan() {
		count += 1
	}

	return count
}

type Masker struct {
	maskDigit     int
	maskCharacter rune
}

func (m *Masker) maskStrings(strs ...*string) {
	for _, s := range strs {
		m.maskString(s)
	}
}

func (m *Masker) maskString(s *string) {
	*s = strings.TrimSpace(*s)

	if len(*s) == 0 {
		return
	}

	tmp := *s
	allowed := 0 // todo: How many characters do we want to expose? 0, 1 or 2?
	length := utf8.RuneCountInString(tmp)
	if length < allowed {
		*s = strings.Repeat(string(m.maskCharacter), allowed)
	}
	*s = strings.Repeat(string(m.maskCharacter), length-allowed) + tmp[length-allowed:]
}

func (m *Masker) maskInts(nums ...*int) {
	for _, n := range nums {
		m.maskInt(n)
	}
}

func (m *Masker) maskInt(n *int) {
	if *n == 0 {
		*n = m.maskDigit
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
		tmp += m.maskDigit
	}

	*n = tmp
}

func (m *Masker) maskBatches(batches []ach.Batcher) {
	for _, b := range batches {
		if header := b.GetHeader(); header != nil {
			m.maskStrings(
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

			m.maskInts(
				//&header.ServiceClassCode, //todo: PII?
				//&header.OriginatorStatusCode, //todo: PII?
				&header.BatchNumber,
			)
		}

		// Entries
		for _, e := range b.GetEntries() {
			m.maskStrings(
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
			m.maskInts(
				&e.Amount,
			)

			if a := e.Addenda02; a != nil {
				m.maskStrings(
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
				m.maskStrings(
					&a.ID,
					&a.PaymentRelatedInformation,
				)
				m.maskInts(
					&a.EntryDetailSequenceNumber,
				)
			}

			if a := e.Addenda98; a != nil {
				m.maskStrings(
					&a.ID,
					//&a.ChangeCode, //todo: PII?
					&a.OriginalTrace,
					&a.OriginalDFI,
					&a.CorrectedData,
					&a.TraceNumber,
				)
			}

			m.maskAddenda99(e.Addenda99)
		}

		// ADVEntries
		for _, e := range b.GetADVEntries() {
			m.maskStrings(
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
			m.maskInts(
				&e.Amount,
				&e.JulianDay,
			)

			m.maskAddenda99(e.Addenda99)
		}

		m.maskBatchControl(b.GetControl())

		if control := b.GetADVControl(); control != nil {
			m.maskStrings(
				&control.ID,
				&control.ODFIIdentification,
			)
			m.maskInts(
				&control.EntryHash,
				&control.TotalDebitEntryDollarAmount,
				&control.TotalCreditEntryDollarAmount,
				&control.BatchNumber,
			)
		}
	}
}

func (m *Masker) maskIATBatches(batches []ach.IATBatch) {
	for _, b := range batches {
		m.maskStrings(&b.ID)

		if h := b.Header; h != nil {
			m.maskStrings(
				&h.ID,
				&h.OriginatorIdentification,
				//&h.StandardEntryClassCode, // todo: PII?
				&h.CompanyEntryDescription,
				&h.EffectiveEntryDate,
				&h.ODFIIdentification,
			)

			m.maskInts(
				// todo: add additional PII fields here
				&h.BatchNumber,
			)
		}

		m.maskBatchControl(b.GetControl())

		for _, e := range b.GetEntries() {
			m.maskStrings(
				&e.ID,
				&e.RDFIIdentification,
				&e.CheckDigit,
				&e.DFIAccountNumber,
				&e.TraceNumber,
				&e.Category,
			)
			m.maskInts(
				&e.Amount,
				&e.TransactionCode,
			)

			if a := e.Addenda10; a != nil {
				m.maskStrings(
					&a.ID,
					//&a.TransactionTypeCode, // todo: PII?
					&a.ForeignTraceNumber,
					&a.Name,
				)
				m.maskInts(
					&a.ForeignPaymentAmount,
					&a.EntryDetailSequenceNumber,
				)
			}

			if a := e.Addenda11; a != nil {
				m.maskStrings(
					&a.ID,
					&a.OriginatorName,
					&a.OriginatorStreetAddress,
				)
				m.maskInts(&a.EntryDetailSequenceNumber)
			}

			if a := e.Addenda12; a != nil {
				m.maskStrings(
					&a.ID,
					&a.OriginatorCityStateProvince,
					&a.OriginatorCountryPostalCode,
				)
				m.maskInts(&a.EntryDetailSequenceNumber)
			}

			if a := e.Addenda13; a != nil {
				m.maskStrings(
					&a.ID,
					&a.ODFIName,
					&a.ODFIIDNumberQualifier,
					&a.ODFIIdentification,
					&a.ODFIBranchCountryCode,
				)
				m.maskInts(&a.EntryDetailSequenceNumber)
			}

			if a := e.Addenda14; a != nil {
				m.maskStrings(
					&a.ID,
					&a.RDFIName,
					&a.RDFIIDNumberQualifier,
					&a.RDFIIdentification,
					&a.RDFIBranchCountryCode,
				)
				m.maskInts(&a.EntryDetailSequenceNumber)
			}

			if a := e.Addenda15; a != nil {
				m.maskStrings(
					&a.ID,
					&a.ReceiverIDNumber,
					&a.ReceiverStreetAddress,
				)
				m.maskInts(&a.EntryDetailSequenceNumber)
			}

			if a := e.Addenda16; a != nil {
				m.maskStrings(
					&a.ID,
					&a.ReceiverCityStateProvince,
					&a.ReceiverCountryPostalCode,
				)
				m.maskInts(&a.EntryDetailSequenceNumber)
			}

			for _, a := range e.Addenda17 {
				m.maskStrings(
					&a.ID,
					&a.PaymentRelatedInformation,
				)
				m.maskInts(&a.EntryDetailSequenceNumber)
			}

			for _, a := range e.Addenda18 {
				m.maskStrings(
					&a.ID,
					&a.ForeignCorrespondentBankName,
					//&a.ForeignCorrespondentBankIDNumberQualifier,
					&a.ForeignCorrespondentBankIDNumber,
					&a.ForeignCorrespondentBankBranchCountryCode,
				)
				m.maskInts(&a.EntryDetailSequenceNumber)
			}

			if a := e.Addenda98; a != nil {
				m.maskStrings(
					&a.ID,
					&a.ChangeCode,
					&a.OriginalTrace,
					&a.OriginalDFI,
					&a.CorrectedData,

					&a.TraceNumber,
				)
			}

			if a := e.Addenda99; a != nil {
				m.maskStrings(
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

func (m *Masker) maskBatchControl(c *ach.BatchControl) {
	if c == nil {
		return
	}

	m.maskStrings(
		&c.ID,
		&c.CompanyIdentification,
		&c.MessageAuthenticationCode,
		&c.ODFIIdentification,
	)
	m.maskInts(
		&c.EntryHash,
		&c.TotalDebitEntryDollarAmount,
		&c.TotalCreditEntryDollarAmount,
		&c.BatchNumber,
	)
}

func (m *Masker) maskAddenda99(a *ach.Addenda99) {
	if a == nil {
		return
	}

	m.maskStrings(
		&a.ID,
		//&a.ReturnCode, //todo: PII?
		&a.OriginalTrace,
		&a.DateOfDeath,
		&a.AddendaInformation,
		&a.TraceNumber,
	)
}
