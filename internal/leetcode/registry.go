package leetcode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/arthrp/leetcode-go"
)

// Problem is a solved LeetCode problem with a sample input.
type Problem struct {
	Number int
	Title  string
	Input  string              // sample input, as shown to the user
	Run    func() (any, error) // runs the solution on the sample input
}

// solvers maps a problem number to the call that runs its solution. Each
// function reads its arguments from the Input column of inputs.txt.
var solvers = map[int]func(a args) any{
	25: func(a args) any {
		return listVals(reverseKGroup(buildList(a.intsArg("head")), a.intArg("k")))
	},
	41: func(a args) any { return firstMissingPositive(a.intsArg("nums")) },
	42: func(a args) any { return trap(a.intsArg("height")) },
	44: func(a args) any { return isMatch(a.strArg("s"), a.strArg("p")) },
	46: func(a args) any { return permute(a.intsArg("nums")) },
	51: func(a args) any { return solveNQueens(a.intArg("n")) },
	60: func(a args) any { return fmt.Sprintf("%q", getPermutation(a.intArg("n"), a.intArg("k"))) },
	65: func(a args) any { return isNumber(a.strArg("s")) },
	68: func(a args) any {
		return fmt.Sprintf("%q", fullJustify(a.strsArg("words"), a.intArg("maxWidth")))
	},
	76: func(a args) any { return fmt.Sprintf("%q", minWindow(a.strArg("s"), a.strArg("t"))) },
	84: func(a args) any { return largestRectangleArea(a.intsArg("heights")) },
	127: func(a args) any {
		return ladderLength(a.strArg("beginWord"), a.strArg("endWord"), a.strsArg("wordList"))
	},
	200:  func(a args) any { return numIslands(a.byteGridArg("grid")) },
	940:  func(a args) any { return distinctSubseqII(a.strArg("s")) },
	1406: func(a args) any { return stoneGameIII(a.intsArg("stoneValue")) },
}

// problems is sorted by Number. Each Run parses its input on every call, so a
// solution that writes into its arguments cannot spoil the next run.
var problems = mustLoadProblems(leetcodego.Inputs, solvers)

// loadProblems reads the Number, Title, and Input columns. Columns are
// separated by a tab or by two or more spaces. Every row needs a solver and
// every solver needs a row.
func loadProblems(text string, solvers map[int]func(args) any) ([]Problem, error) {
	seen := make(map[int]int)
	var out []Problem
	for i, line := range strings.Split(text, "\n") {
		lineNum := i + 1
		line = strings.TrimRight(line, "\r")
		if strings.TrimSpace(line) == "" {
			continue
		}
		cols := splitColumns(line)
		if len(cols) != 3 {
			return nil, fmt.Errorf("inputs.txt:%d: want 3 columns, got %d", lineNum, len(cols))
		}
		if cols[0] == "Number" {
			continue
		}
		number, err := strconv.Atoi(cols[0])
		if err != nil {
			return nil, fmt.Errorf("inputs.txt:%d: bad problem number %q", lineNum, cols[0])
		}
		if prev, ok := seen[number]; ok {
			return nil, fmt.Errorf("inputs.txt:%d: duplicate problem number %d (first on line %d)", lineNum, number, prev)
		}
		seen[number] = lineNum
		solver, ok := solvers[number]
		if !ok {
			return nil, fmt.Errorf("inputs.txt:%d: no solver for problem %d", lineNum, number)
		}
		if cols[1] == "" || cols[2] == "" {
			return nil, fmt.Errorf("inputs.txt:%d: empty title or input", lineNum)
		}
		input, solve := cols[2], solver
		out = append(out, Problem{
			Number: number,
			Title:  cols[1],
			Input:  input,
			Run: func() (any, error) {
				return runSolver(input, solve)
			},
		})
	}
	for number := range solvers {
		if _, ok := seen[number]; !ok {
			return nil, fmt.Errorf("solver %d has no row in inputs.txt", number)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Number < out[j].Number })
	return out, nil
}

func mustLoadProblems(text string, solvers map[int]func(args) any) []Problem {
	loaded, err := loadProblems(text, solvers)
	if err != nil {
		panic(err)
	}
	return loaded
}

// splitColumns splits a row into at most three columns. A column break is a
// tab or a run of two or more spaces, so single spaces inside a title or an
// input stay put.
func splitColumns(line string) []string {
	var cols []string
	start := 0
	for i := 0; i < len(line) && len(cols) < 2; {
		tab := line[i] == '\t'
		spaces := line[i] == ' ' && i+1 < len(line) && line[i+1] == ' '
		if !tab && !spaces {
			i++
			continue
		}
		cols = append(cols, strings.TrimSpace(line[start:i]))
		for i < len(line) && (line[i] == ' ' || line[i] == '\t') {
			i++
		}
		start = i
	}
	cols = append(cols, strings.TrimSpace(line[start:]))
	return cols
}

// runSolver parses input and calls solve. An argError from a mistyped
// argument comes back as an error; any other panic propagates.
func runSolver(input string, solve func(args) any) (result any, err error) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		if ae, ok := r.(*argError); ok {
			err = ae
			return
		}
		panic(r)
	}()
	parsed, err := parseArgs(input)
	if err != nil {
		return nil, fmt.Errorf("parse input: %w", err)
	}
	return solve(parsed), nil
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
