// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package aa provides common holiday definitions that are frequently used.
package aa

import (
	"time"

	"github.com/rickar/cal/v2"
)

var (
	// NewYear represents New Year's Day on 1-Jan
	NewYear = &cal.Holiday{
		Name:  "New Year's Day",
		Month: time.January,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
	}

	// Epiphany represents Epiphany on 6-Jan
	Epiphany = &cal.Holiday{
		Name:  "Epiphany",
		Month: time.January,
		Day:   6,
		Func:  cal.CalcDayOfMonth,
	}

	// MaundyThursday represents Maundy Thursday - three days before Easter
	MaundyThursday = &cal.Holiday{
		Name:   "Maundy Thursday",
		Offset: -3,
		Func:   cal.CalcEasterOffset,
	}

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = &cal.Holiday{
		Name:   "Good Friday",
		Offset: -2,
		Func:   cal.CalcEasterOffset,
	}

	// Easter represents the day of Easter (Sunday)
	Easter = &cal.Holiday{
		Name:   "Easter",
		Offset: 0,
		Func:   cal.CalcEasterOffset,
	}

	// EasterMonday represents Easter Monday - the day after Easter
	EasterMonday = &cal.Holiday{
		Name:   "Easter Monday",
		Offset: 1,
		Func:   cal.CalcEasterOffset,
	}

	// WorkersDay represents International Workers' Day on 1-May
	WorkersDay = &cal.Holiday{
		Name:  "International Workers' Day",
		Month: time.May,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
	}

	// AscensionDay represents Ascension Day on the 39th day after Easter
	AscensionDay = &cal.Holiday{
		Name:   "Ascension Day",
		Offset: 39,
		Func:   cal.CalcEasterOffset,
	}

	// Pentecost represents Pentecoast Sunday on the 49th day after Easter
	Pentecost = &cal.Holiday{
		Name:   "Pentecost",
		Offset: 49,
		Func:   cal.CalcEasterOffset,
	}

	// PentecostMonday represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	PentecostMonday = &cal.Holiday{
		Name:   "Pentecost Monday",
		Offset: 50,
		Func:   cal.CalcEasterOffset,
	}

	// CorpusChristi represents Corpus Christi on the 60th day after Easter
	CorpusChristi = &cal.Holiday{
		Name:   "Corpus Christi",
		Offset: 60,
		Func:   cal.CalcEasterOffset,
	}

	// AssumptionOfMary represents Assumption of Mary on 15-Aug
	AssumptionOfMary = &cal.Holiday{
		Name:  "Assumption of Mary",
		Month: time.August,
		Day:   15,
		Func:  cal.CalcDayOfMonth,
	}

	// AllSaintsDay represents All Saints' Day on 1-Nov
	AllSaintsDay = &cal.Holiday{
		Name:  "All Saints' Day",
		Month: time.November,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
	}

	// ArmisticeDay represents Armistice Day on 11-Nov
	ArmisticeDay = &cal.Holiday{
		Name:  "Armistice Day",
		Month: time.November,
		Day:   11,
		Func:  cal.CalcDayOfMonth,
	}

	// ImmaculateConception represents Immaculate Conception on 8-Dec
	ImmaculateConception = &cal.Holiday{
		Name:  "Immaculate Conception",
		Month: time.December,
		Day:   8,
		Func:  cal.CalcDayOfMonth,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = &cal.Holiday{
		Name:  "Christmas Day",
		Month: time.December,
		Day:   25,
		Func:  cal.CalcDayOfMonth,
	}

	// ChristmasDay2 represents the day after Christmas (Boxing Day / St. Stephen's Day) on 26-Dec
	ChristmasDay2 = &cal.Holiday{
		Name:  "2nd Day of Christmas",
		Month: time.December,
		Day:   26,
		Func:  cal.CalcDayOfMonth,
	}
)
