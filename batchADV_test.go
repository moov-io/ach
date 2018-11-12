// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
)

// mockBatchADVHeader creates a ADV batch header
func mockBatchADVHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 280
	bh.StandardEntryClassCode = "ADV"
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "Accounting"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockBatchADV creates a ADV batch
func mockBatchADV() *BatchADV {
	mockBatch := NewBatchADV(mockBatchADVHeader())
	mockBatch.AddADVEntry(mockADVEntryDetail())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// testBatchADVHeader creates a ADV batch header
func testBatchADVHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchADVHeader())
	_, ok := batch.(*BatchADV)
	if !ok {
		t.Error("Expecting BatchADV")
	}
}

// TestBatchADVHeader tests creating a ADV batch header
func TestBatchADVHeader(t *testing.T) {
	testBatchADVHeader(t)
}

// BenchmarkBatchADVHeader benchmark creating a ADV batch header
func BenchmarkBatchADVHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchADVHeader(b)
	}
}

// TestBatchADVAddendum98 validates Addenda98 returns an error
func TestBatchADVAddendum98(t *testing.T) {
	mockBatch := NewBatchADV(mockBatchADVHeader())
	mockBatch.AddADVEntry(mockADVEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetADVEntries()[0].Category = CategoryNOC
	mockBatch.GetADVEntries()[0].Addenda98 = mockAddenda98
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

// TestBatchADVAddendum99 validates Addenda99 returns an error
func TestBatchADVAddendum99(t *testing.T) {
	mockBatch := NewBatchADV(mockBatchADVHeader())
	mockBatch.AddADVEntry(mockADVEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetADVEntries()[0].Category = CategoryReturn
	mockBatch.GetADVEntries()[0].Addenda99 = mockAddenda99
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

// testBatchADVSEC validates that the standard entry class code is ADV for batchADV
func testBatchADVSEC(t testing.TB) {
	mockBatch := mockBatchADV()
	mockBatch.Header.StandardEntryClassCode = "RCK"
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

// TestBatchADVSEC tests validating that the standard entry class code is ADV for batchADV
func TestBatchADVSEC(t *testing.T) {
	testBatchADVSEC(t)
}

// BenchmarkBatchADVSEC benchmarks validating that the standard entry class code is ADV for batch ADV
func BenchmarkBatchADVSEC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchADVSEC(b)
	}
}

// testBatchADVServiceClassCode validates ServiceClassCode
func testBatchADVServiceClassCode(t testing.TB) {
	mockBatch := mockBatchADV()
	// Batch Header information is required to Create a batch.
	mockBatch.GetHeader().ServiceClassCode = 0
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

// TestBatchADVServiceClassCode tests validating ServiceClassCode
func TestBatchADVServiceClassCode(t *testing.T) {
	testBatchADVServiceClassCode(t)
}

// BenchmarkBatchADVServiceClassCode benchmarks validating ServiceClassCode
func BenchmarkBatchADVServiceClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchADVServiceClassCode(b)
	}
}

// TestBatchADVAddendum99Category validates Addenda99 returns an error
func TestBatchADVAddendum99Category(t *testing.T) {
	mockBatch := NewBatchADV(mockBatchADVHeader())
	mockBatch.AddADVEntry(mockADVEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockBatch.GetADVEntries()[0].Category = CategoryForward
	mockBatch.GetADVEntries()[0].Addenda99 = mockAddenda99
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
