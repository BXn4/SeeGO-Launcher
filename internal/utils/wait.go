package utils

import "time"

func WaitUntil(condition func() bool, timeout time.Duration) bool {
	startTime := time.Now()
	for !condition() {
		if time.Since(startTime) >= timeout {
			return true
		}
		time.Sleep(10 * time.Millisecond)
	}
	return false
}
