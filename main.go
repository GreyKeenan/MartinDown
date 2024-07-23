
package main

import (
	"os"
	"fmt"

	"github.com/GreyKeenan/pj.ghmd/subcommands/index"

)

const VERSION = "v0.0.0"

func main() {
	// identify subcommand

	if (len(os.Args) == 1) {
		fmt.Println(os.Args[0], VERSION)
		return
	}
	
	switch (os.Args[1]) {
		case "index":
			index.Main(os.Args[1:])
		default:
			fmt.Printf("sub-command: \"%v\" not recognized. Currently, \"index\" is the only option.", os.Args[1])
	}
}
