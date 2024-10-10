package main

import "fmt"

func main() {
	keberhasilan := true 
	for i := 1; i <= 5; i++ { 
		var warna1, warna2, warna3, warna4 string
		fmt.Print("Percobaan ", i, " : ")
		fmt.Scanln(&warna1, &warna2, &warna3, &warna4)

		if !(warna1 == "merah" && warna2 == "kuning" && warna3 == "hijau" && warna4 == "ungu") {
			keberhasilan = false 
		}
	}
	
	if keberhasilan {
		fmt.Println("KEBERHASILAN: True")
	} else {
		fmt.Println("KEBERHASILAN: False")
	}
}
