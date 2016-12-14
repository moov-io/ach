# GoLange ACH file reader and writer
[![GoDoc](https://godoc.org/github.com/wadearnold/ach?status.svg)](https://godoc.org/github.com/wadearnold/ach)

## Overview
Package ach reads and writes (ACH) Automated Clearing House files. ACH is the
primary method of electronic money movemenet through the United States.
https://en.wikipedia.org/wiki/Automated_Clearing_House
Their are several kinds of ACH files PPD, PPD+, CCD, CCD+, and CTX; this package currently supports the (PPD) Prearranged Payment and Deposit Entry type transactions described by NACHA Operating Rules
https://www.nacha.org//

## Install

To install the library and command line program, use the following:

```
$ go get -u github.com/wadearnold/ach/...
```

ach is released under the Apache License. See [LICENSE](LICENSE)
