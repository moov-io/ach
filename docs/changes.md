---
layout: page
title: Change files
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Change files

A Notification of Change (NOC) is a non-dollar entry transmitted from a receiving depository financial institution (RDFI) to the originating depository financial institution (ODFI). These are sent in response to outdated or erroneous information in an initial entry. An NOC often occurs due to bank mergers or acquisitions that change account and/or routing numbers. If an RDFI sends an NOC, the ODFI will need to inform their originator promptly.

The Standard Entry Class code for an NOC is "COR". There are a few possible reasons for an NOC, each defined by a "change code". The most common codes are `C01` for an incorrect account number and `C02` for an outdated routing number. We have [a list of supported change codes](#change-codes) below.

NOCs are identified by an [Addenda98](https://pkg.go.dev/github.com/moov-io/ach?tab=doc#Addenda98) record on the EntryDetail with a [ChangeCode](https://pkg.go.dev/github.com/moov-io/ach?tab=doc#ChangeCode) that can be processed.

### Processing

The RDFI must send an NOC within two banking days of the original entry settlement date (to which the NOC is in response to).

The ODFI is responsible for forwarding an NOC to the Originator within two banking days of the NOC settlement date. They must provide the Originator with the following information at a minimum:

- Company name
- Company identification
- Company Entry description
- Effective Entry date
- DFI account number
- Individual name/receiving company name
- Individual identification number

- Change code
- Original Entry trace number
- Original RDFI indentification
- Corrected data

The Originator must make specified changes within six banking days of receiving the above information or prior to initiating another entry to the receiver's account (whichever is later).

### Creation

When creating an NOC entry, add an [Addenda98](https://pkg.go.dev/github.com/moov-io/ach?tab=doc#Addenda98) record onto the EntryDetail with the appropriate change code.

```go
addenda98 := ach.NewAddenda98()
addenda98.ChangeCode = "C01"
addenda98.OriginalTrace = "121042880000001"
addenda98.OriginalDFI = "121042882"
addenda98.CorrectedData = "1918171614"
addenda98.TraceNumber = "91012980000088"

//entry.Addenda98 = addenda98
```

### Change codes

| Code | Reason | Description |
|----|-----|------|
| `C01` | Incorrect bank account number | Bank account number incorrect or formatted incorrectly |
| `C02` | Incorrect transit/routing number | Once valid transit/routing number must be changed |
| `C03` | Incorrect transit/routing number and bank account number | Once valid transit/routing number must be changed and causes a change to bank account number structure |
| `C04` | Bank account name change | Customer has changed name or ODFI submitted name incorrectly |
| `C05` | Incorrect payment code | Entry posted to demand account should contain savings payment codes or vice versa |
| `C06` | Incorrect bank account number and transit code | Bank account number must be changed and payment code should indicate posting to another account type (demand/savings) |
| `C07` | Incorrect transit/routing number, bank account number and payment code | Changes required in three fields indicated |
| `C09` | Incorrect individual ID number | Individual's ID number is incorrect |
| `C10` | Incorrect company name | Company name is no longer valid and should be changed |
| `C11` | Incorrect company identification | Company ID is no longer valid and should be changed |
| `C12` | Incorrect company name and company ID | Both the company name and company id are no longer valid and must be changed |

#### Refused Notification of Change

When ODFIs cannot forward entries to the Originator or a NOC is malformed, invalid or otherwise unable to be processed a Refused NOC may be issued. This will indicate a Refused Change Code to be handled and must be initiated within 15 days of receipt of the NOC.

| Code | Description |
|----|-----|
| `C61` | Misrouted Notification of Change |
| `C62` | Incorrect Trace Number |
| `C63` | Incorrect Company Identification Number |
| `C64` | Incorrect Individual Identification Number or Identification Number |
| `C65` | Incorrectly Formatted Corrected Data |
| `C66` | Incorrect Discretionary Data |
| `C67` | Routing Number not from Original Entry Detail Record |
| `C68` | DFI Account Number not from Original Entry Detail Record |
| `C69` | Incorrect Transaction Code |
