// Graceful shutdown
// Цель задания:
// Научиться правильно останавливать приложения.
// Что нужно сделать:
// -В работе часто возникает потребность правильно останавливать приложения.
// Например, когда наш сервер обслуживает соединения, а нам хочется, чтобы все текущие соединения были обработаны и
// лишь потом произошло выключение сервиса. Для этого существует паттерн graceful shutdown.
// -Напишите приложение, которое выводит квадраты натуральных чисел на экран,
// а после получения сигнала ^С обрабатывает этот сигнал,
// пишет «выхожу из программы» и выходит.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main1() {
	//var wg sync.WaitGroup
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)

	go func(c chan os.Signal) {
		for {
			v := <-c
			if v == syscall.SIGINT || v == syscall.SIGTERM || v == syscall.SIGKILL {
				fmt.Println("\nПрограмма завершает выполнение.")
				os.Exit(0)
			}
		}
	}(ch)

	for {
		var input string
		fmt.Print("Введите число: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Ошибка ввода с клавиатуры:", err)
			continue
		}
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка ввода числа:", err)
			continue
		}
		square := num * num
		fmt.Printf("%d^2 = %d\n", num, square)
	}
}
