package cli

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArgParser(t *testing.T) {
	for _, tc := range []struct {
		Args     []string
		Expected interface{}
	}{{
		Args:     []string{"init"},
		Expected: initCmd{},
	}, {
		Args: []string{"new", "foo", "bar"},
		Expected: newCmd{
			Src: "foo",
			Dst: "bar",
		},
	}, {
		Args:     []string{"update"},
		Expected: initCmd{},
	}} {
		t.Run(strings.Join(tc.Args, " "), func(t *testing.T) {
			require := require.New(t)
			tc := tc

			ctx := parse(tc.Args)
			actual := ctx.Selected().Target.Interface()
			require.EqualValues(tc.Expected, actual)
		})
	}
}
