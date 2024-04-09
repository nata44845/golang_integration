// Горутины. WaitGroup
package main

import (
	"fmt"
	"sync"
)

func putBook(wg *sync.WaitGroup) {
	// time.Sleep(time.Millisecond * 2)
	defer wg.Done()
	fmt.Println("Складываю книги")
}

func deliverBook(wg *sync.WaitGroup) {
	// time.Sleep(time.Millisecond * 4)
	defer wg.Done()
	fmt.Println("Доставляю книги")
}

func burnBook(wg *sync.WaitGroup) {
	// time.Sleep(time.Millisecond * 8)
	fmt.Println("Сжигаю книги")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go putBook(&wg)
	go deliverBook(&wg)

	wg.Wait()
	burnBook(&wg)
}
