package util

import "strconv"

type ResizeHandler interface {
	Resize(width int, height int)
}

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

func PrintMatrix(matrix [][]int) (s string) {
	for _, row := range matrix { // Loop through each row
		for _, col := range row { // Loop through each column in the row
			s += strconv.Itoa(col) + "  " // Print each element with a space
		}
		s += "\n"
	}
	return
}

func MarginHor(s string, amt int) string {
	for i := 0; i < amt; i++ {
		s = " " + s + " "
	}
	return s
}

func Repl(s string, amt int) string {
	dup := s
	for i := 0; i < amt; i++ {
		s = s + dup
	}
	return s
}
