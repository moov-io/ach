// Copyright 2016 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"os"
	"testing"
)

func TestDecod(t *testing.T) {
	f, err := os.Open("./testdata/ppd-debit.ach")
	if err != nil {
		t.Errorf("%s: ", err)
	}
	defer f.Close()
	_, err = Decode(f)
	if err != nil {
		t.Errorf("Can not ach.Decode ach file: %v", err)
	}

}

func TestParseFileHeader(t *testing.T) {
	var line = "101 076401251 0764012510807291511A094101achdestname            companyname                    "
	record, err := parseFileHeader(line)
	if err != nil {
		t.Errorf("ParseFileHeader decode error: %v", err)
	}
	if record.RecordType != "1" {
		t.Errorf("RecordType Expected 1 got: %v", record.RecordType)
	}
	if record.PriorityCode != "01" {
		t.Errorf("PriorityCode Expected 01 got: %v", record.PriorityCode)
	}
	if record.ImmediateDestination != " 076401251" {
		t.Errorf("ImmediateDestination Expected ' 076401251' got: %v", record.ImmediateDestination)
	}
	if record.ImmediateOrigin != " 076401251" {
		t.Errorf("ImmediateOrigin Expected '   076401251' got: %v", record.ImmediateOrigin)
	}
	if record.FileCreationDate != "080729" {
		t.Errorf("FileCreationDate Expected '080729' got:'%v'", record.FileCreationDate)
	}
	if record.FileCreationTime != "1511" {
		t.Errorf("FileCreationTime Expected '1511' got:'%v'", record.FileCreationTime)
	}
	if record.FileIdModifier != "A" {
		t.Errorf("FileIdModifier Expected 'A' got:'%v'", record.FileIdModifier)
	}
	if record.RecordSize != "094" {
		t.Errorf("RecordSize Expected '094' got:'%v'", record.RecordSize)
	}
	if record.BlockingFactor != "10" {
		t.Errorf("BlockingFactor Expected '10' got:'%v'", record.BlockingFactor)
	}
	if record.FormatCode != "1" {
		t.Errorf("FormatCode Expected '1' got:'%v'", record.FormatCode)
	}
	if record.ImmediateDestinationName != "achdestname            " {
		t.Errorf("ImmediateDestinationName Expected 'achdestname           ' got:'%v'", record.ImmediateDestinationName)
	}
	if record.ImmidiateOriginName != "companyname            " {
		t.Errorf("ImmidiateOriginName Expected 'companyname          ' got: '%v'", record.ImmidiateOriginName)
	}
	if record.ReferenceCode != "        " {
		t.Errorf("ReferenceCode Expected '        ' got:'%v'", record.ReferenceCode)
	}
}
