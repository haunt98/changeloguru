package changelog

import "github.com/haunt98/changeloguru/internal/convention"

const (
	addedType  = "Added"
	fixedType  = "Fixed"
	othersType = "Others"
	buildType  = "Build"
)

// The order when generate changelog
var changelogTypes = []string{
	addedType,
	fixedType,
	othersType,
	buildType,
}

func getType(conventionType string) string {
	switch conventionType {
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
