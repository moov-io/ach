package main

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/moov-io/ach"
	"github.com/moov-io/base"
)

type NachaParseEvent struct {
	Nacha string `json:"data"`
}

type NachaParseResponse struct {
	File ach.File `json:"file"`
}

func main() {
	lambda.Start(HandleRequest)
}

// logic to be executed when lambda starts goes here
func HandleRequest(ctx context.Context, event NachaParseEvent) (NachaParseResponse, error) {

	// get NACHA file text from lambda event
	rd := strings.NewReader(event.Nacha)

	// create file from NACHA text
	file, err := ach.NewReader(rd).Read()
	if err != nil {
		return NachaParseResponse{File: file}, err
	}

	// set file ID
	file.ID = base.ID()

	// validate parsed file
	if err := file.Validate(); err != nil {
		return NachaParseResponse{File: file}, err
	}

	//create response object
	parseRes := NachaParseResponse{File: file}

	return parseRes, err
}
