package addenda05_length_test

import (
	"strings"
	"testing"

	"github.com/moov-io/ach"
)

// helpers to create minimal headers per SEC
func minPPDHeader() *ach.BatchHeader {
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = 200
	bh.StandardEntryClassCode = "PPD"
	bh.CompanyName = "COMPANY"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "PAYROLL"
	bh.ODFIIdentification = "12104288"
	return bh
}

func minCCDHeader() *ach.BatchHeader {
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = 200
	bh.StandardEntryClassCode = "CCD"
	bh.CompanyName = "COMPANY"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "PAYMENT"
	bh.ODFIIdentification = "12104288"
	return bh
}

// helper to build a minimal valid entry
func minimalEntryDetail() *ach.EntryDetail {
	ed := ach.NewEntryDetail()
	ed.TransactionCode = 22
	ed.RDFIIdentification = "23138010"
	ed.CheckDigit = "4"
	ed.DFIAccountNumber = "123456789"
	ed.Amount = 100
	ed.IdentificationNumber = "ID0001"
	ed.IndividualName = "Jane Doe"
	ed.TraceNumber = "121042880000001"
	return ed
}

func Test_Addenda05_Length_AppliesTo_PPD_and_CCD(t *testing.T) {
	secs := []struct {
		name     string
		newBatch func() ach.Batcher
	}{
		{"PPD", func() ach.Batcher { return ach.NewBatchPPD(minPPDHeader()) }},
		{"CCD", func() ach.Batcher { return ach.NewBatchCCD(minCCDHeader()) }},
	}

	cases := []struct {
		testName string
		priLen   int
		wantErr  bool
	}{
		{"ok_79_runes", 79, false},
		{"ok_80_runes", 80, false},
		{"fail_81_runes", 81, true},
	}

	for _, sec := range secs {
		for _, tc := range cases {
			t.Run(sec.name+"_"+tc.testName, func(t *testing.T) {
				batch := sec.newBatch()
				ed := minimalEntryDetail()

				add := ach.NewAddenda05()
				add.PaymentRelatedInformation = strings.Repeat("A", tc.priLen)

				ed.AddAddenda05(add)
				ed.AddendaRecordIndicator = 1
				batch.AddEntry(ed)

				if err := batch.Create(); err == nil {
					err = batch.Validate()
					if tc.wantErr && err == nil {
						t.Fatalf("%s: expected error for PRI len=%d runes, got nil", sec.name, tc.priLen)
					}
					if !tc.wantErr && err != nil {
						t.Fatalf("%s: unexpected error for PRI len=%d runes: %v", sec.name, tc.priLen, err)
					}
				} else {
					if !tc.wantErr {
						t.Fatalf("%s: unexpected error from batch.Create() for PRI len=%d: %v", sec.name, tc.priLen, err)
					}
				}
			})
		}
	}
}
