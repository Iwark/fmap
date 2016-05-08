package fmap

import (
	"errors"
	"time"
)

var timeFormats = []string{
	"1/2/2006",
	"1/2/2006 15:4:5",
	"2006-1-2 15:4:5",
	"2006-1-2 15:4",
	"2006-1-2",
	"1-2",
	"15:4:5",
	"15:4",
	"15",
	"15:4:5 Jan 2, 2006 MST",
}

func parseTime(str string) (t time.Time, err error) {
	for _, format := range timeFormats {
		t, err = time.Parse(format, str)
		if err == nil {
			location := t.Location()
			if location.String() == "UTC" {
				location = time.Now().Location()
			}
			pt := []int{t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}
			t = time.Date(pt[5], time.Month(pt[4]), pt[3], pt[2], pt[1], pt[0], 0, location)
			return
		}
	}
	err = errors.New("Can't parse string as time: " + str)
	return
}
