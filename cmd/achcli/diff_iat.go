// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/moov-io/ach"
)

func advEntryKey(e *ach.ADVEntryDetail) string {
	if e == nil {
		return ""
	}
	return e.SequenceNumberField()
}

func (d *differ) diffADVEntries(a, b []*ach.ADVEntryDetail) {
	pairs := align(len(a), len(b), func(i int) string { return advEntryKey(a[i]) }, func(i int) string { return advEntryKey(b[i]) })
	for _, p := range pairs {
		var left, right *ach.ADVEntryDetail
		if p.a >= 0 {
			left = a[p.a]
		}
		if p.b >= 0 {
			right = b[p.b]
		}
		d.diffADVEntry(left, right)
	}
}

func (d *differ) diffADVEntry(a, b *ach.ADVEntryDetail) {
	var buf bytes.Buffer
	inner := &differ{out: &buf, color: d.color, left: d.left, right: d.right}
	inner.oneOrBoth("    ", a == nil, b == nil, advEntryCols(a, b))
	var left99, right99 *ach.Addenda99
	if a != nil {
		left99 = a.Addenda99
	}
	if b != nil {
		right99 = b.Addenda99
	}
	inner.nestedSide("      Addenda99", left99 == nil, right99 == nil, addenda99Cols(left99, right99))
	if buf.Len() == 0 {
		return
	}
	trace := "?"
	if a != nil {
		trace = a.SequenceNumberField()
	} else if b != nil {
		trace = b.SequenceNumberField()
	}
	fmt.Fprintf(d.out, "  ADV Entry %s\n", trace)
	io.Copy(d.out, &buf)
}

func advEntryCols(a, b *ach.ADVEntryDetail) []col {
	var l, r ach.ADVEntryDetail
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
		{"AdviceRoutingNumber", trim(l.AdviceRoutingNumber), trim(r.AdviceRoutingNumber)},
		{"FileIdentification", trim(l.FileIdentification), trim(r.FileIdentification)},
		{"ACHOperatorData", trim(l.ACHOperatorData), trim(r.ACHOperatorData)},
		{"Name", trim(l.IndividualName), trim(r.IndividualName)},
		{"DiscretionaryData", trim(l.DiscretionaryData), trim(r.DiscretionaryData)},
		{"AddendaRecordIndicator", num(l.AddendaRecordIndicator), num(r.AddendaRecordIndicator)},
		{"ACHOperatorRoutingNumber", trim(l.ACHOperatorRoutingNumber), trim(r.ACHOperatorRoutingNumber)},
		{"JulianDay", num(l.JulianDay), num(r.JulianDay)},
		{"SequenceNumber", num(l.SequenceNumber), num(r.SequenceNumber)},
		{"Category", trim(l.Category), trim(r.Category)},
	}
}

func iatBatchKey(b ach.IATBatch) string {
	if b.GetHeader() == nil {
		return ""
	}
	return b.GetHeader().BatchNumberField()
}

func (d *differ) diffIATBatches(a, b []ach.IATBatch) {
	pairs := align(len(a), len(b), func(i int) string { return iatBatchKey(a[i]) }, func(i int) string { return iatBatchKey(b[i]) })
	for _, p := range pairs {
		switch {
		case p.a >= 0 && p.b >= 0:
			d.diffIATBatch(a[p.a], b[p.b])
		case p.a >= 0:
			d.diffIATBatchOnly(false, a[p.a])
		default:
			d.diffIATBatchOnly(true, b[p.b])
		}
	}
}

func iatBatchTitle(a, b ach.IATBatch) string {
	na, nb := iatBatchKey(a), iatBatchKey(b)
	if na == "" {
		na = "?"
	}
	if nb == "" {
		nb = "?"
	}
	if na == nb {
		return "IAT Batch " + na
	}
	return fmt.Sprintf("IAT Batch %s vs %s", na, nb)
}

func (d *differ) diffIATBatch(a, b ach.IATBatch) {
	var buf bytes.Buffer
	inner := &differ{out: &buf, color: d.color, left: d.left, right: d.right}
	inner.nestedSide("  IATBatchHeader", a.GetHeader() == nil, b.GetHeader() == nil, iatBatchHeaderCols(a.GetHeader(), b.GetHeader()))
	inner.diffIATEntries(a.GetEntries(), b.GetEntries())
	inner.nestedSide("  BatchControl", a.GetControl() == nil, b.GetControl() == nil, batchControlCols(a.GetControl(), b.GetControl()))
	if buf.Len() == 0 {
		return
	}
	d.heading(iatBatchTitle(a, b))
	io.Copy(d.out, &buf)
}

func (d *differ) diffIATBatchOnly(right bool, b ach.IATBatch) {
	name := d.left
	if right {
		name = d.right
	}
	d.heading(fmt.Sprintf("%s only in %s", iatBatchTitle(b, b), name))
	if right {
		d.oneOrBoth("  ", true, false, iatBatchHeaderCols(nil, b.GetHeader()))
		d.diffIATEntries(nil, b.GetEntries())
		d.nestedSide("  BatchControl", true, b.GetControl() == nil, batchControlCols(nil, b.GetControl()))
		return
	}
	d.oneOrBoth("  ", false, true, iatBatchHeaderCols(b.GetHeader(), nil))
	d.diffIATEntries(b.GetEntries(), nil)
	d.nestedSide("  BatchControl", b.GetControl() == nil, true, batchControlCols(b.GetControl(), nil))
}

func iatBatchHeaderCols(a, b *ach.IATBatchHeader) []col {
	var l, r ach.IATBatchHeader
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
		{"IATIndicator", trim(l.IATIndicator), trim(r.IATIndicator)},
		{"DestinationCountryCode", trim(l.ISODestinationCountryCode), trim(r.ISODestinationCountryCode)},
		{"ForeignExchangeIndicator", trim(l.ForeignExchangeIndicator), trim(r.ForeignExchangeIndicator)},
		{"ForeignExchangeReferenceIndicator", num(l.ForeignExchangeReferenceIndicator), num(r.ForeignExchangeReferenceIndicator)},
		{"ForeignExchangeReference", trim(l.ForeignExchangeReference), trim(r.ForeignExchangeReference)},
		{"CompanyEntryDescription", trim(l.CompanyEntryDescription), trim(r.CompanyEntryDescription)},
		{"OriginatorIdentification", trim(l.OriginatorIdentification), trim(r.OriginatorIdentification)},
		{"ISOOriginatingCurrencyCode", trim(l.ISOOriginatingCurrencyCode), trim(r.ISOOriginatingCurrencyCode)},
		{"ISODestinationCurrencyCode", trim(l.ISODestinationCurrencyCode), trim(r.ISODestinationCurrencyCode)},
		{"ODFIIdentification", trim(l.ODFIIdentification), trim(r.ODFIIdentification)},
		{"EffectiveEntryDate", trim(l.EffectiveEntryDate), trim(r.EffectiveEntryDate)},
		{"SettlementDate", trim(l.SettlementDate), trim(r.SettlementDate)},
		{"OriginatorStatusCode", num(l.OriginatorStatusCode), num(r.OriginatorStatusCode)},
	}
}

func iatEntryKey(e *ach.IATEntryDetail) string {
	if e == nil {
		return ""
	}
	return e.TraceNumberField()
}

func (d *differ) diffIATEntries(a, b []*ach.IATEntryDetail) {
	pairs := align(len(a), len(b), func(i int) string { return iatEntryKey(a[i]) }, func(i int) string { return iatEntryKey(b[i]) })
	for _, p := range pairs {
		var left, right *ach.IATEntryDetail
		if p.a >= 0 {
			left = a[p.a]
		}
		if p.b >= 0 {
			right = b[p.b]
		}
		d.diffIATEntry(left, right)
	}
}

func (d *differ) diffIATEntry(a, b *ach.IATEntryDetail) {
	var buf bytes.Buffer
	inner := &differ{out: &buf, color: d.color, left: d.left, right: d.right}
	inner.oneOrBoth("    ", a == nil, b == nil, iatEntryCols(a, b))
	inner.diffIATAddenda(a, b)
	if buf.Len() == 0 {
		return
	}
	trace := "?"
	if a != nil && a.TraceNumberField() != "" {
		trace = a.TraceNumberField()
	} else if b != nil {
		trace = b.TraceNumberField()
	}
	fmt.Fprintf(d.out, "  Entry %s\n", trace)
	io.Copy(d.out, &buf)
}

func iatEntryCols(a, b *ach.IATEntryDetail) []col {
	var l, r ach.IATEntryDetail
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
		{"AddendaRecords", num(l.AddendaRecords), num(r.AddendaRecords)},
		{"OFACScreeningIndicator", trim(l.OFACScreeningIndicator), trim(r.OFACScreeningIndicator)},
		{"SecondaryOFACScreeningIndicator", trim(l.SecondaryOFACScreeningIndicator), trim(r.SecondaryOFACScreeningIndicator)},
		{"TraceNumber", trim(l.TraceNumber), trim(r.TraceNumber)},
		{"Category", trim(l.Category), trim(r.Category)},
	}
}

func (d *differ) diffIATAddenda(a, b *ach.IATEntryDetail) {
	var left, right ach.IATEntryDetail
	if a != nil {
		left = *a
	}
	if b != nil {
		right = *b
	}
	d.nestedSide("      Addenda10", left.Addenda10 == nil, right.Addenda10 == nil, addenda10Cols(left.Addenda10, right.Addenda10))
	d.nestedSide("      Addenda11", left.Addenda11 == nil, right.Addenda11 == nil, addenda11Cols(left.Addenda11, right.Addenda11))
	d.nestedSide("      Addenda12", left.Addenda12 == nil, right.Addenda12 == nil, addenda12Cols(left.Addenda12, right.Addenda12))
	d.nestedSide("      Addenda13", left.Addenda13 == nil, right.Addenda13 == nil, addenda13Cols(left.Addenda13, right.Addenda13))
	d.nestedSide("      Addenda14", left.Addenda14 == nil, right.Addenda14 == nil, addenda14Cols(left.Addenda14, right.Addenda14))
	d.nestedSide("      Addenda15", left.Addenda15 == nil, right.Addenda15 == nil, addenda15Cols(left.Addenda15, right.Addenda15))
	d.nestedSide("      Addenda16", left.Addenda16 == nil, right.Addenda16 == nil, addenda16Cols(left.Addenda16, right.Addenda16))
	d.diffAddenda17(left.Addenda17, right.Addenda17)
	d.diffAddenda18(left.Addenda18, right.Addenda18)
	d.nestedSide("      Addenda98", left.Addenda98 == nil, right.Addenda98 == nil, addenda98Cols(left.Addenda98, right.Addenda98))
	d.nestedSide("      Addenda99", left.Addenda99 == nil, right.Addenda99 == nil, addenda99Cols(left.Addenda99, right.Addenda99))
}

func addenda10Cols(a, b *ach.Addenda10) []col {
	var l, r ach.Addenda10
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"TypeCode", trim(l.TypeCode), trim(r.TypeCode)},
		{"TransactionTypeCode", trim(l.TransactionTypeCode), trim(r.TransactionTypeCode)},
		{"ForeignPaymentAmount", num(l.ForeignPaymentAmount), num(r.ForeignPaymentAmount)},
		{"ForeignTraceNumber", trim(l.ForeignTraceNumber), trim(r.ForeignTraceNumber)},
		{"Name", trim(l.Name), trim(r.Name)},
	}
}

func addenda11Cols(a, b *ach.Addenda11) []col {
	var l, r ach.Addenda11
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"OriginatorName", trim(l.OriginatorName), trim(r.OriginatorName)},
		{"OriginatorStreetAddress", trim(l.OriginatorStreetAddress), trim(r.OriginatorStreetAddress)},
	}
}

func addenda12Cols(a, b *ach.Addenda12) []col {
	var l, r ach.Addenda12
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"OriginatorCityStateProvince", trim(l.OriginatorCityStateProvince), trim(r.OriginatorCityStateProvince)},
		{"OriginatorCountryPostalCode", trim(l.OriginatorCountryPostalCode), trim(r.OriginatorCountryPostalCode)},
	}
}

func addenda13Cols(a, b *ach.Addenda13) []col {
	var l, r ach.Addenda13
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"ODFIName", trim(l.ODFIName), trim(r.ODFIName)},
		{"ODFIIDNumberQualifier", trim(l.ODFIIDNumberQualifier), trim(r.ODFIIDNumberQualifier)},
		{"ODFIIdentification", trim(l.ODFIIdentification), trim(r.ODFIIdentification)},
		{"ODFIBranchCountryCode", trim(l.ODFIBranchCountryCode), trim(r.ODFIBranchCountryCode)},
	}
}

func addenda14Cols(a, b *ach.Addenda14) []col {
	var l, r ach.Addenda14
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"RDFIName", trim(l.RDFIName), trim(r.RDFIName)},
		{"RDFIIDNumberQualifier", trim(l.RDFIIDNumberQualifier), trim(r.RDFIIDNumberQualifier)},
		{"RDFIIdentification", trim(l.RDFIIdentification), trim(r.RDFIIdentification)},
		{"RDFIBranchCountryCode", trim(l.RDFIBranchCountryCode), trim(r.RDFIBranchCountryCode)},
	}
}

func addenda15Cols(a, b *ach.Addenda15) []col {
	var l, r ach.Addenda15
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"ReceiverIDNumber", trim(l.ReceiverIDNumber), trim(r.ReceiverIDNumber)},
		{"ReceiverStreetAddress", trim(l.ReceiverStreetAddress), trim(r.ReceiverStreetAddress)},
	}
}

func addenda16Cols(a, b *ach.Addenda16) []col {
	var l, r ach.Addenda16
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"ReceiverCityStateProvince", trim(l.ReceiverCityStateProvince), trim(r.ReceiverCityStateProvince)},
		{"ReceiverCountryPostalCode", trim(l.ReceiverCountryPostalCode), trim(r.ReceiverCountryPostalCode)},
	}
}

func addenda17Key(a *ach.Addenda17) string {
	if a == nil {
		return ""
	}
	return a.SequenceNumberField()
}

func (d *differ) diffAddenda17(a, b []*ach.Addenda17) {
	pairs := align(len(a), len(b), func(i int) string { return addenda17Key(a[i]) }, func(i int) string { return addenda17Key(b[i]) })
	for _, p := range pairs {
		var left, right *ach.Addenda17
		if p.a >= 0 {
			left = a[p.a]
		}
		if p.b >= 0 {
			right = b[p.b]
		}
		d.nestedSide("      Addenda17", left == nil, right == nil, addenda17Cols(left, right))
	}
}

func addenda17Cols(a, b *ach.Addenda17) []col {
	var l, r ach.Addenda17
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"PaymentRelatedInformation", trim(l.PaymentRelatedInformation), trim(r.PaymentRelatedInformation)},
		{"SequenceNumber", num(l.SequenceNumber), num(r.SequenceNumber)},
	}
}

func addenda18Key(a *ach.Addenda18) string {
	if a == nil {
		return ""
	}
	return a.SequenceNumberField()
}

func (d *differ) diffAddenda18(a, b []*ach.Addenda18) {
	pairs := align(len(a), len(b), func(i int) string { return addenda18Key(a[i]) }, func(i int) string { return addenda18Key(b[i]) })
	for _, p := range pairs {
		var left, right *ach.Addenda18
		if p.a >= 0 {
			left = a[p.a]
		}
		if p.b >= 0 {
			right = b[p.b]
		}
		d.nestedSide("      Addenda18", left == nil, right == nil, addenda18Cols(left, right))
	}
}

func addenda18Cols(a, b *ach.Addenda18) []col {
	var l, r ach.Addenda18
	if a != nil {
		l = *a
	}
	if b != nil {
		r = *b
	}
	return []col{
		{"ForeignCorrespondentBankName", trim(l.ForeignCorrespondentBankName), trim(r.ForeignCorrespondentBankName)},
		{"ForeignCorrespondentBankIDNumberQualifier", trim(l.ForeignCorrespondentBankIDNumberQualifier), trim(r.ForeignCorrespondentBankIDNumberQualifier)},
		{"ForeignCorrespondentBankIDNumber", trim(l.ForeignCorrespondentBankIDNumber), trim(r.ForeignCorrespondentBankIDNumber)},
		{"ForeignCorrespondentBankBranchCountryCode", trim(l.ForeignCorrespondentBankBranchCountryCode), trim(r.ForeignCorrespondentBankBranchCountryCode)},
		{"SequenceNumber", num(l.SequenceNumber), num(r.SequenceNumber)},
	}
}
