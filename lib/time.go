package lib

import (
	"time"
)

type Season int

const (
	Unknown Season = iota
	Spring
	Summer
	Autumn
	Winter
)

func (s Season) String() string {
	switch s {
	case Spring:
		return "spring"
	case Summer:
		return "summer"
	case Autumn:
		return "autumn"
	case Winter:
		return "winter"
	}

	return "unknown"
}

func (s Season) StringJa() string {
	switch s {
	case Spring:
		return "春季"
	case Summer:
		return "夏季"
	case Autumn:
		return "秋季"
	case Winter:
		return "冬季(新春)"
	}

	return "unknown"
}

func GetSeason(t time.Time) Season {
	month := t.Month()

	switch {
	case 1 <= month && month <= 3:
		return Winter
	case 4 <= month && month <= 6:
		return Spring
	case 7 <= month && month <= 9:
		return Summer
	case 10 <= month && month <= 12:
		return Autumn
	}

	return Unknown
}

func GetNoWSeason() Season {
	t := time.Now()

	return GetSeason(t)
}
