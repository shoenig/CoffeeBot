package utils

import "rand"

func SecsToNSecs(seconds int64) int64 {
	return seconds * 1000000000
}

// rand rand int [low, high)
func RandInt(low, high int) int {
    dv := high - low
    i := rand.Int() % dv
    return i + low
}
