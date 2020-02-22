package cal

import (
	"math"
	"time"
)

// AddUkraineHolidays adds all Ukraine holidays to the Calendar
func AddUkraineHolidays(c *Calendar) {
	c.AddHoliday(uaHolidays()...)
}

type holidayRule struct {
	name            string
	startYear       int
	endYear         int
	day             int
	month           time.Month
	isEasterDepend  bool
	easterOffsetDay int
}

func uaHolidays() []Holiday {
	var uaHolidays []Holiday

	// Holidays in Ukraine
	// Reference https://en.wikipedia.org/wiki/Public_holidays_in_Ukraine
	var holidayRules = []holidayRule{
		{name: "New Year", month: time.January, day: 1},
		{name: "Labour Day 1", month: time.May, day: 1},
		{name: "Labour Day 2", endYear: 2018, month: time.May, day: 2},
		{name: "Defender Of Ukraine day", startYear: 2015, month: time.October, day: 14},
		{name: "Catholic Christmas day", startYear: 2017, month: time.December, day: 25},
		{name: "Orthodox Christmas day", month: time.January, day: 7},
		{name: "Women Day", month: time.March, day: 8},
		{name: "Victory Day", month: time.May, day: 9},
		{name: "Constitution Day", month: time.June, day: 28},
		{name: "Independence Day", month: time.August, day: 24},
		{name: "Orthodox Easter Day", isEasterDepend: true, easterOffsetDay: 0},
		{name: "Orthodox Pentecost Day", isEasterDepend: true, easterOffsetDay: 49},
	}

	for _, rule := range holidayRules {
		rule := rule
		uaHolidays = append(uaHolidays, holidayByRule(&rule, false), holidayByRule(&rule, true))
	}

	return uaHolidays
}

func holidayByRule(h *holidayRule, isExtendHoliday bool) Holiday {
	return NewHolidayFunc(func(year int, loc *time.Location) (time.Month, int) {
		var d time.Time

		if h.isEasterDepend {
			easter := calculateJulianEaster(year, loc)
			d = easter.AddDate(0, 0, h.easterOffsetDay)
		} else {
			startYear := h.startYear
			endYear := h.endYear
			month := h.month
			day := h.day

			if ((startYear == 0) || (startYear > 0 && year > startYear)) && ((endYear == 0) || (endYear > 0 && year < endYear)) {
				d = time.Date(year, month, day, 0, 0, 0, 0, loc)
			}
		}

		//Extended holiday on first working day for Saturday and Sunday
		if isExtendHoliday {
			switch int(d.Weekday()) {
			case 6:
				d = d.AddDate(0, 0, 2)
			case 0:
				d = d.AddDate(0, 0, 1)
			}
		}

		return d.Month(), d.Day()
	})
}

func calculateJulianEaster(year int, loc *time.Location) time.Time {
	//based on the Meeus Julian algorithm
	a := year % 4
	b := year % 7
	c := year % 19
	d := (19*c + 15) % 30
	e := (2*a + 4*b - d + 34) % 7
	month := math.Floor(float64((d + e + 114) / 31))
	day := ((d + e + 114) % 31) + 1 + 13

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
}
