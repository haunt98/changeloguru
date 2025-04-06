package git

import (
	"fmt"
	"slices"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/hashicorp/go-version"
)

const (
	head = "HEAD"

	defaultCommitCount = 128
	defaultTagCount    = 32
)

// Repository is an abstraction for git-repository
type Repository interface {
	Log(fromRev, toRev string) ([]Commit, error)
	Commit(commitMessage string, paths ...string) error
	SemVerTags() ([]SemVerTag, error)
}

type repo struct {
	gitRepo *git.Repository
}

type stopFn func(*object.Commit) error

// NewRepository return Repository from path
func NewRepository(path string) (Repository, error) {
	gitRepo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	return &repo{
		gitRepo: gitRepo,
	}, nil
}

// Log return all commits between <from revision> and <to revision>
func (r *repo) Log(fromRev, toRev string) ([]Commit, error) {
	if fromRev == "" {
		fromRev = head
	}

	fromHash, err := r.gitRepo.ResolveRevision(plumbing.Revision(fromRev))
	if err != nil {
		return nil, fmt.Errorf("failed to resolve %s: %w", fromRev, err)
	}

	if toRev == "" {
		return r.logWithStopFn(fromHash, nil, nil)
	}

	toHash, err := r.gitRepo.ResolveRevision(plumbing.Revision(toRev))
	if err != nil {
		return nil, fmt.Errorf("failed to resolve %s: %w", toRev, err)
	}

	return r.logWithStopFn(fromHash, nil, stopAtHash(toHash))
}

func (r *repo) Commit(commitMessage string, paths ...string) error {
	gitWorktree, err := r.gitRepo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get git worktree: %w", err)
	}

	for _, path := range paths {
		if err := gitWorktree.AddWithOptions(&git.AddOptions{
			Path: path,
		}); err != nil {
			return fmt.Errorf("failed to git add %s: %w", path, err)
		}
	}

	if _, err := gitWorktree.Commit(commitMessage, &git.CommitOptions{}); err != nil {
		return fmt.Errorf("failed to git commit: %w", err)
	}

	return nil
}

func (r *repo) logWithStopFn(fromHash *plumbing.Hash, beginFn, endFn stopFn) ([]Commit, error) {
	cIter, err := r.gitRepo.Log(&git.LogOptions{
		From: *fromHash,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to git log: %w", err)
	}

	commits := make([]Commit, 0, defaultCommitCount)

	if err := cIter.ForEach(newIterFn(&commits, beginFn, endFn)); err != nil {
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

func newIterFn(commits *[]Commit, beginFn, endFn stopFn) func(c *object.Commit) error {
	if beginFn == nil && endFn == nil {
		return func(c *object.Commit) error {
			commit := newCommit(c)
			*commits = append(*commits, commit)

			return nil
		}
	}

	if beginFn == nil {
		return func(c *object.Commit) error {
			commit := newCommit(c)
			*commits = append(*commits, commit)

			if err := endFn(c); err != nil {
				return err
			}

			return nil
		}
	}

	if endFn == nil {
		return func(c *object.Commit) error {
			if err := beginFn(c); err != nil {
				return err
			}

			commit := newCommit(c)
			*commits = append(*commits, commit)

			return nil
		}
	}

	return func(c *object.Commit) error {
		if err := beginFn(c); err != nil {
			return err
		}

		commit := newCommit(c)
		*commits = append(*commits, commit)

		if err := endFn(c); err != nil {
			return err
		}

		return nil
	}
}

// Return SemVer tags from oldest to latest
func (r *repo) SemVerTags() ([]SemVerTag, error) {
	iter, err := r.gitRepo.Tags()
	if err != nil {
		return nil, err
	}

	vers := make([]*version.Version, 0, defaultTagCount)

	if err := iter.ForEach(func(r *plumbing.Reference) error {
		ver, err := version.NewVersion(r.Name().Short())
		if err != nil {
			// Ignore bad tag
			return nil
		}

		vers = append(vers, ver)

		return nil
	}); err != nil {
		return nil, err
	}

	slices.SortFunc(vers, func(a, b *version.Version) int {
		return a.Compare(b)
	})

	tags := make([]SemVerTag, 0, len(vers))
	for _, ver := range vers {
		tags = append(tags, SemVerTag{
			Version: ver,
		})
	}

	return tags, nil
}
