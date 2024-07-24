
package gfm

import "testing"



func test_CountIndent(t *testing.T, s string, expect int) {
	var got int = CountIndent([]rune(s))
	if (got != expect) {
		t.Fatalf("Got (%v) from (%v) instead of (%v).", got, s, expect)
	}
}
func Test_CountIndent(t *testing.T) {

	test_CountIndent(t, "test", 0)
	test_CountIndent(t, "\t\t    ", CountIndent_Blankline)
	test_CountIndent(t, "    \t x", 9)

}

func test_StripWhitespace_left(t *testing.T, in string, expect string) {
	var gotR []rune = StripWhitespace_left([]rune(in))
	var got string = string(gotR)
	if (got != expect) {
		t.Fatalf("Got (%v) from (%v) instead of (%v).", got, in, expect)
	}
}
func test_StripWhitespace_right(t *testing.T, in string, expect string) {
	var gotR []rune = StripWhitespace_right([]rune(in))
	var got string = string(gotR)
	if (got != expect) {
		t.Fatalf("Got (%v) from (%v) instead of (%v).", got, in, expect)
	}
}
func Test_StripWhitespace_left(t *testing.T) {
	
	test_StripWhitespace_left(t, "", "")
	test_StripWhitespace_left(t, "blurple", "blurple")

	test_StripWhitespace_left(t, " \t word.\t\t", "word.\t\t")
	test_StripWhitespace_left(t, "b ible . ", "b ible . ")
}
func Test_StripWhitespace_right(t *testing.T) {

	test_StripWhitespace_right(t, "", "")
	test_StripWhitespace_right(t, "glug", "glug")

	test_StripWhitespace_right(t, " \t word.\t\t", " \t word.")
	test_StripWhitespace_right(t, "b ible . ", "b ible .")
}

//these all have amazing flawless coverage ik
