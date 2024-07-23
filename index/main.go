
package index

import (
	"os"
	"fmt"
	"errors"
	"bufio"

	"github.com/GreyKeenan/pj.ghmd/mdParser"
)

var bullet string = "1. "

type indexCounter struct {
	levelCounts [6]int
}

func (ic *indexCounter) increment(level int) {
	ic.levelCounts[level]++
	for i := level + 1; i < len(ic.levelCounts); i++ {
		ic.levelCounts[i] = 0
	}
}

func (ic *indexCounter) writeIndexLine(header mdParser.Header) (s string) {

	var level int = header.GetLevel()

	for i := 0; i < level; i++ {
		s += "\t"
	}
	s += bullet

	ic.increment(level)

	s += fmt.Sprint(ic.levelCounts[0])
	for i := 1; i <= level; i++ {
		s += "-" + fmt.Sprint(ic.levelCounts[i])
	}

	s += header.Text
	s += "\n"

	return
}


func Main() (err error) {
	if (len(os.Args) < 3) {
		err = errors.New("No filepath given.")
		return
	}
	
	var f *os.File
	f, err = os.Open(os.Args[2])
	if (err != nil) {
		return
	}
	defer f.Close()


	var s *bufio.Scanner
	s = bufio.NewScanner(f)

	var header mdParser.Header
	var headers []mdParser.Header
	for s.Scan() {
		
		header = mdParser.GetHeader(s.Text())

		if (header.IsHeader()) {
			headers = append(headers, header)
		}

	}
	err = s.Err()
	if (err != nil) {
		return
	}


	for _,v := range headers {
		fmt.Println(v.Text)
	}
	
	return
}
