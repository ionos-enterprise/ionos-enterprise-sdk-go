package profitbricks

import "time"

// DurationOrDefault returns a default value if the given duration is 0.
func DurationOrDefault(value time.Duration, defaultValue time.Duration) time.Duration {
	if value == 0 {
		return defaultValue
	}
	return value
}
