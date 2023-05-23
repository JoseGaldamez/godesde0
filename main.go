package main

import (
	"fmt"

	"github.com/JoseGaldamez/godesde0/variables"
)

func main() {
	variables.ShowIntegers()
	variables.ShowRest()

	state, text := variables.ConvertToText(16485)
	fmt.Println(state, text)
}
