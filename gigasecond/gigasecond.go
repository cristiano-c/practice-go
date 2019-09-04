// Package gigasecond is funny
package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond calculates the moment when someone has lived for 10^9 seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(int(math.Pow(10, 9))))
}
