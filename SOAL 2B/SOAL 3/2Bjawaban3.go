//package main

//import "fmt"

//func main(){
	//var berat1, berat2 float64

	//for {
		//fmt.Print("Masukan berat belanjaan di kedua kantong : ")
		//fmt.Scan(&berat1, &berat2)

		//if berat1 >= 9 || berat2 >= 9 || berat1 - berat2 < 9 {
			//fmt.Print("Proses Selesai")
			//break
		//}
	
	//}
//}



package main

import "fmt"

func main() {
	var berat1, berat2 float64
	var oleng bool

	for {
		fmt.Print("\nMasukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		if berat1 < 0 || berat2 < 0 {
			fmt.Print("Proses Selesai")
			break
		}

		selisih := berat1 - berat2 
		total := berat1 + berat2

		if total > 150 {
			fmt.Print("Proses Selesai.")
			break
		} else if selisih >= 9 || -selisih >= 9 {
			oleng = true
		} else {
			oleng = false
		}

		fmt.Print("Sepeda motor pak Andi akan oleng: ", oleng)
	}
}
