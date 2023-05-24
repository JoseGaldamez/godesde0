package ejerciciosinterfaces

import (
	"fmt"

	"github.com/JoseGaldamez/godesde0/interfaces"
)

func HumansBreathing(human interfaces.Human) {
	human.Breath()
	fmt.Printf("Soy un humano/a %s y estoy respirando\n", human.Gender())
}
