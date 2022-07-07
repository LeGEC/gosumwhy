package gosumwhy

import (
	"io"
	"os"
)

type Options struct {
	AllVersions bool

	Out io.Writer
	Err io.Writer
}

// RunCmd applies the actions described in 'args' to the graph which can be read from 'inGraph'.
// 'inGraph' is expected to present a dependency graph, in the same format as the output of 'go mod graph'.
func RunCmd(inGraph io.Reader, args []string, options *Options) error {
	//default values
	if options == nil {
		options = &Options{}
	}
	if options.Out == nil {
		options.Out = os.Stdout
	}
	if options.Err == nil {
		options.Err = os.Stderr
	}

	// run command :

	// a. read graph from input
	g, err := readGraphFrom(inGraph)
	if err != nil {
		return err
	}

	// b. execute command based on arguments
	if len(args) == 0 || args[0] == "list" {
		List(g, options.Out)
		return nil
	}

	if args[0] == "path" {
		args = args[1:]
	}
	err = Path(g, args, options)
	return err
}
