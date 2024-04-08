// Перепишите задачи 1 и 2, используя пакет ioutil.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	fileName := "file.txt"

	var b bytes.Buffer

	i := 1
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите текст: ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()

		if strings.ToLower(text) == "exit" {
			break
		} else {
			date := time.Now().Format("2006-01-02 15:04:05")
			b.WriteString(fmt.Sprintf("%d %s %s\n", i, date, text))
		}
		i++
	}

	if err := ioutil.WriteFile(fileName, b.Bytes(), 0777); err != nil {
		panic(err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Содержимое файла")
	fmt.Println(string(resultBytes))
}
