package dday

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

func RenderDownloads(g *gocui.Gui) error {
	v, err := g.View("downloads")
	if err != nil {
		return err
	}

	downloads := Resources
	v.Clear()

	// Render each download
	for i, download := range downloads {
		// Line 1: Index + Name
		fmt.Fprintf(v, "%d. %s\n", i+1, download.Name)

		// Line 2: Status
		fmt.Fprintf(v, "Status: %s\n", download.Status)

		// Line 3: Progress Bar with % Indicator
		progress := download.Info.ProgressPercent()
		progressBar := generateProgressBar(progress, 30) // 30 characters wide
		fmt.Fprintf(v, "[%-30s] %.3f%%\n", progressBar, progress)

		// Line 4: Other Download Information
		done, size := formatSize(download.Info.Done), formatSize(download.Info.Size)
		speed := formatSpeed(download.Info.Bandwidth)
		eta := formatETA(download.Info.ETA)
		fmt.Fprintf(v, "Downloaded: %s / %s | Speed: %s | ETA: %s\n", done, size, speed, eta)

		// Padding: Add 2 blank lines before the next block
		if i < len(downloads)-1 {
			fmt.Fprintln(v)
			fmt.Fprintln(v)
		}
	}
	return nil
}

// generateProgressBar creates a simple progress bar string.
func generateProgressBar(progress float64, width int) string {
	filled := int(progress / 100 * float64(width))
	return strings.Repeat("=", filled) + strings.Repeat(" ", width-filled)
}

// formatSize converts size to a human-readable format (MB or GB).
func formatSize(size float64) string {
	if size >= 1e9 {
		return fmt.Sprintf("%.2f GB", size/1e9)
	}
	return fmt.Sprintf("%.2f MB", size/1e6)
}

// formatSpeed converts bandwidth to a human-readable format (MB/s or GB/s).
func formatSpeed(speed float64) string {
	if speed >= 1e9 {
		return fmt.Sprintf("%.2f GB/s", speed/1e9)
	}
	return fmt.Sprintf("%.2f MB/s", speed/1e6)
}

// formatETA converts ETA seconds into HH:MM:SS format.
func formatETA(etaStr string) string {
	etaSeconds, err := parseSeconds(etaStr)
	if err != nil {
		return "Unknown"
	}

	hours := etaSeconds / 3600
	minutes := (etaSeconds % 3600) / 60
	seconds := etaSeconds % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// parseSeconds parses ETA string in seconds.
func parseSeconds(etaStr string) (int, error) {
	var etaSeconds float64
	_, err := fmt.Sscanf(etaStr, "%f seconds", &etaSeconds)
	if err != nil {
		return 0, err
	}
	return int(etaSeconds), nil
}
