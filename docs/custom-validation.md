## Custom Validation

ACH files can vary sometimes from the official NACHA guidelines due to vendor changes. Moov ACH defaults to NACHA guidelines, so to handle this there's an exported `ValidateWith(opts)` method on some structures (`File`, `FileHeader`, etc).

The [`ValidateOpts`](https://godoc.org/github.com/moov-io/ach#ValidateOpts) struct can have fields added in minor releases without breaking API compatibility with callers. The default values in this struct with match NACHA's guidelines.

## Validation Options

The following options can be used with `File.ValidateWith` and `File.SetValidation` to alter the default NACHA validation rules.

### Origin

`RequireABAOrigin bool` can be set to enable routing number validation over the `ImmediateOrigin` file header field. Often the origin can be another value which is significan to the Financial Institution you're uploading ACH Files to.

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

## HTTP Server

The ACH HTTP server can accept `ValidateOpts` when [Validating a file](https://moov-io.github.io/ach/api/#get-/files/{fileID}/validate). This will leverage the above one-off validation methods and return any errors.

**Create a File**
```
curl -XPOST --data @test/testdata/ppd-debit.ach http://localhost:8080/files/create
```
```
{"id":"b1910446fd904abc8b2cee358ffb3673c2cb8a62","error":null}
```

**Apply Custom Validation Rules**

```
curl -XPOST --data '{"requireABAOrigin": true}' http://localhost:8080/files/b1910446fd904abc8b2cee358ffb3673c2cb8a62/validate
```
```
{"error":null}
```
