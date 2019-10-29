package common

import "fmt"

const na string = "N/A"

// Versioning data
var (
	Version    string
	BuildDate  string
	GitCommit  string
	GitBranch  string
	GitSummary string
)

// BuildString ...
func BuildString() string {
	if GitCommit == "" {
		return na
	}
	return fmt.Sprintf("%s@%s, %s", GitBranch, GitCommit, BuildDate)
}

// VersionString ...
func VersionString() string {
	if Version == "" {
		return na
	}
	return Version
}

// VersionedBuildString ...
func VersionedBuildString() string {
	v := Version
	gc := GitCommit
	if v == "" {
		v = na
	}
	if gc == "" {
		gc = na
	}
	return fmt.Sprintf("%s, %s@%s, %s", v, GitBranch, gc, BuildDate)
}
