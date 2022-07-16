package main

import (
	"flag"
	"fmt"
)

func usageFunc() {
	fmt.Fprint(flag.CommandLine.Output(), `	gosumwhy

	Searches the output of 'go mod graph' for dependency paths between modules.

USAGE
	gosumwhy [-f <graphFile>] [-gomod] [-modpath <path>]   # generic options
	gosumwhy ... list
	gosumwhy ... [-allv] path <module> [<module> ...]

DESCRIPTION
	Reads the output of 'go mod graph', and executes one of the two commands
	'list' or 'path'.

	The 'list' will print on stdout the detailed list of modules+versions
	present in the dependency graph.

	The 'path' command requires the user to provide one or more modules to target.
	For each of these modules, a dependency path that leads from the root module to
	that specific module is printed on stdout.
	A module can be either specified with a specific version (e.g: 'that/module@v1.0.1')
	or without (e.g: 'that/module') -- in the latter case, the newest version of that
	module is targetted.
	With the '-allv' option on: 'gosumwhy -allv path module' will print	a dependency
	path to each individual version of 'module'.

	See the OPTIONS and INPUT GRAPH sections below for more details on how gosumwhy
	takes a dependency graph as input.

OPTIONS
`)
	flag.CommandLine.PrintDefaults()
	fmt.Fprint(flag.CommandLine.Output(), `
INPUT GRAPH
	By default, gosumwhy tries to detect if its stdin is an interactive shell or not.
	
	If it is an interactive shell (e.g: not a pipe from another command), it tries to run
	'go mod graph' from the current directory, and uses the output of that command as its
	input graph.
	Otherwise, it reads its input from stdin (e.g: it assumes that stdin is some form of
	pipe, and takes that as an input).

	The command line options described in the OPTIONS section allow to override this
	default behavior.

	If you want to force gosumwhy to read its input from stdin: use '-f -'
	If you want to force gosumwhy to read the graph produced by 'go mod graph': use '-gomod'
	or '-modpath .'

EXAMPLES
	#  print all modules listed as dependencies of current module
	gosumwhy list

	# searches for a dependency path that leads to rsc.io/quote/v3@v3.1.0
	gosumwhy path rsc.io/quote/v3@v3.1.0

	# enumerates all versions of gopkg.in/yaml.v2 (the ones listed as dependencies),
	# and for each version, print a path
	gosumwhy -allv path gopkg.in/yaml.v2
`)
}
