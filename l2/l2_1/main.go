package main

import (
	"fmt"
	"os"
)

func main() {
	text := "Hello smt"
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Не смогли создать файл, ", err)
		return
	}
	defer file.Close()
	file.WriteString(text)
	fmt.Println(file.Name())
}
