// Буферизируемые каналы
package main

import "fmt"

func putBook(rchan chan string) {
	fmt.Println("Складываю книги")
}

func deliverBook(rchan chan string) {
	fmt.Println("Доставляю книги")
}

func burnBook(rchan chan string) {
	fmt.Println("Сжигаю книги")
}

func main() {
	bufChan := make(chan int, 3)

	bufChan <- 1
	bufChan <- 1
	bufChan <- 1

	close(bufChan)

	v, ok := <-bufChan
	fmt.Println(v)
	fmt.Println(ok)

	fmt.Println(len(bufChan), cap(bufChan))
}
