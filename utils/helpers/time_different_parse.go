package helpers

import (
	"fmt"
	"time"
)

func GetDurationTime(published_at time.Time) string {
	d := time.Since(published_at)

	switch {
	case d.Seconds() < 60:
		return fmt.Sprintf("%.0f detik lalu", d.Seconds())
	case d.Minutes() < 60:
		return fmt.Sprintf("%.0f menit lalu", d.Minutes())
	case d.Hours() < 24:
		return fmt.Sprintf("%.0f jam lalu", d.Hours())
	default:
		return fmt.Sprintf("%.0f hari lalu", d.Hours()/24)
	}
}
