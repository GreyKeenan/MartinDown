
package index

import (
	"os"
	"fmt"
	"errors"
	"bufio"

	"github.com/GreyKeenan/pj.ghmd/mdParser"
)


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
	for s.Scan() {
		
		header = mdParser.GetHeader(s.Text())
		if (header.Level != 0) {
			fmt.Println(header.Text)
		}
		
	}
	err = s.Err()
	if (err != nil) {
		return
	}
	
	return
}
