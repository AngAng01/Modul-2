package main
import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukan input string : ")
	fmt.Scanln(&satu)

	fmt.Print("Masukan input string : ")
	fmt.Scanln(&dua)

	fmt.Print("Masukan input string : ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	
	temp = satu // nilai satu sementara disimpan dalam variabel temp
	satu = dua // nilai dua di pindahkan ke satu
	dua = tiga // nilai tiga dipindahkan ke dua
	tiga = temp // nilai temp yang terdapat pada temp di pindah ke tiga

	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}