package times

import "time"

func NowWithLoc(tz string) time.Time {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return time.Now()
	}
	return time.Now().In(loc)
}
