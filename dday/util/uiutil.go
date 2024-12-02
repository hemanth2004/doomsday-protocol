package util

import "strconv"

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

func DeleteElement[T any](slice []T, index int) []T {
	// Handle invalid index cases
	if index < 0 || index >= len(slice) {
		return slice // Return the original slice if index is out of range
	}
	return append(slice[:index], slice[index+1:]...)
}

func DeleteColumn[T any](matrix [][]T, colIndex int) [][]T {
	// Handle edge cases: empty matrix or invalid column index
	if len(matrix) == 0 || colIndex < 0 {
		return matrix
	}

	for i, row := range matrix {
		if colIndex < len(row) { // Check if the column index is valid for this row
			matrix[i] = append(row[:colIndex], row[colIndex+1:]...)
		}
	}

	return matrix
}

func GetColumn[T any](matrix [][]T, colIndex int) []T {
	var column []T

	for _, row := range matrix {
		if colIndex < len(row) { // Ensure the column index is valid for this row
			column = append(column, row[colIndex])
		}
	}

	return column
}

func SetColumn[T any](matrix [][]T, colIndex int, value T) {
	for i, row := range matrix {
		if colIndex < len(row) { // Ensure the column index is valid for this row
			matrix[i][colIndex] = value
		}
	}
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

func Sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

// Ternary Operator
func IfElse[T any](cond bool, exp1, exp2 T) T {
	if cond {
		return exp1
	} else {
		return exp2
	}
}

func MarginHor(s string, amt int) string {
	for i := 0; i < amt; i++ {
		s = " " + s + " "
	}
	return s
}
