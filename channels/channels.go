package channels

import (
	"fmt"
	"strings"
	"time"
)

func MySlowNameWithChannel(name string, ch1 chan bool) {
	letters := strings.Split(name, "")
	for _, letter := range letters {
		time.Sleep(1 * time.Second)
		fmt.Println(letter)
	}

	ch1 <- true
}
