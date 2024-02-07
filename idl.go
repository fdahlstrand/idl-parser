package main

import (
	"fmt"

	"github.com/fdahlstrand/idl-parser/internal/token"
)

func main() {
	t := token.Token{Type: "Foo", Literal: "Bar"}
	fmt.Println("Hello, world!", t)
}
