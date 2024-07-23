
package mdParser

import (
	"unicode"
)

// based on: https://github.github.com/gfm/#characters-and-lines

func IsASCIIPunctuation(r rune) bool {
	switch (r) {
		case '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~':
			return true
		default:
			return false
	}
}
func IsPunctuation(r rune) bool {
	if (IsASCIIPunctuation(r)) {
		return true
	}
	if (unicode.In(r, unicode.Pc, unicode.Pd, unicode.Pe, unicode.Pf, unicode.Pi, unicode.Po, unicode.Ps)) {
		return true
	}

	return false
}
func IsUnicodeWhitespace(r rune) bool {
	if (unicode.In(r, unicode.Zs)) {
		return true
	}
	switch (r) {
		case '\t', '\r', '\n', '\u000c':
			return true
		default:
			return false
	}
}
func IsSpace(r rune) int {
	switch (r) {
		case ' ':
			return 1
		case '\t':
			return 4
		default:
			return false
	}
}
func IsWhitespace(r rune) bool {
	switch (r) {
		case ' ', '\t', '\n', '\u000b', '\u000c', '\r':
			return true
		default:
			return false
	}
}
func IsLineEnding(r rune) bool {
	switch (r) {
		case '\n', '\r':
			return true
		default:
			return false
	}
}


/*
observed rules for automatically linking to headers in GFM:
	(I cant find the details of this in the spec. Is it non-standard?)

	in vscode gfm display:
		leading or trailing spaces are removed
		any sequence of non-leading, non-trailing spaces becomes '-'
		any ?symbols? or ?punctuation? are removed
		all capitalization becomes lowercase

		values entered as unicode literals (&#0000000;) are parsed as that rune.

	in github md preview:
		leading/trailing spaces removed
		multiple spaces become multiple '-' dashes
		actual '-' dashes stay as dashes

		tabs are removed, not converted to '-'

		\u000a displays as a space in the header, but is not converted to a '-'. It is removed
			same for \u000d

		non-ascii unicode characters are still respected as letters/not removed or dashed

		seems like the only thing that gets dashed is spaces, and all (non trailing/leading) spaces get dashed
			including &...; spaces
		
		as for what gets removed, im not as sure. have observed:
			any non-space whitespace is removed from id
			
			anything from the gfm "punctuation" category is removed
				EXCEPT underscores
		
		always lowercase

*/

/*
general todos in order to be accurate
	- recognize code blocks
	- recognize <!-- --> comments

	- ? recognize setext headers

*/
