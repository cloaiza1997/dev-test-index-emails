package functions

import (
	"fmt"
	"time"
)

func FormatTime() (time.Time, string) {
	now := time.Now()

	return now, now.Format("2006-01-02 15:04:05")
}

func HandleReturnError(message string) (bool, int, int, []string) {
	fmt.Println(message)

	return false, 0, 0, []string{}
}
