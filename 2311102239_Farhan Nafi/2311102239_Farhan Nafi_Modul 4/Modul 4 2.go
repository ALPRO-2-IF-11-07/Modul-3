package main

import "fmt"
func hitungSkor(waktu [8]int) (jumlahSoal int, skor int) {
	jumlahSoal = 0
	skor = 0
	for i := 0; i < 8; i++ {
		if waktu[i] <= 301 { 
			jumlahSoal++
			skor += waktu[i]
		}
	}
	return
}

func main() {
	var namaPemenang string
	var maxSoal, minSkor int
	maxSoal = 0
	minSkor = 9999 

	for {
		var nama string
		var waktu [8]int

		fmt.Print("Masukkan nama peserta (atau 'Selesai' untuk mengakhiri): ")
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		fmt.Println("Masukkan waktu pengerjaan 8 soal (dalam menit):")
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		jumlahSoal, skor := hitungSkor(waktu)

		if jumlahSoal > maxSoal || (jumlahSoal == maxSoal && skor < minSkor) {
			maxSoal = jumlahSoal
			minSkor = skor
			namaPemenang = nama
		}
	}

	fmt.Printf("Pemenang: %s\n", namaPemenang)
	fmt.Printf("Jumlah soal yang diselesaikan: %d\n", maxSoal)
	fmt.Printf("Total skor: %d\n", minSkor)
}
