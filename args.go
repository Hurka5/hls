package main

import (
	"github.com/alexflint/go-arg"
	"os"
)

type ArgsStruct struct {

	// Default options
	Paths []string `arg:"positional"`

	// Sort and Filter options
	All     bool `arg:"-a,--all"`
	Reverse bool `arg:"-r,--reverse" `

	//Display options
	Long    bool `arg:"-l,--long"`
	Oneline bool `arg:"-1,--oneline`
	Tree    bool `arg:"-t"--tree`
}

func (ArgsStruct) Version() string {
	return "hls 0.8"
}

var args ArgsStruct

func handleArguments() {

	arg.MustParse(&args)

	if len(args.Paths) == 0 {
		pwd := os.Getenv("PWD")
		args.Paths = append(args.Paths, pwd)
	}

	if args.Long {
		args.Oneline = true
	}
}
