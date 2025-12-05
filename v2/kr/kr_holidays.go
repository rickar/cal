// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package kr provides holiday definitions for the Republic of Korea.
package kr

import (
    "time"

    "github.com/rickar/cal/v2"
    "github.com/rickar/cal/v2/aa"
)

var (
    // Standard KR weekend substitution rules:
    //   https://ko.wikipedia.org/wiki/%EB%8C%80%ED%95%9C%EB%AF%BC%EA%B5%AD%EC%9D%98_%EA%B3%B5%ED%9C%B4%EC%9D%BC
    //   If the Lunar New Year holiday and Chuseok holiday overlap with other public holidays, including Sundays,
    //   the first weekday after the holiday shall be an alternative holiday, and if Children's Day overlaps with other
    //   public holidays or Saturdays, the first weekday shall be an alternative holiday (Article 3.[1] New November 5, 2013).

    //   설날 연휴와 추석 연휴가 일요일을 포함한 다른 공휴일과 겹칠 경우에는 연휴 다음 첫번째 평일을 대체공휴일로 하고,
    //   어린이날이 다른 공휴일이나 토요일과 겹칠 경우에는 그 다음 첫번째 평일을 대체공휴일로 한다.
    //   (관공서의 공휴일에 관한 규정 제3조.[1] 2013년 11월 5일 신설)

    //   Saturdays and Sundays move to next weekday

    weekendAlt = []cal.AltDay{
        {Day: time.Saturday, Offset: 2},
        {Day: time.Sunday, Offset: 1},
    }

    LunarNewYearChuseokAlt = []cal.AltDay{
        {Day: time.Saturday, Offset: 2},
        {Day: time.Sunday, Offset: 2},
    }

    // NewYear represents New Year's Day on 1-Jan
    NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "새해 첫날", Type: cal.ObservancePublic})

    // BeforeLunarNewYear represents 설전날 before on the Lunar 1-Jan
    BeforeLunarNewYear = &cal.Holiday{
        Name:  "설전날",
        Type:  cal.ObservancePublic,
        Month: time.January,
        Day:   1,
        Lunar: true,
        Func:  CalcBeforeDayOfLunarMonth,
    }

    // LunarNewYear represents 설날 on the Lunar 1-Jan
    LunarNewYear = &cal.Holiday{
        Name:     "설날",
        Type:     cal.ObservancePublic,
        Month:    time.January,
        Day:      1,
        Lunar:    true,
        Observed: LunarNewYearChuseokAlt,
        Func:     CalcDayOfLunarMonth,
    }

    // AfterLunarNewYear represents 설다음날 after on the Lunar 1-Jan
    AfterLunarNewYear = &cal.Holiday{
        Name:  "설다음날",
        Type:  cal.ObservancePublic,
        Month: time.January,
        Day:   1,
        Lunar: true,
        Func:  CalcAfterDayOfLunarMonth,
    }

    // IndependenceMovementDay represents 3.1절 on the 1-Mar
    IndependenceMovementDay = &cal.Holiday{
        Name:  "3.1절",
        Type:  cal.ObservancePublic,
        Month: time.March,
        Day:   1,
        Func:  cal.CalcDayOfMonth,
    }

    // ChildrensDay represents 어린이날 on the 5-May
    ChildrensDay = &cal.Holiday{
        Name:     "어린이날",
        Type:     cal.ObservancePublic,
        Month:    time.May,
        Day:      5,
        Observed: weekendAlt,
        Func:     cal.CalcDayOfMonth,
    }

    // BuddhasDay represents 부처님 오신 날 on the Lunar 8-April
    BuddhasDay = &cal.Holiday{
        Name:  "부처님 오신 날",
        Type:  cal.ObservancePublic,
        Month: time.April,
        Day:   8,
        Lunar: true,
        Func:  CalcDayOfLunarMonth,
    }

    // MemorialDay represents 현충일 on the 6-June
    MemorialDay = &cal.Holiday{
        Name:  "현충일",
        Type:  cal.ObservancePublic,
        Month: time.June,
        Day:   6,
        Func:  cal.CalcDayOfMonth,
    }

    // LiberationDay represents 광복절 on the 15-August
    LiberationDay = &cal.Holiday{
        Name:  "광복절",
        Type:  cal.ObservancePublic,
        Month: time.August,
        Day:   15,
        Func:  cal.CalcDayOfMonth,
    }

    // BeforeChuseok represents 추석전날 before on the Lunar 15-August
    BeforeChuseok = &cal.Holiday{
        Name:  "추석전날",
        Type:  cal.ObservancePublic,
        Month: time.August,
        Day:   15,
        Lunar: true,
        Func:  CalcBeforeDayOfLunarMonth,
    }

    // Chuseok represents 추석 on the Lunar 15-August
    Chuseok = &cal.Holiday{
        Name:     "추석",
        Type:     cal.ObservancePublic,
        Month:    time.August,
        Day:      15,
        Lunar:    true,
        Observed: LunarNewYearChuseokAlt,
        Func:     CalcDayOfLunarMonth,
    }

    // AfterChuseok represents 추석전날 after on the Lunar 15-August
    AfterChuseok = &cal.Holiday{
        Name:  "추석다음날",
        Type:  cal.ObservancePublic,
        Month: time.August,
        Day:   15,
        Lunar: true,
        Func:  CalcAfterDayOfLunarMonth,
    }

    // FoundationDay represents 개천절 on the 3-October
    FoundationDay = &cal.Holiday{
        Name:  "개천절",
        Type:  cal.ObservancePublic,
        Month: time.October,
        Day:   3,
        Func:  cal.CalcDayOfMonth,
    }

    // HangeulDay represents 개천절 on the 9-October
    HangeulDay = &cal.Holiday{
        Name:  "한글날",
        Type:  cal.ObservancePublic,
        Month: time.October,
        Day:   9,
        Func:  cal.CalcDayOfMonth,
    }

    // ChristmasDay represents 성탄절 on the 25-Dec
    ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "성탄절", Type: cal.ObservancePublic})

    // Holidays provides a list of the standard national holidays
    Holidays = []*cal.Holiday{
        NewYear,
        BeforeLunarNewYear,
        LunarNewYear,
        AfterLunarNewYear,
        IndependenceMovementDay,
        BuddhasDay,
        ChildrensDay,
        MemorialDay,
        LiberationDay,
        BeforeChuseok,
        Chuseok,
        AfterChuseok,
        FoundationDay,
        HangeulDay,
        ChristmasDay,
    }
)

// CalcDayOfLunarMonth calculates the occurrence of a holiday that is always a
// specific day of the month such as the 5th of November.
func CalcDayOfLunarMonth(h *cal.Holiday, year int) time.Time {
    return Lunar2Solar(year, int(h.Month), h.Day, false)
}

// CalcBeforeDayOfLunarMonth a before day of specific LunarDate.
// 1/1 before day maybe become 12/29 or 12/30
// so treat before day as solar calendar
func CalcBeforeDayOfLunarMonth(h *cal.Holiday, year int) time.Time {
    return Lunar2Solar(year, int(h.Month), h.Day, false).Add(time.Hour * 24 * -1)
}

// CalcAfterDayOfLunarMonth a after day of specific LunarDate.
func CalcAfterDayOfLunarMonth(h *cal.Holiday, year int) time.Time {
    return Lunar2Solar(year, int(h.Month), h.Day, false).Add(time.Hour * 24)
}
