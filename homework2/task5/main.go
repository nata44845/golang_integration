// Напишите программу, которая на вход принимала бы интовое число и для него генерировала бы все возможные комбинации круглых скобок.
package main

import (
	"fmt"
	"strings"
)

func generate_brackets(n int, result []string, x []string, s int) []string {
	if x == nil {
		x = make([]string, n, n)
	}
	if s < 0 {
		return result
	}
	if n == 0 {
		if s == 0 {
			result = append(result, strings.Join(x[:], ""))
		}
	} else {
		x[n-1] = ")"
		result = generate_brackets(n-1, result, x, s+1)
		x[n-1] = "("
		result = generate_brackets(n-1, result, x, s-1)
	}
	return result
}

func main() {
	fmt.Println("Введите число")
	var n int
	fmt.Scan(&n)

	var r []string
	r = generate_brackets(2*n, r, nil, 0)
	str := "["
	n = len(r)
	for i, s := range r {
		str += fmt.Sprintf("\"%s\"", s)
		if i < n-1 {
			str += ", "
		}
	}
	str += "]"
	fmt.Println(str)
}
