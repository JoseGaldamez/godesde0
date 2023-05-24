package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MyMiddleare() {
	fmt.Println("Ejecutando Middleware")
	result := operation(sumar)(2, 3)
	fmt.Println(result)

	result = operation(restar)(2, 3)
	fmt.Println(result)

	result = operation(multiplicar)(2, 3)
	fmt.Println(result)
}

func operation(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operación dentro del middleware")
		return f(a, b)
	}
}
