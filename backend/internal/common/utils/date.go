package utils

import "time"

func Today() string {
	return time.Now().Format("2006-01-02")
}

func Yesterday() string {
	return time.Now().AddDate(0, 0, -1).Format("2006-01-02")
}
