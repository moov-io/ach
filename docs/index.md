# Moov ACH
<!-- Place this tag where you want the button to render. -->
<a class="github-button" href="https://github.com/moov-io/ach" data-size="large" data-show-count="true" aria-label="Star moov-io/ach on GitHub">moov-io/ach</a>
<a href="https://godoc.org/github.com/moov-io/ach"><img src="https://godoc.org/github.com/moov-io/ach?status.svg" /></a>
<a href="https://raw.githubusercontent.com/moov-io/ach/master/LICENSE"><img src="https://img.shields.io/badge/license-Apache2-blue.svg" /></a>

Moov ACH implements a low level Automated Clearing House (ACH) interface for parsing, creating, validating, and merging ACH files. Moov ACH exposes an HTTP API for REST based interaction. Any language which can use HTTP and JSON can leverage the ACH Server. The API's endpoints expose both text and JSON to easily ingest or export either format.

## Running Moov ACH Server

Moov ACH can be deployed in multiple scenarios.

- <a href="#binary-distribution">Binary Distributions</a> are released with every versioned release. Frequently added to the VM/AMI build script for the application needing Moov ACH.
- A <a href="#docker-container">Docker container</a> is built and added to Docker Hub with every versioned released.
- Our hosted [api.moov.io](https://api.moov.io) is updated with every versioned release. Our Kubernetes example is what Moov utilizes in our production environment.

### Binary Distribution

Download the [latest Moov ACH server release](https://github.com/moov-io/ach/releases/latest) for your operating system and run it from a terminal.

For macOS / Linux users:

```sh
$ ./ach-darwin-amd64
ts=2019-06-20T23:23:44.870717Z caller=main.go:75 startup="Starting ach server version v1.3.1"
ts=2019-06-20T23:23:44.871623Z caller=main.go:135 transport=HTTP addr=:8080
ts=2019-06-20T23:23:44.871692Z caller=main.go:125 admin="listening on :9090"
```

For Windows users:

```
> ach.exe
ts=2019-06-20T23:23:44.870717Z caller=main.go:75 startup="Starting ach server version v1.3.1"
ts=2019-06-20T23:23:44.871623Z caller=main.go:135 transport=HTTP addr=:8080
ts=2019-06-20T23:23:44.871692Z caller=main.go:125 admin="listening on :9090
```

Next [Connect to Moov ACH](#connecting-to-moov-ach)

### Docker Container

Moov ACH is dependent on Docker being properly installed and running on your machine. Ensure that Docker is running. If your Docker client has issues connecting to the service review the [Docker getting started guide](https://docs.docker.com/get-started/) if you have any issues.

```sh
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
```

Execute the Docker run command

```sh
$ docker run -p 8080:8080 -p 9090:9090 moov/ach:latest
ts=2019-06-21T17:03:23.782592Z caller=main.go:69 startup="Starting ach server version v1.3.1"
ts=2019-06-21T17:03:23.78314Z caller=main.go:129 transport=HTTP addr=:8080
ts=2019-06-21T17:03:23.783252Z caller=main.go:119 admin="listening on :9090"
```

Next [Connect to Moov ACH](#connecting-to-moov-ach)

## Command Line

On each release there's a `achcli` utility released. This tool can display ACH files in a human-readable format which is easier to read than their plaintext format.

```
$ wget -O achcli https://github.com/moov-io/ach/releases/download/v1.4.0/achcli-darwin-amd64 && chmod +x achcli

$ achcli test/testdata/ppd-debit.ach
Describing ACH file 'test/testdata/ppd-debit.ach'

  Origin     OriginName    Destination  DestinationName       FileCreationDate  FileCreationTime
  121042882  My Bank Name  231380104    Federal Reserve Bank  190624            0000

  BatchCount  BlockCount  EntryAddendaCount  TotalDebitAmount  TotalCreditAmount
  1           1           1                  100000000         0

  BatchNumber  Type  ServiceClass  Addendas  TotalDebits  TotalCredits
  1            PPD   225           1         100000000    0

    TxCode  AccountNumber      Amount     Name                    TraceNumber      Category
    27      12345678           100000000  Receiver Account Name   121042880000001
```

## Connecting to Moov ACH

The Moov ACH service will be running on port `8080` (with an admin port on `9090`).

Confirm that the service is running by issuing the following command or simply visiting the url in your browser [localhost:8080/ping](http://localhost:8080/ping)

```bash
$ curl http://localhost:8080/ping
PONG

$ curl http://localhost:8080/files
{"files":[],"error":null}
```

You can also send [an example PPD ACH file we have](https://github.com/moov-io/ach/blob/master/test/testdata/ppd-valid.json) to any ACH service or read through HTTP examples in [Ruby](https://github.com/moov-io/ruby-ach-demo) and [Go](https://github.com/moov-io/ach/blob/master/examples/http/main.go).

### API documentation

See our [API documentation](https://moov-io.github.io/ach/) for Moov ACH endpoints.

### ACH Admin Port

The port `9090` is bound by ACH for our admin service. This HTTP server has endpoints for Prometheus metrics (`GET /metrics`), readiness (`GET /ready`) and liveness checks (`GET /live`).

## Getting Help

 channel | info
 ------- | -------
 [Project Documentation](https://docs.moov.io/) | Our project documentation available online.
 Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`#ach`) to have an interactive discussion about the development of the project.
