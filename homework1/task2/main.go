// Задание 2. Сортировка пузырьком
// Отсортируйте массив длиной шесть пузырьком.
package main

import (
	"fmt"
)

const size = 6

func bubbleSort(arr [size]int) [size]int {
	finish := false
	for finish == false {
		finish = true
		for i := 0; i < size-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				finish = false
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(bubbleSort([...]int{8, 2, -3, 7, 9, 5}))
}
