package stats

import (
	"encoding/csv"
	"errors"
	"time"
)

// ErrInvalidTime can be returned in csv.ParseError.Err
var ErrInvalidTime = errors.New("Invalid format time")

// Reader is the interface with only the methods of csv.Reader which are used
// by this package
type Reader interface {
	Read() (record []string, err error)
}

type timeWindowReader struct {
	r      *csv.Reader
	from   time.Time
	to     time.Time
	tField uint32
	rowIdx uint64
}

// NewTimeWindowReader returns a Reader whose Read method only returns the
// records which are between the from and to time (both included). It's assumed
// that the field with index 4, have a time (RFC 3339 formatted string), which
// is the one to calculate if it's in the specified time window.
// An error is returned if r is nil or to is previous to from.
func NewTimeWindowReader(r *csv.Reader, from time.Time, to time.Time) (Reader, error) {
	if r == nil {
		return nil, errors.New("Invalid argument. Reader cannot be nil")
	}

	if from.After(to) {
		return nil, errors.New("Invalid argument. 'from' must be previous or equal to 'to'")
	}

	return &timeWindowReader{
		r:      r,
		from:   from.Round(0), // strip monotonic clock
		to:     to.Round(0),   // strip monotonic clock
		tField: 4,
	}, nil
}

// Read reads the csv.Reader records one by one, returning on each call the one
// that is inside of the configured time window, until csv.Reader has no more
// data.
// It behaves as csv.Reader.Read but also it returns ErrInvalidTime error
// if the field which  must contain the time under filtering isn't of the
// expected format.
func (twr *timeWindowReader) Read() ([]string, error) {
	for {
		var rc, err = twr.r.Read()
		if err != nil {
			return nil, err
		}

		tm, err := time.Parse(time.RFC3339, rc[twr.tField])
		if err != nil {
			_, _ = twr.r.ReadAll()
			return nil, &csv.ParseError{
				Line:   int(twr.rowIdx) + 1,
				Column: int(twr.tField),
				Err:    ErrInvalidTime,
			}
		}

		twr.rowIdx++

		if tm.Before(twr.from) || tm.After(twr.to) {
			continue
		}

		return rc, nil
	}
}
