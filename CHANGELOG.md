## v1.41.1 (Released 2024-08-14)

IMPROVEMENTS

- fix: allow ForeignExchangeReferenceIndicator=0 when ForeignExchangeIndicator=FF
- fix: correctly reverse loan transaction codes, test GL credit/debit

BUILD

- build(deps): bump golang from 1.22-alpine to 1.23-alpine
- build(deps): bump golang.org/x/net from 0.27.0 to 0.28.0
- build(deps): bump golang.org/x/text from 0.16.0 to 0.17.0
- build(deps): bump rexml from 3.3.2 to 3.3.3 in /docs
- build(deps-dev): bump github-pages from 231 to 232 in /docs
- build: convert docker-compose to docker compose

## v1.41.0 (Released 2024-07-29)

IMPROVEMENTS

- feat: add AllowZeroEntryAmount validation option

BUILD

- build(deps): bump ossf/scorecard-action from 2.3.3 to 2.4.0
- build(deps): bump rexml from 3.2.8 to 3.3.2 in /docs

## v1.40.4 (Released 2024-07-18)

IMPROVEMENTS

- docs: fix capitlization error. Changed '9' to '('
- docs: update CONTRIBUTING.md
- fix: expand which latin characters are accepted

BUILD

- build(deps): bump github.com/moov-io/base from 0.49.4 to 0.50.0
- fix(deps): update module golang.org/x/net to v0.27.0

## v1.40.3 (Released 2024-06-14)

IMPROVEMENTS

- feat: allow specifying line ending via API
- fix: return checksum error when ImmediateOrigin doesn't validate

BUILD

- build: run oldstable Go in CI
- build(deps): bump github.com/moov-io/base from 0.49.3 to 0.49.4
- build(deps-dev): bump nokogiri from 1.16.2 to 1.16.5 in /docs
- build(deps): bump rexml from 3.2.6 to 3.2.8 in /docs
- build(deps): bump golang.org/x/net from 0.25.0 to 0.26.0

## v1.39.2 (Released 2024-05-10)

This release of moov-io/ach removes the check verifying Amount is non-zero. This has been removed because returned Prenotes
were being incorrectly marked as invalid.

IMPROVEMENTS

- fix: stop checking if .Amount is valid for returns

BUILD

- fix(deps): update module golang.org/x/net to v0.25.0
- build(deps): bump bulma-clean-theme from 0.13.3 to 0.14.0 in /docs
- build(deps): bump golang.org/x/text from 0.14.0 to 0.15.0

## v1.38.1 (Released 2024-04-29)

IMPROVEMENTS

- fix: make -skip-validation flag not depend on -validate flag being present

## v1.38.0 (Released 2024-04-26)

This release of moov-io/ach standardizes `ValidateOpts` query parameters on the HTTP server. Previously slightly different names
were used, which are still supported and deprecated. The GET/POST `/files/{fileID}/validate` endpoints accept `ValidateOpts` as
query parameters with this release.

IMPROVEMENTS

- feat: standardize validate opts query params, accept on GET/POST validate
- fix: nil check BatchHeader when reading JSON
- server/test: verify allowUnorderedBatchNumbers

## v1.37.5 (Released 2024-04-25)

IMPROVEMENTS

- docs: update required fields for CreateFile schema object in openapi.yaml
- fix: add R62 to list of dishonored return codes

BUILD

- build(deps): bump github.com/aws/aws-lambda-go from 1.46.0 to 1.47.0

## v1.37.4 (Released 2024-04-17)

IMPROVEMENTS

- fix: preserve IAT BatchControl.CompanyIdentification when parsing from json

## v1.37.3 (Released 2024-04-12)

IMPROVEMENTS

- fix: CTX allows SavingsReturnNOCDebit

BUILD

- build(deps): bump golang.org/x/net from 0.22.0 to 0.24.0
- build(deps): bump golang.org/x/sync from 0.6.0 to 0.7.0

## v1.37.2 (Released 2024-03-28)

IMPROVEMENTS

- fix: allow "slashed zero" which some OCR software produces
- test: verify trace numbers are kept through HTTP create

## v1.37.1 (Released 2024-03-27)

IMPROVEMENTS

- perf: improve includesValidCharacters lookups
- perf: isTraceNumberODFI

## v1.37.0 (Released 2024-03-13)

The latest update to moov-io/ach introduces significant performance enhancements, particularly to the `MergeFiles(...)` function. Through rigorous testing, we've ensured that the merging process accurately incorporates all input files. However, should any issues arise, we encourage you to [report them](https://github.com/moov-io/ach/issues/new/choose).

Furthermore, this update includes a new function, `MergeDir(...)`, designed to combine a directory of ACH files efficiently, mirroring the functionality of MergeFiles but with enhanced performance.

Key modifications to the behavior of `MergeFiles(..)` include:

- The management of duplicate trace numbers has been refined. They are now allocated to distinct batches within the same output file, in compliance with Nacha regulations, which mandate separation into different batches.
- The merging process now more accurately aligns with the `Conditions` constraints, significantly improving upon previous versions where the limit was seldom fully utilized.
- ValidateOpts are now aggregated across all input files, ensuring that any non-zero value (for example, `true`) applies uniformly to all batches and entries within the file, streamlining the validation process.

IMPROVEMENTS

- cmd/achcli: read files in Nacha or JSON automatically
- docs: Update openapi.yaml - addenda17/addenda18/format for entryHash
- docs: add links to http endpoints where features overlap
- docs: clarify what isn't supported by MergeFiles
- feat: MergeFiles now combines `ValidateOpts` from all input files
- fix: nil check Addenda records inside FileFromJSON
- fix: recreate batch header to avoid mutations
- merge: delve into subdir with custom FS provided
- merge: option to walk sub directories
- refactor: improve merge performance
- test: add checks for pickOutFile, findOutBatch and lineCount
- test: verify entries that are very close still merge properly
- test: verify more numericField inputs

BUILD

- fix(deps): update module golang.org/x/sync to v0.6.0
- build(deps): bump google.golang.org/protobuf from 1.32.0 to 1.33.0
- build(deps): bump golang.org/x/net from 0.21.0 to 0.22.0

## v1.35.1 (Released 2024-03-04)

IMPROVEMENTS

- fix: further reuse memory buffers
- fix: nil check inside setEntryRecordType
- fix: parse large files without newlines

BUILD

- build(deps): bump github.com/stretchr/testify from 1.8.4 to 1.9.0

## v1.35.0 (Released 2024-02-29)

ADDITIONS

- feat: Add `IATCorrectedDataField() string` on `Addenda98` to return the expanded field

IMPROVEMENTS

- all: reduce string allocations
- fix: allow IAT corrections to not have IAT addenda records
- fix: read additional corrected data on IAT corrections

BUILD

- build(deps): bump github.com/prometheus/client_golang

## v1.34.4 (Released 2024-02-26)

IMPROVEMENTS

- fix: check AllowMissingFileHeader and AllowMissingFileControl in reader

## v1.34.3 (Released 2024-02-26)

IMPROVEMENTS

- fix: Correctly assign Category to IAT entries.

BUILD

- build(deps): bump golang.org/x/net from 0.20.0 to 0.21.0
- chore(deps): update dependency github-pages to v231

## v1.34.2 (Released 2024-02-07)

IMPROVEMENTS

- fix: skip reserved fields in Dishonored Returns

BUILD

- build(deps): bump golang from 1.21 to 1.22
- build: experiment with ARM M1 builders and releases
- build: print crashers after fuzzing
- fix(deps): update module github.com/aws/aws-lambda-go to v1.46.0

## v1.34.0 (Released 2024-01-25)

IMPROVEMENTS

- fix: use constructor to create IATEntryDetails in reader
- test/issues: write up scenario for Issue 1312

BUILD

- build(deps): bump actions/download-artifact from 1 to 4
- build(deps): bump actions/upload-artifact from 1 to 4

## v1.33.6 (Released 2024-01-19)

IMPROVEMENTS

- api: fixup EntryDetail required fields
- api: label required fields on EntryDetail
- feat: Addenda99.SetOriginalEntryReturnDate
- feat: Parse Addenda99.AddendaInformation
- feat: improve addenda99 for dishonored returns
- feat: add CheckingReturnNOCCredit to list of valid ATX txn codes (#21)
- fix BatchHeader.EffectiveEntryDateField (#20)
- fix: Don't truncate Addenda99.AddendaInformation
- fix: trim "R" in dishonored and contested setter
- test: verify "!" is allowed in isAlphanumeric

BUILD

- build(deps): bump actions/download-artifact from 1 to 4
- build(deps): bump actions/upload-artifact from 1 to 4
- build(deps): bump github.com/moov-io/base from 0.48.2 to 0.48.3 (#1341)
- build(deps): bump github.com/aws/aws-lambda-go from 1.41.0 to 1.43.0 (#1346)
- build(deps): bump github/codeql-action from 2 to 3 (#1342)
- build(deps): bump github.com/prometheus/client_golang
- build(deps): bump github.com/aws/aws-lambda-go from 1.43.0 to 1.44.0
- build(deps): bump golang.org/x/net from 0.19.0 to 0.20.0
- build(deps): bump github.com/moov-io/base from 0.48.3 to 0.48.5
- build(deps): bump actions/cache from 3 to 4
- build(deps): bump github.com/aws/aws-lambda-go from 1.44.0 to 1.45.0
- chore(deps): update dependency jekyll-feed to v0.17.0 (#1337)

## v1.33.4 (Released 2023-12-07)

IMPROVEMENTS

- all: fix errors found by nilaway
- api: update amount to be long to allow handling of amounts larger than 21M (#1335)
- webui: better overflow layout

BUILD

- build(deps): bump actions/setup-go from 4 to 5
- build(deps): bump golang.org/x/net from 0.18.0 to 0.19.0

## v1.33.3 (Released 2023-11-14)

IMPROVEMENTS

- docs: improve TraceNumber
- fix: check .UnequalAddendaCounts in CTX validation

BUILD

- build(deps): bump golang.org/x/net from 0.17.0 to 0.18.0
- fix(deps): update module github.com/moov-io/base to v0.48.2

## v1.33.2 (Released 2023-10-31)

IMPROVEMENTS

- fix: Parse FileIdModifier from file header; don't assume 'A'
- fix: revert "reduce allocations when parsing FileHeader"
- test/issues: verify FileIDModifier is read

BUILD

- build(deps): bump github.com/moov-io/base from 0.47.0 to 0.47.1

## v1.33.1 (Released 2023-10-26)

This release contains multiple performance improvements ontop of what v1.33.0 has included. The comparison to v1.32.2 is below:

<details>
<summary>v1.32.2 compared to v1.33.1</summary>

```
BenchmarkParsing/Read_../testdata/20110805A.ach-16         	    1840	    632450 ns/op	   97797 B/op	    1953 allocs/op
BenchmarkParsing/Read_../testdata/20110805A.ach-16         	    1909	    606422 ns/op	   57270 B/op	     922 allocs/op

BenchmarkParsing/IAT-16                                    	    4615	    232996 ns/op	   45724 B/op	     410 allocs/op
BenchmarkParsing/IAT-16                                    	    4666	    230543 ns/op	   31823 B/op	     306 allocs/op

BenchmarkFile/String-16                                    	    5512	    206355 ns/op	   35503 B/op	    1282 allocs/op
BenchmarkFile/String-16                                    	    6181	    188959 ns/op	   30261 B/op	     639 allocs/op
```
</details>

IMPROVEMENTS

- feat: optimize parsing of ADVBatchControl and FileControl
- feat: reduce allocations when parsing records
- fix: allocate returned array in ReadFiles and ReadDir once
- fix: use shared buffer when rendering EntryDetail
- test: verify IAT batches are skipped in Iterator

BUILD

- chore(deps): update ossf/scorecard-action action to v2.3.1

## v1.33.0 (Released 2023-10-19)

Moov ACH v1.33.0 has undergone some major performance tuning and improvements. Please let us know if you encounter any issues when using the improved code, thanks!

ADDITIONS

- feat: add a File iterator
- feat: add benchmark for ReadFile and Iterator
- fix: produce entries from iterator that are outside of batches

IMPROVEMENTS

- feat: check for valid JSON before parsing
- fix: avoid nil panic with malformed readers
- fix: process entry rune-by-rune without casting between string -> rune -> string
- fix: skip directories in ReadDir
- perf: reduce allocations in formatting strings
- perf: reduce allocations when checking routing numbers
- perf: use sync.Pool of bytes.Buffer objects grown to max field size
- test: benchmark merging files in groups

BUILD

- cmd/webui: force wasm_exec.js to go1.21.0 version
- fix(deps): update module golang.org/x/net to v0.17.0
- fix(deps): update module golang.org/x/oauth2 to v0.13.0

## v1.32.2 (Released 2023-08-11)

IMPROVEMENTS

- fix: Corrected Data is 29 characters, not 22
- fix: don't print blank line for missing IAT addendas
- test/issues: match IAT output byte-for-byte

BUILD

- build(deps): bump golang from 1.20-alpine to 1.21-alpine

## v1.32.1 (Released 2023-08-07)

IMPROVEMENTS

- feat: detect Windows-1252 encoding and decode to utf-8
- fix: parse field lengths according by runes instead of bytes

BUILD

- build(deps): bump github.com/moov-io/iso4217 from 0.2.1 to 0.3.0
- fix(deps): update module github.com/moov-io/base to v0.45.1

## v1.32.0 (Released 2023-07-19)

IMPROVEMENTS

- feat: add Addenda98Refused
- fix: SetTraceNumber should populate contested and dishonored returns
- fix: allow full range of ASCII and EBCDIC characters accepted by Nacha rules

BUILD

- build: run Go tests on Windows
- build(deps): bump github.com/moov-io/base from 0.44.0 to 0.45.0
- build(deps): bump golang.org/x/text from 0.10.0 to 0.11.0
- chore(deps): update ossf/scorecard-action action to v2.2.0

## v1.31.5 (Released 2023-06-21)

IMPROVEMENTS

- fix: allow .Create() to skip validation when configured

BUILD

- build(deps): bump actions/checkout from 2 to 3
- chore: update github.com/prometheus/client_golang to v1.16.0
- chore: update golang.org/x/text to v0.10.0

## v1.31.4 (Released 2023-06-12)

IMPROVEMENTS

- fix: return empty strings when records are nil
- build: split "make docker" into its own job
- build: remove docker build and openapi generate from Linux CI
- build: replace fuzzer library with Go standard library
- build: switch to the new moov-io/iso4217 library
- build: fix description of addenda99

BUILD

- fix(deps): update github.com/stretchr/testify to v1.8.4
- build(deps): bump github.com/moov-io/base from 0.42.0 to 0.44.0
- build(deps): bump actions/setup-go from 2 to 4
- fix(deps): update github.com/moov-io/iso4217 to v0.2.1

## v1.31.3 (Released 2023-05-22)

IMPROVEMENTS

- build: fix openshift image, upgrade to ubi9
- cmd/server: fix log format flag logic

BUILD

- build(deps): bump bulma-clean-theme from 0.13.2 to 0.13.3 in /docs
- build(deps): bump github.com/stretchr/testify from 1.8.2 to 1.8.3

## v1.31.2 (Released 2023-05-11)

IMPROVEMENTS

- fix: BatchHeaders are different with mismatching CompanyEntryDescription values

BUILD

- build(deps): bump github.com/moov-io/base from 0.41.0 to 0.42.0
- build(deps): bump github.com/aws/aws-lambda-go from 1.40.0 to 1.41.0
- chore(deps): update dependency jekyll-feed to v0.17.0

## v1.31.1 (Released 2023-04-21)

IMPROVEMENTS

- fix: Allow returned prenotes to have a zero'd Amount field
- fix: make CTX setters/getters return correct results
- fix: validate CTX addenda records and indicator correctly
- test: use a static file header for Contested/Dishonored returns

BUILD

- build(deps): bump golang.org/x/oauth2 from 0.5.0 to 0.7.0

## v1.31.0 (Released 2023-04-18)

IMPROVEMENTS

- cmd/achcli: return whatever partial file was parsed
- feat: add AllowInvalidAmounts validation option
- feat: add Merge which can override ValidateOpts of files
- server: support AllowInvalidAmounts validation option

BUILD

- build(deps): bump commonmarker from 0.23.8 to 0.23.9 in /docs
- build(deps): bump github.com/aws/aws-lambda-go from 1.39.1 to 1.40.0
- build(deps): bump github.com/moov-io/base from 0.40.1 to 0.40.2
- build(deps): bump nokogiri from 1.13.10 to 1.14.3 in /docs
- build(deps): update github.com/prometheus/client_golang to v1.15.0

## v1.30.0 (Released 2023-04-10)

Version v1.30.0 of moov-io/ach introduces a new `ErrFileConsecutiveBatchHeaders` error which replaces `ErrFileBatchHeaderInsideBatch` to better reflect when file parsing failed. The library can now accept ACH files without Batch Control records.

IMPROVEMENTS

- fix: accumulate batches without BatchControl records

## v1.29.4 (Released 2023-04-10)

IMPROVEMENTS

- fix: allow IAT as transaction type code of addenda10 records
- fix: keep zero-prefix trace numbers during offset creation

BUILD

- build(deps): bump github.com/moov-io/base from 0.39.0 to 0.40.1
- build(deps): bump golang.org/x/text from v0.8.0 to v0.9.0
- chore(deps): update dependency jekyll-feed to v0.17.0

## v1.29.3 (Released 2023-03-31)

IMPROVEMENTS

- fix: include Offset in API docs and read from JSON

BUILD

- build(deps): bump activesupport from 6.1.7.2 to 7.0.4.3 in /docs
- build(deps): bump bulma-clean-theme from 0.13.1 to 0.13.2 in /docs
- build(deps): bump github.com/aws/aws-lambda-go from 1.38.0 to 1.39.1
- build(deps): bump github.com/stretchr/testify from 1.8.1 to 1.8.2
- build(deps): bump ossf/scorecard-action from 2.1.2 to 2.1.3

## v1.29.2 (Released 2023-02-21)

Shout-out to @ckbaum for contributing a [major improvement to File Flattening](https://github.com/moov-io/ach/pull/1178).

IMPROVEMENTS

- feat: Don't combine two batches if they share any trace numbers.
- feat: Make an effort to preserve original batch ordering.
- feat: expose named errors (and ask for bug reports) from FlattenFile

BUILD

- build: tighten up GITHUB_TOKEN permissions
- build: Apply security best practices from StepSecurity
- build(deps): bump ossf/scorecard-action from 2.0.6 to 2.1.2
- chore(deps): update dependency tzinfo to "~> 2.0"
- fix(deps): update module github.com/juju/ansiterm to v1

## v1.28.1 (Released 2023-02-13)

IMPROVEMENTS

- chore: remove wasteful debug printf
- fix: validate amounts in non-prenote entries
- test: convert log.Fatal over to t.Fatal
- test: verify ValidAmountForCodes works as expected

BUILD

- build: update golang.org/x/sys/unix
- build(deps): bump activesupport from 6.0.3.7 to 6.0.6.1 in /docs
- chore(deps): update golang docker tag to v1.20
- fix(deps): update module github.com/moov-io/base to v0.39.0
- fix(deps): update module golang.org/x/text to v0.7.0

## v1.28.0 (Released 2023-01-12)

ADDITIONS

- Add `ValidateOpts` to BatchControl.
- Add use of `PreserveSpaces` validate opt in BatchControl, BatchHeader, and EntryDetail.

IMPROVEMENTS

- fix: trim DFIAccountNumber by default

BUILD

- chore: generate iso3166 and iso4217 code, no changes
- fix(deps): update module github.com/aws/aws-lambda-go to v1.37.0
- test: use testify assert package for Valid unit tests  change packages to *_test

## v1.27.0 (Released 2023-01-06)

ADDITIONS

- Add `-pretty.amounts` CLI flag to show human readable amounts
- Add `-skip-validation` add option to skip all validations
- Add `-validate` CLI flag to read files with ValidateOpts
- feat: add option to skip all validations

IMPROVEMENTS

- fix: properly count string lengths when padding fields
- test: verify EntryDetail record lengths are valid for all return codes

BUILD

- build: enable Go test shuffle
- fix(deps): update module golang.org/x/oauth2 to v0.4.0
- fix(deps): update module golang.org/x/text to v0.6.0

## v1.26.4 (Released 2022-12-19)

IMPROVEMENTS

- fix: allow missing TerminalIdentificationCode in Addenda02

## v1.26.3 (Released 2022-12-19)

IMPROVEMENTS

- fix: don't re-check whether the file is ADV while writing each batch

BUILD

- build: update moov-io/base to v0.38.1 and tidy up

## v1.26.2 (Released 2022-12-19)

IMPROVEMENTS

- fix: remove unused x-idempotency-key from api docs

BUILD

- fix(deps): update module github.com/moov-io/base to v0.38.0

## v1.26.1 (Released 2022-12-08)

IMPROVEMENTS

- cmd/achcli/describe: mask left-aligned spaces
- cmd/achcli/describe: fixup padding when fields are padded

## v1.26.0 (Released 2022-12-08)

IMPROVEMENTS

- cmd/achcli/describe: return masked account number length unchanged, show spaces
- cmd/achcli: reunify entrydetail transaction type formatting
- cmd/webui: add a Reversal button
- fix: update fields required for reversed files to validate

BUILD

- build(deps): bump nokogiri from 1.13.9 to 1.13.10 in /docs
- fix(deps): update module github.com/aws/aws-lambda-go to v1.36.0
- fix(deps): update module github.com/moov-io/base to v0.37.0

## v1.25.0 (Released 2022-11-29)

IMPROVEMENTS

- cmd/achcli: align tabwriter columns better across batches

## v1.24.0 (Released 2022-11-21)

ADDITIONS

- feat: add `CalculateCheckDigit`
- feat: add `.SetMaxLines(int)` to Reader
- feat: add `UnequalAddendaCounts` validation option

IMPROVEMENTS

- api: add missing unorderedBatchNumbers and allowInvalidCheckDigit

BUILD

- fix(deps): update module github.com/aws/aws-lambda-go to v1.35.0

## v1.23.1 (Released 2022-11-15)

IMPROVEMENTS

- cmd/achcli: support masking corrected data, option to prettify amounts

BUILD

- fix(deps): update module github.com/moov-io/base to v0.36.2
- fix(deps): update module golang.org/x/oauth2 to v0.2.0

## v1.22.2 (Released 2022-11-08)

IMPROVEMENTS

- docs: fix link to creating file over HTTP
- fix: allow mixed-case Addenda10 TransactionType codes
- validator: include missing secondary SEC codes for Addenda10

BUILD

- fix(deps): update module github.com/prometheus/client_golang to v1.14.0

## v1.22.1 (Released 2022-10-31)

IMPROVEMENTS

- docs: Fix spelling mistake in Original RDFI identification
- server: read ValidateOpts before file to segment

## v1.22.0 (Released 2022-10-25)

IMPROVEMENTS

- docs: rename 'payment code' to transaction code
- examples: add micro-entries / micro-deposits (#1103)
- feat: add `ReadJSONFile` and `ReadJSONFileWith` helpers
- feat: add `Reversal()` to files
- feat: add `AllowInvalidCheckDigit` validation option
- feat: add achcli IAT Support (#1100)
- feat: allow empty strings when allow missing file header opt is set

BUILD

- build(deps): bump nokogiri from 1.13.6 to 1.13.9 in /docs
- fix(deps): update module github.com/moov-io/base to v0.36.1
- fix(deps): update module github.com/stretchr/testify to v1.8.1
- fix(deps): update module golang.org/x/oauth2 to v0.1.0

## v1.21.2 (Released 2022-10-04)

IMPROVEMENTS

- fix: copy forward the validate options when segmenting a file
- fix: passing validate opts to http handler
- test: verify ValidateOpts are set in segment http response

BUILD

- meta: add feature request template

## v1.21.0 (Released 2022-10-03)

IMPROVEMENTS

- docs: clarified lambda_json_to_nacha example
- server: read ValidateOpts in SegmentFile endpoints

BUILD

- build: require go1.19.1 in CI/CD

## v1.20.1 (Released 2022-09-21)

IMPROVEMENTS

- fix: lowercase `ID` in OpenAPI spec

## v1.20.0 (Released 2022-09-19)

ADDITIONS

- feat: add `ValidateOpts` for `IATBatch` and `IATEntryDetail`

IMPROVEMENTS

- fix(deps): update golang.org/x/oauth2 digest to f213421
- fix(deps): update module github.com/moov-io/base to v0.35.0
- test: verify we pad to 10 lines in writer

## v1.19.3 (Released 2022-09-01)

Moov ACH v1.19.3 fixes an issue where Entry Hash values were not properly calculated.
- IATBatch hash values would not overflow properly.
- If `RDFIIdentification` was not properly trimmed to the "ABA 8" characters.

IMPROVEMENTS

- fix: calculate EntryHash from ABA8, properly format string value

## v1.19.2 (Released 2022-08-30)

IMPROVEMENTS

- Allow MixedDebitsAndCredits in ARC, BOC, CIE, POP, RCK, TRC, TRC, XCK entries

BUILD

- fix(deps): update golang.org/x/oauth2 digest to 0ebed06

## v1.19.1 (Released 2022-08-22)

This is the same release as v1.19.0 but with deployment issues fixed.

## v1.19.0 (Released 2022-08-22)

IMPROVEMENTS

- Remove recordType and reserved fields from all models
- docs: update OpenAPI spec with /segment endpoint
- feat: add endpoint to segment files without a fileID
- feat: read fileID path parameter on creation
- feat: return ACH file with creation, flatten, and segment endpoints

## v1.18.6 (Released 2022-08-11)

IMPROVEMENTS

- docs: clarify TEL/WEB payment type code field
- fix: contested return addenda parsing of dishonored returns
- server: verify fileID specified in JSON is kept

BUILD

- build: remove deprecated ioutil functions, gofmt from go1.19
- fix(deps): update golang.org/x/oauth2 digest to 128564f
- fix(deps): update module github.com/aws/aws-lambda-go to v1.34.1

## v1.18.5 (Released 2022-07-19)

IMPROVEMENTS

- docs: update ppd-valid.json ID values
- feat: add R36 return code

## v1.18.4 (Released 2022-07-14)

IMPROVEMENTS

- Fix entry category assignment in Reader

## v1.18.3 (Released 2022-07-13)

IMPROVEMENTS

- feat: add R62 and R77 return codes

BUILD

- fix(deps): update module github.com/moov-io/base to v0.33.0

## v1.18.2 (Released 2022-07-07)

IMPROVEMENTS

- fix: reduce memory usage during MergeFiles

BUILD

- fix(deps): update module github.com/aws/aws-lambda-go to v1.32.1
- fix(deps): update module github.com/moov-io/base to v0.32.0
- fix(deps): update module github.com/stretchr/testify to v1.8.0

## v1.18.1 (Released 2022-06-15)

IMPROVEMENTS

- feat: support Refused COR/NOC codes

BUILD

- fix(deps): update module github.com/moov-io/base to v0.31.1

## v1.18.0 (Released 2022-06-07)

IMPROVEMENTS

- Export the SettlementDate field in BatchHeader and IATBatchHeader

BUILD

- fix(deps): update module github.com/moov-io/base to v0.30.0
- fix(deps): update module github.com/stretchr/testify to v1.7.2

## v1.17.0 (Released 2022-06-02)

ADDITIONS

- feat: add build file endpoint

IMPROVEMENTS

- cmd/webui: update wasm_exec.js bundle
- cmd/webui: write library version to page
- feat: set Addenda trace numbers in (EntryDetail).SetTraceNumber

## v1.16.1 (Released 2022-05-23)

IMPROVEMENTS

- server: read unorderedBatchNumbers query param

## v1.16.0 (Released 2022-05-23)

ADDITIONS

- Add `AllowUnorderedBatchNumbers` validate option to skip ascending batch sequence validation

IMPROVEMENTS

- api: add missing validate opts query params, Addenda99 Dishonored and Contested returns
- batch: skip addendaFieldInclusion with Offset entries

BUILD

- build: update base images
- build: update codeql action
- build(deps): bump nokogiri from 1.13.4 to 1.13.6 in /docs
- fix(deps): update module github.com/aws/aws-lambda-go to v1.32.0
- fix(deps): update module github.com/go-kit/log to v0.2.1
- fix(deps): update module github.com/moov-io/base to v0.29.2
- fix(deps): update module github.com/prometheus/client_golang to v1.12.2

## v1.15.1 (Released 2022-05-09)

IMPROVEMENTS

- feat: improve performance of getTraceNumbers

## v1.15.0 (Released 2022-05-03)

ADDITIONS

- merge: allow splitting files on their maximum dollar amounts

IMPROVEMENTS

- docs: provide specific actions of CustomTraceNumbers
- fix(deps): update module github.com/aws/aws-lambda-go to v1.31.1

## v1.14.0 (Released 2022-03-30)

ADDITIONS

- feat: add support for "Dishonored Returns" with addenda records
- feat: support parsing out Contested Dishonored Returns

BUILD

- build: install git in builder image
- meta: update moov-io/base and our logging library

## v1.13.1 (Released 2022-02-28)

IMPROVEMENTS

- fix: better date validation

BUILD

- build(deps): bump nokogiri from 1.12.5 to 1.13.3 in /docs

## v1.13.0 (Released 2022-01-27)

ADDITIONS

- file: include ValidateOpts with JSON representations

IMPROVEMENTS

- merge: carry through ValidateOpts from files
- server: verify balanced file validates

BUILD

- build: enforce a coverage threshold

## v1.12.2 (Released 2021-09-23)

BUG FIXES

- batch: check ValidateOpts prior to TraceNumber ascending verification
- file: pass through ValidateOpts for inner JSON unmarshals

BUILD

- fix(deps): update module github.com/moov-io/base to v0.24.0

## v1.12.1 (Released 2021-09-10)

BUG FIXES

- batch: check ValidateOpts for bypassing trace number ordering requirement

BUILD

- fix(deps): update module github.com/mattn/go-isatty to v0.0.14

## v1.12.0 (Released 2021-09-02)

ADDITIONS

- file: during creation check ValidateOpts for allowing empty headers
- server: include missing FileHeader and missing FileControl query params

IMPROVEMENTS

- file: allow for missing file header and control

BUG FIXES

- file: unmarshal with ValidateOpts if set

## v1.11.0 (Released 2021-08-11)

ADDITIONS

- batch: allow unequal ServiceClassCodes to be used
- server: allow unequalServiceClassCode query param

IMPROVEMENTS

- addenda99: Use updated [R17 title / description from Nacha update](https://www.nacha.org/rules/reversals-and-enforcement) (#971)

BUILD

- fix(deps): update golang.org/x/oauth2 commit hash to faf39c7
- fix(deps): update module github.com/aws/aws-lambda-go to v1.26.0 (#968)
- fix(deps): update module github.com/moov-io/base to v0.22.0
- meta: fixup from gosec linter

## v1.10.1 (Released 2021-07-26)

BUG FIXES

- cmd/achcli/describe: include EffectiveEntryDate on human readable output (#967)

## v1.10.0 (Released 2021-07-23)

ADDITIONS

- Implementing CustomReturnCode validation (#966)
- Keep settlement date when parsing files. (#960)
- file,server: add method for parsing JSON with custom validation

IMPROVEMENTS

- file: update docs on FileFromJSON validation

BUILD

- build: use debian stable's slim image in webui
- build(deps): bump addressable from 2.7.0 to 2.8.0 in /docs
- fix(deps): update github.com/juju/ansiterm commit hash to 9283cdf (#952)
- fix(deps): update module github.com/aws/aws-lambda-go to v1.25.0
- fix(deps): update module github.com/go-kit/kit to v0.11.0
- fix(deps): update module github.com/moov-io/base to v0.21.1

## v1.9.3 (Released 2021-06-25)

BUG FIXES

- api: fixup openapi spec and generate client in CI

BUILD

- fix(deps): update module github.com/moov-io/base to v0.20.1 (#947)

## v1.9.2 (Released 2021-06-17)

BUG FIXES

- ADV: fix entry hash err msg (#939)
- CTX: allow batches to be pre-notifications

IMPROVEMENTS

- docs: clean up api documentation (#938)

BUILD

- fix(deps): update module github.com/aws/aws-lambda-go to v1.24.0 (#935)
- fix(deps): update module github.com/mattn/go-isatty to v0.0.13 (#937)
- fix(deps): update module github.com/prometheus/client_golang to v1.11.0 (#941)

## v1.9.1 (Released 2021-05-20)

ADDITIONS

- webui: Add human-readable ACH parsing (#933)

BUILD

- build(deps): bump nokogiri from 1.11.3 to 1.11.5 in /docs (#934)

## v1.9.0 (Released 2021-05-13)

ADDITIONS

- reader: add `ReadFiles(...)` method for consuming a slice of filepaths (#932)

BUG FIXES

- set batchID from header on POST to create file (#926)

IMPROVEMENTS

- merge: combine entries together when Batch Headers match
- merge: don't include TraceNumbers that collide in merged files

BUILD

- build(deps): bump kramdown from 2.3.0 to 2.3.1 in /docs
- fix(deps): update module github.com/moov-io/base to v0.19.0 (#929)

## v1.8.1 (Released 2021-05-06)

IMPROVEMENTS

- cmd/achcli: include last error message in -reformat

## v1.8.0 (Released 2021-05-03)

BUG FIXES

- reader: reset the `currentBatch` when a `BatchControl` is read to avoid setting EntryDetail records in the wrong batch
- file: force a 10-digit entry hash in `File.Create` and `File.calculateEntryHash` to ensure the file assembled into a valid state

ADDITIONS

- writer: add `BypassValidation` option
- converters: add method to return the least significant digits of a number
- server: accept `customTraceNumbers` as a query param to set the corresponding validation option

IMPROVEMENTS

- file: update documentation for `Read`, `Create`, and `Validate/ValidateWith` to clarify behavior
- docs: describe validation query params on the HTTP request to create a file
- docs: remove function parentheses from code comments
- tests: add a regression test for issue #915

## v1.7.1 (Released 2021-04-16)

There are no changes from v1.7.0. This release fixes issues during the release process.

## v1.7.0 (Released 2021-04-16)

ADDITIONS

- webui: add button to support JSON to ACH conversions
- achcli: add `describe` package to allow human-readable printing to be reused in other tools
- batch: add validation option, `bypassCompanyIdentificationMatch`, to allow batches in which the `CompanyIdentification` field in the batch header and control do not match

## v1.6.5 (Released 2021-04-15)

IMPROVEMENTS

- merge: optimize performance of `lineCount(..)` and reduce errors returned

## v1.6.4 (Released 2021-04-14)

ADDITIONS

- Add option to create a file with zero batches (#884)

IMPROVEMENTS

- iso3166,iso4217: generate lists again after long hiatus

BUG FIXES

- file: sort EntryDetail records when flattening batches

BUILD

- fix(deps): update module github.com/prometheus/client_golang to v1.10.0 (#875)
- fix(deps): update module github.com/moov-io/base to v0.18.2

## v1.6.3 (Released 2021-03-19)

BUG FIXES

- file: updated `File.MergeFiles(..)` to override batch numbers in each file to prevent them from colliding

IMPROVEMENTS

- file: optimized and refactored `File.FlattenBatches(..)` and added benchmarks
- api: added missing required fields in models in the OpenAPI spec and removed unnecessary `omitempty` tags for required fields in Go models
- docs: removed extra spaces in the request body for creating a file to fix broken example
- docs: added a section about where to find Nacha Operating Rules and Nacha's ACH guide for developers
- docs: updated links in README to point to [project documentation site](https://moov-io.github.io/ach)

## v1.6.2 (Released 2021-02-11)

IMPROVEMENTS

- api: update OpenAPI spec to include more details about data types
- file: ability to set custom batch sequence numbers
- docs: new [project documentation site](https://moov-io.github.io/ach) with updated styles and navigation
- docs: add page about Notification of Change (NOC) files
- docs: add read and write examples for CIE batches
- docs: update R10 and R11 return code descriptions
- docs: additional fixes and improvements to the README

## v1.6.1 (Released 2021-01-25)

BUG FIXES

- fileHeader: `ImmediateDestination` and `ImmediateOrigin` should be separated by a space unless bypass validation is set and the routing number is 10 digits

IMPROVEMENTS

- docs: improvements to README by adding a new header, table of contents, related projects, and FAQ
- docs: clarified code comments related to `File.ValidateWith(..)` and fixed minor documentation errors

## v1.6.0 (Released 2021-01-13)

BUG FIXES

- fileHeader: fix `ImmediateDestination` formatting when `BypassDestinationValidation` is set

ADDITIONS

- batch: allow addenda in CIE batches to be optional
- batch: allow POS batches to contain credits
- reader: allow backtick in the valid character set
- build: add tagged-release script to automate the release process

IMPROVEMENTS

- reader: report error when extra characters are found in one-line ACH files
- webui: display error when parsing fails
- deps: remove moov-io/customers and moov-io/paygate as dependencies
- docs: improvements to README by fixing typos, dead links, and formatting

## v1.5.2 (Released 2020-11-16)

BUG FIXES

- server: pass validation options through to ACH reader and underlying batch

IMPROVEMENTS

- build: push ach-webui image and split up commands for each docker build

## v1.5.1 (Released 2020-11-12)

There are no changes from v1.5.0. This release fixes issues during the release process.

## v1.5.0 (Released 2020-11-12)

ADDITIONS

- cmd/webui: initial setup for client-side ACH file parsing to their JSON forms
- entries: allow custom `TransactionCode` validation
- entries: allow custom `TraceNumber` values
- writer: allow setting `Writer.LineEnding` to use custom values for ending each outputted
- server: allow `ValidateOpts` to be set through HTTP query parameters

IMPROVEMENTS

- build: check and lint OpenAPI spec in CI
- api: update summaries of endpoints
- api: clean up OpenAPI spec after running Speccy
- docs: update content for moov-io.github.io/ach/
- docs: add page on return files
- docs: add link to Moov's ACH blog post
- docs: fix filepath example in segmented files
- github: request version in bug reports
- github: add codeowners
- chore(deps): update golang docker tag to v1.15

BUG FIXES

- server: fix batch additions by shimming JSON reading

## v1.4.4 (Released 2020-08-04)

BUG FIXES

- entries: detect overflow when printing large amounts
- reader: only parse IAT entries when we're accumulating an IAT batch

IMPROVEMENTS

- achcli: print BatchControl records as well
- achcli: update help/usage text

## v1.4.3 (Released 2020-07-23)

BUILD

- build: upload achcli binaries on each release

## v1.4.2 (Released 2020-07-23)

IMPROVEMENTS

- api: add example plaintext file for create route
- docs: flip Usage section priority, link to godocs and examples
- file: include the struct field in Unmarshal errors

BUG FIXES

- api: clarify batchNumber in BatchHeader is an integer
- api: include missing CompanyIdentification field on BatchHeader docs
- api: include missing `CompanyIdentification` field on `BatchHeader` docs
- api: quote number-looking example values

## v1.4.1 (Released 2020-07-09)

BUG FIXES

- batch: error if our offset account type is unknown

BUILD

- build: add openshift docker image
- build: enable codeql via github actions
- build: release via Actions, not TravisCI

## v1.4.0 (Released 2020-06-29)

Version v1.4.0 of ACH adds several notable features such as custom validation, a command-line tool `achcli` to describe files, and improvements for verifying NACHA compatibility on slightly malformed files. This release also contains enhanced testing and documentation improvements.

**Custom Validation**

The ACH library (and HTTP server) now supports custom validation with the [`ValidateOpts`](https://godoc.org/github.com/moov-io/ach#ValidateOpts) struct by calling `File.SetValidation(..)` and `Reader.SetValidation(...)`. This offers various options:

- `RequireABAOrigin bool`: Enable or disable routing number validation over the `ImmediateOrigin` file header field
- `BypassOriginValidation`: Skip validation for the `ImmediateOrigin` file header field and allow custom `TraceNumber` values
- `BypassDestinationValidation`: Skip validation for the `ImmediateDestination` file header field and allow custom `TraceNumber` values

The HTTP server also supports reading this struct with camel-cased names when calling the validation route.

**achcli**

`achcli` is a command-line utility for viewing ACH files in a more human readable format. This tool also allows masking `DFIAccountNumber` values with the `-mask` flag.

```
$ achcli -mask 20200601-1002-01.ach
Describing ACH file '20200601-1002-01.ach'

  Origin     OriginName               Destination  DestinationName     FileCreationDate  FileCreationTime
  691000134  ASF APPLICATION SUPERVI  091400606    FIRST BANK & TRUST  181017            0306

  BatchNumber  SECCode  ServiceClassCode  CompanyName  CompanyDiscretionaryData  CompanyIdentification  CompanyEntryDescription
  1            WEB      200               CoinLion                               123456789              TRANSFER

    TransactionCode  RDFIIdentification  AccountNumber  Amount  Name                    TraceNumber      Category
    26               09140060            *******89      12354   Paul Jones              091000017611242  Return

      Addenda99
      ReturnCode  OriginalTrace    DateOfDeath  OriginalDFI  AddendaInformation  TraceNumber
      R01         091400600000001               09100001                         091000017611242

  BatchCount  BlockCount  EntryAddendaCount  TotalDebitAmount  TotalCreditAmount
  1           1           1                  12354             0
```

**Malformed Files**

ACH files with lines that are not 94 characters are now adjusted in-memory (missing or extra spaces) in an attempt to comply with NACHA standards. The underlying file on disk is not modified during this reading.

-----

ADDITIONS

- batches: Add `LiftEffectiveEntryDate()` to offer parsed `time.Time` values of `EffectiveEntryDate`
- cmd/server: add version handler to admin HTTP server
- file: add BypassDestinationValidation to ValidateOpts
- file: add `ValidateWith` to override specific default validations
- file: support setting ValidateOpts on struct for calling Create()
- reader: morph lines to 94 characters if they end in spaces
- server: read `ValidateOpts` in HTTP validate route
- server: return fileID on create errors, enforce marshaled errors as strings
- file: support setting `ValidateOpts` on struct for calling `Create()`
- file: struct unmarshaling works again, it was depreciated for a couple releases
- reader: morph lines to 94 characters with spaces if they are some other length
- reader: allow setting ValidateOpts
- cmd/ach: initial setup of CLI tool to pretty print ACH files

BUG FIXES

- all: replace `Ç` with `C` across the project
- all: use filepath.Join instead of unix paths
- api: fixup flatten files OpenAPI spec
- api: note POST /files/create can return an error
- file: don't validate before flattening batches
- file: keep TraceNumbers when segmenting files
- server: fix segment OpenAPI spec and accept config body
- server: read empty SegmentFileConfiguration
- file: don't validate before flattening batches

IMPROVEMENTS

- api: used shared Error model
- api: use shorter summaries
- api: include AddendaXX, ADV, and IAT records that were missing from OpenAPI spec
- chore(deps): update module prometheus/client_golang to v1.4.1
- chore(deps): update module gorilla/mux to v1.7.4
- reader: write a test for what partial file comes back from invalid EntryDetails
- reader: allow zero-entry files if their controls signify as such
- server: use FoundABug error with mismatched routing
- validators: ensure alpha routing number check digit is invalid
- all: use filepath.Join instead of unix paths
- reader: append a lingering batch even if there's no batch control

BUILD

- chore(deps): update golang docker tag to v1.14
- build: run sonatype-nexus-community/nancy in CI
- build: leverage moov-io/infra's Go linter script

## v1.3.1 (Released 2020-01-22)

BUG FIXES

- api,client: There was a mistaken character in the OpenAPI docs `Ç` which should have been `C`.

IMPROVEMENTS

- build: upgrade golang.org/x/crypto

BUILD

- build: run sonatype-nexus-community/nancy in CI

## v1.3.0 (Released 2020-01-20)

BREAKING CHANGES

- `ImmediateOrigin` values are written with a leading space instead of a zero (`0`) due to post-2013 NACHA guidelines.

BUG FIXES

- `addenda98` fix parsing with no spaces between routing and account number

ADDITIONS

- Add random names, amounts, and memo lines in test file generation script `cmd/writeACH/main.go`

## v1.2.1 (Released 2019-10-11)

BUG FIXES

- fileHeader: Remove requirement of ImmediateOrigin to be a routing number

## v1.2.0 (Released 2019-10-07)

ADDITIONS

- Add `FlattenBatches() (*File, error)` to `ach.File`
   - FlattenBatches [minimizes File Batches by consolidating them](./docs/flatten-batches.md) with the same BatchHeader data into one batch.
- Add `POST /files/:id/flatten` which calls `FlattenBatches()` on a specific ACH file
- Add `POST /files/:id/balance` to [add Offset records](./docs/balanced-offset.md) onto each Batch in an ACH File.
- Addenda98: Add `ChangeCodeField()` for detailed information about a NOC/COR change file ([`ChangeCode`](https://godoc.org/github.com/moov-io/ach#ChangeCode))
- Addenda99: Add `ReturnCodeField()` for detailed information about file returns ([`ReturnCode`](https://godoc.org/github.com/moov-io/ach#ReturnCode))

BUG FIXES

- reader: set EntryDetail.Category to Return when Addenda99 is present
- batch: inspect Entrydetail records for Category
- batch: check ADV entries for Category
- reader: set EntryDetail.Category to NOC when Addenda98 is present
- file: Validate files after reading them from their JSON representation
- server: actaully render new Credit and Debit files from segmentation

IMPROVEMENTS

- created example files for HTTP routes
- file: parse ISO8601 and RFC3339 timestamps in JSON blobs

BUILD

- upgrade to Go 1.13
- build: download CI tools rather than install
- build: update staticcheck with Go 1.13

## v1.1.0 (Released 2019-08-19)

BREAKING CHANGES

In our OpenAPI we've renamed fields generated as `Id` to `ID`, which is more in-line with Go's style conventions.

BUG FIXES

- fileHeader: allow immediate origin to be a 10 digit value (See: [#513](https://github.com/moov-io/ach/pull/513) by [@eduardev](https://github.com/eduardev))
- Fix JSON omitempty typo in `ADVEntryDetail`
- fileHeader: trim padded 0's from ImmediateOriginField() and fixup docs
- batch: only check DNE specifics if the SEC code is DNE
- files: FileHeader validate CheckRoutingNumber
- files: on empty FileCreation dates or times render today's value
- reader: return ErrFileAddendaOutsideBatch from parseAddenda when currentBatch is nil

ADDITIONS

- batch: add Equal method
- Addenda99: Add `ReturnCodeField()` for detailed information about a returned file
- files: support arbitrary merging of ACH files (See [#529](https://github.com/moov-io/ach/issues/529))
- entryDetail: validate that Amount is non-negative
- batch: create Debit and Credit EntryDetail offset records if needed (via `WithOffset`)
- addenda types: Add RuneCountInString check to Parse(record string) function
- file: create debit ach file and credit ach file from a mixed debit and credit ach file (via `SegmentFile`) (see [#528](https://github.com/moov-io/ach/issues/528))
- cmd/server: add environment variables to override command line flags (`-http.addr` and `-log.format`)
- file: support ADV and IAT files in (*File).SegmentFile(...)
- cmd/server: bind HTTP server with TLS if HTTPS_* variables are defined
- cmd/server: add endpoints for segmenting files into debits and credits

BUILD

- vendor: update dependencies
- Fix `moov/achfuzz` docker image build
- api: inline parameters to fix codegen crash
- build: push moov/ach:latest and update docs accordingly
- chore(deps): update module prometheus/client_golang to v1.1.0

## v1.0.0 (Released 2019-03-26)

- No changes from v0.6.0

## v0.6.0 (Released 2019-03-26)

BREAKING CHANGES

- `file.NotificationOfChange` accepts `Batcher` instead of `*BatchCOR` to comply with linter errors

ADDITIONS

- Add const values for `BatchHeader.StandardEntryClassCode` (See [#392](https://github.com/moov-io/ach/issues/392))
- Add const values for `BatchHeader.ServiceClassCode` and `BatchControl.ServiceClassCode`. (See [#391](https://github.com/moov-io/ach/issues/391))
- Add const values for `EntryDetail.TransactionCode` (See [#363](https://github.com/moov-io/ach/issues/363))
- server: Record `ach_files_deleted` metric. (See: [#408](https://github.com/moov-io/ach/pull/408))
- server: log x-request-id header if present. (See: [#407](https://github.com/moov-io/ach/pull/407))
- server: Delete old `ach.File` objects from in-memory repository according to `ACH_FILE_TTL` env variable.
- server: Support `-log.format=json` for JSON formatted logs

BUG FIXES

- Accept File's without an ID specified. Generate a random ID. (See: [#405](https://github.com/moov-io/ach/pull/405))
- server: Fix nil panics. (See: [#406](https://github.com/moov-io/ach/pull/406))
- server: Fix type-casting panics. (See: [#423](https://github.com/moov-io/ach/pull/423))
- server: validate file endpoint returns 400 instead of 500 (See: [#488](https://github.com/moov-io/ach/pull/488))
- server: set CORS headers on `GET /ping` route

BUILD

- `megacheck` is deprecated. staticcheck should be used instead. (See [#430](https://github.com/moov-io/ach/issues/430))
- Automate releases with Docker and binary uploads to release page.
- Update dependencies to their latest versions
- Update to Go 1.12

## v0.5.0 (Released 2018-11-29)

BREAKING CHANGES

- `TraceNumber` has been changed from `int` to a `string`. (See [#366](https://github.com/moov-io/ach/issues/366))
   - Previously zero-prefixed ABA routing numbers would have their leading zero truncated.
- `OriginalTrace` has been changed from `int` to a `string`. (See [#366](https://github.com/moov-io/ach/issues/366))

ADDITIONS

- Support `StandardEntryClassCode` (Batch types):
    - ADV (See [#340](https://github.com/moov-io/ach/issues/340))
    - TRC (See [#346](https://github.com/moov-io/ach/issues/346))
    - TRX (See [#372](https://github.com/moov-io/ach/issues/372))
    - XCK (See [#347](https://github.com/moov-io/ach/issues/347))
- `TransactionCode` match `ServiceClassCode` (See [#56](https://github.com/moov-io/ach/issues/56))
- `Addenda02.TerminalState` validation for BatchPOS and BatchSHR (See [#375](https://github.com/moov-io/ach/issues/375))

REMOVALS

- Remove deprecated functions from `EntryDetail` (See [#385](https://github.com/moov-io/ach/issues/385))

## v0.4.0 (Released 2018-11-06)

BREAKING CHANGES

- `EntryDetail.Addendum` has been broken out into `Addenda02`, `Addenda05`, `Addenda98`, and `Addenda99` fields on `EntryDetail`.
- IAT `EntryDetail.Addendum` has been broken out into Addenda 10-18, 98 and 99.

ADDITIONS

- Support `StandardEntryClassCode` (Batch types):
  - ACK (See [#327](https://github.com/moov-io/ach/issues/327))
  - ATX (See [#327](https://github.com/moov-io/ach/issues/327))
  - DNE (See [#342](https://github.com/moov-io/ach/issues/342))
  - ENR (See [#343](https://github.com/moov-io/ach/issues/343))
- Support NOC for IAT Entries (See [#328](https://github.com/moov-io/ach/issues/328))
- Add `FileFromJson` for reading `File` objects as JSON.
- Add `X-Total-Count` response headers on `GET /files/:id/batches` (See [#280](https://github.com/moov-io/ach/issues/280))

IMPROVEMENTS

- Multiple parsing errors are returned in `Reader.Read() error`'s error message.
- IAT: Validate ISODestinationCountryCode according to ISO 3166
- IAT: Validate ISOOriginatingCurrencyCode and ISODestinationCurrencyCode according to ISO 4217
- build: Run CI tests against Windows
- Verify record lengths in [IAT] BatchHeader, [IAT] BatchControl, FileControl, FileHeader, and [IAT] EntryDetail.

BUG FIXES
- `cmd/server`: don't expect trailing slash on endpoints
- `cmd/server`: Grab write lock on delete requests
- Several panics are fixed from fuzzing
- Reject invalid ABA routing numbers

## v0.3.3 (Released 2018-10-08)

ADDITIONS

- Expose `TypeCode` on Addenda records

IMPROVEMENTS

- Run as unprivileged user in Docker image
- `cmd/readACH`: add -json to print `File` as JSON
- `cmd/writeACH`: validate file before writing
- `cmd/writeACH`: add -json to output `File` as JSON

BUG FIXES

- Fix reading batches out from JSON endpoint
- Fix plaintext ACH file rendering endpoint

## v0.3.2 (Released 2018-10-05)

ADDITIONS

- Handle pre-flight requests and proxy back CORS headers

## v0.3.1 (Released 2018-10-05)

ADDITIONS

- Add `X-Total-Count` on GET responses
- Proxy back CORS headers if sent on the request

BUG FIXES

- Drop requirement for /-suffix on GET /files
- Don't trample content-type in encodeResponse

## v0.3.0 (Released 2018-09-26)

FEATURES

- Added HTTP Server
- SEC Code CIE (Customer-Initiated Entry) ([#209](https://github.com/moov-io/ach/issues/209))
- Support IAT ([#211](https://github.com/moov-io/ach/issues/211))
- IAT Returns ([#233](https://github.com/moov-io/ach/issues/233))
- Support CTX ([#212](https://github.com/moov-io/ach/issues/212))

IMPROVEMENTS

- Added admin HTTP service.
- Added `GET /ping` route (unauthed)

BUG FIXES

- Fixes to parsing and validation.
