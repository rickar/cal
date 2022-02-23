// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package jp provides holiday definitions for Japan.
package jp

import (
	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
	"math"
	"time"
)

var (
	// Standard Japan weekend substitution rules: Sundays move to Monday
	weekendAlt = []cal.AltDay{
		{Day: time.Sunday, Offset: 1},
	}

	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Type: cal.ObservancePublic, Observed: weekendAlt})

	// ComingOfAgeDay represents Coming of Age Day on the 2nd Monday in January
	ComingOfAgeDay = &cal.Holiday{
		Name:    "Coming of Age Day",
		Type:    cal.ObservancePublic,
		Month:   time.January,
		Weekday: time.Monday,
		Offset:  2,
		Func:    cal.CalcWeekdayOffset,
	}

	// NationalFoundationDay represents National Foundation Day on 11-February
	NationalFoundationDay = &cal.Holiday{
		Name:     "National Foundation Day",
		Type:     cal.ObservancePublic,
		Month:    time.February,
		Day:      11,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// TheEmperorsBirthday represents The Emperor's Birthday on 23-February
	TheEmperorsBirthday = &cal.Holiday{
		Name:     "The Emperor's Birthday",
		Type:     cal.ObservancePublic,
		Month:    time.February,
		Day:      23,
		Observed: weekendAlt,
		Func: func(h *cal.Holiday, year int) time.Time {
			if year <= 2019 {
				// Emperor Akihito abdicated in 2019.
				holiday := *h
				holiday.Month = time.December
				holiday.Day = 23

				return cal.CalcDayOfMonth(&holiday, year)
			}

			return cal.CalcDayOfMonth(h, year)
		},
	}

	// VernalEquinoxDay represents Vernal Equinox Day on Around 20-March
	VernalEquinoxDay = &cal.Holiday{
		Name:     "Vernal Equinox Day",
		Type:     cal.ObservancePublic,
		Month:    time.March,
		Observed: weekendAlt,
		Func: func(h *cal.Holiday, year int) time.Time {
			holiday := h
			holiday.Day = calcVernalEquinoxDate(year)

			return cal.CalcDayOfMonth(holiday, year)
		},
	}

	// ShowaDay represents Showa Day on 29-April
	ShowaDay = &cal.Holiday{
		Name:     "Showa Day",
		Type:     cal.ObservancePublic,
		Month:    time.April,
		Day:      29,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// ConstitutionMemorialDay represents Constitution Memorial Day on 3-May
	ConstitutionMemorialDay = &cal.Holiday{
		Name:  "Constitution Memorial Day",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   3,
		Observed: []cal.AltDay{
			{Day: time.Sunday, Offset: 3},
		},
		Func: cal.CalcDayOfMonth,
	}

	// GreeneryDay represents Greenery Day on 4-May
	GreeneryDay = &cal.Holiday{
		Name:  "Greenery Day",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   4,
		Observed: []cal.AltDay{
			{Day: time.Sunday, Offset: 2},
		},
		Func: cal.CalcDayOfMonth,
	}

	// ChildrensDay represents Children's Day on 5-May
	ChildrensDay = &cal.Holiday{
		Name:     "Children's Day",
		Type:     cal.ObservancePublic,
		Month:    time.May,
		Day:      5,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// MarineDay represents Marine Day on the 3rd Monday in July
	MarineDay = &cal.Holiday{
		Name:     "Marine Day",
		Type:     cal.ObservancePublic,
		Month:    time.July,
		Weekday:  time.Monday,
		Offset:   3,
		Observed: weekendAlt,
		Func: func(h *cal.Holiday, year int) time.Time {
			if year == 2020 || year == 2021 {
				// As special arrangement for the 2020 Summer Olympics, the 2020 and 2021 date for Marine Day was moved
				holiday := *h
				holiday.Weekday = time.Thursday
				holiday.Offset = 4

				return cal.CalcWeekdayOffset(&holiday, year)
			}

			return cal.CalcWeekdayOffset(h, year)
		},
	}

	// MountainDay represents Mountain Day on 11-August
	MountainDay = &cal.Holiday{
		Name:      "Mountain Day",
		Type:      cal.ObservancePublic,
		Month:     time.August,
		Day:       11,
		Observed:  weekendAlt,
		Func:      cal.CalcDayOfMonth,
		StartYear: 2016,
	}

	// RespectForTheAgedDay represents Respect for the Aged Day on the 3rd Monday in September
	RespectForTheAgedDay = &cal.Holiday{
		Name:    "Respect for the Aged Day",
		Type:    cal.ObservancePublic,
		Month:   time.September,
		Weekday: time.Monday,
		Offset:  3,
		Func:    cal.CalcWeekdayOffset,
	}

	// AutumnalEquinoxDay represents Autumnal Equinox Day on Around 23-September
	AutumnalEquinoxDay = &cal.Holiday{
		Name:     "Autumnal Equinox Day",
		Type:     cal.ObservancePublic,
		Month:    time.September,
		Observed: weekendAlt,
		Func: func(h *cal.Holiday, year int) time.Time {
			holiday := *h
			holiday.Day = calcAutumnalEquinoxDate(year)

			return cal.CalcDayOfMonth(&holiday, year)
		},
	}

	// SportsDay represents Sports Day on the 2nd Monday in October
	SportsDay = &cal.Holiday{
		Name:    "Respect for the Aged Day",
		Type:    cal.ObservancePublic,
		Month:   time.October,
		Weekday: time.Monday,
		Offset:  2,
		Func: func(h *cal.Holiday, year int) time.Time {
			if year == 2020 || year == 2021 {
				// As special arrangement for the 2020 Summer Olympics, the 2020 and 2021 date for Sports Day was moved
				holiday := *h
				holiday.Month = time.July
				holiday.Weekday = time.Friday
				holiday.Offset = 4

				return cal.CalcWeekdayOffset(&holiday, year)
			}

			return cal.CalcWeekdayOffset(h, year)
		},
	}

	// CultureDay represents Culture Day on 3-November
	CultureDay = &cal.Holiday{
		Name:     "Culture Day",
		Type:     cal.ObservancePublic,
		Month:    time.November,
		Day:      3,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// LaborThanksgivingDay represents Labor Thanksgiving Day on 23-November
	LaborThanksgivingDay = &cal.Holiday{
		Name:     "Labor Thanksgiving Day",
		Type:     cal.ObservancePublic,
		Month:    time.November,
		Day:      23,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// nationalHolidayInSeptember represents National holiday in September
	nationalHolidayInSeptember = &cal.Holiday{
		Name:        "National holiday in September",
		Description: "If the day before and the next day are both holiday, the day becomes national holiday.",
		Type:        cal.ObservancePublic,
		Observed:    weekendAlt,
		Func: func(h *cal.Holiday, year int) time.Time {
			switch year {
			// only dates in September 2015 - 2032 are supported
			case 2015:
				return time.Date(year, 9, 22, 0, 0, 0, 0, cal.DefaultLoc)
			case 2026:
				return time.Date(year, 9, 22, 0, 0, 0, 0, cal.DefaultLoc)
			case 2032:
				return time.Date(year, 9, 21, 0, 0, 0, 0, cal.DefaultLoc)
			default:
				return time.Time{}
			}
		},
	}

	// nationalHoliday2019 represents National holiday in 2019
	nationalHoliday2019 = &cal.Holiday{
		Name: "National holiday in 2019",
		Type: cal.ObservancePublic,
		Func: func(h *cal.Holiday, year int) time.Time {
			if year == 2019 {
				// year 2019 has many national holidays because new emperor enthroned
				return cal.CalcDayOfMonth(h, year)
			}

			return time.Time{}
		},
	}

	exceptionalNationalHolidays = []*cal.Holiday{
		nationalHolidayInSeptember,
		nationalHoliday2019.Clone(&cal.Holiday{Month: time.April, Day: 30}),
		nationalHoliday2019.Clone(&cal.Holiday{Month: time.May, Day: 1, Name: "New Emperor Enthronement Day"}),
		nationalHoliday2019.Clone(&cal.Holiday{Month: time.May, Day: 2}),
		nationalHoliday2019.Clone(&cal.Holiday{Month: time.October, Day: 22, Name: "The New Emperor Enthronement Ceremony"}),
	}

	Holidays = append(
		[]*cal.Holiday{
			NewYear,
			ComingOfAgeDay,
			NationalFoundationDay,
			TheEmperorsBirthday,
			VernalEquinoxDay,
			ShowaDay,
			ConstitutionMemorialDay,
			GreeneryDay,
			ChildrensDay,
			MarineDay,
			MountainDay,
			RespectForTheAgedDay,
			AutumnalEquinoxDay,
			SportsDay,
			CultureDay,
			LaborThanksgivingDay,
		},
		exceptionalNationalHolidays...,
	)
)

func calcVernalEquinoxDate(year int) int {
	val := calcEquinoxBase(year)

	switch {
	case 1851 <= year && year <= 1899:
		val += 19.8277
	case 1900 <= year && year <= 1979:
		val += 20.8357
	case 1980 <= year && year <= 2099:
		val += 20.8431
	case 2100 <= year && year <= 2150:
		val += 21.8510
	}

	return int(math.Floor(val))
}

func calcAutumnalEquinoxDate(year int) int {
	val := calcEquinoxBase(year)

	switch {
	case 1851 <= year && year <= 1899:
		val += 22.2588
	case 1900 <= year && year <= 1979:
		val += 23.2588
	case 1980 <= year && year <= 2099:
		val += 23.2488
	case 2100 <= year && year <= 2150:
		val += 24.2488
	}

	return int(math.Floor(val))
}

func calcEquinoxBase(year int) float64 {
	return 0.242194*float64(year-1980) - math.Floor(float64(year-1980)/4.0)
}
