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

func TestCalculateCommandImpl_ShouldReturnMayorVersion(t *testing.T) {
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

	// Create CalculateCommandImpl with the mock Scm
	calculateCommand := &core.CalculateCommandImpl{Scm: mockScm}

	// Call Execute method
	result, err := calculateCommand.Execute()

	// Assert the result
	assert.Equal(t, "2.1.0", result)
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
	calculateCommand := &core.CalculateCommandImpl{Scm: mockScm}

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
