package changelog

import "github.com/haunt98/changeloguru/pkg/convention"

const (
	addedType  = "Added"
	fixedType  = "Fixed"
	othersType = "Others"
)

func getType(conventionType string) string {
	switch conventionType {
	case convention.FeatType:
		return addedType
	case convention.FixType:
		return fixedType
	default:
		return othersType
	}
}
