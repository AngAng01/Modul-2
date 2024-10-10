package main

import  "fmt"


func main() {
	var k int
	var hasil float64 = 1.0 

	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	for i := 0; i <= k; i++ {
		term := float64((4*i + 2) * (4*i + 2)) / float64((4*i + 1) * (4*i + 3))
		hasil *= term
	}

	fmt.Printf("Nilai f(K) = %.10f\n", hasil)
}
