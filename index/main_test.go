
package index

import (
	"testing"

	"fmt"

	"github.com/GreyKeenan/pj.ghmd/mdParser"
)

func test_getHeader(s string) (header mdParser.Header) {
	header = mdParser.GetHeader(s)
	if (!header.IsHeader()) {
		panic("GetHeader failed")
	}
	return header
}

func test_IC_WIL(t *testing.T, ic *indexCounter, inH mdParser.Header, ex string, exIc [6]int) {
	var s string = ic.buildIndexLine(inH)
	if (s != ex) {
		t.Fatalf("got '%v' from (%v, %v) instead of '%v'", s, inH.GetLevel(), inH.Text, ex)
		return
	}
	if (ic.levelCounts != exIc) {
		t.Fatalf("(ic) got '%v' from (%v, %v) instead of '%v'", ic.levelCounts, inH.GetLevel(), inH.Text, exIc)
		return
	}
}
func Test_indexCounter_buildIndexLine(t *testing.T) {

	var ic indexCounter

	test_IC_WIL(t,
		&ic,
		test_getHeader("### yay"),
		fmt.Sprintf("\t\t%s0-0-1 yay\n",
		bullet),
		[...]int{0, 0, 1, 0, 0, 0},
	)
	test_IC_WIL(t,
		&ic,
		test_getHeader(" # Single Top Boy"),
		fmt.Sprintf("%s1 Single Top Boy\n", bullet),
		[...]int{1, 0, 0, 0, 0, 0},
	)
	test_IC_WIL(t,
		&ic,
		test_getHeader("  ## little next guy"),
		fmt.Sprintf("\t%s1-1 little next guy\n", bullet),
		[...]int{1, 1, 0, 0, 0, 0},
	)
	test_IC_WIL(t,
		&ic,
		test_getHeader("  ## little next next guy"),
		fmt.Sprintf("\t%s1-2 little next next guy\n", bullet),
		[...]int{1, 2, 0, 0, 0, 0},
	)

}
