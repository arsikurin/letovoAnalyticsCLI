package utils

import (
	"errors"
	"regexp"
	"time"
)

func ParseDay(dayRaw string) (time.Weekday, error) {
	isMonday, _ := regexp.MatchString(`(?i)^mo`, dayRaw)
	isTuesday, _ := regexp.MatchString(`(?i)^tu`, dayRaw)
	isWednesday, _ := regexp.MatchString(`(?i)^we`, dayRaw)
	isThursday, _ := regexp.MatchString(`(?i)^th`, dayRaw)
	isFriday, _ := regexp.MatchString(`(?i)^fr`, dayRaw)
	isSaturday, _ := regexp.MatchString(`(?i)^sa`, dayRaw)

	switch true {
	case isMonday:
		return time.Monday, nil
	case isTuesday:
		return time.Tuesday, nil
	case isWednesday:
		return time.Wednesday, nil
	case isThursday:
		return time.Thursday, nil
	case isFriday:
		return time.Friday, nil
	case isSaturday:
		return time.Saturday, nil
	default:
		return -1, errors.New("cannot parse the specific day you specified")
	}
}
