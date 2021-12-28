package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/client"
	"github.com/go-git/go-git/v5/storage/memory"
)

const url = "https://github.com/sjansen/tribble"

func main() {
	ctx := context.TODO()
	fs := memfs.New()
	storer := memory.NewStorage()

	refname, err := findDefaultRefName(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	_, err = git.CloneContext(ctx, storer, fs, &git.CloneOptions{
		URL:           url,
		ReferenceName: refname,
		SingleBranch:  true,
		Depth:         1,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	license, err := fs.Open("LICENSE")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	io.Copy(os.Stdout, license)
}

func findDefaultRefName(url string) (plumbing.ReferenceName, error) {
	var name plumbing.ReferenceName

	e, err := transport.NewEndpoint(url)
	if err != nil {
		return name, err
	}

	cli, err := client.NewClient(e)
	if err != nil {
		return name, err
	}

	s, err := cli.NewUploadPackSession(e, nil)
	if err != nil {
		return name, err
	}

	info, err := s.AdvertisedReferences()
	if err != nil {
		return name, err
	}

	refs, err := info.AllReferences()
	if err != nil {
		return name, err
	}

	return refs["HEAD"].Target(), nil
}
