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
