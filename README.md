wadearnold/ach
===
[![GoDoc](https://godoc.org/github.com/wadearnold/ach?status.svg)](https://godoc.org/github.com/wadearnold/ach)
[![Build Status](https://travis-ci.org/wadearnold/ach.svg?branch=master)](https://travis-ci.org/wadearnold/ach)
[![Coverage Status](https://coveralls.io/repos/github/wadearnold/ach/badge.svg?branch=master)](https://coveralls.io/github/wadearnold/ach?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/wadearnold/ach)](https://goreportcard.com/report/github.com/wadearnold/ach)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/wadearnold/ach/master/LICENSE)


Package 'wadearnold/ach' implements a file decoder and encoder for parsing [ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House
) Automated Clearing House files. ACH is the primary method of electronic money movement throughout the United States.

## Project Status

ACH is at an early stage and under active development. Please star the project if you are interested in its development or join the development. 

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

## References  
* [Wikipeda: Automated Clearing House](http://en.wikipedia.org/wiki/Automated_Clearing_House)
* [Nacha ACH Network: How it Works](https://www.nacha.org/ach-network)
* [Federal ACH Directory](https://www.frbservices.org/EPaymentsDirectory/search.html)

## Format Specification
* [NACHA ACH File Formatting](https://www.nacha.org/system/files/resources/AAP201%20-%20ACH%20File%20Formatting.pdf)
* [PNC ACH File Specification](http://content.pncmc.com/live/pnc/corporate/treasury-management/ach-conversion/ACH-File-Specifications.pdf)
* [Thomson Reuters ACH FIle Structure](http://cs.thomsonreuters.com/ua/acct_pr/acs/cs_us_en/pr/dd/ach_file_structure_and_content.htm)
* [Gusto: How ACH Works: A developer perspective](http://engineering.gusto.com/how-ach-works-a-developer-perspective-part-4/)


## Prior Art
* [ACH:Builder - Tools for Building ACH](http://search.cpan.org/~tkeefer/ACH-Builder-0.03/lib/ACH/Builder.pm)
* [mosscode / ach](https://github.com/mosscode/ach)
* [Helper for building ACH files in Ruby](https://github.com/jm81/ach)

## License
Apache License 2.0 See [LICENSE](LICENSE) for details.
