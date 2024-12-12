// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package ke provides holiday definitions for Kenya.
package ke

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Standard KE weekend substitution rules:
	//   Sundays move to Monday
	weekendAlt = []cal.AltDay{
		{Day: time.Sunday, Offset: 1},
	}

	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{
		Name:     "New Year's Day",
		Type:     cal.ObservancePublic,
		Observed: weekendAlt,
	})

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = aa.GoodFriday.Clone(&cal.Holiday{Name: "Good Friday", Type: cal.ObservancePublic})

	// EasterMonday represents Easter Monday - the day after Easter
	EasterMonday = aa.EasterMonday.Clone(&cal.Holiday{Name: "Easter Monday", Type: cal.ObservancePublic})

	// LabourDay represents Labour Day on 1-May
	LabourDay = aa.WorkersDay.Clone(&cal.Holiday{
		Name:     "Labour Day",
		Type:     cal.ObservancePublic,
		Observed: weekendAlt,
	})

	// MadarakaDay represents Madaraka/Self-Governance Day on 1-Jun
	MadarakaDay = &cal.Holiday{
		Name:     "Madaraka Day",
		Type:     cal.ObservancePublic,
		Month:    time.June,
		Day:      1,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// UtamaduniDay represents Utamaduni Day on 10-Oct
	UtamaduniDay = &cal.Holiday{
		Name:      "Utamaduni Day",
		Type:      cal.ObservancePublic,
		Month:     time.October,
		Day:       10,
		Observed:  weekendAlt,
		Func:      cal.CalcDayOfMonth,
		StartYear: 2022,
		EndYear:   2023,
	}

	// MazingiraDay represents Environment Conservation Day on 10-Oct
	MazingiraDay = &cal.Holiday{
		Name:      "Mazingira Day",
		Type:      cal.ObservancePublic,
		Month:     time.October,
		Day:       10,
		Observed:  weekendAlt,
		Func:      cal.CalcDayOfMonth,
		StartYear: 2024,
	}

	// MashujaaDay represents Mashujaa/Heroes' Day on 20-Oct
	MashujaaDay = &cal.Holiday{
		Name:      "Mashujaa Day",
		Type:      cal.ObservancePublic,
		Month:     time.October,
		Day:       20,
		Observed:  weekendAlt,
		Func:      cal.CalcDayOfMonth,
		StartYear: 2010,
	}

	// JamhuriDay represents Jamhuri/Independence Day on 12-Dec
	JamhuriDay = &cal.Holiday{
		Name:     "Jamhuri Day",
		Type:     cal.ObservancePublic,
		Month:    time.December,
		Day:      12,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{
		Name:     "Christmas Day",
		Type:     cal.ObservancePublic,
		Observed: weekendAlt,
	})

	// BoxingDay represents Boxing Day on 26-Dec
	BoxingDay = aa.ChristmasDay2.Clone(&cal.Holiday{
		Name:     "Boxing Day",
		Type:     cal.ObservancePublic,
		Observed: weekendAlt,
	})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		GoodFriday,
		EasterMonday,
		LabourDay,
		MadarakaDay,
		MazingiraDay,
		MashujaaDay,
		JamhuriDay,
		ChristmasDay,
		BoxingDay,
	}
)
