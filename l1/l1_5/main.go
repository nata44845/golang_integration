package main

import "fmt"

const size = 5

func maximum(input [size]int) int {
	max := input[0]
	for i := 1; i < size; i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}

func minimum(input [size]int) int {
	min := input[0]
	for i := 1; i < size; i++ {
		if input[i] < min {
			min = input[i]
		}
	}
	return min
}

func wrongCalculation(input [size]int) {
	for i := 0; i < 6; i++ {
		fmt.Println(input[i])
	}
}

// max min 10 49 48 20 39
func main() {
	// var numbers [5]int
	a := [size]int{10, 49, 48, 20, 39}
	fmt.Println(a[0], a[1])

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	fmt.Println("Минимум", minimum(a))
	fmt.Println("Максимум", maximum(a))

	wrongCalculation(a)
}
