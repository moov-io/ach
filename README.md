wadearnold/ach
===
[![GoDoc](https://godoc.org/github.com/wadearnold/ach?status.svg)](https://godoc.org/github.com/wadearnold/ach)
[![Build Status](https://travis-ci.org/wadearnold/ach.svg?branch=master)](https://travis-ci.org/wadearnold/ach)

Package 'wadearnold/ach' implements a file decoder and encoder for parsing [ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House
) Automated Clearing House files. ACH is the primary method of electronic money movement throughout the United States.

Their are several kinds of ACH files PPD, PPD+, CCD, CCD+, and CTX; this package currently supports the (PPD) Prearranged Payment and Deposit Entry type transactions described by [NACHA](https://www.nacha.org//) Operating Rules

---
* [Install](#Install)
* [Example](#Example)
* [License](#License)
---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/wadearnold/ach
```

## Example
Let's start registering a couple of URL paths and handlers:

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


## License
Apache License 2.0 See [LICENSE](LICENSE) for details.
