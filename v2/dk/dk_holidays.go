// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package dk provides holiday definitions for Denmark.
package dk

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Nytaarsdag represents New Year's Day on 1-Jan
	Nytaarsdag = aa.NewYear.Clone("Nytårsdag", cal.ObservancePublic, nil)

	// Skaertorsdag represents Maundy Thursday on the Thursday before Easter
	Skaertorsdag = aa.MaundyThursday.Clone("Skærtorsdag", cal.ObservancePublic, nil)

	// Langfredag represents Good Friday on the Friday before Easter
	Langfredag = aa.GoodFriday.Clone("Langfredag", cal.ObservancePublic, nil)

	// AndenPaaskedag represents Easter Monday on the day after Easter
	AndenPaaskedag = aa.EasterMonday.Clone("Anden påskedag", cal.ObservancePublic, nil)

	// StoreBededag represents General Prayer Day on the fourth Friday after Easter
	StoreBededag = &cal.Holiday{
		Name:   "Store bededag",
		Type:   cal.ObservancePublic,
		Offset: 26,
		Func:   cal.CalcEasterOffset,
	}

	// KristiHimmelfartsdag represents Ascension Day on the 39th day after Easter
	KristiHimmelfartsdag = aa.AscensionDay.Clone("Kristi Himmelfartsdag", cal.ObservancePublic, nil)

	// AndenPinsedag represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	AndenPinsedag = aa.PentecostMonday.Clone("Anden Pinsedag", cal.ObservancePublic, nil)

	// Grundlovsdag represents Constitution Day on 5-Jun
	Grundlovsdag = &cal.Holiday{
		Name:  "Grundlovsdag",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   5,
		Func:  cal.CalcDayOfMonth,
	}

	// Juledag represents Christmas Day on 25-Dec
	Juledag = aa.ChristmasDay.Clone("Juledag", cal.ObservancePublic, nil)

	// AndenJuledag represents the second day of Christmas on 26-Dec
	AndenJuledag = aa.ChristmasDay2.Clone("Anden juledag", cal.ObservancePublic, nil)

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Nytaarsdag,
		Skaertorsdag,
		Langfredag,
		AndenPaaskedag,
		StoreBededag,
		KristiHimmelfartsdag,
		AndenPinsedag,
		Grundlovsdag,
		Juledag,
		AndenJuledag,
	}
)
