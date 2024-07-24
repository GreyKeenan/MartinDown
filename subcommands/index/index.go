
package index

import (
	"fmt"
	"os"
	"bufio"

	"github.com/GreyKeenan/pj.ghmd/gfm"
)

func Main(args []string) {

	var err error

	if (len(args) == 1) {
		fmt.Println("No file path given.")
		return
	}

	// check for flags in the future

	var inputFile *os.File
	inputFile, err = os.Open(args[1])
	if (err != nil) {
		fmt.Println("Error opening file:", args[1])
		return
	}
	defer inputFile.Close()

	var headers []gfm.Header

	headers, err = readHeaders(inputFile)
	if (err != nil) {
		fmt.Println("Error reading file:", args[1])
		return
	}


	var outputPath string

	if (len(args) == 2) {
		outputPath = "autoindexed-" + args[1] //TODO ERR if prefixed with dir this doesnt work properly
	} else {
		outputPath = args[2]
	}

	_, err = inputFile.Seek(0, 0)
	if (err != nil) {
		fmt.Println("Failed to write output to:", outputPath) //TODO
		return
	}
	err = write(outputPath, headers, inputFile)
	if (err != nil) {
		fmt.Println("Failed to write output to:", outputPath)
		return
	}

}

func readHeaders(inputFile *os.File) ([]gfm.Header, error) {

	var inputFile_scanner *bufio.Scanner = bufio.NewScanner(inputFile)

	var header gfm.Header
	var headers []gfm.Header

	for inputFile_scanner.Scan() {
		//TODO: identify <!-- --> comments & ``` ... ``` code blocks

		header = gfm.GetHeader([]rune(inputFile_scanner.Text()))
		if (header.IsHeader()) {
			headers = append(headers, header)
		}

	}
	if inputFile_scanner.Err() != nil {
		return nil, inputFile_scanner.Err()
	}

	return headers, nil
}
