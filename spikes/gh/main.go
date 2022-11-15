package main

import (
	"fmt"
	"log"
	"github.com/cli/go-gh"
)

func main() {
		repo, err := gh.CurrentRepository()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s/%s/%s\n", repo.Host(), repo.Owner(), repo.Name())
}
