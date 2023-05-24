package deferPanic

import (
	"fmt"
)

func DemosDefer() {

	fmt.Println("Primer mensaje")

	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Segundo mensaje")
}
