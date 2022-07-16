package gosumwhy

import (
	"regexp"
	"strconv"
	"strings"
)

// Version represents a module and its version. The struct is copy/pasted from golang.org/x/mod/module
type Version struct {
	Path    string
	Version string `json:",omitempty"`
}

// String returns a string representation 'path@version'
func (v Version) String() string {
	if v.Version != "" {
		return v.Path + "@" + v.Version
	}
	return v.Path
}

// LessThan compares two versions. If two
func (v Version) LessThan(o Version) bool {
	if v.Path != o.Path {
		return v.Path < o.Path
	}
	return lessVersionsString(v.Version, o.Version)
}

var v000rx = regexp.MustCompile(`^v(0\.0\.0)-(([0-9]+)-([0-9a-f]+))$`)
var vXYZrx = regexp.MustCompile(`^v([0-9]+(?:\.[0-9]+)*)(?:-(.*))?(\+incompatible)?$`)

func lessSemVerString(v1, v2 string) bool {
	c1 := strings.Split(v1, ".")
	c2 := strings.Split(v2, ".")

	ll := len(c1)
	if ll < len(c2) {
		ll = len(c2)
	}

	for i := 0; i < ll; i++ {
		if i >= len(c2) {
			return false
		}
		if i >= len(c1) {
			return true
		}

		m1, _ := strconv.Atoi(c1[i])
		m2, _ := strconv.Atoi(c2[i])
		if m1 != m2 {
			return m1 < m2
		}
	}
	return false
}

func lessVersionsString(v1, v2 string) bool {
	if v1 == "" {
		return false
	}
	if v2 == "" {
		return true
	}

	chunks1 := v000rx.FindStringSubmatch(v1)
	if chunks1 == nil {
		chunks1 = vXYZrx.FindStringSubmatch(v1)
	}

	chunks2 := v000rx.FindStringSubmatch(v2)
	if chunks2 == nil {
		chunks2 = vXYZrx.FindStringSubmatch(v2)
	}

	if chunks1 == nil && chunks2 == nil {
		return v1 < v2
	}
	if chunks1 == nil {
		return false
	}
	if chunks2 == nil {
		return true
	}

	if chunks1[1] != chunks2[1] {
		return lessSemVerString(chunks1[1], chunks2[1])
	}

	if chunks1[2] != chunks2[2] {
		return chunks1[2] < chunks2[2]
	}

	return chunks1[3] < chunks2[3]
}
