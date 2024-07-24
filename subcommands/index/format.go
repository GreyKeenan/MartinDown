
package index

import (
	"unicode"

	"github.com/GreyKeenan/pj.ghmd/gfm"
)

const output_Leader = "> auto-index ~~courtesy~~ fault of github.com/GreyKeenan/pj.ghmd\n<!-- index generated using github.com/GreyKeenan/pj.ghmd -->\n\n"

/*
func buildHeaderId(ic *indexCounter, level int, splitter string) string {
	var s string = fmt.Sprint(ic.levelCounts[0])
	for i := 1; i <= level; i++ {
		s += splitter + fmt.Sprint(ic.levelCounts[i])
	}

	return s
}
*/

func buildIndexFollower(ic *indexCounter) string {
	var s string

	for i := 0; i < ic.previousLevel; i++ {
		s += "</ul>"
	}
	ic.previousLevel = 0

	return s + "\n" + "\n---\n\n"
}

func buildIndexLine(ic *indexCounter, header gfm.Header) string {
	ic.increment(header.Level)
	var s string 

	for i := header.Level; i > ic.previousLevel; i-- {
		s += "<ul>"
	}
	for i := header.Level; i < ic.previousLevel; i++ {
		s += "</ul>"
	}
	s += "\n"

	s += "<li><a href=\"" + buildIndexAnchor(ic, header) + "\">" + string(header.Text) + "</a></li>\n"
	ic.previousLevel = header.Level
	return s
}

func buildIndexAnchor(ic *indexCounter, header gfm.Header) string {
	var s string = "#"

	//TODO: add id

	var runes []rune = gfm.StripWhitespace_left(header.Text)
	runes = gfm.StripWhitespace_right(runes)

	//TODO: parse &...; here

	for _,v := range runes {
		switch (v) {
			case '_', '-':
				s += string(v)
				continue
			case ' ':
				s += "-"
				continue
		}
		if (gfm.IsUnicodeWhitespace(v) || gfm.IsPunctuation(v)) {
			continue
		}
		s += string(unicode.ToLower(v))
	}

	return s
}
/*

you can inspect the generated html href string on github
	or in url if you just click the link img near the header

observed rules:
	non-leading and non-trailing spaces are converted to '-'
		multiple sequential spaces become multiple dashes on github.com. In vscode, they become single dashes
	&#x20; (hex for space) is converted to '-'
	&#x20; cannot be a leading or trailing space

	'-' remains
	'_' remains

	any other punctuation is omitted
		omitted punctuation will still interrupt leading/trailing spaces
	any other whitespace is omitted
		omitted whitespace will NOT interrupt leading/traling spaces EXCEPT:
		EXCEPT if given as a &...; value. References will always interrupt whitespace
	

	so it seems like the order of operations is:
		1. strip leading/trailing whitespace
		2. decode &...;
		3. replace characters
			a. to lowercase
			b. - -> -
			c. _ -> _
			d. ' ' -> -
			e. other punctuation omitted


	These rules were observed in the edit/preview windows of github.com. They differ for other 'github-style-markdown' program, like in vscode

	There may be more rules to this. I probably havent seen every edge case

	I really wish this was just specified explicitly, jeez

*/

