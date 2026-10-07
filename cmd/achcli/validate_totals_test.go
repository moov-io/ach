// Copyright 2026 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/moov-io/ach"
	"github.com/stretchr/testify/require"
)

func TestValidateTotalsCLI(t *testing.T) {
	if runtime.GOOS == "js" || runtime.GOOS == "wasip1" {
		t.Skip("executable CLI tests require native subprocesses")
	}
	bin := filepath.Join(t.TempDir(), "achcli")
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}
	build := exec.CommandContext(t.Context(), "go", "build", "-o", bin, ".")
	buildOutput, err := build.CombinedOutput()
	require.NoError(t, err, "%s", buildOutput)

	dir := t.TempDir()
	write := func(name string, data []byte) string {
		path := filepath.Join(dir, name)
		require.NoError(t, os.WriteFile(path, data, 0600))
		return path
	}
	fixture := func(parts ...string) []byte {
		parts = append([]string{"..", ".."}, parts...)
		data, err := os.ReadFile(filepath.Join(parts...))
		require.NoError(t, err)
		return data
	}
	ppd := fixture("test", "testdata", "ppd-debit.ach")
	valid := write("valid.ach", ppd)
	partial := write("partial.ach", fixture("test", "testdata", "ppd-debit-invalid-entryDetail-checkDigit.ach"))
	allowCheckDigit := write("allow-check-digit.json", []byte(`{"allowInvalidCheckDigit":true}`))
	bypassBatch := write("bypass-batch.json", []byte(`{"bypassBatchValidation":true}`))
	skipAll := write("skip-all.json", []byte(`{"skipAll":true}`))
	unequalCounts := write("unequal-counts.json", []byte(`{"unequalAddendaCounts":true}`))
	missingControl := write("missing-control.json", []byte(`{"allowMissingFileControl":true}`))
	missingRecords := write("missing-records.json", []byte(`{"allowMissingFileHeader":true,"allowMissingFileControl":true}`))
	missingHeader := write("missing-header.json", []byte(`{"allowMissingFileHeader":true}`))
	badOptions := write("bad-options.json", []byte("{"))
	empty := write("empty.ach", nil)
	missing := filepath.Join(dir, "missing.ach")
	header := append(append([]byte(nil), bytes.SplitN(ppd, []byte("\n"), 2)[0]...), '\n')
	headerOnly := write("header-only.ach", header)
	zeroControl := ach.FileControl{BlockCount: 1}
	zeroTotals := write("zero-totals.ach", append(append([]byte(nil), header...), []byte(zeroControl.String()+"\n")...))
	shortControl := []byte(zeroControl.String()[:55] + "X")
	shortEOF := write("short-control-eof.ach", shortControl)
	shortLF := write("short-control-lf.ach", append(append([]byte(nil), shortControl...), '\n'))

	type testCase struct {
		name string
		args []string
		code int
		want []string
	}
	cases := []testCase{
		{name: "valid", args: []string{"-validate-totals", valid}},
		{name: "verbose", args: []string{"-validate-totals", "-v", valid}, want: []string{"Validated"}},
		{name: "missing", args: []string{"-validate-totals", missing}, code: 1, want: []string{missing}},
		{name: "empty", args: []string{"-validate-totals", empty}, code: 1, want: []string{empty}},
		{name: "empty with missing record options", args: []string{"-validate-totals", "-validate", missingRecords, empty}, code: 1, want: []string{empty, "FileControl"}},
		{name: "header only with missing control option", args: []string{"-validate-totals", "-validate", missingControl, headerOnly}, code: 1, want: []string{headerOnly, "FileControl"}},
		{name: "complete file with missing record options", args: []string{"-validate-totals", "-validate", missingRecords, valid}},
		{name: "stored zero totals", args: []string{"-validate-totals", zeroTotals}},
		{name: "short control at EOF", args: []string{"-validate-totals", "-validate", missingHeader, shortEOF}},
		{name: "short control with LF", args: []string{"-validate-totals", "-validate", missingHeader, shortLF}},
		{name: "partial parse", args: []string{"-validate-totals", partial}, code: 1, want: []string{partial, "calculated check digit"}},
		{name: "permitted options", args: []string{"-validate-totals", "-validate", allowCheckDigit, partial}},
		{name: "no input", args: []string{"-validate-totals"}, code: 1},
		{name: "skip all options", args: []string{"-validate-totals", "-validate", skipAll, valid}, code: 1, want: []string{"SkipAll"}},
		{name: "unequal count options", args: []string{"-validate-totals", "-validate", unequalCounts, valid}, code: 1, want: []string{"UnequalAddendaCounts"}},
		{name: "invalid options", args: []string{"-validate-totals", "-validate", badOptions, valid}, code: 1},
		{name: "missing options", args: []string{"-validate-totals", "-validate", missing, valid}, code: 1},
		{name: "inactive flags", args: []string{"-validate-totals", "-diff=false", "-fix=false", "-merge=false", "-flatten=false", "-mask=false", "-pretty=false", "-skip-validation=false", "-version=false", "-reformat=", "-update-eed=", valid}},
		{name: "totals help short", args: []string{"-validate-totals", "-h", valid}, code: 1, want: []string{"cannot be combined"}},
		{name: "totals help long", args: []string{"-validate-totals", "-help", valid}, code: 1, want: []string{"cannot be combined"}},
		{name: "default help short", args: []string{"-h"}, want: []string{"USAGE"}},
		{name: "default help long", args: []string{"-help"}, want: []string{"USAGE"}},
		{name: "unknown totals flag", args: []string{"-validate-totals", "-unknown", valid}, code: 2, want: []string{"flag provided but not defined"}},
		{name: "unknown default flag", args: []string{"-unknown"}, code: 2, want: []string{"flag provided but not defined"}},
		{name: "invalid totals flag", args: []string{"-validate-totals=bad", valid}, code: 2, want: []string{"invalid boolean value"}},
		{name: "missing flag argument", args: []string{"-validate-totals", "-validate"}, code: 2, want: []string{"flag needs an argument"}},
	}

	fileChanges := []struct {
		field  string
		change func(*ach.FileControl)
	}{
		{"BatchCount", func(c *ach.FileControl) { c.BatchCount++ }},
		{"EntryAddendaCount", func(c *ach.FileControl) { c.EntryAddendaCount++ }},
		{"EntryHash", func(c *ach.FileControl) { c.EntryHash++ }},
		{"TotalDebitEntryDollarAmountInFile", func(c *ach.FileControl) { c.TotalDebitEntryDollarAmountInFile++ }},
		{"TotalCreditEntryDollarAmountInFile", func(c *ach.FileControl) { c.TotalCreditEntryDollarAmountInFile++ }},
	}
	var invalid string
	for _, change := range fileChanges {
		data := rewriteACHRecord(t, ppd, '9', func(record string) string {
			var control ach.FileControl
			control.Parse(record)
			change.change(&control)
			return control.String()
		})
		path := write(change.field+".ach", data)
		if invalid == "" {
			invalid = path
		}
		cases = append(cases,
			testCase{name: change.field, args: []string{"-validate-totals", path}, code: 1, want: []string{path, change.field}},
			testCase{name: change.field + " default", args: []string{path}, want: []string{"Describing"}},
		)
	}
	for _, flag := range []string{"-version", "-diff", "-fix", "-merge", "-flatten", "-reformat=json", "-update-eed=20260102", "-mask", "-mask.accounts", "-mask.corrections", "-mask.names", "-pretty", "-pretty.amounts", "-skip-validation"} {
		cases = append(cases, testCase{name: "conflict " + flag, args: []string{"-validate-totals", flag, valid}, code: 1, want: []string{"cannot be combined"}})
	}
	cases = append(cases,
		testCase{name: "multiple valid", args: []string{"-validate-totals", valid, valid}},
		testCase{name: "valid then invalid", args: []string{"-validate-totals", valid, invalid}, code: 1, want: []string{invalid}},
		testCase{name: "invalid then valid", args: []string{"-validate-totals", invalid, valid}, code: 1, want: []string{invalid}},
		testCase{name: "multiple errors", args: []string{"-validate-totals", invalid, missing}, code: 1, want: []string{invalid, missing}},
	)

	batchChanges := []struct {
		field  string
		change func(*ach.BatchControl)
	}{
		{"EntryAddendaCount", func(c *ach.BatchControl) { c.EntryAddendaCount++ }},
		{"EntryHash", func(c *ach.BatchControl) { c.EntryHash++ }},
		{"TotalDebitEntryDollarAmount", func(c *ach.BatchControl) { c.TotalDebitEntryDollarAmount++ }},
		{"TotalCreditEntryDollarAmount", func(c *ach.BatchControl) { c.TotalCreditEntryDollarAmount++ }},
	}
	for _, change := range batchChanges {
		data := rewriteACHRecord(t, ppd, '8', func(record string) string {
			var control ach.BatchControl
			control.Parse(record)
			change.change(&control)
			return control.String()
		})
		path := write("batch-"+change.field+".ach", data)
		cases = append(cases, testCase{name: "batch " + change.field, args: []string{"-validate-totals", path}, code: 1, want: []string{path, change.field}})
	}
	// Keep file totals consistent with the corrupted batch, so only checking
	// file aggregates cannot pass the nested batch-to-entry regression.
	nested := rewriteACHRecord(t, ppd, '8', func(record string) string {
		var control ach.BatchControl
		control.Parse(record)
		control.TotalDebitEntryDollarAmount++
		return control.String()
	})
	nested = rewriteACHRecord(t, nested, '9', func(record string) string {
		var control ach.FileControl
		control.Parse(record)
		control.TotalDebitEntryDollarAmountInFile++
		return control.String()
	})
	nestedPath := write("nested.ach", nested)
	opts := &ach.ValidateOpts{BypassBatchValidation: true}
	parsed, err := readACHFile(nested, opts)
	require.NoError(t, err)
	require.ErrorContains(t, parsed.ValidateTotals(), "TotalDebitEntryDollarAmount")
	cases = append(cases, testCase{name: "nested totals with reader bypass", args: []string{"-validate-totals", "-validate", bypassBatch, nestedPath}, code: 1, want: []string{nestedPath, "TotalDebitEntryDollarAmount"}})

	jsonData := fixture("test", "testdata", "ppd-valid.json")
	var document map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(jsonData, &document))
	document["fileControl"] = json.RawMessage(`{"batchCount":999}`)
	badJSON, err := json.Marshal(document)
	require.NoError(t, err)
	for _, input := range []struct {
		name string
		data []byte
	}{{"valid-json.json", jsonData}, {"invalid-totals.json", badJSON}, {"json-named-ach.ach", jsonData}, {"malformed-json.ach", []byte("{bad")}} {
		path := write(input.name, input.data)
		cases = append(cases, testCase{name: input.name, args: []string{"-validate-totals", path}, code: 1, want: []string{path}})
	}
	for _, input := range []struct {
		name string
		data []byte
	}{{"LF.ach", bytes.ReplaceAll(ppd, []byte("\r\n"), []byte("\n"))}, {"CRLF.ach", bytes.ReplaceAll(bytes.ReplaceAll(ppd, []byte("\r\n"), []byte("\n")), []byte("\n"), []byte("\r\n"))}} {
		cases = append(cases, testCase{name: input.name, args: []string{"-validate-totals", write(input.name, input.data)}})
	}

	adv := fixture("test", "ach-adv-read", "adv-read.ach")
	badADV := rewriteACHRecord(t, adv, '9', func(record string) string {
		var control ach.ADVFileControl
		control.Parse(record)
		control.BatchCount++
		return control.String()
	})
	cases = append(cases,
		testCase{name: "ADV", args: []string{"-validate-totals", write("adv.ach", adv)}},
		testCase{name: "ADV BatchCount", args: []string{"-validate-totals", write("bad-adv.ach", badADV)}, code: 1, want: []string{"BatchCount"}},
	)
	mixed, err := readACHFile(ppd, nil)
	require.NoError(t, err)
	iat, err := readACHFile(fixture("test", "ach-iat-read", "iat-credit.ach"), nil)
	require.NoError(t, err)
	for _, batch := range iat.IATBatches {
		mixed.AddIATBatch(batch)
	}
	require.NoError(t, mixed.Create())
	var mixedBytes bytes.Buffer
	require.NoError(t, ach.NewWriter(&mixedBytes).Write(mixed))
	badMixed := rewriteACHRecord(t, mixedBytes.Bytes(), '9', func(record string) string {
		var control ach.FileControl
		control.Parse(record)
		control.BatchCount++
		return control.String()
	})
	cases = append(cases,
		testCase{name: "mixed PPD IAT", args: []string{"-validate-totals", write("mixed.ach", mixedBytes.Bytes())}},
		testCase{name: "mixed PPD IAT BatchCount", args: []string{"-validate-totals", write("bad-mixed.ach", badMixed)}, code: 1, want: []string{"BatchCount"}},
	)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			before, err := os.ReadDir(dir)
			require.NoError(t, err)
			inputs := make(map[string][]byte)
			for _, arg := range tc.args {
				if data, err := os.ReadFile(arg); err == nil {
					inputs[arg] = data
				}
			}
			cmd := exec.CommandContext(t.Context(), bin, tc.args...)
			output, err := cmd.CombinedOutput()
			if tc.code == 0 {
				require.NoError(t, err, "%s", output)
			} else {
				var exitErr *exec.ExitError
				require.True(t, errors.As(err, &exitErr), "expected an exited CLI, got %v: %s", err, output)
				require.Equal(t, tc.code, exitErr.ExitCode(), "%s", output)
			}
			for _, want := range tc.want {
				require.Contains(t, string(output), want)
			}
			if len(tc.want) == 0 && tc.code == 0 {
				require.Empty(t, string(output))
			}
			if strings.Contains(strings.Join(tc.args, " "), "-validate-totals") {
				require.NotContains(t, string(output), "Describing")
			}
			for path, data := range inputs {
				after, err := os.ReadFile(path)
				require.NoError(t, err)
				require.Equal(t, data, after, "input changed: %s", path)
			}
			after, err := os.ReadDir(dir)
			require.NoError(t, err)
			require.Equal(t, before, after, "validation created or removed a file")
		})
	}
}

func rewriteACHRecord(t *testing.T, data []byte, recordType byte, rewrite func(string) string) []byte {
	t.Helper()
	for start := 0; start < len(data); {
		length := bytes.IndexByte(data[start:], '\n')
		if length < 0 {
			length = len(data) - start
		}
		end := start + length
		if end > start && data[end-1] == '\r' {
			end--
		}
		record := data[start:end]
		if len(record) > 1 && record[0] == recordType && !bytes.HasPrefix(record, []byte("99")) {
			// The reader pads omitted trailing spaces before parsing records.
			// Keep the fixture's original width after changing its fields.
			padded := string(record)
			if len(record) < 94 {
				padded += strings.Repeat(" ", 94-len(record))
			}
			changed := rewrite(padded)
			require.Len(t, changed, 94)
			out := append([]byte(nil), data[:start]...)
			out = append(out, changed[:len(record)]...)
			return append(out, data[end:]...)
		}
		start += length + 1
	}
	t.Fatalf("missing record type %c in fixture", recordType)
	return nil
}
