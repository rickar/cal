// Package th provides holiday definition for Thailand.
package th

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Standard weekend substitution rules:
	// 		Saturdays move to Monday
	//		Sundays move to Monday

	weekendAlt = []cal.AltDay{
		{Day: time.Saturday, Offset: 2},
		{Day: time.Sunday, Offset: 1},
	}

	// Long weekend substitution rules:
	// 		Saturdays move to Monday
	//		Sundays move to Tuesday

	longWeekendAlt = []cal.AltDay{
		{Day: time.Saturday, Offset: 2},
		{Day: time.Sunday, Offset: 2},
	}

	// NewYear represents New Year's Day on 1 Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{
		Name:     "วันขึ้นปีใหม่",
		Type:     cal.ObservancePublic,
		Observed: weekendAlt,
	})

	// ChakriDay represents Chakri Dynasty Memorial Day on 6 Apr
	ChakriDay = &cal.Holiday{
		Name:     "วันจักรี",
		Type:     cal.ObservancePublic,
		Month:    time.April,
		Day:      6,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// ElderlyDay represents Songkran and Elderly National Day on 13 Apr
	ElderlyDay = &cal.Holiday{
		Name:     "วันมหาสงกรานต์ / วันผู้สูงอายุ",
		Type:     cal.ObservancePublic,
		Month:    time.April,
		Day:      13,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// FamilyDay represents Songkran and Family National Day on 14 Apr
	FamilyDay = &cal.Holiday{
		Name:     "วันสงกรานต์ / วันครอบครัวไทย",
		Type:     cal.ObservancePublic,
		Month:    time.April,
		Day:      14,
		Observed: longWeekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// ThaiNewYear represents Songkran on 15 Apr
	ThaiNewYear = &cal.Holiday{
		Name:     "วันสงกรานต์",
		Type:     cal.ObservancePublic,
		Month:    time.April,
		Day:      15,
		Observed: longWeekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// Songkran is held from 13 Apr to 15 Apr
	SongKranDays = []*cal.Holiday{
		ElderlyDay,
		FamilyDay,
		ThaiNewYear,
	}

	// LaborDay represents international worker's day on 1 May
	LaborDay = aa.WorkersDay.Clone(
		&cal.Holiday{
			Name:     "วันแรงงาน",
			Type:     cal.ObservanceBank,
			Observed: weekendAlt,
		},
	)

	// CoronationDay represents King Rama X's Coronation Day on 4 May
	CoronationDay = &cal.Holiday{
		Name:      "วันฉัตรมงคล",
		Type:      cal.ObservancePublic,
		Month:     time.May,
		Day:       4,
		StartYear: 2019,
		Observed:  weekendAlt,
		Func:      cal.CalcDayOfMonth,
	}

	// QueensBirthday represents Queen Suthida's birthday on 3 June
	QueensBirthday = &cal.Holiday{
		Name:      "วันเฉลิมพระชนมพรรษาสมเด็จพระนางเจ้าฯ พระบรมราชินี",
		Type:      cal.ObservancePublic,
		Month:     time.June,
		Day:       3,
		StartYear: 2019,
		Observed:  weekendAlt,
		Func:      cal.CalcDayOfMonth,
	}

	// KingsBirthday represents King Rama X's birthday on 28 July
	KingsBirthday = &cal.Holiday{
		Name:      "วันเฉลิมพระชนมพรรษาพระบาทสมเด็จพระเจ้าอยู่หัว",
		Type:      cal.ObservancePublic,
		Month:     time.July,
		Day:       28,
		StartYear: 2017,
		Observed:  weekendAlt,
		Func:      cal.CalcDayOfMonth,
	}

	// MothersDay represents Mother National Day on 12 Aug
	MothersDay = &cal.Holiday{
		Name:     "วันแม่แห่งชาติ",
		Type:     cal.ObservancePublic,
		Month:    time.August,
		Day:      12,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// KingBhumibolDay represents King Bhumibol Memorial Day on 13 Oct
	KingBhumibolDay = &cal.Holiday{
		Name:      "วันคล้ายวันสวรรคต พระบาทสมเด็จพระปรมินทรมหาภูมิพลอดุลยเดชมหาราช",
		Type:      cal.ObservancePublic,
		Month:     time.October,
		Day:       13,
		Observed:  weekendAlt,
		StartYear: 2017,
		Func:      cal.CalcDayOfMonth,
	}

	// KingChulalongkornDay represents King Chulalongkorn Memorial Day on 23 Oct
	KingChulalongkornDay = &cal.Holiday{
		Name:     "วันปิยมหาราช",
		Type:     cal.ObservancePublic,
		Month:    time.October,
		Day:      23,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// FathersDay represents Father's National Day on 5 Dec
	FathersDay = &cal.Holiday{
		Name:     "วันพ่อแห่งชาติ",
		Type:     cal.ObservancePublic,
		Month:    time.December,
		Day:      5,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// ConstitutionDay represents Constitution Day on 10 Dec
	ConstitutionDay = &cal.Holiday{
		Name:     "วันรัฐธรรมนูญ",
		Type:     cal.ObservancePublic,
		Month:    time.December,
		Day:      10,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// NewYearEve represents NewYear Eve on 31 Dec
	NewYearEve = &cal.Holiday{
		Name:  "วันสิ้นปี",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   31,
		Func:  cal.CalcDayOfMonth,
	}

	// Holidays provides a list of the standard national holidays
	holidays = []*cal.Holiday{
		NewYear,
		ChakriDay,
		CoronationDay,
		QueensBirthday,
		KingsBirthday,
		MothersDay,
		KingBhumibolDay,
		KingChulalongkornDay,
		FathersDay,
		ConstitutionDay,
		NewYearEve,
	}
	Holidays = append(holidays, SongKranDays...)
)
