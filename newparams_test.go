package ach

import (
	"testing"
)

func TestFileParam(t *testing.T) {
	f := NewFile(
		FileParam{
			ImmediateDestination:     "081000032",
			ImmediateOrigin:          "123456789",
			ImmediateDestinationName: "Your Bank",
			ImmediateOriginName:      "Your Company Inc",
			ReferenceCode:            "A00000"})
	if err := f.Header.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if f.Header.ImmediateOriginName != "Your Company Inc" {
		t.Errorf("FileParam value was not copied to file.Header")
	}
}

func TestBatchParam(t *testing.T) {
	companyName := "Your Company Inc"
	batch := NewBatchPPD(BatchParam{
		ServiceClassCode:        220,
		CompanyName:             companyName,
		StandardEntryClass:      "PPD",
		CompanyIdentification:   "123456789",
		CompanyEntryDescription: "Trans Description",
		CompanyDescriptiveDate:  "Oct 23",
		ODFIIdentification:      "081000032",
	})

	if err := batch.header.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if batch.header.CompanyName != companyName {
		t.Errorf("BatchParam value was not copied to batch.header.CompanyName")
	}
}

/*
		var batch = new nach.Batch({
    serviceClassCode: '220',
    companyName: 'Your Company Inc',
    standardEntryClassCode: 'WEB',
    companyIdentification: '123456789',
    companyEntryDescription: 'Trans Description',
    companyDescriptiveDate: moment(nach.Utils.computeBusinessDay(8)).format('MMM D'),
    effectiveEntryDate: nach.Utils.computeBusinessDay(8),
    originatingDFI: '081000032'
});
*/
