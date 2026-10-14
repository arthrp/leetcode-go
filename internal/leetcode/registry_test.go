package leetcode

import "testing"

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
			p.Run()
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
