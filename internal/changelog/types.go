package changelog

import "github.com/haunt98/changeloguru/internal/convention"

const (
	addedType  = "Added"
	fixedType  = "Fixed"
	othersType = "Others"
	buildType  = "Build"

	depsScope = "deps"
)

// The order when generate changelog
var changelogTypes = []string{
	addedType,
	fixedType,
	othersType,
	buildType,
}

func getType(conventionCommit convention.Commit) string {
	if conventionCommit.Scope == depsScope {
		return buildType
	}

	switch conventionCommit.Type {
	case convention.FeatType:
		return addedType
	case convention.FixType:
		return fixedType
	case convention.BuildType, convention.CIType:
		return buildType
	default:
		return othersType
	}
}
