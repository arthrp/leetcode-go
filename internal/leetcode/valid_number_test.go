package leetcode

import "testing"

func TestIsNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "zero", s: "0", want: true},
		{name: "just e", s: "e", want: false},
		{name: "just dot", s: ".", want: false},
		{name: "digits", s: "2", want: true},
		{name: "leading zeros", s: "0089", want: true},
		{name: "negative decimal", s: "-0.1", want: true},
		{name: "positive decimal", s: "+3.14", want: true},
		{name: "trailing dot", s: "4.", want: true},
		{name: "leading dot", s: "-.9", want: true},
		{name: "exponent", s: "2e10", want: true},
		{name: "negative with capital exponent", s: "-90E3", want: true},
		{name: "positive exponent sign", s: "3e+7", want: true},
		{name: "negative exponent", s: "+6e-1", want: true},
		{name: "decimal exponent", s: "53.5e93", want: true},
		{name: "signed decimal and exponent", s: "-123.456e789", want: true},
		{name: "letters", s: "abc", want: false},
		{name: "digit then letter", s: "1a", want: false},
		{name: "exponent missing digits", s: "1e", want: false},
		{name: "exponent first", s: "e3", want: false},
		{name: "decimal in exponent", s: "99e2.5", want: false},
		{name: "double minus", s: "--6", want: false},
		{name: "minus then plus", s: "-+3", want: false},
		{name: "letter inside", s: "95a54e53", want: false},
		{name: "just plus", s: "+", want: false},
		{name: "plus then dot", s: "+.", want: false},
		{name: "dot then exponent", s: ".e1", want: false},
		{name: "exponent sign missing digits", s: "1e+", want: false},
		{name: "trailing dot before exponent", s: "1.e5", want: true},
		{name: "capital exponent with sign", s: "1E-7", want: true},
		{name: "sign in the middle", s: "1+2", want: false},
		{name: "two exponents", s: "1e5e5", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNumber(tt.s)
			if got != tt.want {
				t.Fatalf("isNumber(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
