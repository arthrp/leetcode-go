package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/arthrp/leetcode-go/internal/leetcode"
)

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		for _, p := range leetcode.Problems() {
			fmt.Printf("%d - %s\n", p.Number, p.Title)
		}
	case 1:
		number, err := strconv.Atoi(args[0])
		if err != nil {
			exitf("%q is not a problem number", args[0])
		}
		p, ok := leetcode.Lookup(number)
		if !ok {
			exitf("no solution for problem %d", number)
		}
		fmt.Printf("%d - %s\nInput: %s\nOutput: %v\n", p.Number, p.Title, p.Input, p.Run())
	default:
		exitf("too many arguments")
	}
}

func exitf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	fmt.Fprintln(os.Stderr, "usage: leetcode [problem number]")
	os.Exit(1)
}
