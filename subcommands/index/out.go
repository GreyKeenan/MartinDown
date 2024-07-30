package index
import (
	"os"
	"bufio"

	"github.com/GreyKeenan/MartinDown/gfm"
)

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

	_, err = outputFile.WriteString(buildIndexLeader())
	if (err != nil) {
		return err
	}


	for _,v := range headers {
		_, err = outputFile.WriteString(buildIndexLine(&ic, v))
		if (err != nil) {
			return err
		}
	}


	_, err = outputFile.WriteString(buildIndexFollower(&ic))
	if (err != nil) {
		return err
	}


	return nil
}
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
