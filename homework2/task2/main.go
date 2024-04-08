// -Напишите программу, которая читает и выводит в консоль строки из файла, созданного в предыдущей практике, без использования ioutil.
// Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "../Task1/file.txt"

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make([]byte, 32)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
	}
}
