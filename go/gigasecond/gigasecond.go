package gigasecond

import "time"

// A gigaSecond is one followed by 9 zeros.
const gigaSecond = 1e9

// AddGigasecond add a gigasecond to the time given.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond * time.Second)
}
