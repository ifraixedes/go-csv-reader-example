package stats

import (
	"encoding/csv"
	"errors"
	"time"
)

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

func (_ *timeWindodReader) Read() ([]string, error) {
	// TODO: WIP
	return nil, nil
}
