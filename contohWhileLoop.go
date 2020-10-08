package main

import "fmt"

func main() {
	var lapar bool

	fmt.Println("Apakah Anda Lapar (True/False)")
	fmt.Scanln(&lapar)

	for lapar {
		fmt.Println("Makan")
		fmt.Println("Apakah Anda Lapar (True/False)")
		fmt.Scanln(&lapar)
	}

	fmt.Println("Makan Berhenti")

}
