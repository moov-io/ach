// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "testing"

// mockBatchTRCHeader creates a BatchTRC BatchHeader
func mockBatchTRCHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = "TRC"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "TRC"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockTRCEntryDetail creates a BatchTRC EntryDetail
func mockTRCEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetCheckSerialNumber("123456789")
	entry.SetProcessControlField("CHECK1")
	entry.SetItemResearchNumber("182726")
	entry.SetItemTypeIndicator("01")
	entry.SetTraceNumber(mockBatchTRCHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchTRC creates a BatchTRC
func mockBatchTRC() *BatchTRC {
	mockBatch := NewBatchTRC(mockBatchTRCHeader())
	mockBatch.AddEntry(mockTRCEntryDetail())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// mockBatchTRCHeaderCredit creates a BatchTRC BatchHeader
func mockBatchTRCHeaderCredit() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = "TRC"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "TRC"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockTRCEntryDetailCredit creates a TRC EntryDetail with a credit entry
func mockTRCEntryDetailCredit() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetCheckSerialNumber("123456789")
	entry.SetProcessControlField("CHECK1")
	entry.SetItemResearchNumber("182726")
	entry.SetTraceNumber(mockBatchTRCHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchTRCCredit creates a BatchTRC with a Credit entry
func mockBatchTRCCredit() *BatchTRC {
	mockBatch := NewBatchTRC(mockBatchTRCHeaderCredit())
	mockBatch.AddEntry(mockTRCEntryDetailCredit())
	return mockBatch
}

// testBatchTRCHeader creates a BatchTRC BatchHeader
func testBatchTRCHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchTRCHeader())
	err, ok := batch.(*BatchTRC)
	if !ok {
		t.Errorf("Expecting BatchTRC got %T", err)
	}
}

// TestBatchTRCHeader tests validating BatchTRC BatchHeader
func TestBatchTRCHeader(t *testing.T) {
	testBatchTRCHeader(t)
}

// BenchmarkBatchTRCHeader benchmarks validating BatchTRC BatchHeader
func BenchmarkBatchTRCHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRCHeader(b)
	}
}

// testBatchTRCCreate validates BatchTRC create
func testBatchTRCCreate(t testing.TB) {
	mockBatch := mockBatchTRC()
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRCCreate tests validating BatchTRC create
func TestBatchTRCCreate(t *testing.T) {
	testBatchTRCCreate(t)
}

// BenchmarkBatchTRCCreate benchmarks validating BatchTRC create
func BenchmarkBatchTRCCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRCCreate(b)
	}
}

// testBatchTRCStandardEntryClassCode validates BatchTRC create for an invalid StandardEntryClassCode
func testBatchTRCStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchTRC()
	mockBatch.Header.StandardEntryClassCode = "WEB"
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "StandardEntryClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchTRCStandardEntryClassCode tests validating BatchTRC create for an invalid StandardEntryClassCode
func TestBatchTRCStandardEntryClassCode(t *testing.T) {
	testBatchTRCStandardEntryClassCode(t)
}

// BenchmarkBatchTRCStandardEntryClassCode benchmarks validating BatchTRC create for an invalid StandardEntryClassCode
func BenchmarkBatchTRCStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRCStandardEntryClassCode(b)
	}
}

// testBatchTRCServiceClassCodeEquality validates service class code equality
func testBatchTRCServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchTRC()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchTRCServiceClassCodeEquality tests validating service class code equality
func TestBatchTRCServiceClassCodeEquality(t *testing.T) {
	testBatchTRCServiceClassCodeEquality(t)
}

// BenchmarkBatchTRCServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchTRCServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRCServiceClassCodeEquality(b)
	}
}

// testBatchTRCMixedCreditsAndDebits validates BatchTRC create for an invalid MixedCreditsAndDebits
func testBatchTRCMixedCreditsAndDebits(t testing.TB) {
	mockBatch := mockBatchTRC()
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchTRCMixedCreditsAndDebits tests validating BatchTRC create for an invalid MixedCreditsAndDebits
func TestBatchTRCMixedCreditsAndDebits(t *testing.T) {
	testBatchTRCMixedCreditsAndDebits(t)
}

// BenchmarkBatchTRCMixedCreditsAndDebits benchmarks validating BatchTRC create for an invalid MixedCreditsAndDebits
func BenchmarkBatchTRCMixedCreditsAndDebits(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRCMixedCreditsAndDebits(b)
	}
}

// testBatchTRCCreditsOnly validates BatchTRC create for an invalid CreditsOnly
func testBatchTRCCreditsOnly(t testing.TB) {
	mockBatch := mockBatchTRC()
	mockBatch.Header.ServiceClassCode = CreditsOnly
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchTRCCreditsOnly tests validating BatchTRC create for an invalid CreditsOnly
func TestBatchTRCCreditsOnly(t *testing.T) {
	testBatchTRCCreditsOnly(t)
}

// BenchmarkBatchTRCCreditsOnly benchmarks validating BatchTRC create for an invalid CreditsOnly
func BenchmarkBatchTRCCreditsOnly(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRCCreditsOnly(b)
	}
}

// testBatchTRCAutomatedAccountingAdvices validates BatchTRC create for an invalid AutomatedAccountingAdvices
func testBatchTRCAutomatedAccountingAdvices(t testing.TB) {
	mockBatch := mockBatchTRC()
	mockBatch.Header.ServiceClassCode = AutomatedAccountingAdvices
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchTRCAutomatedAccountingAdvices tests validating BatchTRC create for an invalid AutomatedAccountingAdvices
func TestBatchTRCAutomatedAccountingAdvices(t *testing.T) {
	testBatchTRCAutomatedAccountingAdvices(t)
}

// BenchmarkBatchTRCAutomatedAccountingAdvices benchmarks validating BatchTRC create for an invalid AutomatedAccountingAdvices
func BenchmarkBatchTRCAutomatedAccountingAdvices(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRCAutomatedAccountingAdvices(b)
	}
}

// testBatchTRCCheckSerialNumber validates BatchTRC CheckSerialNumber is not mandatory
func testBatchTRCCheckSerialNumber(t testing.TB) {
	mockBatch := mockBatchTRC()
	// modify CheckSerialNumber / IdentificationNumber to nothing
	mockBatch.GetEntries()[0].SetCheckSerialNumber("")
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "CheckSerialNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchTRCCheckSerialNumber  tests validating BatchTRC
// CheckSerialNumber / IdentificationNumber is a mandatory field
func TestBatchTRCCheckSerialNumber(t *testing.T) {
	testBatchTRCCheckSerialNumber(t)
}

// BenchmarkBatchTRCCheckSerialNumber benchmarks validating BatchTRC
// CheckSerialNumber / IdentificationNumber is a mandatory field
func BenchmarkBatchTRCCheckSerialNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRCCheckSerialNumber(b)
	}
}

// testBatchTRCTransactionCode validates BatchTRC TransactionCode is not a credit
func testBatchTRCTransactionCode(t testing.TB) {
	mockBatch := mockBatchTRCCredit()
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TransactionCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchTRCTransactionCode tests validating BatchTRC TransactionCode is not a credit
func TestBatchTRCTransactionCode(t *testing.T) {
	testBatchTRCTransactionCode(t)
}

// BenchmarkBatchTRCTransactionCode benchmarks validating BatchTRC TransactionCode is not a credit
func BenchmarkBatchTRCTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRCTransactionCode(b)
	}
}

// testBatchTRCAddendaCount validates BatchTRC Addenda count
func testBatchTRCAddendaCount(t testing.TB) {
	mockBatch := mockBatchTRC()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda05" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchTRCAddendaCount tests validating BatchTRC Addenda count
func TestBatchTRCAddendaCount(t *testing.T) {
	testBatchTRCAddendaCount(t)
}

// BenchmarkBatchTRCAddendaCount benchmarks validating BatchTRC Addenda count
func BenchmarkBatchTRCAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRCAddendaCount(b)
	}
}

// testBatchTRCInvalidBuild validates an invalid batch build
func testBatchTRCInvalidBuild(t testing.TB) {
	mockBatch := mockBatchTRC()
	mockBatch.GetHeader().recordType = "3"
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchTRCInvalidBuild tests validating an invalid batch build
func TestBatchTRCInvalidBuild(t *testing.T) {
	testBatchTRCInvalidBuild(t)
}

// BenchmarkBatchTRCInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchTRCInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRCInvalidBuild(b)
	}
}

// TestBatchTRCAddendum98 validates Addenda98 returns an error
func TestBatchTRCAddendum98(t *testing.T) {
	mockBatch := NewBatchTRC(mockBatchTRCHeader())
	mockBatch.AddEntry(mockTRCEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchTRCAddendum99 validates Addenda99 returns an error
func TestBatchTRCAddendum99(t *testing.T) {
	mockBatch := NewBatchTRC(mockBatchTRCHeader())
	mockBatch.AddEntry(mockTRCEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchTRCAddendum99Category validates Addenda99 returns an error
func TestBatchTRCAddendum99Category(t *testing.T) {
	mockBatch := NewBatchTRC(mockBatchTRCHeader())
	mockBatch.AddEntry(mockTRCEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].Category = CategoryForward
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda99" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchTRCProcessControlField returns an error if ProcessControlField is not defined.
func TestBatchTRCProcessControlField(t *testing.T) {
	mockBatch := NewBatchTRC(mockBatchTRCHeader())
	mockBatch.AddEntry(mockTRCEntryDetail())
	mockBatch.GetEntries()[0].SetProcessControlField("")
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ProcessControlField" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchTRCItemResearchNumber returns an error if ItemResearchNumber is not defined.
func TestBatchTRCItemResearchNumber(t *testing.T) {
	mockBatch := NewBatchTRC(mockBatchTRCHeader())
	mockBatch.AddEntry(mockTRCEntryDetail())
	mockBatch.GetEntries()[0].IndividualName = ""
	mockBatch.GetEntries()[0].SetProcessControlField("CHECK1")
	mockBatch.GetEntries()[0].SetItemResearchNumber("")
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ItemResearchNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchTRCItemTypeIndicator returns an error if ItemTypeIndicator is 01.
func TestBatchTRCItemTypeIndicator(t *testing.T) {
	mockBatch := NewBatchTRC(mockBatchTRCHeader())
	mockBatch.AddEntry(mockTRCEntryDetail())
	mockBatch.Create()

	if mockBatch.GetEntries()[0].ItemTypeIndicator() != "01" {
		t.Error("ItemTypeIndicator does not validate")
	}
}
