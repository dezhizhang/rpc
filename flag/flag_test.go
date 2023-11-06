package flag

import (
	"fmt"
	"os"
	"testing"
)

func TestFlag(t *testing.T) {
	fmt.Println(len(os.Args))

	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v", i, v)
	}
}
