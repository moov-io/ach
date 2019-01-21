// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// +build ignore

// Generates iso4217.go.
//
// This file grabs the ISO 4217 currency codes and writes them
// into source code so we don't rely on any external files (zip,
// json, etc).
//
// The data is pulled from datahub.io as the ISO.org site only offers
// XML.
//
// https://datahub.io/core/currency-codes
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"path/filepath"
	"runtime"
	"time"
)

var (
	downloadUrl    = "https://datahub.io/core/currency-codes/r/codes-all.json"
	outputFilename = filepath.Join("internal", "iso4217", "iso4217.go")
)

// {"AlphabeticCode": "AFN", "Currency": "Afghani", ... }
type currency struct {
	Code string `json:"AlphabeticCode"`
	Name string `json:"Currency"`
}

func main() {
	when := time.Now().Format("2006-01-02T03:04:05Z")
	who, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to get user on %s", runtime.GOOS)
	}

	// Write copyright header
	var buf bytes.Buffer
	fmt.Fprintf(&buf, `// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Generated on %s by %s, any modifications will be overwritten
package iso4217
`, when, who.Username)

	// Download certs
	resp, err := http.Get(downloadUrl)
	if err != nil {
		log.Fatalf("error while downloading %s: %v", downloadUrl, err)
	}
	defer resp.Body.Close()

	var currencies []currency
	if err := json.NewDecoder(resp.Body).Decode(&currencies); err != nil {
		log.Fatalf("error while parsing currency response: %v", err)
	}

	// The JSON file contains duplicates so we need to dedup them..
	cs := make(map[string]bool, 150)

	// Write countries to source code
	fmt.Fprintln(&buf, "var currencyCodes = map[string]bool{")
	for i := range currencies {
		code, name := currencies[i].Code, currencies[i].Name
		if code == "" || name == "" {
			fmt.Printf("SKIPPING: code=%s currency=%s\n", code, name)
			continue
		}
		if _, exists := cs[code]; !exists {
			cs[code] = true // mark as seen
			fmt.Fprintf(&buf, fmt.Sprintf(`"%s": true, // %s`+"\n", code, name))
		}
	}
	fmt.Fprintln(&buf, "}")

	// format source code and write file
	out, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(buf.String())
		log.Fatalf("error formatting output code, err=%v", err)
	}

	err = ioutil.WriteFile(outputFilename, out, 0644)
	if err != nil {
		log.Fatalf("error writing file, err=%v", err)
	}
}
