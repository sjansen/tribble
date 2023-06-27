package main

import (
	"fmt"
	"os"

	"github.com/cli/go-gh/v2"
	"github.com/cli/go-gh/v2/pkg/repository"
)

func die(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {
	repo, err := repository.Current()
	if err != nil {
		die(err)
	}
	fmt.Printf("%s/%s/%s\n", repo.Host, repo.Owner, repo.Name)

	dir, err := os.MkdirTemp("", repo.Name)
	if err != nil {
		die(err)
	}

	// gh repo clone <repository> [<directory>] [-- <gitflags>...]
	args := []string{"repo", "clone", repo.Name, dir, "--", "--bare"}
	stdout, stderr, err := gh.Exec(args...)
	if err != nil {
		die(err)
	}
	fmt.Println(args)
	fmt.Println(stdout.String())
	fmt.Println(stderr.String())
}
