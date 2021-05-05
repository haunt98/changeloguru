package git

import "time"

type Author struct {
	Name  string
	Email string
	When  time.Time
}
