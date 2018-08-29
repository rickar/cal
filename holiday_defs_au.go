// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"time"
)

// Holidays in Australia
var (
	AUNewYear                    = NewHolidayFunc(calculateNewYear)
	AUAustralianDay              = NewHoliday(time.January, 26)
	AUGoodFriday                 = GoodFriday
	AUChristmasDay               = NewHolidayFunc(calculateChristmasDay)
	AUBoxingDays                 = Christmas2
	AUObservedChristmasBoxingDay = NewHolidayFunc(calculateChristmasDayBoxingDay)
	AUEasterMonday               = EasterMonday
	AUAnzacDay                   = NewHolidayFunc(calculateAnzacDay)
)

// AddAustralianHolidays adds all Australian holidays
func AddAustralianHolidays(c *Calendar) {
	c.AddHoliday(
		AUNewYear,
		AUAustralianDay,
		AUGoodFriday,
		AUEasterMonday,
		AUAnzacDay,
		AUChristmasDay,
		AUBoxingDays,
		AUObservedChristmasBoxingDay,
	)
}

//  Holidays associated with the start of the modern Gregorian calendar.
//
//  New Year's Day is on January 1 and is the first day of a new year in the Gregorian calendar,
//  which is used in Australia and many other countries. Due to its geographical position close
//  to the International Date Line, Australia is one of the first countries in the world to welcome the New Year.
//  If it falls on a weekend an additional public holiday is held on the next available weekday.
//
//  https://www.timeanddate.com/holidays/australia/new-year-day
func calculateNewYear(year int, loc *time.Location) (time.Month, int) {
	d := time.Date(year, time.January, 1, 0, 0, 0, 0, loc)
	wd := d.Weekday()
	if wd == 0 {
		d = d.AddDate(0, 0, 1)
	}

	if wd == 6 {
		d = d.AddDate(0, 0, 2)
	}

	return d.Month(), d.Day()
}

// Anzac Day is a national day of remembrance in Australia and New Zealand that broadly commemorates all Australians
// and New Zealanders "who served and died in all wars, conflicts, and peacekeeping operations"
// Observed on 25 April each year. Unlike most other Australian public holidays, If it falls on a weekend it is NOT moved
// to the next available weekday, nor is there an additional public holiday held. However, if it clashes with Easter,
// an additional public holiday is held for Easter.
//
// https://en.wikipedia.org/wiki/Anzac_Day
// https://www.timeanddate.com/holidays/australia/anzac-day
func calculateAnzacDay(year int, loc *time.Location) (time.Month, int) {
	d := time.Date(year, time.April, 25, 0, 0, 0, 0, loc)
	easter := calculateEaster(year, loc)
	emMonth, emDay := calculateEasterMonday(year, loc)
	easterMonday := time.Date(year, emMonth, emDay, 0, 0, 0, 0, loc)

	if d.Equal(easter) || d.Equal(easterMonday) {
		d = easterMonday.AddDate(0, 0, 1)
	}

	return d.Month(), d.Day()
}

// Christmas Day
//
// Christmas day is a public holidays in Australia,
// if it fall on the weekend an additional public holiday is held on the next available weekday.
//
// https://www.timeanddate.com/holidays/australia/christmas-day-holiday
//
func calculateChristmasDay(year int, loc *time.Location) (time.Month, int) {
	d := time.Date(year, time.December, 25, 0, 0, 0, 0, loc)
	wd := d.Weekday()
	if wd == 0 || wd == 6 {
		d = d.AddDate(0, 0, 2)
	}

	return d.Month(), d.Day()
}

// Boxing Day
//
// Boxing day is a public holidays in Australia,
// if it fall on the weekend an additional public holiday is held on the next available weekday.
//
// https://www.timeanddate.com/holidays/australia/boxing-day
//
func calculateBoxingDay(year int, loc *time.Location) (time.Month, int) {
	d := time.Date(year, time.December, 26, 0, 0, 0, 0, loc)
	wd := d.Weekday()
	if wd == 0 || wd == 6 {
		d = d.AddDate(0, 0, 2)
	}

	return d.Month(), d.Day()
}

// Christmas Day
//
// Christmas day is a public holidays in Australia,
// if it fall on the weekend an additional public holiday is held on the next available weekday.
//
// https://www.timeanddate.com/holidays/australia/christmas-day-holiday
//
func calculateChristmasDayBoxingDay(year int, loc *time.Location) (time.Month, int) {
	christmas := time.Date(year, time.December, 25, 0, 0, 0, 0, loc)
	firstAvailableDay := time.Date(year, 12, 27, 0, 0, 0, 0, loc)

	// Sunday or Saturday
	if christmas.Weekday() == 0 || christmas.Weekday() == 6 {
		return firstAvailableDay.Month(), firstAvailableDay.Day()
	}

	return christmas.Month(), christmas.Day()
}
