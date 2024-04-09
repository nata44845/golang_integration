// Знакомство с интерфейсом IO Reader
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Ошибка отктытия файла, ", err)
		return
	}
	defer file.Close()

	// buf := make([]byte, 256)
	// if _, err := io.ReadFull(file, buf); err != nil {
	// 	fmt.Println("Не смогли прочитать последовательность байт из файла, ", err)
	// 	return
	// }
	// fmt.Printf("%s\n", buf)

	buf := make([]byte, 256)

	_, err = file.Read(buf)

	if err != nil {
		fmt.Println("Не смогли прочитать последовательность байт из файла, ", err)
		return
	}
	fmt.Printf("%s\n", buf)
}
