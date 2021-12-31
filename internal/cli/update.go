package cli

import "fmt"

type updateCmd struct{}

func (cmd *updateCmd) Run() error {
	fmt.Println("update")
	return nil
}
