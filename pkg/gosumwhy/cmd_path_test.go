package gosumwhy

import (
	"bytes"
	"strings"
	"testing"
)

func TestPath(t *testing.T) {
	var out bytes.Buffer

	gr, _ := readGraphFrom(goModGraphSample())
	type testCase struct {
		target      Version
		expectedLen int
		hasErr      bool
	}

	testCases := []testCase{
		{Version{"root", ""}, 1, false},
		{Version{"mod1", "v2.0.0"}, 2, false},
		{Version{"mod2", "v1.1.0"}, 2, false},
		{Version{"mod2", "v0.9.0"}, 4, false},
		{Version{"testmodule", "v1.0.0"}, 4, false},
		{Version{"mod2", "v2.0.0"}, 0, true},
	}

	for _, tc := range testCases {
		out.Reset()
		err := Path(gr, []string{tc.target.String()}, &Options{Out: &out})

		if !tc.hasErr && err != nil {
			t.Errorf("unexpected error when searching path to %s: %s", tc.target, err)
		} else if tc.hasErr && err == nil {
			t.Errorf("expected an error when searching path to %s, but got none", tc.target)
		}

		// remove the trailing "\n" from the output:
		output := strings.TrimSpace(out.String())
		lineCount := strings.Count(output, "\n")

		if lineCount != tc.expectedLen {
			t.Errorf("wrong length for path to %s: expected %d  got %d\noutput=\n%s", tc.target, tc.expectedLen, lineCount, output)
		}
		idx := strings.LastIndex(output, "\n")
		last := output[idx+1:]
		if lineCount > 0 && last != tc.target.String() {
			t.Errorf("wrong path, does not lead to target %s:\n%s", tc.target, output)
		}
	}
}

func ExamplePath() {
	// fake reading the module graph for 'rsc.io/quote' from stdin:
	stdin := strings.NewReader(`rsc.io/quote rsc.io/quote/v3@v3.0.0
	rsc.io/quote rsc.io/sampler@v1.3.0
	rsc.io/quote/v3@v3.0.0 rsc.io/smpler@v1.3.0
	rsc.io/sampler@v1.3.0 golang.org/x/text@v.0.0-20170915032832-14c0d48ead0c`)

	gr, _ := readGraphFrom(stdin)

	// search for a path to one or more specific modules, with default options:
	Path(gr, []string{"rsc.io/sampler@v1.3.0", "golang.org/x/text@v0.0.0-20170915032832-14c0d48ead0c"}, nil)

	// Output:
	// ----- rsc.io/sampler@v1.3.0
	// rsc.io/quote
	// rsc.io/sampler@v1.3.0
	// ----- golang.org/x/text@v0.0.0-20170915032832-14c0d48ead0c
	// rsc.io/quote
	// rsc.io/sampler@v1.3.0
	// golang.org/x/text@v0.0.0-20170915032832-14c0d48ead0c
}

func ExamplePath_most_recent() {
	// fake reading the module graph for 'filippo.io/age' (v1.1.0-rc.1) from stdin:
	stdin := strings.NewReader(`filippo.io/age filippo.io/edwards25519@v1.0.0-rc.1
	filippo.io/age golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5
	filippo.io/age golang.org/x/sys@v0.0.0-20210903071746-97244b99971b
	filippo.io/age golang.org/x/term@v0.0.0-20210615171337-6886f2dfbf5b
	golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5 golang.org/x/net@v0.0.0-20210226172049-e18ecbb05110
	golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5 golang.org/x/sys@v0.0.0-20210615035016-665e8c7367d1
	golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5 golang.org/x/term@v0.0.0-20201126162022-7de9c90e9dd1
	golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5 golang.org/x/text@v0.3.3
	golang.org/x/term@v0.0.0-20210615171337-6886f2dfbf5b golang.org/x/sys@v0.0.0-20210615035016-665e8c7367d1
	`)

	gr, _ := readGraphFrom(stdin)

	// search for a path to the most recent version of a module (no specific version named):
	Path(gr, []string{"golang.org/x/sys"}, nil)

	// Output:
	// ----- golang.org/x/sys@v0.0.0-20210903071746-97244b99971b
	// filippo.io/age
	// golang.org/x/sys@v0.0.0-20210903071746-97244b99971b
}

func ExamplePath_allv() {
	// fake reading the module graph for 'filippo.io/age' (v1.1.0-rc.1) from stdin:
	stdin := strings.NewReader(`filippo.io/age filippo.io/edwards25519@v1.0.0-rc.1
	filippo.io/age golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5
	filippo.io/age golang.org/x/sys@v0.0.0-20210903071746-97244b99971b
	filippo.io/age golang.org/x/term@v0.0.0-20210615171337-6886f2dfbf5b
	golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5 golang.org/x/net@v0.0.0-20210226172049-e18ecbb05110
	golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5 golang.org/x/sys@v0.0.0-20210615035016-665e8c7367d1
	golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5 golang.org/x/term@v0.0.0-20201126162022-7de9c90e9dd1
	golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5 golang.org/x/text@v0.3.3
	golang.org/x/term@v0.0.0-20210615171337-6886f2dfbf5b golang.org/x/sys@v0.0.0-20210615035016-665e8c7367d1
	`)

	gr, _ := readGraphFrom(stdin)

	// search for a path to *each version* of a module:
	Path(gr, []string{"golang.org/x/sys"}, &Options{AllVersions: true})

	// two dependency paths are printed, to show why two distinct version of 'golang.org/x/sys'
	// are present in the graph :

	// Output:
	// ----- golang.org/x/sys@v0.0.0-20210903071746-97244b99971b
	// filippo.io/age
	// golang.org/x/sys@v0.0.0-20210903071746-97244b99971b
	// ----- golang.org/x/sys@v0.0.0-20210615035016-665e8c7367d1
	// filippo.io/age
	// golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5
	// golang.org/x/sys@v0.0.0-20210615035016-665e8c7367d1
}
