package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestX(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "ppd-debit-vince.ach"))
	if err != nil {
		panic(err)
	}
	s, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

}
