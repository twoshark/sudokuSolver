package sudokuSolver

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
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

func (s *SudokuSolver) PrintBoardtoLogs() {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if s.board[row][column] == 0 {
				color.New(color.FgRed).Print("0 ")
			} else {
				fmt.Print(s.board[row][column], " ")
			}
		}
		fmt.Print("\n")
	}
}

func (s *SudokuSolver) PrintBoardsSideBySide(original [][]int) {
	fmt.Print("----------Original----------||-----------Solved-----------\n")
	for row := 0; row < 9; row++ {
		if row > 0 && row%3 == 0 {
			fmt.Print("----------------------------||----------------------------\n")
		}
		for column := 0; column < 9; column++ {
			if column%3 == 0 {
				fmt.Print(" | ")
			}
			printCellValue(original, row, column)
		}
		fmt.Print(" >> ")
		for column := 0; column < 9; column++ {
			if column%3 == 0 {
				fmt.Print(" | ")
			}
			if s.board[row][column] == 0 {
				color.New(color.FgRed).Print("0 ")
			} else {
				fmt.Print(s.board[row][column], " ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("----------------------------||----------------------------\n")
}

func printCellValue(arr [][]int, row int, column int) {
	if arr[row][column] == 0 {
		color.New(color.FgRed).Print("0 ")
	} else {
		fmt.Print(arr[row][column], " ")
	}
}
