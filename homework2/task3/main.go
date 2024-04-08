// Напишите программу, создающую текстовый файл только для чтения, и проверьте, что в него нельзя записать данные.

package main

import (
	"fmt"
	"os"
)

func main() {

	text := "Только для чтения"
	fileName := "file.txt"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Unable to create file:", err)
	}
	defer file.Close()

	file.WriteString(text)
	file.Close()

	if err = os.Chmod(fileName, 0400); err != nil {
		fmt.Println("Ошибка изменения разрешений:", err)
	}

	file, err = os.Open(fileName)
	if err != nil {
		fmt.Println("Unable to create file:", err)
	}
	defer file.Close()

	_, err = file.WriteString("Test")
	if err != nil {
		fmt.Println("Ошибка записи:", err)
	}
}
