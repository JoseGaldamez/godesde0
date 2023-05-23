package main

import (
	"fmt"

	"github.com/JoseGaldamez/godesde0/condicionales"
	"github.com/JoseGaldamez/godesde0/ejecicios"
	"github.com/JoseGaldamez/godesde0/iteraciones"
)

func main() {

	condicionales.IfCondicionales()
	condicionales.SwitchCondicionales()

	number, text := ejecicios.CheckIs100("fffsad")
	fmt.Println(number, text)

	//	keyboard.InsertNumber()

	iteraciones.ForExample()

	ejecicios.ShowMultiplyTable()

}
