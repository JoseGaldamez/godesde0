package main

import (
	"fmt"

	"github.com/JoseGaldamez/godesde0/channels"
)

func main() {

	// condicionales.IfCondicionales()
	// condicionales.SwitchCondicionales()

	// number, text := ejecicios.CheckIs100("fffsad")
	// fmt.Println(number, text)

	// keyboard.InsertNumber()

	// iteraciones.ForExample()

	// ejecicios.ShowMultiplyTable()

	// files.AddMultiplyTable()

	// files.ReadMyFile("./files/txt/tables.txt")

	// functions.CallClosure()

	// functions.Exponential(1)

	// arrays.ShowArrays()
	// arrays.CapacityAndLength()

	// maps.ShowMaps()

	// users.AltaUsuario()

	// pedro := new(models.Man)
	// maria := new(models.Woman)

	// HumansBreathing espera un humano, y tanto Pedro como Maria son humanos
	// ejerciciosinterfaces.HumansBreathing(pedro)
	// ejerciciosinterfaces.HumansBreathing(maria)

	// deferPanic.DemoPanic()

	// go gorutines.MySlowName("Jose")

	// fmt.Println("Estoy aqui")
	// var x string
	// fmt.Scanln(&x)

	channel01 := make(chan bool)
	go channels.MySlowNameWithChannel("Jose", channel01)

	state := <-channel01

	if state {
		fmt.Println("Terminó la gorutina")
	}

}
