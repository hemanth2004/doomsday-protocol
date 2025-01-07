package util

import "fmt"

// FormatTime takes an input in seconds (as uint64) and returns a human-readable time string.
func FormatTime(seconds uint64) string {
	switch {
	case seconds >= 86400:
		days := seconds / 86400
		return fmt.Sprintf("%d day%s", days, plural(int(days)))
	case seconds >= 3600:
		hours := seconds / 3600
		return fmt.Sprintf("%d hr%s", hours, plural(int(hours)))
	case seconds >= 60:
		minutes := seconds / 60
		return fmt.Sprintf("%d min%s", minutes, plural(int(minutes)))
	default:
		return fmt.Sprintf("%d sec%s", seconds, plural(int(seconds)))
	}
}

// FormatSize takes an input in bytes and returns a human-readable size string.
func FormatSize(bytes int) string {
	const (
		kb = 1024
		mb = 1024 * kb
		gb = 1024 * mb
		tb = 1024 * gb
	)

	switch {
	case bytes >= tb:
		return fmt.Sprintf("%.1fTB", float64(bytes)/float64(tb))
	case bytes >= gb:
		return fmt.Sprintf("%.1fGB", float64(bytes)/float64(gb))
	case bytes >= mb:
		return fmt.Sprintf("%.1fMB", float64(bytes)/float64(mb))
	case bytes >= kb:
		return fmt.Sprintf("%.1fKB", float64(bytes)/float64(kb))
	default:
		return fmt.Sprintf("%dB", bytes)
	}
}

// FormatSpeed takes an input in bytes per second and returns a human-readable speed string.
func FormatSpeed(bytesPerSecond int) string {
	const (
		kb = 1024
		mb = 1024 * kb
		gb = 1024 * mb
		tb = 1024 * gb
	)

	switch {
	case bytesPerSecond >= tb:
		return fmt.Sprintf("%.1fTB/s", float64(bytesPerSecond)/float64(tb))
	case bytesPerSecond >= gb:
		return fmt.Sprintf("%.1fGB/s", float64(bytesPerSecond)/float64(gb))
	case bytesPerSecond >= mb:
		return fmt.Sprintf("%.1fMB/s", float64(bytesPerSecond)/float64(mb))
	case bytesPerSecond >= kb:
		return fmt.Sprintf("%.1fKB/s", float64(bytesPerSecond)/float64(kb))
	default:
		return fmt.Sprintf("%dB/s", bytesPerSecond)
	}
}

// plural returns "s" if the value is not 1, otherwise an empty string.
func plural(value int) string {
	if value == 1 {
		return ""
	}
	return "s"
}
