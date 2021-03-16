package kr

// Get source from https://github.com/godcong/chronos and rewrite.
// cause upon package does not export functions.

import (
    "time"
)

const (
    MINDATE = -2206513672 // 1900-01-30 unix timestamp
    MAXDATE = 4102326000  // 2099-12-31 unix timestamp

    MINYEAR = 1900
    MAXYEAR = 2100
)

// Lun2Sol converts to solar calendar time
// TODO: implement if needed
// func Lun2Sol(year, month, day int) time.Time {
//     return time.Now()
// }

// Sol2Lun converts to lunar year/month/day and if it is belongs to the leap month
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
