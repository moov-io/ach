// Copyright 2016 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"os"
	"strings"
	"testing"
)

// TestDecode is a complete file decoding test.
func TestDecode(t *testing.T) {
	f, err := os.Open("./testdata/ppd-debit.ach")
	if err != nil {
		t.Errorf("%s: ", err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()
	if err != nil {
		t.Errorf("Can not ach.read ach file: %v", err)
	}
}

// TestParseFileHeader parses a known File Header Record string.
func TestParseFileHeader(t *testing.T) {
	var line = "101 076401251 0764012510807291511A094101achdestname            companyname                    "
	r := NewReader(strings.NewReader(line))
	r.record = line
	record := r.parseFileHeader()

	if record.recordType != "1" {
		t.Errorf("RecordType Expected 1 got: %v", record.recordType)
	}
	if record.priorityCode != "01" {
		t.Errorf("PriorityCode Expected 01 got: %v", record.priorityCode)
	}
	if record.ImmediateDestination() != " 076401251" {
		t.Errorf("ImmediateDestination Expected ' 076401251' got: %v", record.ImmediateDestination())
	}
	if record.ImmediateOrigin() != " 076401251" {
		t.Errorf("ImmediateOrigin Expected '   076401251' got: %v", record.ImmediateOrigin())
	}

	if record.FileCreationDate() != "080729" {
		t.Errorf("FileCreationDate Expected '080729' got:'%v'", record.FileCreationDate())
	}
	if record.FileCreationTime() != "1511" {
		t.Errorf("FileCreationTime Expected '1511' got:'%v'", record.FileCreationTime())
	}

	if record.FileIDModifier != "A" {
		t.Errorf("FileIDModifier Expected 'A' got:'%v'", record.FileIDModifier)
	}
	if record.recordSize != "094" {
		t.Errorf("RecordSize Expected '094' got:'%v'", record.recordSize)
	}
	if record.blockingFactor != "10" {
		t.Errorf("BlockingFactor Expected '10' got:'%v'", record.blockingFactor)
	}
	if record.formatCode != "1" {
		t.Errorf("FormatCode Expected '1' got:'%v'", record.formatCode)
	}
	if record.ImmediateDestinationName != "achdestname            " {
		t.Errorf("ImmediateDestinationName Expected 'achdestname           ' got:'%v'", record.ImmediateDestinationName)
	}
	if record.ImmidiateOriginName != "companyname            " {
		t.Errorf("ImmidiateOriginName Expected 'companyname          ' got: '%v'", record.ImmidiateOriginName)
	}
	if record.ReferenceCode != "        " {
		t.Errorf("ReferenceCode Expected '        ' got:'%v'", record.ReferenceCode)
	}
}

func TestFileHeaderString(t *testing.T) {
	var line = "101 076401251 0764012510807291511A094101achdestname            companyname                    "
	r := NewReader(strings.NewReader(line))
	r.record = line
	record := r.parseFileHeader()

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestParseBatch Header parses a known Batch Header Record string.
func TestParseBatchHeader(t *testing.T) {
	var line = "5225companyname                         origid    PPDCHECKPAYMT000002080730   1076401250000001"
	r := NewReader(strings.NewReader(line))
	r.record = line
	record := r.parseBatchHeader()

	if record.RecordType != "5" {
		t.Errorf("RecordType Expected '5' got: %v", record.RecordType)
	}
	if record.ServiceClassCode != "225" {
		t.Errorf("ServiceClassCode Expected '225' got: %v", record.ServiceClassCode)
	}
	if record.CompanyName != "companyname     " {
		t.Errorf("CompanyName Expected 'companyname    ' got: '%v'", record.CompanyName)
	}
	if record.CompanyDiscretionaryData != "                    " {
		t.Errorf("CompanyDiscretionaryData Expected '                    ' got: %v", record.CompanyDiscretionaryData)
	}
	if record.CompanyIdentification != "origid    " {
		t.Errorf("CompanyIdentification Expected 'origid    ' got: %v", record.CompanyIdentification)
	}
	if record.StandardEntryClassCode != "PPD" {
		t.Errorf("StandardEntryClassCode Expected 'PPD' got: %v", record.StandardEntryClassCode)
	}
	if record.CompanyEntryDescription != "CHECKPAYMT" {
		t.Errorf("CompanyEntryDescription Expected 'CHECKPAYMT' got: %v", record.CompanyEntryDescription)
	}
	if record.CompanyDescriptiveDate != "000002" {
		t.Errorf("CompanyDescriptiveDate Expected '000002' got: %v", record.CompanyDescriptiveDate)
	}
	if record.EffectiveEntryDate != "080730" {
		t.Errorf("EffectiveEntryDate Expected '080730' got: %v", record.EffectiveEntryDate)
	}
	if record.SettlementDate != "   " {
		t.Errorf("SettlementDate Expected '   ' got: %v", record.SettlementDate)
	}
	if record.OriginatorStatusCode != "1" {
		t.Errorf("OriginatorStatusCode Expected 1 got: %v", record.OriginatorStatusCode)
	}
	if record.OdfiIdentification != "07640125" {
		t.Errorf("OdfiIdentification Expected '07640125' got: %v", record.OdfiIdentification)
	}
	if record.BatchNumber != "0000001" {
		t.Errorf("BatchNumber Expected '0000001' got: %v", record.BatchNumber)
	}
}

// TestParseEntryDetail Header parses a known Entry Detail Record string.
func TestParseEntryDetail(t *testing.T) {
	var line = "62705320001912345            0000010500c-1            Arnold Wade           DD0076401255655291"
	r := NewReader(strings.NewReader(line))
	r.record = line
	record := r.parseEntryDetail()

	if record.RecordType != "6" {
		t.Errorf("RecordType Expected '6' got: %v", record.RecordType)
	}
	if record.TransactionCode != "27" {
		t.Errorf("TransactionCode Expected '27' got: %v", record.TransactionCode)
	}
	if record.RdfiIdentification != "05320001" {
		t.Errorf("RdfiIdentification Expected '05320001' got: %v", record.RdfiIdentification)
	}
	if record.CheckDigit != "9" {
		t.Errorf("CheckDigit Expected '9' got: %v", record.CheckDigit)
	}
	if record.DfiAccountNumber != "12345            " {
		t.Errorf("DfiAccountNumber Expected '12345            ' got: %v", record.DfiAccountNumber)
	}
	if record.Amount != "0000010500" {
		t.Errorf("Amount Expected '0000010500' got: %v", record.Amount)
	}

	if record.IndividualIdentificationNumber != "c-1            " {
		t.Errorf("IndividualIdentificationNumber Expected 'c-1            ' got: %v", record.IndividualIdentificationNumber)
	}
	if record.IndividualName != "Arnold Wade           " {
		t.Errorf("IndividualName Expected 'Arnold Wade           ' got: %v", record.IndividualName)
	}
	if record.DiscretionaryData != "DD" {
		t.Errorf("DiscretionaryData Expected 'DD' got: %v", record.DiscretionaryData)
	}
	if record.AddendaRecordIndicator != "0" {
		t.Errorf("AddendaRecordIndicator Expected '0' got: %v", record.AddendaRecordIndicator)
	}
	if record.TraceNumber != "076401255655291" {
		t.Errorf("TraceNumber Expected '076401255655291' got: %v", record.TraceNumber)
	}
	/*
		if record.Addenda != "5" {
			t.Errorf("Addenda Expected '5' got: %v", record.Addenda)
		}
	*/
}

// TestParseAddendaRecord parses a known Addenda Record string.
func TestParseAddendaRecord(t *testing.T) {
	var line = "710WEB                                        DIEGO MAY                                0000001"

	r := NewReader(strings.NewReader(line))
	r.record = line
	record := r.parseAddendaRecord()

	if record.RecordType != "7" {
		t.Errorf("RecordType Expected '7' got: %v", record.RecordType)
	}
	if record.TypeCode != "10" {
		t.Errorf("TypeCode Expected 10 got: %v", record.TypeCode)
	}
	if record.PaymentRelatedInformation != "WEB                                        DIEGO MAY                            " {
		t.Errorf("PaymentRelatedInformation Expected 'WEB                                        DIEGO MAY                            ' got: %v", record.PaymentRelatedInformation)
	}
	if record.SequenceNumber != "    " {
		t.Errorf("SequenceNumber Expected '    ' got: %v", record.SequenceNumber)
	}
	if record.EntryDetailSequenceNumber != "0000001" {
		t.Errorf("EntryDetailSequenceNumber Expected '0000001' got: %v", record.EntryDetailSequenceNumber)
	}
}

// TestParseBatchControl parses a known Batch ControlRecord string.
func TestParseBatchControl(t *testing.T) {
	var line = "82250000010005320001000000010500000000000000origid                             076401250000001"
	r := NewReader(strings.NewReader(line))
	r.record = line
	record := r.parseBatchControl()

	if record.RecordType != "8" {
		t.Errorf("RecordType Expected '8' got: %v", record.RecordType)
	}
	if record.ServiceClassCode != "225" {
		t.Errorf("ServiceClassCode Expected '225' got: %v", record.ServiceClassCode)
	}
	if record.EntryAddendaCount != "000001" {
		t.Errorf("EntryAddendaCount Expected '000001' got: %v", record.EntryAddendaCount)
	}
	if record.EntryHash != "0005320001" {
		t.Errorf("EntryHash Expected '0005320001' got: %v", record.EntryHash)
	}
	if record.TotalDebitEntryDollarAmount != "000000010500" {
		t.Errorf("TotalDebitEntryDollarAmount Expected '000000010500' got: %v", record.TotalDebitEntryDollarAmount)
	}
	if record.TotalCreditEntryDollarAmount != "000000000000" {
		t.Errorf("TotalCreditEntryDollarAmount Expected '000000000000' got: %v", record.TotalCreditEntryDollarAmount)
	}
	if record.CompanyIdentification != "origid    " {
		t.Errorf("CompanyIdentification Expected 'origid    ' got: %v", record.CompanyIdentification)
	}
	if record.MessageAuthenticationCode != "                   " {
		t.Errorf("MessageAuthenticationCode Expected '                   ' got: %v", record.MessageAuthenticationCode)
	}
	if record.Reserved != "      " {
		t.Errorf("Reserved Expected '      ' got: %v", record.Reserved)
	}
	if record.OdfiIdentification != "07640125" {
		t.Errorf("OdfiIdentification Expected '07640125' got: %v", record.OdfiIdentification)
	}
	if record.BatchNumber != "0000001" {
		t.Errorf("BatchNumber Expected '0000001' got: %v", record.BatchNumber)
	}
}

// TestParseFileControl parses a known File Control Record string.
func TestParseFileControl(t *testing.T) {
	var line = "9000001000001000000010005320001000000010500000000000000                                       "
	r := NewReader(strings.NewReader(line))
	r.record = line
	record := r.parseFileControl()

	if record.RecordType != "9" {
		t.Errorf("RecordType Expected '9' got: %v", record.RecordType)
	}
	if record.BatchCount != "000001" {
		t.Errorf("BatchCount Expected '000001' got: %v", record.BatchCount)
	}
	if record.BlockCount != "000001" {
		t.Errorf("BlockCount Expected '000001' got: %v", record.BlockCount)
	}
	if record.EntryAddendaCount != "00000001" {
		t.Errorf("EntryAddendaCount Expected '00000001' got: %v", record.EntryAddendaCount)
	}
	if record.EntryHash != "0005320001" {
		t.Errorf("EntryHash Expected '0005320001' got: %v", record.EntryHash)
	}
	if record.TotalDebitEntryDollarAmountInFile != "000000010500" {
		t.Errorf("TotalDebitEntryDollarAmountInFile Expected '0005000000010500' got: %v", record.TotalDebitEntryDollarAmountInFile)
	}
	if record.TotalCreditEntryDollarAmountInFile != "000000000000" {
		t.Errorf("TotalCreditEntryDollarAmountInFile Expected '000000000000' got: %v", record.TotalCreditEntryDollarAmountInFile)
	}
	if record.Reserved != "                                       " {
		t.Errorf("Reserved Expected '                                       ' got: %v", record.Reserved)
	}

}

func TestRecordTypeUnknown(t *testing.T) {
	var line = "301 076401251 0764012510807291511A094101achdestname            companyname                    "
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !strings.Contains(err.Error(), ErrRecordType.Error()) {
		t.Errorf("Expected RecordType Error got: %v", err)
	}
}

func TestTwoFileHeaders(t *testing.T) {
	var line = "101 076401251 0764012510807291511A094101achdestname            companyname                    "
	var twoHeaders = line + "\n" + line
	r := NewReader(strings.NewReader(twoHeaders))
	_, err := r.Read()

	if !strings.Contains(err.Error(), ErrFileHeader.Error()) {
		t.Errorf("Expected RecordType Error got: %v", err)
	}

}

// Ghetto test function for when I need to prove something to myself.
/*
func TestLines(t *testing.T) {
	f, err := os.Open("./testdata/ppd-debit.ach")
	if err != nil {
		t.Errorf("%s: ", err)
	}
	r := bufio.NewReader(f)
	i := 0
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		i++

		fmt.Printf("%v = %v \n", i, string(line))
	}
}
*/
