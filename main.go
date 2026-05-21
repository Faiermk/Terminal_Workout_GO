package main

import (
	"fmt"
	"tubes/models"
	"tubes/services"
)

func main() {
	var currentUser *models.User

	for currentUser == nil {
		fmt.Println("=== TERMINAL WORKOUT ===")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")

		var pilihAuth int
		fmt.Scan(&pilihAuth)

		switch pilihAuth {
			case 1:
				currentUser = services.Login()
			case 2:
				services.Register()
			case 0: 
				return
		}
	}

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
		currentUser = nil
		for currentUser == nil {
			fmt.Println("\n=== TERMINAL WORKOUT ===")
			fmt.Println("1. Login")
			fmt.Println("2. Registrasi")
			fmt.Println("0. Keluar")
			fmt.Print("Pilih: ")

			var pilihAuth int
			fmt.Scan(&pilihAuth)

			switch pilihAuth {
			case 1:
				currentUser = services.Login()
			case 2:
				services.Register()
			case 0: 
				return
			}
			}
		}
	}
}