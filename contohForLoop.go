package main

import "fmt"

func main() {
	var n int

	fmt.Println("Sudah Berapa Kali Makan ?")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Println("Makan")
	}
	fmt.Println("Makan Berhenti")
}
