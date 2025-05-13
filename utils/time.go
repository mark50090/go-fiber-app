package utils

import (
	"fmt"
	"strconv"
	"time"
)

func TimeStr2UTC(str string, layout string, loc *time.Location) (t time.Time, err error) {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	date, err := time.Parse(layout, str)
	if err != nil {
		return t, err
	}

	newTime := time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		date.Hour(),
		date.Minute(),
		date.Second(),
		0,
		loc,
	)
	t = newTime.UTC()
	return t, nil
}

func Time2Str(t time.Time, layout string) (result string) {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	result = t.Format(layout)
	return result
}

func ParseDate(dateStr string) (time.Time, error) {
	layouts := []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05Z",
		"02-01-2006",
		"02-01-2006 15:04:05",
		"01/02/2006",
		"01/02/2006 15:04:05",
	}

	var parsedTime time.Time
	var err error
	for _, layout := range layouts {
		parsedTime, err = time.ParseInLocation(layout, dateStr, time.Local)
		if err == nil {
			return parsedTime, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse date: %s", dateStr)
}

func Str2Time(str string, layout string) (t time.Time, err error) {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	t, err = time.ParseInLocation(layout, str, time.Local)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func GetStartTimeOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func GetEndTimeOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local)
}

func GetStartEndDateByBudgetYear(budgetYear string) (startDate, endDate time.Time) {

	// Parse budget year to int
	year, err := strconv.Atoi(budgetYear)
	if err != nil {
		return time.Time{}, time.Time{}
	}

	ceYear := year

	// Budget year starts from October 1st of previous year
	startDate = time.Date(ceYear-1, time.October, 1, 0, 0, 0, 0, time.Local)

	// Budget year ends on September 30th of current year
	endDate = time.Date(ceYear, time.September, 30, 23, 59, 59, 0, time.Local)

	return startDate, endDate

}

func GetThaiYear(year string) (thaiYear int) {
	thaiYear, _ = strconv.Atoi(year)
	thaiYear += 543
	return thaiYear
}

func FormatDate(date time.Time) string {
	return date.Local().Format("02/01/2006 15:04:05")
}
