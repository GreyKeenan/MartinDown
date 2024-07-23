
package gfm

type Header struct {
	level int
	Text []rune
}

const (
	header_Not = 0
	Header_Not = -1
	header_MaxIndent = 3
	header_MaxLevel = 6
	header_Rune = '#'
)

func (self *Header) GetLevel() int {
	return self.level - 1
}
func (self *Header) IsHeader() bool {
	return self.level != header_Not
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
			self.level = i
			measured = true
			break
		}
		if (r != header_Rune) {
			return
		}
	}

	if (!measured) {
		self.level = len(runes)
		return
	}
	if (self.level > header_MaxLevel) {
		self.level = header_Not
		return
	}

	self.Text = runes[self.level:]
	return
}
