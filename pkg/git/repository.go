package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

const (
	head = "HEAD"

	defaultCommitCount = 10
)

type Repository interface {
	LogExcludeTo(fromRev, toRev string) ([]Commit, error)
}

var _ Repository = (*repo)(nil)

type repo struct {
	r *git.Repository
}

type stopFn func(*object.Commit) error

func NewRepository(path string) (Repository, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	return &repo{
		r: r,
	}, nil
}

// Get all commits between <from revision> and <to revision> (exclude <to revision>)
func (r *repo) LogExcludeTo(fromRev, toRev string) ([]Commit, error) {
	if fromRev == "" {
		fromRev = head
	}

	fromHash, err := r.r.ResolveRevision(plumbing.Revision(fromRev))
	if err != nil {
		return nil, err
	}

	if toRev == "" {
		return r.logWithStopFnFirst(fromHash, nil)
	}

	toHash, err := r.r.ResolveRevision(plumbing.Revision(toRev))
	if err != nil {
		return nil, err
	}

	return r.logWithStopFnFirst(fromHash, func(c *object.Commit) error {
		if c.Hash == *toHash {
			return storer.ErrStop
		}

		return nil
	})
}

func (r *repo) logWithStopFnFirst(fromHash *plumbing.Hash, fn stopFn) ([]Commit, error) {
	cIter, err := r.r.Log(&git.LogOptions{
		From: *fromHash,
	})
	if err != nil {
		return nil, err
	}

	commits := make([]Commit, 0, defaultCommitCount)

	if err := cIter.ForEach(func(c *object.Commit) error {
		if fn != nil {
			if err := fn(c); err != nil {
				return err
			}
		}

		commit := newCommit(c)
		commits = append(commits, commit)

		return nil
	}); err != nil {
		return nil, err
	}

	return commits, nil
}
