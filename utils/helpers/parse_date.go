package helpers

import (
	"time"

	"github.com/golang-module/carbon/v2"
)

type Date string

const (
	Senin  Date = "Senin"
	Selasa Date = "Selasa"
	Rabu   Date = "Rabu"
	Kamis  Date = "Kamis"
	Jumat  Date = "Jumat"
	Sabtu  Date = "Sabtu"
	Minggu Date = "Minggu"
)
const (
	Mon Date = "Mon"
	Tue Date = "Tue"
	Wed Date = "Wed"
	Thu Date = "Thu"
	Fri Date = "Fri"
	Sat Date = "Sat"
	Sun Date = "Sun"
)

func ParseDateFormat(date *time.Time) string {
	convertCarbon := carbon.CreateFromStdTime(*date)

	return convertCarbon.Format("d M Y H:i:s")
}

func ParseOnlyDate(date *time.Time) string {
	convertCarbon := carbon.CreateFromStdTime(*date)

	return convertCarbon.Format("d M Y")
}

func ParseStringToTime(date string) *time.Time {
	convert := carbon.Parse(date).ToStdTime()
	return &convert
}

func GetDayToTime(date time.Time) string {
	GetDay := carbon.CreateFromStdTime(date).Format("D")

	if GetDay == string(Sun) {
		return string(Minggu)
	}
	if GetDay == string(Mon) {
		return string(Senin)
	}
	if GetDay == string(Tue) {
		return string(Selasa)
	}
	if GetDay == string(Wed) {
		return string(Rabu)
	}
	if GetDay == string(Thu) {
		return string(Kamis)
	}
	if GetDay == string(Fri) {
		return string(Jumat)
	}
	if GetDay == string(Sat) {
		return string(Sabtu)
	}

	return "No date"
}

func ParseClockToTime(time string) time.Time {

	Convert := carbon.Parse("0001-01-01 " + time)

	return Convert.ToStdTime()
}

func ParseTimeToClock(time *time.Time) string {
	convert := carbon.CreateFromStdTime(*time)

	return convert.Format("H:i:s")
}
