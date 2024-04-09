// Как алгоритмически решить задачу с использованием массива
package main

import (
	"fmt"
	"math"
)

const size = 10

func maximum(input [size]int) (int, int) {
	max := math.MinInt
	max2 := math.MinInt
	for i := 0; i < size; i++ {
		if input[i] > max {
			max2 = max
			max = input[i]
		} else if input[i] > max2 {
			max2 = input[i]
		}
	}
	return max, max2
}

func minimum(input [size]int) (int, int) {
	min := math.MaxInt
	min2 := math.MaxInt
	for i := 0; i < size; i++ {
		if input[i] < min {
			min2 = min
			min = input[i]
		} else if input[i] < min2 {
			min2 = input[i]
		}
	}
	return min, min2
}

// max min 10 20 30 40 50 90 80 70 60 2
func main() {
	a := [...]int{10, 20, 30, 40, 50, 90, 80, 70, 60, 2}

	fmt.Println(minimum(a))
	fmt.Println(maximum(a))
}
