package main

import (
	"fmt"
)

func main() {
	maxNumber := 10 // Batas maximum
	number := 1     // Angka Awal

	// Simulasi repeat-until
	for finished := false; !finished; {
		fmt.Println("Angka : ", number) // Cetak angka saat ini
		// Cek kondisi keluar
		if number >= maxNumber {
			finished = true // Ubah kondisi untuk keluar dari loop
		}
		number++ // Increment angka
	}
}
