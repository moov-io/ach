---
layout: page
title: Balanced offset
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Balanced offsets

ACH files can have "offset records" which are [EntryDetail](https://godoc.org/github.com/moov-io/ach#EntryDetail) records that balance a file's debits and credits. This means if there are debits, there is a credit offset matching the sum of the debits. If there are credits, there is a debit offset matching the sum of the credits. They are mutually exclusive.

Offset files are used to offset transactions from a single account inside of the ODFI. The Offset is the last record which settles the funds into a single ODFI account.

Note: Balanced offset files (via `WithOffset`) are only in ACH v1.1.0 and later.

## Setting offset metadata

An ACH [Batch](https://godoc.org/github.com/moov-io/ach#Batch) supports calling [WithOffset](https://godoc.org/github.com/moov-io/ach#Batch.WithOffset) to set offset information such as routing number, account number, account type, and a description.

```go
bh := ach.NewBatchHeader()
// fill in the fields

ed := ach.NewEntryDetail()
// fill in the fields

batch := ach.NewBatch(bh)
batch.AddEntry(ed)
batch.WithOffset(&ach.Offset{
    RoutingNumber: "...",
    AccountNumber: "...",
    AccountType: OffsetSavings,
    Description: "...",
})

// On each batch.Create() call the offset record will be re-tabulated
```
