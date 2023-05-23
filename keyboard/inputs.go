package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number01 int
var number02 int
var leyenda string
var err error

func InsertNumber() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Ingrese numero 1: ")
	if scanner.Scan() {
		number01, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error al convertir el numero")
			panic("El dato ingresado es incorrecto (" + err.Error() + ")")
		}
	}

	fmt.Print("Ingrese numero 2: ")
	if scanner.Scan() {
		number02, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error al convertir el numero")
			panic("El dato ingresado es incorrecto (${err.Error()})")
		}
	}

	fmt.Print("Ingrese leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Printf("El numero 1 es: %d, el numero 2 es: %d, la multiplicación es %d y la leyenda es: %s \n", number01, number02, (number01 * number02), leyenda)

}
