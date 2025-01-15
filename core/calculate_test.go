package core_test

import (
	"errors"
	"testing"

	"github.com/blang/semver/v4"
	"github.com/golang/mock/gomock"
	"github.com/martoc/semver/core"
	"github.com/stretchr/testify/assert"
)

var errExpectedFromTest = errors.New("some error")

func TestCalculateCommandImpl_ShouldReturnSameTaggedVersionShouldNotBumpVersion(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// Create a mock Scm
	mockScm := core.NewMockScm(ctrl)

	// Set up expectations for GetCommitLog method
	mockScm.EXPECT().GetCommitLog().Return([]*core.CommitLog{
		{
			Tags: []*semver.Version{
				{Major: 1, Minor: 0, Patch: 0},
				{Major: 2, Minor: 0, Patch: 2},
				{Major: 2, Minor: 0, Patch: 1},
			},
		},
	}, nil)

	mockScm.EXPECT().Tag("v2", gomock.Any(), true).Return(nil).Times(1)
	mockScm.EXPECT().Tag("v2.0", gomock.Any(), true).Return(nil).Times(1)
	mockScm.EXPECT().Tag("v2.0.2", gomock.Any(), false).Return(nil).Times(1)
	mockScm.EXPECT().Push().Return(nil).Times(1)

	// Create CalculateCommandImpl with the mock Scm
	calculateCommand := &core.CalculateCommandImpl{Scm: mockScm, AddFloatingTags: true, Push: true}

	// Call Execute method
	result, err := calculateCommand.Execute()

	// Assert the result
	assert.Equal(t, core.CalculateOutput{NextVersion: "2.0.2", FloatingVersionMajor: "2", FloatingVersionMinor: "2.0"}, result)
	assert.Nil(t, err)
}

func TestCalculateCommandImpl_ShouldReturnSameTaggedVersionShouldIncreaseMayor(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// Create a mock Scm
	mockScm := core.NewMockScm(ctrl)

	// Set up expectations for GetCommitLog method
	mockScm.EXPECT().GetCommitLog().Return([]*core.CommitLog{
		{
			Tags:    []*semver.Version{},
			Message: "feat!: add new feature",
		},
		{
			Tags: []*semver.Version{
				{Major: 1, Minor: 0, Patch: 0},
				{Major: 2, Minor: 0, Patch: 2},
				{Major: 2, Minor: 0, Patch: 1},
			},
		},
	}, nil)

	mockScm.EXPECT().Tag("v3", gomock.Any(), true).Return(nil).Times(1)
	mockScm.EXPECT().Tag("v3.0", gomock.Any(), true).Return(nil).Times(1)
	mockScm.EXPECT().Tag("v3.0.0", gomock.Any(), false).Return(nil).Times(1)
	mockScm.EXPECT().Push().Return(nil).Times(1)

	// Create CalculateCommandImpl with the mock Scm
	calculateCommand := &core.CalculateCommandImpl{Scm: mockScm, AddFloatingTags: true, Push: true}

	// Call Execute method
	result, err := calculateCommand.Execute()

	// Assert the result
	assert.Equal(t, core.CalculateOutput{NextVersion: "3.0.0", FloatingVersionMajor: "3", FloatingVersionMinor: "3.0"}, result)
	assert.Nil(t, err)
}

func TestCalculateCommandImpl_ShouldIncreaseMinor(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// Create a mock Scm
	mockScm := core.NewMockScm(ctrl)

	// Set up expectations for GetCommitLog method
	mockScm.EXPECT().GetCommitLog().Return([]*core.CommitLog{
		{
			Tags:    []*semver.Version{},
			Message: "feat: add new feature",
		},
		{
			Tags: []*semver.Version{
				{Major: 1, Minor: 0, Patch: 0},
				{Major: 2, Minor: 0, Patch: 2},
				{Major: 2, Minor: 0, Patch: 1},
			},
		},
	}, nil)

	mockScm.EXPECT().Tag("v2", gomock.Any(), true).Return(nil).Times(1)
	mockScm.EXPECT().Tag("v2.1", gomock.Any(), true).Return(nil).Times(1)
	mockScm.EXPECT().Tag("v2.1.0", gomock.Any(), false).Return(nil).Times(1)
	mockScm.EXPECT().Push().Return(nil).Times(1)

	// Create CalculateCommandImpl with the mock Scm
	calculateCommand := &core.CalculateCommandImpl{Scm: mockScm, AddFloatingTags: true, Push: true}

	// Call Execute method
	result, err := calculateCommand.Execute()

	// Assert the result
	assert.Equal(t, core.CalculateOutput{NextVersion: "2.1.0", FloatingVersionMajor: "2", FloatingVersionMinor: "2.1"}, result)
	assert.Nil(t, err)
}

func TestCalculateCommandImpl_ShouldIncreasePatch(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// Create a mock Scm
	mockScm := core.NewMockScm(ctrl)

	// Set up expectations for GetCommitLog method
	mockScm.EXPECT().GetCommitLog().Return([]*core.CommitLog{
		{
			Tags:    []*semver.Version{},
			Message: "fix: add new feature",
		},
		{
			Tags: []*semver.Version{
				{Major: 1, Minor: 0, Patch: 0},
				{Major: 2, Minor: 0, Patch: 2},
				{Major: 2, Minor: 0, Patch: 1},
			},
		},
	}, nil)

	mockScm.EXPECT().Tag("v2", gomock.Any(), true).Return(nil).Times(1)
	mockScm.EXPECT().Tag("v2.0", gomock.Any(), true).Return(nil).Times(1)
	mockScm.EXPECT().Tag("v2.0.3", gomock.Any(), false).Return(nil).Times(1)
	mockScm.EXPECT().Push().Return(nil).Times(1)

	// Create CalculateCommandImpl with the mock Scm
	calculateCommand := &core.CalculateCommandImpl{Scm: mockScm, AddFloatingTags: true, Push: true}

	// Call Execute method
	result, err := calculateCommand.Execute()

	// Assert the result
	assert.Equal(t, core.CalculateOutput{NextVersion: "2.0.3", FloatingVersionMajor: "2", FloatingVersionMinor: "2.0"}, result)
	assert.Nil(t, err)
}

func TestCalculateCommandImpl_ShouldIncreasePatchIfByDefault(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// Create a mock Scm
	mockScm := core.NewMockScm(ctrl)

	// Set up expectations for GetCommitLog method
	mockScm.EXPECT().GetCommitLog().Return([]*core.CommitLog{
		{
			Tags:    []*semver.Version{},
			Message: "whatever: add new feature",
		},
		{
			Tags: []*semver.Version{
				{Major: 1, Minor: 0, Patch: 0},
				{Major: 2, Minor: 0, Patch: 2},
				{Major: 2, Minor: 0, Patch: 1},
			},
		},
	}, nil)

	mockScm.EXPECT().Tag("v2", gomock.Any(), true).Return(nil).Times(1)
	mockScm.EXPECT().Tag("v2.0", gomock.Any(), true).Return(nil).Times(1)
	mockScm.EXPECT().Tag("v2.0.3", gomock.Any(), false).Return(nil).Times(1)
	mockScm.EXPECT().Push().Return(nil).Times(1)

	// Create CalculateCommandImpl with the mock Scm
	calculateCommand := &core.CalculateCommandImpl{Scm: mockScm, AddFloatingTags: true, Push: true}

	// Call Execute method
	result, err := calculateCommand.Execute()

	// Assert the result
	assert.Equal(t, core.CalculateOutput{NextVersion: "2.0.3", FloatingVersionMajor: "2", FloatingVersionMinor: "2.0"}, result)
	assert.Nil(t, err)
}

func TestCalculateCommandImpl_ShouldReturnNextVersion(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// Create a mock Scm
	mockScm := core.NewMockScm(ctrl)

	// Set up expectations for GetCommitLog method
	mockScm.EXPECT().GetCommitLog().Return([]*core.CommitLog{
		{
			Tags:    []*semver.Version{},
			Message: "feat: add new feature",
		},
		{
			Tags: []*semver.Version{
				{Major: 1, Minor: 0, Patch: 0},
				{Major: 2, Minor: 0, Patch: 2},
				{Major: 2, Minor: 0, Patch: 1},
			},
		},
	}, nil)

	mockScm.EXPECT().Tag("v2", gomock.Any(), true).Return(nil).Times(1)
	mockScm.EXPECT().Tag("v2.1", gomock.Any(), true).Return(nil).Times(1)
	mockScm.EXPECT().Tag("v2.1.0", gomock.Any(), false).Return(nil).Times(1)
	mockScm.EXPECT().Push().Return(nil).Times(1)

	// Create CalculateCommandImpl with the mock Scm
	calculateCommand := &core.CalculateCommandImpl{Scm: mockScm, AddFloatingTags: true, Push: true}

	// Call Execute method
	result, err := calculateCommand.Execute()

	// Assert the result
	assert.Equal(t, core.CalculateOutput{NextVersion: "2.1.0", FloatingVersionMajor: "2", FloatingVersionMinor: "2.1"}, result)
	assert.Nil(t, err)
}

func TestCalculateCommandImpl_ShouldReturnError(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// Create a mock Scm
	mockScm := core.NewMockScm(ctrl)

	// Set up expectations for GetCommitLog method
	mockScm.EXPECT().GetCommitLog().Return(nil, errExpectedFromTest).Times(1)

	// Create CalculateCommandImpl with the mock Scm
	calculateCommand := &core.CalculateCommandImpl{Scm: mockScm, AddFloatingTags: true, Push: true}

	// Call Execute method
	result, resultError := calculateCommand.Execute()

	// Assert the result when there's an error
	assert.Equal(t, resultError, errExpectedFromTest)
	assert.Empty(t, result)
}

func TestCalculateCommandBuilder_ShouldSetScm(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// Create a mock Scm
	mockScm := core.NewMockScm(ctrl)

	// Create a CalculateCommandBuilder instance
	builder := &core.CalculateCommandBuilder{}

	// Call SetScm method
	result := builder.SetScm(mockScm)

	// Assert the Scm field is set correctly
	assert.Equal(t, mockScm, result.Scm)
}

func TestCalculateCommandBuilder_Build(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// Create a mock Scm
	mockScm := core.NewMockScm(ctrl)

	// Create a CalculateCommandBuilder instance
	builder := &core.CalculateCommandBuilder{}
	builder.SetScm(mockScm)

	// Call Build method
	command := builder.Build()

	_, ok := command.(*core.CalculateCommandImpl)

	// Assert the Scm field is set correctly
	assert.True(t, ok)
	assert.Equal(t, mockScm, command.(*core.CalculateCommandImpl).Scm) //nolint:forcetypeassert
}

func TestNewCalculateCommandBuilder(t *testing.T) {
	t.Parallel()

	// Call NewCalculateCommandBuilder
	builder := core.NewCalculateCommandBuilder()

	// Assert the builder is not nil
	assert.NotNil(t, builder)
}

func TestCalculateCommandBuilder_SetPath(t *testing.T) {
	t.Parallel()

	// Create a CalculateCommandBuilder instance
	builder := &core.CalculateCommandBuilder{}

	// Call SetPath method
	result := builder.SetPath("/path/to/file")

	// Assert the Path field is set correctly
	assert.Equal(t, "/path/to/file", result.Path)
}

func TestCalculateCommandBuilder_SetScm(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// Create a mock Scm
	mockScm := core.NewMockScm(ctrl)

	// Create a CalculateCommandBuilder instance
	builder := &core.CalculateCommandBuilder{}

	// Call SetScm method
	result := builder.SetScm(mockScm)

	// Assert the Scm field is set correctly
	assert.Equal(t, mockScm, result.Scm)
}

func TestCalculateCommandBuilder_BuildWithNilScm(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// Create a CalculateCommandBuilder instance
	builder := &core.CalculateCommandBuilder{}

	// Call Build method
	command := builder.Build()

	_, ok := command.(*core.CalculateCommandImpl)

	// Assert the Scm field is set correctly
	assert.True(t, ok)
	assert.NotNil(t, command.(*core.CalculateCommandImpl).Scm) //nolint:forcetypeassert
}

func TestCalculateCommandBuilder_BuildWithExistingScm(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// Create a mock Scm
	mockScm := core.NewMockScm(ctrl)

	// Create a CalculateCommandBuilder instance
	builder := &core.CalculateCommandBuilder{}
	builder.SetScm(mockScm)

	// Call Build method
	command := builder.Build()

	_, ok := command.(*core.CalculateCommandImpl)

	// Assert the Scm field is set correctly
	assert.True(t, ok)
	assert.Equal(t, mockScm, command.(*core.CalculateCommandImpl).Scm) //nolint:forcetypeassert
}

func TestCalculateCommandImpl_ShouldReturnSameTaggedVersionShouldIncreaseMayorButSkipTaggingAndPush(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// Create a mock Scm
	mockScm := core.NewMockScm(ctrl)

	// Set up expectations for GetCommitLog method
	mockScm.EXPECT().GetCommitLog().Return([]*core.CommitLog{
		{
			Tags:    []*semver.Version{},
			Message: "feat!: add new feature",
		},
		{
			Tags: []*semver.Version{
				{Major: 1, Minor: 0, Patch: 0},
				{Major: 2, Minor: 0, Patch: 2},
				{Major: 2, Minor: 0, Patch: 1},
			},
		},
	}, nil)

	mockScm.EXPECT().Tag("v3", gomock.Any(), true).Return(nil).Times(0)
	mockScm.EXPECT().Tag("v3.0", gomock.Any(), true).Return(nil).Times(0)
	mockScm.EXPECT().Tag("v3.0.0", gomock.Any(), false).Return(nil).Times(0)
	mockScm.EXPECT().Push().Return(nil).Times(0)

	// Create CalculateCommandImpl with the mock Scm
	calculateCommand := &core.CalculateCommandImpl{Scm: mockScm, AddFloatingTags: true, Push: true, DisableTagging: true}

	// Call Execute method
	result, err := calculateCommand.Execute()

	// Assert the result
	assert.Equal(t, core.CalculateOutput{NextVersion: "3.0.0", FloatingVersionMajor: "3", FloatingVersionMinor: "3.0"}, result)
	assert.Nil(t, err)
}
