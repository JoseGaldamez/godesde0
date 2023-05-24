package functions

// Una funcion que regresa una funcion
func Table(value int) func() int {
	number := value
	secuencia := 0
	return func() int {
		secuencia++
		return number * secuencia
	}
}

func CallClosure() {
	table2 := Table(2)

	for i := 1; i <= 10; i++ {
		println(table2())
	}

}
