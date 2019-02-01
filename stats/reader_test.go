package stats_test

import (
	"encoding/csv"
	"strings"
	"testing"
	"time"

	"github.com/ifraixedes/go-csv-reader-example/stats"
	"github.com/stretchr/testify/assert"
)

func TestNewCSVTimeWindowReader(t *testing.T) {
	type params struct {
		r    *csv.Reader
		from time.Time
		to   time.Time
	}

	var tcases = []struct {
		desc   string
		args   params
		assert func(*testing.T, stats.Reader, error)
	}{
		{
			desc: "successful",
			args: params{
				r:    csv.NewReader(strings.NewReader("")),
				from: time.Time{},
				to:   time.Now(),
			},
			assert: func(t *testing.T, cr stats.Reader, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, cr)
			},
		},
		{
			desc: "error: nil CSV reader",
			args: params{
				r:    nil,
				from: time.Time{},
				to:   time.Now(),
			},
			assert: func(t *testing.T, cr stats.Reader, err error) {
				assert.Error(t, err)
			},
		},
		{
			desc: "error: invalid time window",
			args: params{
				r:    csv.NewReader(strings.NewReader("")),
				from: time.Now(),
				to:   time.Time{},
			},
			assert: func(t *testing.T, cr stats.Reader, err error) {
				assert.Error(t, err)
			},
		},
	}

	for i := range tcases {
		var tc = tcases[i]
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			var r, err = stats.NewTimeWindowReader(tc.args.r, tc.args.from, tc.args.to)
			tc.assert(t, r, err)
		})
	}
}
