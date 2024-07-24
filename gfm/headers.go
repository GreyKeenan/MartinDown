
package gfm

type Header struct {
	Level int
	Text []rune
}

const (
	header_MaxIndent = 3
	header_MaxLevel = 6
	header_Rune = '#'
)
func (self *Header) IsHeader() bool {
	return self.Level != 0
}

// For now, gfm assumes it is given a full line, & that line does not have a newline char at the end

// https://github.github.com/gfm/#atx-headings
func GetHeader(runes []rune) (self Header) {

	var indent int = CountIndent(runes)
	if (indent > header_MaxIndent || indent == CountIndent_Blankline) {
		return
	}
	runes = runes[indent:]
	
	var measured bool
	for i,r := range runes {
		if (CountSpaces(r) > 0 || IsLineEnding(r)) {
			self.Level = i
			measured = true
			break
		}
		if (r != header_Rune) {
			return
		}
	}

	if (!measured) {
		self.Level = len(runes)
		return
	}
	if (self.Level > header_MaxLevel) {
		self.Level = 0
		return
	}

	self.Text = runes[self.Level:]
	return
}
