package arrays

import "fmt"

var tablaSlice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var arrayVector [10]int = [10]int{6, 9, 68, 34, 6, 7, 23, 90, 98, 76}

func ShowSlice() {
	fmt.Println(tablaSlice)

	porcion := arrayVector[3:]   // Slice creado desde un vector, desde la posicion 3 hasta el final
	porcion2 := arrayVector[:6]  // Slice creado desde un vector, desde la posicion 0 hasta la 6
	porcion3 := arrayVector[5:6] // Slice creado desde un vector, desde la posicion 0 hasta la 6

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func CapacityAndLength() {

	elements := make([]int, 5, 20)

	fmt.Println("Capacidad del slice:", cap(elements))
	fmt.Println("Longitud del slice:", len(elements))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Println("Capacidad del slice:", cap(nums))
	fmt.Println("Longitud del slice:", len(nums))
}
