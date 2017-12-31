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
	dayOfWeek := rDay[0]

	// normalize
	if hour >= 24 {
		hour = hour - 24
		day = day + 1
		dayOfWeek = nextDayOfAWeek(dayOfWeek)
	}

	normalized := fmt.Sprintf("%d/%d(%s) %d:%02d", month, day, dayOfWeek, hour, min)

	return normalized, nil
}

func nextDayOfAWeek(day string) string {
	var dayOfAWeek = []string{
		"日",
		"月",
		"火",
		"水",
		"木",
		"金",
		"土",
	}

	counter := 0
	for _, s := range dayOfAWeek {
		if strings.EqualFold(day, s) {
			break
		}

		counter++
	}

	if counter == len(dayOfAWeek)-1 {
		// if dayOfWeek is "土"
		// return "日"
		return dayOfAWeek[0]
	}

	n := counter + 1
	return dayOfAWeek[n]
}

func parseTime(startTime string) (time.Time, error) {
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

	return t, nil
}
