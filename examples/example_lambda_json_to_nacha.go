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

	// get file from lambda event, it has already been marshalled from json to ach.File by Go
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
