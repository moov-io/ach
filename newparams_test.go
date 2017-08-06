package ach

import (
	"testing"
)

func TestFileParam(t *testing.T) {
	f := NewFile(
		FileParam{ImmediateDestination: "081000032",
			ImmediateOrigin:          "123456789",
			ImmediateDestinationName: "Your Bank",
			ImmediateOriginName:      "Your Company Inc",
			ReferenceCode:            "A00000"})
	if err := f.Header.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if f.Header.ImmediateOriginName != "Your Company Inc" {
		t.Errorf("FileParam value was not copied to file.Header")
	}
}
