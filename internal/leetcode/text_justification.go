package leetcode

import "strings"

// 68 - Text Justification
// fullJustify packs words into lines of exactly maxWidth characters.
// Time is O(lines * maxWidth). Extra memory beyond the result is O(maxWidth).
func fullJustify(words []string, maxWidth int) []string {
	var lines []string
	for i := 0; i < len(words); { // i is the first word of the line being packed
		j := i + 1
		width := len(words[i])
		for j < len(words) && width+1+len(words[j]) <= maxWidth { // +1 is the minimum space before the next word
			width += 1 + len(words[j])
			j++
		}

		gaps := j - i - 1 // slots between words[i:j]
		var b strings.Builder
		b.Grow(maxWidth)
		if j == len(words) || gaps == 0 { // last line, or one word: left-justify
			for k := i; k < j; k++ {
				if k > i {
					b.WriteByte(' ') // one space between words
				}
				b.WriteString(words[k])
			}
			b.WriteString(strings.Repeat(" ", maxWidth-b.Len())) // spaces go on the right
		} else { // fully justify: split the spare spaces across the gaps
			spaces := maxWidth - (width - gaps) // width already counted one space per gap
			base, extra := spaces/gaps, spaces%gaps
			for k := i; k < j; k++ {
				b.WriteString(words[k])
				if k == j-1 {
					break
				}
				n := base
				if k-i < extra { // the leftmost gaps take one extra space
					n++
				}
				b.WriteString(strings.Repeat(" ", n))
			}
		}
		lines = append(lines, b.String())
		i = j
	}
	return lines
}
