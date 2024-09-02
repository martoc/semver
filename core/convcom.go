package core

import (
	"regexp"
	"strings"
)

const (
	MAJOR SemanticVersionComponent = iota
	MINOR
	PATCH
)

const (
	commitTypeGroup                 = 2
	breakingChangeGroup             = 3
	lenghtOfMatch                   = 2
	lenghtOfMatchWithBreakingChange = 3
)

type SemanticVersionComponent int

// GetVersionUpdate determines the version update type (MAJOR, MINOR, PATCH) based on the conventional commit message.
func GetVersionUpdate(commitMessage string) SemanticVersionComponent {
	// Regular expression to match commit types, including optional ! before :
	re := regexp.MustCompile(`^(Merged PR \d+: )?(feat|fix|chore|docs|style|refactor|perf|test|BREAKING CHANGE)(!?)(\(.*\))?: .*`)

	// Extract the commit type from the commit message
	match := re.FindStringSubmatch(commitMessage)
	if len(match) > lenghtOfMatch {
		commitType := match[commitTypeGroup]

		// Check if the commit message contains "BREAKING CHANGE:" or has ! in the type
		if strings.Contains(commitMessage, "BREAKING CHANGE") ||
			(len(match) > lenghtOfMatchWithBreakingChange && match[breakingChangeGroup] == "!") {
			return MAJOR
		}

		// Map commit types to version updates
		versionMap := map[string]SemanticVersionComponent{
			"feat":     MINOR,
			"fix":      PATCH,
			"chore":    PATCH,
			"docs":     PATCH,
			"style":    PATCH,
			"refactor": PATCH,
			"perf":     PATCH,
			"test":     PATCH,
		}

		// Return the corresponding version update
		if version, ok := versionMap[commitType]; ok {
			return version
		}
	}

	// Default to PATCH if no match is found
	return PATCH
}
