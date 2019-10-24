package main

import (
	"fmt"

	"github.com/mepaling/convential-commit-test/mymath"
)

func main() {
	Hello("hello, world!")
	Hello(string(mymath.Add(5 + 2)))
}

// Hello prints the string input
// BUG(rsc): no bugs
func Hello(str string) {
	fmt.Printf("%s\n", str)

}
