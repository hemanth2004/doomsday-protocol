package util

// Draw a line
func DrawLine(width int) string {
	var s string = ""
	for i := 0; i < width; i++ {
		s += "─"
	}
	return s
}

// CalculateFlexWidth calculates the width of a flex column in a table.
func CalculateFlexWidth(totalWidth int, fixedWidths []int, flexFactors []int, inputFactor int) int {
	// Calculate total fixed width
	fixedTotal := 0
	for _, width := range fixedWidths {
		fixedTotal += width
	}

	// Remaining width after accounting for fixed widths
	remainingWidth := totalWidth - fixedTotal

	// Add the input factor to the list of flex factors
	totalFlexFactors := 0
	for _, factor := range flexFactors {
		totalFlexFactors += factor
	}
	totalFlexFactors += inputFactor

	// Calculate and return the width of the input flex column
	return (remainingWidth * inputFactor) / totalFlexFactors
}
