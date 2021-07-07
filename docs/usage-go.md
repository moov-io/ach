---
layout: page
title: Go library
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Go library

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and Go v1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help in setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/ach/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/ach.git

# Pull down into the Go Module cache
$ go get -u github.com/moov-io/ach

$ go doc github.com/moov-io/ach BatchHeader
```

The package [`github.com/moov-io/ach`](https://pkg.go.dev/github.com/moov-io/ach) offers a Go-based ACH file reader and writer. To get started, check out a specific example:

### Supported Standard Entry Class (SEC) codes

| SEC Code | Description                                  | Example                                  | Read                | Write                                            |
|----------|---------------------------------------|------------------------------------------|-----------------------------------|------------------------------------|
| ACK      | Acknowledgment Entry for CCD          | [Credit](https://github.com/moov-io/ach/blob/master/test/ach-ack-read/ack-read.ach) | [ACK Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-AckRead) | [ACK Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-AckWrite) |
| ADV      | Automated Accounting Advice           | [Prenote Debit](https://github.com/moov-io/ach/blob/master/test/ach-adv-read/adv-read.ach) | [ADV Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-AdvRead) | [ADV Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-AdvWrite) |
| ARC      | Accounts Receivable Entry             | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-arc-read/arc-debit.ach) | [ARC Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-ArcReadDebit) | [ARC Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-ArcWriteDebit) |
| ATX      | Acknowledgment Entry for CTX          | [Credit](https://github.com/moov-io/ach/blob/master/test/ach-atx-read/atx-read.ach)  | [ATX Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-AtxRead) | [ATX Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-AtxWrite) |
| BOC      | Back Office Conversion                | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-boc-read/boc-debit.ach) | [BOC Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-BocReadDebit) | [BOC Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-BocWriteDebit) |
| CCD      | Corporate credit or debit             | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-ccd-read/ccd-debit.ach) | [CCD Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CcdReadDebit) | [CCD Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CcdWriteDebit) |
| CIE      | Customer-Initiated Entry              | [Credit](https://github.com/moov-io/ach/blob/master/test/ach-cie-read/cie-credit.ach) | [CIE Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CieRead) | [CIE Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CieWrite) |
| COR      | Automated Notification of Change(NOC) | [NOC](https://github.com/moov-io/ach/blob/master/test/ach-cor-read/cor-read.ach)   | [COR Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CorReadCredit) | [COR Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CorWriteCredit) |
| CTX      | Corporate Trade Exchange              | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-ctx-read/ctx-debit.ach) | [CTX Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CtxReadDebit) | [CTX Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-CtxWriteDebit) |
| DNE      | Death Notification Entry              | [DNE](https://github.com/moov-io/ach/blob/master/test/ach-dne-read/dne-read.ach)   | [DNE Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-DneRead) | [DNE Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-DneWrite) |
| ENR      | Automatic Enrollment Entry            | [ENR](https://github.com/moov-io/ach/blob/master/test/ach-enr-read/enr-read.ach)   | [ENR Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-EnrRead) | [ENR Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-EnrWrite) |
| IAT      | International ACH Transactions        | [Credit](https://github.com/moov-io/ach/blob/master/test/ach-iat-read/iat-credit.ach) | [IAT Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-IatReadMixedCreditDebit) | [IAT Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-IatWriteMixedCreditDebit) |
| MTE      | Machine Transfer Entry                | [Credit](https://github.com/moov-io/ach/blob/master/test/ach-mte-read/mte-read.ach)   | [MTE Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-MteReadDebit) | [MTE Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-MteWriteDebit) |
| POP      | Point of Purchase                     | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-pop-read/pop-debit.ach) | [POP Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-PopReadDebit) | [POP Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-PopWriteDebit) |
| POS      | Point of Sale                         | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-pos-read/pos-debit.ach) | [POS Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-PosReadDebit) | [POS Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-PosWriteDebit) |
| PPD      | Prearranged payment and deposits      | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-ppd-read/ppd-debit.ach) [Credit](https://github.com/moov-io/ach/blob/master/test/ach-ppd-read/ppd-credit.ach) | [PPD Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-PpdReadCredit) | [PPD Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-PpdWriteCredit) |
| RCK      | Represented Check Entries             | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-rck-read/rck-debit.ach) | [RCK Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-RckReadDebit) | [RCK Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-RckWriteDebit) |
| SHR      | Shared Network Entry                  | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-shr-read/shr-debit.ach) | [SHR Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-ShrReadDebit) | [SHR Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-ShrWrite) |
| TEL      | Telephone-Initiated Entry             | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-tel-read/tel-debit.ach) | [TEL Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-TelReadDebit) | [TEL Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-TelWriteDebit) |
| TRC      | Truncated Check Entry                 | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-trc-read/trc-debit.ach) | [TRC Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-TrcReadDebit) | [TRC Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-TrcWriteDebit) |
| TRX      | Check Truncation Entries Exchange     | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-trx-read/trx-debit.ach) | [TRX Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-TrxReadDebit) | [TRX Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-TrxWriteDebit) |
| WEB      | Internet-initiated Entries            | [Credit](https://github.com/moov-io/ach/blob/master/test/ach-web-read/web-credit.ach) | [WEB Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-WebReadCredit) | [WEB Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-WebWriteCredit) |
| XCK      | Destroyed Check Entry                 | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-xck-read/xck-debit.ach)  | [XCK Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-XckReadDebit) | [XCK Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-XckWriteDebit) |

### Segment files

| SEC Code | Name                                  | Example                                  | Read                | Write                                            |
|----------|---------------------------------------|------------------------------------------|-----------------------------------|------------------------------------|
| IAT      | International ACH Transactions        | [Credit](https://github.com/moov-io/ach/blob/master/test/ach-iat-read/iat-credit.ach) | [IAT Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-IatReadMixedCreditDebit) | [IAT Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-IatWriteMixedCreditDebit) |
| PPD      | Prearranged payment and deposits      | [Debit](https://github.com/moov-io/ach/blob/master/test/ach-ppd-read/ppd-debit.ach) [Credit](https://github.com/moov-io/ach/blob/master/test/ach-ppd-read/ppd-credit.ach) | [PPD Read](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-PpdReadSegmentFile) | [PPD Write](https://pkg.go.dev/github.com/moov-io/ach/examples#example-package-PpdWriteSegmentFile) |