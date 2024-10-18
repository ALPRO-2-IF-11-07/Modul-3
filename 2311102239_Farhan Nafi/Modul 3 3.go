package main

import (
	"fmt"
	"math"
)

func Jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

func dalem(x, y, cx, cy, r int) bool {
	return Jarak(x, y, cx, cy) <= float64(r)
}

func main() {
	var x1, y1, r1 int
	fmt.Print("Masukkan koordinat pusat (x1, y1) dan radius r1 dari lingkaran 1: ")
	fmt.Scan(&x1, &y1, &r1)

	var x2, y2, r2 int
	fmt.Print("Masukkan koordinat pusat (x2, y2) dan radius r2 dari lingkaran 2: ")
	fmt.Scan(&x2, &y2, &r2)

	var x, y int
	fmt.Print("Masukkan koordinat titik sembarang (x, y): ")
	fmt.Scan(&x, &y)

	inCircle1 := dalem(x, y, x1, y1, r1)
	inCircle2 := dalem(x, y, x2, y2, r2)

	if inCircle1 && inCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
