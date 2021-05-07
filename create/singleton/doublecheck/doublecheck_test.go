package doublecheck

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	ans := Get()
	fmt.Printf("%s\n", ans)
}
