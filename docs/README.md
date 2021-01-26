## Moov ACH

**[GitHub](https://github.com/moov-io/ach)** | **[Running](https://github.com/moov-io/ach#usage)** | **[Configuration](https://github.com/moov-io/ach#configuration-settings)**

### Purpose

Moov ACH implements a file reader and writer written in Go along with an HTTP API for creating, parsing, and validating Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) files. ACH is the primary method of electronic money movement throughout the United States.

### API Documentation

See our [API documentation](https://moov-io.github.io/ach/api/) for Moov ACH endpoints.

### Table of Contents

1. [Introduction](./intro.md)
1. [Create File](./create-file.md)
1. [File Structure](./file-structure.md)
1. [SEC Codes Table](./sec-codes-table.md)

**File Operations**

1. [Custom Validation](./custom-validation.md)
1. [Balanced Offset](./balanced-offset.md)
1. [Flatten Batches](./flatten-batches.md)
1. [Merging Files](./merging-files.md)
1. [Segment a File](./segment-file.md)
1. [Return Files](./returns.md)

**Running ACH**

1. Functions as a Service
   1. [Building for AWS Lambda](./aws/lambda/building-go-for-lambda.md)
      1. [Convert JSON to NACHA format](./aws/lambda/json_to_nacha/lambda_json_to_nacha.md)
      1. [Convert NACHA To JSON format](./aws/lambda/nacha_to_json/lambda_nacha_to_json.md)
   1. [Docker image](https://github.com/moov-io/ach#docker)
   1. [Google Cloud Run Button](https://github.com/moov-io/ach#google-cloud-run-button)

**Production and Monitoring**

1. [Deployment](./deployment.md)
1. [Metrics](./metrics.md)

## Getting Help

If you have ACH-specific questions, NACHA (National Automated Clearing House Association) has their [complete specification](docs/2013-Corporate-Rules-and-Guidelines.pdf) for all file formats and message types.

 channel | info
 ------- | -------
 [Project Documentation](#moov-ach) | Our project documentation available online.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/ach/issues/new) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel to have an interactive discussion about the development of the project.
