// Package core provides the core functionality for version calculation and tagging.
package core

import (
	"strconv"

	"github.com/blang/semver/v4"
	"github.com/martoc/semver/logger"
)

const (
	versionPrefix = "v"
)

// CalculateOutput represents the output of the version calculation.
type CalculateOutput struct {
	NextVersion          string `json:"next_version"`
	FloatingVersionMajor string `json:"floating_version_major"`
	FloatingVersionMinor string `json:"floating_version_minor"`
}

// CalculateCommandBuilder is a builder for creating CalculateCommand instances.
type CalculateCommandBuilder struct {
	Scm             Scm
	Path            string
	AddFloatingTags bool
	Push            bool
	DisableTagging  bool
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

// SetAddFloatingTags sets the AddFloatingTags field of the CalculateCommandBuilder.
// It takes a boolean parameter 'addFloatingTags' and assigns it to the 'AddFloatingTags' field of the CalculateCommandBuilder.
// It returns a pointer to the CalculateCommandBuilder for method chaining.
func (b *CalculateCommandBuilder) SetAddFloatingTags(addFloatingTags bool) *CalculateCommandBuilder {
	b.AddFloatingTags = addFloatingTags

	return b
}

// SetPush sets the Push field of the CalculateCommandBuilder.
// It takes a boolean parameter 'push' and assigns it to the 'Push' field of the CalculateCommandBuilder.
// It returns a pointer to the CalculateCommandBuilder for method chaining.
func (b *CalculateCommandBuilder) SetPush(push bool) *CalculateCommandBuilder {
	b.Push = push

	return b
}

// SetDisableTagging sets the DisableTagging field of the CalculateCommandBuilder.
// It takes a boolean parameter 'disableTagging' and assigns it to the 'DisableTagging' field of the CalculateCommandBuilder.
// It returns a pointer to the CalculateCommandBuilder for method chaining.
func (b *CalculateCommandBuilder) SetDisableTagging(disableTagging bool) *CalculateCommandBuilder {
	b.DisableTagging = disableTagging

	return b
}

// Build returns a Command built from the CalculateCommandBuilder.
// It creates a CalculateCommandImpl with the provided Scm.
func (b *CalculateCommandBuilder) Build() Command {
	if b.Scm == nil {
		b.Scm = NewScmGitBuilder().SetPath(b.Path).Build()
	}

	return &CalculateCommandImpl{
		Scm:             b.Scm,
		AddFloatingTags: b.AddFloatingTags,
		Push:            b.Push,
	}
}

// CalculateCommandImpl represents an implementation of the CalculateCommand interface.
// It contains a Command and Scm field.
type CalculateCommandImpl struct {
	Command
	Scm             Scm
	AddFloatingTags bool
	Push            bool
	DisableTagging  bool
}

// Execute executes the CalculateCommandImpl command and returns the next version tag string and any error encountered.
func (c *CalculateCommandImpl) Execute() (interface{}, error) {
	var output CalculateOutput

	commitLogs, err := c.Scm.GetCommitLog()
	if err != nil {
		return "", err
	}

	nextTag := c.calculateTag(commitLogs)

	if c.AddFloatingTags {
		floatingVersionMajor := strconv.FormatInt(int64(nextTag.Major), 10)

		if !c.DisableTagging {
			err = c.Scm.Tag(versionPrefix+floatingVersionMajor, commitLogs[0].Hash, true) // vx
			if err != nil {
				logger.GetInstance().Println(err)
			}
		}

		output.FloatingVersionMajor = floatingVersionMajor

		floatingVersionMinor := floatingVersionMajor + "." + strconv.FormatInt(int64(nextTag.Minor), 10)

		if !c.DisableTagging {
			err = c.Scm.Tag(versionPrefix+floatingVersionMinor, commitLogs[0].Hash, true) // vx.y
			if err != nil {
				logger.GetInstance().Println(err)
			}
		}

		output.FloatingVersionMinor = floatingVersionMinor
	}

	if !c.DisableTagging {
		err = c.Scm.Tag(versionPrefix+nextTag.String(), commitLogs[0].Hash, false) // vx.y.z
		if err != nil {
			logger.GetInstance().Println(err)
		}
	}

	output.NextVersion = nextTag.String()

	if c.Push && !c.DisableTagging {
		err = c.Scm.Push()
		if err != nil {
			logger.GetInstance().Println(err)

			return "", err
		}
	}

	return output, nil
}

// calculateTag calculates the next version tag based on the commit logs.
// It takes a slice of CommitLog pointers and returns a pointer to the calculated semver.Version.
func (c *CalculateCommandImpl) calculateTag(commitLogs []*CommitLog) *semver.Version {
	nextTag, _ := semver.Make("0.0.0")

	if len(commitLogs) > 0 && len(commitLogs[0].Tags) > 0 {
		nextTag = c.GetGreatestTag(nextTag, commitLogs[0].Tags)

		return &nextTag
	}

	for _, commit := range commitLogs {
		nextTag = c.GetGreatestTag(nextTag, commit.Tags)
	}

	updateType := GetVersionUpdate(commitLogs[0].Message)

	switch updateType {
	case MAJOR:
		nextTag.IncrementMajor() //nolint: errcheck
	case MINOR:
		nextTag.IncrementMinor() //nolint: errcheck
	case PATCH:
		nextTag.IncrementPatch() //nolint: errcheck
	default:
		nextTag.IncrementPatch() //nolint: errcheck
	}

	return &nextTag
}

// GetGreatestTag returns the greatest tag from a list of tags.
// It takes a semver.Version and a slice of semver.Version pointers and returns the greatest semver.Version.
func (c *CalculateCommandImpl) GetGreatestTag(nextTag semver.Version, tags []*semver.Version) semver.Version {
	for _, tag := range tags {
		if nextTag.LT(*tag) {
			nextTag = *tag
		}
	}

	return nextTag
}
