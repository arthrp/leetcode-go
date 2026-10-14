package leetcode

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseArgs(t *testing.T) {
	got, err := parseArgs(`head = [1,2,3,4,5], k = 2`)
	if err != nil {
		t.Fatal(err)
	}
	want := args{"head": []any{1, 2, 3, 4, 5}, "k": 2}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("parseArgs = %#v, want %#v", got, want)
	}
}

func TestParseArgsValues(t *testing.T) {
	tests := []struct {
		in   string
		want args
	}{
		{`n = -3`, args{"n": -3}},
		{`s = "adceb", p = "*a*b"`, args{"s": "adceb", "p": "*a*b"}},
		{`s = "a\"b"`, args{"s": `a"b`}},
		{`nums = [1, 2, 3]`, args{"nums": []any{1, 2, 3}}},
		{`grid = [["1","0"],["0","1"]]`, args{"grid": []any{[]any{"1", "0"}, []any{"0", "1"}}}},
		{`empty = []`, args{"empty": []any{}}},
		{"", args{}},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := parseArgs(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("parseArgs = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestParseArgsErrors(t *testing.T) {
	for _, in := range []string{
		`head [1,2]`,
		`s = "unterminated`,
		`nums = [1,2`,
		`n = 4 leftover`,
		`n = 4,`,
		`= 4`,
		`n =`,
	} {
		t.Run(in, func(t *testing.T) {
			if _, err := parseArgs(in); err == nil {
				t.Fatal("parseArgs returned no error")
			}
		})
	}
}

func TestArgAccessors(t *testing.T) {
	a, err := parseArgs(`n = 4, s = "abc", nums = [1,2], words = ["hot","dot"], grid = [["1","0"],["0","1"]]`)
	if err != nil {
		t.Fatal(err)
	}
	if a.intArg("n") != 4 || a.strArg("s") != "abc" {
		t.Fatalf("scalar accessors: n=%d s=%q", a.intArg("n"), a.strArg("s"))
	}
	if !reflect.DeepEqual(a.intsArg("nums"), []int{1, 2}) {
		t.Fatalf("intsArg = %v", a.intsArg("nums"))
	}
	if !reflect.DeepEqual(a.strsArg("words"), []string{"hot", "dot"}) {
		t.Fatalf("strsArg = %v", a.strsArg("words"))
	}
	wantGrid := [][]byte{{'1', '0'}, {'0', '1'}}
	if !reflect.DeepEqual(a.byteGridArg("grid"), wantGrid) {
		t.Fatalf("byteGridArg = %v", a.byteGridArg("grid"))
	}
}

func TestArgAccessorErrors(t *testing.T) {
	a, err := parseArgs(`n = "x", k = 1, nums = [1,"a"]`)
	if err != nil {
		t.Fatal(err)
	}
	for _, call := range []func(){
		func() { a.intArg("n") },
		func() { a.intArg("missing") },
		func() { a.strArg("k") },
		func() { a.intsArg("k") },
		func() { a.intsArg("nums") },
		func() { a.strsArg("nums") },
		func() { a.byteGridArg("nums") },
	} {
		func() {
			defer func() {
				r := recover()
				if _, ok := r.(*argError); !ok {
					t.Errorf("recover = %v, want *argError", r)
				}
			}()
			call()
			t.Error("accessor returned instead of panicking")
		}()
	}
}

func TestRunSolverRecoversArgError(t *testing.T) {
	_, err := runSolver(`n = "x"`, func(a args) any { return a.intArg("n") })
	if err == nil || !strings.Contains(err.Error(), "want int") {
		t.Fatalf("runSolver error = %v", err)
	}

	defer func() {
		if recover() == nil {
			t.Fatal("runSolver swallowed a non-arg panic")
		}
	}()
	_, _ = runSolver(`n = 1`, func(args) any { panic("boom") })
}
