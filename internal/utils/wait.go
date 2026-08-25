package utils

import (
	"time"
)

// Timeout to condition
// true = TIMED OUT!
func WaitUntil(condition func() bool, timeout time.Duration) bool {
	done := make(chan bool, 1)

	go func() {
		done <- condition()
	}()

	select {
	case result := <-done:
		return !result
	case <-time.After(timeout):
		return true
	}
}
