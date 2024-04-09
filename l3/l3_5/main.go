package main

import "fmt"

func getData() <-chan int {
	dataChan := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			dataChan <- i
		}
		close(dataChan)
	}()
}

func main() {
	dataChan := getData()

	// dataChan <- _
}
