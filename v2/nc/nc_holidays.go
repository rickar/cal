package nc

import (
	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/fr"
	"time"
)

var (
	// NouvelAn represents New Year's Day on 1-Jan
	NouvelAn = fr.NouvelAn.Clone(nil)

	// LundiDePâques represents Easter Monday on the day after Easter
	LundiDePâques = fr.LundiDePâques.Clone(nil)

	// FêteDuTravail represents Labour Day on 1-May
	FêteDuTravail = fr.FêteDuTravail.Clone(nil)

	// FêteDeLaVictoire represents Victory in Europe Day on 8-May
	FêteDeLaVictoire = fr.FêteDeLaVictoire.Clone(nil)

	// Ascension represents Ascension Day on the 39th day after Easter
	Ascension = fr.Ascension.Clone(nil)

	// LundiDePentecôte represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	LundiDePentecôte = fr.LundiDePentecôte.Clone(nil)

	// FêteNationale represents Bastille Day on 14-Jul
	FêteNationale = fr.FêteNationale.Clone(nil)

	// Assomption represents Assumption of Mary on 15-Aug
	Assomption = fr.Assomption.Clone(nil)

	// Toussaint represents All Saints' Day on 1-Nov
	Toussaint = fr.Toussaint.Clone(nil)

	// Armistice1918 represents Armistice Day on 11-Nov
	Armistice1918 = fr.Armistice1918.Clone(nil)

	// Noël represents Christmas Day on 25-Dec
	Noël = fr.Noël.Clone(nil)

	// FêteDeLaCitoyenneté represents the day that New Caledonia became French, the 24-Sept
	FêteDeLaCitoyenneté = &cal.Holiday{
		Name:  "Fête de la citoyenneté",
		Month: time.September,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NouvelAn,
		LundiDePâques,
		FêteDuTravail,
		FêteDeLaVictoire,
		Ascension,
		LundiDePentecôte,
		FêteNationale,
		Assomption,
		Toussaint,
		Armistice1918,
		Noël,
		FêteDeLaCitoyenneté,
	}
)
