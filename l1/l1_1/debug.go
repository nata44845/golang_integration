package main

func doPanic(a int) {
	if a > 1 {
		doPanic(a - 1)
	}
	panic(a)
}

func main() {
	doPanic(5)
}
