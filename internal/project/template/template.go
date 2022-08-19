package template

import (
	"context"
	neturl "net/url"
	"os"
	"path/filepath"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"

	"github.com/sjansen/tribble/internal/git"
)

const configPath = "_tribble/config.toml"
const questionsPath = "_tribble/questions.toml"

type Template struct {
	fs      billy.Filesystem
	url     string
	refname string
}

func New(fs billy.Filesystem) *Template {
	return &Template{
		fs: fs,
	}
}

func Open(ctx context.Context, url, refname string) (t *Template, err error) {
	src, err := neturl.Parse(url)
	if err != nil {
		return nil, err
	} else if src.Scheme == "" {
		src := filepath.Clean(url)
		if _, err := os.Stat(src); err != nil {
			return nil, err
		}
		return &Template{
			fs:  osfs.New(src),
			url: src,
		}, nil
	}

	if refname == "" && src.Scheme != "" {
		r, err := git.FindDefaultRefName(ctx, url)
		if err != nil {
			return nil, err
		}
		refname = r.String()
	}

	fs, err := git.Clone(ctx, url, refname)
	if err != nil {
		return nil, err
	}
	return &Template{
		fs:      fs,
		url:     url,
		refname: refname,
	}, nil
}

func (t *Template) Origin(dst string) (url, refname string, err error) {
	if t.refname == "" {
		origin := t.url
		if !filepath.IsAbs(t.url) {
			src, err := filepath.Abs(t.url)
			if err != nil {
				return "", "", err
			}
			origin, err = filepath.Rel(dst, src)
			if err != nil {
				return "", "", err
			}
		}
		return origin, "", nil
	}
	return t.url, t.refname, nil
}
