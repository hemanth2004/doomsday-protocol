package tableutils

import (
	"fmt"
	"math"
	"unicode/utf8"

	"github.com/evertras/bubble-table/table"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
)

// bubble-table utils

func UpdateTableHeightAndFooter(m table.Model, rowsString [][]string, columns []table.Column, width, viewportHeight int) table.Model {
	linesUsedByEachTableRow := CalculateExtraMultilineRows(columns, rowsString, width-2)
	pageSize := max(1, CalculatePaginationSize(linesUsedByEachTableRow, len(rowsString), viewportHeight))

	startIndex, endIndex := m.VisibleIndices()
	visibleRowsHeight := 0
	for i := startIndex; i < endIndex; i++ {
		visibleRowsHeight += linesUsedByEachTableRow[i]
	}
	unusedHeight := viewportHeight - visibleRowsHeight
	customFooter := fmt.Sprintf("%d / %d ", m.CurrentPage(), m.MaxPages())
	for i := 0; i < unusedHeight-2; i++ {
		customFooter += ""
	}

	return m.
		WithTargetWidth(width).
		WithPageSize(pageSize).
		WithStaticFooter(customFooter)
}

func GetColumnFromKey(columns []table.Column, key string) table.Column {
	for _, col := range columns {
		if col.Key() == key {
			return col
		}
	}
	return table.NewColumn("nil,", "nil", 1)
}

func CalculateColumnWidth(columns []table.Column, totalWidth int, targetColumn table.Column) int {

	fixedTotal := 0
	for _, pair := range columns {
		if !pair.IsFlex() {
			fixedTotal += pair.Width()
		}
	}

	remainingWidth := totalWidth - fixedTotal

	totalFlexFactors := 0
	for _, factor := range columns {
		if factor.IsFlex() {
			totalFlexFactors += factor.FlexFactor()
		}
	}

	totalFlexFactorsWithoutInputFactor := totalFlexFactors - targetColumn.FlexFactor()

	return int(float64(remainingWidth) - ((float64(totalFlexFactorsWithoutInputFactor) / float64(totalFlexFactors)) * float64(remainingWidth)))
	// the -1 is accounting for the table border
}

func CalculateExtraMultilineRows(columns []table.Column, rows [][]string, totalWidth int) []int {
	var allowedWidths []int
	for _, col := range columns {
		if col.IsFlex() {
			allowedWidths = append(allowedWidths, CalculateColumnWidth(columns, totalWidth, col))
		} else {
			allowedWidths = append(allowedWidths, col.Width())
		}
	}

	util.SetColumn(rows, 2, " ")
	extraInEachCell := make([][]int, len(rows))
	for i := range extraInEachCell {
		extraInEachCell[i] = make([]int, len(allowedWidths))
	}

	for i, row := range rows {
		for j, colWidth := range allowedWidths {
			cellContent := row[j]
			lines := int(math.Ceil(float64(utf8.RuneCountInString(cellContent)) / float64(colWidth)))
			extraInEachCell[i][j] = lines
		}
	}

	//debug.Log("Extra In Each Cell:\n" + printMatrix(extraInEachCell))

	maxInEachRow := make([]int, len(rows))
	for i := range rows {
		var maximumExtra int = 0
		for _, extra := range extraInEachCell[i] {
			if extra > maximumExtra {
				maximumExtra = extra
			}
		}
		maxInEachRow[i] = maximumExtra
	}

	return maxInEachRow
}

// Problem of optimization
// Starting info: given in arguements
// Conditions: maximising the page size
// Solution:
//  1. Iterate from pagination size 1 till number of rows
//  2. Example case
//     - total rows = 5 rows
//     - lines used by each row = {1, 2, 3, 2, 1}
//     - allowed viewport = 6 lines
//  Then the page size would be 3 because rows 1, 2 and 3 will be 6, and row 4 and 5 will be 3

func CalculatePaginationSize(linesUsedByEachTableRow []int, totalRows, viewportHeight int) (size int) {
	for i := 1; i <= totalRows; i++ { // i is the potential page size
		currentHeight := 0
		valid := true

		// Check if the current page size (i) fits within the viewport height
		for j := 0; j < i; j++ {
			currentHeight += linesUsedByEachTableRow[j]
			if currentHeight > viewportHeight {
				valid = false
				break
			}
		}

		// If valid, update the size; otherwise, break the loop as no larger page size can work
		if valid {
			size = i
		} else {
			break
		}
	}

	//fmt.Print("Lines In Each Row:", linesUsedByEachTableRow, "\n\r", "Size: ", size, "\n\r")
	return
}
