package helpers

import (
	"time"

	"github.com/golang-module/carbon/v2"
)

func ParseDateFormat(date time.Time) string {
	convertCarbon := carbon.CreateFromStdTime(date)

	return convertCarbon.Format("d M Y H:i:s")
}
