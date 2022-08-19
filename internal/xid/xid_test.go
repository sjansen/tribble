package xid_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/tribble/internal/xid"
)

var TestCases = []struct {
	S  string
	Ok bool
}{
	{"x", true},
	{" x", false},
	{"x y", false},
	{"-x", false},
	{"x-y", false},
	{".x", false},
	{"x.y", false},
	{"/", false},
	{"x/y", false},
	{"0", false},
	{"x0", true},
	{":", false},
	{"x:", false},
	{"_", false},
	{"_x", false},
	{"\u00B7", false},
	{"x\u00B7", true},
	{"\u2118", true},
	{"x\u2118", true},
	{"\uFE4D", false},
	{"x\uFE4D", true},
	{"\U0001D7D8", false},
	{"x\U0001D7D8", true},
	{"Über", true},
}

func TestNormalize(t *testing.T) {
	require := require.New(t)

	for i, tc := range TestCases {
		_, ok := xid.Normalize(tc.S)
		require.Equalf(
			tc.Ok, ok,
			"tc = %q (%d)", tc.S, i,
		)
	}
}
