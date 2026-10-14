# leetcode-go

Go solutions to LeetCode problems. They live in `internal/leetcode`.

Sample inputs live in `inputs.txt`. To add a problem, write the solution and its test, add a `Number    Title    Input` row to `inputs.txt`, and add an entry to `solvers` in `internal/leetcode/registry.go`.

```
go test ./...
go run ./cmd/leetcode
go run ./cmd/leetcode 51
```

With no arguments, `leetcode` lists the solved problems. With a problem number, it runs that solution on its sample input.
