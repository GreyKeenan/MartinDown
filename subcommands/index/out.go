
package index

import (
	"os"
	"unicode"
	"bufio"

	"github.com/GreyKeenan/pj.ghmd/gfm"
)

type indexCounter struct {
	levelCounts [6]int
}

const output_leader = "### Index\n<!-- index generated using github.com/GreyKeenan/pj.ghmd -->\n\n"
const output_breakline = "\n---\n\n"

func (self *indexCounter) increment(level int) {
	self.levelCounts[level]++
	for i := level + 1; i < len(self.levelCounts); i++ {
		self.levelCounts[i] = 0
	}
}

func write(outputPath string, headers []gfm.Header, inputFile *os.File) error {

	var err error

	var outputFile *os.File
	outputFile, err = os.Create(outputPath)
	if (err != nil) {
		return err
	}
	defer outputFile.Close()

	err = write_index(outputFile, headers)
	if (err != nil) {
		return err
	}

	err = write_body(outputFile, inputFile)
	if (err != nil) {
		return err
	}
	
	return nil
}

func write_index(outputFile *os.File, headers []gfm.Header) error {

	var err error

	var ic indexCounter

	_, err = outputFile.WriteString(output_leader)
	if (err != nil) {
		return err
	}

	for _,v := range headers {
		_, err = outputFile.WriteString(buildIndexLine(&ic, v, "* "))
		if (err != nil) {
			return err
		}
	}

	_, err = outputFile.WriteString(output_breakline)
	if (err != nil) {
		return err
	}

	return nil
}

/*
func buildHeaderId(ic *indexCounter, level int, splitter string) string {
	var s string = fmt.Sprint(ic.levelCounts[0])
	for i := 1; i <= level; i++ {
		s += splitter + fmt.Sprint(ic.levelCounts[i])
	}

	return s
}
*/

func buildIndexLine(ic *indexCounter, header gfm.Header, bullet string) string {
	ic.increment(header.Level)
	var s string
	for i := 1; i < header.Level; i++ { //TODO ERR if first header isnt lvl 1, then, it will indent all
		s += "\t"
	}
	return s + bullet + "[" + string(header.Text) + "]" + buildIndexAnchor(ic, header) + "\n"
}

func buildIndexAnchor(ic *indexCounter, header gfm.Header) string {
	var s string = "(#"

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

	s += ")"
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

func write_body(outputFile *os.File, inputFile *os.File) error {
	
	var err error
	var inputFile_scanner *bufio.Scanner = bufio.NewScanner(inputFile)

	for inputFile_scanner.Scan() {
		//TODO: identify <!-- --> comments & ``` ... ``` code blocks
		_, err = outputFile.WriteString(inputFile_scanner.Text() + "\n")
		if (err != nil) {
			return err
		}
	}

	return inputFile_scanner.Err()
}
