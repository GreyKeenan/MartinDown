package sealeye
import (
)

type FlagType uint8
const (
	FlagType_Not = iota
	FlagType_Short 
	FlagType_Long 
)

type Deflagger interface {
	IsFlag(string) FlagType
	Deflag_short(string) string
	Deflag_long(string) string
}


type defaultDeflagger struct { }
func (_ defaultDeflagger) IsFlag(s string) FlagType {
	if (len(s) < 2) {
		return FlagType_Not
	}
	if s[0] == '-' {
		if s[1] == '-' {
			return FlagType_Long
		}
		return FlagType_Short
	}
	return FlagType_Not
}
func (_ defaultDeflagger) Deflag_short(s string) string { return s[1:] }
func (_ defaultDeflagger) Deflag_long(s string) string { return s[2:] }

func NewDefaultDeflagger() Deflagger {
	return defaultDeflagger { }
}
