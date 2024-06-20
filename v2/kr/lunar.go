package kr

// Get source from https://github.com/godcong/chronos and rewrite.
// cause upon package does not export functions.

import (
    "log"
    "sort"
    "time"

    "github.com/rickar/cal/v2"
)

const (
    MINDATE = -2206513672 // 1900-01-30 unix timestamp
    MAXDATE = 4102326000  // 2099-12-31 unix timestamp

    MINYEAR = 1900
    MAXYEAR = 2100
)

// Sol2Lun converts to lunar year/month/day and if it is belongs to the leap month
// Chinese lunar calendar
func Sol2Lun(t time.Time) (ymd [3]int, isLeap bool) {

    offset := offsetDays(MINDATE, t)
    year, offset := lunarYearByOffset(offset)
    leapMonth := leapMonth(year)
    tmp := leapMonth > 0

    var month, days int
    for month = 1; month <= 12; month++ {
        if month == leapMonth+1 && tmp {
            days = lengthOfLeapMonth(year)
            tmp = false
            month--
        } else {
            days = lengthOfMonth(year, month)
        }

        if offset-days <= 0 {
            break
        }

        offset -= days
    }

    ymd[0], ymd[1], ymd[2] = year, month, offset
    return ymd, month == leapMonth
}

// offsetDays count offsets(days) between two timestamps
func offsetDays(timestamp int64, targetDate time.Time) int {
    return int(float64(targetDate.Unix()-timestamp)/86400.0 + 0.5)
}

// lunarYearByOffset calculate lunar year since MINYEAR by offset days
func lunarYearByOffset(offset int) (int, int) {
    var year int
    for year = MINYEAR; year <= MAXYEAR; year++ {
        days := lengthOfYear(year)
        if offset-days < 1 {
            break
        }
        offset -= days
    }

    return year, offset
}

// lengthOfYear length of days of year
func lengthOfYear(year int) int {
    length := 348
    info := getLunarInfo(year)

    // 1000000000000000 to 10000
    for mask := 0x8000; mask > 0x8; mask >>= 1 {
        if info&mask != 0 {
            length++
        }
    }

    return length + lengthOfLeapMonth(year)
}

// lengthOfMonth length of days of year/month
func lengthOfMonth(year, month int) int {
    if month < 1 || month > 12 {
        return -1
    }

    info := getLunarInfo(year)
    if info&(0x10000>>uint32(month)) != 0 {
        return 30
    }

    return 29
}

// lengthOfLeapMonth length of days of year/month if it is leap month
func lengthOfLeapMonth(year int) int {
    if !hasLeapMonth(year) {
        return 0
    }

    info := getLunarInfo(year)
    if info&0x10000 != 0 {
        return 30
    }
    return 29
}

// leapMonth returns leap month if leap year
func leapMonth(year int) int {
    info := getLunarInfo(year)
    return info & 0xf
}

// hasLeapMonth `year` has leap month or not
func hasLeapMonth(year int) bool {
    return leapMonth(year) != 0
}

// isLeapMonth is leap month or not
func isLeapMonth(year, month int) bool {
    return leapMonth(year) == month
}

// getLunarInfo get pre-calcuated data by lunar year
func getLunarInfo(year int) int {
    idx := year - MINYEAR
    if idx < 0 || idx > len(LunarInfoList) {
        return 0
    }
    return LunarInfoList[idx]
}

// Solar2Lunar converts to lunar year/month/day and if it is belongs to the leap month
// Korean lunar calendar
func Solar2Lunar(t time.Time) (ymd [3]int, isLeap bool) {
    const (
        BASEYEAR = 1391
        MINDATE  = 2229156 // 1391-02-05 ( lunisolar 1391-01-01 )
        MAXDATE  = 2470172 // 2050-12-31 ( lunisolar 2050-11-18 )
    )

    // TODO gregorian 1582-10~05 ~ 1582-10-14 dates do not exist.
    days := julian(t)
    if days < MINDATE || days > MAXDATE {
        log.Fatalf("The date is out of range: %d", days)
    }

    days -= MINDATE
    month := sort.Search(len(MonthTable), func(i int) bool { return MonthTable[i] > days }) - 1
    if month > len(MonthTable)-1 {
        log.Fatalf("Out of MonthTable range: %d/%d", month, len(MonthTable))
    }

    year := sort.Search(len(YearTable), func(i int) bool { return YearTable[i] > month }) - 1
    if year > len(YearTable)-1 {
        log.Fatalf("Out of YearTable range: %d/%d", year, len(YearTable))
    }

    var day int
    month, day = month-YearTable[year]+1, days-MonthTable[month]+1
    if LeapTable[year] != 0 && LeapTable[year] <= month {
        if LeapTable[year] == month {
            isLeap = true
        } else {
            isLeap = false
        }
        month--
    } else {
        isLeap = false
    }

    ymd[0], ymd[1], ymd[2] = year+BASEYEAR, month, day
    return ymd, isLeap
}

// Lunar2Solar converts to solar calendar time
func Lunar2Solar(year, month, day int, isLeap bool) time.Time {
    const (
        BASEYEAR = 1391
        MINDATE  = 2229156 // 1391-02-05 ( lunisolar 1391-01-01 )
        MAXDATE  = 2470172 // 2050-12-31 ( lunisolar 2050-11-18 )
    )

    year -= BASEYEAR
    if year < 0 || year >= len(YearTable) {
        log.Fatal("Year is out of range")
    }

    if month < 1 || month > 12 {
        log.Fatal("Month is out of range")
    }

    if isLeap && LeapTable[year]-1 != month {
        log.Fatal("Wrong leap month")
    }

    months := YearTable[year] + month - 1
    if isLeap && month+1 == LeapTable[year] {
        months++
    } else if LeapTable[year] != 0 && LeapTable[year] < month {
        months++
    }

    days := MonthTable[months] + day - 1
    if day < 1 || days >= MonthTable[months+1] {
        log.Fatal("Wrong day")
    }

    y, m, d := julianDayToDate(days + MINDATE)
    return time.Date(y, time.Month(m), d, 0, 0, 0, 0, cal.DefaultLoc)
}

func julianDayToDate(jd int) (year, month, day int) {
    if jd >= 0 {
        if jd >= 2299161 {
            year, month, day = julianDayToGregorianDate(jd)
        } else {
            year, month, day = julianDayToJulianDate(jd)
        }
    }
    return
}

// julian https://github.com/toelsiba/date/blob/master/julian_day.go
// Gregorian calendar starting from October 15, 1582
// Algorithm from Henry F. Fliegel and Thomas C. Van Flandern
func julian(t time.Time) (jd int) {
    year, month, day := t.Year(), int(t.Month()), t.Day()
    e := (1461 * (year + 4800 + (month-14)/12)) / 4
    e += (367 * (month - 2 - 12*((month-14)/12))) / 12
    e += -(3 * ((year + 4900 + (month-14)/12) / 100)) / 4
    e += day - 32075
    return e
}

// Julian calendar until October 4, 1582
// Algorithm from Frequently Asked Questions about Calendars by Claus Toendering
func julianDayToJulianDate(jd int) (year, month, day int) {
    jd += 32082
    dd := (4*jd + 3) / 1461
    ee := jd - (1461*dd)/4
    mm := ((5 * ee) + 2) / 153
    day = ee - (153*mm+2)/5 + 1
    month = mm + 3 - 12*(mm/10)
    year = dd - 4800 + (mm / 10)
    if year <= 0 {
        year--
    }
    return
}

// Gregorian calendar starting from October 15, 1582
// This algorithm is from Henry F. Fliegel and Thomas C. Van Flandern
func julianDayToGregorianDate(jd int) (year, month, day int) {
    ell := uint64(jd) + 68569
    n := (4 * ell) / 146097
    ell = ell - (146097*n+3)/4
    i := (4000 * (ell + 1)) / 1461001
    ell = ell - (1461*i)/4 + 31
    j := (80 * ell) / 2447
    day = int(ell - (2447*j)/80)
    ell = j / 11
    month = int(j + 2 - (12 * ell))
    year = int(100*(n-49) + i + ell)
    return
}
