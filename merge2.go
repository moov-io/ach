package ach

import (
	"fmt"

	rbtree "github.com/sakeven/RbTree"
)

func MergeFiles2(incoming []*File) ([]*File, error) {
	return MergeFilesWith2(incoming, Conditions{
		MaxLines: NACHAFileLineLimit,
	})
}

func MergeFilesWith2(incoming []*File, conditions Conditions) ([]*File, error) {
	// insert incoming files into red-black trees
	// read off sorted input to output files
	// TODO(adam):

	if len(incoming) == 0 {
		return nil, nil
	}

	sorted := &outFile{
		header: incoming[0].Header,
	}

	for i := range incoming {
		outFile := pickOutFile(incoming[i].Header, sorted)
		for j := range incoming[i].Batches {
			bh := incoming[i].Batches[j].GetHeader()
			if bh == nil {
				return nil, fmt.Errorf("incoming[%d].batch[%d] has nil batchHeader", i, j)
			}

			var found bool
			for k := range outFile.batches {
				if *bh == outFile.batches[k].header {
					found = true
					entries := incoming[i].Batches[j].GetEntries()
					for m := range entries {
						if outFile.batches[k].entries.Find(entries[m].TraceNumber) != nil {
							continue // skip for now
						}
						// TODO(adam): doesn't handle duplicates
						outFile.batches[k].entries.Insert(entries[m].TraceNumber, entries[m])
					}
				}
			}
			if !found {
				b := batch{
					header:  *bh,
					entries: rbtree.NewTree[string, *EntryDetail](),
				}
				entries := incoming[i].Batches[j].GetEntries()
				for m := range entries {
					if b.entries.Find(entries[m].TraceNumber) != nil {
						continue // skip for now
					}
					// TODO(adam): doesn't handle duplicates
					b.entries.Insert(entries[m].TraceNumber, entries[m])
				}
				outFile.batches = append(outFile.batches, b)
			}
		}
	}

	// fmt.Printf("sorted: %#v\n", sorted)
	var out []*File
	for {
		// Run through the linked list (sorted.next) until we terminate
		if sorted == nil {
			break
		}

		file := NewFile()
		file.Header = sorted.header

		for i := range sorted.batches {
			bh := sorted.batches[i].header
			bh.BatchNumber += i

			batch, err := NewBatch(&bh)
			if err != nil {
				return nil, err
			}

			// add each entry detail
			for {
				if sorted.batches[i].entries == nil {
					break
				}

				// TODO(adam): shouldn't need to do this...
				if sorted.batches[i].entries.Size() == 0 {
					break
				}

				next := sorted.batches[i].entries.Iterator() // TODO(adam): would this panic?
				if next == nil {
					break
				}
				batch.AddEntry(next.Value)

				sorted.batches[i].entries.Delete(next.Key)
			}

			err = batch.Create()
			if err != nil {
				return nil, err
			}

			file.AddBatch(batch)
		}

		err := file.Create()
		if err != nil {
			return nil, err
		}

		out = append(out, file)

		sorted = sorted.next
	}
	return out, nil
}

type outFile struct {
	header  FileHeader
	batches []batch

	next *outFile
}

type batch struct {
	header  BatchHeader
	entries *rbtree.Tree[string, *EntryDetail]
}

func pickOutFile(fh FileHeader, file *outFile) *outFile {
	// if file == nil {
	// 	return &outFile{
	// 		header: fh,
	// 	}
	// }
	if fh == file.header {
		return file
	}
	if file.next == nil {
		file.next = &outFile{
			header: fh,
		}
		return file.next
	}
	if file.next != nil {
		return pickOutFile(fh, file.next)
	}
	return nil // not possible
}
