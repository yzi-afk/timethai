package timethai_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yzi-afk/timethai"
)

var thaiLocate, _ = time.LoadLocation("Asia/Bangkok")

type testArgs struct {
	t time.Time
}

func TestFormat(t *testing.T) {
	type testArgs struct {
		t      time.Time
		layout string
	}

	tests := []struct {
		name string
		args testArgs
		want string
	}{
		{
			name: "TestFormat",
			args: testArgs{
				t:      time.Date(2023, 11, 18, 0, 0, 0, 0, thaiLocate),
				layout: timethai.DateLayout,
			},
			want: "18 พฤศจิกายน 2566",
		},
		{
			name: "TestFormat",
			args: testArgs{
				t:      time.Date(2023, 11, 1, 0, 0, 0, 0, thaiLocate),
				layout: timethai.DateLayout,
			},
			want: "1 พฤศจิกายน 2566",
		},
		{
			name: "TestFormat",
			args: testArgs{
				t:      time.Date(2023, 11, 1, 0, 0, 0, 0, thaiLocate),
				layout: timethai.ShortMonthLayout,
			},
			want: "พ.ย.",
		},
		{
			name: "TestFormat",
			args: testArgs{
				t:      time.Date(2023, 11, 1, 0, 0, 0, 0, thaiLocate),
				layout: timethai.LongMonthLayout,
			},
			want: "พฤศจิกายน",
		},
		{
			name: "TestFormat",
			args: testArgs{
				t:      time.Date(2023, 11, 1, 0, 0, 0, 0, thaiLocate),
				layout: timethai.ShortWeekDayLayout,
			},
			want: "พ.",
		},
		{
			name: "TestFormat",
			args: testArgs{
				t:      time.Date(2023, 11, 1, 0, 0, 0, 0, thaiLocate),
				layout: timethai.LongWeekDayLayout,
			},
			want: "พุธ",
		},
		{
			name: "TestFormat",
			args: testArgs{
				t:      time.Date(2023, 11, 1, 0, 0, 0, 0, thaiLocate),
				layout: timethai.FullDateTimeLayout,
			},
			want: "วันพุธที่ 1 เดือนพฤศจิกายน พ.ศ. 2566 เวลา 0 นาฬิกา 0 นาที",
		},
		{
			name: "TestFormat",
			args: testArgs{
				t:      time.Date(2023, 11, 1, 0, 0, 0, 0, thaiLocate),
				layout: timethai.DateTimeLayout,
			},
			want: "1 พฤศจิกายน 2566 00:00",
		},
		{
			name: "TestFormat",
			args: testArgs{
				t:      time.Date(2023, 11, 1, 0, 0, 0, 0, thaiLocate),
				layout: timethai.TimeLayout,
			},
			want: "00:00",
		},
		{
			name: "TestFormat",
			args: testArgs{
				t:      time.Date(2023, 1, 1, 0, 0, 0, 0, thaiLocate),
				layout: timethai.ThaiYearLayout,
			},
			want: "2566",
		},
		{
			name: "TestFormat",
			args: testArgs{
				t:      time.Date(2023, 1, 1, 0, 0, 0, 0, thaiLocate),
				layout: time.RFC3339,
			},
			want: "2023-01-01T00:00:00+07:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, timethai.Format(tt.args.t, tt.args.layout))
		})
	}
}

func TestFormatISO8601(t *testing.T) {
	tests := []struct {
		name string
		args testArgs
		want string
	}{
		{
			name: "TestFormatISO8601",
			args: testArgs{
				t: time.Date(2023, 11, 18, 0, 0, 0, 0, thaiLocate),
			},
			want: "2023-11-18T00:00:00+0700",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, timethai.FormatISO8601(tt.args.t))
		})
	}
}

func TestGetShortWeekDay(t *testing.T) {
	tests := []struct {
		name string
		args testArgs
		want string
	}{
		{
			name: "TestGetShortDay",
			args: testArgs{
				t: time.Date(2023, 11, 18, 0, 0, 0, 0, thaiLocate),
			},
			want: "ส.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, timethai.GetShortWeekDay(tt.args.t))
		})
	}
}

func TestGetLongWeekDay(t *testing.T) {
	tests := []struct {
		name string
		args testArgs
		want string
	}{
		{
			name: "TestGetLongWeekDay",
			args: testArgs{
				t: time.Date(2023, 11, 18, 0, 0, 0, 0, thaiLocate),
			},
			want: "เสาร์",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, timethai.GetLongWeekDay(tt.args.t))
		})
	}
}

func TestGetShortMonth(t *testing.T) {
	tests := []struct {
		name string
		args testArgs
		want string
	}{
		{
			name: "TestGetShortMonth_January",
			args: testArgs{
				t: time.Date(2023, 1, 18, 0, 0, 0, 0, thaiLocate),
			},
			want: "ม.ค.",
		},
		{
			name: "TestGetShortMonth",
			args: testArgs{
				t: time.Date(2023, 11, 18, 0, 0, 0, 0, thaiLocate),
			},
			want: "พ.ย.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, timethai.GetShortMonth(tt.args.t))
		})
	}
}

func TestGetLongMonth(t *testing.T) {
	tests := []struct {
		name string
		args testArgs
		want string
	}{
		{
			name: "TestGetLongMonth",
			args: testArgs{
				t: time.Date(2023, 11, 18, 0, 0, 0, 0, thaiLocate),
			},
			want: "พฤศจิกายน",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, timethai.GetLongMonth(tt.args.t))
		})
	}
}

func TestGetBuddhistYear(t *testing.T) {
	tests := []struct {
		name string
		args testArgs
		want int
	}{
		{
			name: "TestGetBuddhistYear",
			args: testArgs{
				t: time.Date(2023, 11, 18, 0, 0, 0, 0, thaiLocate),
			},
			want: 2566,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, timethai.GetBuddhistYear(tt.args.t))
		})
	}
}

func TestLocaleTimeString(t *testing.T) {
	tests := []struct {
		name string
		args testArgs
		want string
	}{
		{
			name: "TestLocaleTimeString",
			args: testArgs{
				t: time.Date(2023, 11, 18, 0, 0, 0, 0, thaiLocate),
			},
			want: "เที่ยงคืน",
		},
		{
			name: "TestLocaleTimeString",
			args: testArgs{
				t: time.Date(2023, 11, 18, 2, 0, 0, 0, thaiLocate),
			},
			want: "ตี 2",
		},
		{
			name: "TestLocaleTimeString",
			args: testArgs{
				t: time.Date(2023, 11, 18, 12, 0, 0, 0, thaiLocate),
			},
			want: "เที่ยง",
		},
		{
			name: "TestLocaleTimeString",
			args: testArgs{
				t: time.Date(2023, 11, 18, 13, 0, 0, 0, thaiLocate),
			},
			want: "บ่ายโมง",
		},
		{
			name: "TestLocaleTimeString",
			args: testArgs{
				t: time.Date(2023, 11, 18, 14, 0, 0, 0, thaiLocate),
			},
			want: "บ่าย 2 โมง",
		},
		{
			name: "TestLocaleTimeString",
			args: testArgs{
				t: time.Date(2023, 11, 18, 16, 0, 0, 0, thaiLocate),
			},
			want: "4 โมงเย็น",
		},
		{
			name: "TestLocaleTimeString",
			args: testArgs{
				t: time.Date(2023, 11, 18, 20, 30, 0, 0, thaiLocate),
			},
			want: "8 ทุ่ม 30 นาที",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, timethai.LocaleTimeString(tt.args.t))
		})
	}
}

func TestLocaleOfficialTimeString(t *testing.T) {
	tests := []struct {
		name string
		args testArgs
		want string
	}{
		{
			name: "TestLocaleOfficialTimeString",
			args: testArgs{
				t: time.Date(2023, 11, 18, 0, 0, 0, 0, thaiLocate),
			},
			want: "0 นาฬิกา 0 นาที",
		}, {
			name: "TestLocaleOfficialTimeString",
			args: testArgs{
				t: time.Date(2023, 11, 18, 20, 30, 0, 0, thaiLocate),
			},
			want: "20 นาฬิกา 30 นาที",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, timethai.LocaleOfficialTimeString(tt.args.t))
		})
	}
}

func TestLocaleDateString(t *testing.T) {
	tests := []struct {
		name string
		args testArgs
		want string
	}{
		{
			name: "TestLocaleDateString",
			args: testArgs{
				t: time.Date(2023, 11, 18, 0, 0, 0, 0, thaiLocate),
			},
			want: "18 พฤศจิกายน 2566",
		},
		{
			name: "TestLocaleDateString",
			args: testArgs{
				t: time.Date(2023, 1, 18, 0, 0, 0, 0, time.UTC),
			},
			want: "18 มกราคม 2566",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, timethai.LocaleDateString(tt.args.t))
		})
	}
}
