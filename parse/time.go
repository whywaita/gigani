package parse

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var weekdayJa = []string{
	"日",
	"月",
	"火",
	"水",
	"木",
	"金",
	"土",
}

var timeRe = regexp.MustCompile(`\d{1,2}`)
var dayRe = regexp.MustCompile(`(\p{Han})`)

var (
	layoutTime   = "1/2(Monday) 15:04"
	layoutRegexp = "%d/%d(%s) %d:%02d"
)

func NormalizeTime(times string) (string, error) {
	// times format is layoutRegexp
	var month, day, hour, min int

	rTime := timeRe.FindAllStringSubmatch(times, -1)
	rDay := dayRe.FindStringSubmatch(times)

	switch {
	case len(rTime) >= 4:
		hour, _ = strconv.Atoi(rTime[2][0])
		min, _ = strconv.Atoi(rTime[3][0])
		day, _ = strconv.Atoi(rTime[1][0])
	case len(rTime) == 2:
		hour = 00
		min = 00
		day, _ = strconv.Atoi(rTime[1][0])
	case len(rTime) == 1:
		// only month
		month, _ = strconv.Atoi(rTime[0][0])
		hour = 00
		min = 00
		day = 1
	default:
		return "", errors.New("NormalizeTime has error, invalid parsed length")
	}

	// get time
	month, _ = strconv.Atoi(rTime[0][0])
	weekday := rDay[0]

	// normalize
	if hour >= 24 {
		hour = hour - 24
		day = day + 1
		weekday = nextWeekday(weekday)
	}

	normalized := fmt.Sprintf("%d/%d(%s) %d:%02d", month, day, weekday, hour, min)

	return normalized, nil
}

func nextWeekday(day string) string {
	counter := 0
	for _, s := range weekdayJa {
		if strings.EqualFold(day, s) {
			break
		}

		counter++
	}

	if counter == len(weekdayJa)-1 {
		// if dayOfWeek is "土"
		// return "日"
		return weekdayJa[0]
	}

	n := counter + 1
	return weekdayJa[n]
}

func ParseTime(times string, year string) (time.Time, error) {
	// day of a week replace japanese to english (for strings.Replace)
	replaceRule := []string{
		"日", "Sunday",
		"月", "Monday",
		"火", "Tuesday",
		"水", "Wednesday",
		"木", "Thursday",
		"金", "Friday",
		"土", "Saturday",
	}

	normalized, err := NormalizeTime(times)
	if err != nil {
		return time.Time{}, err
	}

	r := strings.NewReplacer(replaceRule...)
	s := r.Replace(normalized)

	t, err := time.Parse(layoutTime, s)
	if err != nil {
		return time.Time{}, err
	}

	// year add
	// time.Parse return year is 0000.
	// So, addDate year by title.
	y, err := strconv.Atoi(year)
	if err != nil {
		return time.Time{}, err
	}
	t = t.AddDate(y, 0, 0)

	// if month is 12, maybe that anime is last year.
	if t.Month() == 12 {
		t = t.AddDate(-1, 0, 0)
	}

	return t, nil
}
