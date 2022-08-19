package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
	"unicode"

	"golang.org/x/text/unicode/rangetable"
)

/*
ID_Start = [\p{L}\p{Nl}\p{Other_ID_Start}-\p{Pattern_Syntax}-\p{Pattern_White_Space}]
ID_Continue = [\p{ID_Start}\p{Mn}\p{Mc}\p{Nd}\p{Pc}\p{Other_ID_Continue}-\p{Pattern_Syntax}-\p{Pattern_White_Space}]
*/

var ID_Start, ID_Continue, ID_Excluded *unicode.RangeTable

type data struct {
	ID_Start    string
	ID_Continue string
	ID_Excluded string
}

func init() {
	ID_Start = rangetable.Merge(
		unicode.L,
		unicode.Nl,
		unicode.Other_ID_Start,
	)
	ID_Continue = rangetable.Merge(
		ID_Start,
		unicode.Mn,
		unicode.Mc,
		unicode.Nd,
		unicode.Pc,
		unicode.Other_ID_Continue,
	)
	ID_Excluded = rangetable.Merge(
		unicode.Pattern_Syntax,
		unicode.Pattern_White_Space,
	)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "TEMPLATE [FILENAME]")
	}
	template := os.Args[1]
	filename := template
	if len(os.Args) > 2 {
		filename = os.Args[2]
	} else if strings.HasSuffix(filename, ".tmpl") {
		filename = filename[:len(filename)-5]
	}

	if err := write(template, filename); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func write(src, filename string) error {
	b, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	t, err := template.New("xid").Parse(string(b))
	if err != nil {
		return err
	}

	var f io.WriteCloser
	if filename == "-" {
		f = os.Stdout
	} else {
		f, err = os.Create(filename)
		if err != nil {
			return err
		}
	}
	defer func() {
		if f, ok := f.(interface{ Sync() }); ok {
			f.Sync()
		}
		if f != os.Stdout {
			f.Close()
		}
	}()

	data := &data{
		ID_Start:    fmt.Sprintf("%#v", ID_Start),
		ID_Continue: fmt.Sprintf("%#v", ID_Continue),
		ID_Excluded: fmt.Sprintf("%#v", ID_Excluded),
	}

	t.Execute(f, data)

	return nil
}
