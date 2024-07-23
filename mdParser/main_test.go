
package mdParser

import (
	"testing"
)

func test_SLS_do(t *testing.T, ts string, es string) {
	var s string = StripLeadingSpaces(ts)
	if (s != es) {
		t.Fatalf("got '%v' instead of '%v' from '%v'", s, es, ts)
	}
}
func Test_StripLeadingSpaces(t *testing.T) {
	test_SLS_do(t, "    f", " f")
	test_SLS_do(t, "   x", "x")
	test_SLS_do(t, "  9", "9")
	test_SLS_do(t, " g ", "g ")

	test_SLS_do(t, "", "")
	test_SLS_do(t, " ", "")
	test_SLS_do(t, "  ", "")
	test_SLS_do(t, "    ", " ")
}


func test_GH_do(t *testing.T, in string, exLevel int, exText string) {
	var h Header = GetHeader(in)

	if (h.Level != exLevel) {
		t.Fatalf("h.Level is '%v' instead of '%v' from '%v'", h.Level, exLevel, in)
	}
	if (h.Text != exText) {
		t.Fatalf("h.Text is '%v' instead of '%v' from '%v'", h.Text, exText, in)
	}
}
func Test_GetHeader(t *testing.T) {
	test_GH_do(t, "", 0, "")

	test_GH_do(t, "#", 1, "")
	test_GH_do(t, "####", 4, "")

	test_GH_do(t, "###### title", 6, " title")
	test_GH_do(t, "   # x", 1, " x")
	test_GH_do(t, "    # n", 0, "")
	test_GH_do(t, "x# t", 0, "")
	test_GH_do(t, " ##\tl", 2, "\tl")

	//TODO:
	test_GH_do(t, "####### t", 0, "")
}
