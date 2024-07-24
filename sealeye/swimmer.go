package sealeye
import (
	"os"
	"fmt"
)

type Swimmer struct {
	pos int
	response CommandResponse

	cmd *CommandSpec
	deflagger Deflagger

	args []string
}
func NewSwimmer(root *CommandSpec) Swimmer {
	return Swimmer {
		cmd:root,
		deflagger:NewDefaultDeflagger(),
		args:os.Args,
	}
}
func (self *Swimmer) SetDeflagger(df Deflagger) {
	self.deflagger = df
}
func (self *Swimmer) Help() error {
	
	//TODO? formatter for help
	fmt.Printf("\n%v\n", self.cmd.Help)
	
	return Swimmer_Helped
}

func (self *Swimmer) current() string {
	return self.args[self.pos]
}

func (self *Swimmer) nextArg() bool {
	self.pos++
	return self.pos >= len(self.args)
}

func (self *Swimmer) Swim() (error, CommandResponse, bool) {

	self.response.Keyword = self.current()

	var err error

	var shortFlag string
	var longFlag string
	var exists bool

	var nextCmd *CommandSpec

	for {
		if (self.nextArg()) {
			if len(self.response.Overflow) < self.cmd.OverflowMin {
				return Swimmer_Error_Underflow_End, self.response, true
			}

			return nil, self.response, true
		}

		switch (self.deflagger.IsFlag(self.current())) {
			case FlagType_Not:
			case FlagType_Short:
				shortFlag = self.deflagger.Deflag_short(self.current())
				if ([]rune(shortFlag)[0] == 'h') {
					return self.Help(), self.response, false
				}
				for _,v := range shortFlag {
					longFlag, exists = self.cmd.ShortFlags[v]
					if (!exists) {
						return Swimmer_Error_Flag_NotFound, self.response, false
					}
					err = self.addFlag(longFlag)
					if (err != nil) {
						return err, self.response, false
					}
				}
				continue
			case FlagType_Long:
				longFlag = self.deflagger.Deflag_long(self.current())
				if (longFlag == "help") {
					return self.Help(), self.response, false
				}
				err = self.addFlag(longFlag)
				if (err != nil) {
					return err, self.response, false
				}
				continue
			default:
				return Swimmer_Error_Flag_Type, self.response, false
		}
		
		nextCmd, exists = self.cmd.Subcommands[self.current()]
		if (exists) {

			
			if len(self.response.Overflow) < self.cmd.OverflowMin {
				self.cmd = nextCmd
				return Swimmer_Error_Underflow_Subcommand, self.response, false
			}

			self.cmd = nextCmd
			return nil, self.response, false
		}

		if len(self.response.Overflow) >= self.cmd.OverflowMax {
			return Swimmer_Error_Overflow, self.response, false
		}

		self.response.Overflow = append(self.response.Overflow, self.current())
	}
}

func (self *Swimmer) addFlag(flag string) error {

	var overflow uint8
	var exists bool

	overflow, exists = self.cmd.LongFlags[flag]
	if (!exists) {
		return Swimmer_Error_Flag_Unrecognized
	}

	var flagResponse []string
	flagResponse = append(flagResponse, flag)

	for i := uint8(0); i < overflow; i++ {
		if (self.nextArg()) { //TODO? check if next value is a flag
			return Swimmer_Error_Underflow_Flag
		}
		flagResponse = append(flagResponse, self.current())
	}

	self.response.Flags = append(self.response.Flags, flagResponse)

	return nil
}
