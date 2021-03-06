package main

import (
	"flag"

	u "github.com/araddon/gou"
	"github.com/dataux/dataux/testdata"
)

/*

usage:

	go build && ./importgithub

*/

var (
	eshost *string = flag.String("host", "localhost", "Elasticsearch Server Host Address")
)

func init() {
	u.SetupLogging("debug")
	u.SetColorIfTerminal()
}

func main() {
	flag.Parse()
	testdata.LoadGithubToEsOnce(*eshost)
}
