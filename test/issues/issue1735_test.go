package issues

import (
	"testing"

	"github.com/moov-io/ach"
)

func TestIssue1735_ValidateOptsNotPropagated(t *testing.T) {
	// Test that ValidateOpts are propagated when creating a file from JSON.
	jsonData := []byte(`{
		"fileHeader": {
			"immediateOrigin": "022083649",
			"immediateDestination": "011000015",
			"fileCreationDate": "260120",
			"fileCreationTime": "1030",
			"fileIDModifier": "A"
		},
		"batches": [{
			"batchHeader": {
				"serviceClassCode": 220,
				"companyName": "Test Company ©",
				"companyIdentification": "54321",
				"standardEntryClassCode": "WEB",
				"companyEntryDescription": "Payment",
				"effectiveEntryDate": "260127",
				"originatorStatusCode": 1,
				"ODFIIdentification": "123456780",
				"batchNumber": 1
			},
			"entryDetails": [{
				"transactionCode": 21,
				"RDFIIdentification": "98765432",
				"checkDigit": "0",
				"DFIAccountNumber": "665544",
				"amount": 111111,
				"identificationNumber": "Test",
				"individualName": "John Doe",
				"traceNumber": "12345678123456"
			}]
		}]
	}`)

	// Create ValidateOpts with AllowSpecialCharacters
	validateOpts := &ach.ValidateOpts{
		AllowSpecialCharacters: true,
	}

	// This should succeed but fails
	_, err := ach.FileFromJSONWith(jsonData, validateOpts)
	if err != nil {
		t.Fatalf("Expected success, got error: %v", err)
	}
}
