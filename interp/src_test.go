package interp

import (
	"testing"
)

func Test_effectivePkg(t *testing.T) {
	testCases := []struct {
		desc     string
		root     string
		path     string
		expected string
	}{
		{
			desc:     "path is a subpackage",
			root:     "github.com/foo/plugin/vendor/guthib.com/containous/fromage",
			path:     "guthib.com/containous/fromage/couteau/lol",
			expected: "github.com/foo/plugin/vendor/guthib.com/containous/fromage/couteau/lol",
		},
		{
			desc:     "path is a vendored package",
			root:     "github.com/foo/plugin/vendor/guthib.com/containous/fromage",
			path:     "vendor/guthib.com/containous/vin",
			expected: "github.com/foo/plugin/vendor/guthib.com/containous/fromage/vendor/guthib.com/containous/vin",
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			pkg := effectivePkg(test.root, test.path)

			if pkg != test.expected {
				t.Errorf("Got %s, want %s", pkg, test.expected)
			}
		})
	}
}