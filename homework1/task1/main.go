// Задание 1. Слияние отсортированных массивов
// Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
package main

import (
	"fmt"
)

const n1 = 4
const n2 = 5

func arrayMerge(a [n1]int, b [n2]int) [n1 + n2]int {
	c := [n1 + n2]int{}
	i1 := 0
	i2 := 0
	i := 0
	for i1 < n1 && i2 < n2 {
		if a[i1] < b[i2] {
			c[i] = a[i1]
			i1++
		} else {
			c[i] = b[i2]
			i2++
		}
		i++
	}
	for i1 < n1 {
		c[i] = a[i1]
		i++
		i1++
	}
	for i2 < n2 {
		c[i] = b[i2]
		i++
		i2++
	}
	return c
}

func main() {
	fmt.Println(arrayMerge([n1]int{1, 4, 6, 18}, [n2]int{3, 7, 9, 12, 44}))
}
