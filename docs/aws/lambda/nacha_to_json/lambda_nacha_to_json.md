---
layout: page
title: AWS NACHA to JSON

hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Use AWS Lambda to parse a NACHA-formatted string to JSON
This walkthrough provides instructions on how to use the [`github.com/moov-io/ach`](https://pkg.go.dev/github.com/moov-io/ach) Go library to
parse a NACHA-formatted string in a Lambda event and return the resulting JSON response.

The Lambda event could be triggered by a variety of sources: upload to S3 bucket, API Gateway HTTP request, SQS message, another Lambda, etc. Likewise, the response can be sent to a variety of destinations. This walkthrough doesn't demonstrate a particular trigger or destination, but it can easily be adapted for the use cases mentioned.

Make sure you've built [your project](../building-go-for-lambda.md) for AWS Lambda.

## 1. Create the Go file
Create a new Go file named `main.go` and replace it's contents with the following:
```go
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
```

*`main()`* is invoked when the lambda is triggered <br/>
*`HandleRequest()`* is the callback containing the business logic and accepts a lambda event as a parameter with the shape:
```
{
    "data": " ...NACHA formatted string "
}
```
`NachaParseEvent` mirrors this shape and makes the contents accessible in `HandleRequest()`. This event can be changed to fit
your needs, all you need to do is update the `NachaParseEvent` struct and make sure the trigger for this lambda is passing in the expected event.

<br/><br/>
## 2. Build your file for AWS Lambda
Now that you've saved your Go file, you need to build it for use with AWS Lambdas. See [Build a Go file for AWS Lambda](../building-go-for-lambda.md) for more details.
<br /> <br />

## 3. Create the Lambda function in AWS console
Sign into the AWS Console and head over to the Lambdas section and create a new function.

Select *Author From Scratch*

<img src="/ach/aws/lambda/nacha_to_json/create-function-author-options.PNG" alt="Authorship Details" width="700px"><br /><br />

Under Basic Information, enter a name for your function (e.g. `parse-nacha-to-json`) and select Go as the Runtime

<img src="/ach/aws/lambda/nacha_to_json/create-function-basic-info.PNG" alt="Basic Info" width="700px"><br /><br />

By default, AWS will create a new permissions role for your function with all lambda permissions. This is adequate for this tutorial, but if you plan to access other AWS services from this function you will need to add permissions. Click *Create Function* in the bottom right after you've selected your desired role.

<img src="/ach/aws/lambda/nacha_to_json/create-function-permissions-role.PNG" alt="Permissions Role Selection" width="700px"><br /><br />

## 4. Configure function and upload executable
AWS Lambdas don't currently support inline editing of Go files so you will need to upload the zip file you created in step 2. Click the *Actions* dropdown in the Function Code section, select *Upload a .zip file*, upload your zip and click *Save*.

<img src="/ach/aws/lambda/nacha_to_json/configure-function-upload-zip.PNG" alt="Upload Zip" width="700px"><br /><br />

Now scroll down to the Basic Settings section. Handler needs to be set to the name of your executable file from step 2. In our case, the executable name is `main`. Click *Edit* and change the Handler from `hello` to `main`.

<img src="/ach/aws/lambda/nacha_to_json/configure-function-basic-settings.PNG" alt="Edit Basic Settings" width="700px"><br /><br />

Finally, we will create two test events to confirm our function is working as expected.

The first event will be used to confirm the NACHA-formatted string is parsed correctly. Click the dropdown to the left of the Test button at the top of the page and select *Configure test events*.

<img src="/ach/aws/lambda/nacha_to_json/configure-function-create-parseEvent.PNG" alt="Create parseEvent" width="700px"><br /><br />

Leave the `hello-world` template selected and enter a name for your event (e.g. `parseEvent`).

<img src="/ach/aws/lambda/nacha_to_json/configure-function-configure-parseEvent.PNG" alt="Configure parseEvent" width="700px"><br /><br />

Replace the contents of the code editor with the following:

```json
{
  "data": "101 03130001202313801041908161055A094101Federal Reserve Bank   My Bank Name           12345678\n5225Name on Account                     231380104 PPDREG.SALARY      190816   1121042880000001\n627231380104123456789        0200000000               Debit Account           0121042880000001\n82250000010023138010000200000000000000000000231380104                          121042880000001\n9000001000001000000010023138010000200000000000000000000                                       \n9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999\n9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999\n9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999\n9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999\n9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999"
}
```

This JSON will be handled by `HandleRequest()`. Note the top level `data` field which matches the shape we covered in step 1. Click *Create* at the bottom when you're finished making changes.



Now create your second event, which is designed to cause an error while parsing so you can see how the function returns errors from Go code. Create a new event with a different name (e.g. `parseEventError`) and modify the NACHA string.

You can change the string above by putting a '1' in position 4 (`1011...`). This should result in an error stating that the immediateDestination should be 9 numeric digits.

```
{
  "data": "101103130001202313801041908161055A094101Federal Reserve..."
}
```
<br /><br />
## 5. Run/Test your function
Now you can finally test your function.

Go to the top of the page, select your successful test event (e.g. `parseEvent`) and click *Test*. You should receive a succcessful response similar to the following:

<img src="/ach/aws/lambda/nacha_to_json/test-function-success-parseEvent.PNG" alt="Test parseEvent" width="700px"><br /><br />

Now test your error event (e.g. `parseEventError`). You should receive an error response similar to the following:

<img src="/ach/aws/lambda/nacha_to_json/test-function-error-parseEventError.PNG" alt="Test parseEventError" width="700px"><br /><br />

And that's it! You've successfully created and tested a lambda function that parses a NACHA-formatted string into JSON.
