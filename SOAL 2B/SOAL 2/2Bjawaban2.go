//package main

//import "fmt"

//func main(){
	//var n int
	//var bunga,daftarBunga string

	//fmt.Print("N : ")
	//fmt.Scan(&n)

	//for i := 1; i <= n; i++ {
		//fmt.Print("Bunga ",i, ": ")
		//fmt.Scan(&bunga)

		//daftarBunga+= " " + bunga
	//}

	//fmt.Print("Pita :", daftarBunga)
//}


package main

import "fmt"

func main(){
	var inputBunga,daftarBunga string
	var jumlahBunga int
	

	for {

		jumlahBunga++
		fmt.Print("Bunga ", jumlahBunga , ": ")
		fmt.Scan(&inputBunga)

		if inputBunga == "selesai"{
			break
		} 
		daftarBunga+= " " + inputBunga
	}
	fmt.Print("Pita :", daftarBunga)
	fmt.Print("\nBunga : ", jumlahBunga-1)
}