package changelog

import (
	"fmt"
	"time"

	"github.com/haunt98/clock"
)

func generateVersionHeaderValue(version string, when time.Time) string {
	return fmt.Sprintf("%s (%s)", version, clock.FormatDate(when))
}
