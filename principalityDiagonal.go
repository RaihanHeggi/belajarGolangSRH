package main

import "fmt"

func main() {
	var baris, kolom int

	fmt.Print("Silahkan Masukan banyak baris : ")
	fmt.Scanln(&baris)
	fmt.Print("Silahkan Masukan banyak kolom : ")
	fmt.Scanln(&kolom)

	//cek Matriks Persegi atau Bukan
	if baris == kolom {
		//Masukan Nilai ke Array 2D
		matriks := make([][]int, 0, baris)
		for i := 0; i < baris; i++ {
			isiMatriks := make([]int, kolom)
			for j, _ := range isiMatriks {
				fmt.Scan(&isiMatriks[j])
			}
			matriks = append(matriks, isiMatriks)
		}
		//Hitung Nilai Diagonal Utama
		var diagonalUtama int = 0
		for m := 0; m < baris; m++ {
			for n := 0; n < kolom; n++ {
				if m == n {
					diagonalUtama += matriks[m][n]
				}
			}
		}
		fmt.Println("Jumlah Diagonal Matriks Utama : ", diagonalUtama)
	} else {
		fmt.Println("Matriks Tidak Memiliki Diagonal Utama ")
	}
}
