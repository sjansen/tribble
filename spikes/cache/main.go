package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	ctx := context.TODO()

	repo, err := openFromCache(ctx, "https://github.com/sjansen/tribble")
	if err != nil {
		die(err)
	}

	err = cat(repo, "LICENSE")
	if err != nil {
		die(err)
	}
}

func cat(repo *git.Repository, path string) error {
	ref, err := repo.Head()
	if err != nil {
		return err
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return err
	}

	tree, err := commit.Tree()
	if err != nil {
		return err
	}

	f, err := tree.File(path)
	if err != nil {
		return err
	}

	r, err := f.Reader()
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, r)
	return nil
}

func openFromCache(ctx context.Context, url string) (*git.Repository, error) {
	fmt.Fprintln(os.Stderr, "Updating cache...")
	cache, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}
	cache = filepath.Join(cache, "tribble", "github.com", "sjansen", "tribble.git")

	create := false
	fi, err := os.Stat(cache)
	switch {
	case errors.Is(err, os.ErrNotExist):
		create = true
	case err != nil:
		return nil, err
	case !fi.IsDir():
		return nil, os.ErrExist
	}

	if create {
		return git.PlainCloneContext(ctx, cache, true, &git.CloneOptions{
			URL:      url,
			Progress: os.Stderr,
		})
	}

	repo, err := git.PlainOpen(cache)
	if err != nil {
		return nil, err
	}

	err = repo.FetchContext(ctx, &git.FetchOptions{
		// TODO RefSpecs: ...
		Progress: os.Stderr,
	})
	if errors.Is(err, git.NoErrAlreadyUpToDate) {
		err = nil
	}
	return repo, err
}
