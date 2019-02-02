package stats_test

import (
	"encoding/csv"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/ifraixedes/go-csv-reader-example/stats"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestTimeWindowReader_Read(t *testing.T) {
	t.Run("successful", func(t *testing.T) {
		var records = []string{
			"0,1,2,3,2018-10-31T02:47:31-04:00,5",
			"0,1,2,3,2018-10-30T02:47:31-04:00,5",
			"0,1,2,3,2018-11-01T05:50:28-04:00,5",
			"0,1,2,3,2018-11-02T01:08:28-04:00,5",
		}

		var from, err = time.Parse(time.RFC3339, "2018-10-20T05:50:28-04:00")
		require.NoError(t, err)
		to, err := time.Parse(time.RFC3339, "2018-11-01T00:50:28-04:00")

		var in = strings.Join(records, "\n")
		twr, err := stats.NewTimeWindowReader(csv.NewReader(strings.NewReader(in)), from, to)
		require.NoError(t, err)

		record, err := twr.Read()
		if assert.NoError(t, err) {
			assert.Equal(t, records[0], strings.Join(record, ","))
		}

		record, err = twr.Read()
		if assert.NoError(t, err) {
			assert.Equal(t, records[1], strings.Join(record, ","))
		}

		_, err = twr.Read()
		assert.Equal(t, io.EOF, err)
	})

	t.Run("error: csv.Reader.Read", func(t *testing.T) {
		var records = []string{
			"0,1,2,3,2018-10-31T02:47:31-04:00,5",
			"2,3,2018-10-30T02:47:31-04:00,5",
			"0,1,2,3,2018-11-01T05:50:28-04:00,5",
			"0,1,2,3,2018-11-02T01:08:28-04:00,5",
		}

		var from, err = time.Parse(time.RFC3339, "2018-10-20T05:50:28-04:00")
		require.NoError(t, err)
		to, err := time.Parse(time.RFC3339, "2018-11-01T00:50:28-04:00")

		var in = strings.Join(records, "\n")
		twr, err := stats.NewTimeWindowReader(csv.NewReader(strings.NewReader(in)), from, to)
		require.NoError(t, err)

		_, err = twr.Read()
		assert.NoError(t, err)

		_, err = twr.Read()
		assert.Errorf(t, err, csv.ErrFieldCount.Error())

		_, err = twr.Read()
		assert.Equal(t, io.EOF, err)
	})

	t.Run("error: invalid format time field", func(t *testing.T) {
		var records = []string{
			"0,1,2,3,2018-10-30T02:47:31-04:00,5",
			"0,1,2,3,2018-11-01T05:50:28-04:00,5",
			"0,1,2,3,2018-10-31,5",
			"0,1,2,3,2018-11-02T01:08:28-04:00,5",
		}

		var from, err = time.Parse(time.RFC3339, "2018-10-20T05:50:28-04:00")
		require.NoError(t, err)
		to, err := time.Parse(time.RFC3339, "2018-11-01T00:50:28-04:00")

		var in = strings.Join(records, "\n")
		twr, err := stats.NewTimeWindowReader(csv.NewReader(strings.NewReader(in)), from, to)
		require.NoError(t, err)

		_, err = twr.Read()
		assert.NoError(t, err)

		_, err = twr.Read()
		if assert.Error(t, err) && assert.IsType(t, &csv.ParseError{}, err) {
			var errp = err.(*csv.ParseError)
			assert.Equal(t, &csv.ParseError{
				Line:   3,
				Column: 4,
				Err:    stats.ErrInvalidTime,
			}, errp)
		}

		_, err = twr.Read()
		assert.Equal(t, io.EOF, err)
	})
}
