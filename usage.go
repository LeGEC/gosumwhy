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
	Reads from stdin the output of 'go mod graph', and either lists the complete
	list of modules (and their distinct versions) present in that graph - gosumwhy list
	command - or finds dependency paths that lead to specific module versions -
	gosumwhy path command.

	The first module listed in the input graph is taken as the root module.

	If the 'list' command is executed, it will print on stdout the detailed list of
	modules+versions present in the dependency graph.

	If the 'path' command is executed, for each versioned module listed as argument,
	a dependency path that leads from the root module to that specific module+version
	is printed on stdout.
	The '-allv' option turns any module listed as argument into the list of all versions
	for that module.

OPTIONS
`)
	flag.CommandLine.PrintDefaults()
	fmt.Fprint(flag.CommandLine.Output(), `

EXAMPLES
	# prints the list of all modules listed as the dependencies of current module
	gosumwhy list

	# searches for a sequence of dependencies that leads to the specific version
	# rsc.io/quote/v3@v3.1.0
	gosumwhy path rsc.io/quote/v3@v3.1.0

	# enumerates all versions of gopkg.in/yaml.v2 listed as dependencies, and for
	# each version, print a dependency path that leads to that specific version
	gosumwhy -allv path gopkg.in/yaml.v2
`)
}
