
package index

import (
	"fmt"
	"os"
	"bufio"

	"github.com/GreyKeenan/pj.ghmd/gfm"
)

func Main(args []string) {

	if (len(args) == 1) {
		fmt.Println("No file path given.")
		return
	}

	// check for flags in the future
	
	var err error

	var inputFile *os.File
	inputFile, err = os.Open(args[1])
	if (err != nil) {
		fmt.Println("Could not open file:", args[1])
		return
	}
	defer inputFile.Close()


	var inputFile_scanner *bufio.Scanner
	inputFile_scanner = bufio.NewScanner(inputFile)

	var header gfm.Header

	for inputFile_scanner.Scan() {

		header = gfm.GetHeader([]rune(inputFile_scanner.Text()))
		if (header.IsHeader()) {
			fmt.Printf("header: %d, \"%v\"\n", header.GetLevel(), string(header.Text))
		}

	}
	if inputFile_scanner.Err() != nil {
		panic(inputFile_scanner.Err())
	}
}
