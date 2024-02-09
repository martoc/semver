// Package core provide some functions.
package core

import (
	"github.com/blang/semver/v4"
)

// CalculateCommand is an interface for the CalculateCommandImpl.
type CalculateCommandBuilder struct {
	Scm  Scm
	Path string
}

// NewCalculateCommandBuilder creates a new instance of CalculateCommandBuilder.
// It returns a pointer to the newly created CalculateCommandBuilder.
func NewCalculateCommandBuilder() *CalculateCommandBuilder {
	return &CalculateCommandBuilder{}
}

// SetScm sets the source control management (SCM) for the CalculateCommandBuilder.
// It takes an Scm parameter and returns a pointer to the CalculateCommandBuilder.
func (b *CalculateCommandBuilder) SetScm(scm Scm) *CalculateCommandBuilder {
	b.Scm = scm

	return b
}

// SetPath sets the path for the CalculateCommandBuilder.
// It takes a string parameter 'path' and assigns it to the 'Path' field of the CalculateCommandBuilder.
// It returns a pointer to the CalculateCommandBuilder for method chaining.
func (b *CalculateCommandBuilder) SetPath(path string) *CalculateCommandBuilder {
	b.Path = path

	return b
}

// Build returns a Command built from the CalculateCommandBuilder.
// It creates a CalculateCommandImpl with the provided Scm.
func (b *CalculateCommandBuilder) Build() Command {
	if b.Scm == nil {
		b.Scm = NewScmGitBuilder().SetPath(b.Path).Build()
	}

	return &CalculateCommandImpl{
		Scm: b.Scm,
	}
}

// CalculateCommandImpl represents an implementation of the CalculateCommand interface.
// It contains a Command and Scm field.
type CalculateCommandImpl struct {
	Command
	Scm Scm
}

// Execute executes the CalculateCommandImpl command and returns the next tag version based on the commit logs.
// It retrieves the commit logs using the Scm interface and iterates through each commit to find the highest tag version.
// The highest tag version is returned as a string.
// If an error occurs while retrieving the commit logs, an empty string and the error are returned.
func (c *CalculateCommandImpl) Execute() (string, error) {
	commitLogs, err := c.Scm.GetCommitLog()
	if err != nil {
		return "", err
	}

	nextTag, _ := semver.Make("0.0.0")

	for _, commit := range commitLogs {
		if commit.Tags != nil {
			for _, tag := range commit.Tags {
				if nextTag.LT(*tag) {
					nextTag = *tag
				}
			}
		}
	}

	nextTag.IncrementMinor() //nolint: errcheck

	return nextTag.String(), nil
}
