package cli

import "fmt"

type initCmd struct{}

func (cmd *initCmd) Run() error {
	fmt.Println("init")
	return nil
}
