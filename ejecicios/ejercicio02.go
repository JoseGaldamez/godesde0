package ejecicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ShowMultiplyTable() {

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

	fmt.Println("\n\n[!] Tabla de multiplicar del número: ", number)
	fmt.Println("====================================")
	fmt.Println()
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", number, i, (number * i))
	}

	fmt.Println("\n\n==== Saliendo del sistema ====")

}
