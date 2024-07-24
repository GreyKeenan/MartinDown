package sealeye
import (
	"testing"
	"fmt"
)

func Test_Swimmer(t *testing.T) {
	var cmd = CommandSpec {
		Help: "Look at me, I'm helping!",
	}
	var s Swimmer = NewSwimmer(&cmd)

	s.args = []string { "name", "-harg1", "arg2" }
	status, resp := s.Swim()

	fmt.Printf("%v, %v\n", status, resp.Keyword)
}
