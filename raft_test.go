package lastwood

import (
	"fmt"
	"testing"
)

func TestRaft_String(t *testing.T) {
	r1 := newRaft()
	fmt.Print(r1.String())
}
