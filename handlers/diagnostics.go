package handlers

import "time"

var StartTime time.Time

func UpTime() string {
	return time.Since(StartTime).Round(time.Second).String()
}
