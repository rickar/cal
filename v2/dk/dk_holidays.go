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
	Nytaarsdag = aa.NewYear.Clone(&cal.Holiday{Name: "Nytårsdag", Type: cal.ObservancePublic})

	// Skaertorsdag represents Maundy Thursday on the Thursday before Easter
	Skaertorsdag = aa.MaundyThursday.Clone(&cal.Holiday{Name: "Skærtorsdag", Type: cal.ObservancePublic})

	// Langfredag represents Good Friday on the Friday before Easter
	Langfredag = aa.GoodFriday.Clone(&cal.Holiday{Name: "Langfredag", Type: cal.ObservancePublic})

	// AndenPaaskedag represents Easter Monday on the day after Easter
	AndenPaaskedag = aa.EasterMonday.Clone(&cal.Holiday{Name: "Anden påskedag", Type: cal.ObservancePublic})

	// StoreBededag represents General Prayer Day on the fourth Friday after Easter
	StoreBededag = &cal.Holiday{
		Name:   "Store bededag",
		Type:   cal.ObservancePublic,
		Offset: 26,
		Func:   cal.CalcEasterOffset,
	}

	// KristiHimmelfartsdag represents Ascension Day on the 39th day after Easter
	KristiHimmelfartsdag = aa.AscensionDay.Clone(&cal.Holiday{Name: "Kristi Himmelfartsdag", Type: cal.ObservancePublic})

	// AndenPinsedag represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	AndenPinsedag = aa.PentecostMonday.Clone(&cal.Holiday{Name: "Anden Pinsedag", Type: cal.ObservancePublic})

	// Grundlovsdag represents Constitution Day on 5-Jun
	Grundlovsdag = &cal.Holiday{
		Name:  "Grundlovsdag",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   5,
		Func:  cal.CalcDayOfMonth,
	}

	// Juledag represents Christmas Day on 25-Dec
	Juledag = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Juledag", Type: cal.ObservancePublic})

	// AndenJuledag represents the second day of Christmas on 26-Dec
	AndenJuledag = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Anden juledag", Type: cal.ObservancePublic})

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
