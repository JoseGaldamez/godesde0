package condicionales

import (
	"fmt"
	"runtime"
)

func SwitchCondicionales() {
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("El sistema operativo es Windows")
	case "darwin":
		fmt.Println("El sistema operativo es Mac")
	default:
		fmt.Printf("El sistema operativo es %s \n", os)
	}
}
