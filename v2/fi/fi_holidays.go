// Package fi provides holiday definitions for Finland.
package fi

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Uudenvuodenpäivä represents New Year's Day on 1-Jan
	Uudenvuodenpaiva = aa.NewYear.Clone(&cal.Holiday{Name: "Uudenvuodenpäivä", Type: cal.ObservancePublic})

	// Loppiainen represents Epiphany on 6-Jan
	Loppiainen = aa.Epiphany.Clone(&cal.Holiday{Name: "Loppiainen", Type: cal.ObservancePublic})

	// Pitkäperjantai represents Good Friday on the Friday before Easter
	Pitkaperjantai = aa.GoodFriday.Clone(&cal.Holiday{Name: "Pitkäperjantai", Type: cal.ObservancePublic})

	// Pääsiäispäivä represents the day of Easter
	Paasiaispaiva = aa.Easter.Clone(&cal.Holiday{Name: "Pääsiäispäivä", Type: cal.ObservancePublic})

	// Toinen pääsiäispäivä represents Easter Monday on the day after Easter
	ToinenPaasiaispaiva = aa.EasterMonday.Clone(&cal.Holiday{Name: "Toinen pääsiäispäivä", Type: cal.ObservancePublic})

	// Vappu represents Labour Day on 1-May
	Vappu = aa.WorkersDay.Clone(&cal.Holiday{Name: "Vappu", Type: cal.ObservancePublic})

	// Helatorstai represents Ascension Day on the 39th day after Easter
	Helatorstai = aa.AscensionDay.Clone(&cal.Holiday{Name: "Helatorstai", Type: cal.ObservancePublic})

	// Helluntaipäivä represents Pentecost Sunday on the 49th day after Easter
	Helluntaipaiva = aa.Pentecost.Clone(&cal.Holiday{Name: "Helluntaipäivä", Type: cal.ObservancePublic})

	// Juhannusaatto represents Midsummer's Eve on the day before Midsummer's Day
	Juhannusaatto = &cal.Holiday{
		Name:    "Juhannusaatto",
		Type:    cal.ObservancePublic,
		Month:   time.June,
		Day:     19,
		Offset:  1,
		Weekday: time.Friday,
		Func:    cal.CalcWeekdayFrom,
	}

	// Juhannuspäivä represents Midsummer's Day on the first Saturday from 20-Jun
	Juhannuspaiva = &cal.Holiday{
		Name:    "Juhannuspäivä",
		Type:    cal.ObservancePublic,
		Month:   time.June,
		Day:     20,
		Offset:  1,
		Weekday: time.Saturday,
		Func:    cal.CalcWeekdayFrom,
	}

	// Pyhäinpäivä represents All Saints' Day on the first Saturday from 31-Oct
	Pyhainpaiva = &cal.Holiday{
		Name:    "Pyhäinpäivä",
		Type:    cal.ObservancePublic,
		Month:   time.October,
		Day:     31,
		Offset:  1,
		Weekday: time.Saturday,
		Func:    cal.CalcWeekdayFrom,
	}

	// Itsenäisyyspäivä represents National Day of Finland on 6-Dec
	Itsenaisyyspaiva = &cal.Holiday{
		Name:  "Itsenäisyyspäivä",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   6,
		Func:  cal.CalcDayOfMonth,
	}

	// Jouluaatto represents Christmas Eve on 24-Dec
	Jouluaatto = &cal.Holiday{
		Name:  "Jouluaatto",
		Type:  cal.ObservanceOther,
		Month: time.December,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// Joulupäivä represents Christmas Day on 25-Dec
	Joulupaiva = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Joulupäivä", Type: cal.ObservancePublic})

	// Tapaninpäivä represents the second day of Christmas on 26-Dec
	Tapaninpaiva = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Tapaninpäivä", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Uudenvuodenpaiva,
		Loppiainen,
		Pitkaperjantai,
		Paasiaispaiva,
		ToinenPaasiaispaiva,
		Vappu,
		Helatorstai,
		Helluntaipaiva,
		Juhannusaatto,
		Juhannuspaiva,
		Pyhainpaiva,
		Itsenaisyyspaiva,
		Jouluaatto,
		Joulupaiva,
		Tapaninpaiva,
	}
)
