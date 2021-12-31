package cli

import (
	"fmt"
)

type newCmd struct {
	Src string `kong:"arg,help='Path or URL of project template.'"`
	Dst string `kong:"arg,help='Path to new project.'"`
}

func (cmd *newCmd) Run() error {
	fmt.Println(cmd.Src)
	fmt.Println(cmd.Dst)
	return nil
}
