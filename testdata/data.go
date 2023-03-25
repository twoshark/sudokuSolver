package testdata

type TestData struct {
	Description string
	Input       string
	Output      [][]int
}

var Cases = []TestData{
	{
		Description: "Example From Challenge",
		Input: `0,0,0,2,6,0,7,0,1
6,8,0,0,7,0,0,9,0
1,9,0,0,0,4,5,0,0
8,2,0,1,0,0,0,4,0
0,0,4,6,0,2,9,0,0
0,5,0,0,0,3,0,2,8
0,0,9,3,0,0,0,7,4
0,4,0,0,5,0,0,3,6
7,0,3,0,1,8,0,0,0`,
	},
	{
		Description: "Easy From Web",
		Input: `7,6,0,1,0,9,0,3,8
0,9,0,3,7,4,0,6,2
1,9,0,0,0,4,5,0,0
0,3,0,0,0,0,4,0,0
0,2,7,0,5,0,9,0,3
9,0,5,4,0,0,6,0,0
8,0,1,2,0,6,0,0,0
0,0,0,0,0,1,7,2,6
0,0,0,0,7,4,0,9,0`,
	},
}
