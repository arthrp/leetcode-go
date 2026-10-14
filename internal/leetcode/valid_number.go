package leetcode

// 65 - Valid Number
// isNumber reports whether s is a valid number.
//
// A valid number is an integer or a decimal, each followed by an optional
// exponent. An integer is an optional sign and then digits. A decimal is an
// optional sign and then digits around a dot, with digits on at least one
// side of the dot. An exponent is 'e' or 'E', then an integer.
//
// The scan keeps three facts. digit is true once the current part has a
// digit. dot is true once a '.' has been seen. exp is true once an exponent
// marker has been seen. A sign is allowed only at the start or right after
// the exponent marker. A dot is allowed only before the exponent and only
// once. An exponent marker needs a digit before it and, once accepted, clears
// digit so the exponent must supply digits of its own. Anything else fails.
//
// The string is valid exactly when the current part ends with a digit, so the
// return value is digit. Time is O(len(s)). Extra memory is O(1).
func isNumber(s string) bool {
	digit, dot, exp := false, false, false
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case c >= '0' && c <= '9':
			digit = true
		case c == '+' || c == '-':
			if i > 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		case c == '.':
			if dot || exp {
				return false
			}
			dot = true
		case c == 'e' || c == 'E':
			if exp || !digit {
				return false
			}
			exp, digit = true, false
		default:
			return false
		}
	}
	return digit
}
