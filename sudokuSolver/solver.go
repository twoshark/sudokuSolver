package sudokuSolver

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

type SudokuSolver struct {
	board      [][]int
	candidates [][]*Candidates
	boardLock  sync.Mutex
}

func NewSudokuSolver(board [][]int) *SudokuSolver {
	solver := new(SudokuSolver)
	solver.board = board
	solver.candidates = make([][]*Candidates, 9)
	for i := range solver.candidates {
		solver.candidates[i] = make([]*Candidates, 9)
		for j := range solver.candidates[i] {
			solver.candidates[i][j] = NewCandidates()
		}
	}
	solver.boardLock = sync.Mutex{}
	return solver
}

func (s *SudokuSolver) GetBoard() [][]int {
	s.boardLock.Lock()
	defer s.boardLock.Unlock()
	return s.board
}

func (s *SudokuSolver) UpdateBoard(x, y int, value int) {
	s.boardLock.Lock()
	defer s.boardLock.Unlock()
	s.board[x][y] = value
}

func (s *SudokuSolver) Solve() [][]int {
	log.Print("Original Board:")
	s.PrintBoardtoLogs()
	og := make([][]int, 9)
	for i := range s.board {
		og[i] = make([]int, len(s.board[i]))
		copy(og[i], s.board[i])
	}
	// process candidates and 'easy wins'
	var solvedCellCount uint
	var passCount uint = 1
	var changed bool
	for {
		changed = false
		solvedCellCount = 0
		for x, row := range s.board {
			for y, value := range row {
				if value != 0 {
					solvedCellCount++
					continue
				}

				candidates := s.CellCandidates(x, y)
				switch len(candidates) {
				case 1:
					log.WithFields(log.Fields{"candidates": candidates, "x": x, "y": y}).
						Print("Candidate Selected for Cell")
					s.UpdateBoard(x, y, candidates[0])
					solvedCellCount++
					changed = true
				case 0:
					log.WithFields(log.Fields{"candidates": candidates, "x": x, "y": y}).
						Panic("Solver Error, No Candidates for Cell")
				}
			}
		}
		log.Print("Pass ", passCount, " Completed, Solved Cells: ", solvedCellCount)
		if solvedCellCount == 81 {
			log.Print("Solved Board!")
			s.PrintBoardsSideBySide(og)
			break
		}
		if !changed {
			log.Error("Full confidence values exhausted before solve")
			s.PrintBoardtoLogs()
			break
		}
		passCount++
	}

	return s.board
}
