// Copyright 2026 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/describe/mask"
	"github.com/stretchr/testify/require"
)

func readPPD(t *testing.T) *ach.File {
	t.Helper()
	file, err := ach.ReadFile(testdataPath("ppd-debit.ach"))
	require.NoError(t, err)
	return file
}

func diffText(t *testing.T, left, right *ach.File) string {
	t.Helper()
	var buf bytes.Buffer
	require.NoError(t, writeDiff(&buf, []string{"left.ach", "right.ach"}, left, right, false, mask.Options{}))
	return buf.String()
}

func TestDiffFiles_identical(t *testing.T) {
	out := diffText(t, readPPD(t), readPPD(t))
	require.Empty(t, out)
}

func TestDiffFiles_entryAmount(t *testing.T) {
	left := readPPD(t)
	right := readPPD(t)
	right.Batches[0].GetEntries()[0].Amount = 9999

	out := diffText(t, left, right)
	require.Contains(t, out, "--- left.ach")
	require.Contains(t, out, "+++ right.ach")
	require.Contains(t, out, "Entry 121042880000001")
	require.Contains(t, out, "Amount")
	require.Contains(t, out, "100000000")
	require.Contains(t, out, "9999")
	require.NotContains(t, out, "FileHeader")
}

func TestDiffFiles_fileHeaderAndModifier(t *testing.T) {
	left := readPPD(t)
	right := readPPD(t)
	right.Header.ImmediateOrigin = "999999999"
	right.Header.FileIDModifier = "B"
	right.Header.FileCreationDate = "261116"
	right.Header.FileCreationTime = "0815"

	out := diffText(t, left, right)
	require.Contains(t, out, "FileHeader")
	require.Contains(t, out, "Origin")
	require.Contains(t, out, "121042882")
	require.Contains(t, out, "999999999")
	require.Contains(t, out, "FileIDModifier")
	require.Contains(t, out, "FileCreationDate")
	require.Contains(t, out, "FileCreationTime")
	require.Contains(t, out, "261116")
	require.Contains(t, out, "0815")
}

func TestDiffFiles_individualNameAndAccount(t *testing.T) {
	left := readPPD(t)
	right := readPPD(t)
	ed := right.Batches[0].GetEntries()[0]
	ed.IndividualName = "ALICE SYNTHETIC"
	ed.DFIAccountNumber = "4401234567"

	out := diffText(t, left, right)
	require.Contains(t, out, "Name")
	require.Contains(t, out, "Receiver Account Name")
	require.Contains(t, out, "ALICE SYNTHETIC")
	require.Contains(t, out, "AccountNumber")
	require.Contains(t, out, "12345678")
	require.Contains(t, out, "4401234567")
}

func TestDiffFiles_emptyAddenda02(t *testing.T) {
	left := readPPD(t)
	right := readPPD(t)
	right.Batches[0].GetEntries()[0].Addenda02 = ach.NewAddenda02()

	out := diffText(t, left, right)
	require.Contains(t, out, "Addenda02")
}

func TestDiffFiles_addenda05(t *testing.T) {
	left := readPPD(t)
	right := readPPD(t)
	ed := right.Batches[0].GetEntries()[0]
	ed.AddendaRecordIndicator = 1
	ad := ach.NewAddenda05()
	ad.PaymentRelatedInformation = "invoice 99"
	ad.SequenceNumber = 1
	ed.Addenda05 = []*ach.Addenda05{ad}

	out := diffText(t, left, right)
	require.Contains(t, out, "Addenda05")
	require.Contains(t, out, "invoice 99")
	require.Contains(t, out, "AddendaRecordIndicator")
}

func TestDiffFiles_addedEntry(t *testing.T) {
	left := readPPD(t)
	right := readPPD(t)
	src := right.Batches[0].GetEntries()[0]
	added := *src
	added.TraceNumber = "121042880000002"
	added.Amount = 50
	right.Batches[0].AddEntry(&added)

	out := diffText(t, left, right)
	require.Contains(t, out, "Entry 121042880000002")
	require.Contains(t, out, "50")
	require.NotContains(t, out, "Entry 121042880000001")
}

func TestDiffFiles_replacedEntryKeepsSharedTrace(t *testing.T) {
	left := readPPD(t)
	right := readPPD(t)

	shared := *left.Batches[0].GetEntries()[0]
	shared.TraceNumber = "121042880000002"
	shared.Amount = 200
	left.Batches[0].AddEntry(&shared)

	replaced := right.Batches[0].GetEntries()[0]
	replaced.TraceNumber = "121042880000003"
	replaced.Amount = 50
	replaced.IndividualName = "BOB REPLACEMENT"
	rightShared := shared
	right.Batches[0].AddEntry(&rightShared)

	out := diffText(t, left, right)
	require.Contains(t, out, "Entry 121042880000001")
	require.Contains(t, out, "Entry 121042880000003")
	require.NotContains(t, out, "Entry 121042880000002")

	i1 := strings.Index(out, "Entry 121042880000001")
	i3 := strings.Index(out, "Entry 121042880000003")
	require.GreaterOrEqual(t, i1, 0)
	require.GreaterOrEqual(t, i3, 0)
	require.Less(t, i1, i3)
	leftOnly := out[i1:i3]
	rightOnly := out[i3:]
	require.Contains(t, leftOnly, "\n    - ")
	require.NotContains(t, leftOnly, "\n    + ")
	require.Contains(t, rightOnly, "\n    + ")
	require.NotContains(t, rightOnly, "\n    - ")
	require.Contains(t, leftOnly, "Receiver Account Name")
	require.Contains(t, rightOnly, "BOB REPLACEMENT")
}

func TestDiffFiles_ppdWithStrayADVEntries(t *testing.T) {
	left := readPPD(t)
	right := readPPD(t)
	right.Batches[0].GetEntries()[0].Amount = 9999
	left.Batches[0].AddADVEntry(ach.NewADVEntryDetail())

	out := diffText(t, left, right)
	require.Contains(t, out, "Batch 0000001 PPD")
	require.Contains(t, out, "Amount")
	require.Contains(t, out, "100000000")
	require.Contains(t, out, "9999")
	require.Contains(t, out, "\n    - ")
	require.Contains(t, out, "\n    + ")
	require.NotContains(t, out, "only in right.ach")
}

func TestDiffFiles_batchHeader(t *testing.T) {
	left := readPPD(t)
	right := readPPD(t)
	right.Batches[0].GetHeader().CompanyEntryDescription = "BONUS"

	out := diffText(t, left, right)
	require.Contains(t, out, "Batch 0000001 PPD")
	require.Contains(t, out, "EntryDescription")
	require.Contains(t, out, "REG.SALARY")
	require.Contains(t, out, "BONUS")
}

func TestDiffFiles_typedNilBatcher(t *testing.T) {
	left := readPPD(t)
	var typed *ach.BatchPPD
	left.Batches = []ach.Batcher{typed}
	require.NotPanics(t, func() {
		out := diffText(t, left, &ach.File{})
		require.NotContains(t, out, "PPD")
		require.NotContains(t, out, "only in")
	})
}

func TestDiffFiles_ppdVsIAT(t *testing.T) {
	left := readPPD(t)
	right, err := ach.ReadFile(testdataPath("iat-debit.ach"))
	require.NoError(t, err)
	out := diffText(t, left, right)
	require.Contains(t, out, "only in left.ach")
	require.Contains(t, out, "IAT Batch")
}

func TestDiffFiles_nilBatcher(t *testing.T) {
	left := readPPD(t)
	right := readPPD(t)
	right.Batches = []ach.Batcher{nil, right.Batches[0]}

	require.NotPanics(t, func() {
		out := diffText(t, left, right)
		require.Empty(t, out)
	})

	onlyNil := readPPD(t)
	onlyNil.Batches = []ach.Batcher{nil}
	out := diffText(t, left, onlyNil)
	require.Contains(t, out, "only in left.ach")
	require.NotContains(t, out, "only in right.ach")
}

func TestDiffFiles_mixedADVAndPPD(t *testing.T) {
	left := readPPD(t)
	right, err := ach.ReadFile(testdataPath("adv.ach"))
	require.NoError(t, err)
	require.NotEmpty(t, right.Batches)
	right.Batches[0].GetHeader().BatchNumber = left.Batches[0].GetHeader().BatchNumber

	out := diffText(t, left, right)
	require.Contains(t, out, "only in left.ach")
	require.Contains(t, out, "only in right.ach")
	require.Contains(t, out, "PPD")
	require.Contains(t, out, "ADV")
	require.Contains(t, out, "Receiver Account Name")
}

func TestDiffFiles_addedBatch(t *testing.T) {
	left := readPPD(t)
	right := readPPD(t)
	extra := readPPD(t)
	extra.Batches[0].GetHeader().BatchNumber = 2
	right.AddBatch(extra.Batches[0])

	out := diffText(t, left, right)
	require.Contains(t, out, "Batch 0000002 PPD only in right.ach")
	require.Contains(t, out, "Entry 121042880000001")
	require.Contains(t, out, "BatchControl")
	require.NotContains(t, out, "only in left.ach")
	ctrl := out[strings.Index(out, "BatchControl"):]
	require.Contains(t, ctrl, "\n    + ")
	require.NotContains(t, ctrl, "\n    - ")
}

func TestDiffFiles_iatSettlementDate(t *testing.T) {
	left, err := ach.ReadFile(testdataPath("iat-debit.ach"))
	require.NoError(t, err)
	right, err := ach.ReadFile(testdataPath("iat-debit.ach"))
	require.NoError(t, err)
	right.IATBatches[0].GetHeader().SettlementDate = "123"

	out := diffText(t, left, right)
	require.Contains(t, out, "IAT Batch")
	require.Contains(t, out, "SettlementDate")
	require.Contains(t, out, "123")
}

func TestDiffFiles_advAddenda99(t *testing.T) {
	raw := testdataBytes(t, "adv-return.json")
	left, err := ach.FileFromJSON(raw)
	require.NoError(t, err)
	right, err := ach.FileFromJSON(raw)
	require.NoError(t, err)
	require.NotEmpty(t, right.Batches[0].GetADVEntries())
	require.NotNil(t, right.Batches[0].GetADVEntries()[0].Addenda99)
	right.Batches[0].GetADVEntries()[0].Addenda99.ReturnCode = "R01"

	out := diffText(t, left, right)
	require.Contains(t, out, "ADV Entry")
	require.Contains(t, out, "Addenda99")
	require.Contains(t, out, "R14")
	require.Contains(t, out, "R01")
}

func TestDiffFiles_advJulianDay(t *testing.T) {
	left, err := ach.ReadFile(testdataPath("adv.ach"))
	require.NoError(t, err)
	right, err := ach.ReadFile(testdataPath("adv.ach"))
	require.NoError(t, err)
	require.NotEmpty(t, right.Batches[0].GetADVEntries())
	right.Batches[0].GetADVEntries()[0].JulianDay = 200

	out := diffText(t, left, right)
	require.Contains(t, out, "ADV Entry")
	require.Contains(t, out, "JulianDay")
	require.Contains(t, out, "200")
}

func TestWriteDiff_color(t *testing.T) {
	left := readPPD(t)
	right := readPPD(t)
	right.Batches[0].GetEntries()[0].Amount = 9999

	var buf bytes.Buffer
	require.NoError(t, writeDiff(&buf, []string{"left.ach", "right.ach"}, left, right, true, mask.Options{}))
	out := buf.String()
	require.Contains(t, out, "\x1b[31m")
	require.Contains(t, out, "\x1b[32m")

	plain := diffText(t, left, right)
	require.NotContains(t, plain, "\x1b[")
}

func TestDiffFiles_iatAmount(t *testing.T) {
	left, err := ach.ReadFile(testdataPath("iat-debit.ach"))
	require.NoError(t, err)
	right, err := ach.ReadFile(testdataPath("iat-debit.ach"))
	require.NoError(t, err)
	right.IATBatches[0].GetEntries()[0].Amount = 42

	out := diffText(t, left, right)
	require.Contains(t, out, "IAT Batch")
	require.Contains(t, out, "Amount")
	require.Contains(t, out, "42")
}

func TestDiffFiles_fileControlTotals(t *testing.T) {
	left := readPPD(t)
	right := readPPD(t)
	right.Batches[0].GetEntries()[0].Amount = 9999
	right.Batches[0].GetControl().TotalDebitEntryDollarAmount = 9999
	right.Control.TotalDebitEntryDollarAmountInFile = 9999

	out := diffText(t, left, right)
	require.Contains(t, out, "Amount")
	require.Contains(t, out, "BatchControl")
	require.Contains(t, out, "TotalDebits")
	require.Contains(t, out, "FileControl")
	require.Contains(t, out, "TotalDebitAmount")
}

func TestAlign_uniqueKeysOutOfOrder(t *testing.T) {
	a := []string{"1", "2"}
	b := []string{"2", "1"}
	got := align(len(a), len(b), func(i int) string { return a[i] }, func(i int) string { return b[i] })
	require.Equal(t, []idxPair{{a: 0, b: 1}, {a: 1, b: 0}}, got)
}

func TestAlign_extraOnRight(t *testing.T) {
	a := []string{"1"}
	b := []string{"1", "2"}
	got := align(len(a), len(b), func(i int) string { return a[i] }, func(i int) string { return b[i] })
	require.Equal(t, []idxPair{{a: 0, b: 0}, {a: -1, b: 1}}, got)
}

func TestAlign_duplicateKeysPositional(t *testing.T) {
	a := []string{"1", "1"}
	b := []string{"1", "1"}
	got := align(len(a), len(b), func(i int) string { return a[i] }, func(i int) string { return b[i] })
	require.Equal(t, []idxPair{{a: 0, b: 0}, {a: 1, b: 1}}, got)
}

func TestAlign_differentLeftoverKeysAreOneSided(t *testing.T) {
	a := []string{"1", "2"}
	b := []string{"2", "3"}
	got := align(len(a), len(b), func(i int) string { return a[i] }, func(i int) string { return b[i] })
	require.Equal(t, []idxPair{{a: 1, b: 0}, {a: 0, b: -1}, {a: -1, b: 1}}, got)
}

func pad94(s string) string {
	n := len([]rune(s))
	if n >= 94 {
		return string([]rune(s)[:94])
	}
	return s + strings.Repeat(" ", 94-n)
}

func joinACH(lines ...string) string {
	for i := range lines {
		lines[i] = pad94(lines[i])
	}
	return strings.Join(lines, "\n") + "\n"
}

func TestDiffFiles_issue1931Sample(t *testing.T) {
	left := joinACH(
		"101 231380104 1210428822611160815A094101FED RESERVE SYNTH      NORTHWIND PAYROLL",
		"5220NORTHWIND PAYROL                    1987654321PPDPAYROLL         261117   1121042880000001",
		"6222313801044401234567       0000002543EMP0001        ALICE SYNTHETIC         0121042880000001",
		"822000000100231380100000000000000000000025431987654321                         121042880000001",
		"9000001000001000000010023138010000000000000000000002543",
	)
	right := strings.ReplaceAll(left, "0000002543", "0000009999")

	parse := func(raw string) *ach.File {
		t.Helper()
		r := ach.NewReader(strings.NewReader(raw))
		file, err := r.Read()
		require.NoError(t, err)
		return &file
	}

	out := diffText(t, parse(left), parse(right))
	require.Contains(t, out, "Amount")
	require.Contains(t, out, "2543")
	require.Contains(t, out, "9999")
	require.Contains(t, out, "Entry 121042880000001")
}

func TestDiffCLI_issue1931(t *testing.T) {
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

	left := joinACH(
		"101 231380104 1210428822611160815A094101FED RESERVE SYNTH      NORTHWIND PAYROLL",
		"5220NORTHWIND PAYROL                    1987654321PPDPAYROLL         261117   1121042880000001",
		"6222313801044401234567       0000002543EMP0001        ALICE SYNTHETIC         0121042880000001",
		"822000000100231380100000000000000000000025431987654321                         121042880000001",
		"9000001000001000000010023138010000000000000000000002543",
	)
	right := strings.ReplaceAll(left, "0000002543", "0000009999")
	dir := t.TempDir()
	a := filepath.Join(dir, "a.ach")
	b := filepath.Join(dir, "b.ach")
	require.NoError(t, os.WriteFile(a, []byte(left), 0600))
	require.NoError(t, os.WriteFile(b, []byte(right), 0600))

	cmd := exec.CommandContext(t.Context(), bin, "-diff", a, b)
	out, err := cmd.CombinedOutput()
	require.NoError(t, err, "%s", out)
	text := string(out)
	require.Contains(t, text, "Amount")
	require.Contains(t, text, "2543")
	require.Contains(t, text, "9999")

	same := exec.CommandContext(t.Context(), bin, "-diff", a, a)
	sameOut, err := same.CombinedOutput()
	require.NoError(t, err, "%s", sameOut)
	require.Empty(t, string(sameOut))
}

func TestDiffFiles_wrongPathCount(t *testing.T) {
	err := diffFiles([]string{"only.ach"}, nil, mask.Options{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "expected 2 files")
}

func TestDiffFiles_missingFile(t *testing.T) {
	err := diffFiles([]string{"no-such-left.ach", "no-such-right.ach"}, nil, mask.Options{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "problem reading")
}

func diffTextMasked(t *testing.T, left, right *ach.File, opts mask.Options) string {
	t.Helper()
	var buf bytes.Buffer
	require.NoError(t, writeDiff(&buf, []string{"left.ach", "right.ach"}, left, right, false, opts))
	return buf.String()
}

func TestDiffFiles_mask(t *testing.T) {
	all := mask.Options{MaskAccountNumbers: true, MaskNames: true, MaskCorrectedData: true}

	cases := []struct {
		name     string
		opts     mask.Options
		edit     func(left, right *ach.File)
		contains []string
		hidden   []string
	}{
		{
			name: "entry name and account",
			opts: all,
			edit: func(_, right *ach.File) {
				ed := right.Batches[0].GetEntries()[0]
				ed.IndividualName = "ALICE SYNTHETIC"
				ed.DFIAccountNumber = "4401234567"
			},
			contains: []string{"Name", "Re****** Ac***** Na**", "AL*** SY*******", "AccountNumber", "****5678", "******4567"},
			hidden:   []string{"Receiver Account Name", "ALICE SYNTHETIC", "12345678", "4401234567"},
		},
		{
			name: "only names",
			opts: mask.Options{MaskNames: true},
			edit: func(_, right *ach.File) {
				ed := right.Batches[0].GetEntries()[0]
				ed.IndividualName = "ALICE SYNTHETIC"
				ed.DFIAccountNumber = "4401234567"
			},
			contains: []string{"AL*** SY*******", "12345678", "4401234567"},
			hidden:   []string{"ALICE SYNTHETIC"},
		},
		{
			// The raw values differ only in digits the mask hides: the row is
			// still printed so the difference is not lost.
			name: "difference hidden by the mask is still reported",
			opts: all,
			edit: func(left, right *ach.File) {
				left.Batches[0].GetEntries()[0].DFIAccountNumber = "1112345678"
				right.Batches[0].GetEntries()[0].DFIAccountNumber = "9992345678"
			},
			contains: []string{"Entry 121042880000001", "AccountNumber", "******5678"},
			hidden:   []string{"1112345678", "9992345678"},
		},
		{
			name: "corrected data",
			opts: all,
			edit: func(left, right *ach.File) {
				l := ach.NewAddenda98()
				l.CorrectedData = "1918164748"
				left.Batches[0].GetEntries()[0].Addenda98 = l
				r := ach.NewAddenda98()
				r.CorrectedData = "1918169999"
				right.Batches[0].GetEntries()[0].Addenda98 = r
			},
			contains: []string{"Addenda98", "CorrectedData", "******4748", "******9999"},
			hidden:   []string{"1918164748", "1918169999"},
		},
		{
			name: "DNE customer SSN",
			opts: all,
			edit: func(left, right *ach.File) {
				for i, f := range []*ach.File{left, right} {
					f.Batches[0].GetHeader().StandardEntryClassCode = ach.DNE
					ad := ach.NewAddenda05()
					ad.PaymentRelatedInformation = []string{
						`DATE OF DEATH*010218*CUSTOMER SSN*123456789*AMOUNT*12.34\`,
						`DATE OF DEATH*010218*CUSTOMER SSN*987654321*AMOUNT*12.34\`,
					}[i]
					ad.SequenceNumber = 1
					f.Batches[0].GetEntries()[0].Addenda05 = []*ach.Addenda05{ad}
				}
			},
			contains: []string{"Addenda05", "CUSTOMER SSN******6789", "CUSTOMER SSN******4321"},
			hidden:   []string{"123456789", "987654321"},
		},
		{
			name: "ENR account and name",
			opts: all,
			edit: func(left, right *ach.File) {
				for i, f := range []*ach.File{left, right} {
					f.Batches[0].GetHeader().StandardEntryClassCode = ach.ENR
					ad := ach.NewAddenda05()
					ad.PaymentRelatedInformation = []string{
						`22*12200004*3*123987654321*777777777*DOE*JOHN*1\`,
						`22*12200004*3*123987650000*777777777*DOE*JOHN*1\`,
					}[i]
					ad.SequenceNumber = 1
					f.Batches[0].GetEntries()[0].Addenda05 = []*ach.Addenda05{ad}
				}
			},
			contains: []string{"Addenda05", "********4321", "********0000"},
			hidden:   []string{"123987654321", "123987650000", "777777777", "JOHN"},
		},
		{
			name: "IAT account",
			opts: all,
			edit: func(left, right *ach.File) {
				right.IATBatches[0].GetEntries()[0].DFIAccountNumber = "4401234567"
			},
			contains: []string{"IAT Batch", "AccountNumber", "******4567"},
			hidden:   []string{"4401234567"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			name := "ppd-debit.ach"
			if strings.HasPrefix(tc.name, "IAT") {
				name = "iat-debit.ach"
			}
			left, err := ach.ReadFile(testdataPath(name))
			require.NoError(t, err)
			right, err := ach.ReadFile(testdataPath(name))
			require.NoError(t, err)
			tc.edit(left, right)

			out := diffTextMasked(t, left, right, tc.opts)
			for _, s := range tc.contains {
				require.Contains(t, out, s)
			}
			for _, s := range tc.hidden {
				require.NotContains(t, out, s)
			}
		})
	}
}

func TestDiffCLI_mask(t *testing.T) {
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

	left := joinACH(
		"101 231380104 1210428822611160815A094101FED RESERVE SYNTH      NORTHWIND PAYROLL",
		"5220NORTHWIND PAYROL                    1987654321PPDPAYROLL         261117   1121042880000001",
		"6222313801044401234567       0000002543EMP0001        ALICE SYNTHETIC         0121042880000001",
		"822000000100231380100000000000000000000025431987654321                         121042880000001",
		"9000001000001000000010023138010000000000000000000002543",
	)
	right := strings.NewReplacer("4401234567", "4401239999", "ALICE SYNTHETIC", "ALICE CHANGED  ").Replace(left)
	dir := t.TempDir()
	a := filepath.Join(dir, "a.ach")
	b := filepath.Join(dir, "b.ach")
	require.NoError(t, os.WriteFile(a, []byte(left), 0600))
	require.NoError(t, os.WriteFile(b, []byte(right), 0600))

	out, err := exec.CommandContext(t.Context(), bin, "-mask", "-diff", a, b).CombinedOutput()
	require.NoError(t, err, "%s", out)
	text := string(out)
	require.Contains(t, text, "AccountNumber")
	require.Contains(t, text, "******4567")
	require.Contains(t, text, "******9999")
	require.Contains(t, text, "AL*** SY*******")
	require.Contains(t, text, "AL*** CH*****")
	for _, s := range []string{"4401234567", "4401239999", "SYNTHETIC", "CHANGED"} {
		require.NotContains(t, text, s)
	}

	out, err = exec.CommandContext(t.Context(), bin, "-mask.accounts", "-diff", a, b).CombinedOutput()
	require.NoError(t, err, "%s", out)
	text = string(out)
	require.Contains(t, text, "******9999")
	require.Contains(t, text, "ALICE CHANGED")
	require.NotContains(t, text, "4401239999")
}
