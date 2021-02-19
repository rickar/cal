// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package no provides holiday definitions for Norway.
package no

import (
	"time"

	"github.com/Tamh/cal/v2"
	"github.com/Tamh/cal/v2/aa"
)

var (
	// FoersteNyttaarsdag represents New Year's Day on 1-Jan
	FoersteNyttaarsdag = aa.NewYear.Clone(&cal.Holiday{Name: "Første nyttårsdag", Type: cal.ObservancePublic})

	// Skjaertorsdag represents Maundy Thursday on the Thursday before Easter
	Skjaertorsdag = aa.MaundyThursday.Clone(&cal.Holiday{Name: "Skjærtorsdag", Type: cal.ObservancePublic})

	// Langfredag represents Good Friday on the Friday before Easter
	Langfredag = aa.GoodFriday.Clone(&cal.Holiday{Name: "Langfredag", Type: cal.ObservancePublic})

	// AndrePaaskedag represents Easter Monday on the day after Easter
	AndrePaaskedag = aa.EasterMonday.Clone(&cal.Holiday{Name: "Andre påskedag", Type: cal.ObservancePublic})

	// Arbeiderenesdag represents Labour Day on 1-May
	Arbeiderenesdag = aa.WorkersDay.Clone(&cal.Holiday{Name: "Arbeidernes dag", Type: cal.ObservancePublic})

	// Grunnlovsdag represents Constitution Day on 17-May
	Grunnlovsdag = &cal.Holiday{
		Name:  "Grunnlovsdag",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   17,
		Func:  cal.CalcDayOfMonth,
	}

	// Kristihimmelfartsdag represents Ascension Day on the 39th day after Easter
	Kristihimmelfartsdag = aa.AscensionDay.Clone(&cal.Holiday{Name: "Kristi Himmelfartsdag", Type: cal.ObservancePublic})

	// AndrePinsedag represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	AndrePinsedag = aa.PentecostMonday.Clone(&cal.Holiday{Name: "Andre pinsedag", Type: cal.ObservancePublic})

	// FoersteJuledag represents Christmas Day on 25-Dec
	FoersteJuledag = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Første juledag", Type: cal.ObservancePublic})

	// AndreJuledag represents the second day of Christmas on 26-Dec
	AndreJuledag = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Andre juledag", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		FoersteNyttaarsdag,
		Skjaertorsdag,
		Langfredag,
		AndrePaaskedag,
		Arbeiderenesdag,
		Grunnlovsdag,
		Kristihimmelfartsdag,
		AndrePinsedag,
		FoersteJuledag,
		AndreJuledag,
	}
)
