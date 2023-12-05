package helpers

import (
	"time"

	"github.com/golang-module/carbon/v2"
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

func ParseClockToTime(time string) time.Time {

	Convert := carbon.Parse("0001-01-01 " + time)

	return Convert.ToStdTime()
}

func ParseTimeToClock(time *time.Time) string {
	convert := carbon.CreateFromStdTime(*time)

	return convert.Format("H:i:s")
}
