package bg

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Standard BG weekend substitution rules:
	//   Saturdays move to Monday
	//   Sundays move to Monday
	weekendAlt = []cal.AltDay{
		{Day: time.Saturday, Offset: 2},
		{Day: time.Sunday, Offset: 1},
	}
	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "Нова година", Type: cal.ObservancePublic, Observed: weekendAlt})

	// LiberationDay represents Liberation Day on 3-Mar
	LiberationDay = &cal.Holiday{
		Name:     "Ден на Освобождението на България от османско иго - национален празник",
		Type:     cal.ObservancePublic,
		Month:    time.March,
		Day:      3,
		Func:     cal.CalcDayOfMonth,
		Observed: weekendAlt,
	}

	// GoodFriday represents Good Friday - two days before Easter
	OrthodoxGoodFriday = &cal.Holiday{
		Name:   "Велики петък",
		Type:   cal.ObservancePublic,
		Offset: -2,
		Julian: true,
		Func:   cal.CalcEasterOffset,
	}

	// EasterMonday represents Easter Monday on the day after Easter
	OrthodoxEasterMonday = &cal.Holiday{
		Name:   "Великден",
		Type:   cal.ObservancePublic,
		Offset: 1,
		Julian: true,
		Func:   cal.CalcEasterOffset,
	}

	// LabourDay represents Labour Day on 1-May
	LabourDay = &cal.Holiday{
		Name:  "Ден на труда и на международната работническа солидарност",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   1,
		Func: func(h *cal.Holiday, year int) time.Time {
			// May 1st is quite close to OrthodoxEaster so we need to make an extra check
			// If LabourDay falls on a Saturday or Sunday and Easter is on the Friday before that,
			// then LabourDay is observed on the Tuesday after Easter.
			easter := cal.CalcEasterOffset(OrthodoxGoodFriday, year)
			labourDay := cal.CalcDayOfMonth(h, year)
			daysDiff := labourDay.Sub(easter).Hours() / 24
			if daysDiff <= 2 {
				holiday := *h
				holiday.Offset = 2
				holiday.Julian = true
				return cal.CalcEasterOffset(&holiday, year)
			}
			return labourDay
		},
		Observed: weekendAlt,
	}

	// StGeorgesDay represents St. George's Day on 6-May
	StGeorgesDay = &cal.Holiday{
		Name:     "Гергьовден, Ден на храбростта и Българската армия",
		Type:     cal.ObservancePublic,
		Month:    time.May,
		Day:      6,
		Func:     cal.CalcDayOfMonth,
		Observed: weekendAlt,
	}

	// StCyrilAndMethodiusDay represents St. Cyril and St. Methodius Day on 24-May
	StCyrilAndMethodiusDay = &cal.Holiday{
		Name:     "Ден на светите братя Кирил и Методий, на българската азбука, просвета и култура и на славянската книжовност",
		Type:     cal.ObservancePublic,
		Month:    time.May,
		Day:      24,
		Func:     cal.CalcDayOfMonth,
		Observed: weekendAlt,
	}

	// UnificationDay represents Unification Day on 6-Sep
	UnificationDay = &cal.Holiday{
		Name:     "Ден на Съединението",
		Type:     cal.ObservancePublic,
		Month:    time.September,
		Day:      6,
		Func:     cal.CalcDayOfMonth,
		Observed: weekendAlt,
	}

	// IndependenceDay represents Independence Day on 22-Sep
	IndependenceDay = &cal.Holiday{
		Name:     "Ден на Независимостта на България",
		Type:     cal.ObservancePublic,
		Month:    time.September,
		Day:      22,
		Func:     cal.CalcDayOfMonth,
		Observed: weekendAlt,
	}

	// ChristmasEve represents Christmas Eve 24-Dec
	ChristmasEve = &cal.Holiday{
		Name:     "Бъдни вечер",
		Type:     cal.ObservancePublic,
		Month:    time.December,
		Day:      24,
		Func:     cal.CalcDayOfMonth,
		Observed: weekendAlt,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Коледа", Type: cal.ObservancePublic, Observed: []cal.AltDay{
		{Day: time.Saturday, Offset: 2},
		{Day: time.Sunday, Offset: 2},
	}})

	// ChristmasDay2 represents Christmas Day 2 on 26-Dec
	ChristmasDay2 = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Коледа 2", Type: cal.ObservancePublic, Observed: []cal.AltDay{
		{Day: time.Saturday, Offset: 2},
		{Day: time.Sunday, Offset: 2},
		{Day: time.Monday, Offset: 2},
		{Day: time.Tuesday, Offset: 1}}})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		LiberationDay,
		OrthodoxGoodFriday,
		OrthodoxEasterMonday,
		LabourDay,
		StGeorgesDay,
		StCyrilAndMethodiusDay,
		UnificationDay,
		IndependenceDay,
		ChristmasEve,
		ChristmasDay,
		ChristmasDay2,
	}
)
