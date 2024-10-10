package main

import "fmt"

func main() {
	var tahun int
	var kabisat bool

	fmt.Print("Tahun : ")
	fmt.Scan(&tahun)

	if tahun % 400 == 0  || tahun % 4 == 0 && tahun % 100 != 0  {
		kabisat = true
		fmt.Print("Kabisat : ", kabisat)
	} else {
		kabisat = false
		fmt.Print("Kabisat : ", kabisat)
	}	
}
