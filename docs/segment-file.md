---
layout: page
title: Segmenting files
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Segmenting files

ACH File creation should default to NACHA rules by supporting both credits and debits in the same ACH file.  Some legacy systems require the file to be segmented and contain only debits or only credits in a file.

Segmenting files is a post file creation process that takes an input of an ACH.File and returns two ACH.Files that have been segmented (debits and credits).

Note: Segmented files (via `SegmentFile()`) are only in ACH v1.1.0 and later.

## Creating segmented files

An ACH [File](https://godoc.org/github.com/moov-io/ach#File) supports calling [SegmentFile](https://godoc.org/github.com/moov-io/ach#File.SegmentFile) to create a debit ach file and credit ach file from an ach file that contains mixed debits and credits.

```go
// open an ACH file(from the ./examples/ directory)
f, err := os.Open(filepath.Join("testdata", "ppd-mixedDebitCredit.ach"))
if err != nil {
	log.Fatal(err)
}

// read file
r := ach.NewReader(f)
achFile, err := r.Read()
if err != nil {
	fmt.Printf("Issue reading file: %+v \n", err)
}

// Segment File
creditFile, debitFile, err := achFile.SegmentFile()
if err != nil {
	fmt.Printf("Could not segment the file: %v", err)
}

// write the file to std out. Anything io.Writer
w := ach.NewWriter(os.Stdout)
if err := w.Write(creditFile); err != nil {
	log.Fatalf("Unexpected error: %s\n", err)
}
if err := w.Write(debitFile); err != nil {
	log.Fatalf("Unexpected error: %s\n", err)
}

w.Flush()
```
