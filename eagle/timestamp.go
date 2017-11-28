package eagle

import (
	"strconv"
	"time"
)

type timestamp string

// Seconds is the seconds since the unix epoch. Returns 0 if error
func (t timestamp) Time() time.Time {
	if v, err := strconv.ParseUint(string(t), 0, 64); err != nil {
		return time.Unix(0, 0)
	} else {
		// seconds Jan 1, 2000 00:00:00 since the unix epoch
		return time.Unix(int64(v+946598400), 0)
	}
}
