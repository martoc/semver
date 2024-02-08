package core_test

import (
	"testing"
	"time"

	"github.com/blang/semver/v4"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/golang/mock/gomock"
	"github.com/martoc/semver/core"
	"github.com/stretchr/testify/assert"
)

func TestScmGit_GetCommitLog(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// Create a mock GitRepo
	mockRepo := core.NewMockGitRepo(ctrl)
	mockCommitIter := core.NewMockCommitIter(ctrl)
	mockReferenceIter := core.NewMockReferenceIter(ctrl)

	// Set up expectations on the mock Repo
	mockRepo.EXPECT().PlainOpen(gomock.Any()).Return(nil)
	mockRepo.EXPECT().Head().Return(plumbing.NewHashReference("refs/branches/main",
		plumbing.NewHash("e574dfaecd0a2a1d666c19f813c9a8f573fc121b")), nil)

	commit := &object.Commit{
		Hash:    plumbing.NewHash("e574dfaecd0a2a1d666c19f813c9a8f573fc121b"),
		Message: "Commit message",
		Author: object.Signature{
			Name: "Sarah Connor",
			When: time.Now(),
		},
	}

	mockCommitIter.EXPECT().Next().Return(commit, nil)
	mockCommitIter.EXPECT().Next().Return(nil, nil)
	mockRepo.EXPECT().Log(&git.LogOptions{From: commit.Hash}).Return(mockCommitIter, nil)

	tagRef := plumbing.NewReferenceFromStrings("refs/tags/v1.0.0", "e574dfaecd0a2a1d666c19f813c9a8f573fc121b")
	mockReferenceIter.EXPECT().Next().Return(tagRef, nil)
	mockReferenceIter.EXPECT().Next().Return(nil, nil)
	mockRepo.EXPECT().Tags().Return(mockReferenceIter, nil)
	mockRepo.EXPECT().CommitObject(tagRef.Hash()).Return(commit, nil)

	// Create the ScmGit instance with the mock Repo
	scm := core.ScmGit{
		Path: "/path/to/repo",
		Repo: mockRepo,
	}

	// Call the method under test
	commitLogs, err := scm.GetCommitLog()

	// Assert the results
	assert.NoError(t, err)
	assert.Len(t, commitLogs, 1)

	expectedCommitLog := &core.CommitLog{
		Hash:    "e574dfaecd0a2a1d666c19f813c9a8f573fc121b",
		Message: "Commit message",
		Tags: []*semver.Version{
			{
				Major: 1,
				Minor: 0,
				Patch: 0,
			},
		},
		Head:   true,
		Author: "Sarah Connor",
		Date:   commit.Author.When,
	}

	assert.Equal(t, expectedCommitLog, commitLogs[0])
}
