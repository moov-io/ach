package issues

import (
	"os"
	"testing"
	"time"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestIssue1647(t *testing.T) {
	// Create a new ACH file
	file := ach.NewFile()
	file.Header = ach.NewFileHeader()
	file.Header.ImmediateDestination = "123456780" // Routing number
	file.Header.ImmediateOrigin = "987654320"      // Routing number
	file.Header.FileCreationDate = time.Now().Format("060102")
	file.Header.FileCreationTime = time.Now().Format("1504")
	file.Header.FileIDModifier = "A"

	// Create a PPD batch
	bh := &ach.BatchHeader{
		ServiceClassCode:        220, // Credit
		CompanyName:             "Test Company",
		CompanyIdentification:   "companyID",
		StandardEntryClassCode:  "PPD",
		CompanyEntryDescription: "PAYROLL",
		EffectiveEntryDate:      time.Now().Format("060102"),
		ODFIIdentification:      "12345678",
		OriginatorStatusCode:    1,
	}
	batch, err := ach.NewBatch(bh)
	require.NoError(t, err)

	// Create an Entry Detail Record with addendaRecordIndicator set to 1
	entry := ach.NewEntryDetail()
	entry.TransactionCode = ach.CheckingCredit // 22
	entry.RDFIIdentification = "12345678"
	entry.CheckDigit = "0"
	entry.DFIAccountNumber = "12345678901234567"
	entry.Amount = 10000 // $100.00
	entry.IndividualName = "John Doe"
	entry.AddendaRecordIndicator = 1 // Indicates an addenda record should follow, but we won't add one
	entry.SetTraceNumber(bh.ODFIIdentification, 2)

	err = entry.Validate()
	require.ErrorContains(t, err, "AddendaRecordIndicator 1 does not match addendas found")

	// Add the entry to the batch
	batch.AddEntry(entry)
	err = batch.Create()
	require.ErrorContains(t, err, "AddendaRecordIndicator 1 does not match addendas found")

	// Add the batch to the file
	file.AddBatch(batch)
	err = file.Create()
	require.NoError(t, err)

	// Validate the file
	err = file.Validate()
	require.ErrorContains(t, err, "AddendaRecordIndicator 1 does not match addendas found")

	// For completeness, test the output ACH file format
	err = ach.NewWriter(os.Stdout).Write(file)
	require.ErrorContains(t, err, "AddendaRecordIndicator 1 does not match addendas found")
}
