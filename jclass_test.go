package class

import (
	"fmt"
	"os"
	"testing"
)

func TestParseClassJvm7(t *testing.T) {
	f, _ := os.Open("examples/HelloWorld.class")
	defer f.Close()

	c, err := Parse(f)
	if err != nil {
		panic(err)
	}

	f, _ = os.Create("examples/Dumped.class")
	defer f.Close()

	fmt.Println(c.Dump(f))
}

//with Functional interface
func TestParseClassJvm8(t *testing.T) {
	f, _ := os.Open("examples/DumpedSetStatusImpl.class")
	defer f.Close()

	c, err := Parse(f)
	if err != nil {
		panic(err)
	}

	f, _ = os.Create("examples/DumpedSetStatusImpl.class")
	defer f.Close()

	fmt.Println(c.Dump(f))
}
