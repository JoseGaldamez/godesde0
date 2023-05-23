package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var State bool
var Salary float32
var Pi float64
var Age int
var Bithday time.Time

func ShowRest() {
	Name = "Jose"
	State = true
	Salary = 100.5
	Pi = 3.1416
	Age = 26
	Bithday = time.Now()

	fmt.Println("Nombre:", Name, "\nEstado:", State, "\nSalario:", Salary, "\nPi:", Pi, "\nEdad:", Age, "\nFecha:", Bithday)

}

func ConvertToText(number int) (bool, string) {
	text := strconv.Itoa(number)
	return true, text
}
