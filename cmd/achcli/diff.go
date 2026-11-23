// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/describe/mask"

	"github.com/juju/ansiterm"
)

func diffFiles(paths []string, validateOpts *ach.ValidateOpts, maskOpts mask.Options) error {
	if len(paths) != 2 {
		return fmt.Errorf("expected 2 files, but got %d", len(paths))
	}
	f1, f2, err := readTwoFiles(paths, validateOpts)
	if err != nil {
		return err
	}
	return writeDiff(os.Stdout, paths, f1, f2, stdoutIsTTY(), maskOpts)
}

func readTwoFiles(paths []string, validateOpts *ach.ValidateOpts) (*ach.File, *ach.File, error) {
	f1, err := readIncomingFile(paths[0], validateOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("problem reading %s: %v", paths[0], err)
	}
	f2, err := readIncomingFile(paths[1], validateOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("problem reading %s: %v", paths[1], err)
	}
	return f1, f2, nil
}

func stdoutIsTTY() bool {
	fi, err := os.Stdout.Stat()
	return err == nil && fi.Mode()&os.ModeCharDevice != 0
}

func writeDiff(out io.Writer, paths []string, f1, f2 *ach.File, color bool, maskOpts mask.Options) error {
	var buf bytes.Buffer
	d := &differ{out: &buf, color: color, left: paths[0], right: paths[1], needBlank: true, mask: maskOpts}
	d.diff(f1, f2)
	if buf.Len() == 0 {
		return nil
	}
	fmt.Fprintf(out, "--- %s\n+++ %s\n", paths[0], paths[1])
	_, err := io.Copy(out, &buf)
	return err
}

type col struct {
	name  string
	left  string
	right string
}

type idxPair struct {
	a int
	b int
}

type differ struct {
	out       io.Writer
	color     bool
	left      string
	right     string
	needBlank bool

	// mask hides sensitive values when they are printed. Fields are still
	// compared on their raw values, so a difference is never hidden.
	mask mask.Options
	// sec is the SEC code of the batch being compared, used to mask
	// ENR and DNE payment information in Addenda05 records.
	sec string
}

// child returns a differ writing to out with the same settings as d.
func (d *differ) child(out io.Writer) *differ {
	return &differ{out: out, color: d.color, left: d.left, right: d.right, mask: d.mask, sec: d.sec}
}

func (d *differ) diff(f1, f2 *ach.File) {
	if f1 == nil {
		f1 = &ach.File{}
	}
	if f2 == nil {
		f2 = &ach.File{}
	}
	d.section("FileHeader", fileHeaderCols(f1.Header, f2.Header))
	d.diffBatches(f1.Batches, f2.Batches)
	d.diffIATBatches(f1.IATBatches, f2.IATBatches)
	d.section("FileControl", fileControlCols(f1.Control, f2.Control))
	d.section("ADVFileControl", advFileControlCols(f1.ADVControl, f2.ADVControl))
}

func (d *differ) heading(s string) {
	if d.needBlank {
		fmt.Fprintln(d.out)
	}
	fmt.Fprintln(d.out, s)
	d.needBlank = true
}

func (d *differ) section(title string, cols []col) {
	if !hasDiff(cols) {
		return
	}
	d.heading(title)
	d.tableBoth("  ", d.masked(differing(cols)))
}

func indentOf(title string) string {
	i := 0
	for i < len(title) && title[i] == ' ' {
		i++
	}
	return title[:i] + "  "
}

func (d *differ) nestedSide(title string, leftNil, rightNil bool, cols []col) {
	if leftNil && rightNil {
		return
	}
	if !hasDiff(cols) && leftNil == rightNil {
		return
	}
	fmt.Fprintln(d.out, title)
	d.oneOrBoth(indentOf(title), leftNil, rightNil, cols)
}

func hasDiff(cols []col) bool {
	for _, c := range cols {
		if c.left != c.right {
			return true
		}
	}
	return false
}

func differing(cols []col) []col {
	out := make([]col, 0, len(cols))
	for _, c := range cols {
		if c.left != c.right {
			out = append(out, c)
		}
	}
	return out
}

// masked replaces the values of sensitive columns with their masked form,
// following the same rules as achcli -mask when describing a file.
func (d *differ) masked(cols []col) []col {
	for i := range cols {
		fn := d.maskFunc(cols[i].name)
		if fn == nil {
			continue
		}
		cols[i].left = maskValue(fn, cols[i].left)
		cols[i].right = maskValue(fn, cols[i].right)
	}
	return cols
}

func (d *differ) maskFunc(name string) func(string) string {
	switch name {
	case "AccountNumber":
		if d.mask.MaskAccountNumbers {
			return mask.Number
		}
	case "Name":
		if d.mask.MaskNames {
			return mask.Name
		}
	case "IdentificationNumber":
		if d.mask.MaskIdentification {
			return mask.Number
		}
	case "CorrectedData":
		if d.mask.MaskCorrectedData {
			return mask.Number
		}
	case "PaymentRelatedInformation":
		return d.maskPaymentInformation()
	}
	return nil
}

// maskPaymentInformation masks the account numbers, names and SSNs that ENR and
// DNE batches carry in Addenda05 records.
func (d *differ) maskPaymentInformation() func(string) string {
	switch d.sec {
	case ach.ENR:
		if !d.mask.MaskNames && !d.mask.MaskAccountNumbers {
			return nil
		}
		return func(v string) string {
			info, err := ach.ParseENRPaymentInformation(&ach.Addenda05{PaymentRelatedInformation: v})
			if err != nil || info == nil {
				return v
			}
			if d.mask.MaskNames {
				info.IndividualName = mask.Name(info.IndividualName)
			}
			if d.mask.MaskAccountNumbers {
				info.IndividualIdentification = mask.Number(info.IndividualIdentification)
				info.DFIAccountNumber = mask.Number(info.DFIAccountNumber)
			}
			return info.String()
		}
	case ach.DNE:
		if !d.mask.MaskNames && !d.mask.MaskAccountNumbers {
			return nil
		}
		return func(v string) string {
			info, err := ach.ParseDNEPaymentInformation(&ach.Addenda05{PaymentRelatedInformation: v})
			if err != nil || info == nil {
				return v
			}
			info.CustomerSSN = mask.Number(info.CustomerSSN)
			return info.String()
		}
	}
	return nil
}

// maskValue leaves empty values empty, so a field that is only set on one
// side does not look set on both.
func maskValue(fn func(string) string, v string) string {
	if v == "" {
		return v
	}
	return fn(v)
}

func names(cols []col) []string {
	out := make([]string, len(cols))
	for i, c := range cols {
		out[i] = c.name
	}
	return out
}

func sides(cols []col, right bool) []string {
	out := make([]string, len(cols))
	for i, c := range cols {
		if right {
			out[i] = c.right
		} else {
			out[i] = c.left
		}
	}
	return out
}

func (d *differ) oneOrBoth(indent string, leftNil, rightNil bool, cols []col) {
	cols = d.masked(differing(cols))
	if len(cols) == 0 {
		return
	}
	switch {
	case leftNil && !rightNil:
		d.tableOne(indent, true, cols)
	case rightNil && !leftNil:
		d.tableOne(indent, false, cols)
	default:
		d.tableBoth(indent, cols)
	}
}

func (d *differ) tableBoth(indent string, cols []col) {
	if len(cols) == 0 {
		return
	}
	tw := ansiterm.NewTabWriter(d.out, 0, 0, 2, ' ', 0)
	d.writeRow(tw, indent+"  ", names(cols), false, ansiterm.Default)
	d.writeRow(tw, indent+"- ", sides(cols, false), d.color, ansiterm.Red)
	d.writeRow(tw, indent+"+ ", sides(cols, true), d.color, ansiterm.Green)
	tw.Flush()
}

func (d *differ) tableOne(indent string, right bool, cols []col) {
	prefix := "- "
	color := ansiterm.Red
	if right {
		prefix = "+ "
		color = ansiterm.Green
	}
	tw := ansiterm.NewTabWriter(d.out, 0, 0, 2, ' ', 0)
	d.writeRow(tw, indent+"  ", names(cols), false, ansiterm.Default)
	d.writeRow(tw, indent+prefix, sides(cols, right), d.color, color)
	tw.Flush()
}

func (d *differ) writeRow(w io.Writer, prefix string, vals []string, colored bool, color ansiterm.Color) {
	fmt.Fprint(w, prefix)
	for i, v := range vals {
		if i > 0 {
			fmt.Fprint(w, "\t")
		}
		if colored {
			aw := ansiterm.NewWriter(w)
			aw.SetColorCapable(true)
			ansiterm.Foreground(color).Fprint(aw, v)
		} else {
			fmt.Fprint(w, v)
		}
	}
	fmt.Fprintln(w)
}

func align(na, nb int, keyA, keyB func(int) string) []idxPair {
	usedA := make([]bool, na)
	usedB := make([]bool, nb)
	shared := zipSharedKeys(na, nb, usedA, usedB, keyA, keyB)
	empty := zipKey(na, nb, usedA, usedB, keyA, keyB, "")
	rest := unusedOneSided(na, nb, usedA, usedB)
	out := make([]idxPair, 0, len(shared)+len(empty)+len(rest))
	out = append(out, shared...)
	out = append(out, empty...)
	return append(out, rest...)
}

func zipSharedKeys(na, nb int, usedA, usedB []bool, keyA, keyB func(int) string) []idxPair {
	out := make([]idxPair, 0, na)
	seen := make(map[string]bool)
	for i := 0; i < na; i++ {
		if usedA[i] {
			continue
		}
		k := keyA(i)
		if k == "" || seen[k] {
			continue
		}
		seen[k] = true
		out = append(out, zipKey(na, nb, usedA, usedB, keyA, keyB, k)...)
	}
	return out
}

func zipKey(na, nb int, usedA, usedB []bool, keyA, keyB func(int) string, k string) []idxPair {
	as := unusedWithKey(na, usedA, keyA, k)
	bs := unusedWithKey(nb, usedB, keyB, k)
	n := min(len(as), len(bs))
	out := make([]idxPair, n)
	for i := 0; i < n; i++ {
		usedA[as[i]] = true
		usedB[bs[i]] = true
		out[i] = idxPair{a: as[i], b: bs[i]}
	}
	return out
}

func unusedWithKey(n int, used []bool, key func(int) string, want string) []int {
	out := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if !used[i] && key(i) == want {
			out = append(out, i)
		}
	}
	return out
}

func unusedOneSided(na, nb int, usedA, usedB []bool) []idxPair {
	out := make([]idxPair, 0, na+nb)
	for i := 0; i < na; i++ {
		if !usedA[i] {
			out = append(out, idxPair{a: i, b: -1})
		}
	}
	for j := 0; j < nb; j++ {
		if !usedB[j] {
			out = append(out, idxPair{a: -1, b: j})
		}
	}
	return out
}

func trim(v string) string { return strings.TrimSpace(v) }

func batcherPresent(b ach.Batcher) bool {
	if b == nil {
		return false
	}
	v := reflect.ValueOf(b)
	return v.Kind() != reflect.Ptr || !v.IsNil()
}

func num(i int) string { return strconv.Itoa(i) }

func fileHeaderCols(a, b ach.FileHeader) []col {
	return []col{
		{"Origin", trim(a.ImmediateOrigin), trim(b.ImmediateOrigin)},
		{"OriginName", trim(a.ImmediateOriginName), trim(b.ImmediateOriginName)},
		{"Destination", trim(a.ImmediateDestination), trim(b.ImmediateDestination)},
		{"DestinationName", trim(a.ImmediateDestinationName), trim(b.ImmediateDestinationName)},
		{"FileCreationDate", trim(a.FileCreationDate), trim(b.FileCreationDate)},
		{"FileCreationTime", trim(a.FileCreationTime), trim(b.FileCreationTime)},
		{"FileIDModifier", trim(a.FileIDModifier), trim(b.FileIDModifier)},
		{"ReferenceCode", trim(a.ReferenceCode), trim(b.ReferenceCode)},
	}
}

func fileControlCols(a, b ach.FileControl) []col {
	return []col{
		{"BatchCount", num(a.BatchCount), num(b.BatchCount)},
		{"BlockCount", num(a.BlockCount), num(b.BlockCount)},
		{"EntryAddendaCount", num(a.EntryAddendaCount), num(b.EntryAddendaCount)},
		{"EntryHash", num(a.EntryHash), num(b.EntryHash)},
		{"TotalDebitAmount", num(a.TotalDebitEntryDollarAmountInFile), num(b.TotalDebitEntryDollarAmountInFile)},
		{"TotalCreditAmount", num(a.TotalCreditEntryDollarAmountInFile), num(b.TotalCreditEntryDollarAmountInFile)},
	}
}

func advFileControlCols(a, b ach.ADVFileControl) []col {
	return []col{
		{"BatchCount", num(a.BatchCount), num(b.BatchCount)},
		{"BlockCount", num(a.BlockCount), num(b.BlockCount)},
		{"EntryAddendaCount", num(a.EntryAddendaCount), num(b.EntryAddendaCount)},
		{"EntryHash", num(a.EntryHash), num(b.EntryHash)},
		{"TotalDebitAmount", num(a.TotalDebitEntryDollarAmountInFile), num(b.TotalDebitEntryDollarAmountInFile)},
		{"TotalCreditAmount", num(a.TotalCreditEntryDollarAmountInFile), num(b.TotalCreditEntryDollarAmountInFile)},
	}
}
