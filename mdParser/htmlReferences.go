
package mdParser

import (
)

const (
	NONREFERENCE = 0
	ENTITY = iota
	CHARACTER
)

const nonrune = '0'

// https://github.github.com/gfm/#entity-and-numeric-character-references

func CheckReference(s string) (int, rune) {

	var runes []rune = []rune(s)

	if (s[0] != '&') {
		return NONREFERENCE, nonrune
	}
	if (s[1] != '#') {
		return checkReference_entity(runes)
	}

	if (s[2] == 'x' || s[2] == 'X') {
		return checkReference_hexadecimal
	}
	
	return checkReference_decimal
}

//TODO
func checkReference_entity(runes []rune) (int, rune) {
	return NONREFERENCE, nonrune
}
func checkReference_hexadecimal(runes []rune) (int, rune) {
	return NONREFERENCE, nonrune
}
func checkReference_decimal(runes []rune) (int, rune) {
	return NONREFERENCE, nonrune
}
