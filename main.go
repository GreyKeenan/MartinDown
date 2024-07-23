
package main

import (
	"os"
	"fmt"

	"github.com/GreyKeenan/pj.ghmd/index"
)


func check(err error) {
	if (err != nil) {
		panic(err)
	}
}

func main() {
	
	var err error

	if (len(os.Args) < 2) {
		fmt.Println("ghmd v0.0.0")
		return
	}
	switch (os.Args[1]) {
		case "index":
			err = index.Main()
			check(err)
		default:
			panic("Unrecognized sub-command. Currently, 'index' is the only option.")
	}
}
/*
get filename from
get stream from file
read through stream and build index
print index to a file
	adds header-ids in case of duplicate names
read through stream again and append to that file. Insert header-ids here
*/
