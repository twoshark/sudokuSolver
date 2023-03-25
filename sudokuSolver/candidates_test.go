package sudokuSolver

import (
	"sort"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type CandidateTestCase struct {
	Description string
	Input       CandidateInput
	Expected    []int
}

type CandidateInput struct {
	BoardId int
	X       int
	Y       int
}

type Coordinates struct {
}

func (suite *SolverTestSuite) TestGetCandidates() {
	cases := []CandidateTestCase{
		{
			Description: "Board 0: 0,0",
			Input: CandidateInput{
				BoardId: 0,
				X:       0,
				Y:       0,
			},
			Expected: []int{3, 4, 5},
		},
		{
			Description: "Board 0: 0,5",
			Input: CandidateInput{
				BoardId: 0,
				X:       0,
				Y:       5,
			},
			Expected: []int{5, 9},
		},
	}

	for _, testCase := range cases {
		log.Print("Test Case: ", testCase.Description)
		candidates := suite.Solvers[testCase.Input.BoardId].CellCandidates(testCase.Input.X, testCase.Input.Y)
		sort.Ints(candidates)
		assert.Equal(suite.T(), testCase.Expected, candidates)
	}
}
