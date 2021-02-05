---
layout: page
title: AWS JSON to NACHA

hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Use AWS Lambda to parse JSON to a NACHA-formatted string
This walkthrough provides instructions on how to use the [`github.com/moov-io/ach`](https://pkg.go.dev/github.com/moov-io/ach) Go library to
parse a JSON object in a Lambda event and return the resulting NACHA-formatted string.

The Lambda event could be triggered by a variety of sources: upload to S3 bucket, API Gateway HTTP request, SQS message, another Lambda, etc. Likewise, the response can be sent to a variety of destinations. This walkthrough doesn't demonstrate a particular trigger or destination, but it can easily be adapted for the use cases mentioned.

Make sure you've built [your project](../building-go-for-lambda.md) for AWS Lambda.


## 1. Create the Go file
Create a new Go file named `main.go` and replace it's contents with the following:
```go
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
```

*`main()`* is invoked when the lambda is triggered <br/>
*`HandleRequest()`* is the callback containing the business logic and accepts a lambda event as a parameter with the shape:
```
{
    "data": { ...JSON representing ACH file }
}
```
`JsonParseEvent` mirrors this shape and makes the contents accessible in `HandleRequest()`. This event can be changed to fit
your needs, all you need to do is update the `JsonParseEvent` struct and make sure the trigger for this lambda is passing in the expected event.

<br/><br/>
## 2. Build your file for AWS Lambda
Now that you've saved your Go file, you need to build it for use with AWS Lambdas. See [Build a Go file for AWS Lambda](../building-go-for-lambda.md) for more details.
<br /> <br />

## 3. Create the Lambda function in AWS console
Sign into the AWS Console and head over to the Lambdas section and create a new function.

Select *Author From Scratch*

<img src="/ach/aws/lambda/json_to_nacha/create-function-author-options.PNG" alt="Authorship Details" width="700px"><br /><br />

Under Basic Information, enter a name for your function (e.g. `parse-json-to-nacha`) and select Go as the Runtime

<img src="/ach/aws/lambda/json_to_nacha/create-function-basic-info.PNG" alt="Basic Info" width="700px"><br /><br />

By default, AWS will create a new permissions role for your function with all lambda permissions. This is adequate for this tutorial, but if you plan to access other AWS services from this function you will need to add permissions. Click *Create Function* in the bottom right after you've selected your desired role.

<img src="/ach/aws/lambda/json_to_nacha/create-function-permissions-role.PNG" alt="Permissions Role Selection" width="700px"><br /><br />

## 4. Configure function and upload executable
AWS Lambdas don't currently support inline editing of Go files so you will need to upload the zip file you created in step 2. Click the *Actions* dropdown in the Function Code section, select *Upload a .zip file*, upload your zip and click *Save*.

<img src="/ach/aws/lambda/json_to_nacha/configure-function-upload-zip.PNG" alt="Upload Zip" width="700px"><br /><br />

Now scroll down to the Basic Settings section. Handler needs to be set to the name of your executable file from step 2. In our case, the executable name is `main`. Click *Edit* and change the Handler from `hello` to `main`.

<img src="/ach/aws/lambda/json_to_nacha/configure-function-basic-settings.PNG" alt="Edit Basic Settings" width="700px"><br /><br />

Finally, we will create two test events to confirm our function is working as expected.

The first event will be used to confirm the JSON is parsed correctly. Click the dropdown to the left of the Test button at the top of the page and select *Configure test events*.

<img src="/ach/aws/lambda/json_to_nacha/configure-function-create-parseEvent.PNG" alt="Create parseEvent" width="700px"><br /><br />

Leave the `hello-world` template selected and enter a name for your event (e.g. `parseEvent`).

<img src="/ach/aws/lambda/json_to_nacha/configure-function-configure-parseEvent.PNG" alt="Configure parseEvent" width="700px"><br /><br />

Replace the contents of the code editor with the following:

```json
{
  "data": {
    "id": "",
    "fileHeader": {
      "id": "",
      "immediateDestination": "031300012",
      "immediateOrigin": "231380104",
      "fileCreationDate": "190816",
      "fileCreationTime": "1055",
      "fileIDModifier": "A",
      "immediateDestinationName": "Federal Reserve Bank",
      "immediateOriginName": "My Bank Name",
      "referenceCode": "12345678"
    },
    "batches": [
      {
        "batchHeader": {
          "id": "",
          "serviceClassCode": 225,
          "companyName": "Name on Account",
          "companyIdentification": "231380104",
          "standardEntryClassCode": "PPD",
          "companyEntryDescription": "REG.SALARY",
          "effectiveEntryDate": "190816",
          "originatorStatusCode": 1,
          "ODFIIdentification": "12104288",
          "batchNumber": 1
        },
        "entryDetails": [
          {
            "id": "",
            "transactionCode": 27,
            "RDFIIdentification": "23138010",
            "checkDigit": "4",
            "DFIAccountNumber": "123456789        ",
            "amount": 200000000,
            "identificationNumber": "               ",
            "individualName": "Debit Account         ",
            "discretionaryData": "  ",
            "traceNumber": "121042880000001"
          }
        ],
        "batchControl": {
          "id": "",
          "serviceClassCode": 225,
          "entryAddendaCount": 1,
          "entryHash": 23138010,
          "totalDebit": 200000000,
          "totalCredit": 0,
          "companyIdentification": "231380104",
          "ODFIIdentification": "12104288",
          "batchNumber": 1
        },
        "offset": null
      }
    ],
    "IATBatches": null,
    "fileControl": {
      "id": "",
      "batchCount": 1,
      "blockCount": 1,
      "entryAddendaCount": 1,
      "entryHash": 23138010,
      "totalDebit": 200000000,
      "totalCredit": 0
    },
    "fileADVControl": {
      "id": "",
      "batchCount": 0,
      "entryAddendaCount": 0,
      "entryHash": 0,
      "totalDebit": 0,
      "totalCredit": 0
    },
    "NotificationOfChange": null,
    "ReturnEntries": null
  }
}
```

This JSON will be handled by `HandleRequest()`. Note the top level `data` field which matches the shape we covered in step 1. Click *Create* at the bottom when you're finished making changes.

Now create your second event, which is designed to cause an error while parsing so you can see how the function returns errors from Go code.  Create a new event with a different name (e.g. `parseEventError`) and modify the JSON.

Change the JSON above by setting immediateDestination to be an empty string. This should result in an error stating that the immediateDestination should be 9 numeric digits.

```
{
  "data": {
    "id": "",
    "fileHeader": {
      "id": "",
      "immediateDestination": "",
	  ...
}
```

## 5. Run/Test your function
Now you can finally test your function.

Go to the top of the page, select your successful test event (e.g. `parseEvent`) and click *Test*. You should receive a succcessful response similar to the following:

<img src="/ach/aws/lambda/json_to_nacha/test-function-success-parseEvent.PNG" alt="Test parseEvent" width="700px"><br /><br />

Now test your error event (e.g. `parseEventError`). You should receive an error response similar to the following:

<img src="/ach/aws/lambda/json_to_nacha/test-function-error-parseEventError.PNG" alt="Test parseEventError" width="700px"><br /><br />

And that's it! You've successfully created and tested a lambda function that parses JSON into a NACHA-formatted string.
