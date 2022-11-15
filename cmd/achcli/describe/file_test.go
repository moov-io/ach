// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package describe

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatAmount(t *testing.T) {
	require.Equal(t, "12345", formatAmount(false, 12345))
	require.Equal(t, "123.45", formatAmount(true, 12345))
	require.Equal(t, "1,234,567.89", formatAmount(true, 123456789))
}

func TestMaskNumber(t *testing.T) {
	require.Equal(t, "****", maskNumber(""))
	require.Equal(t, "****", maskNumber("1"))
	require.Equal(t, "****", maskNumber("12"))
	require.Equal(t, "****", maskNumber("123"))
	require.Equal(t, "****", maskNumber("1234"))
	require.Equal(t, "*2345", maskNumber("12345"))
	require.Equal(t, "**3456", maskNumber("123456"))
	require.Equal(t, "***4567", maskNumber("1234567"))
	require.Equal(t, "****5678", maskNumber("12345678"))
	require.Equal(t, "*****6789", maskNumber("123456789"))
	require.Equal(t, "******7890", maskNumber("1234567890"))
}

func TestMaskName(t *testing.T) {
	require.Equal(t, "", maskName(""))
	require.Equal(t, "* * *", maskName("a a a"))
	require.Equal(t, "* * *", maskName(" a  a  a "))
	require.Equal(t, "Jo** ***", maskName("John Doe"))
	require.Equal(t, "Jo** Sm*** **", maskName("John Smith Jr"))
	require.Equal(t, "Al******* Lo********** ** ***", maskName("Alexander Longnameiton Jr III"))
}
