# ach: validate Addenda05 PRI length in BatchPPD.Validate

## Problem
`BatchPPD.Validate()` did not surface invalid `Addenda05` records.  
Specifically, an `Addenda05` with a `PaymentRelatedInformation` (PRI) field longer than the NACHA limit of **80 characters** could slip through validation. This risked producing or accepting malformed ACH files.

## Fix
- In `BatchPPD.Validate()`:
  - Enforce `len(Addenda05.PaymentRelatedInformation) <= 80`.
  - Call `a.Validate()` for each attached `Addenda05` so any other format errors are surfaced.

## Impact
- Stricter correctness for PPD batches carrying `Addenda05`.
- No public API changes.
- Minimal runtime overhead.

---

## Reproduction (before this patch)

### Minimal test case

```go
package ach

import (
	"strings"
	"testing"
)

func TestRepro_Addenda05TooLong_NotSurfacedBeforeFix(t *testing.T) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 200
	bh.StandardEntryClassCode = "PPD"
	bh.CompanyName = "COMPANY"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "PAYROLL"
	bh.ODFIIdentification = "12104288"

	batch := NewBatchPPD(bh)

	ed := NewEntryDetail()
	ed.TransactionCode = 22
	ed.RDFIIdentification = "23138010"
	ed.CheckDigit = "4"
	ed.DFIAccountNumber = "123456789"
	ed.Amount = 100
	ed.IdentificationNumber = "ID0001"
	ed.IndividualName = "Jane Doe"
	ed.TraceNumber = "121042880000001"

	add := NewAddenda05()
	// Invalid: >80 chars
	add.PaymentRelatedInformation = strings.Repeat("X", 81)

	ed.AddAddenda05(add)
	ed.AddendaRecordIndicator = 1
	batch.AddEntry(ed)

	// BEFORE FIX: Validate() would return nil
	_ = batch.Create()
	if err := batch.Validate(); err == nil {
		t.Fatalf("Before fix: expected error for PRI >80, got nil")
	}
}
```

### CLI
```bash
go test -run TestRepro_Addenda05TooLong_NotSurfacedBeforeFix -v .
```

**Before patch:**  
- `batch.Validate()` returned nil even with invalid PRI.

---

## What this PR changes

- Preserves existing checks (`validateOpts`, `verify()`, SEC code, `InvalidEntries()`).
- Iterates over entries’ `Addenda05`:
  - Returns error if PRI length > 80.
  - Calls `a.Validate()` for additional format checks.

---

## Tests

**New test:** `TestBatchPPD_Addenda05_InvalidFormat_IsCaughtByBatchValidate`  
- Creates a PPD batch with an `Addenda05` whose PRI is 81 chars.  
- Passes if the error is caught in either `batch.Create()` or `batch.Validate()`.  

### Run locally
```bash
go test -run TestBatchPPD_Addenda05_InvalidFormat_IsCaughtByBatchValidate -v .
```

---

## Edge Cases
- Entries with `AddendaRecordIndicator=1` but no `Addenda05` → skipped safely.
- Multiple addenda per entry → each validated.
- Other addenda types (98/99, 02, 10, etc.) → unaffected.
- Backward-compatible: only rejects invalid data.

---

## Performance
- Linear in the number of attached addenda (typically small).
- Negligible overhead.

---

## Documentation
- Enforces NACHA’s 80-char PRI limit.  
- Internal validation is now aligned with spec; no user-facing docs changed.

---

## Manual Verification
- ✅ New unit test fails on pre-patch code, passes with patch.
- ✅ `go test ./...` passes on Linux/macOS.
- ✅ No API changes.
- ✅ Error message is clear and specific.
