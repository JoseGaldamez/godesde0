package main

import (
	"fmt"

	"runtime"

	"github.com/JoseGaldamez/godesde0/ejecicios"
)

func main() {
	if os := runtime.GOOS; os == "windows" || os == "darwin" {
		fmt.Println("El sistema operativo es Windows")
	} else {
		fmt.Println("El sistema operativo no es Windows, es", os)
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("El sistema operativo es Windows")
	case "darwin":
		fmt.Println("El sistema operativo es Mac")
	default:
		fmt.Printf("El sistema operativo es %s \n", os)
	}

	number, text := ejecicios.CheckIs100("500")
	fmt.Println(number, text)

}
