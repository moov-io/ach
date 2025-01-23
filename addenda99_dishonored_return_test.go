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

	"github.com/stretchr/testify/require"
)

func mockAddenda99Dishonored() *Addenda99Dishonored {
	addenda99 := NewAddenda99Dishonored()
	addenda99.DishonoredReturnReasonCode = "R68"
	addenda99.OriginalEntryTraceNumber = "059999990000301"
	addenda99.OriginalReceivingDFIIdentification = "12391871"
	addenda99.ReturnTraceNumber = "123918710000001"
	addenda99.ReturnSettlementDate = "179"
	addenda99.ReturnReasonCode = "01"
	addenda99.AddendaInformation = "Untimely Return"
	addenda99.TraceNumber = "059999990000001"
	return addenda99
}

func TestAddenda99Dishonored(t *testing.T) {
	addenda99Dishonored := NewAddenda99Dishonored()
	line := "799R6909100001137143222042712114530   1211453000251201170506                   091000011371432"
	addenda99Dishonored.Parse(line)

	expected := &Addenda99Dishonored{
		TypeCode:                           "99",
		DishonoredReturnReasonCode:         "R69",
		OriginalEntryTraceNumber:           "091000011371432",
		OriginalReceivingDFIIdentification: "12114530",
		ReturnTraceNumber:                  "121145300025120",
		ReturnSettlementDate:               "117",
		ReturnReasonCode:                   "05",
		AddendaInformation:                 "06                   ",
		TraceNumber:                        "091000011371432",
	}
	require.Equal(t, expected, addenda99Dishonored)

	addenda99 := NewAddenda99()
	addenda99.ReturnCode = "R69"
	addenda99.OriginalTrace = "091000011371432"
	addenda99.DateOfDeath = "220427"
	addenda99.OriginalDFI = "12114530"
	addenda99.SetDishonoredAddendaInformation("121145300025120", "117", "R05", "06")
	addenda99.TraceNumber = "091000011371432"
	require.Equal(t, line, addenda99.String())

	require.Equal(t, "121145300025120", addenda99.AddendaInformationReturnTraceNumber())
	require.Equal(t, "117", addenda99.AddendaInformationReturnSettlementDate())
	require.Equal(t, "R05", addenda99.AddendaInformationReturnReasonCode())
	require.Equal(t, "06                   ", addenda99.AddendaInformationExtra())
}

func TestAddenda99Dishonored_Example(t *testing.T) {
	addenda99Dishonored := NewAddenda99Dishonored()
	line := "799R68059999990000301      12391871   12391871000000117901                     059999990000001"
	addenda99Dishonored.Parse(line)

	expected := &Addenda99Dishonored{
		TypeCode:                           "99",
		DishonoredReturnReasonCode:         "R68",
		OriginalEntryTraceNumber:           "059999990000301",
		OriginalReceivingDFIIdentification: "12391871",
		ReturnTraceNumber:                  "123918710000001",
		ReturnSettlementDate:               "179",
		ReturnReasonCode:                   "01",
		AddendaInformation:                 "                     ",
		TraceNumber:                        "059999990000001",
	}
	require.Equal(t, expected, addenda99Dishonored)
}

func TestAddenda99Dishonored__Fields(t *testing.T) {
	addenda99 := mockAddenda99Dishonored()

	// shorten some fields
	addenda99.OriginalEntryTraceNumber = "0599999900301"
	addenda99.ReturnTraceNumber = "123918710001"
	addenda99.TraceNumber = "05999900001"

	require.Equal(t, "R68", addenda99.DishonoredReturnReasonCodeField())
	require.Equal(t, "000599999900301", addenda99.OriginalEntryTraceNumberField())
	require.Equal(t, "12391871", addenda99.OriginalReceivingDFIIdentificationField())
	require.Equal(t, "000123918710001", addenda99.ReturnTraceNumberField())
	require.Equal(t, "179", addenda99.ReturnSettlementDateField())
	require.Equal(t, "01", addenda99.ReturnReasonCodeField())
	require.Equal(t, "Untimely Return      ", addenda99.AddendaInformationField())
	require.Equal(t, "000005999900001", addenda99.TraceNumberField())

	line := addenda99.String()
	dishonored := NewAddenda99Dishonored()
	dishonored.Parse(line)

	require.Equal(t, "000599999900301", dishonored.OriginalEntryTraceNumberField())
	require.Equal(t, "12391871", dishonored.OriginalReceivingDFIIdentificationField())
	require.Equal(t, "000123918710001", dishonored.ReturnTraceNumberField())
	require.Equal(t, "179", dishonored.ReturnSettlementDateField())
	require.Equal(t, "01", dishonored.ReturnReasonCodeField())
	require.Equal(t, "Untimely Return      ", dishonored.AddendaInformationField())
	require.Equal(t, "000005999900001", dishonored.TraceNumberField())
}
