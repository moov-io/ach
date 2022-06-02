---
layout: page
title: Custom validation
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Custom validation

ACH files can vary sometimes from the official NACHA guidelines due to vendor changes. Moov ACH defaults to NACHA guidelines, so to handle this there's an exported `ValidateWith(opts)` method on some structures (`File`, `FileHeader`, etc).

The [`ValidateOpts`](https://godoc.org/github.com/moov-io/ach#ValidateOpts) struct can have fields added in minor releases without breaking API compatibility with callers. The default values in this struct match with NACHA's guidelines.

## Validation options

The following options can be used with `File.ValidateWith` and `File.SetValidation` to alter the default NACHA validation rules.

### Origin

`RequireABAOrigin bool` can be set to enable routing number validation over the `ImmediateOrigin` file header field. Often the origin can be another value which is significant to the Financial Institution you're uploading ACH Files to.

```
// One-time change to the NACHA validation rules
opts := &ValidateOpts{RequireABAOrigin: true}
if err := file.Header.ValidateWith(opts); err != nil {
    // do something...
}
```

`BypassOriginValidation bool` can be set to skip all validation for the `ImmediateOrigin` file header field. This also allows for custom TraceNumbers which aren't prefixed with a routing number as required by the NACHA specification.

```
// Override the default validation rules on an *ach.File object.
file.SetValidation(&ValidateOpts{
    BypassOriginValidation: true,
})
if err := file.Validate(); err != nil {
    // do something...
}
```

### Destination

`BypassDestinationValidation bool` can be set to skip validation for the `ImmediateDestination` file header field.

### Transaction Codes

```
// CheckTransactionCode allows for custom validation of TransactionCode values
CheckTransactionCode func(code int) error
```

### Trace Numbers

The Nacha/ACH spec requires that trace numbers follow a few rules. This validation option disables them.

- Ascending order of trace numbers within batches
- Trace numbers beginning with their ODFI's routing number
- `AddendaRecordIndicator` is set correctly

```
// CustomTraceNumbers disables validation of TraceNumbers
CustomTraceNumbers bool `json:"customTraceNumbers"`
```

### Batches

```
// AllowZeroBatches allows the file to have zero batches
AllowZeroBatches bool `json:"allowZeroBatches"`

// BypassCompanyIdentificationMatch allows batches in which the Company Identification field
// in the batch header and control do not match.
BypassCompanyIdentificationMatch bool `json:"bypassCompanyIdentificationMatch"`

// UnequalServiceClassCode skips equality checks for the ServiceClassCode in each pair of BatchHeader
// and BatchControl records.
UnequalServiceClassCode bool `json:"unequalServiceClassCode"`

// AllowUnorderedBatchNumebrs allows a file to be read with unordered batch numbers.
AllowUnorderedBatchNumbers bool `json:"allowUnorderedBatchNumbers"`
```

### File Header

```
// AllowMissingFileHeader allows a file to be read without a FileHeader record.
AllowMissingFileHeader bool `json:"allowMissingFileHeader"`
```

### File Control

```
// AllowMissingFileControl allows a file to be read without a FileControl record.
AllowMissingFileControl bool `json:"allowMissingFileControl"`
```

### Returns

```
// CustomReturnCodes can be set to skip validation for the Return Code field in an Addenda99
// This allows for non-standard/deprecated return codes (e.g. R97)
CustomReturnCodes bool `json:"customReturnCodes"`
```

## Reader

An `ach.Reader` can have custom validation rules as well, simply set them prior to reading.

```
r := ach.NewReader(fd) // create an ach.Reader from an os.File
r.SetValidation(&ach.ValidateOpts{
    BypassDestinationValidation: true,
})
file, err := r.Read()
if err != nil {
    // do something...
}
```

## JSON Options

The JSON representation includes the `ValidateOpts` if specified on the `*ach.File` instance.

## HTTP server

The ACH HTTP server can accept `ValidateOpts` when [Validating a file](https://moov-io.github.io/ach/api/#get-/files/{fileID}/validate). This will leverage the above one-off validation methods and return any errors.

**Create a file**
```
curl -X POST --data-binary @test/testdata/ppd-debit.ach http://localhost:8080/files/create
```
```
{"id":"b1910446fd904abc8b2cee358ffb3673c2cb8a62","error":null}
```

**Apply custom validation rules**

```
curl -X POST --data-binary '{"requireABAOrigin": true}' http://localhost:8080/files/b1910446fd904abc8b2cee358ffb3673c2cb8a62/validate
```
```
{"error":null}
```
