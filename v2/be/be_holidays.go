// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package be provides holiday definitions for Belgium.
package be

import (
	"time"

	"github.com/Tamh/cal/v2"
	"github.com/Tamh/cal/v2/aa"
)

var (
	// Nieuwjaar represents New Year's Day on 1-Jan
	Nieuwjaar = aa.NewYear.Clone(&cal.Holiday{Name: "Nieuwjaarsdag", Type: cal.ObservancePublic})

	// Paasmaandag represents Easter Monday on the day after Easter
	Paasmaandag = aa.EasterMonday.Clone(&cal.Holiday{Name: "Paasmaandag", Type: cal.ObservancePublic})

	// DagVanDeArbeid represents Labor Day on the first Monday in May
	DagVanDeArbeid = aa.WorkersDay.Clone(&cal.Holiday{Name: "Dag van de Arbeid", Type: cal.ObservancePublic})

	// OnzeLieveHeerHemelvaart represents Ascension Day on the 39th day after Easter
	OnzeLieveHeerHemelvaart = aa.AscensionDay.Clone(&cal.Holiday{Name: "Onze Lieve Heer Hemelvaart", Type: cal.ObservancePublic})

	// Pinkstermaandag represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	Pinkstermaandag = aa.PentecostMonday.Clone(&cal.Holiday{Name: "Pinkstermaandag", Type: cal.ObservancePublic})

	// NationaleFeestdag represents Belgian National Day on 21-Jul
	NationaleFeestdag = &cal.Holiday{
		Name:  "Nationale Feestdag",
		Type:  cal.ObservancePublic,
		Month: time.July,
		Day:   21,
		Func:  cal.CalcDayOfMonth,
	}

	// OnzeLieveVrouwHemelvaart represents Assumption of Mary on 15-Aug
	OnzeLieveVrouwHemelvaart = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Onze Lieve Vrouw Hemelvaart", Type: cal.ObservancePublic})

	// Allerheiligen represents All Saints' Day on 1-Nov
	Allerheiligen = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Allerheiligen", Type: cal.ObservancePublic})

	// Wapenstilstand represents Armistice Day on 11-Nov
	Wapenstilstand = aa.ArmisticeDay.Clone(&cal.Holiday{Name: "Wapenstilstand", Type: cal.ObservancePublic})

	// Kerstmis represents Christmas Day on 25-Dec
	Kerstmis = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Kerstmis", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Nieuwjaar,
		Paasmaandag,
		DagVanDeArbeid,
		OnzeLieveHeerHemelvaart,
		Pinkstermaandag,
		NationaleFeestdag,
		OnzeLieveVrouwHemelvaart,
		Allerheiligen,
		Wapenstilstand,
		Kerstmis,
	}
)
