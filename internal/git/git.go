package git

import (
	"context"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/client"
	"github.com/go-git/go-git/v5/storage/memory"
)

func Clone(ctx context.Context, url, refname string) (billy.Filesystem, error) {
	// TODO verify refname if provided
	fs := memfs.New()
	storer := memory.NewStorage()
	_, err := git.CloneContext(ctx, storer, fs, &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName(refname),
		SingleBranch:  true,
		Depth:         1,
	})
	if err != nil {
		return nil, err
	}

	return fs, nil
}

func FindDefaultRefName(ctx context.Context, url string) (plumbing.ReferenceName, error) {
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

	info, err := s.AdvertisedReferencesContext(ctx)
	if err != nil {
		return name, err
	}

	refs, err := info.AllReferences()
	if err != nil {
		return name, err
	}

	return refs["HEAD"].Target(), nil
}
