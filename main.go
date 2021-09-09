package main

import (
	"flag"
	"time"
)

type CmdOpts struct {
	Type     string
	UnixTime int64
	Args     []string
}

func main() {
	unixTime := time.Now().Unix()

	opts := CmdOpts{}
	//flag.Usage = flagHelpMessage
	flag.StringVar(&opts.Type, "type", "events", "[TODO: add explain]")
	flag.Int64Var(&opts.UnixTime, "time", unixTime, "[TODO: add explain]")
	flag.Parse()

	if opts.Type == "events" {
		fetchEvents(&opts)
	}
}
