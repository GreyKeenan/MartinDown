package sealeye
import (
	"fmt"
)

const OverflowMax_Unbound = -1

type CommandSpec struct {
	Help string
	Subcommands map[string]*CommandSpec
	LongFlags map[string]uint8 //uint8 is the overflow of the flag
	ShortFlags map[rune]string
	OverflowMin int
	OverflowMax int
}

type CommandResponse struct {
	Keyword string
	Flags [][]string
	Overflow []string
}

func (self *CommandResponse) String() string {
	var s string = self.Keyword + "\nFlags:\n"

	for _,v := range self.Flags {
		s += fmt.Sprintln(v)
	}
	s += "Overflow:" + fmt.Sprint(self.Overflow)

	return s
}
