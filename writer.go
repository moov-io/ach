// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ach

import (
	"bufio"
	"io"
	"strings"
)

// Writer writes a File to an io.Writer.
// The File is validated against Nacha guidelines unless BypassValidation is enabled.
type Writer struct {
	w          *bufio.Writer
	lineNum    int    //current line being written
	LineEnding string // configurable line ending to support different consumer requirements
	// BypassValidation can be set to skip file validation and will allow non-compliant Nacha files to be written.
	BypassValidation bool
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w:          bufio.NewWriter(w),
		LineEnding: "\n", //set default line ending
	}
}

// Writer writes a single ach.file record to w
func (w *Writer) Write(file *File) error {
	if !w.BypassValidation {
		if err := file.Validate(); err != nil {
			return err
		}
	}

	w.lineNum = 0
	// Iterate over all records in the file
	if err := w.writeString(&file.Header); err != nil {
		return err
	}

	isADV := file.IsADV()

	if err := w.writeBatch(file, isADV); err != nil {
		return err
	}

	if err := w.writeIATBatch(file); err != nil {
		return err
	}

	if !isADV {
		if err := w.writeString(&file.Control); err != nil {
			return err
		}
	} else {
		if err := w.writeString(&file.ADVControl); err != nil {
			return err
		}
	}

	// pad the final block
	for i := 0; i < (10-(w.lineNum%10)) && w.lineNum%10 != 0; i++ {
		if _, err := w.w.WriteString(strings.Repeat("9", 94) + w.LineEnding); err != nil {
			return err
		}
	}

	return w.w.Flush()
}

// Flush writes any buffered data to the underlying io.Writer.
func (w *Writer) Flush() error {
	return w.w.Flush()
}

func (w *Writer) writeBatch(file *File, isADV bool) error {
	for _, batch := range file.Batches {
		if err := w.writeString(batch.GetHeader()); err != nil {
			return err
		}
		if !isADV {
			for _, entry := range batch.GetEntries() {
				if err := w.writeString(entry); err != nil {
					return err
				}
				if err := w.writeString(entry.Addenda02); err != nil {
					return err
				}
				for _, addenda05 := range entry.Addenda05 {
					if err := w.writeString(addenda05); err != nil {
						return err
					}
				}
				if err := w.writeString(entry.Addenda98); err != nil {
					return err
				}
				if err := w.writeString(entry.Addenda98Refused); err != nil {
					return err
				}
				if err := w.writeString(entry.Addenda99); err != nil {
					return err
				}
				if err := w.writeString(entry.Addenda99Dishonored); err != nil {
					return err
				}
				if err := w.writeString(entry.Addenda99Contested); err != nil {
					return err
				}
			}
		} else {
			for _, entry := range batch.GetADVEntries() {
				if err := w.writeString(entry); err != nil {
					return err
				}
				if err := w.writeString(entry.Addenda99); err != nil {
					return err
				}
			}
		}

		if batch.GetHeader().StandardEntryClassCode != ADV {
			if err := w.writeString(batch.GetControl()); err != nil {
				return err
			}
		} else {
			if err := w.writeString(batch.GetADVControl()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (w *Writer) writeIATBatch(file *File) error {
	for _, iatBatch := range file.IATBatches {
		if err := w.writeString(iatBatch.GetHeader()); err != nil {
			return err
		}
		for _, entry := range iatBatch.GetEntries() {
			if err := w.writeString(entry); err != nil {
				return err
			}
			if err := w.writeString(entry.Addenda10); err != nil {
				return err
			}
			if err := w.writeString(entry.Addenda11); err != nil {
				return err
			}
			if err := w.writeString(entry.Addenda12); err != nil {
				return err
			}
			if err := w.writeString(entry.Addenda13); err != nil {
				return err
			}
			if err := w.writeString(entry.Addenda14); err != nil {
				return err
			}
			if err := w.writeString(entry.Addenda15); err != nil {
				return err
			}
			if err := w.writeString(entry.Addenda16); err != nil {
				return err
			}
			// IAT Addenda17
			for _, addenda17 := range entry.Addenda17 {
				if err := w.writeString(addenda17); err != nil {
					return err
				}
			}
			// IAT Addenda18
			for _, addenda18 := range entry.Addenda18 {
				if err := w.writeString(addenda18); err != nil {
					return err
				}
			}
			if err := w.writeString(entry.Addenda98); err != nil {
				return err
			}

			if err := w.writeString(entry.Addenda99); err != nil {
				return err
			}

		}
		if err := w.writeString(iatBatch.GetControl()); err != nil {
			return err
		}
	}
	return nil
}

type writeEntry interface {
	String() string
}

func (w *Writer) writeString(entry writeEntry) error {
	if entry == nil {
		return nil
	}

	line := entry.String()
	if line == "" {
		return nil
	}

	if _, err := w.w.WriteString(line + w.LineEnding); err != nil {
		return err
	}

	w.lineNum++
	return nil
}
