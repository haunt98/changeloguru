package changelog

import (
	"fmt"
	"time"

	"github.com/make-go-great/date-go"
)

func generateVersionHeaderValue(ver string, when time.Time) string {
	return fmt.Sprintf("%s (%s)", ver, date.FormatDateByDefault(when, time.Local))
}
