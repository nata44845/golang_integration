// Небуферизируемые каналы
package main

import (
	"fmt"
)

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

	respChan := make(chan string)

	go putBook(respChan)
	go deliverBook(respChan)
	burnBook(respChan)

	strChan := make(chan string)
	go func() {
		<-strChan
	}()

	strChan <- "" //Ошибка блокировки deadlock
	close(strChan)

	str, ok := <-strChan
	fmt.Println(str)
	fmt.Println(ok)

	intChan := make(chan int)
	go func() {
		for i := 0; i < 4; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	// for {
	// 	val, ok := <-intChan
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(val) //deadlock
	// }

	for val := range intChan {
		fmt.Println(val)
	}

}
