package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Print("Masukkan a, b, c, d (dengan syarat a ≥ c dan b ≥ d): ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	Pac := permutation(a, c)
	Cac := combination(a, c)

	Pbd := permutation(b, d)
	Cbd := combination(b, d)

	fmt.Printf("Hasil Permutasi Dan Kombinasi A Terhadap C: %d %d\n", Pac, Cac)
	fmt.Printf("Hasil Permutasi Dan Kombinasi B Terhadap D: %d %d\n", Pbd, Cbd)
}
