---
layout: page
title: Command line
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Command line

On each release there's an `achcli` utility released. This tool can display ACH files in a human-readable format which is easier to read than their plaintext format. It also allows masking `DFIAccountNumber` values with the `-mask` flag.

## Options

```
$ achcli -help
achcli is a tool for displaying Nacha formatted ACH files in a human readable format.

USAGE
   achcli [-mask] [-pretty] [-validate opts.json] path/to/file.ach

EXAMPLES
  achcli -diff first.ach second.ach    Show the difference between two ACH files
  achcli -mask file.ach                Print file details with personally identifiable information partially removed
  achcli -reformat=json first.ach      Convert an incoming ACH file into another format (options: ach, json)
  achcli -validate opts.json file.ach  Read an ACH File with the provided ValidateOpts
  achcli -version                      Print the version of achcli (Example: v1.38.0)
  achcli 20060102.ach                  Summarize an ACH file for human readability

FLAGS
  -diff
        Compare two files against each other
  -flatten
        Flatten batches in each file
  -mask
        Mask/hide full account numbers and individual names
  -mask.accounts
        Mask/hide full account numbers
  -mask.corrections
        Mask/Hide Corrected Data in Addenda98 records
  -mask.names
        Mask/hide full individual names
  -merge
        Merge files before describing
  -pretty
        Display all values in their human readable format
  -pretty.amounts
        Display human readable amounts instead of exact values
  -reformat string
        Reformat an incoming ACH file to another format
  -skip-validation
        Skip all validation checks
  -v    Print verbose details about each ACH file
  -validate string
        Path to config file in json format to enable validation opts
  -version
        Print moov-io/ach cli version
```

## Install and Usage

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
