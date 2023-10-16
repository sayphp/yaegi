package main

import (
	"fmt"

	"github.com/sayphp/yaegi/_test/p6"
)

func main() {
	t := p6.IPPrefixSlice{}
	fmt.Println(t)
	b, e := t.MarshalJSON()
	fmt.Println(string(b), e)
}
