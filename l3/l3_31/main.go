// Паттерн Конвейер
package main

import (
	"fmt"
)

func putBook() chan string {
	firstChan := make(chan string)
	go func() {
		firstChan <- "Складываю книги"
	}()
	return firstChan
}

func deliverBook(firstChan chan string) chan string {
	fmt.Println(<-firstChan)
	secondChan := make(chan string)
	go func() {
		secondChan <- "Везу тележку"
	}()
	return secondChan
}

func burnBook(secondChan chan string) chan string {
	fmt.Println(<-secondChan)
	thirdChan := make(chan string)
	go func() {
		thirdChan <- "Сжигаю книги"
	}()
	return thirdChan
}

func main() {

	fc := putBook()
	sc := deliverBook(fc)
	tc := burnBook(sc)

	fmt.Println(<-tc)
}
