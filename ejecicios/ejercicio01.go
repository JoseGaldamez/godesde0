package ejecicios

import "strconv"

func CheckIs100(stringNumber string) (int, string) {
	number, error := strconv.Atoi(stringNumber)

	if error != nil {
		return 0, "Error al convertir el numero"
	}

	if number > 100 {
		return number, "Es mayor a 100"
	} else if number < 100 {
		return number, "Es menor a 100"
	} else {
		return number, "Es igual a 100"
	}
}
