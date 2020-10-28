package timer

import (
	"fmt"
	"time"
)

func FormatDuration(d time.Duration) string {
	neg := false
	if d < 0 {
		neg = true
		d *= -1
	}

	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	result := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	if neg {
		result = "-" + result
	}
	return result
}
