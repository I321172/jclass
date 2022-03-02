package class

import (
	"fmt"
	"os"
	"testing"
)

func TestParseClass(t *testing.T) {
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
