// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package ecb provides holiday definitions for the European Central Bank.
package ecb

import (
	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone("New Year's Day", cal.ObservanceBank, nil)

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = aa.GoodFriday.Clone("Good Friday", cal.ObservanceBank, nil)

	// EasterMonday represents Easter Monday - the day after Easter
	EasterMonday = aa.EasterMonday.Clone("Easter Monday", cal.ObservanceBank, nil)

	// LabourDay represents Labour Day on 1-May
	LabourDay = aa.WorkersDay.Clone("Labour Day", cal.ObservanceBank, nil)

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone("Christmas Day", cal.ObservanceBank, nil)

	// ChristmasHoliday represents the day after Christmas on 26-Dec
	ChristmasHoliday = aa.ChristmasDay2.Clone("Christmas Holiday", cal.ObservanceBank, nil)

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
