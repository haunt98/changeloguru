package git

import "time"

type Author struct {
	When  time.Time
	Name  string
	Email string
}
