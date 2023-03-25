package stringBoardParser

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golangSudokuSolver/testdata"
	"testing"
)

type ParseTestSuite struct {
	suite.Suite
}

type ParseTestCase struct {
	Description string
	Input       string
	Output      [][]int
}

func (suite *ParseTestSuite) SetupTest()    {}
func (suite *ParseTestSuite) TearDownTest() {}

func (suite *ParseTestSuite) TestParse() {
	cases := []ParseTestCase{
		{
			Description: testdata.Cases[0].Description,
			Input:       testdata.Cases[0].Input,
			Output: [][]int{
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
			Description: testdata.Cases[1].Description,
			Input:       testdata.Cases[1].Input,
			Output: [][]int{
				{7, 6, 0, 1, 0, 9, 0, 3, 8},
				{0, 9, 0, 3, 7, 4, 0, 6, 2},
				{1, 9, 0, 0, 0, 4, 5, 0, 0},
				{0, 3, 0, 0, 0, 0, 4, 0, 0},
				{0, 2, 7, 0, 5, 0, 9, 0, 3},
				{9, 0, 5, 4, 0, 0, 6, 0, 0},
				{8, 0, 1, 2, 0, 6, 0, 0, 0},
				{0, 0, 0, 0, 0, 1, 7, 2, 6},
				{0, 0, 0, 0, 7, 4, 0, 9, 0},
			},
		},
	}

	for _, testCase := range cases {
		output, err := ParseInput(testCase.Input)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), testCase.Output, output)
	}
}

func TestParserTestSuite(t *testing.T) {
	suite.Run(t, new(ParseTestSuite))
}
