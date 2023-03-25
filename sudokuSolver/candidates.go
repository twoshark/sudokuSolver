package sudokuSolver

import (
	"sync"
)

type Candidates struct {
	values []int
	lock   sync.Mutex
}

func NewCandidates(candidates ...int) *Candidates {
	c := new(Candidates)
	c.values = make([]int, 0, 9)
	for _, candidate := range candidates {
		c.values = append(c.values, candidate)
	}
	c.lock = sync.Mutex{}
	return c
}

func (c *Candidates) UpdateCandidates(candidates []int) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.values = candidates
}

func (s *SudokuSolver) CellCandidates(x, y int) []int {
	taken := s.gatherTakenValues(x, y)
	candidates := make([]int, 0, 9)
	for key := range taken {
		if key != 0 && !taken[key] {
			candidates = append(candidates, key)
		}
	}
	s.candidates[x][y].UpdateCandidates(candidates)
	return candidates
}

func (s *SudokuSolver) gatherTakenValues(x int, y int) map[int]bool {
	claimInt := make(chan int, 27)
	defer close(claimInt)

	taken := make(map[int]bool)
	for i := 0; i <= 9; i++ {
		taken[i] = false
	}

	go s.RowCandidates(x, claimInt)
	go s.ColumnCandidates(y, claimInt)
	go s.BoxCandidates(x/3, y/3, claimInt)

	for i := 0; i < 27; i++ {
		select {
		case claim := <-claimInt:
			taken[claim] = true
		}
	}
	return taken
}

func (s *SudokuSolver) RowCandidates(x int, claimInt chan int) {
	row := s.GetBoard()[x]
	rowTaken := make([]int, 0, 9)
	for i := 0; i < len(row); i++ {
		claimInt <- row[i]
		if row[i] != 0 {
			rowTaken = append(rowTaken, row[i])
		}
	}
}

func (s *SudokuSolver) ColumnCandidates(y int, claimInt chan int) {
	board := s.GetBoard()
	colTaken := make([]int, 0, 9)
	for _, row := range board {
		claimInt <- row[y]
		if row[y] != 0 {
			colTaken = append(colTaken, row[y])
		}
	}
}

func (s *SudokuSolver) BoxCandidates(boxGridX, boxGridY int, claimInt chan int) {
	board := s.GetBoard()
	boxTaken := make([]int, 0, 9)
	for i := 3 * boxGridX; i < 3*boxGridX+3; i++ {
		for j := 3 * boxGridY; j < 3*boxGridY+3; j++ {
			claimInt <- board[i][j]
			if board[i][j] != 0 {
				boxTaken = append(boxTaken, board[i][j])
			}
		}
	}
}
