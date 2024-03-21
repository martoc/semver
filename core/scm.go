package core

import (
	"time"

	"github.com/blang/semver/v4"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/martoc/semver/logger"
)

//go:generate ${GOPATH}/bin/mockgen -destination=./commit_iter_mock.go -package=core github.com/go-git/go-git/v5/plumbing/object CommitIter
//go:generate ${GOPATH}/bin/mockgen -destination=./reference_iter_mock.go -package=core github.com/go-git/go-git/v5/plumbing/storer ReferenceIter
//go:generate ${GOPATH}/bin/mockgen -source=scm.go -destination=./scm_mock.go -package=core

// CommitLog represents a commit in the Git repository.
type CommitLog struct {
	Hash       string            // The commit hash.
	Tags       []*semver.Version // The tags associated with the commit.
	Message    string            // The commit message.
	Date       time.Time         // The commit date.
	Author     string            // The author of the commit.
	Head       bool              // Indicates if the commit is the HEAD commit.
	BranchName string            // The name of the branch the commit belongs to.
}

// Scm is an interface that defines the methods for interacting with a source control management system.
type Scm interface {
	GetCommitLog() ([]*CommitLog, error) // GetCommitLog retrieves the commit history of the Git repository.
	Tag(name, hash string, floating bool) error
	Push() error
}

// GitRepo is an interface that defines the methods for interacting with a Git repository.
type GitRepo interface {
	PlainOpen(string) error
	Head() (*plumbing.Reference, error)
	Log(*git.LogOptions) (object.CommitIter, error)
	Tags() (storer.ReferenceIter, error)
	CommitObject(plumbing.Hash) (*object.Commit, error)
	CreateTag(name string, hash plumbing.Hash, opts *git.CreateTagOptions) (*plumbing.Reference, error)
	DeleteTag(name string) error
	Push(opts *git.PushOptions) error
}

// GitRepoImpl is an implementation of the GitRepo interface.
type GitRepoImpl struct {
	repo *git.Repository
}

// PlainOpen opens a Git repository at the specified path.
// It initializes the GitRepoImpl struct with the opened repository.
// Returns an error if the repository cannot be opened.
func (g *GitRepoImpl) PlainOpen(path string) error {
	repo, err := git.PlainOpen(path)
	g.repo = repo

	return err
}

// Head returns the reference to the HEAD commit in the Git repository.
func (g *GitRepoImpl) Head() (*plumbing.Reference, error) {
	return g.repo.Head()
}

// Log returns a commit iterator for the Git repository, based on the provided options.
// It retrieves the commit history of the repository.
// The returned commit iterator can be used to iterate over the commits in the repository.
// The options parameter allows specifying various options for filtering and sorting the commit history.
// It returns the commit iterator and an error if any occurred.
func (g *GitRepoImpl) Log(options *git.LogOptions) (object.CommitIter, error) {
	return g.repo.Log(options)
}

// Tags returns an iterator over the tags in the Git repository.
// It returns a storer.ReferenceIter that can be used to iterate over the tags.
// If an error occurs, it is returned along with the iterator.
func (g *GitRepoImpl) Tags() (storer.ReferenceIter, error) {
	return g.repo.Tags()
}

// CommitObject returns the commit object with the given hash.
func (g *GitRepoImpl) CommitObject(hash plumbing.Hash) (*object.Commit, error) {
	return g.repo.CommitObject(hash)
}

// CreateTag creates a new tag with the given name and hash in the Git repository.
// It returns a reference to the newly created tag and any error encountered.
func (g *GitRepoImpl) CreateTag(name string, hash plumbing.Hash, opts *git.CreateTagOptions) (*plumbing.Reference, error) {
	return g.repo.CreateTag(name, hash, opts)
}

// DeleteTag deletes the tag with the given name from the Git repository.
// It returns an error if the tag deletion fails.
func (g *GitRepoImpl) DeleteTag(name string) error {
	return g.repo.DeleteTag(name)
}

// Push pushes the changes to the remote repository.
func (g *GitRepoImpl) Push(opts *git.PushOptions) error {
	return g.repo.Push(opts)
}

// ScmGit is an implementation of the Scm interface for Git repositories.
type ScmGit struct {
	Path string
	Repo GitRepo
}

// ScmGitBuilder is a builder for creating ScmGit instances.
type ScmGitBuilder struct {
	Path string
	Repo GitRepo
}

// NewScmGitBuilder creates a new ScmGitBuilder instance.
func NewScmGitBuilder() *ScmGitBuilder {
	return &ScmGitBuilder{}
}

// SetPath sets the path of the Git repository.
func (b *ScmGitBuilder) SetPath(path string) *ScmGitBuilder {
	b.Path = path

	return b
}

// SetRepo sets the Git repository implementation.
func (b *ScmGitBuilder) SetRepo(repo GitRepo) *ScmGitBuilder {
	b.Repo = repo

	return b
}

// Build creates a new Scm instance based on the builder configuration.
func (b *ScmGitBuilder) Build() Scm {
	if b.Repo == nil {
		b.Repo = &GitRepoImpl{}
	}

	return &ScmGit{
		Path: b.Path,
		Repo: b.Repo,
	}
}

// GetCommitLog retrieves the commit history of the Git repository.
// It returns a slice of CommitLog structs representing each commit,
// along with associated information such as the commit hash, message,
// tags, author, and date.
// If an error occurs during the retrieval process, it is returned as the second value.
func (s *ScmGit) GetCommitLog() ([]*CommitLog, error) {
	// Open the Git repository
	err := s.Repo.PlainOpen(s.Path)
	if err != nil {
		return nil, err
	}

	// Get the HEAD reference
	ref, err := s.Repo.Head()
	if err != nil {
		return nil, err
	}

	// Retrieve the commit history starting from HEAD
	commitIter, err := s.Repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return nil, err
	}

	commitLogs := []*CommitLog{}
	// Iterate over commits and display commit information using for loop
	for {
		commit, errCommitIter := commitIter.Next()
		if commit == nil || errCommitIter != nil {
			break
		}

		isHead := false
		// Get the tags associated with the commit
		tags, errTags := s.Repo.Tags()
		if errTags != nil {
			logger.GetInstance().Error(errTags)
		}

		tagNames := s.getTags(commit, tags)

		if ref.Hash() == commit.Hash {
			isHead = true
		}

		commitLogs = append(commitLogs, &CommitLog{
			Hash:    commit.Hash.String(),
			Message: commit.Message,
			Tags:    tagNames,
			Head:    isHead,
			Author:  commit.Author.Name,
			Date:    commit.Author.When,
		})
	}

	return commitLogs, nil //nolint:nilerr
}

// getTags returns a slice of semver.Version representing the tags associated with the given commit.
// It takes a commit object and a reference iterator as parameters.
// The function iterates over the tags and checks if the tag's commit hash matches the given commit's hash.
// If a match is found, the tag name is cleaned and converted into a semver.Version object.
// The cleaned tag names are then appended to the tagNames slice.
// Finally, the function returns the tagNames slice.
func (s *ScmGit) getTags(commit *object.Commit, tags storer.ReferenceIter) []*semver.Version {
	tagNames := []*semver.Version{}

	for {
		tag, errTagIter := tags.Next()
		if tag == nil || errTagIter != nil {
			break
		}

		tagCommit, errCommit := s.Repo.CommitObject(tag.Hash())
		if errCommit != nil {
			logger.GetInstance().Error(tag.Name(), " - ", tag.Hash(), ": ", errCommit)
		}

		if tagCommit != nil && tagCommit.Hash == commit.Hash {
			version, errSemver := semver.Make(s.cleanVersion(tag.Name().Short()))
			if errSemver != nil {
				logger.GetInstance().Error(tag.Name().Short(), ": ", errSemver)
			} else {
				tagNames = append(tagNames, &version)
			}
		}
	}

	return tagNames
}

// cleanVersion removes the leading 'v' character from the given tagName if it exists.
// It returns the cleaned tagName.
func (s *ScmGit) cleanVersion(tagName string) string {
	if tagName[0] == 'v' {
		return tagName[1:]
	}

	return tagName
}

// Tag creates a new tag with the given name and hash in the Git repository.
// It returns an error if the tag creation fails.
func (s *ScmGit) Tag(name, hash string, floating bool) error {
	commitHash := plumbing.NewHash(hash)

	if floating {
		err := s.Repo.DeleteTag(name)
		if err != nil {
			logger.GetInstance().Println(err)
		}
	}

	_, err := s.Repo.CreateTag(name, commitHash, nil)

	return err
}

// Push pushes the changes to the remote repository.
// It returns an error if the push operation fails.
func (s *ScmGit) Push() error {
	return s.Repo.Push(&git.PushOptions{
		RemoteName: "origin",
		RefSpecs: []config.RefSpec{
			config.RefSpec("refs/tags/*:refs/tags/*"), // Push tags to the remote repository
		},
	})
}
