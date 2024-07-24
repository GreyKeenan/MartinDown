package main
import (

	"github.com/GreyKeenan/pj.ghmd/sealeye"
)

func buildCommand_root() *sealeye.CommandSpec {
	var cmd = sealeye.CommandSpec { 
		Help: "TODO: HELP", //TODO
		OverflowMin: 0,
		OverflowMax: 0,
		LongFlags: map[string]uint8 {
			"version": 0,
		},
		ShortFlags: map[rune]string {
			'v': "version",
		},
		Subcommands: map[string]*sealeye.CommandSpec {
			"index": buildCommand_index(),
		},
	}
	return &cmd
}


func buildCommand_index() *sealeye.CommandSpec {
	var cmd = sealeye.CommandSpec {
		Help: "HELP TODO", //TODO
		OverflowMin: 1,
		OverflowMax: 2,
		LongFlags: map[string]uint8 {
			"bullet": 1, //string
			"ids": 1, //none, headers-only, index
			"id-position": 1, //leading, trailing
			"reformat": 0,
			"nohtml": 0,
			"unlinked": 0,
			"undecorated": 0,
			"in-place": 0,
		},
		ShortFlags: map[rune]string {
			'b': "bullet",
			'f': "reformat",
			'i': "in-place",
		},
	}
	return &cmd
}
