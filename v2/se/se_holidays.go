// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package se provides holiday definitions for Sweden.
package se

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Nyarsdagen represents New Year's Day on 1-Jan
	Nyarsdagen = aa.NewYear.Clone("Nyårsdagen", cal.ObservancePublic, nil)

	// TrettondedagJul represents Epiphany on 6-Jan
	TrettondedagJul = aa.Epiphany.Clone("Trettondedag jul", cal.ObservancePublic, nil)

	// Langfredagen represents Good Friday on the Friday before Easter
	Langfredagen = aa.GoodFriday.Clone("Långfredagen", cal.ObservancePublic, nil)

	// AnnandagPask represents Easter Monday on the day after Easter
	AnnandagPask = aa.EasterMonday.Clone("Annandag påsk", cal.ObservancePublic, nil)

	// ForstaMaj represents Labour Day on 1-May
	ForstaMaj = aa.WorkersDay.Clone("Första Maj", cal.ObservancePublic, nil)

	// KristiHimmelfardsdag represents Ascension Day on the 39th day after Easter
	KristiHimmelfardsdag = aa.AscensionDay.Clone("Kristi himmelsfärds dag", cal.ObservancePublic, nil)

	// Nationaldagen represents National Day of Sweden on 6-Jun
	Nationaldagen = &cal.Holiday{
		Name:  "Sveriges nationaldag",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   6,
		Func:  cal.CalcDayOfMonth,
	}

	// Midsommarafton represents Midsummer's Eve on the day before Midsummer's Day
	Midsommarafton = &cal.Holiday{
		Name:    "Midsommarafton",
		Type:    cal.ObservanceOther,
		Month:   time.June,
		Day:     19,
		Offset:  1,
		Weekday: time.Friday,
		Func:    cal.CalcWeekdayFrom,
	}

	// Midsommardagen represents Midsummer's Day on the first Saturday from 20-Jun
	Midsommardagen = &cal.Holiday{
		Name:    "Midsommardagen",
		Type:    cal.ObservancePublic,
		Month:   time.June,
		Day:     20,
		Offset:  1,
		Weekday: time.Saturday,
		Func:    cal.CalcWeekdayFrom,
	}

	// AllaHelgonsDag represents All Saints' Day on the first Saturday from 31-Oct
	AllaHelgonsDag = &cal.Holiday{
		Name:    "Alla helgons dag",
		Type:    cal.ObservancePublic,
		Month:   time.October,
		Day:     31,
		Offset:  1,
		Weekday: time.Saturday,
		Func:    cal.CalcWeekdayFrom,
	}

	// Julafton represents Christmas Eve on 24-Dec
	Julafton = &cal.Holiday{
		Name:  "Julafton",
		Type:  cal.ObservanceOther,
		Month: time.December,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// Juldagen represents Christmas Day on 25-Dec
	Juldagen = aa.ChristmasDay.Clone("Juldagen", cal.ObservancePublic, nil)

	// AnnandagJul represents the second day of Christmas on 26-Dec
	AnnandagJul = aa.ChristmasDay2.Clone("Annandag jul", cal.ObservancePublic, nil)

	// Nyarsafton represents New Year's Eve on 31-Dec
	Nyarsafton = &cal.Holiday{
		Name:  "Nyårsafton",
		Type:  cal.ObservanceOther,
		Month: time.December,
		Day:   31,
		Func:  cal.CalcDayOfMonth,
	}

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Nyarsdagen,
		TrettondedagJul,
		Langfredagen,
		AnnandagPask,
		ForstaMaj,
		KristiHimmelfardsdag,
		Nationaldagen,
		Midsommarafton,
		Midsommardagen,
		AllaHelgonsDag,
		Julafton,
		Juldagen,
		AnnandagJul,
		Nyarsafton,
	}
)
