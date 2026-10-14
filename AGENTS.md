# Project layout

- Solutions live in `internal/leetcode`, one file per problem plus a `_test.go`.
- Sample inputs live in `inputs.txt` at the repo root. It has a header row, then one `Number    Title    Input` row per problem, with columns separated by 4 spaces. Inputs use LeetCode syntax, e.g. `head = [1,2,3,4,5], k = 2`. The root file `inputs.go` embeds it at build time.
- `internal/leetcode/registry.go` maps each problem number to a solver in `solvers`. A solver reads its arguments with the accessors in `args.go` (`intArg`, `strArg`, `intsArg`, `strsArg`, `byteGridArg`, `treeArg`).
- To add a problem, write the solution and its test, add a row to `inputs.txt`, and add a `solvers` entry. Every row needs a solver and every solver needs a row, or the package panics at init.
- Check changes with `go test ./...` and `go run ./cmd/leetcode <number>`.

# Comments

- Keep the comment above a solution function short: the problem number and title, 1-2 lines on what the function does, and its time and extra-memory complexity in big O notation.
- Explain the approach in comments next to the code lines they describe, not in the top comment.
