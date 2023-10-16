package main

import (
	"github.com/sayphp/yaegi/interp"
)

func main() {
	i := interp.New(interp.Opt{})
	i.Eval(`println("Hello")`)
}

// Output:
// Hello
