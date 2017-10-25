package ach

// written

import (
	"flag"
	"strings"
	"time"
)

const (
	timeFormat = "060102" // for date of death
)

func init() {
	flag.Lookup("alsologtostderr").Value.Set("true")
}

// ReturnAddenda
type ReturnAddenda struct {
	// RecordType defines the type of record in the block. entryAddendaPos 7
	recordType string
	// TypeCode Addenda types code '99'
	TypeCode string

	ReturnCode         string
	OriginalTrace      int
	DateOfDeath        *time.Time
	OriginalDFI        string
	AddendaInformation string
	Trace              int

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
}

// Parse takes the input record string and parses the ReturnAddenda values
func (returnAddenda *ReturnAddenda) Parse(record string) {
	// 1-1 Always "7"
	returnAddenda.recordType = "7"
	// 2-3 Defines the specific explanation and format for the addenda information contained in the same record
	returnAddenda.TypeCode = record[1:3]
	// 4-6
	returnAddenda.ReturnCode = record[3:6]
	// 7-21
	returnAddenda.OriginalTrace = returnAddenda.parseNumField(record[6:21])
	// 22-27, might be a date or blank
	result, err := time.Parse(timeFormat, strings.TrimSpace(record[21:27]))
	returnAddenda.DateOfDeath = &result
	if err != nil { // honestly it's best to just assume it's a blank date
		returnAddenda.DateOfDeath = nil
	}
	// 28-35
	returnAddenda.OriginalDFI = record[27:35]
	// 36-79
	returnAddenda.AddendaInformation = strings.TrimSpace(record[35:79])
	// 80-94
	returnAddenda.Trace = returnAddenda.parseNumField(record[79:94])
}

// implement later
func (returnAddenda *ReturnAddenda) Validate() error {
	return nil
}

// implement later
func (returnAddenda *ReturnAddenda) convertDateOfDeath() error {
	return nil
}
