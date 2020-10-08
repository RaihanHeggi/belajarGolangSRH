package main

import "fmt"

func main() {
	var kekuatanBenteng, banyakPercobaan int
	var hancur bool = false

	fmt.Print("Berapa Kekuatan Awal Benteng : ")
	fmt.Scanln(&kekuatanBenteng)

	for !hancur {
		if kekuatanBenteng >= 0 {
			kekuatanBenteng -= 3
			if kekuatanBenteng >= 0 {
				fmt.Println("Sisa Kekuatan Benteng : ", kekuatanBenteng)
			}
			banyakPercobaan++
		} else {
			fmt.Println("Gerbang Berhasil Di Dobrak")
			hancur = true
		}
	}
	fmt.Println("Benteng Hancur dalam ", banyakPercobaan, " Kali")
}
