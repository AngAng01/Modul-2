package main

import "fmt"

func main() {
	var warna1, warna2, warna3, warna4 string
	var kebenaran string
	kebenaran := "merah kuning hijau ungu"

		for i := 1; i < 6; i ++{
			fmt.Print("Percobaan ", i , " : " )
			fmt.Scanln(&warna1, &warna2, &warna3, &warna4)
		}
		
		hasil := warna1 + " " + warna2 + " " + warna3 + " " + warna4 

		if hasil == kebenaran {
			fmt.Println("KEBERHASILAN : True")
			return
		} else {
			fmt.Println("KEBERHASILAN : False")
		}
	

}