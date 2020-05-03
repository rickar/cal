// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package be provides holiday definitions for Belgium.
package be

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Nieuwjaar represents New Year's Day on 1-Jan
	Nieuwjaar = aa.NewYear.Clone("Nieuwjaarsdag", cal.ObservancePublic, nil)

	// Paasmaandag represents Easter Monday on the day after Easter
	Paasmaandag = aa.EasterMonday.Clone("Paasmaandag", cal.ObservancePublic, nil)

	// DagVanDeArbeid represents Labor Day on the first Monday in May
	DagVanDeArbeid = aa.WorkersDay.Clone("Dag van de Arbeid", cal.ObservancePublic, nil)

	// OnzeLieveHeerHemelvaart represents Ascension Day on the 39th day after Easter
	OnzeLieveHeerHemelvaart = aa.AscensionDay.Clone("Onze Lieve Heer Hemelvaart", cal.ObservancePublic, nil)

	// Pinkstermaandag represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	Pinkstermaandag = aa.PentecostMonday.Clone("Pinkstermaandag", cal.ObservancePublic, nil)

	// NationaleFeestdag represents Belgian National Day on 21-Jul
	NationaleFeestdag = &cal.Holiday{
		Name:  "Nationale Feestdag",
		Type:  cal.ObservancePublic,
		Month: time.July,
		Day:   21,
		Func:  cal.CalcDayOfMonth,
	}

	// OnzeLieveVrouwHemelvaart represents Assumption of Mary on 15-Aug
	OnzeLieveVrouwHemelvaart = aa.AssumptionOfMary.Clone("Onze Lieve Vrouw Hemelvaart", cal.ObservancePublic, nil)

	// Allerheiligen represents All Saints' Day on 1-Nov
	Allerheiligen = aa.AllSaintsDay.Clone("Allerheiligen", cal.ObservancePublic, nil)

	// Wapenstilstand represents Armistice Day on 11-Nov
	Wapenstilstand = aa.ArmisticeDay.Clone("Wapenstilstand", cal.ObservancePublic, nil)

	// Kerstmis represents Christmas Day on 25-Dec
	Kerstmis = aa.ChristmasDay.Clone("Kerstmis", cal.ObservancePublic, nil)

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
