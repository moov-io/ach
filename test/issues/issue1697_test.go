package issues

import (
	"testing"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestIssue1697(t *testing.T) {
	correction := ach.NewAddenda98()
	correction.ChangeCode = "C06"
	correction.CorrectedData = "123456789 111            22"

	data := correction.ParseCorrectedData()
	require.NotNil(t, data)

	require.Equal(t, "123456789 111", data.AccountNumber)
	require.Empty(t, data.RoutingNumber)
	require.Empty(t, data.Name)
	require.Equal(t, ach.CheckingCredit, data.TransactionCode)
	require.Empty(t, data.Identification)
}
