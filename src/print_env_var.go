package main

import (
	"fmt"
	"os"
)

func main() {
	some_var := os.Getenv("somevar")
	fmt.Printf("I've read somevar, it holds: %v\n", some_var)
}
