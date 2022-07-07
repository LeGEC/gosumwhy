package gosumwhy

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Version struct {
	Path    string
	Version string `json:",omitempty"`
}

func (v Version) String() string {
	if v.Version != "" {
		return v.Path + "@" + v.Version
	}
	return v.Path
}

func (v Version) LessThan(o Version) bool {
	if v.Path != o.Path {
		return v.Path < o.Path
	}
	return lessVersionsString(v.Version, o.Version)
}

type Graph struct {
	mod      map[string]Version
	versions map[Version][]Version

	root   Version
	uses   map[Version][]Version
	usedBy map[Version][]Version
}

func newGraph() *Graph {
	g := Graph{
		mod:      make(map[string]Version, 0),
		versions: make(map[Version][]Version),
		uses:     make(map[Version][]Version),
		usedBy:   make(map[Version][]Version),
	}

	return &g
}

func readGraphFrom(r io.Reader) (*Graph, error) {
	g := newGraph()

	s := bufio.NewScanner(r)
	l := 0
	for s.Scan() {
		l++
		line := s.Text()
		line = strings.TrimSpace(line)
		if line == "" || line[0] == '#' {
			continue
		}

		idx := strings.IndexByte(line, ' ')
		if idx <= 0 {
			return nil, fmt.Errorf("line %d: invalid format", l)
		}
		moda := line[:idx]
		modb := line[idx+1:]

		g.addEdge(moda, modb)
	}
	g.tidy()
	if s.Err() != nil {
		return nil, s.Err()
	}

	var keys []Version
	for k := range g.versions {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i].LessThan(keys[j]) })

	return g, nil
}

func (g *Graph) addMod(spec string) Version {
	if v, ok := g.mod[spec]; ok {
		return v
	}

	s := make([]byte, len(spec))
	copy(s, spec)

	v := spec2version(spec)

	g.mod[spec] = v
	if g.root.Path == "" {
		g.root = v
	}

	bare := Version{Path: v.Path}
	g.versions[bare] = append(g.versions[bare], v)

	return v
}

func (g *Graph) addEdge(moda, modb string) {
	a := g.addMod(moda)
	b := g.addMod(modb)

	g.uses[a] = append(g.uses[a], b)
	g.usedBy[b] = append(g.usedBy[b], a)
}

func (g *Graph) tidy() {
	for k, vs := range g.versions {
		sort.Slice(vs, func(i, j int) bool { return vs[i].LessThan(vs[j]) })
		g.versions[k] = vs
	}

	for k, nodes := range g.uses {
		sort.Slice(nodes, func(i, j int) bool { return nodes[i].LessThan(nodes[j]) })
		g.uses[k] = nodes
	}

	for k, nodes := range g.usedBy {
		sort.Slice(nodes, func(i, j int) bool { return nodes[i].LessThan(nodes[j]) })
		g.usedBy[k] = nodes
	}
}

func (g *Graph) pathTo(v Version) ([]Version, error) {
	if v == g.root {
		return []Version{v}, nil
	}

	pf := newPathFinder(g)
	path, err := pf.findPath(g.root, v)
	if err != nil {
		return nil, err
	}

	return path, nil
}

func spec2version(spec string) Version {
	idx := strings.IndexByte(spec, '@')

	v := Version{Path: spec}
	if idx >= 0 {
		v.Path = spec[:idx]
		v.Version = spec[idx+1:]
	}

	return v
}
