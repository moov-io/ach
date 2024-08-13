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

package issues

import (
	"testing"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestIssue1462(t *testing.T) {
	// "For Entries originated in a fixed-value amount, the Foreign Exchange
	//  Reference Field will be space filled."
	//
	// ForeignExchangeReferenceIndicator is also a Required field, which according
	// to the Nacha specification allows ODFI's to leave this value as blank.
	//
	// In order to strike a balance between ODFI/RDFI processing we're going to
	// allow ForeignExchangeReferenceIndicator = 0 when ForeignExchangeIndicator == "FF".
	//
	// https://github.com/moov-io/ach/issues/1462

	bh := ach.NewIATBatchHeader()
	bh.ServiceClassCode = ach.CreditsOnly
	bh.IATIndicator = ach.IATCOR
	bh.ForeignExchangeIndicator = "FF"
	bh.ForeignExchangeReferenceIndicator = 0 // parsed as "" or " "
	bh.ISODestinationCountryCode = "US"
	bh.OriginatorIdentification = "123456789"
	bh.StandardEntryClassCode = ach.COR
	bh.CompanyEntryDescription = "TRADEPAYMT"
	bh.ISOOriginatingCurrencyCode = "CAD"
	bh.ISODestinationCurrencyCode = "USD"
	bh.ODFIIdentification = "91012980"
	require.NoError(t, bh.Validate())

	// Don't allow it for FV or VF
	bh.ForeignExchangeIndicator = "FV"
	require.ErrorContains(t, bh.Validate(), "ForeignExchangeReferenceIndicator 0 is a required field")
	bh.ForeignExchangeIndicator = "VF"
	require.ErrorContains(t, bh.Validate(), "ForeignExchangeReferenceIndicator 0 is a required field")
}
