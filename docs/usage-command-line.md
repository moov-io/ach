---
layout: page
title: Command line
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Command line

On each release there's an `achcli` utility released. This tool can display ACH files in a human-readable format which is easier to read than their plaintext format. It also allows masking `DFIAccountNumber` values with the `-mask` flag.

```
$ wget -O achcli https://github.com/moov-io/ach/releases/download/v1.6.1/achcli-darwin-amd64 && chmod +x achcli

$ achcli test/testdata/ppd-debit.ach
Describing ACH file 'test/testdata/ppd-debit.ach'

  Origin     OriginName    Destination  DestinationName       FileCreationDate  FileCreationTime
  121042882  My Bank Name  231380104    Federal Reserve Bank  190624            0000

  BatchNumber  SECCode  ServiceClassCode  CompanyName      DiscretionaryData  Identification  EntryDescription  DescriptiveDate
  1            PPD      225 (Debits Only)  Name on Account                     121042882       REG.SALARY        

    TransactionCode   RDFIIdentification  AccountNumber      Amount     Name                    TraceNumber      Category
    27 (Checking Debit)  23138010            12345678           100000000  Receiver Account Name   121042880000001  

  ServiceClassCode  EntryAddendaCount  EntryHash  TotalDebits  TotalCredits  MACCode  ODFIIdentification  BatchNumber
  225 (Debits Only)  1                  23138010   100000000    0                      12104288            1

  BatchCount  BlockCount  EntryAddendaCount  TotalDebitAmount  TotalCreditAmount
  1           1           1                  100000000         0
```