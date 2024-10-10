package main

import "fmt"

func main(){
	var beratGram, costGram, total, costKg int

	fmt.Print("Berat Parsel (gram) : ")
	fmt.Scan(&beratGram)

	kg := beratGram / 1000
	gram := beratGram % 1000

	fmt.Print("Detail berat : ", kg ," kg + ", gram , " gr")

	costKg = kg * 10000

	if kg < 10 {
		if gram < 500 {
			costGram = gram * 15
		} else {
			costGram = gram * 5
		}
	}
		
	total = costKg + costGram
	fmt.Print("\nDetail biaya : ", "Rp. ", costKg , " + ", "Rp. ", costGram )
	fmt.Print("\nTotal Biaya : Rp. ", total)
}