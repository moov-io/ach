---
layout: page
title: Calculating Effective Entry Date 
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Calculating Effective Entry Date 

## 1. Transmission Windows for Same-Day ACH

There are [three same-day transmission deadlines with the Federal Reserve](https://www.frbservices.org/resources/resource-centers/same-day-ach/fedach-processing-schedule.html):

* **1st Cutoff**: 10:30 a.m. ET
* **2nd Cutoff**: 2:45 p.m. ET
* **3rd Cutoff**: 4:45 p.m. ET

Files must be submitted before the last cutoff time to be eligible for same-day processing. At each cutoff the Federal Reserve processes uploaded files and distributes them to financial institutions. 

If a file is submitted after the last cutoff, entries that list the current banking day as their Effective Entry Date will be settled on the next banking day.

## 2. Same-Day ACH Transaction Limit

* Entries qualify for same-day processing if their amount is **$1,000,000** or less.
* In a batch containing entries over the limit, those entries will have their Effective Entry Date adjusted to the next banking day and remaining items will still be eligible for same-day settlement

## 3. Credit Entries

For credit entries, the Effective Entry Date must be either:
* The **same banking day** as the processing date for **Same Day Entries**, or
* **One or two banking days** following the processing date for other entries

## 4. Debit Entries

For debit entries, the Effective Entry Date must be either:
* The **same banking day** as the processing date for **Same Day Entries**, or
* **One banking day** following the processing date for other entries

**Note:** Receiving banks may place holds on ACH debits per their funds availability policies, separate from these settlement times.

## 5. Effective Entry Date - Edge Cases

### ENR Entries
* For Automated Enrollment Entries (ENR), the Effective Entry Date field must be **space-filled**
* ENR entries cannot be Same Day Entries

### Notifications of Change
* The ACH Operator does **not edit** the Effective Entry Date field for Notifications of Change

### Return Entries
* The ACH Operator may verify that the Effective Entry Date is properly formatted and may replace the field's existing content with the current processing date if the Effective Entry Date is invalid

# Calculating Effective Entry Dates Using moov-io/base Functions

## Overview

The moov-io/base package provides functionality for calculating valid Effective Entry Dates when creating ACH batch headers, ensuring compliance with Nacha rules and Federal Reserve processing schedules.

## Creating Batch Headers

The [`NewBatchHeader`](https://pkg.go.dev/github.com/moov-io/ach#NewBatchHeader) function creates a [`BatchHeader`](https://pkg.go.dev/github.com/moov-io/ach#BatchHeader) object that requires several key fields to be populated according to Nacha rules:

```go
bh := ach.NewBatchHeader()
bh.ServiceClassCode = "220"                  // 220 for credits only
bh.StandardEntryClassCode = ach.PPD              // PPD, CCD, WEB, etc.
bh.CompanyName = "ACME Corporation"          // Your company name
bh.CompanyIdentification = "121042882"       // Your company ID
bh.CompanyEntryDescription = "PAYROLL"       // Transaction description
bh.ODFIIdentification = "12104288"           // Your bank's routing number

bh.EffectiveEntryDate = "240102"             // YYMMDD format (e.g. 2024-01-02)
```

## Banking Day Functions

The [`moov-io/base.Time`](https://pkg.go.dev/github.com/moov-io/base#NewTime) package provides these essential functions for determining valid banking days:

* [`AddBankingTime(hours, minutes, seconds int)`](https://pkg.go.dev/github.com/moov-io/base#Time.AddBankingTime): Increments time by specified duration while respecting banking hours (9am-5pm ET)
* [`AddBankingDay(d int)`](https://pkg.go.dev/github.com/moov-io/base#Time.AddBankingDay): Adds the specified number of valid banking days to a date
* [`IsBankingDay()`](https://pkg.go.dev/github.com/moov-io/base#Time.IsBankingDay): Validates if a date is a banking day by checking both holidays and weekends
* [`IsBusinessDay()`](https://pkg.go.dev/github.com/moov-io/base#Time.IsBusinessDay): Determines if a date is a standard business day
* [`IsHoliday()`](https://pkg.go.dev/github.com/moov-io/base#Time.IsHoliday): Checks if a date falls on a Federal Reserve holiday
* [`IsWeekend()`](https://pkg.go.dev/github.com/moov-io/base#Time.IsWeekend): Verifies if a date falls on Saturday or Sunday

## Transmission Windows

The package handles Same-Day ACH transmission windows with time-based validations. Here's how to implement the cutoff time checks:

```go
package main

import (
   "fmt"
   "github.com/moov-io/ach"
   "github.com/moov-io/base"
)

func main() {
   // Create batch header
   bh := ach.NewBatchHeader()
   bh.ServiceClassCode = "220"
   bh.StandardEntryClassCode = ach.PPD
   bh.CompanyName = "ACME Corporation"
   bh.CompanyIdentification = "121042882"
   bh.CompanyEntryDescription = "PAYROLL"
   bh.ODFIIdentification = "12104288"

   amount := 500000 // Example amount
   bh.EffectiveEntryDate = calculateEffectiveEntryDate(amount)
   
   fmt.Printf("Settlement Date: %s\n", bh.EffectiveEntryDate)
}

func calculateEffectiveEntryDate(amount int) string {
   now := base.Now()
   
   // Create cutoff time - in real world scenarios, there's typically an internal buffer before Fed cutoff
   thirdCutoff := now.AddBankingTime(16, 15, 0)  // 4:15 PM ET

   if now.IsBankingDay() {
       // Same day settlement is available if under limit and before cutoff
       if amount <= 1000000 && now.Before(thirdCutoff.Time) {
           return now.Format("060102")
       }
   }

   // After cutoff or non-banking day - settle on future date
   if amount <= 1000000 {
       // Same day entries settle next banking day
       return now.AddBankingDay(1).Format("060102")
   }
   // Non same-day entries must settle two banking days out
   return now.AddBankingDay(2).Format("060102")
}
