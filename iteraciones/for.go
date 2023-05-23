package iteraciones

import "fmt"

func ForExample() {

	for i := 100; i >= 0; i -= 5 {
		fmt.Println("Hola ", i)
		if i == 50 {
			// se salta la iteracion
			continue
		}

		if i == 25 {
			// se aborta el ciclo
			break
		}
	}

}
