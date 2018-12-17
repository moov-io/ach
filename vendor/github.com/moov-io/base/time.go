// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package base

import (
	"errors"
	"time"
)

const (
	iso8601Format = "2006-01-02T15:04:05Z07:00"
)

// Time is an time.Time struct that encodes and decodes in ISO 8601.
//
// ISO 8601 is usable by a large array of libraries whereas RFC 3339 support
// isn't often part of language standard libraries.
type Time struct {
	time.Time
}

// Now returns a Time object with the current clock time set.
func Now() Time {
	return Time{
		Time: time.Now().UTC().Truncate(1 * time.Second),
	}
}

// NewTime wraps a time.Time value in Moov's base.Time struct.
// If you need the underlying time.Time value call .Time:
//
// now := Now()
// fmt.Println(start.Sub(now.Time))
func NewTime(t time.Time) Time {
	return Time{t}
}

func (t Time) MarshalJSON() ([]byte, error) {
	var bs []byte
	bs = append(bs, '"')

	t.Time = t.Time.Truncate(1 * time.Second) // drop milliseconds
	bs = t.AppendFormat(bs, iso8601Format)

	bs = append(bs, '"')
	return bs, nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	tt, err := time.Parse(`"`+iso8601Format+`"`, string(data))
	if err != nil || tt.IsZero() {
		// Try in RFC3339 format (default Go time)
		tt, _ = time.Parse(time.RFC3339, string(data))
		*t = NewTime(tt)
	}
	t.Time = tt.Truncate(1 * time.Second) // drop millis

	// Return an error if nothing was parsed.
	if t.Time.IsZero() {
		return errors.New("empty date time")
	}

	return nil
}

// Equal compares two Time values. Time values are considered equal if they both truncate
// to the same year/month/day and hour/minute/second.
func (t Time) Equal(other Time) bool {
	t1 := t.Time.Truncate(1 * time.Second)
	t2 := other.Time.Truncate(1 * time.Second)
	return t1.Equal(t2)
}
