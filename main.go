package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/LeGEC/gosumwhy/pkg/gosumwhy"
)

type CliOptions struct {
	gosumwhy.Options

	RunGoMod   bool
	ModulePath string
	GraphFile  string
}

var options CliOptions

func setupFlags() {
	options.Out = os.Stdout
	options.Err = os.Stderr

	flag.Usage = usageFunc

	// flags are defined here
	// if any code changes here, update the documentation in usage.go
	flag.StringVar(&options.GraphFile, "f", "", "read graph from that `file`\n"+
		"providing '-' as a filename (as in 'gosumshy -f -') forces reading from stdin\n")
	flag.BoolVar(&options.RunGoMod, "gomod", false, "if provided, run 'go mod graph' and read its output. See also -modpath\n")
	flag.StringVar(&options.ModulePath, "modpath", "", "run 'go mod graph' on the go module located at that `path`. Implies -gomod\n")
	flag.BoolVar(&options.AllVersions, "allv", false, "if set, the path command will print a dependency path for *each* version listed for that module\n")
}

func main() {
	setupFlags()
	flag.Parse()

	if options.ModulePath != "" {
		options.RunGoMod = true
	}
	inGraph, err := openGraphFile(options)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	args := flag.Args()
	err = gosumwhy.RunCmd(inGraph, args, &options.Options)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(2)
	}
}
