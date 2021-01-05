package git

import (
	"fmt"

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
	Log(fromRev, toRev string) ([]Commit, error)
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

// Get all commits between <from revision> and <to revision>
func (r *repo) Log(fromRev, toRev string) ([]Commit, error) {
	if fromRev == "" {
		fromRev = head
	}

	fromHash, err := r.r.ResolveRevision(plumbing.Revision(fromRev))
	if err != nil {
		return nil, fmt.Errorf("failed to resolve %s: %w", fromRev, err)
	}

	if toRev == "" {
		return r.logWithStopFnLast(fromHash, nil)
	}

	toHash, err := r.r.ResolveRevision(plumbing.Revision(toRev))
	if err != nil {
		return nil, fmt.Errorf("failed to resolve %s: %w", toRev, err)
	}

	return r.logWithStopFnLast(fromHash, stopAtHash(toHash))
}

func (r *repo) logWithStopFnLast(fromHash *plumbing.Hash, fn stopFn) ([]Commit, error) {
	cIter, err := r.r.Log(&git.LogOptions{
		From: *fromHash,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to git log: %w", err)
	}

	commits := make([]Commit, 0, defaultCommitCount)

	if err := cIter.ForEach(func(c *object.Commit) error {
		commit := newCommit(c)
		commits = append(commits, commit)

		if fn == nil {
			return nil
		}

		return fn(c)
	}); err != nil {
		return nil, fmt.Errorf("failed to iterate each git log: %w", err)
	}

	return commits, nil
}

func stopAtHash(hash *plumbing.Hash) stopFn {
	return func(c *object.Commit) error {
		if c.Hash == *hash {
			return storer.ErrStop
		}

		return nil
	}
}
