package leetcode

// Problem is a solved LeetCode problem with a sample input.
type Problem struct {
	Number int
	Title  string
	Input  string     // sample input, as shown to the user
	Run    func() any // runs the solution on the sample input
}

// problems is sorted by Number. Each Run builds its input on every call, so a
// solution that writes into its arguments cannot spoil the next run.
var problems = []Problem{
	{
		Number: 25,
		Title:  "Reverse Nodes in k-Group",
		Input:  "head = [1,2,3,4,5], k = 2",
		Run: func() any {
			return listVals(reverseKGroup(buildList([]int{1, 2, 3, 4, 5}), 2))
		},
	},
	{
		Number: 42,
		Title:  "Trapping Rain Water",
		Input:  "height = [0,1,0,2,1,0,1,3,2,1,2,1]",
		Run: func() any {
			return trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
		},
	},
	{
		Number: 44,
		Title:  "Wildcard Matching",
		Input:  `s = "adceb", p = "*a*b"`,
		Run: func() any {
			return isMatch("adceb", "*a*b")
		},
	},
	{
		Number: 46,
		Title:  "Permutations",
		Input:  "nums = [1,2,3]",
		Run: func() any {
			return permute([]int{1, 2, 3})
		},
	},
	{
		Number: 51,
		Title:  "N-Queens",
		Input:  "n = 4",
		Run: func() any {
			return solveNQueens(4)
		},
	},
	{
		Number: 127,
		Title:  "Word Ladder",
		Input:  `beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]`,
		Run: func() any {
			return ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
		},
	},
	{
		Number: 200,
		Title:  "Number of Islands",
		Input:  `grid = [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]`,
		Run: func() any {
			return numIslands([][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			})
		},
	},
	{
		Number: 940,
		Title:  "Distinct Subsequences II",
		Input:  `s = "abc"`,
		Run: func() any {
			return distinctSubseqII("abc")
		},
	},
	{
		Number: 1406,
		Title:  "Stone Game III",
		Input:  "stoneValue = [1,2,3,7]",
		Run: func() any {
			return stoneGameIII([]int{1, 2, 3, 7})
		},
	},
}

// Problems returns the solved problems, sorted by number.
func Problems() []Problem {
	out := make([]Problem, len(problems))
	copy(out, problems)
	return out
}

// Lookup returns the problem with the given number.
func Lookup(number int) (Problem, bool) {
	for _, p := range problems {
		if p.Number == number {
			return p, true
		}
	}
	return Problem{}, false
}
