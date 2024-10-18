package main

import (
	"fmt"
)

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func fogoh(x int) int {
	return f(g(h(x)))
}

func gohof(x int) int {
	return g(h(f(x)))
}

func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	var a, b, c int
	fmt.Print("Masukkan nilai a, b, c: ")
	fmt.Scan(&a, &b, &c)

	result1 := fogoh(a)
	result2 := gohof(b)
	result3 := hofog(c)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
