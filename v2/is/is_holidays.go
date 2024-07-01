package is

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Nyarsdagur represents New Year's Day on 1-Jan
	Nyarsdagur = aa.NewYear.Clone(&cal.Holiday{Name: "Nýársdagur", Type: cal.ObservancePublic})

	// Skirdagur represents Maundy Thursday on the Thursday before Easter
	Skirdagur = aa.MaundyThursday.Clone(&cal.Holiday{Name: "Skírdagur", Type: cal.ObservancePublic})

	// Langifostudagur represents Good Friday on the Friday before Easter
	Langifostudagur = aa.GoodFriday.Clone(&cal.Holiday{Name: "Föstudagurinn langi", Type: cal.ObservancePublic})

	// Annaripaskum represents Easter Monday on the day after Easter
	Annaripaskum = aa.EasterMonday.Clone(&cal.Holiday{Name: "Annar í páskum", Type: cal.ObservancePublic})

	// Sumardagurinn represents the First Day of Summer on the first Thursday after 18-Apr
	Sumardagurinn = &cal.Holiday{
		Name:    "Sumardagurinn fyrsti",
		Type:    cal.ObservancePublic,
		Month:   time.April,
		Day:     19,
		Offset:  1,
		Weekday: time.Thursday,
		Func:    cal.CalcWeekdayFrom,
	}

	// Verkalydsdagurinn represents Labour Day on 1-May
	Verkalydsdagurinn = aa.WorkersDay.Clone(&cal.Holiday{Name: "Verkalýðsdagurinn", Type: cal.ObservancePublic})

	// Uppstigningardagur represents Ascension Day on the 39th day after Easter
	Uppstigningardagur = aa.AscensionDay.Clone(&cal.Holiday{Name: "Uppstigningardagur", Type: cal.ObservancePublic})

	// Annarihvit represents Whit Monday on the day after Pentecost
	Annarihvit = aa.PentecostMonday.Clone(&cal.Holiday{Name: "Annar í hvítasunnu", Type: cal.ObservancePublic})

	// Thjodhatid represents Independence Day on 17-Jun
	Thjodhatid = &cal.Holiday{
		Name:  "Þjóðhátíðardagurinn",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   17,
		Func:  cal.CalcDayOfMonth,
	}

	// Verslunarmannahelgi represents Commerce Day on the first Monday in August
	Verslunarmannahelgi = &cal.Holiday{
		Name:    "Frídagur verslunarmanna",
		Type:    cal.ObservancePublic,
		Month:   time.August,
		Offset:  1,
		Weekday: time.Monday,
		Func:    cal.CalcWeekdayOffset,
	}

	// Adfangadagur represents Christmas Eve on 24-Dec
	Adfangadagur = &cal.Holiday{
		Name:  "Aðfangadagur",
		Type:  cal.ObservanceOther,
		Month: time.December,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// Joladagur represents Christmas Day on 25-Dec
	Joladagur = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Jóladagur", Type: cal.ObservancePublic})

	// Annarijolum represents the second day of Christmas on 26-Dec
	Annarijolum = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Annar í jólum", Type: cal.ObservancePublic})

	// Gamarsdagur represents New Year's Eve on 31-Dec
	Gamarsdagur = &cal.Holiday{
		Name:  "Gamlársdagur",
		Type:  cal.ObservanceOther,
		Month: time.December,
		Day:   31,
		Func:  cal.CalcDayOfMonth,
	}

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Nyarsdagur,
		Skirdagur,
		Langifostudagur,
		Annaripaskum,
		Sumardagurinn,
		Verkalydsdagurinn,
		Uppstigningardagur,
		Annarihvit,
		Thjodhatid,
		Verslunarmannahelgi,
		Adfangadagur,
		Joladagur,
		Annarijolum,
		Gamarsdagur,
	}
)
