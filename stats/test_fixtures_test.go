package stats_test

import (
	"fmt"
	"time"

	"github.com/ifraixedes/go-csv-reader-example/stats"
)

// Time window: 2018-10-31T03:43:46-04:00 - 2018-11-01T21:25:40-04:00
// Total builds: 53
// Builds succeeded: 30
// Top users: [userA, userB, userC, userD, userE]
// Top error codes: [ 4, 3, 5, 2, 7]
var expectedBuilds = stats.Builds{
	From: func() time.Time {
		t, err := time.Parse(time.RFC3339, "2018-10-31T03:43:46-04:00")
		if err != nil {
			panic(err)
		}
		return t
	}(),
	To: func() time.Time {
		t, err := time.Parse(time.RFC3339, "2018-11-01T21:25:40-04:00")
		if err != nil {
			panic(err)
		}
		return t
	}(),
	Num:         53,
	RateSuccess: 30.0 / 53.0,
	TopUsers:    [...]string{"userA", "userB", "userC", "userD", "userE"},
	TopErrCodes: [...]uint8{4, 3, 5, 2, 7},
}

// Num: 15
var recordsUserA = []string{
	"bid1,userA,not-used,not-used,2018-10-31T05:45:46-04:00,no-used,0,no-used",
	"bid1,userA,not-used,not-used,2018-10-31T06:45:46-04:00,no-used,1,no-used",
	"bid1,userA,not-used,not-used,2018-10-31T06:50:46-04:00,no-used,0,no-used",
	"bid1,userA,not-used,not-used,2018-10-31T07:00:00-04:00,no-used,2,no-used",
	"bid1,userA,not-used,not-used,2018-10-31T07:02:15-04:00,no-used,0,no-used",
	"bid1,userA,not-used,not-used,2018-10-31T11:02:15-04:00,no-used,0,no-used",
	"bid1,userA,not-used,not-used,2018-10-31T11:30:34-04:00,no-used,3,no-used",
	"bid1,userA,not-used,not-used,2018-10-31T11:32:34-04:00,no-used,4,no-used",
	"bid1,userA,not-used,not-used,2018-10-31T11:35:50-04:00,no-used,5,no-used",
	"bid1,userA,not-used,not-used,2018-10-31T11:38:04-04:00,no-used,0,no-used",
	"bid1,userA,not-used,not-used,2018-11-01T01:25:40-04:00,no-used,0,no-used",
	"bid1,userA,not-used,not-used,2018-11-01T02:25:40-04:00,no-used,0,no-used",
	"bid1,userA,not-used,not-used,2018-11-01T03:25:40-04:00,no-used,0,no-used",
	"bid1,userA,not-used,not-used,2018-11-01T04:25:40-04:00,no-used,0,no-used",
	"bid1,userA,not-used,not-used,2018-11-01T05:25:40-04:00,no-used,0,no-used",
}

// Num: 13
var recordsUserB = []string{
	"bid1,userB,not-used,not-used,2018-10-31T07:00:00-04:00,no-used,2,no-used",
	"bid1,userB,not-used,not-used,2018-10-31T07:02:15-04:00,no-used,0,no-used",
	"bid1,userB,not-used,not-used,2018-10-31T11:02:15-04:00,no-used,0,no-used",
	"bid1,userB,not-used,not-used,2018-10-31T11:30:34-04:00,no-used,3,no-used",
	"bid1,userB,not-used,not-used,2018-10-31T11:32:34-04:00,no-used,4,no-used",
	"bid1,userB,not-used,not-used,2018-10-31T11:35:50-04:00,no-used,5,no-used",
	"bid1,userB,not-used,not-used,2018-10-31T11:38:04-04:00,no-used,5,no-used",
	"bid1,userB,not-used,not-used,2018-11-01T01:25:40-04:00,no-used,0,no-used",
	"bid1,userB,not-used,not-used,2018-11-01T02:25:40-04:00,no-used,6,no-used",
	"bid1,userB,not-used,not-used,2018-11-01T03:25:40-04:00,no-used,0,no-used",
	"bid1,userB,not-used,not-used,2018-11-01T04:25:40-04:00,no-used,8,no-used",
	"bid1,userB,not-used,not-used,2018-11-01T11:55:05-04:00,no-used,0,no-used",
	"bid1,userB,not-used,not-used,2018-11-01T13:47:49-04:00,no-used,2,no-used",
}

// Num: 10
var recordsUserC = []string{
	"bid1,userC,not-used,not-used,2018-10-31T07:02:15-04:00,no-used,0,no-used",
	"bid1,userC,not-used,not-used,2018-10-31T11:02:15-04:00,no-used,0,no-used",
	"bid1,userC,not-used,not-used,2018-10-31T11:30:34-04:00,no-used,3,no-used",
	"bid1,userC,not-used,not-used,2018-10-31T11:32:34-04:00,no-used,4,no-used",
	"bid1,userC,not-used,not-used,2018-10-31T11:35:50-04:00,no-used,5,no-used",
	"bid1,userC,not-used,not-used,2018-10-31T11:38:04-04:00,no-used,0,no-used",
	"bid1,userC,not-used,not-used,2018-11-01T01:25:40-04:00,no-used,7,no-used",
	"bid1,userC,not-used,not-used,2018-11-01T02:25:40-04:00,no-used,7,no-used",
	"bid1,userC,not-used,not-used,2018-11-01T03:25:40-04:00,no-used,0,no-used",
	"bid1,userC,not-used,not-used,2018-11-01T04:25:40-04:00,no-used,0,no-used",
}

// Num: 8
var recordsUserD = []string{
	"bid1,userD,not-used,not-used,2018-10-31T07:02:15-04:00,no-used,0,no-used",
	"bid1,userD,not-used,not-used,2018-10-31T11:02:15-04:00,no-used,0,no-used",
	"bid1,userD,not-used,not-used,2018-10-31T11:30:34-04:00,no-used,3,no-used",
	"bid1,userD,not-used,not-used,2018-10-31T11:32:34-04:00,no-used,4,no-used",
	"bid1,userD,not-used,not-used,2018-11-01T01:25:40-04:00,no-used,0,no-used",
	"bid1,userD,not-used,not-used,2018-11-01T02:25:40-04:00,no-used,0,no-used",
	"bid1,userD,not-used,not-used,2018-11-01T03:25:40-04:00,no-used,0,no-used",
	"bid1,userD,not-used,not-used,2018-11-01T04:25:40-04:00,no-used,0,no-used",
}

// Num: 6
var recordsUserE = []string{
	"bid1,userE,not-used,not-used,2018-10-31T11:02:15-04:00,no-used,0,no-used",
	"bid1,userE,not-used,not-used,2018-10-31T11:30:34-04:00,no-used,3,no-used",
	"bid1,userE,not-used,not-used,2018-10-31T11:32:34-04:00,no-used,4,no-used",
	"bid1,userE,not-used,not-used,2018-11-01T01:25:40-04:00,no-used,0,no-used",
	"bid1,userE,not-used,not-used,2018-11-01T02:25:40-04:00,no-used,0,no-used",
	"bid1,userE,not-used,not-used,2018-11-01T03:25:40-04:00,no-used,0,no-used",
}

// Num: 1
var recordsUserF = []string{
	"bid1,userF,not-used,not-used,2018-10-31T11:32:34-04:00,no-used,4,no-used",
}

func genRecord(bt time.Time, user string, exitCode uint8) string {
	return fmt.Sprintf("bid1,%s,not-used,not-used,%s,not-used,%d,no-used",
		user, bt.Format(time.RFC3339), exitCode,
	)
}
