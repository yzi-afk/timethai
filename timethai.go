package timethai

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	DateLayout         = time.DateOnly
	TimeLayout         = time.TimeOnly
	DateTimeLayout     = DateLayout + " " + TimeLayout
	ShortMonthLayout   = "ม.ค."
	LongMonthLayout    = "มกราคม"
	ShortWeekDayLayout = "อา."
	LongWeekDayLayout  = "อาทิตย์"
	FullDateTimeLayout = "วัน" + LongWeekDayLayout + "ที่ 02 เดือน" + LongMonthLayout + " พ.ศ. 2006 เวลา 15:04 น."
	ThaiYearLayout     = "2006"
)

var shortWeekDay = map[int]string{
	0: "อา.",
	1: "จ.",
	2: "อ.",
	3: "พ.",
	4: "พฤ.",
	5: "ศ.",
	6: "ส.",
}

var longWeekDay = map[int]string{
	0: "อาทิตย์",
	1: "จันทร์",
	2: "อังคาร",
	3: "พุธ",
	4: "พฤหัสบดี",
	5: "ศุกร์",
	6: "เสาร์",
}

var shortMonth = map[int]string{
	1:  "ม.ค.",
	2:  "ก.พ.",
	3:  "มี.ค.",
	4:  "เม.ย.",
	5:  "พ.ค.",
	6:  "มิ.ย.",
	7:  "ก.ค.",
	8:  "ส.ค.",
	9:  "ก.ย.",
	10: "ต.ค.",
	11: "พ.ย.",
	12: "ธ.ค.",
}

var longMonth = map[int]string{
	1:  "มกราคม",
	2:  "กุมภาพันธ์",
	3:  "มีนาคม",
	4:  "เมษายน",
	5:  "พฤษภาคม",
	6:  "มิถุนายน",
	7:  "กรกฎาคม",
	8:  "สิงหาคม",
	9:  "กันยายน",
	10: "ตุลาคม",
	11: "พฤศจิกายน",
	12: "ธันวาคม",
}

func Format(t time.Time, layout string) string {
	switch layout {
	case DateLayout:
		return FormatDate(t)
	case TimeLayout:
		return FormatTime(t)
	case DateTimeLayout:
		return FormatDateTime(t)
	case ShortMonthLayout:
		return GetShortMonth(t)
	case LongMonthLayout:
		return GetLongMonth(t)
	case ShortWeekDayLayout:
		return GetShortWeekDay(t)
	case LongWeekDayLayout:
		return GetLongWeekDay(t)
	case FullDateTimeLayout:
		return FormatFullDateTime(t)
	case ThaiYearLayout:
		return strconv.Itoa(GetBuddhistYear(t))
	default:
		return t.Format(layout)
	}
}

func FormatISO8601(t time.Time) string {
	return t.Format("2006-01-02T15:04:05-0700")
}

func FormatDate(t time.Time) string {
	return fmt.Sprintf("%d %s %d", t.Day(), GetLongMonth(t), GetBuddhistYear(t))
}

func FormatTime(t time.Time) string {
	return fmt.Sprintf("%02d:%02d", t.Hour(), t.Minute())
}

func FormatDateTime(t time.Time) string {
	return fmt.Sprintf("%s %s", FormatDate(t), FormatTime(t))
}

func FormatFullDateTime(t time.Time) string {
	return fmt.Sprintf("วัน%sที่ %d เดือน%s พ.ศ. %d เวลา %s", GetLongWeekDay(t), t.Day(), GetLongMonth(t), GetBuddhistYear(t), LocaleOfficialTimeString(t))
}

func GetShortWeekDay(t time.Time) string {
	return shortWeekDay[int(t.Weekday())]
}

func GetLongWeekDay(t time.Time) string {
	return longWeekDay[int(t.Weekday())]
}

func GetShortMonth(t time.Time) string {
	return shortMonth[int(t.Month())]
}

func GetLongMonth(t time.Time) string {
	return longMonth[int(t.Month())]
}

func GetBuddhistYear(t time.Time) int {
	return t.Year() + 543
}

func LocaleTimeString(t time.Time) string {
	const halfDay = 12
	var str strings.Builder
	hour := t.Hour()
	minute := t.Minute()
	switch hour {
	case 0:
		str.WriteString("เที่ยงคืน")
	case 1, 2, 3, 4, 5:
		str.WriteString("ตี " + strconv.Itoa(hour))
	case 6, 7, 8, 9, 10, 11:
		str.WriteString(strconv.Itoa(hour) + " โมงเช้า")
	case 12:
		str.WriteString("เที่ยง")
	case 13:
		str.WriteString("บ่ายโมง")
	case 14, 15:
		str.WriteString("บ่าย " + strconv.Itoa(hour-halfDay) + " โมง")
	case 16, 17, 18:
		str.WriteString(strconv.Itoa(hour-halfDay) + " โมงเย็น")
	case 19, 20, 21, 22, 23:
		str.WriteString(strconv.Itoa(hour-halfDay) + " ทุ่ม")
	}

	if minute > 0 {
		str.WriteString(" ")
		str.WriteString(strconv.Itoa(minute))
		str.WriteString(" นาที")
	}

	return str.String()
}

func LocaleOfficialTimeString(t time.Time) string {
	var str strings.Builder
	hour := t.Hour()
	minute := t.Minute()

	str.WriteString(strconv.Itoa(hour))
	str.WriteString(" นาฬิกา")

	str.WriteString(" ")
	str.WriteString(strconv.Itoa(minute))
	str.WriteString(" นาที")

	return str.String()
}

func LocaleDateString(t time.Time) string {
	var str strings.Builder
	str.WriteString(strconv.Itoa(t.Day()))
	str.WriteString(" ")
	str.WriteString(GetLongMonth(t))
	str.WriteString(" ")
	str.WriteString(strconv.Itoa(GetBuddhistYear(t)))
	return str.String()
}
