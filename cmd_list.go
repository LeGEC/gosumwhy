package gosumwhy

import (
	"fmt"
	"io"
	"sort"
)

// List writes on 'out' the list of all modules that appear in the graph 'g', and for each module the list of
// all versions mentioned in that graph.
func List(g *Graph, out io.Writer) {
	root := g.root

	var mods []Version
	for k := range g.versions {
		if k == root {
			continue
		}
		mods = append(mods, k)
	}

	sort.Slice(mods, func(i, j int) bool { return mods[i].LessThan(mods[j]) })

	printVersions(out, root, g.versions[root])
	for _, m := range mods {
		printVersions(out, m, g.versions[m])
	}
}

func printVersions(out io.Writer, mod Version, versions []Version) {
	fmt.Fprintf(out, "%s\n", mod)
	for i := len(versions) - 1; i >= 0; i-- {
		v := versions[i]
		fmt.Fprintf(out, "\t%s\n", v)
	}
	/*
		if len(versions) == 1 {
			fmt.Fprintf(out, "\t[1 version]\n")
		} else {
			fmt.Fprintf(out, "\t[%d versions]\n", len(versions))
		}
	*/
}
