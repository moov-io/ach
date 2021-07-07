---
layout: page
title: Flatten batches
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Flatten batches

File creation honors the order and number of Batch's that are created by the client. This is not optimized for the smallest file size and cost, that could be sent to the Federal Reserve for processing.

Flattening batches is a post file creation process that takes an input of an ACH.File and returns an ACH.File which has been optimized by flattening (merging) Batch Headers that have the same information and appending all Entries into a single Batch.

Note: FlattenBatches (via `FlattenBatches()`) is only in ACH v1.2.0 and later.

## Creating flattened batches

An ACH [File](https://godoc.org/github.com/moov-io/ach#File) supports calling [FlattenBatches](https://godoc.org/github.com/moov-io/ach#File.FlattenBatches):

```go
// Open a file for reading. Any io.Reader can be used.
f, err := os.Open(filepath.Join("test", "testdata", "flattenBatchesOneBatchHeader.ach"))

if err != nil {
	t.Fatal(err)
}
r := NewReader(f)
achFile, err := r.Read()
if err != nil {
	t.Fatalf("Issue reading file: %+v \n", err)
}

of, err := achFile.FlattenBatches()

if err != nil {
	t.Fatalf("Could not flatten the file: %+v \n", err)
}

if err := of.Validate(); err != nil {
	t.Fatalf("Flattened file did not validate: %+v \n", err)
}

// Write the file to stdout. Any io.Writer can be used.
w := ach.NewWriter(os.Stdout)
if err := w.Write(of); err != nil {
	log.Fatalf("Unexpected error: %s\n", err)
}

w.Flush()
```
