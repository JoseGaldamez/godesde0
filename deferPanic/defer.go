package deferPanic

import (
	"fmt"
	"log"
)

func DemosDefer() {

	fmt.Println("Primer mensaje")

	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Segundo mensaje")
}

func DemoPanic() {

	defer func() {
		reco := recover()
		if reco != nil {
			fmt.Println("Se puede hacer un recover del panic")
			log.Fatalf("Ocurrió un error que generó panic\n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
