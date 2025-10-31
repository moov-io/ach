// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ach

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/moov-io/base"
)

// First position of all Record Types. These codes are uniquely assigned to
// the first byte of each row in a file.
const (
	fileHeaderPos   = "1"
	batchHeaderPos  = "5"
	entryDetailPos  = "6"
	entryAddendaPos = "7"
	batchControlPos = "8"
	fileControlPos  = "9"

	// RecordLength character count of each line representing a letter in a file
	RecordLength = 94
)

// FileError is an error describing issues validating a file
type FileError struct {
	FieldName string
	Value     string
	Msg       string
}

func (e FileError) Error() string {
	return fmt.Sprintf("%s %s", e.FieldName, e.Msg)
}

// File contains the structures of a parsed ACH File.
type File struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a file.
	ID string `json:"id"`

	Header     FileHeader     `json:"fileHeader"`
	Batches    []Batcher      `json:"batches"`
	IATBatches []IATBatch     `json:"IATBatches"`
	Control    FileControl    `json:"fileControl"`
	ADVControl ADVFileControl `json:"fileADVControl"`

	// NotificationOfChange (Notification of change) is a slice of references to BatchCOR in file.Batches
	NotificationOfChange []Batcher `json:"NotificationOfChange"`

	// ReturnEntries is a slice of references to file.Batches that contain return entries
	ReturnEntries []Batcher `json:"ReturnEntries"`

	validateOpts *ValidateOpts
}

// NewFile constructs a file template.
func NewFile() *File {
	return &File{
		Header:  NewFileHeader(),
		Control: NewFileControl(),
	}
}

type file struct {
	ID string `json:"id"`
}

type fileHeader struct {
	Header FileHeader `json:"fileHeader"`
}

type fileControl struct {
	Control FileControl `json:"fileControl"`
}

type advFileControl struct {
	ADVControl ADVFileControl `json:"advFileControl"`
}

// FileFromJSON attempts to return a *File object assuming the input is valid JSON.
//
// Callers should always check for a nil-error before using the returned file.
//
// The File returned may not be valid and an error may be returned from validation.
// Invalid files may be rejected by Financial Institutions or ACH tools.
//
// Date and Time fields in formats: RFC 3339 and ISO 8601 will be parsed and rewritten
// as their YYMMDD (year, month, day) or hhmm (hour, minute) formats.
func FileFromJSON(bs []byte) (*File, error) {
	return FileFromJSONWith(bs, nil)
}

// ReadJSONFile will consume the specified filepath and parse the contents as a JSON formatted ACH file.
func ReadJSONFile(path string) (*File, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return FileFromJSON(bs)
}

// ReadJSONFileWith will consume the specified filepath and parse the contents
// as a JSON formatted ACH file with custom ValidateOpts.
func ReadJSONFileWith(path string, opts *ValidateOpts) (*File, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return FileFromJSONWith(bs, opts)
}

// FileFromJSONWith attempts to return a *File object assuming the input is valid JSON.
//
// It allows custom validation overrides, so the file may not be Nacha compliant
// after parsing. Invalid files may be rejected by Financial Institutions or ACH tools.
//
// Callers should always check for a nil-error before using the returned file.
//
// Date and Time fields in formats: RFC 3339 and ISO 8601 will be parsed and rewritten
// as their YYMMDD (year, month, day) or hhmm (hour, minute) formats.
func FileFromJSONWith(bs []byte, opts *ValidateOpts) (*File, error) {
	if len(bs) == 0 {
		return nil, errors.New("no JSON data provided")
	}
	if !json.Valid(bs) {
		return nil, fmt.Errorf("problem reading File: %w", ErrInvalidJSON)
	}

	// Read the ValidateOpts first
	validateOpts, err := readValidateOpts(bs)
	if err != nil {
		return nil, fmt.Errorf("reading validate opts: %w", err)
	}

	out := NewFile()
	out.SetValidation(opts.merge(validateOpts))

	// read file root level
	var f file
	if err := json.NewDecoder(bytes.NewReader(bs)).Decode(&f); err != nil {
		return nil, fmt.Errorf("problem reading File: %v", err)
	}
	out.ID = f.ID

	// Read FileHeader
	header := fileHeader{
		Header: out.Header,
	}
	if err := json.NewDecoder(bytes.NewReader(bs)).Decode(&header); err != nil {
		return nil, fmt.Errorf("problem reading FileHeader: %v", err)
	}
	out.Header = header.Header

	// Build resulting file
	if err := out.setBatchesFromJSON(bs); err != nil {
		return nil, err
	}

	// Overwrite various timestamps with their ACH formatted values
	out.overwriteDateTimeFields()

	if !out.IsADV() {
		// Read FileControl
		control := fileControl{
			Control: NewFileControl(),
		}
		if err := json.NewDecoder(bytes.NewReader(bs)).Decode(&control); err != nil {
			return nil, fmt.Errorf("problem reading FileControl: %v", err)
		}
		out.Control = control.Control
	} else {
		// Read ADVFileControl
		advControl := advFileControl{
			ADVControl: NewADVFileControl(),
		}
		if err := json.NewDecoder(bytes.NewReader(bs)).Decode(&advControl); err != nil {
			return nil, fmt.Errorf("problem reading ADVFileControl: %v", err)
		}
		out.ADVControl = advControl.ADVControl
	}

	if !out.IsADV() {
		out.Control.BatchCount = len(out.Batches)
	} else {
		out.ADVControl.BatchCount = len(out.Batches)
	}

	if err := out.Create(); err != nil {
		return out, err
	}
	if err := out.Validate(); err != nil {
		return out, err
	}
	return out, nil
}

// MarshalJSON will produce a JSON blob with the ACH file's fields and validation settings.
func (f *File) MarshalJSON() ([]byte, error) {
	type Aux struct {
		File
		ValidateOpts *ValidateOpts `json:"validateOpts"`
	}
	return json.Marshal(Aux{
		File:         *f,
		ValidateOpts: f.validateOpts,
	})
}

// UnmarshalJSON parses a JSON blob with ach.FileFromJSON
func (f *File) UnmarshalJSON(p []byte) error {
	// merge any validate opts with the current file
	opts, err := readValidateOpts(p)
	if err != nil {
		return err
	}
	f.SetValidation(f.validateOpts.merge(opts))

	// Read the file
	file, err := FileFromJSONWith(p, f.validateOpts)
	if err != nil {
		return err
	}
	if file != nil {
		*f = *file
	}

	return nil
}

func readValidateOpts(p []byte) (*ValidateOpts, error) {
	type Aux struct {
		ValidateOpts *ValidateOpts `json:"validateOpts"`
	}
	var opts Aux
	err := json.Unmarshal(p, &opts)
	if err != nil {
		return nil, err
	}
	return opts.ValidateOpts, nil
}

type batchesJSON struct {
	Batches []*Batch `json:"batches"`
}

type iatBatchesJSON struct {
	IATBatches []IATBatch `json:"iatBatches"`
}

func setEntryRecordType(e *EntryDetail) {
	if e == nil {
		return
	}
	if e.Addenda02 != nil {
		e.Addenda02.TypeCode = "02"
	}
	for _, a := range e.Addenda05 {
		if a != nil {
			a.TypeCode = "05"
		}
	}
	if e.Addenda98 != nil {
		e.Addenda98.TypeCode = "98"
	}
	if e.Addenda98Refused != nil {
		e.Addenda98Refused.TypeCode = "98"
	}
	if e.Addenda99 != nil {
		e.Addenda99.TypeCode = "99"
	}
	if e.Addenda99Dishonored != nil {
		e.Addenda99Dishonored.TypeCode = "99"
	}
	if e.Addenda99Contested != nil {
		e.Addenda99Contested.TypeCode = "99"
	}
}

func setADVEntryRecordType(e *ADVEntryDetail) {
	if e == nil {
		return
	}
	if e.Addenda99 == nil {
		e.Category = CategoryForward
	}
}

func setIATEntryRecordType(e *IATEntryDetail) {
	if e == nil {
		return
	}
	// these values need to be inferred from the json field names
	if e.Addenda10 != nil {
		e.Addenda10.TypeCode = "10"
	}
	if e.Addenda11 != nil {
		e.Addenda11.TypeCode = "11"
	}
	if e.Addenda12 != nil {
		e.Addenda12.TypeCode = "12"
	}
	if e.Addenda13 != nil {
		e.Addenda13.TypeCode = "13"
	}
	if e.Addenda14 != nil {
		e.Addenda14.TypeCode = "14"
	}
	if e.Addenda15 != nil {
		e.Addenda15.TypeCode = "15"
	}
	if e.Addenda16 != nil {
		e.Addenda16.TypeCode = "16"
	}
	for _, a := range e.Addenda17 {
		if a != nil {
			a.TypeCode = "17"
		}
	}
	for _, a := range e.Addenda18 {
		if a != nil {
			a.TypeCode = "18"
		}
	}
	if e.Addenda98 != nil {
		e.Addenda98.TypeCode = "98"
	}
	if e.Addenda99 != nil {
		e.Addenda99.TypeCode = "99"
	}
}

// setBatchesFromJson takes bs as JSON and attempts to read out all the Batches within.
//
// We have to break this out as Batcher is an interface (and can't be read by Go's
// json struct tag decoding).
func (f *File) setBatchesFromJSON(bs []byte) error {
	var batches batchesJSON
	var iatBatches iatBatchesJSON

	if err := json.Unmarshal(bs, &batches); err != nil {
		return err
	}
	// Clear out any nil batches
	for i := range f.Batches {
		if f.Batches[i] == nil {
			f.Batches = append(f.Batches[:i], f.Batches[i+1:]...)
		}
	}
	// Add new batches to file
	for i := range batches.Batches {
		if batches.Batches[i] == nil || batches.Batches[i].Header == nil {
			continue
		}

		batch := *batches.Batches[i]
		batch.SetID(batch.Header.ID)
		batch.SetValidation(f.validateOpts)

		for _, e := range batch.Entries {
			e.SetValidation(f.validateOpts)

			// these values need to be inferred from the json field names
			setEntryRecordType(e)

			// A few SEC codes don't follow the standard columns so we have to smush
			// them together as the JSON doesn't support ReceivingCompany separate
			// from IndividualName.
			switch batch.GetHeader().StandardEntryClassCode {
			case ATX, CTX:
				addendaIndicator := e.AddendaRecordIndicator
				addendaField, _ := strconv.Atoi(e.CATXAddendaRecordsField())

				individualName := e.IndividualName

				if addendaIndicator > 0 && addendaField == 0 {
					e.SetCATXAddendaRecords(addendaIndicator)
				}
				if addendaIndicator == 0 && addendaField > 0 {
					e.SetCATXAddendaRecords(addendaField)
				}

				if addendaField == 0 {
					e.SetCATXReceivingCompany(individualName)
				}
			}
		}
		for _, e := range batch.ADVEntries {
			setADVEntryRecordType(e)
		}

		if err := batch.build(); err != nil {
			return batch.Error("Invalid Batch", err, batch.Header.ID)
		}

		// Attach a batch with the correct type
		f.Batches = append(f.Batches, ConvertBatchType(batch))
	}

	if err := json.Unmarshal(bs, &iatBatches); err != nil {
		return err
	}

	// Add new iatBatches to file
	for i := range iatBatches.IATBatches {
		iatBatch := iatBatches.IATBatches[i]

		if iatBatch.Header == nil {
			continue
		}

		iatBatch.ID = iatBatch.Header.ID
		iatBatch.SetValidation(f.validateOpts)

		for _, e := range iatBatch.Entries {
			setIATEntryRecordType(e)
		}

		if err := iatBatch.build(); err != nil {
			return iatBatch.Error("from JSON", err)
		}
		f.IATBatches = append(f.IATBatches, iatBatch)
	}

	return nil
}

// overwriteDateTimeFields will scan through fields in a File for Date / Time
// values which are not in their ACH format (YYMMDD, hhmm). It'll attempt to parse
// various formats and overwrite them to the expected values (YYMMDD, hhmm).
func (f *File) overwriteDateTimeFields() {
	// File header
	if t, err := datetimeParse(f.Header.FileCreationDate); err == nil {
		f.Header.FileCreationDate = t.Format("060102")
	}
	if t, err := datetimeParse(f.Header.FileCreationTime); err == nil {
		f.Header.FileCreationTime = t.Format("1504")
	}

	// Batches
	for i := range f.Batches {
		// BatchHeader
		header := f.Batches[i].GetHeader()
		if t, err := datetimeParse(strings.TrimPrefix(header.CompanyDescriptiveDate, "SD")); err == nil {
			header.CompanyDescriptiveDate = "SD" + t.Format("1504")
		}
		if t, err := datetimeParse(header.EffectiveEntryDate); err == nil {
			header.EffectiveEntryDate = t.Format("060102")
		}
		f.Batches[i].SetHeader(header)
	}

	// TODO(adam): Addenda99 has DateOfDeath which is hard to parse and overwrite with Batcher.GetEntries() copying structs

	// IAT Batches
	for i := range f.IATBatches {
		if t, err := datetimeParse(f.IATBatches[i].Header.EffectiveEntryDate); err == nil {
			f.IATBatches[i].Header.EffectiveEntryDate = t.Format("060102")
		}
	}
}

var datetimeformats = []string{
	"2006-01-02T15:04:05.999Z", // Default javascript (new Date).toISOString()
	"2006-01-02T15:04:05Z",     // ISO 8601 without milliseconds
	time.RFC3339,               // Go default
	"01/02/2006",               // DD/MM/YYYY
}

func datetimeParse(v string) (time.Time, error) {
	for i := range datetimeformats {
		if t, err := time.Parse(datetimeformats[i], v); err == nil && !t.IsZero() {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unknown format: %s", v)
}

// Create will modify the File to tabulate and assemble it into a valid state.
// This includes setting any posting dates, sequence numbers, counts, and sums.
//
// Create requires a FileHeader and at least one Batch if validateOpts.AllowZeroBatches is false.
//
// Since each Batch may modify computable fields in the File, any calls to
// Batch.Create should be done before Create.
//
// To check if the File is Nacha compliant, call Validate or ValidateWith.
func (f *File) Create() error {
	opts := f.validateOpts
	if opts == nil {
		opts = &ValidateOpts{}
	}
	if !opts.SkipAll {
		// Requires a valid FileHeader to build FileControl
		if !opts.AllowMissingFileHeader {
			if err := f.Header.Validate(); err != nil {
				return err
			}
		}

		// If AllowZeroBatches is false, require at least one Batch in the new file.
		if !opts.AllowZeroBatches && (len(f.Batches) <= 0 && len(f.IATBatches) <= 0) {
			return ErrFileNoBatches
		}
	}

	if !f.IsADV() {
		// add 2 for FileHeader/control and reset if build was called twice do to error
		totalRecordsInFile := 2
		batchSeq := 1
		fileEntryAddendaCount := 0
		fileEntryHashSum := 0
		totalDebitAmount := 0
		totalCreditAmount := 0

		for i, batch := range f.Batches {
			// create ascending batch numbers unless batch number has been provided
			if f.Batches[i].GetHeader().BatchNumber <= 1 {
				f.Batches[i].GetHeader().BatchNumber = batchSeq
				f.Batches[i].GetControl().BatchNumber = batchSeq
			}
			batchSeq++
			// sum file entry and addenda records. Assume batch.Create batch properly calculated control
			fileEntryAddendaCount = fileEntryAddendaCount + batch.GetControl().EntryAddendaCount
			// add 2 for Batch header/control + entry added count
			totalRecordsInFile = totalRecordsInFile + 2 + batch.GetControl().EntryAddendaCount
			// sum hash from batch control. Assume Batch.Build properly calculated field.
			fileEntryHashSum = fileEntryHashSum + batch.GetControl().EntryHash
			totalDebitAmount = totalDebitAmount + batch.GetControl().TotalDebitEntryDollarAmount
			totalCreditAmount = totalCreditAmount + batch.GetControl().TotalCreditEntryDollarAmount
		}
		for i, iatBatch := range f.IATBatches {
			// create ascending batch numbers
			if f.IATBatches[i].GetHeader().BatchNumber <= 1 {
				f.IATBatches[i].GetHeader().BatchNumber = batchSeq
				f.IATBatches[i].GetControl().BatchNumber = batchSeq
			}
			batchSeq++
			// sum file entry and addenda records. Assume batch.Create batch properly calculated control
			fileEntryAddendaCount = fileEntryAddendaCount + iatBatch.GetControl().EntryAddendaCount
			// add 2 for Batch header/control + entry added count
			totalRecordsInFile = totalRecordsInFile + 2 + iatBatch.GetControl().EntryAddendaCount
			// sum hash from batch control. Assume Batch.Build properly calculated field.
			fileEntryHashSum = fileEntryHashSum + iatBatch.GetControl().EntryHash
			totalDebitAmount = totalDebitAmount + iatBatch.GetControl().TotalDebitEntryDollarAmount
			totalCreditAmount = totalCreditAmount + iatBatch.GetControl().TotalCreditEntryDollarAmount
		}

		// create FileControl from calculated values
		fc := NewFileControl()
		fc.ID = f.ID
		fc.BatchCount = batchSeq - 1
		// blocking factor of 10 is static default value in f.Header.blockingFactor.
		if (totalRecordsInFile % 10) != 0 {
			fc.BlockCount = totalRecordsInFile/10 + 1
		} else {
			fc.BlockCount = totalRecordsInFile / 10
		}
		fc.EntryAddendaCount = fileEntryAddendaCount

		// If greater than 10 digits, truncate
		fc.EntryHash = fc.converters.leastSignificantDigits(fileEntryHashSum, 10)

		fc.TotalDebitEntryDollarAmountInFile = totalDebitAmount
		fc.TotalCreditEntryDollarAmountInFile = totalCreditAmount
		f.Control = fc
	} else {
		if err := f.createFileADV(); err != nil {
			return err
		}
	}
	return nil
}

// AddBatch appends a Batch to the ach.File
func (f *File) AddBatch(batch Batcher) []Batcher {
	if batch == nil {
		return f.Batches
	}
	if batch.Category() == CategoryNOC {
		f.NotificationOfChange = append(f.NotificationOfChange, batch)
	}
	if batch.Category() == CategoryReturn {
		f.ReturnEntries = append(f.ReturnEntries, batch)
	}
	f.Batches = append(f.Batches, batch)
	return f.Batches
}

// RemoveBatch will delete a given Batcher from an ach.File
func (f *File) RemoveBatch(batch Batcher) {
	if batch.Category() == CategoryNOC {
		for i := 0; i < len(f.NotificationOfChange); i++ {
			if f.NotificationOfChange[i].Equal(batch) {
				f.NotificationOfChange = append(f.NotificationOfChange[:i], f.NotificationOfChange[i+1:]...)
				i--
			}
		}
	}
	if batch.Category() == CategoryReturn {
		for i := 0; i < len(f.ReturnEntries); i++ {
			if f.ReturnEntries[i].Equal(batch) {
				f.ReturnEntries = append(f.ReturnEntries[:i], f.ReturnEntries[i+1:]...)
				i--
			}
		}
	}
	for i := 0; i < len(f.Batches); i++ {
		if f.Batches[i].Equal(batch) {
			f.Batches = append(f.Batches[:i], f.Batches[i+1:]...)
			i--
		}
	}
}

// AddIATBatch appends a IATBatch to the ach.File
func (f *File) AddIATBatch(iatBatch IATBatch) []IATBatch {
	f.IATBatches = append(f.IATBatches, iatBatch)
	return f.IATBatches
}

// SetHeader allows for header to be built.
func (f *File) SetHeader(h FileHeader) *File {
	f.Header = h
	return f
}

// Validate performs checks on each record according to Nacha guidelines.
// Validate will never modify the File.
//
// ValidateOpts may be set to bypass certain rules and will only be applied to the FileHeader.
// The underlying Batches and Entries on this File will use their own ValidateOpts if they are set.
//
// The first error encountered is returned.
func (f *File) Validate() error {
	return f.ValidateWith(f.validateOpts)
}

func (f *File) GetValidation() *ValidateOpts {
	if f == nil {
		return nil
	}
	return f.validateOpts
}

// SetValidation stores ValidateOpts on the File which are to be used to override
// the default NACHA validation rules.
func (f *File) SetValidation(opts *ValidateOpts) {
	if f == nil {
		return
	}

	f.validateOpts = opts
	f.Header.SetValidation(opts)
}

// ValidateOpts contains specific overrides from the default set of validations
// performed on a NACHA file, records and various fields within.
type ValidateOpts struct {
	// SkipAll will disable all validation checks of a File. It has no effect when set on records.
	SkipAll bool `json:"skipAll"`

	// RequireABAOrigin can be set to enable routing number validation
	// over the ImmediateOrigin file header field.
	RequireABAOrigin bool `json:"requireABAOrigin"`

	// BypassOriginValidation can be set to skip validation for the
	// ImmediateOrigin file header field.
	//
	// This also allows for custom TraceNumbers which aren't prefixed with
	// a routing number as required by the NACHA specification.
	BypassOriginValidation bool `json:"bypassOriginValidation"`

	// BypassDestinationValidation can be set to skip validation for the
	// ImmediateDestination file header field.
	//
	// This also allows for custom TraceNumbers which aren't prefixed with
	// a routing number as required by the NACHA specification.
	BypassDestinationValidation bool `json:"bypassDestinationValidation"`

	// CheckTransactionCode allows for custom validation of TransactionCode values
	//
	// Note: Functions cannot be serialized into/from JSON, so this check cannot be used from config files.
	CheckTransactionCode func(code int) error `json:"-"`

	// CustomTraceNumbers disables Nacha specified checks of TraceNumbers:
	// - Ascending order of trace numbers within batches
	// - Trace numbers beginning with their ODFI's routing number
	// - AddendaRecordIndicator is set correctly
	CustomTraceNumbers bool `json:"customTraceNumbers"`

	// AllowZeroBatches allows the file to have zero batches
	AllowZeroBatches bool `json:"allowZeroBatches"`

	// AllowMissingFileHeader allows a file to be read without a FileHeader record.
	AllowMissingFileHeader bool `json:"allowMissingFileHeader"`

	// AllowMissingFileControl allows a file to be read without a FileControl record.
	AllowMissingFileControl bool `json:"allowMissingFileControl"`

	// BypassCompanyIdentificationMatch allows batches in which the Company Identification field
	// in the batch header and control do not match.
	BypassCompanyIdentificationMatch bool `json:"bypassCompanyIdentificationMatch"`

	// CustomReturnCodes can be set to skip validation for the Return Code field in an Addenda99
	// This allows for non-standard/deprecated return codes (e.g. R97)
	CustomReturnCodes bool `json:"customReturnCodes"`

	// UnequalServiceClassCode skips equality checks for the ServiceClassCode in each pair of BatchHeader
	// and BatchControl records.
	UnequalServiceClassCode bool `json:"unequalServiceClassCode"`

	// AllowUnorderedBatchNumebrs allows a file to be read with unordered batch numbers.
	AllowUnorderedBatchNumbers bool `json:"allowUnorderedBatchNumbers"`

	// AllowInvalidCheckDigit allows the CheckDigit field in EntryDetail to differ from
	// the expected calculation
	AllowInvalidCheckDigit bool `json:"allowInvalidCheckDigit"`

	// UnequalAddendaCounts skips checking that Addenda Count fields match their expected and computed values.
	UnequalAddendaCounts bool `json:"unequalAddendaCounts"`

	// PreserveSpaces keeps the spacing before and after values that normally have spaces trimmed during parsing.
	PreserveSpaces bool `json:"preserveSpaces"`

	// AllowInvalidAmounts will skip verifying the Amount is valid for the TransactionCode and entry type.
	AllowInvalidAmounts bool `json:"allowInvalidAmounts"`

	// AllowZeroEntryAmount will skip enforcing the entry Amount to be non-zero
	AllowZeroEntryAmount bool `json:"allowZeroEntryAmount"`

	// AllowSpecialCharacters will permit a wider range of UTF-8 characters in alphanumeric fields
	AllowSpecialCharacters bool `json:"allowSpecialCharacters"`

	// AllowEmptyIndividualName will skip verifying IndividualName fields are populated
	// for SEC codes that require the field to be non-blank (and non-zero)
	AllowEmptyIndividualName bool `json:"allowEmptyIndividualName"`

	// BypassBatchValidation will skip validation for batches in a file and only validate file header and control info
	BypassBatchValidation bool `json:"bypassBatchValidation"`

	// SkipFileCreationValidation will skip validation of the FileCreationTime and FileCreationDate fields in a file header
	SkipFileCreationValidation bool `json:"skipFileCreationValidation"`
}

// merge will combine two ValidateOpts structs and keep any non-zero field values.
func (v *ValidateOpts) merge(other *ValidateOpts) *ValidateOpts {
	// If either ValidateOpts is nil return the other
	if v == nil {
		return other
	}
	if other == nil {
		return v
	}

	out := &ValidateOpts{
		SkipAll:                          v.SkipAll || other.SkipAll,
		RequireABAOrigin:                 v.RequireABAOrigin || other.RequireABAOrigin,
		BypassOriginValidation:           v.BypassOriginValidation || other.BypassOriginValidation,
		BypassDestinationValidation:      v.BypassDestinationValidation || other.BypassDestinationValidation,
		CustomTraceNumbers:               v.CustomTraceNumbers || other.CustomTraceNumbers,
		AllowZeroBatches:                 v.AllowZeroBatches || other.AllowZeroBatches,
		AllowMissingFileHeader:           v.AllowMissingFileHeader || other.AllowMissingFileHeader,
		AllowMissingFileControl:          v.AllowMissingFileControl || other.AllowMissingFileControl,
		BypassCompanyIdentificationMatch: v.BypassCompanyIdentificationMatch || other.BypassCompanyIdentificationMatch,
		CustomReturnCodes:                v.CustomReturnCodes || other.CustomReturnCodes,
		UnequalServiceClassCode:          v.UnequalServiceClassCode || other.UnequalServiceClassCode,
		AllowUnorderedBatchNumbers:       v.AllowUnorderedBatchNumbers || other.AllowUnorderedBatchNumbers,
		AllowInvalidCheckDigit:           v.AllowInvalidCheckDigit || other.AllowInvalidCheckDigit,
		UnequalAddendaCounts:             v.UnequalAddendaCounts || other.UnequalAddendaCounts,
		PreserveSpaces:                   v.PreserveSpaces || other.PreserveSpaces,
		AllowInvalidAmounts:              v.AllowInvalidAmounts || other.AllowInvalidAmounts,
		AllowZeroEntryAmount:             v.AllowZeroEntryAmount || other.AllowZeroEntryAmount,
		AllowSpecialCharacters:           v.AllowSpecialCharacters || other.AllowSpecialCharacters,
		AllowEmptyIndividualName:         v.AllowEmptyIndividualName || other.AllowEmptyIndividualName,
		BypassBatchValidation:            v.BypassBatchValidation || other.BypassBatchValidation,
		SkipFileCreationValidation:       v.SkipFileCreationValidation || other.SkipFileCreationValidation,
	}

	if v.CheckTransactionCode != nil {
		out.CheckTransactionCode = v.CheckTransactionCode
	}
	if other.CheckTransactionCode != nil {
		out.CheckTransactionCode = other.CheckTransactionCode
	}

	return out
}

// ValidateWith performs checks on each record according to Nacha guidelines.
// ValidateWith will never modify the File.
//
// ValidateOpts may be set to bypass certain rules and will only be applied to the FileHeader.
// opts passed in will override ValidateOpts set by SetValidation.
// The underlying Batches and Entries on this File will use their own ValidateOpts if they are set.
//
// The first error encountered is returned.
func (f *File) ValidateWith(opts *ValidateOpts) error {
	if opts == nil {
		opts = &ValidateOpts{}
	}

	if opts.SkipAll {
		return nil
	}

	if !opts.AllowMissingFileHeader {
		if err := f.Header.ValidateWith(opts); err != nil {
			return err
		}
	}

	if !f.IsADV() {
		// The value of the Batch Count Field is equal to the number of Company/Batch/Header Records in the file.
		if f.Control.BatchCount != (len(f.Batches) + len(f.IATBatches)) {
			return NewErrFileCalculatedControlEquality("BatchCount", len(f.Batches), f.Control.BatchCount)
		}

		if !opts.BypassBatchValidation {
			for _, b := range f.Batches {
				if err := b.Validate(); err != nil {
					return err
				}
			}
		}

		if !opts.AllowMissingFileControl {
			if err := f.Control.Validate(); err != nil {
				return err
			}
		}
		if err := f.isEntryAddendaCount(false); err != nil {
			return err
		}
		if err := f.isFileAmount(false); err != nil {
			return err
		}
		if !opts.AllowUnorderedBatchNumbers {
			if err := f.isSequenceAscending(); err != nil {
				return err
			}
		}
		return f.isEntryHash(false)
	}

	// File contains ADV batches BatchADV

	// The value of the Batch Count Field is equal to the number of Company/Batch/Header Records in the file.
	if f.ADVControl.BatchCount != len(f.Batches) {
		return NewErrFileCalculatedControlEquality("BatchCount", len(f.Batches), f.ADVControl.BatchCount)
	}
	if !opts.AllowMissingFileControl {
		if err := f.ADVControl.Validate(); err != nil {
			return err
		}
	}
	if err := f.isEntryAddendaCount(true); err != nil {
		return err
	}
	if err := f.isFileAmount(true); err != nil {
		return err
	}
	return f.isEntryHash(true)
}

// isEntryAddendaCount is prepared by hashing the RDFI's 8-digit Routing Number in each entry.
// The Entry Hash provides a check against inadvertent alteration of data
func (f *File) isEntryAddendaCount(IsADV bool) error {
	// IsADV
	// true: the file contains ADV batches
	// false: the file contains other batch types

	count := 0

	// we assume that each batch block has already validated the addenda count is accurate in batch control.

	if !IsADV {
		for _, batch := range f.Batches {
			count += batch.GetControl().EntryAddendaCount
		}
		for _, iatBatch := range f.IATBatches {
			count += iatBatch.GetControl().EntryAddendaCount
		}
		if f.Control.EntryAddendaCount != count {
			if f.validateOpts != nil && f.validateOpts.UnequalAddendaCounts {
				return nil
			}
			return NewErrFileCalculatedControlEquality("EntryAddendaCount", count, f.Control.EntryAddendaCount)
		}
	} else {
		for _, batch := range f.Batches {
			count += batch.GetADVControl().EntryAddendaCount
		}
		if f.ADVControl.EntryAddendaCount != count {
			if f.validateOpts != nil && f.validateOpts.UnequalAddendaCounts {
				return nil
			}
			return NewErrFileCalculatedControlEquality("EntryAddendaCount", count, f.ADVControl.EntryAddendaCount)
		}
	}
	return nil
}

// isFileAmount The Total Debit and Credit Entry Dollar Amounts Fields contain accumulated
// Entry Detail debit and credit totals within the file
func (f *File) isFileAmount(IsADV bool) error {
	// IsADV
	// true: the file contains ADV batches
	// false: the file contains other batch types

	debit := 0
	credit := 0

	if !IsADV {
		for _, batch := range f.Batches {
			debit += batch.GetControl().TotalDebitEntryDollarAmount
			credit += batch.GetControl().TotalCreditEntryDollarAmount
		}
		// IAT
		for _, iatBatch := range f.IATBatches {
			debit += iatBatch.GetControl().TotalDebitEntryDollarAmount
			credit += iatBatch.GetControl().TotalCreditEntryDollarAmount
		}

		if f.Control.TotalDebitEntryDollarAmountInFile != debit {
			return NewErrFileCalculatedControlEquality("TotalDebitEntryDollarAmountInFile", debit, f.Control.TotalDebitEntryDollarAmountInFile)
		}
		if f.Control.TotalCreditEntryDollarAmountInFile != credit {
			return NewErrFileCalculatedControlEquality("TotalCreditEntryDollarAmountInFile", credit, f.Control.TotalCreditEntryDollarAmountInFile)
		}
	} else {
		for _, batch := range f.Batches {
			debit += batch.GetADVControl().TotalDebitEntryDollarAmount
			credit += batch.GetADVControl().TotalCreditEntryDollarAmount
		}

		if f.ADVControl.TotalDebitEntryDollarAmountInFile != debit {
			return NewErrFileCalculatedControlEquality("TotalDebitEntryDollarAmountInFile", debit, f.ADVControl.TotalDebitEntryDollarAmountInFile)
		}
		if f.ADVControl.TotalCreditEntryDollarAmountInFile != credit {
			return NewErrFileCalculatedControlEquality("TotalCreditEntryDollarAmountInFile", credit, f.ADVControl.TotalCreditEntryDollarAmountInFile)

		}
	}
	return nil
}

// isEntryHash validates the hash by recalculating the result
func (f *File) isEntryHash(IsADV bool) error {
	// IsADV
	// true: the file contains ADV batches
	// false: the file contains other batch types but not ADV

	hashField := f.calculateEntryHash(IsADV)

	if !IsADV {
		if hashField != f.Control.EntryHash {
			return NewErrFileCalculatedControlEquality("EntryHash", hashField, f.Control.EntryHash)
		}
	} else {
		if hashField != f.ADVControl.EntryHash {
			return NewErrFileCalculatedControlEquality("EntryHash", hashField, f.ADVControl.EntryHash)
		}
	}
	return nil
}

// calculateEntryHash This field is prepared by hashing the 8-digit Routing Number in each batch.
// The Entry Hash provides a check against inadvertent alteration of data
func (f *File) calculateEntryHash(IsADV bool) int {
	// IsADV
	// true: the file contains ADV batches
	// false: the file contains other batch types but not ADV

	hash := 0

	if !IsADV {
		for _, batch := range f.Batches {
			hash = hash + batch.GetControl().EntryHash
		}
		// IAT
		for _, iatBatch := range f.IATBatches {
			hash = hash + iatBatch.GetControl().EntryHash
		}
	} else {
		for _, batch := range f.Batches {
			hash = hash + batch.GetADVControl().EntryHash
		}
	}

	// Ensure the entry hash cannot exceed 10 digits
	// If greater than 10 digits, truncate
	return f.Control.leastSignificantDigits(hash, 10)
}

// IsADV determines if the File is a File containing ADV batches
func (f *File) IsADV() bool {
	for i := range f.Batches {
		if v := f.Batches[i].GetHeader(); v == nil {
			f.Batches[i].SetHeader(NewBatchHeader())
		}
		if v := f.Batches[i].GetControl(); v == nil {
			f.Batches[i].SetControl(NewBatchControl())
		}
		if f.Batches[i].GetHeader().StandardEntryClassCode == ADV {
			return true
		}
	}
	return false
}

func (f *File) createFileADV() error {
	// add 2 for FileHeader/control and reset if build was called twice do to error
	totalRecordsInFile := 2
	batchSeq := 1
	fileEntryAddendaCount := 0
	fileEntryHashSum := 0
	totalDebitAmount := 0
	totalCreditAmount := 0

	for i, batch := range f.Batches {
		// create ascending batch numbers

		if batch.GetHeader().StandardEntryClassCode != ADV {
			return ErrFileADVOnly
		}

		if f.Batches[i].GetHeader().BatchNumber <= 1 {
			f.Batches[i].GetHeader().BatchNumber = batchSeq
			f.Batches[i].GetADVControl().BatchNumber = batchSeq
		}
		batchSeq++
		// sum file entry and addenda records. Assume batch.Create batch properly calculated control
		fileEntryAddendaCount = fileEntryAddendaCount + batch.GetADVControl().EntryAddendaCount
		// add 2 for Batch header/control + entry added count
		totalRecordsInFile = totalRecordsInFile + 2 + batch.GetADVControl().EntryAddendaCount
		// sum hash from batch control. Assume Batch.Build properly calculated field.
		fileEntryHashSum = fileEntryHashSum + batch.GetADVControl().EntryHash
		totalDebitAmount = totalDebitAmount + batch.GetADVControl().TotalDebitEntryDollarAmount
		totalCreditAmount = totalCreditAmount + batch.GetADVControl().TotalCreditEntryDollarAmount
	}

	fc := NewADVFileControl()
	fc.ID = f.ID
	fc.BatchCount = batchSeq - 1
	// blocking factor of 10 is static default value in f.Header.blockingFactor.
	if (totalRecordsInFile % 10) != 0 {
		fc.BlockCount = totalRecordsInFile/10 + 1
	} else {
		fc.BlockCount = totalRecordsInFile / 10
	}
	fc.EntryAddendaCount = fileEntryAddendaCount
	fc.EntryHash = fileEntryHashSum
	fc.TotalDebitEntryDollarAmountInFile = totalDebitAmount
	fc.TotalCreditEntryDollarAmountInFile = totalCreditAmount
	f.ADVControl = fc

	return nil
}

// SegmentFile takes a valid ACH File and returns 2 segmented ACH Files, one ACH File containing credit entries
// and one ACH File containing debit entries.  The return is 2 Files a Credit File and Debit File, or an error.
//
// Callers should always check for a nil-error before using the returned file.
//
// The File returned may not be valid and callers should confirm with Validate. Invalid files may be rejected
// by other Financial Institutions or ACH tools.
func (f *File) SegmentFile(_ *SegmentFileConfiguration) (*File, *File, error) {
	if err := f.Validate(); err != nil {
		return nil, nil, err
	}

	creditFile := NewFile()
	debitFile := NewFile()

	if f.validateOpts != nil {
		creditFile.SetValidation(f.validateOpts)
		debitFile.SetValidation(f.validateOpts)
	}

	if f.Batches != nil {
		err := f.segmentFileBatches(creditFile, debitFile)
		if err != nil {
			return nil, nil, err
		}
	}

	if f.IATBatches != nil {
		f.segmentFileIATBatches(creditFile, debitFile)
	}

	// Additional Sorting to be FI specific
	if len(creditFile.Batches) != 0 || len(creditFile.IATBatches) != 0 {
		f.addFileHeaderData(creditFile)
		if err := creditFile.Create(); err != nil {
			return nil, nil, err
		}
		if err := creditFile.Validate(); err != nil {
			return nil, nil, err
		}
	}
	if len(debitFile.Batches) != 0 || len(debitFile.IATBatches) != 0 {
		f.addFileHeaderData(debitFile)
		if err := debitFile.Create(); err != nil {
			return nil, nil, err
		}
		if err := debitFile.Validate(); err != nil {
			return nil, nil, err
		}
	}
	return creditFile, debitFile, nil
}

func (f *File) segmentFileBatches(creditFile, debitFile *File) error {
	for _, batch := range f.Batches {
		bh := batch.GetHeader()

		var creditBatch Batcher
		var debitBatch Batcher

		switch bh.StandardEntryClassCode {
		case ADV:
			switch bh.ServiceClassCode {
			case AutomatedAccountingAdvices:
				bh := createSegmentFileBatchHeader(AutomatedAccountingAdvices, bh)
				creditBatch, _ = NewBatch(bh)
				debitBatch, _ = NewBatch(bh)

				entries := batch.GetADVEntries()
				for _, entry := range entries {
					err := segmentFileBatchAddADVEntry(creditBatch, debitBatch, entry)
					if err != nil {
						return err
					}
				}
				// Add the Entry to its Batch
				if creditBatch != nil && len(creditBatch.GetADVEntries()) > 0 {
					_ = creditBatch.Create()
					creditFile.AddBatch(creditBatch)
				}

				if debitBatch != nil && len(debitBatch.GetADVEntries()) > 0 {
					_ = debitBatch.Create()
					debitFile.AddBatch(debitBatch)
				}
			}
		default:
			switch bh.ServiceClassCode {
			case MixedDebitsAndCredits:
				cbh := createSegmentFileBatchHeader(CreditsOnly, bh)
				creditBatch, _ = NewBatch(cbh)

				dbh := createSegmentFileBatchHeader(DebitsOnly, bh)
				debitBatch, _ = NewBatch(dbh)

				entries := batch.GetEntries()
				for _, entry := range entries {
					err := segmentFileBatchAddEntry(creditBatch, debitBatch, entry)
					if err != nil {
						return err
					}
				}

				if creditBatch != nil && len(creditBatch.GetEntries()) > 0 {
					_ = creditBatch.Create()
					creditFile.AddBatch(creditBatch)
				}
				if debitBatch != nil && len(debitBatch.GetEntries()) > 0 {
					_ = debitBatch.Create()
					debitFile.AddBatch(debitBatch)
				}
			case CreditsOnly:
				creditFile.AddBatch(batch)
			case DebitsOnly:
				debitFile.AddBatch(batch)
			}
		}
	}
	return nil
}

// segmentFileIATBatches segments IAT batches debits and credits into debit and credit files
func (f *File) segmentFileIATBatches(creditFile, debitFile *File) {
	for _, iatb := range f.IATBatches {
		IATBh := iatb.GetHeader()

		switch IATBh.ServiceClassCode {
		case MixedDebitsAndCredits:
			cbh := createSegmentFileIATBatchHeader(CreditsOnly, IATBh)
			creditIATBatch := NewIATBatch(cbh)

			dbh := createSegmentFileIATBatchHeader(DebitsOnly, IATBh)
			debitIATBatch := NewIATBatch(dbh)

			entries := iatb.GetEntries()
			for _, IATEntry := range entries {
				IATEntry.TraceNumber = "" // unset so Batch.build generates a TraceNumber
				switch IATEntry.TransactionCode {
				case CheckingCredit, CheckingReturnNOCCredit, CheckingPrenoteCredit, CheckingZeroDollarRemittanceCredit,
					SavingsCredit, SavingsReturnNOCCredit, SavingsPrenoteCredit, SavingsZeroDollarRemittanceCredit,
					GLCredit, GLReturnNOCCredit, GLPrenoteCredit, GLZeroDollarRemittanceCredit,
					LoanCredit, LoanReturnNOCCredit, LoanPrenoteCredit, LoanZeroDollarRemittanceCredit:
					creditIATBatch.AddEntry(IATEntry)
				case CheckingDebit, CheckingReturnNOCDebit, CheckingPrenoteDebit, CheckingZeroDollarRemittanceDebit,
					SavingsDebit, SavingsReturnNOCDebit, SavingsPrenoteDebit, SavingsZeroDollarRemittanceDebit,
					GLDebit, GLReturnNOCDebit, GLPrenoteDebit, GLZeroDollarRemittanceDebit,
					LoanDebit, LoanReturnNOCDebit:
					debitIATBatch.AddEntry(IATEntry)
				}
			}

			if len(creditIATBatch.GetEntries()) > 0 {
				_ = creditIATBatch.Create()
				creditFile.AddIATBatch(creditIATBatch)
			}
			if len(debitIATBatch.GetEntries()) > 0 {
				_ = debitIATBatch.Create()
				debitFile.AddIATBatch(debitIATBatch)
			}
		case CreditsOnly:
			creditFile.AddIATBatch(iatb)
		case DebitsOnly:
			debitFile.AddIATBatch(iatb)
		}
	}

}

// createSegmentFileBatchHeader adds BatchHeader data for a debit/credit Segment File
func createSegmentFileBatchHeader(serviceClassCode int, bh *BatchHeader) *BatchHeader {
	nbh := NewBatchHeader()
	nbh.ID = base.ID()
	nbh.ServiceClassCode = serviceClassCode
	nbh.CompanyName = bh.CompanyName
	nbh.CompanyDiscretionaryData = bh.CompanyDiscretionaryData
	nbh.CompanyIdentification = bh.CompanyIdentification
	nbh.StandardEntryClassCode = bh.StandardEntryClassCode
	nbh.CompanyEntryDescription = bh.CompanyEntryDescription
	nbh.CompanyDescriptiveDate = bh.CompanyDescriptiveDate
	nbh.EffectiveEntryDate = bh.EffectiveEntryDate
	nbh.SettlementDate = bh.SettlementDate
	if serviceClassCode == AutomatedAccountingAdvices {
		nbh.OriginatorStatusCode = 0 // ADV requires this be 0
	} else {
		nbh.OriginatorStatusCode = bh.OriginatorStatusCode
	}
	nbh.ODFIIdentification = bh.ODFIIdentification
	return nbh
}

// createSegmentFileIATBatchHeader adds IATBatchHeader data for a debit/credit Segment File
func createSegmentFileIATBatchHeader(serviceClassCode int, IATBh *IATBatchHeader) *IATBatchHeader {
	nbh := NewIATBatchHeader()
	nbh.ID = base.ID()
	nbh.ServiceClassCode = serviceClassCode
	nbh.ForeignExchangeIndicator = IATBh.ForeignExchangeIndicator
	nbh.ForeignExchangeReferenceIndicator = IATBh.ForeignExchangeReferenceIndicator
	nbh.ISODestinationCountryCode = IATBh.ISODestinationCountryCode
	nbh.OriginatorIdentification = IATBh.OriginatorIdentification
	nbh.StandardEntryClassCode = IATBh.StandardEntryClassCode
	nbh.CompanyEntryDescription = IATBh.CompanyEntryDescription
	nbh.ISOOriginatingCurrencyCode = IATBh.ISOOriginatingCurrencyCode
	nbh.ISODestinationCurrencyCode = IATBh.ISODestinationCurrencyCode
	nbh.ODFIIdentification = IATBh.ODFIIdentification
	return nbh
}

// addFileHeaderData adds FileHeader data for a debit/credit Segment File
func (f *File) addFileHeaderData(file *File) *File {
	file.ID = base.ID()
	file.Header.ID = base.ID()
	file.Header.ImmediateOrigin = f.Header.ImmediateOrigin
	file.Header.ImmediateDestination = f.Header.ImmediateDestination
	file.Header.FileCreationDate = time.Now().Format("060102")
	file.Header.FileCreationTime = time.Now().AddDate(0, 0, 1).Format("1504") // HHmm
	file.Header.FileIDModifier = f.Header.FileIDModifier
	file.Header.ImmediateDestinationName = f.Header.ImmediateDestinationName
	file.Header.ImmediateOriginName = f.Header.ImmediateOriginName
	return file
}

// segmentFileBatchAddEntry adds entries to batches in a segmented file
// Applies to All SEC Codes except ADV (Automated Accounting Advice)
func segmentFileBatchAddEntry(creditBatch, debitBatch Batcher, entry *EntryDetail) error {
	switch entry.TransactionCode {
	case CheckingCredit, CheckingReturnNOCCredit, CheckingPrenoteCredit, CheckingZeroDollarRemittanceCredit,
		SavingsCredit, SavingsReturnNOCCredit, SavingsPrenoteCredit, SavingsZeroDollarRemittanceCredit,
		GLCredit, GLReturnNOCCredit, GLPrenoteCredit, GLZeroDollarRemittanceCredit,
		LoanCredit, LoanReturnNOCCredit, LoanPrenoteCredit, LoanZeroDollarRemittanceCredit:
		if creditBatch == nil {
			return errors.New("missing creditBatch")
		}
		creditBatch.AddEntry(entry)

	case CheckingDebit, CheckingReturnNOCDebit, CheckingPrenoteDebit, CheckingZeroDollarRemittanceDebit,
		SavingsDebit, SavingsReturnNOCDebit, SavingsPrenoteDebit, SavingsZeroDollarRemittanceDebit,
		GLDebit, GLReturnNOCDebit, GLPrenoteDebit, GLZeroDollarRemittanceDebit,
		LoanDebit, LoanReturnNOCDebit:
		if debitBatch == nil {
			return errors.New("missing debitBatch")
		}
		debitBatch.AddEntry(entry)
	}
	return nil
}

// segmentFileBatchAddADVEntry adds entries to batches in a segment file for SEC Code ADV (Automated Accounting Advice)
func segmentFileBatchAddADVEntry(creditBatch Batcher, debitBatch Batcher, entry *ADVEntryDetail) error {
	switch entry.TransactionCode {
	case CreditForDebitsOriginated, CreditForCreditsReceived, CreditForCreditsRejected, CreditSummary:
		if creditBatch == nil {
			return errors.New("missing creditBatch")
		}
		creditBatch.AddADVEntry(entry)

	case DebitForCreditsOriginated, DebitForDebitsReceived, DebitForDebitsRejectedBatches, DebitSummary:
		if debitBatch == nil {
			return errors.New("missing debitBatch")
		}
		debitBatch.AddADVEntry(entry)
	}
	return nil
}

// Validates that the batch numbers are ascending
func (f *File) isSequenceAscending() error {
	lastSeq := 0
	for _, batch := range f.Batches {
		current := batch.GetHeader().BatchNumber
		if f.validateOpts == nil || !f.validateOpts.CustomTraceNumbers {
			if current <= lastSeq {
				return NewErrFileBatchNumberAscending(lastSeq, current)
			}
		}

		lastSeq = current
	}

	return nil
}
