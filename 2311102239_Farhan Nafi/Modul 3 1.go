package main

import (
	"fmt"
)

func faktor(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktor(n-1)
}

func muta(n, r int) int {
	if r > n {
		return 0
	}
	return faktor(n) / faktor(n-r)
}

func combine(n, r int) int {
	if r > n {
		return 0
	}
	return faktor(n) / (faktor(r) * faktor(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Print("Masukkan nilai a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	pac := muta(a, c)
	cac := combine(a, c)
	pbd := muta(b, d)
	cbd := combine(b, d)

	fmt.Println(pac, cac)
	fmt.Println(pbd, cbd)
}
