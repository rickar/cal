// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package ie provides holiday definitions for the Republic Of Ireland.
package ie

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{})

	// SaintBrigidDay represents Saint Patrick's Day on 17-Mar
	SaintBrigidDay = &cal.Holiday{
		Name:      "Saint Brigid’s Day",
		Month:     time.February,
		Weekday:   time.Monday,
		Offset:    1,
		Func:      CalcIfFirstFallsOnFriday,
		StartYear: 2023,
	}

	// ExtraPublicHoliday2022 represents extra public holiday in 2022
	ExtraPublicHoliday2022 = &cal.Holiday{
		Name:      "Extra Public Holiday 2022",
		Month:     time.March,
		Day:       18,
		StartYear: 2022,
		EndYear:   2022,
		Func:      cal.CalcDayOfMonth,
	}

	// SaintPatrickDay represents Saint Patrick's Day on 17-Mar
	SaintPatrickDay = &cal.Holiday{
		Name:  "Saint Patrick's Day",
		Month: time.March,
		Day:   17,
		Func:  cal.CalcDayOfMonth,
	}

	// EasterMonday represents Easter Monday - the day after Easter
	EasterMonday = aa.EasterMonday.Clone(&cal.Holiday{})

	// FirstMondayMay represents the first Monday in May
	FirstMondayMay = &cal.Holiday{
		Name:    "First Monday in May",
		Month:   time.May,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// FirstMondayJune represents the first Monday in June
	FirstMondayJune = &cal.Holiday{
		Name:    "First Monday in June",
		Month:   time.June,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// FirstMondayAugust represents the first Monday in August
	FirstMondayAugust = &cal.Holiday{
		Name:    "First Monday in August",
		Month:   time.August,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// LastMondayInOctober represents the last Monday in October
	LastMondayInOctober = &cal.Holiday{
		Name:    "Last Monday in October",
		Type:    cal.ObservancePublic,
		Month:   time.October,
		Weekday: time.Monday,
		Offset:  -1,
		Func:    cal.CalcWeekdayOffset,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{})

	// SaintStephenDay represents Saint Stephen's Day on 26-Dec
	SaintStephenDay = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Saint Stephen's Day"})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		ExtraPublicHoliday2022,
		SaintBrigidDay,
		SaintPatrickDay,
		EasterMonday,
		FirstMondayMay,
		FirstMondayJune,
		FirstMondayAugust,
		LastMondayInOctober,
		ChristmasDay,
		SaintStephenDay,
	}
)

func CalcIfFirstFallsOnFriday(h *cal.Holiday, year int) time.Time {
	fistFriday := cal.DayStart(cal.WeekdayN(year, h.Month, time.Friday, 1))
	// if 1st falls on a Friday
	if fistFriday.Day() == 1 {
		return fistFriday
	}

	return cal.DayStart(cal.WeekdayN(year, h.Month, h.Weekday, h.Offset))
}
