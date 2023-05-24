package functions

import "fmt"

func Exponential(value int) {

	if value > 1000000 {
		return
	}

	fmt.Println(value)
	Exponential(value * 2)
}
