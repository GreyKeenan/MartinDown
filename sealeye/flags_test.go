package sealeye
import (
	"testing"
)

func test_compare(t *testing.T, in string, got string, exp string) {
	if (got != exp) {
		t.Fatalf("Got '%v' from '%v' instead of '%v'", got, in, exp)
	}
}
func test_defaultDeflagger(t *testing.T, in string, exp string) {
	var Deflagger_Default Deflagger = NewDefaultDeflagger()
	switch (Deflagger_Default.IsFlag(in)) {
		case FlagType_Not:
			test_compare(t, in, in, exp)
		case FlagType_Short:
			test_compare(t, in, Deflagger_Default.Deflag_short(in), exp)
		case FlagType_Long:
			test_compare(t, in, Deflagger_Default.Deflag_long(in), exp)
		default:
			t.Fatalf("IsFlag returned unknown FlagType")
	}
}
func Test_defaultDeflagger(t *testing.T) {
	test_defaultDeflagger(t, "", "")
	
	test_defaultDeflagger(t, "--flag", "flag")
	test_defaultDeflagger(t, "-manyshortflags", "manyshortflags")
	test_defaultDeflagger(t, "notaflagatall", "notaflagatall")

	test_defaultDeflagger(t, "--flags-with-hyphens", "flags-with-hyphens")
	test_defaultDeflagger(t, "-flags-with-hyphens", "flags-with-hyphens")
	test_defaultDeflagger(t, "flags-with-hyphens", "flags-with-hyphens")
}
