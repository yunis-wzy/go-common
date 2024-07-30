package tools

import (
	"fmt"
	"testing"
)

func TestTwoList (t *testing.T) {
	src := []string{"1", "2", "3", "4", "5"}
	dst := []string{"3", "4", "5", "6", "7"}
	add, del, set := TwoList(src, dst)
	println(src, dst, set)
	fmt.Println(del, set, add)
}