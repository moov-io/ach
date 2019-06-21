moov-io/ach
===
[![GoDoc](https://godoc.org/github.com/moov-io/ach?status.svg)](https://godoc.org/github.com/moov-io/ach)
[![Build Status](https://travis-ci.com/moov-io/ach.svg?branch=master)](https://travis-ci.com/moov-io/ach)
[![Coverage Status](https://codecov.io/gh/moov-io/ach/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/ach)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/ach)](https://goreportcard.com/report/github.com/moov-io/ach)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/ach/master/LICENSE)

Package `github.com/moov-io/ach` implements a file reader and writer written in Go along with a HTTP API for creating, parsing and validating Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) files. ACH is the primary method of electronic money movement throughout the United States.

If you're looking for a complete implementation of ACH origination (file creation), OFAC checks, micro-deposits, SFTP uploading, and other featues the [moov-io/paygate](https://github.com/moov-io/paygate) project aims to be a full system for ACH transfers.

Docs: [docs.moov.io](https://docs.moov.io/en/latest/) | [api docs](https://api.moov.io/apps/ach/)

## Project Status

Moov ACH is under active development and in production for multiple companies. Please star the project if you are interested in its progress. We've built an HTTP API for ACH file creation and validation. Currently we support generating and parsing all Standard Entry Class (SEC) codes. If you have layers above ACH to simplify tasks or found bugs we would appreciate an issue or pull request. Thanks!

## Usage

The ACH project implements a Go library and HTTP server for creating and modifying ACH files. For a complete ACH origination service checkout [moov-io/paygate](https://github.com/moov-io/paygate).

### Go library

`github.com/moov-io/ach` offers a Go based ACH file reader and writer. To get started checkout a specific example:

<details>
<summary>Supported Standard Entry Class (SEC) codes</summary>

| SEC Code | Name                                  | Example                                  | Read                | Write                                            |
|----------|---------------------------------------|------------------------------------------|-----------------------------------|------------------------------------|
| ACK      | Acknowledgment Entry for CCD          | [Credit](test/ach-ack-read/ack-read.ach) | [ACK Read](test/ach-ack-read/main.go) | [ACK Write](test/ach-ack-write/main.go) |
| ADV      | Automated Accounting Advice           | [Prenote Debit](test/ach-adv-read/adv-read.ach) | [ADV Read](test/ach-adv-read/main.go) | [ADV Write](test/ach-adv-write/main.go) |
| ARC      | Accounts Receivable Entry             | [Debit](test/ach-arc-read/arc-debit.ach) | [ARC Read](test/ach-arc-read/main.go) | [ARC Write](test/ach-arc-write/main.go) |
| ATX      | Acknowledgment Entry for CTX          | [Credit](test/ach-atx-read/atx-read.ach)  | [ATX Read](test/ach-atx-read/main.go) | [ATX Write](test/ach-atx-write/main.go) |
| BOC      | Back Office Conversion                | [Debit](test/ach-boc-read/boc-debit.ach) | [BOC Read](test/ach-boc-read/main.go) | [BOC Write](test/ach-boc-write/main.go) |
| CCD      | Corporate credit or debit             | [Debit](test/ach-ccd-read/ccd-debit.ach) | [CCD Read](test/ach-ccd-read/main.go) | [CCD Write](test/ach-ccd-write/main.go) |
| CIE      | Customer-Initiated Entry              | [Credit](test/ach-cie-read/cie-credit.ach) | [CIE Read](test/ach-cie-read/main.go) | [CIE Write](test/ach-cie-write/main.go) |
| COR      | Automated Notification of Change(NOC) | [NOC](test/ach-cor-read/cor-read.ach)   | [COR Read](test/ach-cor-read/main.go) | [COR Write](test/ach-cor-write/main.go) |
| CTX      | Corporate Trade Exchange              | [Debit](test/ach-ctx-read/ctx-debit.ach) | [CTX Read](test/ach-ctx-read/main.go) | [CTX Write](test/ach-ctx-write/main.go) |
| DNE      | Death Notification Entry              | [DNE](test/ach-dne-read/dne-read.ach)   | [DNE Read](test/ach-dne-read/main.go) | [DNE Write](test/ach-dne-write/main.go) |
| ENR      | Automatic Enrollment Entry            | [ENR](test/ach-enr-read/enr-read.ach)   | [ENR Read](test/ach-enr-read/main.go) | [ENR Write](test/ach-enr-write/main.go) |
| IAT      | International ACH Transactions        | [Credit](test/ach-iat-read/iat-credit.ach) | [IAT Read](test/ach-iat-read/main.go) | [IAT Write](test/ach-iat-write/main.go) |
| MTE      | Machine Transfer Entry                | [Credit](test/ach-mte-read/mte-read.ach)   | [MTE Read](test/ach-mte-read/main.go) | [MTE Write](test/ach-mte-write/main.go) |
| POP      | Point of Purchase                     | [Debit](test/ach-pop-read/pop-debit.ach) | [POP Read](test/ach-pop-read/main.go) | [POP Write](test/ach-pop-write/main.go) |
| POS      | Point of Sale                         | [Debit](test/ach-pos-read/pos-debit.ach) | [POS Read](test/ach-pos-read/main.go) | [POS Write](test/ach-pos-write/main.go) |
| PPD      | Prearranged payment and deposits      | [Debit](test/ach-ppd-read/ppd-debit.ach) [Credit](test/ach-ppd-read/ppd-credit.ach) | [PPD Read](test/ach-ppd-read/main.go) | [PPD Write](test/ach-ppd-write/main.go) |
| RCK      | Represented Check Entries             | [Debit](test/ach-rck-read/rck-debit.ach) | [RCK Read](test/ach-rck-read/main.go) | [RCK Write](test/ach-rck-write/main.go) |
| SHR      | Shared Network Entry                  | [Debit](test/ach-shr-read/shr-debit.ach) | [SHR Read](test/ach-shr-read/main.go) | [SHR Write](test/ach-shr-write/main.go) |
| TEL      | Telephone-Initiated Entry             | [Debit](test/ach-tel-read/tel-debit.ach) | [TEL Read](test/ach-tel-read/main.go) | [TEL Write](test/ach-tel-write/main.go) |
| TRC      | Truncated Check Entry                 | [Debit](test/ach-trc-read/trc-debit.ach) | [TRC Read](test/ach-trc-read/main.go) | [TRC Write](test/ach-trc-write/main.go) |
| TRX      | Check Truncation Entries Exchange     | [Debit](test/ach-trx-read/trx-debit.ach) | [TRX Read](test/ach-trx-read/main.go) | [TRX Write](test/ach-trx-write/main.go) |
| WEB      | Internet-initiated Entries            | [Credit](test/ach-web-read/web-credit.ach) | [WEB Read](test/ach-web-read/main.go) | [WEB Write](test/ach-web-write/main.go) |
| XCK      | Destroyed Check Entry                 | [Debit](test/ach-xck-read/xck-debit.ach)  | [XCK Read](test/ach-xck-read/main.go) | [XCK Write](test/ach-xck-write/main.go) |

</details>

### HTTP API

`github.com/moov-io/ach/server` offers a HTTP and JSON API for creating and editing files. If you're using Go the `ach.File` type can be used, otherwise just send properly formatted JSON. We have an [example JSON file](test/testdata/ppd-valid.json), but each SEC type will generate differnet JSON.

Examples: [Go](examples/http/main.go) | [Ruby](https://github.com/moov-io/ruby-ach-demo)

- [Create an ACH file for a payment and get the raw file](https://github.com/moov-io/ruby-ach-demo)

## Getting Started

- [Running ACH Server](https://docs.moov.io/en/latest/tutorials/ach-server/)
- [ACH Server metrics](documentation/metrics.md)
- [Merging ACH files](https://docs.moov.io/en/latest/ach/merging-files/)

### Docker

We publish a [public docker image `moov/ach`](https://hub.docker.com/r/moov/ach/tags) on Docker Hub with each tagged release of ACH. No configuration is required to serve on `:8080` and metrics at `:9090/metrics` in Prometheus format.

```
$ docker run -p 8080:8080 -p 9090:9090 moov/ach:latest
ts=2019-06-20T23:58:44.4931106Z caller=main.go:75 startup="Starting ach server version v1.0.2"
ts=2019-06-20T23:58:44.5010238Z caller=main.go:135 transport=HTTP addr=:8080
ts=2019-06-20T23:58:44.5018409Z caller=main.go:125 admin="listening on :9090"

$ curl localhost:8080/files
{"files":[],"error":null}
```

### From Source

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and thus requires Go 1.11+. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/ach/releases) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/ach.git

# Pull down into the Go Module cache
$ go get -u github.com/moov-io/ach

$ go doc github.com/moov-io/ach BatchHeader
```

### Configuration

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `ACH_FILE_TTL` | Time to live (TTL) for `*ach.File` objects stored in the in-memory repository. | 0 = No TTL / Never delete files (Example: `240m`) |

Note: By design ACH **does not persist** (save) any data about the files, batches or entry details created. The only storage occurs in memory of the process and upon restart ACH will have no files, batches, or data saved. Also, no in memory encryption of the data is performed.

## Getting Help

If you have ACH specific questions NACHA (National Automated Clearing House Association) has their [complete specification](documentation/2013-Corporate-Rules-and-Guidelines.pdf) for all file formats and message types.

 channel | info
 ------- | -------
 [Project Documentation](https://docs.moov.io/en/latest/) | Our project documentation available online.
 Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce an problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](http://moov-io.slack.com/) | Join our slack channel to have an interactive discussion about the development of the project. [Request an invite to the slack channel](https://join.slack.com/t/moov-io/shared_invite/enQtNDE5NzIwNTYxODEwLTRkYTcyZDI5ZTlkZWRjMzlhMWVhMGZlOTZiOTk4MmM3MmRhZDY4OTJiMDVjOTE2MGEyNWYzYzY1MGMyMThiZjg)

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows
- Rasberry Pi

Note: 32-bit platforms have known issues and are not supported.

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) to get started!

Note: This project uses Go Modules, which requires Go 1.11 or higher, but we ship the vendor directory in our repository.

### Releasing

To make a release of ach simply open a pull request with `CHANGELOG.md` and `version.go` updated with the next version number and details. You'll also need to push the tag (i.e. `git push origin v1.0.0`) to origin in order for CI to make the release.

### Fuzzing

We currently run fuzzing over ACH in the form of a [`moov/achfuzz`](https://hub.docker.com/r/moov/achfuzz) Docker image. You can [read more](./test/fuzz-reader/README.md) or run the image and report crasher examples to [`security@moov.io`](mailto:security@moov.io). Thanks!

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
