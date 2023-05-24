package users

import (
	"fmt"
	"time"

	"github.com/JoseGaldamez/godesde0/models"
)

func AltaUsuario() {
	fmt.Println("==== Alta de usuario ====")

	user := new(models.User)
	user.AddUser(10, "Jose Galdamez", time.Now(), true)

	fmt.Println(user)

}
