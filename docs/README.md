## Moov ACH

**GitHub** | **[Running](./index.md#running-moov-ach-server)** | **[Configuration](https://github.com/moov-io/ach#configuration)**

### Purpose

Moov ACH implements a file reader and writer written in Go along with a HTTP API for creating, parsing and validating Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) files. ACH is the primary method of electronic money movement throughout the United States.

### API documentation

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
1. [Segment-File](./segment-file.md)

**Production and Monitoring**

1. [Deployment](./deployment.md)
1. [Metrics](./metrics.md)

## Getting Help

If you have ACH specific questions NACHA (National Automated Clearing House Association) has their [complete specification](docs/2013-Corporate-Rules-and-Guidelines.pdf) for all file formats and message types.

 channel | info
 ------- | -------
 [Project Documentation](./docs/#moov-ach) | Our project documentation available online.
 Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/ach/issues/new) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel to have an interactive discussion about the development of the project.
