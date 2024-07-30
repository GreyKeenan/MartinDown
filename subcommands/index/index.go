
package index

import (
	"fmt"
	"os"
	"bufio"

	"github.com/GreyKeenan/MartinDown/gfm"
	"github.com/GreyKeenan/MartinDown/sealeye"
)

func Main(cmd sealeye.CommandResponse) {

	var err error

	var inputFile *os.File
	inputFile, err = os.Open(cmd.Overflow[0])
	if (err != nil) {
		fmt.Println("Error opening file:", cmd.Overflow[0])
		return
	}
	defer inputFile.Close()

	var headers []gfm.Header

	headers, err = readHeaders(inputFile)
	if (err != nil) {
		fmt.Println("Error reading file:", cmd.Overflow[0])
		return
	}


	var outputPath string

	if (len(cmd.Overflow) == 1) {
		outputPath = cmd.Overflow[0] + "-autoindex.md"
	} else {
		outputPath = cmd.Overflow[1]
	}

	_, err = inputFile.Seek(0, 0)
	if (err != nil) {
		fmt.Println("Failed to write output to:", outputPath)
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
