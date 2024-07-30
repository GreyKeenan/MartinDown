
package main

import (
	"fmt"

	"github.com/GreyKeenan/MartinDown/subcommands/index"
	"github.com/GreyKeenan/MartinDown/sealeye"

)

const VERSION = "v0.0.0"

func main() {

	var err error

	var swimmer sealeye.Swimmer = sealeye.NewSwimmer(buildCommand_root())

	var cmd sealeye.CommandResponse
	var finished bool

	err, cmd, finished = swimmer.Swim()
	if err != nil {
		panic(err)
	}
	if finished {
		fmt.Println(VERSION)
		return
	}
	
	for !finished {
		err, cmd, finished = swimmer.Swim()
		if err != nil {
			panic(err)
		}
		fmt.Println(cmd.String())

		switch (cmd.Keyword) {
			case "index":
				index.Main(cmd)
				return
			default:
				panic("Unrecognized Subcommand Keyword")
		}
	}

}
