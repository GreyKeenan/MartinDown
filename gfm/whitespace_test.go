
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
