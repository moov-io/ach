---
layout: page
title: ACH file setup
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Create a file

Creating an Automated Clearing House (ACH) file can be done several ways:

- [Using Go and our generated client](#go-client)
- [Uploading a JSON representation](#upload-a-json-representation)
- [Uploading a raw ACH file](#upload-a-json-representation)

## Go client

We have an example of [using our Go client and uploading the JSON representation](https://github.com/moov-io/ach/blob/master/examples/http/main.go). The basic idea follows this structure:

1. Create a [BatchHeader](https://godoc.org/github.com/moov-io/ach#BatchHeader) record with `ach.NewBatchHeader()`.
1. Create an [EntryDetail](https://godoc.org/github.com/moov-io/ach#EntryDetail) record with `ach.NewEntryDetail()`.
1. Create a [Batch](https://godoc.org/github.com/moov-io/ach#Batch) from our `BatchHeader` and `EntryDetail`.
   1. Using a constructor like `batch := ach.NewBatchPPD(batchHeader)` and adding the batch with `batch.AddEntry(entry)`.
   1. Call and verify `batch.Create()` returns no error.
1. Create our ACH File record `file := ach.NewFile()` and [FileHeader](https://godoc.org/github.com/moov-io/ach#FileHeader) with `ach.NewFileHeader()`
1. Add the `FileHeader` (via `file.SetHeader(fileHeader)`) and `Batch` records to the file (via `file.AddBatch(batch)`).
   1. Call and verify `file.Create()` returns no error.
1. Encode the `File` to JSON (via `json.NewEncoder(&buf).Encode(&file)`) for a `net/http` request.

## Upload a JSON representation

In Ruby we have an example of [creating an ACH file from JSON](https://github.com/moov-io/ruby-ach-demo/blob/master/main.rb). The JSON structure corresponds to our [API endpoint for creating files](https://api.moov.io/#operation/createFile) that the ACH HTTP server expects.

We have [example ACH files](https://github.com/moov-io/ach/blob/master/test/testdata/ppd-valid.json) in JSON.

Note: The header `Content-Type: application/json` must be set to parse the file as JSON, otherwise Nacha's format will be assumed.

### Validate options

When creating a file the server supports query parameters for setting `ValidateOpts` values.

Example: `POST /files/create?requireABAOrigin=true&bypassDestination=true`

| Query Param                        | Validation Option                  |
|------------------------------------|------------------------------------|
| `allowMissingFileControl`          | `AllowMissingFileControl`          |
| `allowMissingFileHeader`           | `AllowMissingFileHeader`           |
| `allowZeroBatches`                 | `AllowZeroBatches`                 |
| `bypassCompanyIdentificationMatch` | `BypassCompanyIdentificationMatch` |
| `bypassDestination`                | `BypassDestinationValidation`      |
| `bypassOrigin`                     | `BypassOriginValidation`           |
| `customReturnCodes`                | `CustomReturnCodes`                |
| `customTraceNumbers`               | `CustomTraceNumbers`               |
| `requireABAOrigin`                 | `RequireABAOrigin`                 |
| `unequalServiceClassCode`          | `UnequalServiceClassCode`          |
| `unorderedBatchNumbers`            | `AllowUnorderedBatchNumbers`       |

## Upload a raw ACH file

Our ACH HTTP server also handles [uploading raw ACH files](https://api.moov.io/#operation/createFile) which is the NACHA text format.  We have example files in their NACHA format and example code for creating and reading the files.

| SEC Code | Description                                  | Example ACH File                                 | Read                | Write                                            |
|----------|---------------------------------------|------------------------------------------|-----------------------------------|------------------------------------|
| ACK      | Acknowledgment Entry for CCD          | [Credit](https://github.com/moov-io/ach/blob/master/test/ach-ack-read/ack-read.ach) | [ACK Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-AckRead) | [ACK Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-AckWrite) |
| ADV      | Automated Accounting Advice           | [Prenote Debit](https://github.com/moov-io/ach/blob/master/test/ach-adv-read/adv-read.ach) | [ADV Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-AdvRead) | [ADV Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-AdvWrite) |
| ARC      | Accounts Receivable Entry             | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-arc-read/arc-debit.ach) | [ARC Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-ArcReadDebit) | [ARC Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-ArcWriteDebit) |
| ATX      | Acknowledgment Entry for CTX          | [Credit](https://github.com/moov-io/ach/blob/master/test/ach-atx-read/atx-read.ach)  | [ATX Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-AtxRead) | [ATX Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-AtxWrite) |
| BOC      | Back Office Conversion                | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-boc-read/boc-debit.ach) | [BOC Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-BocReadDebit) | [BOC Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-BocWriteDebit) |
| CCD      | Corporate credit or debit             | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-ccd-read/ccd-debit.ach) | [CCD Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CcdReadDebit) | [CCD Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CcdWriteDebit) |
| CIE      | Customer-Initiated Entry              | [Credit](https://github.com/moov-io/ach/blob/master/test/ach-cie-read/cie-credit.ach) | [CIE Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CieRead) | [CIE Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CieWrite) |
| COR      | Automated Notification of Change(NOC) | [NOC](https://github.com/moov-io/ach/blob/master/test/ach-cor-read/cor-read.ach)   | [COR Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CorReadCredit) | [COR Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CorWriteCredit) |
| CTX      | Corporate Trade Exchange              | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-ctx-read/ctx-debit.ach) | [CTX Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CtxReadDebit) | [CTX Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CtxWriteDebit) |
| DNE      | Death Notification Entry              | [DNE](https://github.com/moov-io/ach/blob/master/test/ach-dne-read/dne-read.ach)   | [DNE Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-DneRead) | [DNE Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-DneWrite) |
| ENR      | Automatic Enrollment Entry            | [ENR](https://github.com/moov-io/ach/blob/master/test/ach-enr-read/enr-read.ach)   | [ENR Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-EnrRead) | [ENR Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-EnrWrite) |
| IAT      | International ACH Transactions        | [Credit](https://github.com/moov-io/ach/blob/master/test/ach-iat-read/iat-credit.ach) | [IAT Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-IatReadMixedCreditDebit) | [IAT Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-IatWriteMixedCreditDebit) |
| MTE      | Machine Transfer Entry                | [Credit](https://github.com/moov-io/ach/blob/master/test/ach-mte-read/mte-read.ach)   | [MTE Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-MteReadDebit) | [MTE Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-MteWriteDebit) |
| POP      | Point of Purchase                     | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-pop-read/pop-debit.ach) | [POP Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-PopReadDebit) | [POP Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-PopWriteDebit) |
| POS      | Point of Sale                         | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-pos-read/pos-debit.ach) | [POS Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-PosReadDebit) | [POS Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-PosWriteDebit) |
| PPD      | Prearranged payment and deposits      | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-ppd-read/ppd-debit.ach) [Credit](https://github.com/moov-io/ach/blob/master/test/ach-ppd-read/ppd-credit.ach) | [PPD Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-PpdReadCredit) | [PPD Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-PpdWriteCredit) |
| RCK      | Represented Check Entries             | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-rck-read/rck-debit.ach) | [RCK Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-RckReadDebit) | [RCK Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-RckWriteDebit) |
| SHR      | Shared Network Entry                  | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-shr-read/shr-debit.ach) | [SHR Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-ShrReadDebit) | [SHR Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-ShrWrite) |
| TEL      | Telephone-Initiated Entry             | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-tel-read/tel-debit.ach) | [TEL Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-TelReadDebit) | [TEL Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-TelWriteDebit) |
| TRC      | Truncated Check Entry                 | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-trc-read/trc-debit.ach) | [TRC Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-TrcReadDebit) | [TRC Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-TrcWriteDebit) |
| TRX      | Check Truncation Entries Exchange     | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-trx-read/trx-debit.ach) | [TRX Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-TrxReadDebit) | [TRX Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-TrxWriteDebit) |
| WEB      | Internet-initiated Entries            | [Credit](https://github.com/moov-io/ach/blob/master/test/ach-web-read/web-credit.ach) | [WEB Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-WebReadCredit) | [WEB Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-WebWriteCredit) |
| XCK      | Destroyed Check Entry                 | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-xck-read/xck-debit.ach)  | [XCK Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-XckReadDebit) | [XCK Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-XckWriteDebit) |


Note: The header `Content-Type: text/plain` should be set.
