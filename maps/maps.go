package maps

import "fmt"

func ShowMaps() {
	countries := make(map[string]string)

	countries["Mexico"] = "D.F."
	countries["Argentina"] = "Buenos Aires"
	countries["Colombia"] = "Bogota"

	fmt.Println(countries)
	fmt.Println(countries["Mexico"])

	campeonato := map[string]int{
		"Barcelona":   30,
		"Chivas":      59,
		"Real Madrid": 45,
		"Olimpia":     39,
		"River Plate": 36,
	}

	delete(campeonato, "Real Madrid")

	for equipo, puntaje := range campeonato {
		fmt.Printf("> El equipo %s, tiene un puntaje de: %d \n", equipo, puntaje)
	}

	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)

}
