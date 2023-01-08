package main

import (
	"fmt"
	luhn "github.com/golang-infrastructure/go-Luhn"
)

func main() {
	b := luhn.Check("8703008427912866")
	fmt.Println(b)
	// Output:
	// true

	b = luhn.Check("8703008427912861")
	fmt.Println(b)
	// Output:
	// false
}
