---
layout: page
title: Reversal Files
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Reversal Files

Often ACH files are sent in error and need to be undone. Nacha has a process for creating a "reversal file" which needs to follow a few steps.
All other fields should be left untouched so the receiver of a reversal can match the Entry Detail records.

- Replace the value in Company Entry Description with `REVERSAL`
- Update Effective Entry Date in the future depending on same-day or standard ACH.
- Update each Transaction Code to undo fund movement.

## Creation

The Moov ACH library offers a `Reversal(..)` function on files.

```go
func (f *File) Reversal(effectiveEntryDate time.Time) error
    Reversal will transform a File into a Nacha compliant reversal which can be
    transmitted to undo fund movement.
```

```go
file, err := ach.ReadFile(..) // read from audit storage
if err != nil {
    // handle error
}

err = file.Reversal()
if err != nil {
    // handle error
}

// Now `file` has all Batches and Entries converted into Reversals
```
