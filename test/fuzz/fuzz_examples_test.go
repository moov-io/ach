package main

import (
	"testing"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestFuzzCrashers(t *testing.T) {
	t.Run("read json", func(t *testing.T) {
		input := []string{
			`{"BAtChes":[{"entrYDetAils":[null]}]}`,
			`{"BAtChes":[{"entrYDetAils":[{"AddendA05":[null]}]}]}`,
		}
		for i := range input {
			require.NotPanics(t, func() {
				ach.FileFromJSON([]byte(input[i]))
			})
		}
	})
}
