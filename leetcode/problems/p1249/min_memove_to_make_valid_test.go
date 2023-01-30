package p1249

import (
	"fmt"
	"testing"
)

func Test1249_1(t *testing.T) {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
	fmt.Println(minRemoveToMakeValid("a)b(c)d"))
	fmt.Println(minRemoveToMakeValid("))(("))
}
