//go:generate go run ../../scripts/xid/main.go generated.go.tmpl
//go:generate gofmt -s -w generated.go
package xid

import (
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func Normalize(s string) (string, bool) {
	ok := true
	s = norm.NFKC.String(s)
	for i, r := range s {
		if i == 0 {
			if unicode.Is(id_excluded, r) {
				return s, false
			}
			if !unicode.Is(id_start, r) {
				return s, false
			}
		} else {
			if unicode.Is(id_excluded, r) {
				return s, false
			}
			if !unicode.Is(id_continue, r) {
				return s, false
			}
		}
	}
	return s, ok
}
