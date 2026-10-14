package leetcode

import (
	"strings"
	"testing"
)

func TestProblems(t *testing.T) {
	problems := Problems()
	if len(problems) == 0 {
		t.Fatal("Problems() returned no problems")
	}

	seen := make(map[int]bool, len(problems))
	for i, p := range problems {
		if seen[p.Number] {
			t.Errorf("duplicate problem number %d", p.Number)
		}
		seen[p.Number] = true
		if i > 0 && p.Number <= problems[i-1].Number {
			t.Errorf("problems out of order: %d follows %d", p.Number, problems[i-1].Number)
		}
		if p.Title == "" || p.Input == "" {
			t.Errorf("problem %d is missing a title or input", p.Number)
		}
		if p.Run == nil {
			t.Errorf("problem %d has no Run", p.Number)
			continue
		}
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("problem %d Run panicked: %v", p.Number, r)
				}
			}()
			if _, err := p.Run(); err != nil {
				t.Errorf("problem %d Run: %v", p.Number, err)
			}
		}()
	}
}

func TestLookup(t *testing.T) {
	p, ok := Lookup(51)
	if !ok {
		t.Fatal("Lookup(51) found nothing")
	}
	if p.Number != 51 || p.Title != "N-Queens" {
		t.Fatalf("Lookup(51) = %d %q", p.Number, p.Title)
	}

	if _, ok := Lookup(9999); ok {
		t.Fatal("Lookup(9999) found a problem")
	}
}

func TestRunParsesEachCall(t *testing.T) {
	p, ok := Lookup(200)
	if !ok {
		t.Fatal("Lookup(200) found nothing")
	}
	first, err := p.Run()
	if err != nil {
		t.Fatal(err)
	}
	second, err := p.Run()
	if err != nil {
		t.Fatal(err)
	}
	if first != 3 || second != 3 {
		t.Fatalf("Run() = %v then %v, want 3 both times", first, second)
	}
}

func TestLoadProblems(t *testing.T) {
	solvers := map[int]func(args) any{
		1: func(a args) any { return a.intArg("n") },
		2: func(a args) any { return a.intArg("n") },
	}
	text := joinLines(
		"Number    Title    Input",
		"",
		"2    Two    n = 2",
		"1    One    n = 1",
	)
	loaded, err := loadProblems(text, solvers)
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded) != 2 || loaded[0].Number != 1 || loaded[1].Number != 2 {
		t.Fatalf("order = %d, %d", loaded[0].Number, loaded[1].Number)
	}
	if loaded[0].Title != "One" || loaded[0].Input != "n = 1" {
		t.Fatalf("problem 1 = %q %q", loaded[0].Title, loaded[0].Input)
	}
	got, err := loaded[1].Run()
	if err != nil || got != 2 {
		t.Fatalf("Run() = %v, %v", got, err)
	}

	tabbed, err := loadProblems("7\tTab Title\tn = 7\n", map[int]func(args) any{
		7: func(a args) any { return a.intArg("n") },
	})
	if err != nil {
		t.Fatal(err)
	}
	if tabbed[0].Title != "Tab Title" || tabbed[0].Input != "n = 7" {
		t.Fatalf("tabbed row = %q %q", tabbed[0].Title, tabbed[0].Input)
	}
}

func TestLoadProblemsErrors(t *testing.T) {
	one := map[int]func(args) any{
		1: func(args) any { return nil },
	}
	tests := []struct {
		name    string
		text    string
		solvers map[int]func(args) any
		want    string
	}{
		{"columns", "1    Only two columns", one, "want 3 columns"},
		{"number", "nope    Title    n = 1", one, "bad problem number"},
		{"duplicate", joinLines("1    A    n = 1", "1    B    n = 2"), one, "duplicate problem number 1"},
		{"unknown", "9    Missing    n = 1", one, "no solver for problem 9"},
		{"no row", "1    A    n = 1", map[int]func(args) any{
			1: func(args) any { return nil },
			2: func(args) any { return nil },
		}, "solver 2 has no row"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := loadProblems(tt.text, tt.solvers)
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("error = %v, want substring %q", err, tt.want)
			}
		})
	}
}

func joinLines(lines ...string) string {
	return strings.Join(lines, "\n")
}
