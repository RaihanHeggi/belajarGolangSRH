package main

import "fmt"

func main() {
	var lapar, feeling bool
	var warna string

	fmt.Print("Apakah Kucing Lapar (True/False) : ")
	fmt.Scanln(&lapar)
	fmt.Print("Apakah Kucing Jatuh Cinta (True/False) : ")
	fmt.Scanln(&feeling)
	fmt.Print("Apa Warna Kucing : ")
	fmt.Scanln(&warna)

	if (warna != "Hitam") && (lapar || feeling) {
		fmt.Println("Kucing Mengeong")
	} else {
		fmt.Println("Kucing Membisu")
	}

}
