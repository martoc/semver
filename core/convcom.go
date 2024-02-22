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

type SemanticVersionComponent int

// GetVersionUpdate determines the version update type (MAJOR, MINOR, PATCH) based on the conventional commit message.
func GetVersionUpdate(commitMessage string) SemanticVersionComponent {
	// Regular expression to match commit types, including optional ! before :
	re := regexp.MustCompile(`^(feat|fix|chore|docs|style|refactor|perf|test|BREAKING CHANGE)(!?)(\(.*\))?: .*`)

	// Extract the commit type from the commit message
	match := re.FindStringSubmatch(commitMessage)
	if len(match) > 1 {
		commitType := match[1]

		// Check if the commit message contains "BREAKING CHANGE:" or has ! in the type
		if strings.Contains(commitMessage, "BREAKING CHANGE") || (len(match) > 2 && match[2] == "!") {
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
