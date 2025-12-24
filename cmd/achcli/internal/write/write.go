package write

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/internal/read"
)

func File(w io.Writer, file *ach.File, format read.Format) error {
	switch format {
	case read.FormatNacha:
		return ach.NewWriter(w).Write(file)
	case read.FormatJSON:
		return json.NewEncoder(w).Encode(file)
	}
	return fmt.Errorf("unknown format %v", format)
}
