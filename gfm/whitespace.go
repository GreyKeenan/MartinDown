
package gfm

const CountIndent_Blankline = -1

func CountSpaces(r rune) int {
	switch (r) {
		case ' ':
			return 1
		case '\t':
			return 4
		default:
			return 0
	}
}
func CountIndent(runes []rune) int {
	
	var tally, x int
	for _,r := range runes {
		x = CountSpaces(r)
		if (x == 0) {
			return tally
		}
		tally += x
	}

	return CountIndent_Blankline
}
