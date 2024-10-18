package main

import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2 
		} else {
			n = 3*n + 1 
		}
	}
	fmt.Printf("1\n")
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	if n > 0 && n < 1000000 {
		cetakDeret(n)
	} else {
		fmt.Println("Masukkan bilangan yang lebih kecil dari 1000000.")
	}
}
