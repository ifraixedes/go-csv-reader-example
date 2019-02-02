package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ifraixedes/go-csv-reader-example/stats"
)

func main() {
	var in, err = parseInput()
	if err != nil {
		exit(err)
	}

	var r = csv.NewReader(in.csv)
	b, err := stats.ComputeBuilds(r, in.twFrom, in.twTo)
	if err != nil {
		exit(err)
	}

	printBuilds(*b)
}

type input struct {
	csv    *os.File
	twFrom time.Time
	twTo   time.Time
}

func parseInput() (*input, error) {
	var (
		csvfp = flag.String("c", "", "CSV file path")
		tws   = flag.String("s", (time.Time{}).Format(time.RFC822), "Start time & date of the time window (default any). Format must be RFC822.")
		twe   = flag.String("e", time.Now().Format(time.RFC822), "End time & date of the time window (default current time). Format must be RFC822.")
	)

	flag.Parse()

	if *csvfp == "" {
		exit(errors.New("CSV file path must be indicated"))
	}

	var from, err = time.Parse(time.RFC822, *tws)
	if err != nil {
		exit(errors.New("Invalid start time & date format"))
	}

	to, err := time.Parse(time.RFC822, *twe)
	if err != nil {
		exit(errors.New("Invalid end time & date format"))
	}

	f, err := os.Open(*csvfp)
	if err != nil {
		perr, ok := err.(*os.PathError)
		if ok {
			exit(fmt.Errorf("Error while opening the CSV (%s): %s", perr.Path, perr.Err.Error()))
		} else {
			exit(fmt.Errorf("Error while opening the CSV (%s): %s", *csvfp, err.Error()))
		}
	}

	return &input{
		csv:    f,
		twFrom: from,
		twTo:   to,
	}, nil
}

func printBuilds(b stats.Builds) {
	var (
		topUsers    []string
		topErrCodes []uint8
	)

	for _, u := range b.TopUsers {
		if u == "" {
			break
		}
		topUsers = append(topUsers, u)
	}

	var topUsersMsg = ""
	if len(topUsers) > 0 {
		topUsersMsg = fmt.Sprintf("%v", topUsers)
	}

	for _, c := range b.TopErrCodes {
		if c == 0 {
			break
		}
		topErrCodes = append(topErrCodes, c)
	}

	var topErrCodesMsg = ""
	if len(topErrCodes) > 0 {
		topErrCodesMsg = fmt.Sprintf("%v", topErrCodes)
	}

	var successRateMsg = ""
	if b.Num > 0 {
		successRateMsg = fmt.Sprintf("%.2f%%", b.RateSuccess*100)
	}

	fmt.Printf(`
Remote Builder service builds stats
====================================
Applied time Window:      %s - %s
Number of Builds:         %d
Success rate:             %s
Top 5 users:              %s
Top 5 error exit codes:   %s
	`,
		b.From.Format(time.RFC850), b.To.Format(time.RFC850),
		b.Num,
		successRateMsg,
		topUsersMsg,
		topErrCodesMsg,
	)
}

func exit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "There has been an error.\n%s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
