// Пакеты IO и OS. Угадай число
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main1() {
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Не смогли создать файл, ", err)
		return
	}
	defer file.Close()

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(101)
	fmt.Println("Введите число от 1 до 100")
	file.WriteString("Введите число от 1 до 100\n")
	for {
		var answer int
		for {
			_, _ = fmt.Scan(&answer)
			file.WriteString(fmt.Sprintf("Введено число %d\n", answer))
			if answer < 1 || answer > 100 {
				fmt.Println("Число должно быть в диапазоне от 1 до 100")
				file.WriteString("Число должно быть в диапазоне от 1 до 100\n")
			} else {
				break
			}
		}
		if answer == n {
			fmt.Println("Ура! Число угадано.")
			file.WriteString("Ура! Число угадано.\n")
			return
		} else if answer < n {
			fmt.Println(fmt.Sprintf("Загаденное число больше %d", answer))
			file.WriteString(fmt.Sprintf("Загаденное число больше %d\n", answer))
		} else {
			fmt.Println(fmt.Sprintf("Загаденное число меньше %d", answer))
			file.WriteString(fmt.Sprintf("Загаденное число меньше %d\n", answer))
		}
	}
}
