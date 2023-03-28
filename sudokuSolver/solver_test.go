package sudokuSolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SolverTestSuite struct {
	suite.Suite
	Solvers []*SudokuSolver
}

type SolverTestData struct {
	Description string
	Input       string
	Board       [][]int
	Candidates  struct{}
}

var Cases = []SolverTestData{
	{
		Description: "Example from Peer",
		Input: `0,0,0,2,6,0,7,0,1
6,8,0,0,7,0,0,9,0
1,9,0,0,0,4,5,0,0
8,2,0,1,0,0,0,4,0
0,0,4,6,0,2,9,0,0
0,5,0,0,0,3,0,2,8
0,0,9,3,0,0,0,7,4
0,4,0,0,5,0,0,3,6
7,0,3,0,1,8,0,0,0`,
		Board: [][]int{
			{0, 0, 0, 2, 6, 0, 7, 0, 1},
			{6, 8, 0, 0, 7, 0, 0, 9, 0},
			{1, 9, 0, 0, 0, 4, 5, 0, 0},
			{8, 2, 0, 1, 0, 0, 0, 4, 0},
			{0, 0, 4, 6, 0, 2, 9, 0, 0},
			{0, 5, 0, 0, 0, 3, 0, 2, 8},
			{0, 0, 9, 3, 0, 0, 0, 7, 4},
			{0, 4, 0, 0, 5, 0, 0, 3, 6},
			{7, 0, 3, 0, 1, 8, 0, 0, 0},
		},
	},
	{
		Description: "Easy From Web",
		Input: `3,8,6,0,0,4,7,0,0
0,0,9,0,0,0,2,0,0
0,2,0,1,0,3,8,0,5
0,7,0,0,3,0,6,2,0
0,5,2,0,0,1,0,0,4
9,4,0,2,7,0,0,0,0
2,3,0,7,4,9,5,8,6
8,0,0,0,1,0,4,0,0
4,0,0,0,0,0,0,0,2`,
		Board: [][]int{
			{3, 8, 6, 0, 0, 4, 7, 0, 0},
			{0, 0, 9, 0, 0, 0, 2, 0, 0},
			{0, 2, 0, 1, 0, 3, 8, 0, 5},
			{0, 7, 0, 0, 3, 0, 6, 2, 0},
			{0, 5, 2, 0, 0, 1, 0, 0, 4},
			{9, 4, 0, 2, 7, 0, 0, 0, 0},
			{2, 3, 0, 7, 4, 9, 5, 8, 6},
			{8, 0, 0, 0, 1, 0, 4, 0, 0},
			{4, 0, 0, 0, 0, 0, 0, 0, 2},
		},
	},
	{
		Description: "Medium From Web",
		Input: `9,0,0,0,6,0,7,0,4
0,0,3,0,0,9,0,0,0
4,0,5,0,0,0,6,0,0
0,0,7,8,3,0,9,0,0
3,0,0,7,4,0,0,0,0
2,0,4,0,0,6,5,7,0
6,0,0,0,0,4,0,5,0
0,0,0,2,0,7,0,6,0
0,0,9,0,0,0,0,8,7`,
		Board: [][]int{
			{9, 0, 0, 0, 6, 0, 7, 0, 4},
			{0, 0, 3, 0, 0, 9, 0, 0, 0},
			{4, 0, 5, 0, 0, 0, 6, 0, 0},
			{0, 0, 7, 8, 3, 0, 9, 0, 0},
			{3, 0, 0, 7, 4, 0, 0, 0, 0},
			{2, 0, 4, 0, 0, 6, 5, 7, 0},
			{6, 0, 0, 0, 0, 4, 0, 5, 0},
			{0, 0, 0, 2, 0, 7, 0, 6, 0},
			{0, 0, 9, 0, 0, 0, 0, 8, 7},
		},
	},
}

func TestSolverTestSuite(t *testing.T) {
	suite.Run(t, new(SolverTestSuite))
}

func (suite *SolverTestSuite) SetupTest() {
	suite.Solvers = make([]*SudokuSolver, len(Cases))
	for i, testCase := range Cases {
		suite.Solvers[i] = NewSudokuSolver(testCase.Board)
	}
}
func (suite *SolverTestSuite) TearDownTest() {}

func (suite *SolverTestSuite) TestParse() {
	for _, testCase := range Cases {
		output, err := ParseInput(testCase.Input)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), testCase.Board, output)
	}
}

func (suite *SolverTestSuite) TestNewSudokuSolver() {
	for i, testCase := range Cases {
		assert.NotNil(suite.T(), suite.Solvers[i])
		assert.Equal(suite.T(), testCase.Board, suite.Solvers[i].board)
	}
}

func (suite *SolverTestSuite) TestSolve() {
	for _, testCase := range Cases {
		board, err := ParseInput(testCase.Input)
		if err != nil {
			panic(err)
		}
		assert.Equal(suite.T(), board, testCase.Board)
		solver := NewSudokuSolver(board)
		_ = solver.Solve()
	}
}
