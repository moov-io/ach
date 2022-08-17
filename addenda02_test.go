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
	"testing"

	"github.com/moov-io/base"
)

// mockAddenda02 creates a mock Addenda02 record
func mockAddenda02() *Addenda02 {
	addenda02 := NewAddenda02()
	addenda02.ReferenceInformationOne = "REFONEA"
	addenda02.ReferenceInformationTwo = "REF"
	addenda02.TerminalIdentificationCode = "TERM02"
	addenda02.TransactionSerialNumber = "100049"
	addenda02.TransactionDate = "0612"
	addenda02.AuthorizationCodeOrExpireDate = "123456"
	addenda02.TerminalLocation = "Target Store 0049"
	addenda02.TerminalCity = "PHILADELPHIA"
	addenda02.TerminalState = "PA"
	addenda02.TraceNumber = "121042880000123"
	return addenda02
}

// TestMockAddenda02 validates mockAddenda02
func TestMockAddenda02(t *testing.T) {
	addenda02 := mockAddenda02()
	if err := addenda02.Validate(); err != nil {
		t.Error("mockAddenda02 does not validate and will break other tests")
	}
}

// testAddenda02ValidTypeCode validates Addenda02 TypeCode
func testAddenda02ValidTypeCode(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TypeCode = "65"
	err := addenda02.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02ValidTypeCode tests validating Addenda02 TypeCode
func TestAddenda02ValidTypeCode(t *testing.T) {
	testAddenda02ValidTypeCode(t)
}

// BenchmarkAddenda02ValidTypeCode benchmarks validating Addenda02 TypeCode
func BenchmarkAddenda02ValidTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02ValidTypeCode(b)
	}
}

// testAddenda02TypeCode02 TypeCode is 02 if TypeCode is a valid TypeCode
func testAddenda02TypeCode02(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TypeCode = "05"
	err := addenda02.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02TypeCode02 tests TypeCode is 02 if TypeCode is a valid TypeCode
func TestAddenda02TypeCode02(t *testing.T) {
	testAddenda02TypeCode02(t)
}

// BenchmarkAddenda02TypeCode02 benchmarks TypeCode is 02 if TypeCode is a valid TypeCode
func BenchmarkAddenda02TypeCode02(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TypeCode02(b)
	}
}

// testAddenda02FieldInclusionTypeCode validates TypeCode fieldInclusion
func testAddenda02FieldInclusionTypeCode(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TypeCode = ""
	err := addenda02.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02FieldInclusionTypeCode tests validating TypeCode fieldInclusion
func TestAddenda02FieldInclusionTypeCode(t *testing.T) {
	testAddenda02FieldInclusionTypeCode(t)
}

// BenchmarkAddenda02FieldInclusionTypeCode benchmarks validating TypeCode fieldInclusion
func BenchmarkAddenda02FieldInclusionTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02FieldInclusionTypeCode(b)
	}
}

// testAddenda02TerminalIdentificationCode validates TerminalIdentificationCode is required
func testAddenda02TerminalIdentificationCode(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TerminalIdentificationCode = ""
	err := addenda02.Validate()
	if !base.Match(err, ErrFieldRequired) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02TerminalIdentificationCode tests validating TerminalIdentificationCode is required
func TestAddenda02TerminalIdentificationCode(t *testing.T) {
	testAddenda02TerminalIdentificationCode(t)
}

// BenchmarkAddenda02TerminalIdentificationCode benchmarks validating TerminalIdentificationCode is required
func BenchmarkAddenda02TerminalIdentificationCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TerminalIdentificationCode(b)
	}
}

// testAddenda02TransactionSerialNumber validates TransactionSerialNumber is required
func testAddenda02TransactionSerialNumber(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TransactionSerialNumber = ""
	err := addenda02.Validate()
	if !base.Match(err, ErrFieldRequired) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02TransactionSerialNumber tests validating TransactionSerialNumber is required
func TestAddenda02TransactionSerialNumber(t *testing.T) {
	testAddenda02TransactionSerialNumber(t)
}

// BenchmarkAddenda02TransactionSerialNumber benchmarks validating TransactionSerialNumber is required
func BenchmarkAddenda02TransactionSerialNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TransactionSerialNumber(b)
	}
}

// testAddenda02TransactionDate validates TransactionDate is required
func testAddenda02TransactionDate(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TransactionDate = ""
	err := addenda02.Validate()
	if !base.Match(err, ErrFieldRequired) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02TransactionDate tests validating TransactionDate is required
func TestAddenda02TransactionDate(t *testing.T) {
	testAddenda02TransactionDate(t)
}

// BenchmarkAddenda02TransactionDate benchmarks validating TransactionDate is required
func BenchmarkAddenda02TransactionDate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TransactionDate(b)
	}
}

// testAddenda02TerminalLocation validates TerminalLocation is required
func testAddenda02TerminalLocation(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TerminalLocation = ""
	err := addenda02.Validate()
	if !base.Match(err, ErrFieldRequired) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02TerminalLocation tests validating TerminalLocation is required
func TestAddenda02TerminalLocation(t *testing.T) {
	testAddenda02TerminalLocation(t)
}

// BenchmarkAddenda02TerminalLocation benchmarks validating TerminalLocation is required
func BenchmarkAddenda02TerminalLocation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TerminalLocation(b)
	}
}

// testAddenda02TerminalCity validates TerminalCity is required
func testAddenda02TerminalCity(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TerminalCity = ""
	err := addenda02.Validate()
	if !base.Match(err, ErrFieldRequired) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02TerminalCity tests validating TerminalCity is required
func TestAddenda02TerminalCity(t *testing.T) {
	testAddenda02TerminalCity(t)
}

// BenchmarkAddenda02TerminalCity benchmarks validating TerminalCity is required
func BenchmarkAddenda02TerminalCity(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TerminalCity(b)
	}
}

// testAddenda02TerminalState validates TerminalState is required
func testAddenda02TerminalState(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TerminalState = ""
	err := addenda02.Validate()
	if !base.Match(err, ErrFieldRequired) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02TerminalState tests validating TerminalState is required
func TestAddenda02TerminalState(t *testing.T) {
	testAddenda02TerminalState(t)
}

// BenchmarkAddenda02TerminalState benchmarks validating TerminalState is required
func BenchmarkAddenda02TerminalState(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TerminalState(b)
	}
}

// TestAddenda02String validates that a known parsed Addenda02 record can be return to a string of the same value
func testAddenda02String(t testing.TB) {
	addenda02 := NewAddenda02()
	var line = "702REFONEAREFTERM021000490612123456Target Store 0049          PHILADELPHIA   PA121042880000123"
	addenda02.Parse(line)
	if addenda02.String() != line {
		t.Errorf("Strings do not match")
	}
	if addenda02.TypeCode != "02" {
		t.Errorf("TypeCode Expected 02 got: %v", addenda02.TypeCode)
	}
}

// TestAddenda02String tests validating that a known parsed Addenda02 record can be return to a string of the same value
func TestAddenda02String(t *testing.T) {
	testAddenda02String(t)
}

// BenchmarkAddenda02String benchmarks validating that a known parsed Addenda02 record can be return to a string of the same value
func BenchmarkAddenda02String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02String(b)
	}
}

// testAddenda02TransactionDateMonth validates the month is valid for transactionDate
func testAddenda02TransactionDateMonth(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TransactionDate = "1306"
	err := addenda02.Validate()
	if !base.Match(err, ErrValidMonth) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02TransactionDateMonth tests validating the month is valid for transactionDate
func TestAddenda02TransactionDateMonth(t *testing.T) {
	testAddenda02TransactionDateMonth(t)
}

// BenchmarkAddenda02TransactionDateMonth benchmarks validating the month is valid for transactionDate
func BenchmarkAddenda02TransactionDateMonth(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TransactionDateMonth(b)
	}
}

// testAddenda02TransactionDateDay validates the day is valid for transactionDate
func testAddenda02TransactionDateDay(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TransactionDate = "0205"
	err := addenda02.Validate()
	// no error expected
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02TransactionDateDay tests validating the day is valid for transactionDate
func TestAddenda02TransactionDateDay(t *testing.T) {
	testAddenda02TransactionDateDay(t)
}

// BenchmarkAddenda02TransactionDateDay benchmarks validating the day is valid for transactionDate
func BenchmarkAddenda02TransactionDateDay(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TransactionDateDay(b)
	}
}

// testAddenda02TransactionDateFeb validates the day is valid for transactionDate
func testAddenda02TransactionDateFeb(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TransactionDate = "0230"
	err := addenda02.Validate()
	if !base.Match(err, ErrValidDay) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02TransactionDateFeb tests validating the day is valid for transactionDate
func TestAddenda02TransactionDateFeb(t *testing.T) {
	testAddenda02TransactionDateFeb(t)
}

// BenchmarkAddenda02TransactionDateFeb benchmarks validating the day is valid for transactionDate
func BenchmarkAddenda02TransactionDateFeb(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TransactionDateFeb(b)
	}
}

// testAddenda02TransactionDate30Day validates the day is valid for transactionDate
func testAddenda02TransactionDate30Day(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TransactionDate = "0630"
	err := addenda02.Validate()
	// no error expected
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02TransactionDate30Day tests validating the day is valid for transactionDate
func TestAddenda02TransactionDate30Day(t *testing.T) {
	testAddenda02TransactionDate30Day(t)
}

// BenchmarkAddenda02TransactionDate30Day benchmarks validating the day is valid for transactionDate
func BenchmarkAddenda02TransactionDate30Day(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TransactionDate30Day(b)
	}
}

// testAddenda02TransactionDate31Day validates the day is valid for transactionDate
func testAddenda02TransactionDate31Day(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TransactionDate = "0131"
	err := addenda02.Validate()
	// no error expected
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02TransactionDate31Day tests validating the day is valid for transactionDate
func TestAddenda02TransactionDate31Day(t *testing.T) {
	testAddenda02TransactionDate31Day(t)
}

// BenchmarkAddenda02TransactionDate31Day benchmarks validating the day is valid for transactionDate
func BenchmarkAddenda02TransactionDate31Day(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TransactionDate31Day(b)
	}
}

// testAddenda02TransactionDateInvalidDay validates the day is invalid for transactionDate
func testAddenda02TransactionDateInvalidDay(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TransactionDate = "1039"
	err := addenda02.Validate()
	if !base.Match(err, ErrValidDay) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda02TransactionDateInvalidDay tests validating the day is invalid for transactionDate
func TestAddenda02TransactionDateInvalidDay(t *testing.T) {
	testAddenda02TransactionDateInvalidDay(t)
}

// BenchmarkAddenda02TransactionDateInvalidDay benchmarks validating the day is invalid for transactionDate
func BenchmarkAddenda02TransactionDateInvalidDay(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TransactionDateInvalidDay(b)
	}
}

// testReferenceInformationOneAlphaNumeric validates ReferenceInformationOne is alphanumeric
func testReferenceInformationOneAlphaNumeric(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.ReferenceInformationOne = "®"
	err := addenda02.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestReferenceInformationOneAlphaNumeric tests validating ReferenceInformationOne is alphanumeric
func TestReferenceInformationOneAlphaNumeric(t *testing.T) {
	testReferenceInformationOneAlphaNumeric(t)
}

// BenchmarkReferenceInformationOneAlphaNumeric benchmarks validating ReferenceInformationOne is alphanumeric
func BenchmarkReferenceInformationOneAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testReferenceInformationOneAlphaNumeric(b)
	}
}

// testReferenceInformationTwoAlphaNumeric validates ReferenceInformationTwo is alphanumeric
func testReferenceInformationTwoAlphaNumeric(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.ReferenceInformationTwo = "®"
	err := addenda02.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestReferenceInformationTwoAlphaNumeric tests validating ReferenceInformationTwo is alphanumeric
func TestReferenceInformationTwoAlphaNumeric(t *testing.T) {
	testReferenceInformationTwoAlphaNumeric(t)
}

// BenchmarkReferenceInformationTwoAlphaNumeric benchmarks validating ReferenceInformationTwo is alphanumeric
func BenchmarkReferenceInformationTwoAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testReferenceInformationTwoAlphaNumeric(b)
	}
}

// testTerminalIdentificationCodeAlphaNumeric validates TerminalIdentificationCode is alphanumeric
func testTerminalIdentificationCodeAlphaNumeric(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TerminalIdentificationCode = "®"
	err := addenda02.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestTerminalIdentificationCodeAlphaNumeric tests validating TerminalIdentificationCode is alphanumeric
func TestTerminalIdentificationCodeAlphaNumeric(t *testing.T) {
	testTerminalIdentificationCodeAlphaNumeric(t)
}

// BenchmarkTerminalIdentificationCodeAlphaNumeric benchmarks validating TerminalIdentificationCode is alphanumeric
func BenchmarkTerminalIdentificationCodeAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testTerminalIdentificationCodeAlphaNumeric(b)
	}
}

// testTransactionSerialNumberAlphaNumeric validates TransactionSerialNumber is alphanumeric
func testTransactionSerialNumberAlphaNumeric(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TransactionSerialNumber = "®"
	err := addenda02.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestTransactionSerialNumberAlphaNumeric tests validating TransactionSerialNumber is alphanumeric
func TestTransactionSerialNumberAlphaNumeric(t *testing.T) {
	testTransactionSerialNumberAlphaNumeric(t)
}

// BenchmarkTransactionSerialNumberAlphaNumeric benchmarks validating TransactionSerialNumber is alphanumeric
func BenchmarkTransactionSerialNumberAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testTransactionSerialNumberAlphaNumeric(b)
	}
}

// testAuthorizationCodeOrExpireDateAlphaNumeric validates AuthorizationCodeOrExpireDate is alphanumeric
func testAuthorizationCodeOrExpireDateAlphaNumeric(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.AuthorizationCodeOrExpireDate = "®"
	err := addenda02.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAuthorizationCodeOrExpireDateAlphaNumeric tests validating AuthorizationCodeOrExpireDate is alphanumeric
func TestAuthorizationCodeOrExpireDateAlphaNumeric(t *testing.T) {
	testAuthorizationCodeOrExpireDateAlphaNumeric(t)
}

// BenchmarkAuthorizationCodeOrExpireDateAlphaNumeric benchmarks validating AuthorizationCodeOrExpireDate is alphanumeric
func BenchmarkAuthorizationCodeOrExpireDateAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAuthorizationCodeOrExpireDateAlphaNumeric(b)
	}
}

// testTerminalLocationAlphaNumeric validates TerminalLocation is alphanumeric
func testTerminalLocationAlphaNumeric(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TerminalLocation = "®"
	err := addenda02.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestTerminalLocationAlphaNumeric tests validating TerminalLocation is alphanumeric
func TestTerminalLocationAlphaNumeric(t *testing.T) {
	testTerminalLocationAlphaNumeric(t)
}

// BenchmarkTerminalLocationAlphaNumeric benchmarks validating TerminalLocation is alphanumeric
func BenchmarkTerminalLocationAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testTerminalLocationAlphaNumeric(b)
	}
}

// testTerminalCityAlphaNumeric validates TerminalCity is alphanumeric
func testTerminalCityAlphaNumeric(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TerminalCity = "®"
	err := addenda02.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestTerminalCityAlphaNumeric tests validating TerminalCity is alphanumeric
func TestTerminalCityAlphaNumeric(t *testing.T) {
	testTerminalCityAlphaNumeric(t)
}

// BenchmarkTerminalCityAlphaNumeric benchmarks validating TerminalCity is alphanumeric
func BenchmarkTerminalCityAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testTerminalCityAlphaNumeric(b)
	}
}

// testTerminalStateAlphaNumeric validates TerminalState is alphanumeric
func testTerminalStateAlphaNumeric(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TerminalState = "®"
	err := addenda02.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestTerminalStateAlphaNumeric tests validating TerminalState is alphanumeric
func TestTerminalStateAlphaNumeric(t *testing.T) {
	testTerminalStateAlphaNumeric(t)
}

// BenchmarkTerminalStateAlphaNumeric benchmarks validating TerminalState is alphanumeric
func BenchmarkTerminalStateAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testTerminalStateAlphaNumeric(b)
	}
}

// TestAddenda02RuneCountInString validates RuneCountInString
func TestAddenda02RuneCountInString(t *testing.T) {
	addenda02 := NewAddenda02()
	var line = "702REFONEAREFTERM021000490612123456Target Store 0049          PHILADELPHIA"
	addenda02.Parse(line)

	if addenda02.ReferenceInformationOne != "" {
		t.Error("Parsed with an invalid RuneCountInString not equal to 94")
	}
}
