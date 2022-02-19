package utils

import (
	log "github.com/sirupsen/logrus"
	"regexp"
	"time"
)

func ParseDay(dayRaw string) time.Weekday {
	isMonday, _ := regexp.MatchString(`(?i)^mo`, dayRaw)
	isTuesday, _ := regexp.MatchString(`(?i)^tu`, dayRaw)
	isWednesday, _ := regexp.MatchString(`(?i)^we`, dayRaw)
	isThursday, _ := regexp.MatchString(`(?i)^th`, dayRaw)
	isFriday, _ := regexp.MatchString(`(?i)^fr`, dayRaw)
	isSaturday, _ := regexp.MatchString(`(?i)^sa`, dayRaw)

	switch true {
	case isMonday:
		return time.Monday
	case isTuesday:
		return time.Tuesday
	case isWednesday:
		return time.Wednesday
	case isThursday:
		return time.Thursday
	case isFriday:
		return time.Friday
	case isSaturday:
		return time.Saturday
	default:
		log.Fatalln("Cannot parse the specific day you provided")
		return -1
	}
}
