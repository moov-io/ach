package ach

import "testing"

// TODO make all the mock values cor fields

func mockBatchCORHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "COR"
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "123456789"
	bh.CompanyEntryDescription = "Vndr Pay"
	bh.ODFIIdentification = 6200001
	return bh
}

func mockCOREntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 27
	entry.SetRDFI(9101298)
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 5000000
	entry.IdentificationNumber = "location #23"
	entry.SetReceivingCompany("Best Co. #23")
	entry.TraceNumber = 123456789
	entry.DiscretionaryData = "S"
	return entry
}

// TODO make a addendaNOC for COR batches

func mockBatchCOR() *BatchCOR {
	mockBatch := NewBatchCOR()
	mockBatch.SetHeader(mockBatchCORHeader())
	mockBatch.AddEntry(mockCOREntryDetail())
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

func TestBatchCORSEC(t *testing.T) {
	mockBatch := mockBatchCOR()
	mockBatch.header.StandardEntryClassCode = "RCK"
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
