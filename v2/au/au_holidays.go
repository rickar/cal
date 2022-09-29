// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package au provides holiday definitions for Australia.
package au

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Standard Australian weekend substitution rules:
	//   Saturdays move to Monday
	//   Sundays move to Monday
	weekendAlt = []cal.AltDay{
		{Day: time.Saturday, Offset: 2},
		{Day: time.Sunday, Offset: 1},
	}

	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "New Year's Day", Type: cal.ObservancePublic, Observed: weekendAlt})

	// AustraliaDay represents Australia Day on 26-Jan
	AustraliaDay = &cal.Holiday{
		Name:     "Australia Day",
		Type:     cal.ObservancePublic,
		Month:    time.January,
		Day:      26,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = aa.GoodFriday.Clone(&cal.Holiday{Name: "Good Friday", Type: cal.ObservancePublic})

	// EasterSaturday represents the day before Easter, which falls on a Saturday.
	EasterSaturday = &cal.Holiday{
		Name:   "Easter Saturday",
		Offset: -1,
		Func:   cal.CalcEasterOffset,
	}

	// EasterSunday represents Easter, which falls on a Sunday.
	EasterSunday = &cal.Holiday{
		Name:   "Easter Sunday",
		Offset: 0,
		Func:   cal.CalcEasterOffset,
	}

	// EasterMonday represents Easter Monday - the day after Easter
	EasterMonday = aa.EasterMonday.Clone(&cal.Holiday{Name: "Easter Monday", Type: cal.ObservancePublic})

	// LabourDayWa represents Labour Day in WA on the first Monday of March
	LabourDayWa = &cal.Holiday{
		Name:    "Labour Day",
		Type:    cal.ObservancePublic,
		Month:   time.March,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// LabourDayVic represents Labour Day in VIC on the second Monday of March
	LabourDayVic = &cal.Holiday{
		Name:    "Labour Day",
		Type:    cal.ObservancePublic,
		Month:   time.March,
		Weekday: time.Monday,
		Offset:  2,
		Func:    cal.CalcWeekdayOffset,
	}

	// LabourDayTas represents Eight Hours Day in TAS on the second Monday of March
	LabourDayTas = &cal.Holiday{
		Name:    "Eight Hours Day",
		Type:    cal.ObservancePublic,
		Month:   time.March,
		Weekday: time.Monday,
		Offset:  2,
		Func:    cal.CalcWeekdayOffset,
	}

	// CanberraDay represents Canberra Day in ACT on the second Monday of March
	CanberraDay = &cal.Holiday{
		Name:    "Canberra Day",
		Type:    cal.ObservancePublic,
		Month:   time.March,
		Weekday: time.Monday,
		Offset:  2,
		Func:    cal.CalcWeekdayOffset,
	}

	// MarchPublicHoliday represents March Public Holiday in SA on the second Monday of March
	MarchPublicHoliday = &cal.Holiday{
		Name:    "March Public Holiday",
		Type:    cal.ObservancePublic,
		Month:   time.March,
		Weekday: time.Monday,
		Offset:  2,
		Func:    cal.CalcWeekdayOffset,
	}

	// AnzacDay represents ANZAC Day on 25-Apr
	AnzacDay = &cal.Holiday{
		Name:  "ANZAC Day",
		Type:  cal.ObservancePublic,
		Month: time.April,
		Day:   25,
		Func:  cal.CalcDayOfMonth,
	}

	// AnzacDayActWa represents ANZAC Day for ACT and WA who observe a public holiday if it falls on a weekend
	AnzacDayActWa = AnzacDay.Clone(&cal.Holiday{Observed: weekendAlt})

	// AnzacDayNtQldSa represents ANZAC Day for NT, QLD and SA who observe a public holiday on Monday if ANZAC day falls on a Sunday
	AnzacDayNtQldSa = AnzacDay.Clone(&cal.Holiday{Observed: []cal.AltDay{{Day: time.Sunday, Offset: 1}}})

	// LabourDayNtQld represents May Day in NT and QLD on the first Monday of May
	LabourDayNtQld = &cal.Holiday{
		Name:    "Labour Day / May Day",
		Type:    cal.ObservancePublic,
		Month:   time.May,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// ReconciliationDay represents Reconciliation Day in ACT on the first Monday after or on 27-May
	ReconciliationDay = &cal.Holiday{
		Name:      "Reconciliation Day",
		Type:      cal.ObservancePublic,
		Month:     time.May,
		Day:       27,
		Weekday:   time.Monday,
		Offset:    1,
		Func:      cal.CalcWeekdayFrom,
		StartYear: 2018,
	}

	// WesternAustraliaDay represents Western Australia Day on the first Monday in June
	WesternAustraliaDay = &cal.Holiday{
		Name:    "Western Australia Day",
		Type:    cal.ObservancePublic,
		Month:   time.June,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// QueensBirthday represents Queen's Birthday on the second Monday in June
	QueensBirthday = &cal.Holiday{
		Name:    "Queen's Birthday",
		Type:    cal.ObservancePublic,
		Month:   time.June,
		Weekday: time.Monday,
		Offset:  2,
		Func:    cal.CalcWeekdayOffset,
	}

	// PicnicDay represents Picnic Day in NT on the first Monday in August
	PicnicDay = &cal.Holiday{
		Name:    "Picnic Day",
		Type:    cal.ObservancePublic,
		Month:   time.August,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// QueensBirthdayWa represents Queen's Birthday in WA on the last Monday in September
	QueensBirthdayWa = &cal.Holiday{
		Name:    "Queen's Birthday",
		Type:    cal.ObservancePublic,
		Month:   time.September,
		Weekday: time.Monday,
		Offset:  -1,
		Func:    cal.CalcWeekdayOffset,
	}

	// FridayBeforeAflFinal represents the Friday before the AFL Grand Final;
	// normally on the Friday before the last Saturday of September but subject to AFL schedules
	FridayBeforeAflFinal = &cal.Holiday{
		Name:      "Friday before the AFL Grand Final",
		Type:      cal.ObservancePublic,
		Func:      calcFridayBeforeAflFinal,
		StartYear: 2015,
	}

	// QueensBirthdayQld represents Queen's Birthday in QLD on the first Monday in October
	QueensBirthdayQld = &cal.Holiday{
		Name:    "Queen's Birthday",
		Type:    cal.ObservancePublic,
		Month:   time.October,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// LabourDayActNswSa represents Labour Day in ACT, NSW, and SA on the first Monday in October
	LabourDayActNswSa = &cal.Holiday{
		Name:    "Labour Day",
		Type:    cal.ObservancePublic,
		Month:   time.October,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// MelbourneCup represents Melbourne Cup day on the first Tuesday in November
	MelbourneCup = &cal.Holiday{
		Name:    "Melbourne Cup",
		Type:    cal.ObservancePublic,
		Month:   time.November,
		Weekday: time.Tuesday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Christmas Day", Type: cal.ObservanceBank, Observed: weekendAlt})

	// BoxingDay represents Boxing Day on 26-Dec
	BoxingDay = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Boxing Day", Type: cal.ObservanceBank,
		Observed: []cal.AltDay{
			{Day: time.Saturday, Offset: 2},
			{Day: time.Sunday, Offset: 2},
			{Day: time.Monday, Offset: 1}}})

	// ProclamationDay represents Proclamation Day on 26-Dec
	ProclamationDay = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Proclamation Day", Type: cal.ObservanceBank,
		Observed: []cal.AltDay{
			{Day: time.Saturday, Offset: 2},
			{Day: time.Sunday, Offset: 2},
			{Day: time.Monday, Offset: 1}}})

	// MourningDay2022 represents the National Day of Mourning for Her Majesty the Queen.
	MourningDay2022 = &cal.Holiday{
		Name:      "National Day of Mourning for Her Majesty the Queen",
		Type:      cal.ObservancePublic,
		Month:     time.September,
		Day:       22,
		StartYear: 2022,
		EndYear:   2022,
		Func:      cal.CalcDayOfMonth,
	}

	// HolidaysACT provides a list of standard holidays in the Australian Capital Territory region.
	HolidaysACT = []*cal.Holiday{
		NewYear,
		AustraliaDay,
		CanberraDay,
		GoodFriday,
		EasterSaturday,
		EasterSunday,
		EasterMonday,
		AnzacDayActWa,
		ReconciliationDay,
		QueensBirthday,
		MourningDay2022,
		LabourDayActNswSa,
		ChristmasDay,
		BoxingDay,
	}

	// HolidaysNSW provides a list of standard holidays in the New South Wales region.
	HolidaysNSW = []*cal.Holiday{
		NewYear,
		AustraliaDay,
		GoodFriday,
		EasterSaturday,
		EasterSunday,
		EasterMonday,
		AnzacDay,
		QueensBirthday,
		MourningDay2022,
		LabourDayActNswSa,
		ChristmasDay,
		BoxingDay,
	}

	// HolidaysNT provides a list of standard holidays in the Northern Territory region.
	HolidaysNT = []*cal.Holiday{
		NewYear,
		AustraliaDay,
		GoodFriday,
		EasterSaturday,
		EasterMonday,
		AnzacDayNtQldSa,
		LabourDayNtQld,
		QueensBirthday,
		PicnicDay,
		MourningDay2022,
		ChristmasDay,
		BoxingDay,
	}

	// HolidaysQLD provides a list of standard holidays in the Queensland region.
	HolidaysQLD = []*cal.Holiday{
		NewYear,
		AustraliaDay,
		GoodFriday,
		EasterSaturday,
		EasterSunday,
		EasterMonday,
		AnzacDayNtQldSa,
		LabourDayNtQld,
		MourningDay2022,
		QueensBirthdayQld,
		ChristmasDay,
		BoxingDay,
	}

	// HolidaysSA provides a list of standard holidays in the South Australia region.
	HolidaysSA = []*cal.Holiday{
		NewYear,
		AustraliaDay,
		MarchPublicHoliday,
		GoodFriday,
		EasterSaturday,
		EasterMonday,
		AnzacDayNtQldSa,
		QueensBirthday,
		MourningDay2022,
		LabourDayActNswSa,
		ChristmasDay,
		ProclamationDay,
	}

	// HolidaysTAS provides a list of standard holidays in the Tasmania region.
	HolidaysTAS = []*cal.Holiday{
		NewYear,
		AustraliaDay,
		LabourDayTas,
		GoodFriday,
		EasterMonday,
		AnzacDay,
		QueensBirthday,
		MourningDay2022,
		ChristmasDay,
		BoxingDay,
	}

	// HolidaysVIC provides a list of standard holidays in the Victoria region.
	HolidaysVIC = []*cal.Holiday{
		NewYear,
		AustraliaDay,
		LabourDayVic,
		GoodFriday,
		EasterSaturday,
		EasterSunday,
		EasterMonday,
		AnzacDay,
		QueensBirthday,
		MourningDay2022,
		FridayBeforeAflFinal,
		MelbourneCup,
		ChristmasDay,
		BoxingDay,
	}

	// HolidaysWA provides a list of standard holidays in the Western Australia region.
	HolidaysWA = []*cal.Holiday{
		NewYear,
		AustraliaDay,
		LabourDayWa,
		GoodFriday,
		EasterMonday,
		AnzacDayActWa,
		WesternAustraliaDay,
		QueensBirthdayWa,
		MourningDay2022,
		ChristmasDay,
		BoxingDay,
	}
)

func calcFridayBeforeAflFinal(_ *cal.Holiday, year int) time.Time {
	switch year {
	case 2015:
		return time.Date(year, time.October, 2, 0, 0, 0, 0, cal.DefaultLoc)
	case 2016:
		return time.Date(year, time.September, 30, 0, 0, 0, 0, cal.DefaultLoc)
	case 2020:
		return time.Date(year, time.October, 23, 0, 0, 0, 0, cal.DefaultLoc)
	default:
		aflFinalDay := cal.DayStart(cal.WeekdayN(year, time.September, time.Saturday, -1))
		return aflFinalDay.AddDate(0, 0, -1)
	}
}
