package parse

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var timeRe = regexp.MustCompile(`\d{1,2}`)
var dayRe = regexp.MustCompile(`(\p{Han})`)

var weekday = []string{
	"日",
	"月",
	"火",
	"水",
	"木",
	"金",
	"土",
}

func NormalizeTime(times string) (string, error) {
	rTime := timeRe.FindAllStringSubmatch(times, -1)
	rDay := dayRe.FindStringSubmatch(times)

	if len(rTime) < 4 {
		return "", errors.New("NormalizeTime has error")
	}

	// get time
	month, _ := strconv.Atoi(rTime[0][0])
	day, _ := strconv.Atoi(rTime[1][0])
	hour, _ := strconv.Atoi(rTime[2][0])
	min, _ := strconv.Atoi(rTime[3][0])
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
	for _, s := range weekday {
		if strings.EqualFold(day, s) {
			break
		}

		counter++
	}

	if counter == len(weekday)-1 {
		// if dayOfWeek is "土"
		// return "日"
		return weekday[0]
	}

	n := counter + 1
	return weekday[n]
}

func parseTime(startTime string, year string) (time.Time, error) {
	// day of a week replace jaoanese to english (for strings.Replace)
	replaceRule := []string{
		"日", "Sunday",
		"月", "Monday",
		"火", "Tuesday",
		"水", "Wednesday",
		"木", "Thursday",
		"金", "Friday",
		"土", "Saturday",
	}

	normalized, err := NormalizeTime(startTime)
	if err != nil {
		return time.Time{}, err
	}

	r := strings.NewReplacer(replaceRule...)
	s := r.Replace(normalized)

	layout := "1/2(Monday) 15:04"

	t, err := time.Parse(layout, s)
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
