package main

import (
	"fmt"
)

func main() {
	var angka int
	total := 0
	fmt.Println("Masukkan angka positif untuk di jumlahkan. Masukkan angka negatif untuk keluar.")

	//Simulasi while loop
	for {
		fmt.Print("Masukkan angka : ")
		fmt.Scanln(&angka)
		if angka < 0 { // keluar dari loop jika angka negatif
			break
		}
		total += angka // tambahkan angka ke total
	}
	fmt.Printf("Total jumlah angka yang dimasukan adalah : %d\n", total)
}
