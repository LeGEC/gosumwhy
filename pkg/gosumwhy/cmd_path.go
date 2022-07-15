package gosumwhy

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// Path searches for a path from the root module of 'g' to each of the modules mentioned in 'targets'.
// The path for each target is written to 'options.Out'.
// If 'options.AllVersions' is set: for each module listed in 'targets', a path to each version listed in graph 'g'
// for that module is written to 'out'.
// Otherwise: for each module+version listed in targets, a dependency path to that specific module+version is written to 'out'.
// A module can be specified either with a specific version (e.g: rsc.io/quote@v3.1.0) or without (e.g: rsc.io/quote).
// In the latter case: the newest version for that module is targeted.
func Path(g *Graph, targets []string, options *Options) error {
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

	if len(targets) == 0 {
		return errors.New("no target module provided")
	}

	c := &cmd{
		out: options.Out,
		err: options.Err,
	}

	var err error
	if options.AllVersions {
		err = c.printPathsToAllVersions(g, targets)
	} else {
		err = c.printPaths(g, targets)
	}

	return err
}

type cmd struct {
	out io.Writer
	err io.Writer
}

func (c *cmd) findPath(g *Graph, vtarget Version) error {
	path, err := g.pathTo(vtarget)
	if err != nil {
		return err
	}

	fmt.Fprintf(c.out, "----- %s\n", vtarget)
	for _, p := range path {
		fmt.Fprintf(c.out, "%s\n", p)
	}
	return nil
}

func (c *cmd) printPathsToAllVersions(g *Graph, targets []string) error {
	visited := make(map[string]bool)
	for _, target := range targets {
		vtarget := spec2version(target)
		if visited[vtarget.Path] {
			continue
		}
		visited[vtarget.Path] = true

		m := Version{Path: vtarget.Path}
		versions := g.versions[m]
		for i := len(versions) - 1; i >= 0; i-- {
			vtarget := versions[i]
			err := c.findPath(g, vtarget)
			if err != nil {
				if len(targets) == 1 {
					return err
				}
				fmt.Fprintf(c.err, "*** %s\n", err)
			}
		}
	}

	return nil
}

func (c *cmd) printPaths(g *Graph, targets []string) error {
	for _, target := range targets {
		vtarget := spec2version(target)
		if vtarget.Version == "" {
			// get latest version for that module:
			versions := g.versions[vtarget]
			if len(versions) > 0 {
				vtarget = versions[len(versions)-1]
			}
		}

		err := c.findPath(g, vtarget)
		if err != nil {
			if len(targets) == 1 {
				return err
			}
			fmt.Fprintf(c.err, "*** %s\n", err)
		}
	}

	return nil
}
