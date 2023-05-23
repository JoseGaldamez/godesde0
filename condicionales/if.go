package condicionales

import (
	"fmt"
	"runtime"
)

func IfCondicionales() {

	if os := runtime.GOOS; os == "windows" || os == "darwin" {
		fmt.Println("El sistema operativo es Windows")
	} else {
		fmt.Println("El sistema operativo no es Windows, es", os)
	}
}
