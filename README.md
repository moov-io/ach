# ACH file reader and writer in GoLang
[![GoDoc](https://godoc.org/github.com/wadearnold/ach?status.svg)](https://godoc.org/github.com/wadearnold/ach)
[![Build Status](https://travis-ci.org/wadearnold/ach.svg?branch=master)](https://travis-ci.org/wadearnold/ach)

## Overview
Package ach reads and writes [ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House
) Automated Clearing House files. ACH is the
primary method of electronic money movement through the United States.

Their are several kinds of ACH files PPD, PPD+, CCD, CCD+, and CTX; this package currently supports the (PPD) Prearranged Payment and Deposit Entry type transactions described by [NACHA](https://www.nacha.org//) Operating Rules


## Install

To install the library and command line program, use the following:

```
$ go get -u github.com/wadearnold/ach/...
```

## License
ach is released under the Apache License. See [LICENSE](LICENSE)
