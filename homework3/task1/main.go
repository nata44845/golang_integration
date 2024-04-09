/*Цели задания:
-Научиться работать с каналами и горутинами.
-Понять, как должно происходить общение между потоками.
Что нужно сделать:
-Реализуйте паттерн-конвейер:
Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
-Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
-Произведение: следующая горутина умножает квадрат числа на 2.
-При вводе «стоп» выполнение программы останавливается.
**/

package main

import (
	"fmt"
	"strconv"
)

func Sqr(x int) chan int {
	fmt.Println("Ввод: ", x)
	firstChan := make(chan int)
	go func() {
		firstChan <- x
		firstChan <- (x * x)
	}()
	return firstChan
}

func Mult2(firstChan chan int) chan int {
	x := <-firstChan
	fmt.Println("Квадрат: ", <-firstChan)
	secondChan := make(chan int)
	go func() {
		secondChan <- (x * 2)
	}()
	return secondChan
}

func main() {

	for {
		fmt.Print("Введите число (или 'стоп' для завершения): ")
		var input string
		fmt.Scanln(&input)

		if input == "стоп" || input == "stop" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка ввода числа:", err)
			continue
		}

		fc := Sqr(num)
		sc := Mult2(fc)
		fmt.Println("Произведение на 2:", <-sc)
	}
}
