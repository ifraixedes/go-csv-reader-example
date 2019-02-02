package stats

import (
	"encoding/csv"
	"errors"
	"io"
	"sort"
	"strconv"
	"time"
)

// ErrInvalidRecord can be returned by the computation functions if a CSV record
// has at least one field which isn't of the expected format.
var ErrInvalidRecord = errors.New("Record has at least one field of an unexpected format")

// Record has the typed fields which a CSV record has.
// NOTE it only has the fields used by the current computation functions.
type Record struct {
	UserID   string
	ExecEnd  time.Time
	ExitCode uint8
}

// NewRecordFromCSV returns a new Record from a CSV record.
// It can returns ErrInvalidRecord if the execution date or exit code fails or
// it doesn't have enough number of fields.
func NewRecordFromCSV(rec []string) (*Record, error) {
	if len(rec) < 6 {
		return nil, ErrInvalidRecord
	}

	var tm, err = time.Parse(time.RFC3339, rec[4])
	if err != nil {
		return nil, ErrInvalidRecord
	}

	code, err := strconv.ParseInt(rec[6], 10, 8)
	if err != nil {
		return nil, ErrInvalidRecord
	}

	return &Record{
		UserID:   rec[1],
		ExecEnd:  tm,
		ExitCode: uint8(code),
	}, nil
}

// Builds contains the stats of the remote build service in a time window.
type Builds struct {
	From        time.Time
	To          time.Time
	Num         uint64
	TopUsers    [5]string
	RateSuccess float32
	TopErrCodes [5]uint8
}

// ComputeBuilds calculate the stats of the remote build server of r records
// pending to read considering the passed time window.
func ComputeBuilds(r *csv.Reader, from time.Time, to time.Time) (*Builds, error) {
	var twr, err = NewTimeWindowReader(r, from, to)
	if err != nil {
		return nil, err
	}

	var (
		csvr          []string
		nBuilds       uint64
		nBuildsFailed uint64
		usersM        = map[string]int{}
		usersNBuilds  = []struct {
			u string
			n uint64
		}{}
		errCodesM       = map[uint8]int{}
		errCodesNBuilds = []struct {
			c uint8
			n uint64
		}{}
	)

	for csvr, err = twr.Read(); err == nil; csvr, err = twr.Read() {
		var rec *Record
		rec, err = NewRecordFromCSV(csvr)
		if err != nil {
			break
		}

		nBuilds++

		if i, ok := usersM[rec.UserID]; ok {
			usersNBuilds[i].n++
		} else {
			usersM[rec.UserID] = len(usersNBuilds)
			usersNBuilds = append(usersNBuilds,
				struct {
					u string
					n uint64
				}{rec.UserID, 1},
			)
		}

		if rec.ExitCode > 0 {
			nBuildsFailed++

			if i, ok := errCodesM[rec.ExitCode]; ok {
				errCodesNBuilds[i].n++
			} else {
				errCodesM[rec.ExitCode] = len(errCodesNBuilds)
				errCodesNBuilds = append(errCodesNBuilds,
					struct {
						c uint8
						n uint64
					}{rec.ExitCode, 1},
				)
			}
		}
	}

	if err != nil && err != io.EOF {
		return nil, err
	}

	var b = Builds{
		From:        from,
		To:          to,
		Num:         nBuilds,
		RateSuccess: float32(nBuilds-nBuildsFailed) / float32(nBuilds),
	}

	sort.Slice(usersNBuilds, func(i, j int) bool {
		return usersNBuilds[i].n > usersNBuilds[j].n
	})

	sort.Slice(errCodesNBuilds, func(i, j int) bool {
		return errCodesNBuilds[i].n > errCodesNBuilds[j].n
	})

	var n = len(usersNBuilds)
	if n > 5 {
		n = 5
	}
	for i := 0; i < n; i++ {
		b.TopUsers[i] = usersNBuilds[i].u
	}

	n = len(errCodesNBuilds)
	if n > 5 {
		n = 5
	}
	for i := 0; i < n; i++ {
		b.TopErrCodes[i] = errCodesNBuilds[i].c
	}

	return &b, nil
}
