---
layout: page
title: Merging files
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Merging files

Moov ACH supports merging an arbitrary set of ACH files together. This is useful for optimizing cost and network efficiency. NACHA files are limited to 10,000 lines (in their text format) and so Moov ACH merges valid files together where the FileHeaders match the same ABA routing numbers.

An example of merging ACH files can be seen below. Assuming we have two ACH files to merge (`first.ach` and `second.ach`) on disk, let's read them and produce a merged file.


```go
package main

import (
    "log"
    "os"
    "github.com/moov-io/ach"
)

func main() {
    // first.ach and second.ach need to have the same Origin and Destination ABA routing numbers
    // in order to be merged into one ACH file, otherwise two ACH files would be returned.

    first, err := readACHFilepath("first.ach")
    if err != nil {
        log.Fatalf("first.ach parsing error: %v", err)
    }
    second, err := readACHFilepath("second.ach")
    if err != nil {
        log.Fatalf("second.ach parsing error: %v", err)
    }

    mergedFiles, err := ach.MergeFiles([]*ach.File{first, second})
    if err != nil {
        log.Fatalf("problem merging ACH files: %v", err)
    }
    log.Printf("merged into %d ACH files", len(mergedFiles))
}

func readACHFilepath(path string) (*ach.File, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	file, err := ach.NewReader(fd).Read()
	if err != nil {
		return nil, err
	}
	return &file, nil
}
```

Running the following code would produce (saved as `merge.go`):

```
$ go run merge.go
2019/05/23 13:07:37 merged into 1 ACH files
```
