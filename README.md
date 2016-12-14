wadearnold/ach
===
[![GoDoc](https://godoc.org/github.com/wadearnold/ach?status.svg)](https://godoc.org/github.com/wadearnold/ach)
[![Build Status](https://travis-ci.org/wadearnold/ach.svg?branch=master)](https://travis-ci.org/wadearnold/ach)
[![Coverage Status](https://coveralls.io/repos/github/wadearnold/ach/badge.svg?branch=master)](https://coveralls.io/github/wadearnold/ach?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/wadearnold/ach)](https://goreportcard.com/report/github.com/wadearnold/ach)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/wadearnold/ach/master/LICENSE)


Package 'wadearnold/ach' implements a file decoder and encoder for parsing [ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House
) Automated Clearing House files. ACH is the primary method of electronic money movement throughout the United States.

Their are several kinds of ACH files PPD, PPD+, CCD, CCD+, and CTX; this package currently supports the (PPD) Prearranged Payment and Deposit Entry type transactions described by [NACHA](https://www.nacha.org//) Operating Rules

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/wadearnold/ach
```

## Example
Let's start parsing an ACH PPD file from the test data:

```go
package main
import (
  "fmt"
  "log"
  "wadearnold/ach"

)
func main() {
  f, err := os.Open("./testdata/ppd-debit.ach")
	if err != nil {
		log.Fatal("%s: ", err)
	}
	defer f.Close()

  achfile, err := ach.Decode(f)
  if err != nil {
		fmt.PrintF("Cannot decode: %v", err)
	}
  fmt.PrintF("ACH Destination from: %v", achfile.FileHeaderRecord.ImmediateDestinationName)

```

# Contributing

We use GitHub to manage reviews of pull requests.

* If you have a trivial fix or improvement, go ahead and create a pull
  request, addressing (with `@...`) one or more of the maintainers
  (see [AUTHORS.md](AUTHORS.md)) in the description of the pull request.

* If you plan to do something more involved, first propose your ideas
  in a Github issue. This will avoid unnecessary work and surely give
  you and us a good deal of inspiration.

* Relevant coding style guidelines are the [Go Code Review
  Comments](https://code.google.com/p/go-wiki/wiki/CodeReviewComments)
  and the _Formatting and style_ section of Peter Bourgon's [Go: Best
  Practices for Production
  Environments](http://peter.bourgon.org/go-in-production/#formatting-and-style).

## License
Apache License 2.0 See [LICENSE](LICENSE) for details.

## Prior Art
[Gusto: How ACH Works: A developer perspective](http://engineering.gusto.com/how-ach-works-a-developer-perspective-part-4/)

[mosscode / ach](https://github.com/mosscode/ach)
