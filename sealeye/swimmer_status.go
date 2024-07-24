package sealeye
import (
)

//go:generate stringer -type=Swimmer_Status

type Swimmer_Status int 
const (
	Swimmer_Error Swimmer_Status = iota
	Swimmer_Error_Flag
	Swimmer_Error_Flag_Type
	Swimmer_Error_Flag_Unrecognized
	Swimmer_Error_Flag_NotFound
	Swimmer_Error_Underflow
	Swimmer_Error_Underflow_Flag
	Swimmer_Error_Underflow_End
	Swimmer_Error_Underflow_Subcommand
	Swimmer_Error_Overflow
)
const ( //TODO errs should only be errs. Not helped
	Swimmer_Helped Swimmer_Status = 0 - iota
)

func (self Swimmer_Status) Error() string {
	return self.String()
}
