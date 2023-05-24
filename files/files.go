package files

import (
	"bufio"
	"fmt"
	"io/ioutil" // this is deprecated
	"os"

	"github.com/JoseGaldamez/godesde0/ejecicios"
)

func SaveMultiplicationTable() {
	text, number := ejecicios.ShowMultiplyTable()
	tableName := "./files/txt/table_" + fmt.Sprint(number) + ".txt"
	file, err := os.Create(tableName)

	if err != nil {
		fmt.Println("Error al crear el archivo", err.Error())
	}

	fmt.Fprintln(file, text)
	file.Close()

}

func AddMultiplyTable() {
	tableName := "./files/txt/tables.txt"
	text, _ := ejecicios.ShowMultiplyTable()

	if !AppendFile(tableName, text) {
		fmt.Println("Error al escribir el archivo")
	}

}

func AppendFile(fileName string, content string) bool {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo", err.Error())
		return false
	}

	_, err = file.WriteString(content)

	if err != nil {
		fmt.Println("Error al escribir en el archivo", err.Error())
		return false
	}

	file.Close()

	return true
}

func ReadMyFile(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo", err.Error())
		return
	}

	fmt.Println(string(file))

	fmt.Println("================== Con Scanner ===================")

	fileScanner, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo con OS", err.Error())
		return
	}

	scanner := bufio.NewScanner(fileScanner)

	for scanner.Scan() {
		// Emprime linea por linea
		fmt.Println(">", scanner.Text())

	}

	fileScanner.Close()

}
