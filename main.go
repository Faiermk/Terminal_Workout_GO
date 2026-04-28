package main

import (
	"fmt"
	"tubes/services"
)

func main() {
	var pilih int

	for {
		fmt.Println("=== APLIKASI WORKOUT ===")
		fmt.Println("1. Tambah")
		fmt.Println("2. Lihat")
		fmt.Println("3. Update")
		fmt.Println("4. Hapus")
		fmt.Println("5. Search Sequential")
		fmt.Println("6. Search Binary")
		fmt.Println("7. Sorting")
		fmt.Println("8. Rekomendasi")
		fmt.Println("9. Laporan")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			services.TambahWorkout()
		case 2:
			services.LihatWorkout()
		case 3:
			services.UpdateWorkout()
		case 4:
			services.HapusWorkout()
		case 5:
			services.SequentialSearch()
		case 6:
			services.BinarySearch()
		case 7:
			services.SelectionSort()
		case 8:
			services.RekomendasiWorkout()
		case 9:
			services.Laporan()
		case 0:
			return
		}
	}
}