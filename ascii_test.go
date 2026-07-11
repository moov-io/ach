// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.

package ach

import (
	"bufio"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestEntryDetailASCIIParity(t *testing.T) {
	records := []string{
		"",
		"6short",
		"62705320001912345            0000010500c-1            Arnold Wade           DD0076401255655291",
		"62705320001912345            0000010500c-1            Arnold Wadé           DD0076401255655291",
	}
	records = append(records, fixtureRecords(t, '6', "")...)

	for _, record := range records {
		got, want := NewEntryDetail(), NewEntryDetail()
		got.Parse(record)
		want.parseRuneSafe(record)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("fast and slow EntryDetail parsing differ for %q\nfast: %#v\nslow: %#v", record, got, want)
		}
	}
}

func TestAddenda05ASCIIParity(t *testing.T) {
	records := []string{
		"",
		"7short",
		"705WEB                                        DIEGO MAY                            00010000001",
		"705WEB                                        DIÉGO MAY                            00010000001",
	}
	records = append(records, fixtureRecords(t, '7', "05")...)

	for _, record := range records {
		got, want := NewAddenda05(), NewAddenda05()
		got.Parse(record)
		want.parseRuneSafe(record)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("fast and slow Addenda05 parsing differ for %q\nfast: %#v\nslow: %#v", record, got, want)
		}
	}
}

func fixtureRecords(t *testing.T, recordType byte, subtype string) []string {
	t.Helper()
	paths, err := filepath.Glob(filepath.Join("test", "testdata", "*.ach"))
	if err != nil {
		t.Fatal(err)
	}

	var records []string
	for _, path := range paths {
		fd, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		scanner := bufio.NewScanner(fd)
		for scanner.Scan() {
			line := strings.TrimSuffix(scanner.Text(), "\r")
			if len(line) == 94 && line[0] == recordType && (subtype == "" || line[1:3] == subtype) {
				records = append(records, line)
				if len(records) == 20 {
					break
				}
			}
		}
		if err := scanner.Err(); err != nil {
			fd.Close()
			t.Fatal(err)
		}
		fd.Close()
		if len(records) == 20 {
			break
		}
	}
	if len(records) == 0 {
		t.Fatalf("no fixture records found for type %q subtype %q", recordType, subtype)
	}
	return records
}

var (
	benchmarkEntryDetail EntryDetail
	benchmarkAddenda05   Addenda05
	benchmarkPosition    int
	benchmarkASCII       bool
	benchmarkEntrySpans  entryDetailSpans
)

type entryDetailSpans struct {
	transactionCode        string
	rdfiIdentification     string
	checkDigit             string
	dfiAccountNumber       string
	amount                 string
	identificationNumber   string
	individualName         string
	discretionaryData      string
	addendaRecordIndicator string
	traceNumber            string
}

func BenchmarkEntryDetailParsePaths(b *testing.B) {
	record := "62705320001912345            0000010500c-1            Arnold Wade           DD0076401255655291"
	b.Run("rune_safe", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			var parsed EntryDetail
			parsed.parseRuneSafe(record)
			benchmarkEntryDetail = parsed
		}
	})
	b.Run("ascii", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			var parsed EntryDetail
			parsed.Parse(record)
			benchmarkEntryDetail = parsed
		}
	})
}

func BenchmarkEntryDetailParseComponents(b *testing.B) {
	record := "62705320001912345            0000010500c-1            Arnold Wade           DD0076401255655291"
	b.Run("rune_boundaries", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			if utf8.RuneCountInString(record) != 94 {
				b.Fatal("invalid benchmark record")
			}
			byteIndex := 0
			for i := 0; i < 94; i++ {
				_, size := utf8.DecodeRuneInString(record[byteIndex:])
				byteIndex += size
			}
			benchmarkPosition = byteIndex
		}
	})
	b.Run("ascii_check", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			benchmarkASCII = isASCII(record)
		}
	})
	b.Run("field_decode", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			var parsed EntryDetail
			parsed.parseASCII(record)
			benchmarkEntryDetail = parsed
		}
	})
}

func BenchmarkEntryDetailPositionMechanism(b *testing.B) {
	record := "62705320001912345            0000010500c-1            Arnold Wade           DD0076401255655291"
	b.Run("rune_offsets_and_slices", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			if utf8.RuneCountInString(record) != 94 {
				b.Fatal("invalid benchmark record")
			}
			positions := make([]int, 95)
			byteIndex := 0
			for i := 0; i < 94; i++ {
				positions[i] = byteIndex
				_, size := utf8.DecodeRuneInString(record[byteIndex:])
				byteIndex += size
			}
			positions[94] = byteIndex
			benchmarkEntrySpans = entryDetailSpans{
				transactionCode:        record[positions[1]:positions[3]],
				rdfiIdentification:     record[positions[3]:positions[11]],
				checkDigit:             record[positions[11]:positions[12]],
				dfiAccountNumber:       record[positions[12]:positions[29]],
				amount:                 record[positions[29]:positions[39]],
				identificationNumber:   record[positions[39]:positions[54]],
				individualName:         record[positions[54]:positions[76]],
				discretionaryData:      record[positions[76]:positions[78]],
				addendaRecordIndicator: record[positions[78]:positions[79]],
				traceNumber:            record[positions[79]:positions[94]],
			}
		}
	})
	b.Run("byte_offsets_and_slices", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			benchmarkEntrySpans = entryDetailSpans{
				transactionCode:        record[1:3],
				rdfiIdentification:     record[3:11],
				checkDigit:             record[11:12],
				dfiAccountNumber:       record[12:29],
				amount:                 record[29:39],
				identificationNumber:   record[39:54],
				individualName:         record[54:76],
				discretionaryData:      record[76:78],
				addendaRecordIndicator: record[78:79],
				traceNumber:            record[79:94],
			}
		}
	})
}

func BenchmarkAddenda05ParsePaths(b *testing.B) {
	record := "705WEB                                        DIEGO MAY                            00010000001"
	b.Run("rune_safe", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			var parsed Addenda05
			parsed.parseRuneSafe(record)
			benchmarkAddenda05 = parsed
		}
	})
	b.Run("ascii", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			var parsed Addenda05
			parsed.Parse(record)
			benchmarkAddenda05 = parsed
		}
	})
}
