// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package describe

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCodeMap(t *testing.T) {
	// Entries
	require.Equal(t, "(Checking Credit)", entry("Checking", credit))
	require.Equal(t, "(Checking Debit) ", entry("Checking", debit))
	require.Equal(t, "(Savings Credit) ", entry("Savings", credit))
	require.Equal(t, "(Savings Debit)  ", entry("Savings", debit))

	// Notice of Change
	require.Equal(t, "(Checking Return NOC Debit) ", noc("Checking", debit))
	require.Equal(t, "(Checking Return NOC Credit)", noc("Checking", credit))
	require.Equal(t, "(Savings Return NOC Debit)  ", noc("Savings", debit))
	require.Equal(t, "(Savings Return NOC Credit) ", noc("Savings", credit))

	// pre-notifications
	require.Equal(t, "(Checking Prenote Debit) ", prenote("Checking", debit))
	require.Equal(t, "(Checking Prenote Credit)", prenote("Checking", credit))
	require.Equal(t, "(Savings Prenote Debit)  ", prenote("Savings", debit))
	require.Equal(t, "(Savings Prenote Credit) ", prenote("Savings", credit))

	// Remittance
	require.Equal(t, "(Checking Zero Dollar Remittance Debit) ", remittance("Checking", debit))
	require.Equal(t, "(Checking Zero Dollar Remittance Credit)", remittance("Checking", credit))
	require.Equal(t, "(Savings Zero Dollar Remittance Debit)  ", remittance("Savings", debit))
	require.Equal(t, "(Savings Zero Dollar Remittance Credit) ", remittance("Savings", credit))
}
