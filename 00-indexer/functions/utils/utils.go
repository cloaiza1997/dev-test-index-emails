package functions

import (
	"time"
)

func FormatTime() (time.Time, string) {
	now := time.Now()

	return now, now.Format("2006-01-02 15:04:05")
}
