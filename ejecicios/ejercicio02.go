package ejecicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ShowMultiplyTable() (string, int) {

	var text string

	number := 0
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Ingrese un numero (0 para salir): ")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El valor ingresado no es válido.")
				fmt.Println("Intente nuevamente.")
				continue
			}

			break
		}
	}
	if number == 0 {
		os.Exit(0)
	}

	text += fmt.Sprintln("\n\n[!] Tabla de multiplicar del número: ", number)
	text += fmt.Sprintln("====================================")
	fmt.Println()
	for i := 1; i <= 10; i++ {
		text += fmt.Sprintf("%d x %d = %d \n", number, i, (number * i))
	}

	return text, number

}
