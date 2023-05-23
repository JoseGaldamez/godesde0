package main

import (
	"fmt"

	"github.com/JoseGaldamez/godesde0/condicionales"
	"github.com/JoseGaldamez/godesde0/ejecicios"
)

func main() {

	condicionales.IfCondicionales()
	condicionales.SwitchCondicionales()

	number, text := ejecicios.CheckIs100("500")
	fmt.Println(number, text)

}
