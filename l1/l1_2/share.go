package main

var a int

func changeA(b int) {
	a = b
}

func main() {
	for i := 0; i < 10; i++ {
		changeA(i)
	}
}
