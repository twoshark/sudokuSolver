package stringBoardParser

import (
	"errors"
	"strconv"
	"strings"
)

func ParseInput(input string) ([][]int, error) {
	// parse rows and check count
	rows := strings.Split(input, "\n")
	if len(rows) != 9 {
		return nil, errors.New("bad input, too many rows")
	}

	//allocate array
	arr := make([][]int, 9)
	for row, rowArr := range rows {
		rowItems := strings.Split(rowArr, ",")
		if len(rowItems) > 9 {
			return nil, errors.New("too many columns")
		}
		arr[row] = make([]int, 9)
		for col, colVal := range rowItems {
			intItem, err := strconv.Atoi(colVal)
			if err != nil {
				panic(err)
			}
			arr[row][col] = intItem
		}
	}
	return arr, nil
}
