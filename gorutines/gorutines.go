package gorutines

import (
	"fmt"
	"strings"
	"time"
)

func MySlowName(name string) {
	letters := strings.Split(name, "")

	for _, letter := range letters {
		time.Sleep(1 * time.Second)
		fmt.Println(letter)
	}

}
