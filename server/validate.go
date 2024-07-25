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

package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/moov-io/ach"
)

const (
	skipAll                          = "skipAll"
	requireABAOrigin                 = "requireABAOrigin"
	bypassOrigin                     = "bypassOrigin"
	bypassOriginValidation           = "bypassOriginValidation"
	bypassDestination                = "bypassDestination"
	bypassDestinationValidation      = "bypassDestinationValidation"
	customTraceNumbers               = "customTraceNumbers"
	allowZeroBatches                 = "allowZeroBatches"
	allowMissingFileHeader           = "allowMissingFileHeader"
	allowMissingFileControl          = "allowMissingFileControl"
	bypassCompanyIdentificationMatch = "bypassCompanyIdentificationMatch"
	customReturnCodes                = "customReturnCodes"
	unequalServiceClassCode          = "unequalServiceClassCode"
	unorderedBatchNumbers            = "unorderedBatchNumbers"
	allowUnorderedBatchNumbers       = "allowUnorderedBatchNumbers"
	allowInvalidCheckDigit           = "allowInvalidCheckDigit"
	unequalAddendaCounts             = "unequalAddendaCounts"
	preserveSpaces                   = "preserveSpaces"
	allowInvalidAmounts              = "allowInvalidAmounts"
	allowZeroEntryAmount             = "allowZeroEntryAmount"
)

// readValidateOpts parses ValidateOpts from the URL query parameters and from the request body.
// A copy of the request body is returned. Callers are responsible for closing the body.
//
// Query parameters override the JSON body
func readValidateOpts(request *http.Request) (io.Reader, *ach.ValidateOpts, error) {
	validationNames := []string{
		skipAll,
		requireABAOrigin,
		bypassOrigin,
		bypassOriginValidation,
		bypassDestination,
		bypassDestinationValidation,
		customTraceNumbers,
		allowZeroBatches,
		allowMissingFileHeader,
		allowMissingFileControl,
		bypassCompanyIdentificationMatch,
		customReturnCodes,
		unequalServiceClassCode,
		unorderedBatchNumbers,
		allowUnorderedBatchNumbers,
		allowInvalidCheckDigit,
		unequalAddendaCounts,
		preserveSpaces,
		allowInvalidAmounts,
		allowZeroEntryAmount,
	}

	var buf bytes.Buffer
	bs, _ := io.ReadAll(io.TeeReader(request.Body, &buf))

	opts := &ach.ValidateOpts{}
	json.Unmarshal(bs, opts)

	for _, name := range validationNames {
		q := request.URL.Query()
		if q == nil {
			continue
		}
		input := q.Get(name)
		if input == "" {
			continue
		}

		yes, err := strconv.ParseBool(input)
		if err != nil {
			return nil, nil, fmt.Errorf("%s is an invalid boolean: %v", name, err)
		}
		switch name {
		case skipAll:
			opts.SkipAll = yes
		case requireABAOrigin:
			opts.RequireABAOrigin = yes
		case bypassOrigin, bypassOriginValidation:
			opts.BypassOriginValidation = yes
		case bypassDestination, bypassDestinationValidation:
			opts.BypassDestinationValidation = yes
		case customTraceNumbers:
			opts.CustomTraceNumbers = yes
		case allowZeroBatches:
			opts.AllowZeroBatches = yes
		case allowMissingFileHeader:
			opts.AllowMissingFileHeader = yes
		case allowMissingFileControl:
			opts.AllowMissingFileControl = yes
		case bypassCompanyIdentificationMatch:
			opts.BypassCompanyIdentificationMatch = yes
		case customReturnCodes:
			opts.CustomReturnCodes = yes
		case unequalServiceClassCode:
			opts.UnequalServiceClassCode = yes
		case unorderedBatchNumbers, allowUnorderedBatchNumbers:
			opts.AllowUnorderedBatchNumbers = yes
		case allowInvalidCheckDigit:
			opts.AllowInvalidCheckDigit = yes
		case unequalAddendaCounts:
			opts.UnequalAddendaCounts = yes
		case preserveSpaces:
			opts.PreserveSpaces = yes
		case allowInvalidAmounts:
			opts.AllowInvalidAmounts = yes
		case allowZeroEntryAmount:
			opts.AllowZeroEntryAmount = yes
		}
	}

	return &buf, opts, nil
}
