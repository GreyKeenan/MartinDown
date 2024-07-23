
package mdParser

//honestly might make more since as a module than a package but this is fine since its incomplete

import (
)

type Header struct {
	level int
	Text string
}
func (self *Header) IsHeader() bool {
	return self.level != 0
}
func (self *Header) GetLevel() int {
	return self.level - 1
}

func StripLeadingSpaces(s string) string {

	for i,v := range s {
		if (i > 2) {
			break
		}
		if (v != ' ') {
			return s[i:]
		}
	}

	if (len(s) < 4) {
		return ""
	}

	return s[3:] //TODO does this work with non ascii | and below
		//it does not. looking here: https://groups.google.com/g/golang-nuts/c/YyKlLwuWt3w for a solution
		//can cast string to []rune ?? Yes: https://go.dev/ref/spec#Conversions
			// it will convert properly. Thats so slick
}


func GetHeader(s string) (h Header) { // only ATX, not SETEXT
	//TODO: doesnt respect quoted/code-blocked non-headers
	s = StripLeadingSpaces(s)

	var measured bool
	for i,v := range s {
		if (v == ' ' || v == '\t' || h.level == '\n') { //am using with Scanner which strips the \n
			h.level = i
			measured = true
			break
		}
		if (v != '#') {
			h.level = 0
			measured = true
			break
		}
		continue
	}

	if (!measured) {
		h.level = len(s)
		return 
	}
	
	if (h.level > 6) {
		h.level = 0
		return
	}
	if (h.level != 0) {
		h.Text = s[h.level:]
	}

	return
}
