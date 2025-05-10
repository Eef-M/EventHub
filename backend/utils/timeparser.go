package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var jakartaTimezone = time.FixedZone("WIB", 7*60*60)

func ParseDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02"
	t, err := time.ParseInLocation(layout, dateStr, jakartaTimezone)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date: %v", err)
	}
	return t, nil
}

func ParseTime(timeStr string) (time.Time, error) {
	layout := "15:04"
	t, err := time.ParseInLocation(layout, timeStr, jakartaTimezone)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time: %v", err)
	}
	return t, nil
}

func RandomDateBetween(start, end time.Time, r *rand.Rand) time.Time {
	diff := end.Sub(start)
	randomSeconds := time.Duration(r.Int63n(int64(diff.Seconds()))) * time.Second
	return start.Add(randomSeconds)
}

func RandomTimeWIB(r *rand.Rand) time.Time {
	hour := 8 + r.Intn(13)
	minute := r.Intn(60)
	return time.Date(0, 1, 1, hour, minute, 0, 0, jakartaTimezone)
}
