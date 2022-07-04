package changelog

import (
	"fmt"
	"time"

	"github.com/make-go-great/date-go"
)

func generateVersionHeaderValue(version string, when time.Time) string {
	return fmt.Sprintf("%s (%s)", version, date.FormatDateByDefault(when, time.Local))
}
