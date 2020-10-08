package main

import "fmt"

func main() {
	var kenyang bool

	fmt.Println("Apakah anda Lapar (True/False)")
	fmt.Scanln(&kenyang)

	for lapar := false; !lapar; {
		fmt.Println("Makan")
		fmt.Println("Apakah anda Lapar (True/False)")
		fmt.Scanln(&kenyang)

		lapar = true
	}
	fmt.Println("Makan Berhenti")
}
