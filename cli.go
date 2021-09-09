package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type CmdOpts struct {
	Type     string
	UnixTime int64
	Limit    int64
	Args     []string
}

func ParseOpts() (*CmdOpts, error) {
	opts := CmdOpts{}
	unixTime := time.Now().Unix()
	flag.Usage = flagHelpMessage
	flag.StringVar(&opts.Type, "type", "events", "type: [events]")
	flag.Int64Var(&opts.Limit, "limit", 3, "upper limit to display:[0-6]")
	flag.Int64Var(&opts.UnixTime, "time", unixTime, "default time is UnixTime")
	flag.Parse()

	return &opts, nil
}

func flagHelpMessage() {
	cmd := os.Args[0]
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s [OPTIONS]", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Examples:")
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s -limit 5 -time 1422019499", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Options:")
	flag.PrintDefaults()
}
