package functions

import "fmt"

func Calculate() {

	var n3 int = 32
	var n4 int = 243

	plus := func(n1 int, n2 int) int {
		return n1 + n2 + n3 + n4
	}

	fmt.Println(plus(10, 25))
}
