package gosumwhy

import "fmt"

type pathFinder struct {
	g *Graph
	d map[Version]int
}

func newPathFinder(g *Graph) *pathFinder {
	return &pathFinder{
		g: g,
		d: make(map[Version]int),
	}
}

func (h *pathFinder) findPath(from, to Version) ([]Version, error) {
	fwd := h.g.uses
	bak := h.g.usedBy

	h.d = make(map[Version]int)
	h.d[from] = 1

	var stack = []Version{from}

searchloop:
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]

		dist := h.d[cur]

		for _, next := range fwd[cur] {
			if h.d[next] != 0 && h.d[next] <= dist+1 {
				continue
			}

			h.d[next] = dist + 1
			stack = append(stack, next)

			if next == to {
				break searchloop
			}
		}
	}

	if h.d[to] == 0 {
		// 'to' was not found by searching exhaustively through links described by 'fwd' starting from 'from'
		return nil, fmt.Errorf("module '%s' was not found in dependencies for '%s'", to, from)
	}

	node := to
	dist := h.d[node]
	var result = make([]Version, dist)
	for dist > 0 {
		dist--

		result[dist] = node
		for _, p := range bak[node] {
			if h.d[p] != dist {
				continue
			}

			node = p
			break
		}
	}

	return result, nil
}

func (h *pathFinder) computeDistances(from Version, targets []Version) {
	fwd := h.g.uses

	h.d = make(map[Version]int)
	h.d[from] = 1

	var stack = []Version{from}

	// compute distances to all nodes in graph in a breadth first order,
	// stop computing when all targets have their distance computed
fillloop:
	for i := 0; i < len(targets); i++ {
		tgt := targets[i]
		if h.d[tgt] > 0 {
			continue
		}

		for len(stack) > 0 {
			if h.d[tgt] > 0 {
				continue fillloop
			}

			cur := stack[0]
			stack = stack[1:]

			dist := h.d[cur]

			for _, next := range fwd[cur] {
				if h.d[next] != 0 && h.d[next] <= dist+1 {
					continue
				}

				h.d[next] = dist + 1
				stack = append(stack, next)
			}
		}
	}
}

func (h *pathFinder) computeAllDistances(from Version) {
	fwd := h.g.uses

	h.d = make(map[Version]int)
	h.d[from] = 1

	var stack = []Version{from}

	// compute distances to all nodes in graph in a breadth first order
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]

		dist := h.d[cur]

		for _, next := range fwd[cur] {
			if h.d[next] != 0 && h.d[next] <= dist+1 {
				continue
			}

			h.d[next] = dist + 1
			stack = append(stack, next)
		}
	}
}

func (h *pathFinder) extractPath(from, to Version) ([]Version, error) {
	bak := h.g.usedBy
	if h.d[to] == 0 {
		// 'to' was not found by searching exhaustively through links described by 'fwd' starting from 'from'
		return nil, fmt.Errorf("module '%s' was not found in dependencies for '%s'", to, from)
	}

	node := to
	dist := h.d[node]
	var result = make([]Version, dist)
	for dist > 1 {
		dist--

		found := false
		result[dist] = node
		for _, p := range bak[node] {
			if h.d[p] != dist {
				continue
			}

			found = true
			node = p
			break
		}

		if !found {
			// if '.extractPath(from, to)' is called after '.computeAllDistances(from)', this should not happen
			return nil, fmt.Errorf("could not find a path from '%s' to '%s", from, to)
		}
	}
	result[0] = node

	return result, nil
}
