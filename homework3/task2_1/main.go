package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	msg := make(chan string, 1)
	go func() {
		fmt.Print("Введите число: ")
		for {
			var s string
			fmt.Scan(&s)
			msg <- s
		}
	}()
loop:
	for {
		select {
		case <-ch:
			fmt.Println("\nПрограмма завершает выполнение.")
			break loop
		case s := <-msg:
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("\nОшибка ввода числа:", err)
				continue
			}
			fmt.Println("Число: ", num)
			fmt.Println("Квадрат числа: ", num*num)
		}
	}
}
