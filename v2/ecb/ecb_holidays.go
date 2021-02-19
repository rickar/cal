// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package ecb provides holiday definitions for the European Central Bank.
package ecb

import (
	"github.com/Tamh/cal/v2"
	"github.com/Tamh/cal/v2/aa"
)

var (
	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "New Year's Day", Type: cal.ObservanceBank})

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = aa.GoodFriday.Clone(&cal.Holiday{Name: "Good Friday", Type: cal.ObservanceBank})

	// EasterMonday represents Easter Monday - the day after Easter
	EasterMonday = aa.EasterMonday.Clone(&cal.Holiday{Name: "Easter Monday", Type: cal.ObservanceBank})

	// LabourDay represents Labour Day on 1-May
	LabourDay = aa.WorkersDay.Clone(&cal.Holiday{Name: "Labour Day", Type: cal.ObservanceBank})

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Christmas Day", Type: cal.ObservanceBank})

	// ChristmasHoliday represents the day after Christmas on 26-Dec
	ChristmasHoliday = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Christmas Holiday", Type: cal.ObservanceBank})

	// Holidays provides a list of the standard ECB holidays
	Holidays = []*cal.Holiday{
		NewYear,
		GoodFriday,
		EasterMonday,
		LabourDay,
		ChristmasDay,
		ChristmasHoliday,
	}
)
