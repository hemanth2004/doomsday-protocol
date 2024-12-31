package util

func Sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func TrimString(str string, maxLength int) string {
	// Handle invalid maxLength
	if maxLength < 0 {
		return str
	}
	if len(str) > maxLength {
		return str[:maxLength]
	}
	return str
}

// Ternary Operator
func IfElse[T any](cond bool, exp1, exp2 T) T {
	if cond {
		return exp1
	} else {
		return exp2
	}
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
