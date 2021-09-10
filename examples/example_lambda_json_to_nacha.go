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

//go:build examples
// +build examples

package main

import (
	"bytes"
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/moov-io/ach"
	"github.com/moov-io/base"
)

type JsonParseEvent struct {
	Json ach.File `json:"data"`
}

func main() {
	lambda.Start(HandleRequest)
}

// logic to be executed when lambda starts goes here
func HandleRequest(ctx context.Context, event JsonParseEvent) (string, error) {

	// get file from lambda event, it has already been marshaled from json to ach.File by Go
	file := event.Json

	// set file ID
	file.ID = base.ID()

	// validate parsed file
	err := file.Validate()
	if err != nil {
		return "", err
	}

	// create buffer to contain NACHA text
	buf := new(bytes.Buffer)

	// write ach.File to buffer
	err = ach.NewWriter(buf).Write(&file)
	if err != nil {
		return buf.String(), err
	}

	// get NACHA text from buffer
	parseRes := buf.String()

	return parseRes, err
}
