package arrays

import "fmt"

var table [10]int
var matriz [10][30]int

func ShowArrays() {
	table[7] = 33
	table[2] = 54

	fmt.Println(table)

	matriz[5][2] = 1
	matriz[8][24] = 1
	matriz[5][15] = 1

	fmt.Println(matriz)

}
