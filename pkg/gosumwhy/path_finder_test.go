package gosumwhy

import (
	"testing"
)

func TestFindPath(t *testing.T) {

	type testCase struct {
		target      Version
		expectedLen int
		hasErr      bool
	}

	gr, _ := readGraphFrom(goModGraphSample())
	testCases := []testCase{
		{Version{"root", ""}, 1, false},
		{Version{"mod1", "v2.0.0"}, 2, false},
		{Version{"mod2", "v1.1.0"}, 2, false},
		{Version{"mod2", "v0.9.0"}, 4, false},
		{Version{"testmodule", "v1.0.0"}, 4, false},
		{Version{"mod2", "v2.0.0"}, 0, true},
	}

	pf := newPathFinder(gr)
	pf.computeAllDistances(gr.root)

	for _, tc := range testCases {
		path, err := pf.extractPath(gr.root, tc.target)
		if !tc.hasErr && err != nil {
			t.Errorf("unexpected error when searching path to %s: %s", tc.target, err)
		} else if tc.hasErr && err == nil {
			t.Errorf("expected an error when searching path to %s, but got none", tc.target)
		}

		if len(path) != tc.expectedLen {
			t.Errorf("wrong length for path to %s: expected %d  got %d (path=%v)", tc.target, tc.expectedLen, len(path), path)
		}
		if len(path) > 0 && path[len(path)-1] != tc.target {
			t.Errorf("wrong path, does not lead to target %s: %v", tc.target, path)
		}
	}

}
