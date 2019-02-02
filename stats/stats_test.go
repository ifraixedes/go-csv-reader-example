package stats_test

import (
	"encoding/csv"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ifraixedes/go-csv-reader-example/stats"
)

func TestNewRecordFromCSV(t *testing.T) {
	var tcases = []struct {
		desc   string
		argRec []string
		assert func(*testing.T, *stats.Record, error)
	}{
		{
			desc:   "successful",
			argRec: []string{"bid1", "userE", "not-used", "not-used", "2018-10-31T11:02:15-04:00", "no-used", "3"},
			assert: func(t *testing.T, r *stats.Record, err error) {
				assert.NoError(t, err)
				if assert.NotNil(t, r) {
					var tm, terr = time.Parse(time.RFC3339, "2018-10-31T11:02:15-04:00")
					require.NoError(t, terr)

					assert.Equal(t, &stats.Record{
						UserID:   "userE",
						ExecEnd:  tm,
						ExitCode: 3,
					}, r)
				}
			},
		},
		{
			desc:   "error: invalid exec time",
			argRec: []string{"bid1", "userE", "not-used", "not-used", "2018-10-31T11", "no-used", "3"},
			assert: func(t *testing.T, r *stats.Record, err error) {
				assert.Error(t, err)
				assert.Equal(t, stats.ErrInvalidRecord, err)
			},
		},
		{
			desc:   "error: invalid exit code",
			argRec: []string{"bid1", "userE", "not-used", "not-used", "2018-10-31T11:02:15-04:00", "no-used", "no-numeric"},
			assert: func(t *testing.T, r *stats.Record, err error) {
				assert.Error(t, err)
				assert.Equal(t, stats.ErrInvalidRecord, err)
			},
		},
		{
			desc:   "error: invalid number of fields",
			argRec: []string{"bid1", "userE", "not-used", "not-used"},
			assert: func(t *testing.T, r *stats.Record, err error) {
				assert.Error(t, err)
				assert.Equal(t, stats.ErrInvalidRecord, err)
			},
		},
	}

	for i := range tcases {
		var tc = tcases[i]
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			var r, err = stats.NewRecordFromCSV(tc.argRec)
			tc.assert(t, r, err)
		})
	}
}

func TestComputeBuilds(t *testing.T) {
	t.Run("successful", func(t *testing.T) {
		var recOutTWindow = make([]string, rand.Intn(10)+10)
		for i := 0; i < len(recOutTWindow); i++ {
			var bt = expectedBuilds.To.Add(time.Duration(rand.Int()+1) * time.Second)
			recOutTWindow[i] = genRecord(bt, "userA", uint8(rand.Intn(50)))
		}

		var records = append([]string{}, recordsUserA...)
		records = append(records, recordsUserB...)
		records = append(records, recordsUserC...)
		records = append(records, recordsUserD...)
		records = append(records, recordsUserE...)
		records = append(records, recordsUserF...)
		records = append(records, recOutTWindow...)

		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})

		var in = strings.NewReader(strings.Join(records, "\n"))
		var b, err = stats.ComputeBuilds(csv.NewReader(in), expectedBuilds.From, expectedBuilds.To)
		assert.NoError(t, err)
		assert.Equal(t, &expectedBuilds, b)
	})

	t.Run("error: invalid record", func(t *testing.T) {
		t.Skipf("TO BE IMPLEMENTED")
	})

	t.Run("error: reader returned error", func(t *testing.T) {
		t.Skipf("TO BE IMPLEMENTED")
	})
}
