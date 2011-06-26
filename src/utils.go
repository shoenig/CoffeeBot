package utils

import "fmt"
import "rand"
import "time"

func SecsToNSecs(seconds int64) int64 {
	return seconds * 1000000000
}

// rand rand int [low, high)
func RandInt(low, high int) int {
	dv := high - low
	i := rand.Int() % dv
	return i + low
}

// returns HH:MM
func SimpleTime() string {
	t := time.LocalTime()
	return fmt.Sprintf(" (%v:%v)", t.Hour, t.Minute)
}
