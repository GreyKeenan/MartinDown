
package mdParser

import (
)

type Header struct {
	Level int
	Text string
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

	return s[3:] //TODO does this work with non ascii
}


func GetHeader(s string) (h Header) {
	s = StripLeadingSpaces(s)

	var measured bool
	for i,v := range s {
		if (v == ' ' || v == '\t' || h.Level == '\n') { //TODO: \n is not on the end because of scan()
			h.Level = i
			measured = true
			break
		}
		if (v != '#') {
			h.Level = 0
			measured = true
			break
		}
		continue
	}

	if (!measured) {
		h.Level = len(s)
		return 
	}
	
	if (h.Level > 6) {
		h.Level = 0
		return
	}
	if (h.Level != 0) {
		h.Text = s[h.Level:]
	}

	return
}
