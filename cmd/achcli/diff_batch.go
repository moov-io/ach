// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/moov-io/ach"
)

func batchKey(b ach.Batcher) string {
	if !batcherPresent(b) || b.GetHeader() == nil {
		return ""
	}
	return b.GetHeader().BatchNumberField()
}

func batchNumber(b ach.Batcher) string {
	if n := batchKey(b); n != "" {
		return n
	}
	return "?"
}

func batchSEC(b ach.Batcher) string {
	if !batcherPresent(b) || b.GetHeader() == nil {
		return ""
	}
	return b.GetHeader().StandardEntryClassCode
}

func (d *differ) diffBatches(a, b []ach.Batcher) {
	pairs := align(len(a), len(b), func(i int) string { return batchKey(a[i]) }, func(i int) string { return batchKey(b[i]) })
	for _, p := range pairs {
		var left, right ach.Batcher
		if p.a >= 0 {
			left = a[p.a]
		}
		if p.b >= 0 {
			right = b[p.b]
		}
		d.diffBatch(left, right)
	}
}

func (d *differ) diffBatch(a, b ach.Batcher) {
	if !batcherPresent(a) && !batcherPresent(b) {
		return
	}
	prev := d.sec
	d.sec = maskedSEC(a, b)
	defer func() { d.sec = prev }()

	if !batcherPresent(a) {
		d.diffBatchOnly(true, b)
		return
	}
	if !batcherPresent(b) {
		d.diffBatchOnly(false, a)
		return
	}
	if isADVBatch(a) != isADVBatch(b) {
		d.diffBatchOnly(false, a)
		d.diffBatchOnly(true, b)
		return
	}
	var buf bytes.Buffer
	inner := d.child(&buf)
	inner.diffBatchRecords(a, b)
	if buf.Len() == 0 {
		return
	}
	d.heading(batchTitle(a, b))
	io.Copy(d.out, &buf)
}

func (d *differ) diffBatchRecords(a, b ach.Batcher) {
	d.nestedSide("  BatchHeader", a.GetHeader() == nil, b.GetHeader() == nil, batchHeaderCols(a.GetHeader(), b.GetHeader()))
	adv := isADVBatch(a)
	if adv || len(a.GetADVEntries()) > 0 || len(b.GetADVEntries()) > 0 {
		d.diffADVEntries(a.GetADVEntries(), b.GetADVEntries())
		d.nestedSide("  ADVBatchControl", a.GetADVControl() == nil, b.GetADVControl() == nil, advBatchControlCols(a.GetADVControl(), b.GetADVControl()))
	}
	if !adv || len(a.GetEntries()) > 0 || len(b.GetEntries()) > 0 {
		d.diffEntries(a.GetEntries(), b.GetEntries())
		d.nestedSide("  BatchControl", a.GetControl() == nil, b.GetControl() == nil, batchControlCols(a.GetControl(), b.GetControl()))
	}
}

func (d *differ) diffBatchOnly(right bool, b ach.Batcher) {
	if !batcherPresent(b) {
		return
	}
	name := d.left
	if right {
		name = d.right
	}
	d.heading(fmt.Sprintf("%s only in %s", batchTitle(b, b), name))
	if right {
		d.oneOrBoth("  ", true, false, batchHeaderCols(nil, b.GetHeader()))
	} else {
		d.oneOrBoth("  ", false, true, batchHeaderCols(b.GetHeader(), nil))
	}
	adv := isADVBatch(b)
	if adv || len(b.GetADVEntries()) > 0 {
		d.diffADVEntries(nilSideADV(right, b.GetADVEntries()))
		d.nestedSide("  ADVBatchControl", right, !right, oneSidedADVControl(right, b.GetADVControl()))
	}
	if !adv || len(b.GetEntries()) > 0 {
		d.diffEntries(nilSideEntries(right, b.GetEntries()))
		d.nestedSide("  BatchControl", right, !right, oneSidedControl(right, b.GetControl()))
	}
}

func batchTitle(a, b ach.Batcher) string {
	na, nb := batchNumber(a), batchNumber(b)
	sec := batchSEC(a)
	if sec == "" {
		sec = batchSEC(b)
	}
	if na == nb {
		return strings.TrimSpace("Batch " + na + " " + sec)
	}
	return fmt.Sprintf("Batch %s vs %s %s", na, nb, sec)
}

// maskedSEC returns ENR or DNE when either batch has that SEC code, as those
// batches carry personal data in their Addenda05 records.
func maskedSEC(a, b ach.Batcher) string {
	for _, sec := range []string{batchSEC(a), batchSEC(b)} {
		if sec == ach.ENR || sec == ach.DNE {
			return sec
		}
	}
	return ""
}

func isADVBatch(b ach.Batcher) bool {
	if !batcherPresent(b) {
		return false
	}
	if h := b.GetHeader(); h != nil && strings.TrimSpace(h.StandardEntryClassCode) != "" {
		return h.StandardEntryClassCode == ach.ADV
	}
	return len(b.GetADVEntries()) > 0
}

func nilSideEntries(right bool, entries []*ach.EntryDetail) ([]*ach.EntryDetail, []*ach.EntryDetail) {
	if right {
		return nil, entries
	}
	return entries, nil
}

func nilSideADV(right bool, entries []*ach.ADVEntryDetail) ([]*ach.ADVEntryDetail, []*ach.ADVEntryDetail) {
	if right {
		return nil, entries
	}
	return entries, nil
}

func oneSidedControl(right bool, bc *ach.BatchControl) []col {
	if right {
		return batchControlCols(nil, bc)
	}
	return batchControlCols(bc, nil)
}

func oneSidedADVControl(right bool, bc *ach.ADVBatchControl) []col {
	if right {
		return advBatchControlCols(nil, bc)
	}
	return advBatchControlCols(bc, nil)
}

func batchHeaderCols(a, b *ach.BatchHeader) []col {
	var l, r ach.BatchHeader
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"BatchNumber", num(l.BatchNumber), num(r.BatchNumber)},
		{"SECCode", trim(l.StandardEntryClassCode), trim(r.StandardEntryClassCode)},
		{"ServiceClassCode", num(l.ServiceClassCode), num(r.ServiceClassCode)},
		{"CompanyName", trim(l.CompanyName), trim(r.CompanyName)},
		{"DiscretionaryData", trim(l.CompanyDiscretionaryData), trim(r.CompanyDiscretionaryData)},
		{"Identification", trim(l.CompanyIdentification), trim(r.CompanyIdentification)},
		{"EntryDescription", trim(l.CompanyEntryDescription), trim(r.CompanyEntryDescription)},
		{"EffectiveEntryDate", trim(l.EffectiveEntryDate), trim(r.EffectiveEntryDate)},
		{"DescriptiveDate", trim(l.CompanyDescriptiveDate), trim(r.CompanyDescriptiveDate)},
		{"SettlementDate", trim(l.SettlementDate), trim(r.SettlementDate)},
		{"OriginatorStatusCode", num(l.OriginatorStatusCode), num(r.OriginatorStatusCode)},
		{"ODFIIdentification", trim(l.ODFIIdentification), trim(r.ODFIIdentification)},
	}
}

func batchControlCols(a, b *ach.BatchControl) []col {
	var l, r ach.BatchControl
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"ServiceClassCode", num(l.ServiceClassCode), num(r.ServiceClassCode)},
		{"EntryAddendaCount", num(l.EntryAddendaCount), num(r.EntryAddendaCount)},
		{"EntryHash", num(l.EntryHash), num(r.EntryHash)},
		{"TotalDebits", num(l.TotalDebitEntryDollarAmount), num(r.TotalDebitEntryDollarAmount)},
		{"TotalCredits", num(l.TotalCreditEntryDollarAmount), num(r.TotalCreditEntryDollarAmount)},
		{"CompanyIdentification", trim(l.CompanyIdentification), trim(r.CompanyIdentification)},
		{"MACCode", trim(l.MessageAuthenticationCode), trim(r.MessageAuthenticationCode)},
		{"ODFIIdentification", trim(l.ODFIIdentification), trim(r.ODFIIdentification)},
		{"BatchNumber", num(l.BatchNumber), num(r.BatchNumber)},
	}
}

func advBatchControlCols(a, b *ach.ADVBatchControl) []col {
	var l, r ach.ADVBatchControl
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"ServiceClassCode", num(l.ServiceClassCode), num(r.ServiceClassCode)},
		{"EntryAddendaCount", num(l.EntryAddendaCount), num(r.EntryAddendaCount)},
		{"EntryHash", num(l.EntryHash), num(r.EntryHash)},
		{"TotalDebits", num(l.TotalDebitEntryDollarAmount), num(r.TotalDebitEntryDollarAmount)},
		{"TotalCredits", num(l.TotalCreditEntryDollarAmount), num(r.TotalCreditEntryDollarAmount)},
		{"ACHOperatorData", trim(l.ACHOperatorData), trim(r.ACHOperatorData)},
		{"ODFIIdentification", trim(l.ODFIIdentification), trim(r.ODFIIdentification)},
		{"BatchNumber", num(l.BatchNumber), num(r.BatchNumber)},
	}
}

func entryKey(e *ach.EntryDetail) string {
	if e == nil {
		return ""
	}
	return e.TraceNumberField()
}

func (d *differ) diffEntries(a, b []*ach.EntryDetail) {
	pairs := align(len(a), len(b), func(i int) string { return entryKey(a[i]) }, func(i int) string { return entryKey(b[i]) })
	for _, p := range pairs {
		var left, right *ach.EntryDetail
		if p.a >= 0 {
			left = a[p.a]
		}
		if p.b >= 0 {
			right = b[p.b]
		}
		d.diffEntry(left, right)
	}
}

func entryTrace(a, b *ach.EntryDetail) string {
	if a != nil && a.TraceNumberField() != "" {
		return a.TraceNumberField()
	}
	if b != nil {
		return b.TraceNumberField()
	}
	return "?"
}

func (d *differ) diffEntry(a, b *ach.EntryDetail) {
	var buf bytes.Buffer
	inner := d.child(&buf)
	inner.oneOrBoth("    ", a == nil, b == nil, entryCols(a, b))
	inner.diffEntryAddenda(a, b)
	if buf.Len() == 0 {
		return
	}
	fmt.Fprintf(d.out, "  Entry %s\n", entryTrace(a, b))
	io.Copy(d.out, &buf)
}

func entryCols(a, b *ach.EntryDetail) []col {
	var l, r ach.EntryDetail
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"TransactionCode", num(l.TransactionCode), num(r.TransactionCode)},
		{"RDFIIdentification", trim(l.RDFIIdentification), trim(r.RDFIIdentification)},
		{"CheckDigit", trim(l.CheckDigit), trim(r.CheckDigit)},
		{"AccountNumber", trim(l.DFIAccountNumber), trim(r.DFIAccountNumber)},
		{"Amount", num(l.Amount), num(r.Amount)},
		{"IdentificationNumber", trim(l.IdentificationNumber), trim(r.IdentificationNumber)},
		{"Name", trim(l.IndividualName), trim(r.IndividualName)},
		{"DiscretionaryData", trim(l.DiscretionaryData), trim(r.DiscretionaryData)},
		{"AddendaRecordIndicator", num(l.AddendaRecordIndicator), num(r.AddendaRecordIndicator)},
		{"TraceNumber", trim(l.TraceNumber), trim(r.TraceNumber)},
		{"Category", trim(l.Category), trim(r.Category)},
	}
}

func addenda05Key(a *ach.Addenda05) string {
	if a == nil {
		return ""
	}
	return a.SequenceNumberField()
}

func (d *differ) diffEntryAddenda(a, b *ach.EntryDetail) {
	var left, right ach.EntryDetail
	if a != nil {
		left = *a
	}
	if b != nil {
		right = *b
	}
	d.nestedSide("      Addenda02", left.Addenda02 == nil, right.Addenda02 == nil, addenda02Cols(left.Addenda02, right.Addenda02))
	d.diffAddenda05(left.Addenda05, right.Addenda05)
	d.nestedSide("      Addenda98", left.Addenda98 == nil, right.Addenda98 == nil, addenda98Cols(left.Addenda98, right.Addenda98))
	d.nestedSide("      Addenda98Refused", left.Addenda98Refused == nil, right.Addenda98Refused == nil, addenda98RefusedCols(left.Addenda98Refused, right.Addenda98Refused))
	d.nestedSide("      Addenda99", left.Addenda99 == nil, right.Addenda99 == nil, addenda99Cols(left.Addenda99, right.Addenda99))
	d.nestedSide("      Addenda99Dishonored", left.Addenda99Dishonored == nil, right.Addenda99Dishonored == nil, addenda99DishonoredCols(left.Addenda99Dishonored, right.Addenda99Dishonored))
	d.nestedSide("      Addenda99Contested", left.Addenda99Contested == nil, right.Addenda99Contested == nil, addenda99ContestedCols(left.Addenda99Contested, right.Addenda99Contested))
}

func (d *differ) diffAddenda05(a, b []*ach.Addenda05) {
	pairs := align(len(a), len(b), func(i int) string { return addenda05Key(a[i]) }, func(i int) string { return addenda05Key(b[i]) })
	for _, p := range pairs {
		var left, right *ach.Addenda05
		if p.a >= 0 {
			left = a[p.a]
		}
		if p.b >= 0 {
			right = b[p.b]
		}
		d.nestedSide("      Addenda05", left == nil, right == nil, addenda05Cols(left, right))
	}
}

func addenda02Cols(a, b *ach.Addenda02) []col {
	var l, r ach.Addenda02
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"ReferenceInfoOne", trim(l.ReferenceInformationOne), trim(r.ReferenceInformationOne)},
		{"ReferenceInfoTwo", trim(l.ReferenceInformationTwo), trim(r.ReferenceInformationTwo)},
		{"TerminalIdentification", trim(l.TerminalIdentificationCode), trim(r.TerminalIdentificationCode)},
		{"TransactionSerial", trim(l.TransactionSerialNumber), trim(r.TransactionSerialNumber)},
		{"Date", trim(l.TransactionDate), trim(r.TransactionDate)},
		{"AuthCodeOrExpires", trim(l.AuthorizationCodeOrExpireDate), trim(r.AuthorizationCodeOrExpireDate)},
		{"Location", trim(l.TerminalLocation), trim(r.TerminalLocation)},
		{"City", trim(l.TerminalCity), trim(r.TerminalCity)},
		{"State", trim(l.TerminalState), trim(r.TerminalState)},
		{"TraceNumber", trim(l.TraceNumber), trim(r.TraceNumber)},
	}
}

func addenda05Cols(a, b *ach.Addenda05) []col {
	var l, r ach.Addenda05
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"PaymentRelatedInformation", trim(l.PaymentRelatedInformation), trim(r.PaymentRelatedInformation)},
		{"SequenceNumber", num(l.SequenceNumber), num(r.SequenceNumber)},
		{"EntryDetailSequenceNumber", num(l.EntryDetailSequenceNumber), num(r.EntryDetailSequenceNumber)},
	}
}

func addenda98Cols(a, b *ach.Addenda98) []col {
	var l, r ach.Addenda98
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"ChangeCode", trim(l.ChangeCode), trim(r.ChangeCode)},
		{"OriginalTrace", trim(l.OriginalTrace), trim(r.OriginalTrace)},
		{"OriginalDFI", trim(l.OriginalDFI), trim(r.OriginalDFI)},
		{"CorrectedData", trim(l.CorrectedData), trim(r.CorrectedData)},
		{"TraceNumber", trim(l.TraceNumber), trim(r.TraceNumber)},
	}
}

func addenda98RefusedCols(a, b *ach.Addenda98Refused) []col {
	var l, r ach.Addenda98Refused
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"RefusedChangeCode", trim(l.RefusedChangeCode), trim(r.RefusedChangeCode)},
		{"OriginalTrace", trim(l.OriginalTrace), trim(r.OriginalTrace)},
		{"OriginalDFI", trim(l.OriginalDFI), trim(r.OriginalDFI)},
		{"CorrectedData", trim(l.CorrectedData), trim(r.CorrectedData)},
		{"ChangeCode", trim(l.ChangeCode), trim(r.ChangeCode)},
		{"TraceSequenceNumber", trim(l.TraceSequenceNumber), trim(r.TraceSequenceNumber)},
	}
}

func addenda99Cols(a, b *ach.Addenda99) []col {
	var l, r ach.Addenda99
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"ReturnCode", trim(l.ReturnCode), trim(r.ReturnCode)},
		{"OriginalTrace", trim(l.OriginalTrace), trim(r.OriginalTrace)},
		{"DateOfDeath", trim(l.DateOfDeath), trim(r.DateOfDeath)},
		{"OriginalDFI", trim(l.OriginalDFI), trim(r.OriginalDFI)},
		{"AddendaInformation", trim(l.AddendaInformation), trim(r.AddendaInformation)},
		{"TraceNumber", trim(l.TraceNumber), trim(r.TraceNumber)},
	}
}

func addenda99DishonoredCols(a, b *ach.Addenda99Dishonored) []col {
	var l, r ach.Addenda99Dishonored
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"DishonoredReturnReasonCode", trim(l.DishonoredReturnReasonCode), trim(r.DishonoredReturnReasonCode)},
		{"OriginalEntryTraceNumber", trim(l.OriginalEntryTraceNumber), trim(r.OriginalEntryTraceNumber)},
		{"OriginalReceivingDFIIdentification", trim(l.OriginalReceivingDFIIdentification), trim(r.OriginalReceivingDFIIdentification)},
		{"ReturnTraceNumber", trim(l.ReturnTraceNumber), trim(r.ReturnTraceNumber)},
		{"ReturnSettlementDate", trim(l.ReturnSettlementDate), trim(r.ReturnSettlementDate)},
		{"ReturnReasonCode", trim(l.ReturnReasonCode), trim(r.ReturnReasonCode)},
		{"AddendaInformation", trim(l.AddendaInformation), trim(r.AddendaInformation)},
		{"TraceNumber", trim(l.TraceNumber), trim(r.TraceNumber)},
	}
}

func addenda99ContestedCols(a, b *ach.Addenda99Contested) []col {
	var l, r ach.Addenda99Contested
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"ContestedReturnCode", trim(l.ContestedReturnCode), trim(r.ContestedReturnCode)},
		{"OriginalEntryTraceNumber", trim(l.OriginalEntryTraceNumber), trim(r.OriginalEntryTraceNumber)},
		{"DateOriginalEntryReturned", trim(l.DateOriginalEntryReturned), trim(r.DateOriginalEntryReturned)},
		{"OriginalReceivingDFIIdentification", trim(l.OriginalReceivingDFIIdentification), trim(r.OriginalReceivingDFIIdentification)},
		{"OriginalSettlementDate", trim(l.OriginalSettlementDate), trim(r.OriginalSettlementDate)},
		{"ReturnTraceNumber", trim(l.ReturnTraceNumber), trim(r.ReturnTraceNumber)},
		{"ReturnSettlementDate", trim(l.ReturnSettlementDate), trim(r.ReturnSettlementDate)},
		{"ReturnReasonCode", trim(l.ReturnReasonCode), trim(r.ReturnReasonCode)},
		{"DishonoredReturnTraceNumber", trim(l.DishonoredReturnTraceNumber), trim(r.DishonoredReturnTraceNumber)},
		{"DishonoredReturnSettlementDate", trim(l.DishonoredReturnSettlementDate), trim(r.DishonoredReturnSettlementDate)},
		{"DishonoredReturnReasonCode", trim(l.DishonoredReturnReasonCode), trim(r.DishonoredReturnReasonCode)},
		{"TraceNumber", trim(l.TraceNumber), trim(r.TraceNumber)},
	}
}
