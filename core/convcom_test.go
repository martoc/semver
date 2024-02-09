package core_test

import (
	"testing"

	"github.com/martoc/semver/core"
	"github.com/stretchr/testify/assert"
)

func Test_GetVersionUpdate_FeatureUpdate(t *testing.T) {
	t.Parallel()

	commitMessage := "feat: add new feature"
	expectedUpdate := core.MINOR
	result := core.GetVersionUpdate(commitMessage)
	assert.Equal(t, expectedUpdate, result,
		"Unexpected version update for feature update.")
}

func Test_GetVersionUpdate_BugFixUpdate(t *testing.T) {
	t.Parallel()

	commitMessage := "fix: resolve a bug"
	expectedUpdate := core.PATCH
	result := core.GetVersionUpdate(commitMessage)
	assert.Equal(t, expectedUpdate, result,
		"Unexpected version update for bug fix update.")
}

func Test_GetVersionUpdate_BreakingChangeChore(t *testing.T) {
	t.Parallel()

	commitMessage := "chore!: make breaking change"
	expectedUpdate := core.MAJOR
	result := core.GetVersionUpdate(commitMessage)
	assert.Equal(t, expectedUpdate, result,
		"Unexpected version update for breaking change in chore.")
}

func Test_GetVersionUpdate_RefactorUpdate(t *testing.T) {
	t.Parallel()

	commitMessage := "refactor: improve code"
	expectedUpdate := core.PATCH
	result := core.GetVersionUpdate(commitMessage)
	assert.Equal(t, expectedUpdate, result,
		"Unexpected version update for refactor update.")
}

func Test_GetVersionUpdate_DocsUpdateWithBreakingChange(t *testing.T) {
	t.Parallel()

	commitMessage := "docs: correct spelling with BREAKING CHANGE"
	expectedUpdate := core.MAJOR
	result := core.GetVersionUpdate(commitMessage)
	assert.Equal(t, expectedUpdate, result,
		"Unexpected version update for docs update with breaking change.")
}

func Test_GetVersionUpdate_TestUpdateWithBreakingChange(t *testing.T) {
	t.Parallel()

	commitMessage := "test: implement new functionality\n\nBREAKING CHANGE: modify test framework"
	expectedUpdate := core.MAJOR
	result := core.GetVersionUpdate(commitMessage)
	assert.Equal(t, expectedUpdate, result,
		"Unexpected version update for test update with breaking change.")
}

func Test_GetVersionUpdate_RefactorUpdateWithBreakingChange(t *testing.T) {
	t.Parallel()

	commitMessage := "refactor!: implement new functionality\n\nBREAKING CHANGE: modify test framework"
	expectedUpdate := core.MAJOR
	result := core.GetVersionUpdate(commitMessage)
	assert.Equal(t, expectedUpdate, result,
		"Unexpected version update for refactor update with breaking change.")
}

func Test_GetVersionUpdate_RefactorUpdateWithBreakingChangeAndTicket(t *testing.T) {
	t.Parallel()

	commitMessage := "refactor!(TICKET-12344): implement new functionality\n\nBREAKING CHANGE: modify test framework"
	expectedUpdate := core.MAJOR
	result := core.GetVersionUpdate(commitMessage)
	assert.Equal(t, expectedUpdate, result,
		"Unexpected version update for refactor update with breaking change and ticket.")
}
