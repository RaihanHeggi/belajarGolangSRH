package main

import "fmt"

func main() {
	var peluruPenembakA, peluruPenembakB, peluruPenembakC int
	var banyakBalon int

	fmt.Print("Masukkan Jumlah Balon : ")
	fmt.Scanln(&banyakBalon)

	for i := 0; i <= banyakBalon; i++ {
		if banyakBalon - i == 3 {

		}else {
			for j := 0; j <= 4; j++ {
				if j == 0 && i == 0 {
					peluruPenembakA += 1
					peluruPenembakB += 1
					peluruPenembakC += 1
					break
				} else if j == 2 {
					peluruPenembakA += 1
				} else if j == 3 {
					peluruPenembakB += 1
				} else {
					peluruPenembakC += 1
				}
			}	
		}
	}

	fmt.Println(peluruPenembakA)
	fmt.Println(peluruPenembakB)
	fmt.Println(peluruPenembakC)

}
