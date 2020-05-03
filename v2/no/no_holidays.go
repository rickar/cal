// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package no provides holiday definitions for Norway.
package no

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// FoersteNyttaarsdag represents New Year's Day on 1-Jan
	FoersteNyttaarsdag = aa.NewYear.Clone("Første nyttårsdag", cal.ObservancePublic, nil)

	// Skjaertorsdag represents Maundy Thursday on the Thursday before Easter
	Skjaertorsdag = aa.MaundyThursday.Clone("Skjærtorsdag", cal.ObservancePublic, nil)

	// Langfredag represents Good Friday on the Friday before Easter
	Langfredag = aa.GoodFriday.Clone("Langfredag", cal.ObservancePublic, nil)

	// AndrePaaskedag represents Easter Monday on the day after Easter
	AndrePaaskedag = aa.EasterMonday.Clone("Andre påskedag", cal.ObservancePublic, nil)

	// Arbeiderenesdag represents Labour Day on 1-May
	Arbeiderenesdag = aa.WorkersDay.Clone("Arbeidernes dag", cal.ObservancePublic, nil)

	// Grunnlovsdag represents Constitution Day on 17-May
	Grunnlovsdag = &cal.Holiday{
		Name:  "Grunnlovsdag",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   17,
		Func:  cal.CalcDayOfMonth,
	}

	// Kristihimmelfartsdag represents Ascension Day on the 39th day after Easter
	Kristihimmelfartsdag = aa.AscensionDay.Clone("Kristi Himmelfartsdag", cal.ObservancePublic, nil)

	// AndrePinsedag represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	AndrePinsedag = aa.PentecostMonday.Clone("Andre pinsedag", cal.ObservancePublic, nil)

	// FoersteJuledag represents Christmas Day on 25-Dec
	FoersteJuledag = aa.ChristmasDay.Clone("Første juledag", cal.ObservancePublic, nil)

	// AndreJuledag represents the second day of Christmas on 26-Dec
	AndreJuledag = aa.ChristmasDay2.Clone("Andre juledag", cal.ObservancePublic, nil)

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
