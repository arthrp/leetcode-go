package leetcode

import (
	"fmt"
	"strconv"
)

// args holds the named values parsed from a LeetCode sample input, such as
// `head = [1,2,3,4,5], k = 2`. Values are int, string, or []any.
type args map[string]any

// argError is the panic value the accessors raise for a missing or mistyped
// argument. Run recovers it and returns it as an error.
type argError struct {
	msg string
}

func (e *argError) Error() string { return e.msg }

// parseArgs parses a LeetCode sample input into named arguments.
func parseArgs(s string) (args, error) {
	p := parser{s: s}
	return p.parseList()
}

type parser struct {
	s string
	i int
}

func (p *parser) parseList() (args, error) {
	out := make(args)
	p.skipSpace()
	if p.done() {
		return out, nil
	}
	for {
		name, err := p.ident()
		if err != nil {
			return nil, err
		}
		p.skipSpace()
		if p.done() || p.s[p.i] != '=' {
			return nil, fmt.Errorf("missing = after %q", name)
		}
		p.i++
		p.skipSpace()
		v, err := p.value()
		if err != nil {
			return nil, err
		}
		if _, ok := out[name]; ok {
			return nil, fmt.Errorf("duplicate argument %q", name)
		}
		out[name] = v
		p.skipSpace()
		if p.done() {
			return out, nil
		}
		if p.s[p.i] != ',' {
			return nil, fmt.Errorf("unexpected %q", p.s[p.i:])
		}
		p.i++
		p.skipSpace()
		if p.done() {
			return nil, fmt.Errorf("trailing comma")
		}
	}
}

func (p *parser) ident() (string, error) {
	if p.done() || !isIdentStart(p.s[p.i]) {
		return "", fmt.Errorf("expected argument name")
	}
	start := p.i
	p.i++
	for !p.done() && isIdentCont(p.s[p.i]) {
		p.i++
	}
	return p.s[start:p.i], nil
}

func (p *parser) value() (any, error) {
	if p.done() {
		return nil, fmt.Errorf("missing value")
	}
	switch p.s[p.i] {
	case '"':
		return p.str()
	case '[':
		return p.array()
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return p.integer()
	default:
		return nil, fmt.Errorf("unexpected %q", p.s[p.i:])
	}
}

func (p *parser) integer() (int, error) {
	start := p.i
	if p.s[p.i] == '-' {
		p.i++
	}
	digits := p.i
	for !p.done() && p.s[p.i] >= '0' && p.s[p.i] <= '9' {
		p.i++
	}
	if p.i == digits {
		return 0, fmt.Errorf("bad integer %q", p.s[start:])
	}
	n, err := strconv.Atoi(p.s[start:p.i])
	if err != nil {
		return 0, fmt.Errorf("bad integer %q", p.s[start:p.i])
	}
	return n, nil
}

func (p *parser) str() (string, error) {
	start := p.i
	p.i++ // opening quote
	for !p.done() {
		switch p.s[p.i] {
		case '\\':
			p.i += 2
			if p.i > len(p.s) {
				return "", fmt.Errorf("unterminated string")
			}
		case '"':
			p.i++
			s, err := strconv.Unquote(p.s[start:p.i])
			if err != nil {
				return "", fmt.Errorf("bad string %q: %w", p.s[start:p.i], err)
			}
			return s, nil
		default:
			p.i++
		}
	}
	return "", fmt.Errorf("unterminated string")
}

func (p *parser) array() ([]any, error) {
	p.i++ // '['
	p.skipSpace()
	if !p.done() && p.s[p.i] == ']' {
		p.i++
		return []any{}, nil
	}
	var out []any
	for {
		if p.done() {
			return nil, fmt.Errorf("unterminated array")
		}
		v, err := p.value()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
		p.skipSpace()
		if p.done() {
			return nil, fmt.Errorf("unterminated array")
		}
		switch p.s[p.i] {
		case ']':
			p.i++
			return out, nil
		case ',':
			p.i++
			p.skipSpace()
		default:
			return nil, fmt.Errorf("expected comma in array")
		}
	}
}

func (p *parser) skipSpace() {
	for !p.done() && (p.s[p.i] == ' ' || p.s[p.i] == '\t') {
		p.i++
	}
}

func (p *parser) done() bool { return p.i >= len(p.s) }

func isIdentStart(b byte) bool {
	return b == '_' || b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}

func isIdentCont(b byte) bool {
	return isIdentStart(b) || b >= '0' && b <= '9'
}

func (a args) must(name string) any {
	v, ok := a[name]
	if !ok {
		panic(&argError{fmt.Sprintf("missing argument %q", name)})
	}
	return v
}

func (a args) intArg(name string) int {
	v := a.must(name)
	n, ok := v.(int)
	if !ok {
		panic(&argError{fmt.Sprintf("argument %q is %T, want int", name, v)})
	}
	return n
}

func (a args) strArg(name string) string {
	v := a.must(name)
	s, ok := v.(string)
	if !ok {
		panic(&argError{fmt.Sprintf("argument %q is %T, want string", name, v)})
	}
	return s
}

func (a args) intsArg(name string) []int {
	items := a.sliceArg(name)
	out := make([]int, len(items))
	for i, item := range items {
		n, ok := item.(int)
		if !ok {
			panic(&argError{fmt.Sprintf("argument %q[%d] is %T, want int", name, i, item)})
		}
		out[i] = n
	}
	return out
}

func (a args) strsArg(name string) []string {
	items := a.sliceArg(name)
	out := make([]string, len(items))
	for i, item := range items {
		s, ok := item.(string)
		if !ok {
			panic(&argError{fmt.Sprintf("argument %q[%d] is %T, want string", name, i, item)})
		}
		out[i] = s
	}
	return out
}

// byteGridArg reads a 2D array of one-character strings as a byte grid.
func (a args) byteGridArg(name string) [][]byte {
	rows := a.sliceArg(name)
	out := make([][]byte, len(rows))
	for r, row := range rows {
		cells, ok := row.([]any)
		if !ok {
			panic(&argError{fmt.Sprintf("argument %q[%d] is %T, want array", name, r, row)})
		}
		out[r] = make([]byte, len(cells))
		for c, cell := range cells {
			s, ok := cell.(string)
			if !ok || len(s) != 1 {
				panic(&argError{fmt.Sprintf("argument %q[%d][%d] is %v, want a one-character string", name, r, c, cell)})
			}
			out[r][c] = s[0]
		}
	}
	return out
}

func (a args) sliceArg(name string) []any {
	v := a.must(name)
	items, ok := v.([]any)
	if !ok {
		panic(&argError{fmt.Sprintf("argument %q is %T, want array", name, v)})
	}
	return items
}
